(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.rosgraph_msgs==="undefined"){g.rosgraph_msgs={};}g.rosgraph_msgs.Clock=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Time.js'></script>");

function Clock() {
    this.clock = tinyros.Time();
};

Clock.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.clock.sec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.clock.sec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.clock.sec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.clock.sec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.clock.nsec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.clock.nsec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.clock.nsec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.clock.nsec) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

Clock.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.clock.sec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.clock.sec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.clock.sec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.clock.sec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.clock.nsec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.clock.nsec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.clock.nsec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.clock.nsec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

Clock.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 4
    return length;
};

Clock.prototype.echo = function() { return ""; };

Clock.prototype.getType = function() { return "rosgraph_msgs/Clock"; };

Clock.prototype.getMD5 = function() { return "d3bedbe03b904b8181e3fef4bbe0a73e"; };

Clock.prototype.getID = function() { return 0; };

Clock.prototype.setID = function(id) { };

return function() { return new Clock(); };

});
