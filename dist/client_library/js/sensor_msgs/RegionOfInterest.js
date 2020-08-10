(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.RegionOfInterest=f();}})(function(){var define,module,exports;'use strict';

function RegionOfInterest() {
    this.x_offset = 0;
    this.y_offset = 0;
    this.height = 0;
    this.width = 0;
    this.do_rectify = false;
};

RegionOfInterest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.x_offset) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.x_offset) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.x_offset) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.x_offset) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.y_offset) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.y_offset) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.y_offset) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.y_offset) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.height) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.height) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.height) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.height) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.width) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.width) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.width) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.width) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset] = this.do_rectify === false ? 0 : 1;
    offset += 1;
    return offset;
};

RegionOfInterest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.x_offset = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.x_offset |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.x_offset |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.x_offset |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.y_offset = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.y_offset |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.y_offset |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.y_offset |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.height = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.height |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.height |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.height |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.width = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.width |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.width |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.width |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.do_rectify = buff[offset] !== 0 ? true : false;
    offset += 1;
    return offset;
};

RegionOfInterest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 4
    length += 4
    length += 4
    length += 1
    return length;
};

RegionOfInterest.prototype.echo = function() { return ""; };

RegionOfInterest.prototype.getType = function() { return "sensor_msgs/RegionOfInterest"; };

RegionOfInterest.prototype.getMD5 = function() { return "8370dc286f915405c906299aef5bb442"; };

RegionOfInterest.prototype.getID = function() { return 0; };

RegionOfInterest.prototype.setID = function(id) { };

return function() { return new RegionOfInterest(); };

});
