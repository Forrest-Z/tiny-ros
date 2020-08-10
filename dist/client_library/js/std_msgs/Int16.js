(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Int16=f();}})(function(){var define,module,exports;'use strict';

function Int16() {
    this.data = 0;
};

Int16.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.data) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.data) >> (8 * 1)) & 0xFF;
    offset += 2;
    return offset;
};

Int16.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.data |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    offset += 2;
    return offset;
};

Int16.prototype.serializedLength = function() {
    var length = 0;
    length += 2
    return length;
};

Int16.prototype.echo = function() { return ""; };

Int16.prototype.getType = function() { return "std_msgs/Int16"; };

Int16.prototype.getMD5 = function() { return "156735359f7b6d347f113a1090134af3"; };

Int16.prototype.getID = function() { return 0; };

Int16.prototype.setID = function(id) { };

return function() { return new Int16(); };

});
