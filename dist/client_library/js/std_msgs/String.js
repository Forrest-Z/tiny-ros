(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.String=f();}})(function(){var define,module,exports;'use strict';

function String() {
    this.data = "";
};

String.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var encoder_data = new TextEncoder('utf8');
    var utf8array_data = encoder_data.encode(this.data);
    buff[offset + 0] = (utf8array_data.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_data.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_data.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_data.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_data.length; i++) {
        buff[offset + i] = utf8array_data[i];
    }
    offset += utf8array_data.length;
    return offset;
};

String.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_data |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_data |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_data |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_data = new TextDecoder('utf8');
    this.data = decoder_data.decode(buff.slice(offset, offset + length_data));
    offset += length_data;
    return offset;
};

String.prototype.serializedLength = function() {
    var length = 0;
    var encoder_data = new TextEncoder('utf8');
    var utf8array_data = encoder_data.encode(this.data);
    length += 4;
    length += utf8array_data.length;
    return length;
};

String.prototype.echo = function() { return ""; };

String.prototype.getType = function() { return "std_msgs/String"; };

String.prototype.getMD5 = function() { return "44b822292b4a9ed05e241aa225458f97"; };

String.prototype.getID = function() { return 0; };

String.prototype.setID = function(id) { };

return function() { return new String(); };

});
