(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_srvs==="undefined"){g.std_srvs={};}g.std_srvs.SetBoolRequest=f();}})(function(){var define,module,exports;'use strict';

function SetBoolRequest() {
    this.__id__ = 0;
    this.data = false;
};

SetBoolRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset] = this.data === false ? 0 : 1;
    offset += 1;
    return offset;
};

SetBoolRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.data = buff[offset] !== 0 ? true : false;
    offset += 1;
    return offset;
};

SetBoolRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    return length;
};

SetBoolRequest.prototype.echo = function() { return ""; };

SetBoolRequest.prototype.getType = function() { return "std_srvs/SetBool"; };

SetBoolRequest.prototype.getMD5 = function() { return "2600271ce244b6b0d236894ec6f04373"; };

SetBoolRequest.prototype.getID = function() { return this.__id__; };

SetBoolRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetBoolRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_srvs==="undefined"){g.std_srvs={};}g.std_srvs.SetBoolResponse=f();}})(function(){var define,module,exports;'use strict';

function SetBoolResponse() {
    this.__id__ = 0;
    this.success = false;
    this.message = "";
};

SetBoolResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset] = this.success === false ? 0 : 1;
    offset += 1;
    var encoder_message = new TextEncoder('utf8');
    var utf8array_message = encoder_message.encode(this.message);
    buff[offset + 0] = (utf8array_message.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_message.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_message.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_message.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_message.length; i++) {
        buff[offset + i] = utf8array_message[i];
    }
    offset += utf8array_message.length;
    return offset;
};

SetBoolResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.success = buff[offset] !== 0 ? true : false;
    offset += 1;
    var length_message = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_message |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_message |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_message |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_message = new TextDecoder('utf8');
    this.message = decoder_message.decode(buff.slice(offset, offset + length_message));
    offset += length_message;
    return offset;
};

SetBoolResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    var encoder_message = new TextEncoder('utf8');
    var utf8array_message = encoder_message.encode(this.message);
    length += 4;
    length += utf8array_message.length;
    return length;
};

SetBoolResponse.prototype.echo = function() { return ""; };

SetBoolResponse.prototype.getType = function() { return "std_srvs/SetBool"; };

SetBoolResponse.prototype.getMD5 = function() { return "51cf1b5cb67d107350442299d694fd1d"; };

SetBoolResponse.prototype.getID = function() { return this.__id__; };

SetBoolResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetBoolResponse(); };

});
