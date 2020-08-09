#include "Server.h"
#include "HTTPSocket.h"
#include "WebSocket.h"
#include "Extensions.h"
#include "Parser.h"

#include <cstring>
#include <openssl/sha.h>
#include <openssl/ssl.h>

void base64(unsigned char *src, char *dst)
{
    static const char *b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
    for (int i = 0; i < 18; i += 3) {
        *dst++ = b64[(src[i] >> 2) & 63];
        *dst++ = b64[((src[i] & 3) << 4) | ((src[i + 1] & 240) >> 4)];
        *dst++ = b64[((src[i + 1] & 15) << 2) | ((src[i + 2] & 192) >> 6)];
        *dst++ = b64[src[i + 2] & 63];
    }
    *dst++ = b64[(src[18] >> 2) & 63];
    *dst++ = b64[((src[18] & 3) << 4) | ((src[19] & 240) >> 4)];
    *dst++ = b64[((src[19] & 15) << 2)];
    *dst++ = '=';
}

namespace uWS {

void Server::acceptHandler(uv_poll_t *p, int status, int events)
{
    if (status < 0) {
        return;
    }

    Server *server = (Server *) p->data;

    socklen_t listenAddrLength = sizeof(sockaddr_in);
    uv_os_fd_t serverFd;
    uv_fileno((uv_handle_t *) p, &serverFd);
    uv_os_fd_t clientFd = accept(serverFd, (sockaddr *) &server->listenAddr, &listenAddrLength);
    if (clientFd == -1) {
        return;
    }

#ifdef __APPLE__
    int noSigpipe = 1;
    setsockopt(clientFd, SOL_SOCKET, SO_NOSIGPIPE, &noSigpipe, sizeof(int));
#endif

    new HTTPSocket(clientFd, server);
}

void Server::upgradeHandler(Server *server)
{
    server->upgradeQueueMutex.lock();

    while (!server->upgradeQueue.empty()) {
        UpgradeRequest upgradeRequest = server->upgradeQueue.front();
        server->upgradeQueue.pop();

        unsigned char shaInput[] = "XXXXXXXXXXXXXXXXXXXXXXXX258EAFA5-E914-47DA-95CA-C5AB0DC85B11";
        memcpy(shaInput, upgradeRequest.sslKey.c_str(), 24);
        unsigned char shaDigest[SHA_DIGEST_LENGTH];
        SHA1(shaInput, sizeof(shaInput) - 1, shaDigest);

        memcpy(server->upgradeBuffer, "HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: ", 97);
        base64(shaDigest, server->upgradeBuffer + 97);
        memcpy(server->upgradeBuffer + 125, "\r\n", 2);
        size_t upgradeResponseLength = 127;

        // Note: This could be moved into Extensions.cpp as a "decorator" if we get more complex extension support
        PerMessageDeflate *perMessageDeflate = nullptr;
        ExtensionsParser extensionsParser(upgradeRequest.extensions.c_str());
        if ((server->options & PERMESSAGE_DEFLATE) && extensionsParser.perMessageDeflate) {
            std::string response;
            perMessageDeflate = new PerMessageDeflate(extensionsParser, server->options, response);
            response.append("\r\n\r\n");
            memcpy(server->upgradeBuffer + 127, response.data(), response.length());
            upgradeResponseLength += response.length();
        } else {
            memcpy(server->upgradeBuffer + 127, "\r\n", 2);
            upgradeResponseLength += 2;
        }

        uv_poll_t *clientPoll = new uv_poll_t;
        WebSocket webSocket(clientPoll);
        webSocket.initPoll(server, upgradeRequest.fd, upgradeRequest.ssl, perMessageDeflate);
        webSocket.write(server->upgradeBuffer, upgradeResponseLength, false);

        if (server->clients) {
            webSocket.link(server->clients);
        }
        server->clients = clientPoll;
        server->connectionCallback(webSocket);
    }

    server->upgradeQueueMutex.unlock();
}

void Server::closeHandler(Server *server)
{
    if (!server->master) {
        uv_close((uv_handle_t *) &server->upgradeAsync, [](uv_handle_t *a) {});
        uv_close((uv_handle_t *) &server->closeAsync, [](uv_handle_t *a) {});
    }

    if (server->listenPoll) {
        uv_os_fd_t listenFd;
        uv_fileno((uv_handle_t *) server->listenPoll, &listenFd);
        ::close(listenFd);
        uv_poll_stop(server->listenPoll);
        uv_close((uv_handle_t *) server->listenPoll, [](uv_handle_t *handle) {
            delete (uv_poll_t *) handle;
        });
    }

    for (WebSocket webSocket = server->clients; webSocket; webSocket = webSocket.next()) {
        webSocket.close(server->forceClose);
    }
}

Server::Server(int port, bool master, int options, int maxPayload) : master(master), options(options), maxPayload(maxPayload)
{
    loop = master ? uv_default_loop() : uv_loop_new();

    recvBuffer = new char[LARGE_BUFFER_SIZE + Parser::CONSUME_POST_PADDING];
    upgradeBuffer = new char[LARGE_BUFFER_SIZE];
    inflateBuffer = new char[LARGE_BUFFER_SIZE];
    sendBuffer = new char[SHORT_BUFFER_SIZE];

    onConnection([](WebSocket webSocket) {});
    onDisconnection([](WebSocket webSocket, int code, char *message, size_t length) {});
    onMessage([](WebSocket webSocket, char *message, size_t length, OpCode opCode) {});
    onUpgrade([this](uv_os_fd_t fd, const char *secKey, void *ssl, const char *extensions, size_t extensionsLength) {
        upgrade(fd, secKey, ssl, extensions, extensionsLength);
    });

    if (port) {
        uv_os_fd_t listenFd = socket(AF_INET, SOCK_STREAM, 0);
        listenAddr.sin_family = AF_INET;
        listenAddr.sin_addr.s_addr = INADDR_ANY;
        listenAddr.sin_port = htons(port);

        listenPoll = new uv_poll_t;
        listenPoll->data = this;
    
        int opt = 1;
        setsockopt(listenFd, SOL_SOCKET, SO_REUSEADDR, (const char *)&opt, sizeof(opt));

        if (bind(listenFd, (sockaddr *) &listenAddr, sizeof(sockaddr_in)) | listen(listenFd, 10)) {
            throw ERR_LISTEN;
        }

        uv_poll_init_socket(loop, listenPoll, listenFd);
        uv_poll_start(listenPoll, UV_READABLE, acceptHandler);
    }

    if (!master) {
        upgradeAsync.data = this;
        closeAsync.data = this;

        uv_async_init(loop, &closeAsync, [](uv_async_t *a) {
            closeHandler((Server *) a->data);
        });

        uv_async_init(loop, &upgradeAsync, [](uv_async_t *a) {
            upgradeHandler((Server *) a->data);
        });
    }
}

Server::~Server()
{
    delete [] recvBuffer;
    delete [] upgradeBuffer;
    delete [] sendBuffer;
    delete [] inflateBuffer;

    if (!master) {
        uv_loop_delete(loop);
    }
}

void Server::onUpgrade(std::function<void (uv_os_fd_t, const char *, void *, const char *, size_t)> upgradeCallback)
{
    this->upgradeCallback = upgradeCallback;
}

void Server::onConnection(std::function<void (WebSocket)> connectionCallback)
{
    this->connectionCallback = connectionCallback;
}

void Server::onDisconnection(std::function<void (WebSocket, int, char *, size_t)> disconnectionCallback)
{
    this->disconnectionCallback = disconnectionCallback;
}

void Server::onMessage(std::function<void (WebSocket, char *, size_t, OpCode)> messageCallback)
{
    this->messageCallback = messageCallback;
}

void Server::close(bool force)
{
    forceClose = force;
    if (master) {
        closeHandler(this);
    } else {
        uv_async_send(&closeAsync);
    }
}

void Server::upgrade(uv_os_fd_t fd, const char *secKey, void *ssl, const char *extensions, size_t extensionsLength)
{
    upgradeQueueMutex.lock();
    upgradeQueue.push({fd, std::string(secKey, 24), ssl, std::string(extensions, extensionsLength)});
    upgradeQueueMutex.unlock();

    if (master) {
        upgradeHandler(this);
    } else {
        uv_async_send(&upgradeAsync);
    }
}

void Server::broadcast(char *data, size_t length, OpCode opCode)
{
    for (WebSocket webSocket = clients; webSocket; webSocket = webSocket.next()) {
        webSocket.send(data, length, opCode);
    }
}

void Server::run()
{
    uv_run(loop, UV_RUN_DEFAULT);
}

}
