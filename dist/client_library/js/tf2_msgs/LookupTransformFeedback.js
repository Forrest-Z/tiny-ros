(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf2_msgs==="undefined"){g.tf2_msgs={};}g.tf2_msgs.LookupTransformFeedback=f();}})(function(){var define,module,exports;'use strict';

function LookupTransformFeedback() {
};

LookupTransformFeedback.prototype.serialize = function(buff, idx) {
    var offset = idx;
    return offset;
};

LookupTransformFeedback.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    return offset;
};

LookupTransformFeedback.prototype.serializedLength = function() {
    var length = 0;
    return length;
};

LookupTransformFeedback.prototype.echo = function() { return ""; };

LookupTransformFeedback.prototype.getType = function() { return "tf2_msgs/LookupTransformFeedback"; };

LookupTransformFeedback.prototype.getMD5 = function() { return "e6217f8a8e77aa218a8d6f594d08ba08"; };

LookupTransformFeedback.prototype.getID = function() { return 0; };

LookupTransformFeedback.prototype.setID = function(id) { };

return function() { return new LookupTransformFeedback(); };

});
