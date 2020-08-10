(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.OccupancyGrid=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/MapMetaData.js'></script>");

function OccupancyGrid() {
    this.header = std_msgs.Header();
    this.info = nav_msgs.MapMetaData();
    this.data = new Array();
};

OccupancyGrid.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.info.serialize(buff, offset);
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

OccupancyGrid.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.info.deserialize(buff, offset);
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

OccupancyGrid.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.info.serializedLength();
    var length_data = this.data.length;
    length += 4;
    for (var i = 0; i < length_data; i++) {
        length += 1
    }
    return length;
};

OccupancyGrid.prototype.echo = function() { return ""; };

OccupancyGrid.prototype.getType = function() { return "nav_msgs/OccupancyGrid"; };

OccupancyGrid.prototype.getMD5 = function() { return "e489a26457224a97799696f3642f16a0"; };

OccupancyGrid.prototype.getID = function() { return 0; };

OccupancyGrid.prototype.setID = function(id) { };

return function() { return new OccupancyGrid(); };

});
