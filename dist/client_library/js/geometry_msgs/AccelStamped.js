(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.AccelStamped=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Accel.js'></script>");

function AccelStamped() {
    this.header = std_msgs.Header();
    this.accel = geometry_msgs.Accel();
};

AccelStamped.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.accel.serialize(buff, offset);
    return offset;
};

AccelStamped.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.accel.deserialize(buff, offset);
    return offset;
};

AccelStamped.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.accel.serializedLength();
    return length;
};

AccelStamped.prototype.echo = function() { return ""; };

AccelStamped.prototype.getType = function() { return "geometry_msgs/AccelStamped"; };

AccelStamped.prototype.getMD5 = function() { return "fa35432963826361a1073b1df905a559"; };

AccelStamped.prototype.getID = function() { return 0; };

AccelStamped.prototype.setID = function(id) { };

return function() { return new AccelStamped(); };

});
