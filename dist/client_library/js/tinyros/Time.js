(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tinyros==="undefined"){g.tinyros={};}g.tinyros.Time=f();}})(function(){var define,module,exports;'use strict';

function Time() {
    this.sec = 0;
    this.nsec = 0;
};

Time.dds_ms = 0;

Time.dds_start_ms = 0;

Time.dds_last_ms = 0;

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

Time.prototype.dds = function() {
    var time = this.now();
    var offset = time.toMSec();
    offset = offset > Time.dds_start_ms && Time.dds_start_ms > 0 ? offset - Time.dds_start_ms : 0;
    time.sec = +(offset / 1000);
    time.nsec = +((offset % 1000) * 1000000);
    time.sec = +(Time.dds_ms / 1000);
    time.nsec = +((Time.dds_ms % 1000) * 1000000);

    var nsec_part = +(time.nsec % 1000000000);
    var sec_part = +(time.nsec / 1000000000);
    time.sec += sec_part;
    time.nsec = nsec_part;
    return time;
};

Time.prototype.setdds_ms = function(ms) {
    Time.dds_ms = ms;
};

Time.prototype.getdds_ms = function() {
    return Time.dds_ms;
};

Time.prototype.setdds_last_ms = function(ms) {
    Time.dds_last_ms = ms;
};

Time.prototype.getdds_last_ms = function() {
    return Time.dds_last_ms;
};

Time.prototype.setdds_start_ms = function(ms) {
    Time.dds_start_ms = ms;
};

Time.prototype.getdds_start_ms = function() {
    return Time.dds_start_ms;
};

return function() { return new Time(); };

});
