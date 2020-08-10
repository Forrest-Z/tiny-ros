(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.ProjectedMap=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/OccupancyGrid.js'></script>");

function ProjectedMap() {
    this.map = nav_msgs.OccupancyGrid();
    this.min_z = 0.0;
    this.max_z = 0.0;
};

ProjectedMap.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.map.serialize(buff, offset);
    var float64Array_min_z = new Float64Array(1);
    var uInt8Float64Array_min_z = new Uint8Array(float64Array_min_z.buffer);
    float64Array_min_z[0] = +this.min_z;
    buff[offset + 0] = uInt8Float64Array_min_z[0];
    buff[offset + 1] = uInt8Float64Array_min_z[1];
    buff[offset + 2] = uInt8Float64Array_min_z[2];
    buff[offset + 3] = uInt8Float64Array_min_z[3];
    buff[offset + 4] = uInt8Float64Array_min_z[4];
    buff[offset + 5] = uInt8Float64Array_min_z[5];
    buff[offset + 6] = uInt8Float64Array_min_z[6];
    buff[offset + 7] = uInt8Float64Array_min_z[7];
    offset += 8;
    var float64Array_max_z = new Float64Array(1);
    var uInt8Float64Array_max_z = new Uint8Array(float64Array_max_z.buffer);
    float64Array_max_z[0] = +this.max_z;
    buff[offset + 0] = uInt8Float64Array_max_z[0];
    buff[offset + 1] = uInt8Float64Array_max_z[1];
    buff[offset + 2] = uInt8Float64Array_max_z[2];
    buff[offset + 3] = uInt8Float64Array_max_z[3];
    buff[offset + 4] = uInt8Float64Array_max_z[4];
    buff[offset + 5] = uInt8Float64Array_max_z[5];
    buff[offset + 6] = uInt8Float64Array_max_z[6];
    buff[offset + 7] = uInt8Float64Array_max_z[7];
    offset += 8;
    return offset;
};

ProjectedMap.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.map.deserialize(buff, offset);
    var float64Array_min_z = new Float64Array(1);
    var uInt8Float64Array_min_z = new Uint8Array(float64Array_min_z.buffer);
    uInt8Float64Array_min_z[0] = buff[offset + 0];
    uInt8Float64Array_min_z[1] = buff[offset + 1];
    uInt8Float64Array_min_z[2] = buff[offset + 2];
    uInt8Float64Array_min_z[3] = buff[offset + 3];
    uInt8Float64Array_min_z[4] = buff[offset + 4];
    uInt8Float64Array_min_z[5] = buff[offset + 5];
    uInt8Float64Array_min_z[6] = buff[offset + 6];
    uInt8Float64Array_min_z[7] = buff[offset + 7];
    this.min_z = float64Array_min_z[0];
    offset += 8;
    var float64Array_max_z = new Float64Array(1);
    var uInt8Float64Array_max_z = new Uint8Array(float64Array_max_z.buffer);
    uInt8Float64Array_max_z[0] = buff[offset + 0];
    uInt8Float64Array_max_z[1] = buff[offset + 1];
    uInt8Float64Array_max_z[2] = buff[offset + 2];
    uInt8Float64Array_max_z[3] = buff[offset + 3];
    uInt8Float64Array_max_z[4] = buff[offset + 4];
    uInt8Float64Array_max_z[5] = buff[offset + 5];
    uInt8Float64Array_max_z[6] = buff[offset + 6];
    uInt8Float64Array_max_z[7] = buff[offset + 7];
    this.max_z = float64Array_max_z[0];
    offset += 8;
    return offset;
};

ProjectedMap.prototype.serializedLength = function() {
    var length = 0;
    length += this.map.serializedLength();
    length += 8
    length += 8
    return length;
};

ProjectedMap.prototype.echo = function() { return ""; };

ProjectedMap.prototype.getType = function() { return "map_msgs/ProjectedMap"; };

ProjectedMap.prototype.getMD5 = function() { return "cbd5598c259cc16f5aa07335587a7367"; };

ProjectedMap.prototype.getID = function() { return 0; };

ProjectedMap.prototype.setID = function(id) { };

return function() { return new ProjectedMap(); };

});
