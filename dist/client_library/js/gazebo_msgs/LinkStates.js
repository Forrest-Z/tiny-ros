(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.LinkStates=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Pose.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Twist.js'></script>");

function LinkStates() {
    this.name = new Array();
    this.pose = new Array();
    this.twist = new Array();
};

LinkStates.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var length_name = this.name.length;
    buff[offset + 0] = (length_name >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_name >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_name >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_name >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_name; i++) {
        var encoder_namei = new TextEncoder('utf8');
        var utf8array_namei = encoder_namei.encode(this.name[i]);
        buff[offset + 0] = (utf8array_namei.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_namei.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_namei.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_namei.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_namei.length; i++) {
            buff[offset + i] = utf8array_namei[i];
        }
        offset += utf8array_namei.length;
    }
    var length_pose = this.pose.length;
    buff[offset + 0] = (length_pose >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_pose >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_pose >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_pose >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_pose; i++) {
        offset += this.pose[i].serialize(buff, offset);
    }
    var length_twist = this.twist.length;
    buff[offset + 0] = (length_twist >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_twist >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_twist >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_twist >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_twist; i++) {
        offset += this.twist[i].serialize(buff, offset);
    }
    return offset;
};

LinkStates.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_name |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.name = new Array(length_name);
    for (var i = 0; i < length_name; i++) {
        var length_namei = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_namei |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_namei |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_namei |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_namei = new TextDecoder('utf8');
        this.name[i] = decoder_namei.decode(buff.slice(offset, offset + length_namei));
        offset += length_namei;
    }
    var length_pose = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_pose |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_pose |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_pose |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.pose = new Array(length_pose);
    for (var i = 0; i < length_pose; i++) {
        this.pose[i] = geometry_msgs.Pose();
        offset += this.pose[i].deserialize(buff, offset);
    }
    var length_twist = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_twist |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_twist |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_twist |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.twist = new Array(length_twist);
    for (var i = 0; i < length_twist; i++) {
        this.twist[i] = geometry_msgs.Twist();
        offset += this.twist[i].deserialize(buff, offset);
    }
    return offset;
};

LinkStates.prototype.serializedLength = function() {
    var length = 0;
    var length_name = this.name.length;
    length += 4;
    for (var i = 0; i < length_name; i++) {
        var encoder_namei = new TextEncoder('utf8');
        var utf8array_namei = encoder_namei.encode(this.name[i]);
        length += 4;
        length += utf8array_namei.length;
    }
    var length_pose = this.pose.length;
    length += 4;
    for (var i = 0; i < length_pose; i++) {
        length += this.pose[i].serializedLength();
    }
    var length_twist = this.twist.length;
    length += 4;
    for (var i = 0; i < length_twist; i++) {
        length += this.twist[i].serializedLength();
    }
    return length;
};

LinkStates.prototype.echo = function() { return ""; };

LinkStates.prototype.getType = function() { return "gazebo_msgs/LinkStates"; };

LinkStates.prototype.getMD5 = function() { return "a6f8cc7b3dee31015716313fe2d419eb"; };

LinkStates.prototype.getID = function() { return 0; };

LinkStates.prototype.setID = function(id) { };

return function() { return new LinkStates(); };

});
