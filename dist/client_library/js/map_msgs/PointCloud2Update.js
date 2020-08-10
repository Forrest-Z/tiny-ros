(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.PointCloud2Update=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/sensor_msgs/PointCloud2.js'></script>");

function PointCloud2Update() {
    this.header = std_msgs.Header();
    this.type = 0;
    this.points = sensor_msgs.PointCloud2();

    // ENUM{
    this.ADD = 0;
    this.DELETE = 1;
    // }ENUM
};

PointCloud2Update.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    buff[offset + 0] = ((+this.type) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.type) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.type) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.type) >> (8 * 3)) & 0xFF;
    offset += 4;
    offset += this.points.serialize(buff, offset);
    return offset;
};

PointCloud2Update.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    this.type = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.type |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.type |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.type |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    offset += this.points.deserialize(buff, offset);
    return offset;
};

PointCloud2Update.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 4
    length += this.points.serializedLength();
    return length;
};

PointCloud2Update.prototype.echo = function() { return ""; };

PointCloud2Update.prototype.getType = function() { return "map_msgs/PointCloud2Update"; };

PointCloud2Update.prototype.getMD5 = function() { return "e79dfbefd7336861352e1bc7148491c4"; };

PointCloud2Update.prototype.getID = function() { return 0; };

PointCloud2Update.prototype.setID = function(id) { };

return function() { return new PointCloud2Update(); };

});
