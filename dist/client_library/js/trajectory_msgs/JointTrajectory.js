(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.trajectory_msgs==="undefined"){g.trajectory_msgs={};}g.trajectory_msgs.JointTrajectory=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/trajectory_msgs/JointTrajectoryPoint.js'></script>");

function JointTrajectory() {
    this.header = std_msgs.Header();
    this.joint_names = new Array();
    this.points = new Array();
};

JointTrajectory.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var length_joint_names = this.joint_names.length;
    buff[offset + 0] = (length_joint_names >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_joint_names >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_joint_names >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_joint_names >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_joint_names; i++) {
        var encoder_joint_namesi = new TextEncoder('utf8');
        var utf8array_joint_namesi = encoder_joint_namesi.encode(this.joint_names[i]);
        buff[offset + 0] = (utf8array_joint_namesi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_joint_namesi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_joint_namesi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_joint_namesi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_joint_namesi.length; i++) {
            buff[offset + i] = utf8array_joint_namesi[i];
        }
        offset += utf8array_joint_namesi.length;
    }
    var length_points = this.points.length;
    buff[offset + 0] = (length_points >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_points >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_points >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_points >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_points; i++) {
        offset += this.points[i].serialize(buff, offset);
    }
    return offset;
};

JointTrajectory.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_joint_names = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_joint_names |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_joint_names |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_joint_names |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.joint_names = new Array(length_joint_names);
    for (var i = 0; i < length_joint_names; i++) {
        var length_joint_namesi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_joint_namesi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_joint_namesi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_joint_namesi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_joint_namesi = new TextDecoder('utf8');
        this.joint_names[i] = decoder_joint_namesi.decode(buff.slice(offset, offset + length_joint_namesi));
        offset += length_joint_namesi;
    }
    var length_points = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_points |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_points |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_points |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.points = new Array(length_points);
    for (var i = 0; i < length_points; i++) {
        this.points[i] = trajectory_msgs.JointTrajectoryPoint();
        offset += this.points[i].deserialize(buff, offset);
    }
    return offset;
};

JointTrajectory.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var length_joint_names = this.joint_names.length;
    length += 4;
    for (var i = 0; i < length_joint_names; i++) {
        var encoder_joint_namesi = new TextEncoder('utf8');
        var utf8array_joint_namesi = encoder_joint_namesi.encode(this.joint_names[i]);
        length += 4;
        length += utf8array_joint_namesi.length;
    }
    var length_points = this.points.length;
    length += 4;
    for (var i = 0; i < length_points; i++) {
        length += this.points[i].serializedLength();
    }
    return length;
};

JointTrajectory.prototype.echo = function() { return ""; };

JointTrajectory.prototype.getType = function() { return "trajectory_msgs/JointTrajectory"; };

JointTrajectory.prototype.getMD5 = function() { return "33e07cd7f5a6df021dad1271b3770d66"; };

JointTrajectory.prototype.getID = function() { return 0; };

JointTrajectory.prototype.setID = function(id) { };

return function() { return new JointTrajectory(); };

});
