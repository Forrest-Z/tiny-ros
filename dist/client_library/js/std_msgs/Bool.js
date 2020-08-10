(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Bool=f();}})(function(){var define,module,exports;'use strict';

function Bool() {
    this.data = false;
};

Bool.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset] = this.data === false ? 0 : 1;
    offset += 1;
    return offset;
};

Bool.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.data = buff[offset] !== 0 ? true : false;
    offset += 1;
    return offset;
};

Bool.prototype.serializedLength = function() {
    var length = 0;
    length += 1
    return length;
};

Bool.prototype.echo = function() { return ""; };

Bool.prototype.getType = function() { return "std_msgs/Bool"; };

Bool.prototype.getMD5 = function() { return "cf6f397ea93618cea833f64b7eef203e"; };

Bool.prototype.getID = function() { return 0; };

Bool.prototype.setID = function(id) { };

return function() { return new Bool(); };

});
