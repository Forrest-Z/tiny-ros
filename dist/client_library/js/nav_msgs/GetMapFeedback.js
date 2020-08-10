(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GetMapFeedback=f();}})(function(){var define,module,exports;'use strict';

function GetMapFeedback() {
};

GetMapFeedback.prototype.serialize = function(buff, idx) {
    var offset = idx;
    return offset;
};

GetMapFeedback.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    return offset;
};

GetMapFeedback.prototype.serializedLength = function() {
    var length = 0;
    return length;
};

GetMapFeedback.prototype.echo = function() { return ""; };

GetMapFeedback.prototype.getType = function() { return "nav_msgs/GetMapFeedback"; };

GetMapFeedback.prototype.getMD5 = function() { return "f561626803919fb2f269eb497bfdfea4"; };

GetMapFeedback.prototype.getID = function() { return 0; };

GetMapFeedback.prototype.setID = function(id) { };

return function() { return new GetMapFeedback(); };

});
