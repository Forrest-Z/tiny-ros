(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.UInt32=f();}})(function(){var define,module,exports;'use strict';

function UInt32() {
    this.data = 0;
};

UInt32.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.data) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.data) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.data) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.data) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

UInt32.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.data |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.data |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.data |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

UInt32.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

UInt32.prototype.echo = function() { return ""; };

UInt32.prototype.getType = function() { return "std_msgs/UInt32"; };

UInt32.prototype.getMD5 = function() { return "d4e8dc9c9e9d5076e594a3e342c2d4e3"; };

UInt32.prototype.getID = function() { return 0; };

UInt32.prototype.setID = function(id) { };

return function() { return new UInt32(); };

});
