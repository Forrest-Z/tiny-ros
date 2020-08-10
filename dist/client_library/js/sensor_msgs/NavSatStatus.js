(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.NavSatStatus=f();}})(function(){var define,module,exports;'use strict';

function NavSatStatus() {
    this.status = 0;
    this.service = 0;

    // ENUM{
    this.STATUS_NO_FIX =   -1        ;
    this.STATUS_FIX =       0        ;
    this.STATUS_SBAS_FIX =  1        ;
    this.STATUS_GBAS_FIX =  2        ;
    this.SERVICE_GPS =      1;
    this.SERVICE_GLONASS =  2;
    this.SERVICE_COMPASS =  4      ;
    this.SERVICE_GALILEO =  8;
    // }ENUM
};

NavSatStatus.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.status) >> (8 * 0)) & 0xFF;
    offset += 1;
    buff[offset + 0] = ((+this.service) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.service) >> (8 * 1)) & 0xFF;
    offset += 2;
    return offset;
};

NavSatStatus.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.status = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    this.service = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.service |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    offset += 2;
    return offset;
};

NavSatStatus.prototype.serializedLength = function() {
    var length = 0;
    length += 1
    length += 2
    return length;
};

NavSatStatus.prototype.echo = function() { return ""; };

NavSatStatus.prototype.getType = function() { return "sensor_msgs/NavSatStatus"; };

NavSatStatus.prototype.getMD5 = function() { return "85ed59cf80532c1c15a2cf735d06279b"; };

NavSatStatus.prototype.getID = function() { return 0; };

NavSatStatus.prototype.setID = function(id) { };

return function() { return new NavSatStatus(); };

});
