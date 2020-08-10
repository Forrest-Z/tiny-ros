(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Float64=f();}})(function(){var define,module,exports;'use strict';

function Float64() {
    this.data = 0.0;
};

Float64.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var float64Array_data = new Float64Array(1);
    var uInt8Float64Array_data = new Uint8Array(float64Array_data.buffer);
    float64Array_data[0] = +this.data;
    buff[offset + 0] = uInt8Float64Array_data[0];
    buff[offset + 1] = uInt8Float64Array_data[1];
    buff[offset + 2] = uInt8Float64Array_data[2];
    buff[offset + 3] = uInt8Float64Array_data[3];
    buff[offset + 4] = uInt8Float64Array_data[4];
    buff[offset + 5] = uInt8Float64Array_data[5];
    buff[offset + 6] = uInt8Float64Array_data[6];
    buff[offset + 7] = uInt8Float64Array_data[7];
    offset += 8;
    return offset;
};

Float64.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var float64Array_data = new Float64Array(1);
    var uInt8Float64Array_data = new Uint8Array(float64Array_data.buffer);
    uInt8Float64Array_data[0] = buff[offset + 0];
    uInt8Float64Array_data[1] = buff[offset + 1];
    uInt8Float64Array_data[2] = buff[offset + 2];
    uInt8Float64Array_data[3] = buff[offset + 3];
    uInt8Float64Array_data[4] = buff[offset + 4];
    uInt8Float64Array_data[5] = buff[offset + 5];
    uInt8Float64Array_data[6] = buff[offset + 6];
    uInt8Float64Array_data[7] = buff[offset + 7];
    this.data = float64Array_data[0];
    offset += 8;
    return offset;
};

Float64.prototype.serializedLength = function() {
    var length = 0;
    length += 8
    return length;
};

Float64.prototype.echo = function() { return ""; };

Float64.prototype.getType = function() { return "std_msgs/Float64"; };

Float64.prototype.getMD5 = function() { return "3439fe880debfd63cf4e61e62e285345"; };

Float64.prototype.getID = function() { return 0; };

Float64.prototype.setID = function(id) { };

return function() { return new Float64(); };

});
