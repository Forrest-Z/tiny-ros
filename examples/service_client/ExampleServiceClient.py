import os
import sys
import time
print('Tinyros modules: %s' % os.path.abspath(os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))))
sys.path.append(os.path.abspath(os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))))
import tinyros
import tinyros_hello.srv.Test

def main():
    tinyros.init("PyExampleServiceClient", "127.0.0.1")
    client = tinyros.ServiceClient("test_srv", tinyros_hello.srv.Test.Request, tinyros_hello.srv.Test.Response)
    tinyros.nh().serviceClient(client)
    while True:
        req = tinyros_hello.srv.Test.Request()
        res = tinyros_hello.srv.Test.Response()
        req.input = "hello world!"
        if client.call(req, res):
            print('Service responsed with "%s"' % res.output)
        else:
            print("Service call failed.")
        time.sleep(1)

if __name__ == '__main__':
    main()

