(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.shape_msgs==="undefined"){g.shape_msgs={};}g.shape_msgs.SolidPrimitive=f();}})(function(){var define,module,exports;'use strict';

function SolidPrimitive() {
    this.type = 0;
    this.dimensions = new Array();

    // ENUM{
    this.BOX = 1;
    this.SPHERE = 2;
    this.CYLINDER = 3;
    this.CONE = 4;
    this.BOX_X = 0;
    this.BOX_Y = 1;
    this.BOX_Z = 2;
    this.SPHERE_RADIUS = 0;
    this.CYLINDER_HEIGHT = 0;
    this.CYLINDER_RADIUS = 1;
    this.CONE_HEIGHT = 0;
    this.CONE_RADIUS = 1;
    // }ENUM
};

SolidPrimitive.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.type) >> (8 * 0)) & 0xFF;
    offset += 1;
    var length_dimensions = this.dimensions.length;
    buff[offset + 0] = (length_dimensions >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_dimensions >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_dimensions >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_dimensions >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_dimensions; i++) {
        var float64Array_dimensionsi = new Float64Array(1);
        var uInt8Float64Array_dimensionsi = new Uint8Array(float64Array_dimensionsi.buffer);
        float64Array_dimensionsi[0] = +this.dimensions[i];
        buff[offset + 0] = uInt8Float64Array_dimensionsi[0];
        buff[offset + 1] = uInt8Float64Array_dimensionsi[1];
        buff[offset + 2] = uInt8Float64Array_dimensionsi[2];
        buff[offset + 3] = uInt8Float64Array_dimensionsi[3];
        buff[offset + 4] = uInt8Float64Array_dimensionsi[4];
        buff[offset + 5] = uInt8Float64Array_dimensionsi[5];
        buff[offset + 6] = uInt8Float64Array_dimensionsi[6];
        buff[offset + 7] = uInt8Float64Array_dimensionsi[7];
        offset += 8;
    }
    return offset;
};

SolidPrimitive.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.type = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    var length_dimensions = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_dimensions |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_dimensions |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_dimensions |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.dimensions = new Array(length_dimensions);
    for (var i = 0; i < length_dimensions; i++) {
        var float64Array_dimensionsi = new Float64Array(1);
        var uInt8Float64Array_dimensionsi = new Uint8Array(float64Array_dimensionsi.buffer);
        uInt8Float64Array_dimensionsi[0] = buff[offset + 0];
        uInt8Float64Array_dimensionsi[1] = buff[offset + 1];
        uInt8Float64Array_dimensionsi[2] = buff[offset + 2];
        uInt8Float64Array_dimensionsi[3] = buff[offset + 3];
        uInt8Float64Array_dimensionsi[4] = buff[offset + 4];
        uInt8Float64Array_dimensionsi[5] = buff[offset + 5];
        uInt8Float64Array_dimensionsi[6] = buff[offset + 6];
        uInt8Float64Array_dimensionsi[7] = buff[offset + 7];
        this.dimensions[i] = float64Array_dimensionsi[0];
        offset += 8;
    }
    return offset;
};

SolidPrimitive.prototype.serializedLength = function() {
    var length = 0;
    length += 1
    var length_dimensions = this.dimensions.length;
    length += 4;
    for (var i = 0; i < length_dimensions; i++) {
        length += 8
    }
    return length;
};

SolidPrimitive.prototype.echo = function() { return ""; };

SolidPrimitive.prototype.getType = function() { return "shape_msgs/SolidPrimitive"; };

SolidPrimitive.prototype.getMD5 = function() { return "f40805922065416be24909fd8fd751b5"; };

SolidPrimitive.prototype.getID = function() { return 0; };

SolidPrimitive.prototype.setID = function(id) { };

return function() { return new SolidPrimitive(); };

});
