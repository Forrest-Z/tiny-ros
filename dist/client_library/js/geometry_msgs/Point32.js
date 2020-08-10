(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.Point32=f();}})(function(){var define,module,exports;'use strict';

function Point32() {
    this.x = 0.0;
    this.y = 0.0;
    this.z = 0.0;
};

Point32.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var float32Array_x = new Float32Array(1);
    var uInt8Float32Array_x = new Uint8Array(float32Array_x.buffer);
    float32Array_x[0] = +this.x;
    buff[offset + 0] = uInt8Float32Array_x[0];
    buff[offset + 1] = uInt8Float32Array_x[1];
    buff[offset + 2] = uInt8Float32Array_x[2];
    buff[offset + 3] = uInt8Float32Array_x[3];
    offset += 4;
    var float32Array_y = new Float32Array(1);
    var uInt8Float32Array_y = new Uint8Array(float32Array_y.buffer);
    float32Array_y[0] = +this.y;
    buff[offset + 0] = uInt8Float32Array_y[0];
    buff[offset + 1] = uInt8Float32Array_y[1];
    buff[offset + 2] = uInt8Float32Array_y[2];
    buff[offset + 3] = uInt8Float32Array_y[3];
    offset += 4;
    var float32Array_z = new Float32Array(1);
    var uInt8Float32Array_z = new Uint8Array(float32Array_z.buffer);
    float32Array_z[0] = +this.z;
    buff[offset + 0] = uInt8Float32Array_z[0];
    buff[offset + 1] = uInt8Float32Array_z[1];
    buff[offset + 2] = uInt8Float32Array_z[2];
    buff[offset + 3] = uInt8Float32Array_z[3];
    offset += 4;
    return offset;
};

Point32.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var float32Array_x = new Float32Array(1);
    var uInt8Float32Array_x = new Uint8Array(float32Array_x.buffer);
    uInt8Float32Array_x[0] = buff[offset + 0];
    uInt8Float32Array_x[1] = buff[offset + 1];
    uInt8Float32Array_x[2] = buff[offset + 2];
    uInt8Float32Array_x[3] = buff[offset + 3];
    this.x = float32Array_x[0];
    offset += 4;
    var float32Array_y = new Float32Array(1);
    var uInt8Float32Array_y = new Uint8Array(float32Array_y.buffer);
    uInt8Float32Array_y[0] = buff[offset + 0];
    uInt8Float32Array_y[1] = buff[offset + 1];
    uInt8Float32Array_y[2] = buff[offset + 2];
    uInt8Float32Array_y[3] = buff[offset + 3];
    this.y = float32Array_y[0];
    offset += 4;
    var float32Array_z = new Float32Array(1);
    var uInt8Float32Array_z = new Uint8Array(float32Array_z.buffer);
    uInt8Float32Array_z[0] = buff[offset + 0];
    uInt8Float32Array_z[1] = buff[offset + 1];
    uInt8Float32Array_z[2] = buff[offset + 2];
    uInt8Float32Array_z[3] = buff[offset + 3];
    this.z = float32Array_z[0];
    offset += 4;
    return offset;
};

Point32.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 4
    length += 4
    return length;
};

Point32.prototype.echo = function() { return ""; };

Point32.prototype.getType = function() { return "geometry_msgs/Point32"; };

Point32.prototype.getMD5 = function() { return "b17f2230f465fce816e3773d7d59a841"; };

Point32.prototype.getID = function() { return 0; };

Point32.prototype.setID = function(id) { };

return function() { return new Point32(); };

});
