(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tinyros==="undefined"){g.tinyros={};}g.tinyros.Publisher=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros_msgs/TopicInfo.js'></script>");

function Publisher(topic_name, msg_type) {
    this.id = -1;
    this.negotiated = false;
    this.nh = null;
    this.msg = msg_type;
    this.topic = topic_name;
    this.endpoint = tinyros_msgs.TopicInfo().ID_PUBLISHER;
};

Publisher.prototype.publish = function(msg) {
    if (this.nh !== null) {
        return this.nh.publish(this.id, msg);
    } else {
        return -1;
    }
};
  
Publisher.prototype.getMsgType = function() {
    return this.msg().getType();
};

Publisher.prototype.getMsgMD5 = function() {
    return this.msg().getMD5();
};

Publisher.prototype.getEndpointType = function() {
    return this.endpoint;
};

Publisher.prototype.negotiated = function() {
    return this.negotiated;
};

return function(_1,_2) { return new Publisher(_1,_2); };

});





