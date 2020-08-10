(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Float32=f();}})(function(){var define,module,exports;'use strict';

function Float32() {
    this.data = 0.0;
};

Float32.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var float32Array_data = new Float32Array(1);
    var uInt8Float32Array_data = new Uint8Array(float32Array_data.buffer);
    float32Array_data[0] = +this.data;
    buff[offset + 0] = uInt8Float32Array_data[0];
    buff[offset + 1] = uInt8Float32Array_data[1];
    buff[offset + 2] = uInt8Float32Array_data[2];
    buff[offset + 3] = uInt8Float32Array_data[3];
    offset += 4;
    return offset;
};

Float32.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var float32Array_data = new Float32Array(1);
    var uInt8Float32Array_data = new Uint8Array(float32Array_data.buffer);
    uInt8Float32Array_data[0] = buff[offset + 0];
    uInt8Float32Array_data[1] = buff[offset + 1];
    uInt8Float32Array_data[2] = buff[offset + 2];
    uInt8Float32Array_data[3] = buff[offset + 3];
    this.data = float32Array_data[0];
    offset += 4;
    return offset;
};

Float32.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

Float32.prototype.echo = function() { return ""; };

Float32.prototype.getType = function() { return "std_msgs/Float32"; };

Float32.prototype.getMD5 = function() { return "2aff5d2343e8e80ceea1362fc770035c"; };

Float32.prototype.getID = function() { return 0; };

Float32.prototype.setID = function(id) { };

return function() { return new Float32(); };

});
