(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Byte=f();}})(function(){var define,module,exports;'use strict';

function Byte() {
    this.data = 0;
};

Byte.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.data) >> (8 * 0)) & 0xFF;
    offset += 1;
    return offset;
};

Byte.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    return offset;
};

Byte.prototype.serializedLength = function() {
    var length = 0;
    length += 1
    return length;
};

Byte.prototype.echo = function() { return ""; };

Byte.prototype.getType = function() { return "std_msgs/Byte"; };

Byte.prototype.getMD5 = function() { return "8c5affe485c5af9bd37408a1a905a49c"; };

Byte.prototype.getID = function() { return 0; };

Byte.prototype.setID = function(id) { };

return function() { return new Byte(); };

});
