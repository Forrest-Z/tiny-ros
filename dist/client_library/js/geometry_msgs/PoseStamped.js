(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.PoseStamped=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Pose.js'></script>");

function PoseStamped() {
    this.header = std_msgs.Header();
    this.pose = geometry_msgs.Pose();
};

PoseStamped.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.pose.serialize(buff, offset);
    return offset;
};

PoseStamped.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.pose.deserialize(buff, offset);
    return offset;
};

PoseStamped.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.pose.serializedLength();
    return length;
};

PoseStamped.prototype.echo = function() { return ""; };

PoseStamped.prototype.getType = function() { return "geometry_msgs/PoseStamped"; };

PoseStamped.prototype.getMD5 = function() { return "c7084e6b27c3d6e62efd9bf6d2f6540f"; };

PoseStamped.prototype.getID = function() { return 0; };

PoseStamped.prototype.setID = function(id) { };

return function() { return new PoseStamped(); };

});
