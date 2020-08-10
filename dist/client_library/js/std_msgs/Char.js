(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Char=f();}})(function(){var define,module,exports;'use strict';

function Char() {
    this.data = ' ';
};

Char.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset] = this.data.charCodeAt() & 0xFF;
    offset += 1;
    return offset;
};

Char.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.data = String.fromCharCode(buff[offset]);
    offset += 1;
    return offset;
};

Char.prototype.serializedLength = function() {
    var length = 0;
    length += 1
    return length;
};

Char.prototype.echo = function() { return ""; };

Char.prototype.getType = function() { return "std_msgs/Char"; };

Char.prototype.getMD5 = function() { return "18741a80dcc02fcae20538073fcb0872"; };

Char.prototype.getID = function() { return 0; };

Char.prototype.setID = function(id) { };

return function() { return new Char(); };

});
