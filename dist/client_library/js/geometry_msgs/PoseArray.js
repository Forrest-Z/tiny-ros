(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.PoseArray=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Pose.js'></script>");

function PoseArray() {
    this.header = std_msgs.Header();
    this.poses = new Array();
};

PoseArray.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var length_poses = this.poses.length;
    buff[offset + 0] = (length_poses >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_poses >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_poses >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_poses >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_poses; i++) {
        offset += this.poses[i].serialize(buff, offset);
    }
    return offset;
};

PoseArray.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_poses = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_poses |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_poses |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_poses |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.poses = new Array(length_poses);
    for (var i = 0; i < length_poses; i++) {
        this.poses[i] = geometry_msgs.Pose();
        offset += this.poses[i].deserialize(buff, offset);
    }
    return offset;
};

PoseArray.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var length_poses = this.poses.length;
    length += 4;
    for (var i = 0; i < length_poses; i++) {
        length += this.poses[i].serializedLength();
    }
    return length;
};

PoseArray.prototype.echo = function() { return ""; };

PoseArray.prototype.getType = function() { return "geometry_msgs/PoseArray"; };

PoseArray.prototype.getMD5 = function() { return "184f43246f3bc9cb5d0613694e6641a6"; };

PoseArray.prototype.getID = function() { return 0; };

PoseArray.prototype.setID = function(id) { };

return function() { return new PoseArray(); };

});
