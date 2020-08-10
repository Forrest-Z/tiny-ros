(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Int64=f();}})(function(){var define,module,exports;'use strict';

function Int64() {
    this.data = 0;
};

Int64.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.data) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.data) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.data) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.data) >> (8 * 3)) & 0xFF;
    buff[offset + 4] = ((+this.data) >> (8 * 4)) & 0xFF;
    buff[offset + 5] = ((+this.data) >> (8 * 5)) & 0xFF;
    buff[offset + 6] = ((+this.data) >> (8 * 6)) & 0xFF;
    buff[offset + 7] = ((+this.data) >> (8 * 7)) & 0xFF;
    offset += 8;
    return offset;
};

Int64.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.data |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.data |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.data |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    this.data |= +((buff[offset + 4] & 0xFF) << (8 * 4));
    this.data |= +((buff[offset + 5] & 0xFF) << (8 * 5));
    this.data |= +((buff[offset + 6] & 0xFF) << (8 * 6));
    this.data |= +((buff[offset + 7] & 0xFF) << (8 * 7));
    offset += 8;
    return offset;
};

Int64.prototype.serializedLength = function() {
    var length = 0;
    length += 8
    return length;
};

Int64.prototype.echo = function() { return ""; };

Int64.prototype.getType = function() { return "std_msgs/Int64"; };

Int64.prototype.getMD5 = function() { return "3dd57e1dfe8b2ebcc4764deb44bda243"; };

Int64.prototype.getID = function() { return 0; };

Int64.prototype.setID = function(id) { };

return function() { return new Int64(); };

});
