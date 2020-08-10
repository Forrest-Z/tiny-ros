(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.ChannelFloat32=f();}})(function(){var define,module,exports;'use strict';

function ChannelFloat32() {
    this.name = "";
    this.values = new Array();
};

ChannelFloat32.prototype.serialize = function(buff, idx) {
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
    var length_values = this.values.length;
    buff[offset + 0] = (length_values >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_values >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_values >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_values >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_values; i++) {
        var float32Array_valuesi = new Float32Array(1);
        var uInt8Float32Array_valuesi = new Uint8Array(float32Array_valuesi.buffer);
        float32Array_valuesi[0] = +this.values[i];
        buff[offset + 0] = uInt8Float32Array_valuesi[0];
        buff[offset + 1] = uInt8Float32Array_valuesi[1];
        buff[offset + 2] = uInt8Float32Array_valuesi[2];
        buff[offset + 3] = uInt8Float32Array_valuesi[3];
        offset += 4;
    }
    return offset;
};

ChannelFloat32.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_name = new TextDecoder('utf8');
    this.name = decoder_name.decode(buff.slice(offset, offset + length_name));
    offset += length_name;
    var length_values = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_values |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_values |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_values |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.values = new Array(length_values);
    for (var i = 0; i < length_values; i++) {
        var float32Array_valuesi = new Float32Array(1);
        var uInt8Float32Array_valuesi = new Uint8Array(float32Array_valuesi.buffer);
        uInt8Float32Array_valuesi[0] = buff[offset + 0];
        uInt8Float32Array_valuesi[1] = buff[offset + 1];
        uInt8Float32Array_valuesi[2] = buff[offset + 2];
        uInt8Float32Array_valuesi[3] = buff[offset + 3];
        this.values[i] = float32Array_valuesi[0];
        offset += 4;
    }
    return offset;
};

ChannelFloat32.prototype.serializedLength = function() {
    var length = 0;
    var encoder_name = new TextEncoder('utf8');
    var utf8array_name = encoder_name.encode(this.name);
    length += 4;
    length += utf8array_name.length;
    var length_values = this.values.length;
    length += 4;
    for (var i = 0; i < length_values; i++) {
        length += 4
    }
    return length;
};

ChannelFloat32.prototype.echo = function() { return ""; };

ChannelFloat32.prototype.getType = function() { return "sensor_msgs/ChannelFloat32"; };

ChannelFloat32.prototype.getMD5 = function() { return "c4cf01c81334c609dca1afd3a227daff"; };

ChannelFloat32.prototype.getID = function() { return 0; };

ChannelFloat32.prototype.setID = function(id) { };

return function() { return new ChannelFloat32(); };

});
