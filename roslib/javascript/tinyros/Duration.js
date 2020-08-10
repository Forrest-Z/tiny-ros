(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tinyros==="undefined"){g.tinyros={};}g.tinyros.Duration=f();}})(function(){var define,module,exports;'use strict';

function Duration() {
    this.sec = 0;
    this.nsec = 0;
};

Duration.prototype.toSec = function() {
    return (this.sec + this.nsec * 1e-9);
};

Duration.prototype.toMSec = function() {
    return (this.sec * 1000 + this.nsec * 1e-6);
};

Duration.prototype.toNSec = function() {
    return (this.sec * 1e9 + this.nsec);
};

return function() { return new Duration(); };

});

