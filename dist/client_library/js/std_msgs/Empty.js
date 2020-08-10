(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Empty=f();}})(function(){var define,module,exports;'use strict';

function Empty() {
};

Empty.prototype.serialize = function(buff, idx) {
    var offset = idx;
    return offset;
};

Empty.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    return offset;
};

Empty.prototype.serializedLength = function() {
    var length = 0;
    return length;
};

Empty.prototype.echo = function() { return ""; };

Empty.prototype.getType = function() { return "std_msgs/Empty"; };

Empty.prototype.getMD5 = function() { return "140bdcb7bbc50b4a936e90ff2350c8d3"; };

Empty.prototype.getID = function() { return 0; };

Empty.prototype.setID = function(id) { };

return function() { return new Empty(); };

});
