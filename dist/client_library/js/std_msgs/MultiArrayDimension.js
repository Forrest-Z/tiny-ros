(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.MultiArrayDimension=f();}})(function(){var define,module,exports;'use strict';

function MultiArrayDimension() {
    this.label = "";
    this.size = 0;
    this.stride = 0;
};

MultiArrayDimension.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var encoder_label = new TextEncoder('utf8');
    var utf8array_label = encoder_label.encode(this.label);
    buff[offset + 0] = (utf8array_label.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_label.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_label.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_label.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_label.length; i++) {
        buff[offset + i] = utf8array_label[i];
    }
    offset += utf8array_label.length;
    buff[offset + 0] = ((+this.size) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.size) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.size) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.size) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.stride) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.stride) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.stride) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.stride) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

MultiArrayDimension.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_label = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_label |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_label |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_label |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_label = new TextDecoder('utf8');
    this.label = decoder_label.decode(buff.slice(offset, offset + length_label));
    offset += length_label;
    this.size = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.size |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.size |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.size |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.stride = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.stride |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.stride |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.stride |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

MultiArrayDimension.prototype.serializedLength = function() {
    var length = 0;
    var encoder_label = new TextEncoder('utf8');
    var utf8array_label = encoder_label.encode(this.label);
    length += 4;
    length += utf8array_label.length;
    length += 4
    length += 4
    return length;
};

MultiArrayDimension.prototype.echo = function() { return ""; };

MultiArrayDimension.prototype.getType = function() { return "std_msgs/MultiArrayDimension"; };

MultiArrayDimension.prototype.getMD5 = function() { return "c2aacf83d49c7aa4a8504bd32158e990"; };

MultiArrayDimension.prototype.getID = function() { return 0; };

MultiArrayDimension.prototype.setID = function(id) { };

return function() { return new MultiArrayDimension(); };

});
