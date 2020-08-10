(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.Path=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/PoseStamped.js'></script>");

function Path() {
    this.header = std_msgs.Header();
    this.poses = new Array();
};

Path.prototype.serialize = function(buff, idx) {
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

Path.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_poses = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_poses |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_poses |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_poses |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.poses = new Array(length_poses);
    for (var i = 0; i < length_poses; i++) {
        this.poses[i] = geometry_msgs.PoseStamped();
        offset += this.poses[i].deserialize(buff, offset);
    }
    return offset;
};

Path.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var length_poses = this.poses.length;
    length += 4;
    for (var i = 0; i < length_poses; i++) {
        length += this.poses[i].serializedLength();
    }
    return length;
};

Path.prototype.echo = function() { return ""; };

Path.prototype.getType = function() { return "nav_msgs/Path"; };

Path.prototype.getMD5 = function() { return "4a185240c929c496a7e0d6202e3c89af"; };

Path.prototype.getID = function() { return 0; };

Path.prototype.setID = function(id) { };

return function() { return new Path(); };

});
