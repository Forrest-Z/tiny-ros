(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.Imu=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Quaternion.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Vector3.js'></script>");

function Imu() {
    this.header = std_msgs.Header();
    this.orientation = geometry_msgs.Quaternion();
    this.orientation_covariance = [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0];
    this.angular_velocity = geometry_msgs.Vector3();
    this.angular_velocity_covariance = [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0];
    this.linear_acceleration = geometry_msgs.Vector3();
    this.linear_acceleration_covariance = [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0];
};

Imu.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.orientation.serialize(buff, offset);
    for (var i = 0; i < 9; i++) {
        var float64Array_orientation_covariancei = new Float64Array(1);
        var uInt8Float64Array_orientation_covariancei = new Uint8Array(float64Array_orientation_covariancei.buffer);
        float64Array_orientation_covariancei[0] = +this.orientation_covariance[i];
        buff[offset + 0] = uInt8Float64Array_orientation_covariancei[0];
        buff[offset + 1] = uInt8Float64Array_orientation_covariancei[1];
        buff[offset + 2] = uInt8Float64Array_orientation_covariancei[2];
        buff[offset + 3] = uInt8Float64Array_orientation_covariancei[3];
        buff[offset + 4] = uInt8Float64Array_orientation_covariancei[4];
        buff[offset + 5] = uInt8Float64Array_orientation_covariancei[5];
        buff[offset + 6] = uInt8Float64Array_orientation_covariancei[6];
        buff[offset + 7] = uInt8Float64Array_orientation_covariancei[7];
        offset += 8;
    }
    offset += this.angular_velocity.serialize(buff, offset);
    for (var i = 0; i < 9; i++) {
        var float64Array_angular_velocity_covariancei = new Float64Array(1);
        var uInt8Float64Array_angular_velocity_covariancei = new Uint8Array(float64Array_angular_velocity_covariancei.buffer);
        float64Array_angular_velocity_covariancei[0] = +this.angular_velocity_covariance[i];
        buff[offset + 0] = uInt8Float64Array_angular_velocity_covariancei[0];
        buff[offset + 1] = uInt8Float64Array_angular_velocity_covariancei[1];
        buff[offset + 2] = uInt8Float64Array_angular_velocity_covariancei[2];
        buff[offset + 3] = uInt8Float64Array_angular_velocity_covariancei[3];
        buff[offset + 4] = uInt8Float64Array_angular_velocity_covariancei[4];
        buff[offset + 5] = uInt8Float64Array_angular_velocity_covariancei[5];
        buff[offset + 6] = uInt8Float64Array_angular_velocity_covariancei[6];
        buff[offset + 7] = uInt8Float64Array_angular_velocity_covariancei[7];
        offset += 8;
    }
    offset += this.linear_acceleration.serialize(buff, offset);
    for (var i = 0; i < 9; i++) {
        var float64Array_linear_acceleration_covariancei = new Float64Array(1);
        var uInt8Float64Array_linear_acceleration_covariancei = new Uint8Array(float64Array_linear_acceleration_covariancei.buffer);
        float64Array_linear_acceleration_covariancei[0] = +this.linear_acceleration_covariance[i];
        buff[offset + 0] = uInt8Float64Array_linear_acceleration_covariancei[0];
        buff[offset + 1] = uInt8Float64Array_linear_acceleration_covariancei[1];
        buff[offset + 2] = uInt8Float64Array_linear_acceleration_covariancei[2];
        buff[offset + 3] = uInt8Float64Array_linear_acceleration_covariancei[3];
        buff[offset + 4] = uInt8Float64Array_linear_acceleration_covariancei[4];
        buff[offset + 5] = uInt8Float64Array_linear_acceleration_covariancei[5];
        buff[offset + 6] = uInt8Float64Array_linear_acceleration_covariancei[6];
        buff[offset + 7] = uInt8Float64Array_linear_acceleration_covariancei[7];
        offset += 8;
    }
    return offset;
};

Imu.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.orientation.deserialize(buff, offset);
    for (var i = 0; i < 9; i++) {
        var float64Array_orientation_covariancei = new Float64Array(1);
        var uInt8Float64Array_orientation_covariancei = new Uint8Array(float64Array_orientation_covariancei.buffer);
        uInt8Float64Array_orientation_covariancei[0] = buff[offset + 0];
        uInt8Float64Array_orientation_covariancei[1] = buff[offset + 1];
        uInt8Float64Array_orientation_covariancei[2] = buff[offset + 2];
        uInt8Float64Array_orientation_covariancei[3] = buff[offset + 3];
        uInt8Float64Array_orientation_covariancei[4] = buff[offset + 4];
        uInt8Float64Array_orientation_covariancei[5] = buff[offset + 5];
        uInt8Float64Array_orientation_covariancei[6] = buff[offset + 6];
        uInt8Float64Array_orientation_covariancei[7] = buff[offset + 7];
        this.orientation_covariance[i] = float64Array_orientation_covariancei[0];
        offset += 8;
    }
    offset += this.angular_velocity.deserialize(buff, offset);
    for (var i = 0; i < 9; i++) {
        var float64Array_angular_velocity_covariancei = new Float64Array(1);
        var uInt8Float64Array_angular_velocity_covariancei = new Uint8Array(float64Array_angular_velocity_covariancei.buffer);
        uInt8Float64Array_angular_velocity_covariancei[0] = buff[offset + 0];
        uInt8Float64Array_angular_velocity_covariancei[1] = buff[offset + 1];
        uInt8Float64Array_angular_velocity_covariancei[2] = buff[offset + 2];
        uInt8Float64Array_angular_velocity_covariancei[3] = buff[offset + 3];
        uInt8Float64Array_angular_velocity_covariancei[4] = buff[offset + 4];
        uInt8Float64Array_angular_velocity_covariancei[5] = buff[offset + 5];
        uInt8Float64Array_angular_velocity_covariancei[6] = buff[offset + 6];
        uInt8Float64Array_angular_velocity_covariancei[7] = buff[offset + 7];
        this.angular_velocity_covariance[i] = float64Array_angular_velocity_covariancei[0];
        offset += 8;
    }
    offset += this.linear_acceleration.deserialize(buff, offset);
    for (var i = 0; i < 9; i++) {
        var float64Array_linear_acceleration_covariancei = new Float64Array(1);
        var uInt8Float64Array_linear_acceleration_covariancei = new Uint8Array(float64Array_linear_acceleration_covariancei.buffer);
        uInt8Float64Array_linear_acceleration_covariancei[0] = buff[offset + 0];
        uInt8Float64Array_linear_acceleration_covariancei[1] = buff[offset + 1];
        uInt8Float64Array_linear_acceleration_covariancei[2] = buff[offset + 2];
        uInt8Float64Array_linear_acceleration_covariancei[3] = buff[offset + 3];
        uInt8Float64Array_linear_acceleration_covariancei[4] = buff[offset + 4];
        uInt8Float64Array_linear_acceleration_covariancei[5] = buff[offset + 5];
        uInt8Float64Array_linear_acceleration_covariancei[6] = buff[offset + 6];
        uInt8Float64Array_linear_acceleration_covariancei[7] = buff[offset + 7];
        this.linear_acceleration_covariance[i] = float64Array_linear_acceleration_covariancei[0];
        offset += 8;
    }
    return offset;
};

Imu.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.orientation.serializedLength();
    for (var i = 0; i < 9; i++) {
        length += 8
    }
    length += this.angular_velocity.serializedLength();
    for (var i = 0; i < 9; i++) {
        length += 8
    }
    length += this.linear_acceleration.serializedLength();
    for (var i = 0; i < 9; i++) {
        length += 8
    }
    return length;
};

Imu.prototype.echo = function() { return ""; };

Imu.prototype.getType = function() { return "sensor_msgs/Imu"; };

Imu.prototype.getMD5 = function() { return "a42c1ab94665a5807834c0ea19a6d16a"; };

Imu.prototype.getID = function() { return 0; };

Imu.prototype.setID = function(id) { };

return function() { return new Imu(); };

});
