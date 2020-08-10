(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tinyros==="undefined"){g.tinyros={};}g.tinyros.Subscriber=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros_msgs/TopicInfo.js'></script>");

function Subscriber(topic_name, msg_type, callback) {
    this.id = -1;
    this.cb = callback;
    this.negotiated = false;
    this.srv_flag = false;
    this.msg = msg_type;
    this.topic = topic_name;
    this.endpoint = tinyros_msgs.TopicInfo().ID_SUBSCRIBER;
};

Subscriber.prototype.callback = function(data) {
    var msg = this.msg();
    msg.deserialize(data, 0);
    this.cb(msg);
};

Subscriber.prototype.getMsgType = function() {
    return this.msg().getType();
};

Subscriber.prototype.getMsgMD5 = function() {
    return this.msg().getMD5();
};

Subscriber.prototype.getEndpointType = function() {
    return this.endpoint;
};

Subscriber.prototype.negotiated = function() {
    return this.negotiated;
};

return function(_1,_2,_3) { return new Subscriber(_1,_2,_3); };

});

