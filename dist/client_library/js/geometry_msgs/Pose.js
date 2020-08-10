(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.Pose=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Point.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Quaternion.js'></script>");

function Pose() {
    this.position = geometry_msgs.Point();
    this.orientation = geometry_msgs.Quaternion();
};

Pose.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.position.serialize(buff, offset);
    offset += this.orientation.serialize(buff, offset);
    return offset;
};

Pose.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.position.deserialize(buff, offset);
    offset += this.orientation.deserialize(buff, offset);
    return offset;
};

Pose.prototype.serializedLength = function() {
    var length = 0;
    length += this.position.serializedLength();
    length += this.orientation.serializedLength();
    return length;
};

Pose.prototype.echo = function() { return ""; };

Pose.prototype.getType = function() { return "geometry_msgs/Pose"; };

Pose.prototype.getMD5 = function() { return "0b42fb88be8cac0efa6e446e13befcae"; };

Pose.prototype.getID = function() { return 0; };

Pose.prototype.setID = function(id) { };

return function() { return new Pose(); };

});
