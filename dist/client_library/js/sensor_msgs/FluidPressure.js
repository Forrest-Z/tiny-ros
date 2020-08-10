(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.FluidPressure=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function FluidPressure() {
    this.header = std_msgs.Header();
    this.fluid_pressure = 0.0;
    this.variance = 0.0;
};

FluidPressure.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var float64Array_fluid_pressure = new Float64Array(1);
    var uInt8Float64Array_fluid_pressure = new Uint8Array(float64Array_fluid_pressure.buffer);
    float64Array_fluid_pressure[0] = +this.fluid_pressure;
    buff[offset + 0] = uInt8Float64Array_fluid_pressure[0];
    buff[offset + 1] = uInt8Float64Array_fluid_pressure[1];
    buff[offset + 2] = uInt8Float64Array_fluid_pressure[2];
    buff[offset + 3] = uInt8Float64Array_fluid_pressure[3];
    buff[offset + 4] = uInt8Float64Array_fluid_pressure[4];
    buff[offset + 5] = uInt8Float64Array_fluid_pressure[5];
    buff[offset + 6] = uInt8Float64Array_fluid_pressure[6];
    buff[offset + 7] = uInt8Float64Array_fluid_pressure[7];
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

FluidPressure.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var float64Array_fluid_pressure = new Float64Array(1);
    var uInt8Float64Array_fluid_pressure = new Uint8Array(float64Array_fluid_pressure.buffer);
    uInt8Float64Array_fluid_pressure[0] = buff[offset + 0];
    uInt8Float64Array_fluid_pressure[1] = buff[offset + 1];
    uInt8Float64Array_fluid_pressure[2] = buff[offset + 2];
    uInt8Float64Array_fluid_pressure[3] = buff[offset + 3];
    uInt8Float64Array_fluid_pressure[4] = buff[offset + 4];
    uInt8Float64Array_fluid_pressure[5] = buff[offset + 5];
    uInt8Float64Array_fluid_pressure[6] = buff[offset + 6];
    uInt8Float64Array_fluid_pressure[7] = buff[offset + 7];
    this.fluid_pressure = float64Array_fluid_pressure[0];
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

FluidPressure.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 8
    length += 8
    return length;
};

FluidPressure.prototype.echo = function() { return ""; };

FluidPressure.prototype.getType = function() { return "sensor_msgs/FluidPressure"; };

FluidPressure.prototype.getMD5 = function() { return "0fdea137019d78ebf8c2cb91c31a458a"; };

FluidPressure.prototype.getID = function() { return 0; };

FluidPressure.prototype.setID = function(id) { };

return function() { return new FluidPressure(); };

});
