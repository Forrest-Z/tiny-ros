(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf2_msgs==="undefined"){g.tf2_msgs={};}g.tf2_msgs.TF2Error=f();}})(function(){var define,module,exports;'use strict';

function TF2Error() {
    this.error = 0;
    this.error_string = "";

    // ENUM{
    this.NO_ERROR =  0;
    this.LOOKUP_ERROR =  1;
    this.CONNECTIVITY_ERROR =  2;
    this.EXTRAPOLATION_ERROR =  3;
    this.INVALID_ARGUMENT_ERROR =  4;
    this.TIMEOUT_ERROR =  5;
    this.TRANSFORM_ERROR =  6;
    // }ENUM
};

TF2Error.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.error) >> (8 * 0)) & 0xFF;
    offset += 1;
    var encoder_error_string = new TextEncoder('utf8');
    var utf8array_error_string = encoder_error_string.encode(this.error_string);
    buff[offset + 0] = (utf8array_error_string.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_error_string.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_error_string.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_error_string.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_error_string.length; i++) {
        buff[offset + i] = utf8array_error_string[i];
    }
    offset += utf8array_error_string.length;
    return offset;
};

TF2Error.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.error = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    var length_error_string = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_error_string |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_error_string |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_error_string |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_error_string = new TextDecoder('utf8');
    this.error_string = decoder_error_string.decode(buff.slice(offset, offset + length_error_string));
    offset += length_error_string;
    return offset;
};

TF2Error.prototype.serializedLength = function() {
    var length = 0;
    length += 1
    var encoder_error_string = new TextEncoder('utf8');
    var utf8array_error_string = encoder_error_string.encode(this.error_string);
    length += 4;
    length += utf8array_error_string.length;
    return length;
};

TF2Error.prototype.echo = function() { return ""; };

TF2Error.prototype.getType = function() { return "tf2_msgs/TF2Error"; };

TF2Error.prototype.getMD5 = function() { return "ed32adf5a372962d977aea0e5630d1d6"; };

TF2Error.prototype.getID = function() { return 0; };

TF2Error.prototype.setID = function(id) { };

return function() { return new TF2Error(); };

});
