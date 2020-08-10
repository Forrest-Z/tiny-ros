(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.ColorRGBA=f();}})(function(){var define,module,exports;'use strict';

function ColorRGBA() {
    this.r = 0.0;
    this.g = 0.0;
    this.b = 0.0;
    this.a = 0.0;
};

ColorRGBA.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var float32Array_r = new Float32Array(1);
    var uInt8Float32Array_r = new Uint8Array(float32Array_r.buffer);
    float32Array_r[0] = +this.r;
    buff[offset + 0] = uInt8Float32Array_r[0];
    buff[offset + 1] = uInt8Float32Array_r[1];
    buff[offset + 2] = uInt8Float32Array_r[2];
    buff[offset + 3] = uInt8Float32Array_r[3];
    offset += 4;
    var float32Array_g = new Float32Array(1);
    var uInt8Float32Array_g = new Uint8Array(float32Array_g.buffer);
    float32Array_g[0] = +this.g;
    buff[offset + 0] = uInt8Float32Array_g[0];
    buff[offset + 1] = uInt8Float32Array_g[1];
    buff[offset + 2] = uInt8Float32Array_g[2];
    buff[offset + 3] = uInt8Float32Array_g[3];
    offset += 4;
    var float32Array_b = new Float32Array(1);
    var uInt8Float32Array_b = new Uint8Array(float32Array_b.buffer);
    float32Array_b[0] = +this.b;
    buff[offset + 0] = uInt8Float32Array_b[0];
    buff[offset + 1] = uInt8Float32Array_b[1];
    buff[offset + 2] = uInt8Float32Array_b[2];
    buff[offset + 3] = uInt8Float32Array_b[3];
    offset += 4;
    var float32Array_a = new Float32Array(1);
    var uInt8Float32Array_a = new Uint8Array(float32Array_a.buffer);
    float32Array_a[0] = +this.a;
    buff[offset + 0] = uInt8Float32Array_a[0];
    buff[offset + 1] = uInt8Float32Array_a[1];
    buff[offset + 2] = uInt8Float32Array_a[2];
    buff[offset + 3] = uInt8Float32Array_a[3];
    offset += 4;
    return offset;
};

ColorRGBA.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var float32Array_r = new Float32Array(1);
    var uInt8Float32Array_r = new Uint8Array(float32Array_r.buffer);
    uInt8Float32Array_r[0] = buff[offset + 0];
    uInt8Float32Array_r[1] = buff[offset + 1];
    uInt8Float32Array_r[2] = buff[offset + 2];
    uInt8Float32Array_r[3] = buff[offset + 3];
    this.r = float32Array_r[0];
    offset += 4;
    var float32Array_g = new Float32Array(1);
    var uInt8Float32Array_g = new Uint8Array(float32Array_g.buffer);
    uInt8Float32Array_g[0] = buff[offset + 0];
    uInt8Float32Array_g[1] = buff[offset + 1];
    uInt8Float32Array_g[2] = buff[offset + 2];
    uInt8Float32Array_g[3] = buff[offset + 3];
    this.g = float32Array_g[0];
    offset += 4;
    var float32Array_b = new Float32Array(1);
    var uInt8Float32Array_b = new Uint8Array(float32Array_b.buffer);
    uInt8Float32Array_b[0] = buff[offset + 0];
    uInt8Float32Array_b[1] = buff[offset + 1];
    uInt8Float32Array_b[2] = buff[offset + 2];
    uInt8Float32Array_b[3] = buff[offset + 3];
    this.b = float32Array_b[0];
    offset += 4;
    var float32Array_a = new Float32Array(1);
    var uInt8Float32Array_a = new Uint8Array(float32Array_a.buffer);
    uInt8Float32Array_a[0] = buff[offset + 0];
    uInt8Float32Array_a[1] = buff[offset + 1];
    uInt8Float32Array_a[2] = buff[offset + 2];
    uInt8Float32Array_a[3] = buff[offset + 3];
    this.a = float32Array_a[0];
    offset += 4;
    return offset;
};

ColorRGBA.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 4
    length += 4
    length += 4
    return length;
};

ColorRGBA.prototype.echo = function() { return ""; };

ColorRGBA.prototype.getType = function() { return "std_msgs/ColorRGBA"; };

ColorRGBA.prototype.getMD5 = function() { return "3c740aa311a337bfa4133c69405e0aed"; };

ColorRGBA.prototype.getID = function() { return 0; };

ColorRGBA.prototype.setID = function(id) { };

return function() { return new ColorRGBA(); };

});
