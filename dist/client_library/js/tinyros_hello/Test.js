(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tinyros_hello==="undefined"){g.tinyros_hello={};}g.tinyros_hello.TestRequest=f();}})(function(){var define,module,exports;'use strict';

function TestRequest() {
    this.__id__ = 0;
    this.input = "";
};

TestRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_input = new TextEncoder('utf8');
    var utf8array_input = encoder_input.encode(this.input);
    buff[offset + 0] = (utf8array_input.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_input.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_input.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_input.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_input.length; i++) {
        buff[offset + i] = utf8array_input[i];
    }
    offset += utf8array_input.length;
    return offset;
};

TestRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_input = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_input |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_input |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_input |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_input = new TextDecoder('utf8');
    this.input = decoder_input.decode(buff.slice(offset, offset + length_input));
    offset += length_input;
    return offset;
};

TestRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_input = new TextEncoder('utf8');
    var utf8array_input = encoder_input.encode(this.input);
    length += 4;
    length += utf8array_input.length;
    return length;
};

TestRequest.prototype.echo = function() { return ""; };

TestRequest.prototype.getType = function() { return "tinyros_hello/Test"; };

TestRequest.prototype.getMD5 = function() { return "26ee7a44335f1f7b55a5a7490460807d"; };

TestRequest.prototype.getID = function() { return this.__id__; };

TestRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new TestRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tinyros_hello==="undefined"){g.tinyros_hello={};}g.tinyros_hello.TestResponse=f();}})(function(){var define,module,exports;'use strict';

function TestResponse() {
    this.__id__ = 0;
    this.output = "";
};

TestResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_output = new TextEncoder('utf8');
    var utf8array_output = encoder_output.encode(this.output);
    buff[offset + 0] = (utf8array_output.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_output.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_output.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_output.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_output.length; i++) {
        buff[offset + i] = utf8array_output[i];
    }
    offset += utf8array_output.length;
    return offset;
};

TestResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_output = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_output |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_output |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_output |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_output = new TextDecoder('utf8');
    this.output = decoder_output.decode(buff.slice(offset, offset + length_output));
    offset += length_output;
    return offset;
};

TestResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_output = new TextEncoder('utf8');
    var utf8array_output = encoder_output.encode(this.output);
    length += 4;
    length += utf8array_output.length;
    return length;
};

TestResponse.prototype.echo = function() { return ""; };

TestResponse.prototype.getType = function() { return "tinyros_hello/Test"; };

TestResponse.prototype.getMD5 = function() { return "e2f6296e6ea9df7406f4fac9fb52d44b"; };

TestResponse.prototype.getID = function() { return this.__id__; };

TestResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new TestResponse(); };

});
