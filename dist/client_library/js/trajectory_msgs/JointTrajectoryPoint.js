(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.trajectory_msgs==="undefined"){g.trajectory_msgs={};}g.trajectory_msgs.JointTrajectoryPoint=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Duration.js'></script>");

function JointTrajectoryPoint() {
    this.positions = new Array();
    this.velocities = new Array();
    this.accelerations = new Array();
    this.effort = new Array();
    this.time_from_start = tinyros.Duration();
};

JointTrajectoryPoint.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var length_positions = this.positions.length;
    buff[offset + 0] = (length_positions >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_positions >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_positions >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_positions >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_positions; i++) {
        var float64Array_positionsi = new Float64Array(1);
        var uInt8Float64Array_positionsi = new Uint8Array(float64Array_positionsi.buffer);
        float64Array_positionsi[0] = +this.positions[i];
        buff[offset + 0] = uInt8Float64Array_positionsi[0];
        buff[offset + 1] = uInt8Float64Array_positionsi[1];
        buff[offset + 2] = uInt8Float64Array_positionsi[2];
        buff[offset + 3] = uInt8Float64Array_positionsi[3];
        buff[offset + 4] = uInt8Float64Array_positionsi[4];
        buff[offset + 5] = uInt8Float64Array_positionsi[5];
        buff[offset + 6] = uInt8Float64Array_positionsi[6];
        buff[offset + 7] = uInt8Float64Array_positionsi[7];
        offset += 8;
    }
    var length_velocities = this.velocities.length;
    buff[offset + 0] = (length_velocities >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_velocities >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_velocities >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_velocities >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_velocities; i++) {
        var float64Array_velocitiesi = new Float64Array(1);
        var uInt8Float64Array_velocitiesi = new Uint8Array(float64Array_velocitiesi.buffer);
        float64Array_velocitiesi[0] = +this.velocities[i];
        buff[offset + 0] = uInt8Float64Array_velocitiesi[0];
        buff[offset + 1] = uInt8Float64Array_velocitiesi[1];
        buff[offset + 2] = uInt8Float64Array_velocitiesi[2];
        buff[offset + 3] = uInt8Float64Array_velocitiesi[3];
        buff[offset + 4] = uInt8Float64Array_velocitiesi[4];
        buff[offset + 5] = uInt8Float64Array_velocitiesi[5];
        buff[offset + 6] = uInt8Float64Array_velocitiesi[6];
        buff[offset + 7] = uInt8Float64Array_velocitiesi[7];
        offset += 8;
    }
    var length_accelerations = this.accelerations.length;
    buff[offset + 0] = (length_accelerations >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_accelerations >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_accelerations >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_accelerations >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_accelerations; i++) {
        var float64Array_accelerationsi = new Float64Array(1);
        var uInt8Float64Array_accelerationsi = new Uint8Array(float64Array_accelerationsi.buffer);
        float64Array_accelerationsi[0] = +this.accelerations[i];
        buff[offset + 0] = uInt8Float64Array_accelerationsi[0];
        buff[offset + 1] = uInt8Float64Array_accelerationsi[1];
        buff[offset + 2] = uInt8Float64Array_accelerationsi[2];
        buff[offset + 3] = uInt8Float64Array_accelerationsi[3];
        buff[offset + 4] = uInt8Float64Array_accelerationsi[4];
        buff[offset + 5] = uInt8Float64Array_accelerationsi[5];
        buff[offset + 6] = uInt8Float64Array_accelerationsi[6];
        buff[offset + 7] = uInt8Float64Array_accelerationsi[7];
        offset += 8;
    }
    var length_effort = this.effort.length;
    buff[offset + 0] = (length_effort >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_effort >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_effort >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_effort >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_effort; i++) {
        var float64Array_efforti = new Float64Array(1);
        var uInt8Float64Array_efforti = new Uint8Array(float64Array_efforti.buffer);
        float64Array_efforti[0] = +this.effort[i];
        buff[offset + 0] = uInt8Float64Array_efforti[0];
        buff[offset + 1] = uInt8Float64Array_efforti[1];
        buff[offset + 2] = uInt8Float64Array_efforti[2];
        buff[offset + 3] = uInt8Float64Array_efforti[3];
        buff[offset + 4] = uInt8Float64Array_efforti[4];
        buff[offset + 5] = uInt8Float64Array_efforti[5];
        buff[offset + 6] = uInt8Float64Array_efforti[6];
        buff[offset + 7] = uInt8Float64Array_efforti[7];
        offset += 8;
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

JointTrajectoryPoint.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_positions = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_positions |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_positions |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_positions |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.positions = new Array(length_positions);
    for (var i = 0; i < length_positions; i++) {
        var float64Array_positionsi = new Float64Array(1);
        var uInt8Float64Array_positionsi = new Uint8Array(float64Array_positionsi.buffer);
        uInt8Float64Array_positionsi[0] = buff[offset + 0];
        uInt8Float64Array_positionsi[1] = buff[offset + 1];
        uInt8Float64Array_positionsi[2] = buff[offset + 2];
        uInt8Float64Array_positionsi[3] = buff[offset + 3];
        uInt8Float64Array_positionsi[4] = buff[offset + 4];
        uInt8Float64Array_positionsi[5] = buff[offset + 5];
        uInt8Float64Array_positionsi[6] = buff[offset + 6];
        uInt8Float64Array_positionsi[7] = buff[offset + 7];
        this.positions[i] = float64Array_positionsi[0];
        offset += 8;
    }
    var length_velocities = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_velocities |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_velocities |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_velocities |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.velocities = new Array(length_velocities);
    for (var i = 0; i < length_velocities; i++) {
        var float64Array_velocitiesi = new Float64Array(1);
        var uInt8Float64Array_velocitiesi = new Uint8Array(float64Array_velocitiesi.buffer);
        uInt8Float64Array_velocitiesi[0] = buff[offset + 0];
        uInt8Float64Array_velocitiesi[1] = buff[offset + 1];
        uInt8Float64Array_velocitiesi[2] = buff[offset + 2];
        uInt8Float64Array_velocitiesi[3] = buff[offset + 3];
        uInt8Float64Array_velocitiesi[4] = buff[offset + 4];
        uInt8Float64Array_velocitiesi[5] = buff[offset + 5];
        uInt8Float64Array_velocitiesi[6] = buff[offset + 6];
        uInt8Float64Array_velocitiesi[7] = buff[offset + 7];
        this.velocities[i] = float64Array_velocitiesi[0];
        offset += 8;
    }
    var length_accelerations = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_accelerations |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_accelerations |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_accelerations |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.accelerations = new Array(length_accelerations);
    for (var i = 0; i < length_accelerations; i++) {
        var float64Array_accelerationsi = new Float64Array(1);
        var uInt8Float64Array_accelerationsi = new Uint8Array(float64Array_accelerationsi.buffer);
        uInt8Float64Array_accelerationsi[0] = buff[offset + 0];
        uInt8Float64Array_accelerationsi[1] = buff[offset + 1];
        uInt8Float64Array_accelerationsi[2] = buff[offset + 2];
        uInt8Float64Array_accelerationsi[3] = buff[offset + 3];
        uInt8Float64Array_accelerationsi[4] = buff[offset + 4];
        uInt8Float64Array_accelerationsi[5] = buff[offset + 5];
        uInt8Float64Array_accelerationsi[6] = buff[offset + 6];
        uInt8Float64Array_accelerationsi[7] = buff[offset + 7];
        this.accelerations[i] = float64Array_accelerationsi[0];
        offset += 8;
    }
    var length_effort = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_effort |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_effort |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_effort |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.effort = new Array(length_effort);
    for (var i = 0; i < length_effort; i++) {
        var float64Array_efforti = new Float64Array(1);
        var uInt8Float64Array_efforti = new Uint8Array(float64Array_efforti.buffer);
        uInt8Float64Array_efforti[0] = buff[offset + 0];
        uInt8Float64Array_efforti[1] = buff[offset + 1];
        uInt8Float64Array_efforti[2] = buff[offset + 2];
        uInt8Float64Array_efforti[3] = buff[offset + 3];
        uInt8Float64Array_efforti[4] = buff[offset + 4];
        uInt8Float64Array_efforti[5] = buff[offset + 5];
        uInt8Float64Array_efforti[6] = buff[offset + 6];
        uInt8Float64Array_efforti[7] = buff[offset + 7];
        this.effort[i] = float64Array_efforti[0];
        offset += 8;
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

JointTrajectoryPoint.prototype.serializedLength = function() {
    var length = 0;
    var length_positions = this.positions.length;
    length += 4;
    for (var i = 0; i < length_positions; i++) {
        length += 8
    }
    var length_velocities = this.velocities.length;
    length += 4;
    for (var i = 0; i < length_velocities; i++) {
        length += 8
    }
    var length_accelerations = this.accelerations.length;
    length += 4;
    for (var i = 0; i < length_accelerations; i++) {
        length += 8
    }
    var length_effort = this.effort.length;
    length += 4;
    for (var i = 0; i < length_effort; i++) {
        length += 8
    }
    length += 4
    length += 4
    return length;
};

JointTrajectoryPoint.prototype.echo = function() { return ""; };

JointTrajectoryPoint.prototype.getType = function() { return "trajectory_msgs/JointTrajectoryPoint"; };

JointTrajectoryPoint.prototype.getMD5 = function() { return "38679319321341510f6fde7d7f745eb0"; };

JointTrajectoryPoint.prototype.getID = function() { return 0; };

JointTrajectoryPoint.prototype.setID = function(id) { };

return function() { return new JointTrajectoryPoint(); };

});
