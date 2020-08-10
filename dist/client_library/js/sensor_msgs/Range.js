(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.Range=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function Range() {
    this.header = std_msgs.Header();
    this.radiation_type = 0;
    this.field_of_view = 0.0;
    this.min_range = 0.0;
    this.max_range = 0.0;
    this.range = 0.0;

    // ENUM{
    this.ULTRASOUND = 0;
    this.INFRARED = 1;
    // }ENUM
};

Range.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    buff[offset + 0] = ((+this.radiation_type) >> (8 * 0)) & 0xFF;
    offset += 1;
    var float32Array_field_of_view = new Float32Array(1);
    var uInt8Float32Array_field_of_view = new Uint8Array(float32Array_field_of_view.buffer);
    float32Array_field_of_view[0] = +this.field_of_view;
    buff[offset + 0] = uInt8Float32Array_field_of_view[0];
    buff[offset + 1] = uInt8Float32Array_field_of_view[1];
    buff[offset + 2] = uInt8Float32Array_field_of_view[2];
    buff[offset + 3] = uInt8Float32Array_field_of_view[3];
    offset += 4;
    var float32Array_min_range = new Float32Array(1);
    var uInt8Float32Array_min_range = new Uint8Array(float32Array_min_range.buffer);
    float32Array_min_range[0] = +this.min_range;
    buff[offset + 0] = uInt8Float32Array_min_range[0];
    buff[offset + 1] = uInt8Float32Array_min_range[1];
    buff[offset + 2] = uInt8Float32Array_min_range[2];
    buff[offset + 3] = uInt8Float32Array_min_range[3];
    offset += 4;
    var float32Array_max_range = new Float32Array(1);
    var uInt8Float32Array_max_range = new Uint8Array(float32Array_max_range.buffer);
    float32Array_max_range[0] = +this.max_range;
    buff[offset + 0] = uInt8Float32Array_max_range[0];
    buff[offset + 1] = uInt8Float32Array_max_range[1];
    buff[offset + 2] = uInt8Float32Array_max_range[2];
    buff[offset + 3] = uInt8Float32Array_max_range[3];
    offset += 4;
    var float32Array_range = new Float32Array(1);
    var uInt8Float32Array_range = new Uint8Array(float32Array_range.buffer);
    float32Array_range[0] = +this.range;
    buff[offset + 0] = uInt8Float32Array_range[0];
    buff[offset + 1] = uInt8Float32Array_range[1];
    buff[offset + 2] = uInt8Float32Array_range[2];
    buff[offset + 3] = uInt8Float32Array_range[3];
    offset += 4;
    return offset;
};

Range.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    this.radiation_type = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    var float32Array_field_of_view = new Float32Array(1);
    var uInt8Float32Array_field_of_view = new Uint8Array(float32Array_field_of_view.buffer);
    uInt8Float32Array_field_of_view[0] = buff[offset + 0];
    uInt8Float32Array_field_of_view[1] = buff[offset + 1];
    uInt8Float32Array_field_of_view[2] = buff[offset + 2];
    uInt8Float32Array_field_of_view[3] = buff[offset + 3];
    this.field_of_view = float32Array_field_of_view[0];
    offset += 4;
    var float32Array_min_range = new Float32Array(1);
    var uInt8Float32Array_min_range = new Uint8Array(float32Array_min_range.buffer);
    uInt8Float32Array_min_range[0] = buff[offset + 0];
    uInt8Float32Array_min_range[1] = buff[offset + 1];
    uInt8Float32Array_min_range[2] = buff[offset + 2];
    uInt8Float32Array_min_range[3] = buff[offset + 3];
    this.min_range = float32Array_min_range[0];
    offset += 4;
    var float32Array_max_range = new Float32Array(1);
    var uInt8Float32Array_max_range = new Uint8Array(float32Array_max_range.buffer);
    uInt8Float32Array_max_range[0] = buff[offset + 0];
    uInt8Float32Array_max_range[1] = buff[offset + 1];
    uInt8Float32Array_max_range[2] = buff[offset + 2];
    uInt8Float32Array_max_range[3] = buff[offset + 3];
    this.max_range = float32Array_max_range[0];
    offset += 4;
    var float32Array_range = new Float32Array(1);
    var uInt8Float32Array_range = new Uint8Array(float32Array_range.buffer);
    uInt8Float32Array_range[0] = buff[offset + 0];
    uInt8Float32Array_range[1] = buff[offset + 1];
    uInt8Float32Array_range[2] = buff[offset + 2];
    uInt8Float32Array_range[3] = buff[offset + 3];
    this.range = float32Array_range[0];
    offset += 4;
    return offset;
};

Range.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 1
    length += 4
    length += 4
    length += 4
    length += 4
    return length;
};

Range.prototype.echo = function() { return ""; };

Range.prototype.getType = function() { return "sensor_msgs/Range"; };

Range.prototype.getMD5 = function() { return "54d647e3a481f5b87ce39d1b97e84d53"; };

Range.prototype.getID = function() { return 0; };

Range.prototype.setID = function(id) { };

return function() { return new Range(); };

});
