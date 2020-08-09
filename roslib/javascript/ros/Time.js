(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.ros==="undefined"){g.ros={};}g.ros.Time=f();}})(function(){var define,module,exports;'use strict';

function Time() {
    this.sec = 0;
    this.nsec = 0;
};

Time.prototype.now = function() {
    var now = new Time();
    var ms = +Date.now();
    now.sec = +(ms/1000);
    now.nsec = +((ms % 1000) * 1000000);
    return now;
};

Time.prototype.toSec = function() {
    return (this.sec + this.nsec * 1e-9);
};

Time.prototype.toMSec = function() {
    return (this.sec * 1000 + this.nsec * 1e-6);
};

Time.prototype.toNSec = function() {
    return (this.sec * 1e9 + this.nsec);
};

Time.prototype.DDS = function() {
    return new Time();
};

return function() { return new Time(); };

});
