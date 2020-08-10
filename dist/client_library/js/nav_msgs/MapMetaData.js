(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.MapMetaData=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Time.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Pose.js'></script>");

function MapMetaData() {
    this.map_load_time = tinyros.Time();
    this.resolution = 0.0;
    this.width = 0;
    this.height = 0;
    this.origin = geometry_msgs.Pose();
};

MapMetaData.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.map_load_time.sec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.map_load_time.sec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.map_load_time.sec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.map_load_time.sec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.map_load_time.nsec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.map_load_time.nsec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.map_load_time.nsec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.map_load_time.nsec) >> (8 * 3)) & 0xFF;
    offset += 4;
    var float32Array_resolution = new Float32Array(1);
    var uInt8Float32Array_resolution = new Uint8Array(float32Array_resolution.buffer);
    float32Array_resolution[0] = +this.resolution;
    buff[offset + 0] = uInt8Float32Array_resolution[0];
    buff[offset + 1] = uInt8Float32Array_resolution[1];
    buff[offset + 2] = uInt8Float32Array_resolution[2];
    buff[offset + 3] = uInt8Float32Array_resolution[3];
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
    offset += this.origin.serialize(buff, offset);
    return offset;
};

MapMetaData.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.map_load_time.sec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.map_load_time.sec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.map_load_time.sec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.map_load_time.sec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.map_load_time.nsec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.map_load_time.nsec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.map_load_time.nsec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.map_load_time.nsec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var float32Array_resolution = new Float32Array(1);
    var uInt8Float32Array_resolution = new Uint8Array(float32Array_resolution.buffer);
    uInt8Float32Array_resolution[0] = buff[offset + 0];
    uInt8Float32Array_resolution[1] = buff[offset + 1];
    uInt8Float32Array_resolution[2] = buff[offset + 2];
    uInt8Float32Array_resolution[3] = buff[offset + 3];
    this.resolution = float32Array_resolution[0];
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
    offset += this.origin.deserialize(buff, offset);
    return offset;
};

MapMetaData.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += this.origin.serializedLength();
    return length;
};

MapMetaData.prototype.echo = function() { return ""; };

MapMetaData.prototype.getType = function() { return "nav_msgs/MapMetaData"; };

MapMetaData.prototype.getMD5 = function() { return "328f5a1f2242fff4676d48189bd8b309"; };

MapMetaData.prototype.getID = function() { return 0; };

MapMetaData.prototype.setID = function(id) { };

return function() { return new MapMetaData(); };

});
