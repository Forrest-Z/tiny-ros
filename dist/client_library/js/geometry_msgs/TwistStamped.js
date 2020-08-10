(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.TwistStamped=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Twist.js'></script>");

function TwistStamped() {
    this.header = std_msgs.Header();
    this.twist = geometry_msgs.Twist();
};

TwistStamped.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.twist.serialize(buff, offset);
    return offset;
};

TwistStamped.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.twist.deserialize(buff, offset);
    return offset;
};

TwistStamped.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.twist.serializedLength();
    return length;
};

TwistStamped.prototype.echo = function() { return ""; };

TwistStamped.prototype.getType = function() { return "geometry_msgs/TwistStamped"; };

TwistStamped.prototype.getMD5 = function() { return "2e3e0a57a69306091cb5c65e92d048e1"; };

TwistStamped.prototype.getID = function() { return 0; };

TwistStamped.prototype.setID = function(id) { };

return function() { return new TwistStamped(); };

});
