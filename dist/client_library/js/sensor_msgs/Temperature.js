(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.Temperature=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function Temperature() {
    this.header = std_msgs.Header();
    this.temperature = 0.0;
    this.variance = 0.0;
};

Temperature.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var float64Array_temperature = new Float64Array(1);
    var uInt8Float64Array_temperature = new Uint8Array(float64Array_temperature.buffer);
    float64Array_temperature[0] = +this.temperature;
    buff[offset + 0] = uInt8Float64Array_temperature[0];
    buff[offset + 1] = uInt8Float64Array_temperature[1];
    buff[offset + 2] = uInt8Float64Array_temperature[2];
    buff[offset + 3] = uInt8Float64Array_temperature[3];
    buff[offset + 4] = uInt8Float64Array_temperature[4];
    buff[offset + 5] = uInt8Float64Array_temperature[5];
    buff[offset + 6] = uInt8Float64Array_temperature[6];
    buff[offset + 7] = uInt8Float64Array_temperature[7];
    offset += 8;
    var float64Array_variance = new Float64Array(1);
    var uInt8Float64Array_variance = new Uint8Array(float64Array_variance.buffer);
    float64Array_variance[0] = +this.variance;
    buff[offset + 0] = uInt8Float64Array_variance[0];
    buff[offset + 1] = uInt8Float64Array_variance[1];
    buff[offset + 2] = uInt8Float64Array_variance[2];
    buff[offset + 3] = uInt8Float64Array_variance[3];
    buff[offset + 4] = uInt8Float64Array_variance[4];
    buff[offset + 5] = uInt8Float64Array_variance[5];
    buff[offset + 6] = uInt8Float64Array_variance[6];
    buff[offset + 7] = uInt8Float64Array_variance[7];
    offset += 8;
    return offset;
};

Temperature.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var float64Array_temperature = new Float64Array(1);
    var uInt8Float64Array_temperature = new Uint8Array(float64Array_temperature.buffer);
    uInt8Float64Array_temperature[0] = buff[offset + 0];
    uInt8Float64Array_temperature[1] = buff[offset + 1];
    uInt8Float64Array_temperature[2] = buff[offset + 2];
    uInt8Float64Array_temperature[3] = buff[offset + 3];
    uInt8Float64Array_temperature[4] = buff[offset + 4];
    uInt8Float64Array_temperature[5] = buff[offset + 5];
    uInt8Float64Array_temperature[6] = buff[offset + 6];
    uInt8Float64Array_temperature[7] = buff[offset + 7];
    this.temperature = float64Array_temperature[0];
    offset += 8;
    var float64Array_variance = new Float64Array(1);
    var uInt8Float64Array_variance = new Uint8Array(float64Array_variance.buffer);
    uInt8Float64Array_variance[0] = buff[offset + 0];
    uInt8Float64Array_variance[1] = buff[offset + 1];
    uInt8Float64Array_variance[2] = buff[offset + 2];
    uInt8Float64Array_variance[3] = buff[offset + 3];
    uInt8Float64Array_variance[4] = buff[offset + 4];
    uInt8Float64Array_variance[5] = buff[offset + 5];
    uInt8Float64Array_variance[6] = buff[offset + 6];
    uInt8Float64Array_variance[7] = buff[offset + 7];
    this.variance = float64Array_variance[0];
    offset += 8;
    return offset;
};

Temperature.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 8
    length += 8
    return length;
};

Temperature.prototype.echo = function() { return ""; };

Temperature.prototype.getType = function() { return "sensor_msgs/Temperature"; };

Temperature.prototype.getMD5 = function() { return "898a0e5950c8e4073c0c3cf2d7e7bf26"; };

Temperature.prototype.getID = function() { return 0; };

Temperature.prototype.setID = function(id) { };

return function() { return new Temperature(); };

});
