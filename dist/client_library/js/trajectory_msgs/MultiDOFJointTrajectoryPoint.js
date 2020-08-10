(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.trajectory_msgs==="undefined"){g.trajectory_msgs={};}g.trajectory_msgs.MultiDOFJointTrajectoryPoint=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Transform.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Twist.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Duration.js'></script>");

function MultiDOFJointTrajectoryPoint() {
    this.transforms = new Array();
    this.velocities = new Array();
    this.accelerations = new Array();
    this.time_from_start = tinyros.Duration();
};

MultiDOFJointTrajectoryPoint.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var length_transforms = this.transforms.length;
    buff[offset + 0] = (length_transforms >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_transforms >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_transforms >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_transforms >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_transforms; i++) {
        offset += this.transforms[i].serialize(buff, offset);
    }
    var length_velocities = this.velocities.length;
    buff[offset + 0] = (length_velocities >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_velocities >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_velocities >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_velocities >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_velocities; i++) {
        offset += this.velocities[i].serialize(buff, offset);
    }
    var length_accelerations = this.accelerations.length;
    buff[offset + 0] = (length_accelerations >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_accelerations >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_accelerations >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_accelerations >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_accelerations; i++) {
        offset += this.accelerations[i].serialize(buff, offset);
    }
    buff[offset + 0] = ((+this.time_from_start.sec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.time_from_start.sec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.time_from_start.sec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.time_from_start.sec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.time_from_start.nsec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.time_from_start.nsec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.time_from_start.nsec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.time_from_start.nsec) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

MultiDOFJointTrajectoryPoint.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_transforms = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_transforms |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_transforms |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_transforms |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.transforms = new Array(length_transforms);
    for (var i = 0; i < length_transforms; i++) {
        this.transforms[i] = geometry_msgs.Transform();
        offset += this.transforms[i].deserialize(buff, offset);
    }
    var length_velocities = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_velocities |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_velocities |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_velocities |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.velocities = new Array(length_velocities);
    for (var i = 0; i < length_velocities; i++) {
        this.velocities[i] = geometry_msgs.Twist();
        offset += this.velocities[i].deserialize(buff, offset);
    }
    var length_accelerations = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_accelerations |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_accelerations |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_accelerations |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.accelerations = new Array(length_accelerations);
    for (var i = 0; i < length_accelerations; i++) {
        this.accelerations[i] = geometry_msgs.Twist();
        offset += this.accelerations[i].deserialize(buff, offset);
    }
    this.time_from_start.sec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.time_from_start.sec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.time_from_start.sec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.time_from_start.sec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.time_from_start.nsec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.time_from_start.nsec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.time_from_start.nsec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.time_from_start.nsec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

MultiDOFJointTrajectoryPoint.prototype.serializedLength = function() {
    var length = 0;
    var length_transforms = this.transforms.length;
    length += 4;
    for (var i = 0; i < length_transforms; i++) {
        length += this.transforms[i].serializedLength();
    }
    var length_velocities = this.velocities.length;
    length += 4;
    for (var i = 0; i < length_velocities; i++) {
        length += this.velocities[i].serializedLength();
    }
    var length_accelerations = this.accelerations.length;
    length += 4;
    for (var i = 0; i < length_accelerations; i++) {
        length += this.accelerations[i].serializedLength();
    }
    length += 4
    length += 4
    return length;
};

MultiDOFJointTrajectoryPoint.prototype.echo = function() { return ""; };

MultiDOFJointTrajectoryPoint.prototype.getType = function() { return "trajectory_msgs/MultiDOFJointTrajectoryPoint"; };

MultiDOFJointTrajectoryPoint.prototype.getMD5 = function() { return "f8b4a74b416279b5c5d565029308ff08"; };

MultiDOFJointTrajectoryPoint.prototype.getID = function() { return 0; };

MultiDOFJointTrajectoryPoint.prototype.setID = function(id) { };

return function() { return new MultiDOFJointTrajectoryPoint(); };

});
