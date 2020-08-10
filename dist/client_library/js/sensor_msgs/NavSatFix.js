(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.NavSatFix=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/sensor_msgs/NavSatStatus.js'></script>");

function NavSatFix() {
    this.header = std_msgs.Header();
    this.status = sensor_msgs.NavSatStatus();
    this.latitude = 0.0;
    this.longitude = 0.0;
    this.altitude = 0.0;
    this.position_covariance = [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0];
    this.position_covariance_type = 0;

    // ENUM{
    this.COVARIANCE_TYPE_UNKNOWN =  0;
    this.COVARIANCE_TYPE_APPROXIMATED =  1;
    this.COVARIANCE_TYPE_DIAGONAL_KNOWN =  2;
    this.COVARIANCE_TYPE_KNOWN =  3;
    // }ENUM
};

NavSatFix.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.status.serialize(buff, offset);
    var float64Array_latitude = new Float64Array(1);
    var uInt8Float64Array_latitude = new Uint8Array(float64Array_latitude.buffer);
    float64Array_latitude[0] = +this.latitude;
    buff[offset + 0] = uInt8Float64Array_latitude[0];
    buff[offset + 1] = uInt8Float64Array_latitude[1];
    buff[offset + 2] = uInt8Float64Array_latitude[2];
    buff[offset + 3] = uInt8Float64Array_latitude[3];
    buff[offset + 4] = uInt8Float64Array_latitude[4];
    buff[offset + 5] = uInt8Float64Array_latitude[5];
    buff[offset + 6] = uInt8Float64Array_latitude[6];
    buff[offset + 7] = uInt8Float64Array_latitude[7];
    offset += 8;
    var float64Array_longitude = new Float64Array(1);
    var uInt8Float64Array_longitude = new Uint8Array(float64Array_longitude.buffer);
    float64Array_longitude[0] = +this.longitude;
    buff[offset + 0] = uInt8Float64Array_longitude[0];
    buff[offset + 1] = uInt8Float64Array_longitude[1];
    buff[offset + 2] = uInt8Float64Array_longitude[2];
    buff[offset + 3] = uInt8Float64Array_longitude[3];
    buff[offset + 4] = uInt8Float64Array_longitude[4];
    buff[offset + 5] = uInt8Float64Array_longitude[5];
    buff[offset + 6] = uInt8Float64Array_longitude[6];
    buff[offset + 7] = uInt8Float64Array_longitude[7];
    offset += 8;
    var float64Array_altitude = new Float64Array(1);
    var uInt8Float64Array_altitude = new Uint8Array(float64Array_altitude.buffer);
    float64Array_altitude[0] = +this.altitude;
    buff[offset + 0] = uInt8Float64Array_altitude[0];
    buff[offset + 1] = uInt8Float64Array_altitude[1];
    buff[offset + 2] = uInt8Float64Array_altitude[2];
    buff[offset + 3] = uInt8Float64Array_altitude[3];
    buff[offset + 4] = uInt8Float64Array_altitude[4];
    buff[offset + 5] = uInt8Float64Array_altitude[5];
    buff[offset + 6] = uInt8Float64Array_altitude[6];
    buff[offset + 7] = uInt8Float64Array_altitude[7];
    offset += 8;
    for (var i = 0; i < 9; i++) {
        var float64Array_position_covariancei = new Float64Array(1);
        var uInt8Float64Array_position_covariancei = new Uint8Array(float64Array_position_covariancei.buffer);
        float64Array_position_covariancei[0] = +this.position_covariance[i];
        buff[offset + 0] = uInt8Float64Array_position_covariancei[0];
        buff[offset + 1] = uInt8Float64Array_position_covariancei[1];
        buff[offset + 2] = uInt8Float64Array_position_covariancei[2];
        buff[offset + 3] = uInt8Float64Array_position_covariancei[3];
        buff[offset + 4] = uInt8Float64Array_position_covariancei[4];
        buff[offset + 5] = uInt8Float64Array_position_covariancei[5];
        buff[offset + 6] = uInt8Float64Array_position_covariancei[6];
        buff[offset + 7] = uInt8Float64Array_position_covariancei[7];
        offset += 8;
    }
    buff[offset + 0] = ((+this.position_covariance_type) >> (8 * 0)) & 0xFF;
    offset += 1;
    return offset;
};

NavSatFix.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.status.deserialize(buff, offset);
    var float64Array_latitude = new Float64Array(1);
    var uInt8Float64Array_latitude = new Uint8Array(float64Array_latitude.buffer);
    uInt8Float64Array_latitude[0] = buff[offset + 0];
    uInt8Float64Array_latitude[1] = buff[offset + 1];
    uInt8Float64Array_latitude[2] = buff[offset + 2];
    uInt8Float64Array_latitude[3] = buff[offset + 3];
    uInt8Float64Array_latitude[4] = buff[offset + 4];
    uInt8Float64Array_latitude[5] = buff[offset + 5];
    uInt8Float64Array_latitude[6] = buff[offset + 6];
    uInt8Float64Array_latitude[7] = buff[offset + 7];
    this.latitude = float64Array_latitude[0];
    offset += 8;
    var float64Array_longitude = new Float64Array(1);
    var uInt8Float64Array_longitude = new Uint8Array(float64Array_longitude.buffer);
    uInt8Float64Array_longitude[0] = buff[offset + 0];
    uInt8Float64Array_longitude[1] = buff[offset + 1];
    uInt8Float64Array_longitude[2] = buff[offset + 2];
    uInt8Float64Array_longitude[3] = buff[offset + 3];
    uInt8Float64Array_longitude[4] = buff[offset + 4];
    uInt8Float64Array_longitude[5] = buff[offset + 5];
    uInt8Float64Array_longitude[6] = buff[offset + 6];
    uInt8Float64Array_longitude[7] = buff[offset + 7];
    this.longitude = float64Array_longitude[0];
    offset += 8;
    var float64Array_altitude = new Float64Array(1);
    var uInt8Float64Array_altitude = new Uint8Array(float64Array_altitude.buffer);
    uInt8Float64Array_altitude[0] = buff[offset + 0];
    uInt8Float64Array_altitude[1] = buff[offset + 1];
    uInt8Float64Array_altitude[2] = buff[offset + 2];
    uInt8Float64Array_altitude[3] = buff[offset + 3];
    uInt8Float64Array_altitude[4] = buff[offset + 4];
    uInt8Float64Array_altitude[5] = buff[offset + 5];
    uInt8Float64Array_altitude[6] = buff[offset + 6];
    uInt8Float64Array_altitude[7] = buff[offset + 7];
    this.altitude = float64Array_altitude[0];
    offset += 8;
    for (var i = 0; i < 9; i++) {
        var float64Array_position_covariancei = new Float64Array(1);
        var uInt8Float64Array_position_covariancei = new Uint8Array(float64Array_position_covariancei.buffer);
        uInt8Float64Array_position_covariancei[0] = buff[offset + 0];
        uInt8Float64Array_position_covariancei[1] = buff[offset + 1];
        uInt8Float64Array_position_covariancei[2] = buff[offset + 2];
        uInt8Float64Array_position_covariancei[3] = buff[offset + 3];
        uInt8Float64Array_position_covariancei[4] = buff[offset + 4];
        uInt8Float64Array_position_covariancei[5] = buff[offset + 5];
        uInt8Float64Array_position_covariancei[6] = buff[offset + 6];
        uInt8Float64Array_position_covariancei[7] = buff[offset + 7];
        this.position_covariance[i] = float64Array_position_covariancei[0];
        offset += 8;
    }
    this.position_covariance_type = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    return offset;
};

NavSatFix.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.status.serializedLength();
    length += 8
    length += 8
    length += 8
    for (var i = 0; i < 9; i++) {
        length += 8
    }
    length += 1
    return length;
};

NavSatFix.prototype.echo = function() { return ""; };

NavSatFix.prototype.getType = function() { return "sensor_msgs/NavSatFix"; };

NavSatFix.prototype.getMD5 = function() { return "adc27ca9d8634ed087021b82f3c43576"; };

NavSatFix.prototype.getID = function() { return 0; };

NavSatFix.prototype.setID = function(id) { };

return function() { return new NavSatFix(); };

});
