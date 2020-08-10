(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GetMapGoal=f();}})(function(){var define,module,exports;'use strict';

function GetMapGoal() {
};

GetMapGoal.prototype.serialize = function(buff, idx) {
    var offset = idx;
    return offset;
};

GetMapGoal.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    return offset;
};

GetMapGoal.prototype.serializedLength = function() {
    var length = 0;
    return length;
};

GetMapGoal.prototype.echo = function() { return ""; };

GetMapGoal.prototype.getType = function() { return "nav_msgs/GetMapGoal"; };

GetMapGoal.prototype.getMD5 = function() { return "b39e6b705afaad0184bd2c87f4bd870f"; };

GetMapGoal.prototype.getID = function() { return 0; };

GetMapGoal.prototype.setID = function(id) { };

return function() { return new GetMapGoal(); };

});
