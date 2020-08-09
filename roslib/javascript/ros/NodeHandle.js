(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.ros==="undefined"){g.ros={};}g.ros.NodeHandle=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/ros/Hardware.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/String.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros_msgs/TopicInfo.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros_msgs/Log.js'></script>");

const MODE_FIRST_FF       = 0;
const MODE_PROTOCOL_VER   = 1;
const PROTOCOL_VER        = 0xb9;
const MODE_SIZE_L         = 2;
const MODE_SIZE_L1        = 3;
const MODE_SIZE_H         = 4;
const MODE_SIZE_H1        = 5;
const MODE_SIZE_CHECKSUM  = 6;
const MODE_TOPIC_L        = 7;
const MODE_TOPIC_L1       = 8;
const MODE_TOPIC_H        = 9;
const MODE_TOPIC_H1       = 10;
const MODE_MESSAGE        = 11;
const MODE_MSG_CHECKSUM   = 12;

const INPUT_SIZE  = (64 * 1024);
const OUTPUT_SIZE = (64 * 1024);

const MAX_SUBSCRIBERS = 100;
const MAX_PUBLISHERS  = 100;

const WS_PORT_NUM = "11318";

function NodeHandle() {
    this.ip_addr = "127.0.0.1";
    this.node_name = "";
    this.hardware = null;
    this.ssl = false;
    this.connected = false;
    this.reconnect = false;
    this.publishers = new Map();
    this.subscribers = new Map();
};

NodeHandle.prototype.onopen = function(evt) {
    var msg = std_msgs.String();
    msg.data = this.nh.node_name;
    this.nh.publish(tinyros_msgs.TopicInfo().ID_SESSION_ID, msg);
    this.nh.negotiateTopics();
    this.nh.connected = true;
};

NodeHandle.prototype.onmessage = function(evt) {
    var message_in = new Uint8Array(evt.data);
    if (message_in.length >= 12) {
        var topic = 0, bytes = 0, index= 0, checksum = 0;
        do {
            index = 0;
            if (message_in[index++] !== 0xff) {
                break;
            }

            if (message_in[index++] !== 0xb9) {
                break;
            }

            bytes = message_in[index];
            bytes += message_in[index + 1] << 8;
            bytes += message_in[index + 2] << 16;
            bytes += message_in[index + 3] << 24;
            checksum = message_in[index];
            checksum += message_in[index + 1];
            checksum += message_in[index + 2];
            checksum += message_in[index + 3];
            checksum += message_in[index + 4];
            index += 5;
          
            if((checksum % 256) !== 255) {
                break;
            }

            topic = message_in[index];
            topic += message_in[index + 1] << 8;
            topic += message_in[index + 2] << 16;
            topic += message_in[index + 3] << 24;
            checksum = message_in[index];
            checksum += message_in[index + 1];
            checksum += message_in[index + 2];
            checksum += message_in[index + 3];
            index += 4;
            
            if (message_in.length < (index + bytes + 1)) {
                break;
            }
            
            if(bytes > 0) {
                for (var i=0; i < bytes + 1; i++) {
                    checksum += message_in[index + i];
                }
            } else {
                checksum += message_in[index];
            }

            if ((checksum % 256) === 255) {
                var TopicInfo = tinyros_msgs.TopicInfo();
                if (topic === TopicInfo.ID_PUBLISHER) {
                    this.nh.negotiateTopics();
                } else {
                    if(this.nh.subscribers.has(topic)) {
                        this.nh.subscribers.get(topic).callback(message_in, index);
                    }
                }
            }
        } while(0);
    }
};

NodeHandle.prototype.onerror = function(evt) {
    console.log("WebSocket onerror!");
    this.nh.connected = false;
    if (this.nh.reconnect) return;
    this.nh.reconnect = true;
    setTimeout(function() {
        if (this.nh.ssl) {
            this.nh.hardware = new WebSocket("wss://" + this.nh.ip_addr + ":" + WS_PORT_NUM);
            this.nh.hardware.nh = this.nh;
        } else {
            this.nh.hardware = new WebSocket("ws://" + this.nh.ip_addr + ":" + WS_PORT_NUM);
            this.nh.hardware.nh = this.nh;
        }
        this.nh.hardware.binaryType = "arraybuffer";
        this.nh.hardware.onopen = this.nh.onopen;
        this.nh.hardware.onmessage = this.nh.onmessage;
        this.nh.hardware.onclose = this.nh.onclose;
        this.nh.hardware.onerror = this.nh.onerror;
        this.nh.reconnect = false;
    }, 1000);
};

NodeHandle.prototype.onclose = function(evt) {
    console.log("WebSocket colosed->" + evt.code + ":" + evt.reason);
    this.nh.connected = false;
    if (this.nh.reconnect) return;
    this.nh.reconnect = true;
    setTimeout(function() {
        if (this.nh.ssl) {
            this.nh.hardware = new WebSocket("wss://" + this.nh.ip_addr + ":" + WS_PORT_NUM);
            this.nh.hardware.nh = this.nh;
        } else {
            this.nh.hardware = new WebSocket("ws://" + this.nh.ip_addr + ":" + WS_PORT_NUM);
            this.nh.hardware.nh = this.nh;
        }
        this.nh.hardware.binaryType = "arraybuffer";
        this.nh.hardware.onopen = this.nh.onopen;
        this.nh.hardware.onmessage = this.nh.onmessage;
        this.nh.hardware.onclose = this.nh.onclose;
        this.nh.hardware.onerror = this.nh.onerror;
        this.nh.reconnect = false;
    }, 1000);
};


NodeHandle.prototype.initNode = function(node_name, ip_addr, ssl) {
    this.node_name = node_name;
    
    if (typeof ssl !== 'undefined') {
        this.ssl = ssl;
    }

    if (typeof ip_addr !== 'undefined') {
        this.ip_addr = ip_addr;
    }
    
    if (this.ssl) {
        this.hardware = new WebSocket("wss://" + this.ip_addr + ":" + WS_PORT_NUM);
        this.hardware.nh = this;
    } else {
        this.hardware = new WebSocket("ws://" + this.ip_addr + ":" + WS_PORT_NUM);
        this.hardware.nh = this;
    }
    
    this.hardware.binaryType = "arraybuffer";
    this.hardware.onopen = this.onopen;
    this.hardware.onmessage = this.onmessage;
    this.hardware.onclose = this.onclose;
    this.hardware.onerror = this.onerror;
};

NodeHandle.prototype.publish = function(id, msg) {
    if (!this.ok()) {
        console.log("NodeHandle.publish: !ok()");
        return -1;
    }
    
    var length = msg.serializedLength() + 12;
    if (length > OUTPUT_SIZE) {
        console.log("NodeHandle.publish: overflow!");
        return -1;
    }
    
    var buff = new Uint8Array(length);
    var l = msg.serialize(buff, 11);
    buff[0] = 0xFF;
    buff[1] = PROTOCOL_VER;
    buff[2] = ((l-11) >> 0 ) & 0xFF;
    buff[3] = ((l-11) >> 8 ) & 0xFF;
    buff[4] = ((l-11) >> 16) & 0xFF;
    buff[5] = ((l-11) >> 24) & 0xFF;
    buff[6] = (255 - ((buff[2] + buff[3] + buff[4] + buff[5]) % 256)) & 0xFF;
    buff[7] = (id >> 0 ) & 0xFF;
    buff[8] = (id >> 8 ) & 0xFF;
    buff[9] = (id >> 16) & 0xFF;
    buff[10] = (id >> 24) & 0xFF;

    var chk = 0;
    for (var i = 7; i < l; i++) {
        chk += buff[i];
    }
    
    buff[l] = (255 - (chk % 256)) & 0xFF;

    this.hardware.send(buff);
    
    return 0;
};

NodeHandle.prototype.advertise = function(p) {
    if (this.publishers.size < MAX_PUBLISHERS) {
        p.id = this.publishers.size + 100 + MAX_SUBSCRIBERS;
        p.nh = this;
        this.publishers.set(p.id , p);
        this.negotiateTopicsP(p);
        return true;
    }

    return false;
};

NodeHandle.prototype.subscribe = function(s) {
    if (this.subscribers.size < MAX_SUBSCRIBERS) {
        s.id = this.subscribers.size + 100;
        this.subscribers.set(s.id , s);
        this.negotiateTopicsS(s);
        return true;
    }
    return false;
};

NodeHandle.prototype.negotiateTopicsP = function(p) {
    var ti = tinyros_msgs.TopicInfo();
    ti.topic_id = p.id;
    ti.topic_name = p.topic;
    ti.message_type = p.getMsgType();
    ti.md5sum = p.getMsgMD5();
    ti.buffer_size = OUTPUT_SIZE;
    ti.node = this.node_name;
    this.publish(p.endpoint, ti);
};

NodeHandle.prototype.negotiateTopicsS = function(s) {
    var ti = tinyros_msgs.TopicInfo();
    ti.topic_id = s.id;
    ti.topic_name = s.topic;
    ti.message_type = s.getMsgType();
    ti.md5sum = s.getMsgMD5();
    ti.buffer_size = INPUT_SIZE;
    ti.node = this.node_name;
    this.publish(s.endpoint, ti);
};

NodeHandle.prototype.negotiateTopics = function() {
    for (var value of this.publishers.values()) {
        this.negotiateTopicsP(value);
    }
    for (var value of this.subscribers.values()) {
        this.negotiateTopicsS(value);
    }
};

NodeHandle.prototype.ok = function() {
    return this.connected;
};

NodeHandle.prototype.log = function(level, msg) {
    var l = tinyros_msgs.Log();
    l.level = level;
    l.msg = "[" + this.node_name + "] " + msg;
    this.publish(tinyros_msgs.TopicInfo().ID_LOG, l);
};

return function() { return new NodeHandle(); };

});

