(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.UInt16=f();}})(function(){var define,module,exports;'use strict';

function UInt16() {
    this.data = 0;
};

UInt16.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.data) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.data) >> (8 * 1)) & 0xFF;
    offset += 2;
    return offset;
};

UInt16.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.data |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    offset += 2;
    return offset;
};

UInt16.prototype.serializedLength = function() {
    var length = 0;
    length += 2
    return length;
};

UInt16.prototype.echo = function() { return ""; };

UInt16.prototype.getType = function() { return "std_msgs/UInt16"; };

UInt16.prototype.getMD5 = function() { return "a4caad902eedb84b728d8369c50ffc39"; };

UInt16.prototype.getID = function() { return 0; };

UInt16.prototype.setID = function(id) { };

return function() { return new UInt16(); };

});
