(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Int8=f();}})(function(){var define,module,exports;'use strict';

function Int8() {
    this.data = 0;
};

Int8.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.data) >> (8 * 0)) & 0xFF;
    offset += 1;
    return offset;
};

Int8.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    return offset;
};

Int8.prototype.serializedLength = function() {
    var length = 0;
    length += 1
    return length;
};

Int8.prototype.echo = function() { return ""; };

Int8.prototype.getType = function() { return "std_msgs/Int8"; };

Int8.prototype.getMD5 = function() { return "36c967b84bec6cd7c260bffc7f4dfbe0"; };

Int8.prototype.getID = function() { return 0; };

Int8.prototype.setID = function(id) { };

return function() { return new Int8(); };

});
