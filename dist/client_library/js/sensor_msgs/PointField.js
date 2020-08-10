(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.PointField=f();}})(function(){var define,module,exports;'use strict';

function PointField() {
    this.name = "";
    this.offset = 0;
    this.datatype = 0;
    this.count = 0;

    // ENUM{
    this.INT8 =  1;
    this.UINT8 =  2;
    this.INT16 =  3;
    this.UINT16 =  4;
    this.INT32 =  5;
    this.UINT32 =  6;
    this.FLOAT32 =  7;
    this.FLOAT64 =  8;
    // }ENUM
};

PointField.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var encoder_name = new TextEncoder('utf8');
    var utf8array_name = encoder_name.encode(this.name);
    buff[offset + 0] = (utf8array_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_name.length; i++) {
        buff[offset + i] = utf8array_name[i];
    }
    offset += utf8array_name.length;
    buff[offset + 0] = ((+this.offset) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.offset) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.offset) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.offset) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.datatype) >> (8 * 0)) & 0xFF;
    offset += 1;
    buff[offset + 0] = ((+this.count) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.count) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.count) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.count) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

PointField.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_name = new TextDecoder('utf8');
    this.name = decoder_name.decode(buff.slice(offset, offset + length_name));
    offset += length_name;
    this.offset = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.offset |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.offset |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.offset |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.datatype = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    this.count = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.count |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.count |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.count |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

PointField.prototype.serializedLength = function() {
    var length = 0;
    var encoder_name = new TextEncoder('utf8');
    var utf8array_name = encoder_name.encode(this.name);
    length += 4;
    length += utf8array_name.length;
    length += 4
    length += 1
    length += 4
    return length;
};

PointField.prototype.echo = function() { return ""; };

PointField.prototype.getType = function() { return "sensor_msgs/PointField"; };

PointField.prototype.getMD5 = function() { return "039974f05fdf0470d9dc865fd01dcc3e"; };

PointField.prototype.getID = function() { return 0; };

PointField.prototype.setID = function(id) { };

return function() { return new PointField(); };

});
