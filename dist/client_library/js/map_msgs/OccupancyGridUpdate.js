(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.OccupancyGridUpdate=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function OccupancyGridUpdate() {
    this.header = std_msgs.Header();
    this.x = 0;
    this.y = 0;
    this.width = 0;
    this.height = 0;
    this.data = new Array();
};

OccupancyGridUpdate.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    buff[offset + 0] = ((+this.x) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.x) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.x) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.x) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.y) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.y) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.y) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.y) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.width) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.width) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.width) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.width) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.height) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.height) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.height) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.height) >> (8 * 3)) & 0xFF;
    offset += 4;
    var length_data = this.data.length;
    buff[offset + 0] = (length_data >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_data >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_data >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_data >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_data; i++) {
        buff[offset + 0] = ((+this.data[i]) >> (8 * 0)) & 0xFF;
        offset += 1;
    }
    return offset;
};

OccupancyGridUpdate.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    this.x = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.x |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.x |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.x |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.y = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.y |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.y |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.y |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.width = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.width |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.width |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.width |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.height = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.height |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.height |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.height |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_data |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_data |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_data |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.data = new Array(length_data);
    for (var i = 0; i < length_data; i++) {
        this.data[i] = +((buff[offset + 0] & 0xFF) << (8 * 0));
        offset += 1;
    }
    return offset;
};

OccupancyGridUpdate.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 4
    length += 4
    length += 4
    length += 4
    var length_data = this.data.length;
    length += 4;
    for (var i = 0; i < length_data; i++) {
        length += 1
    }
    return length;
};

OccupancyGridUpdate.prototype.echo = function() { return ""; };

OccupancyGridUpdate.prototype.getType = function() { return "map_msgs/OccupancyGridUpdate"; };

OccupancyGridUpdate.prototype.getMD5 = function() { return "159b2d7856932f2e2cad9b082ed99ec2"; };

OccupancyGridUpdate.prototype.getID = function() { return 0; };

OccupancyGridUpdate.prototype.setID = function(id) { };

return function() { return new OccupancyGridUpdate(); };

});
