(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tinyros_hello==="undefined"){g.tinyros_hello={};}g.tinyros_hello.TinyrosHello=f();}})(function(){var define,module,exports;'use strict';

function TinyrosHello() {
    this.hello = "";
};

TinyrosHello.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var encoder_hello = new TextEncoder('utf8');
    var utf8array_hello = encoder_hello.encode(this.hello);
    buff[offset + 0] = (utf8array_hello.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_hello.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_hello.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_hello.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_hello.length; i++) {
        buff[offset + i] = utf8array_hello[i];
    }
    offset += utf8array_hello.length;
    return offset;
};

TinyrosHello.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_hello = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_hello |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_hello |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_hello |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_hello = new TextDecoder('utf8');
    this.hello = decoder_hello.decode(buff.slice(offset, offset + length_hello));
    offset += length_hello;
    return offset;
};

TinyrosHello.prototype.serializedLength = function() {
    var length = 0;
    var encoder_hello = new TextEncoder('utf8');
    var utf8array_hello = encoder_hello.encode(this.hello);
    length += 4;
    length += utf8array_hello.length;
    return length;
};

TinyrosHello.prototype.echo = function() { return ""; };

TinyrosHello.prototype.getType = function() { return "tinyros_hello/TinyrosHello"; };

TinyrosHello.prototype.getMD5 = function() { return "0c68e66a217802ad0c9f648b7a69d090"; };

TinyrosHello.prototype.getID = function() { return 0; };

TinyrosHello.prototype.setID = function(id) { };

return function() { return new TinyrosHello(); };

});
