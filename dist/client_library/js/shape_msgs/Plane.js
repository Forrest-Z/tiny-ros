(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.shape_msgs==="undefined"){g.shape_msgs={};}g.shape_msgs.Plane=f();}})(function(){var define,module,exports;'use strict';

function Plane() {
    this.coef = [0.0,0.0,0.0,0.0];
};

Plane.prototype.serialize = function(buff, idx) {
    var offset = idx;
    for (var i = 0; i < 4; i++) {
        var float64Array_coefi = new Float64Array(1);
        var uInt8Float64Array_coefi = new Uint8Array(float64Array_coefi.buffer);
        float64Array_coefi[0] = +this.coef[i];
        buff[offset + 0] = uInt8Float64Array_coefi[0];
        buff[offset + 1] = uInt8Float64Array_coefi[1];
        buff[offset + 2] = uInt8Float64Array_coefi[2];
        buff[offset + 3] = uInt8Float64Array_coefi[3];
        buff[offset + 4] = uInt8Float64Array_coefi[4];
        buff[offset + 5] = uInt8Float64Array_coefi[5];
        buff[offset + 6] = uInt8Float64Array_coefi[6];
        buff[offset + 7] = uInt8Float64Array_coefi[7];
        offset += 8;
    }
    return offset;
};

Plane.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    for (var i = 0; i < 4; i++) {
        var float64Array_coefi = new Float64Array(1);
        var uInt8Float64Array_coefi = new Uint8Array(float64Array_coefi.buffer);
        uInt8Float64Array_coefi[0] = buff[offset + 0];
        uInt8Float64Array_coefi[1] = buff[offset + 1];
        uInt8Float64Array_coefi[2] = buff[offset + 2];
        uInt8Float64Array_coefi[3] = buff[offset + 3];
        uInt8Float64Array_coefi[4] = buff[offset + 4];
        uInt8Float64Array_coefi[5] = buff[offset + 5];
        uInt8Float64Array_coefi[6] = buff[offset + 6];
        uInt8Float64Array_coefi[7] = buff[offset + 7];
        this.coef[i] = float64Array_coefi[0];
        offset += 8;
    }
    return offset;
};

Plane.prototype.serializedLength = function() {
    var length = 0;
    for (var i = 0; i < 4; i++) {
        length += 8
    }
    return length;
};

Plane.prototype.echo = function() { return ""; };

Plane.prototype.getType = function() { return "shape_msgs/Plane"; };

Plane.prototype.getMD5 = function() { return "770421286b7c90effe8aac9f1c37eac0"; };

Plane.prototype.getID = function() { return 0; };

Plane.prototype.setID = function(id) { };

return function() { return new Plane(); };

});
