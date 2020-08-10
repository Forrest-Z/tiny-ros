(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.diagnostic_msgs==="undefined"){g.diagnostic_msgs={};}g.diagnostic_msgs.KeyValue=f();}})(function(){var define,module,exports;'use strict';

function KeyValue() {
    this.key = "";
    this.value = "";
};

KeyValue.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var encoder_key = new TextEncoder('utf8');
    var utf8array_key = encoder_key.encode(this.key);
    buff[offset + 0] = (utf8array_key.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_key.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_key.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_key.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_key.length; i++) {
        buff[offset + i] = utf8array_key[i];
    }
    offset += utf8array_key.length;
    var encoder_value = new TextEncoder('utf8');
    var utf8array_value = encoder_value.encode(this.value);
    buff[offset + 0] = (utf8array_value.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_value.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_value.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_value.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_value.length; i++) {
        buff[offset + i] = utf8array_value[i];
    }
    offset += utf8array_value.length;
    return offset;
};

KeyValue.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_key = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_key |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_key |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_key |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_key = new TextDecoder('utf8');
    this.key = decoder_key.decode(buff.slice(offset, offset + length_key));
    offset += length_key;
    var length_value = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_value |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_value |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_value |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_value = new TextDecoder('utf8');
    this.value = decoder_value.decode(buff.slice(offset, offset + length_value));
    offset += length_value;
    return offset;
};

KeyValue.prototype.serializedLength = function() {
    var length = 0;
    var encoder_key = new TextEncoder('utf8');
    var utf8array_key = encoder_key.encode(this.key);
    length += 4;
    length += utf8array_key.length;
    var encoder_value = new TextEncoder('utf8');
    var utf8array_value = encoder_value.encode(this.value);
    length += 4;
    length += utf8array_value.length;
    return length;
};

KeyValue.prototype.echo = function() { return ""; };

KeyValue.prototype.getType = function() { return "diagnostic_msgs/KeyValue"; };

KeyValue.prototype.getMD5 = function() { return "1baa904b80c685c77d1a42a872ca1d07"; };

KeyValue.prototype.getID = function() { return 0; };

KeyValue.prototype.setID = function(id) { };

return function() { return new KeyValue(); };

});
