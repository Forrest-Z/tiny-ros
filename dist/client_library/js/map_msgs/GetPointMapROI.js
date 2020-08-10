(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.GetPointMapROIRequest=f();}})(function(){var define,module,exports;'use strict';

function GetPointMapROIRequest() {
    this.__id__ = 0;
    this.x = 0.0;
    this.y = 0.0;
    this.z = 0.0;
    this.r = 0.0;
    this.l_x = 0.0;
    this.l_y = 0.0;
    this.l_z = 0.0;
};

GetPointMapROIRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var float64Array_x = new Float64Array(1);
    var uInt8Float64Array_x = new Uint8Array(float64Array_x.buffer);
    float64Array_x[0] = +this.x;
    buff[offset + 0] = uInt8Float64Array_x[0];
    buff[offset + 1] = uInt8Float64Array_x[1];
    buff[offset + 2] = uInt8Float64Array_x[2];
    buff[offset + 3] = uInt8Float64Array_x[3];
    buff[offset + 4] = uInt8Float64Array_x[4];
    buff[offset + 5] = uInt8Float64Array_x[5];
    buff[offset + 6] = uInt8Float64Array_x[6];
    buff[offset + 7] = uInt8Float64Array_x[7];
    offset += 8;
    var float64Array_y = new Float64Array(1);
    var uInt8Float64Array_y = new Uint8Array(float64Array_y.buffer);
    float64Array_y[0] = +this.y;
    buff[offset + 0] = uInt8Float64Array_y[0];
    buff[offset + 1] = uInt8Float64Array_y[1];
    buff[offset + 2] = uInt8Float64Array_y[2];
    buff[offset + 3] = uInt8Float64Array_y[3];
    buff[offset + 4] = uInt8Float64Array_y[4];
    buff[offset + 5] = uInt8Float64Array_y[5];
    buff[offset + 6] = uInt8Float64Array_y[6];
    buff[offset + 7] = uInt8Float64Array_y[7];
    offset += 8;
    var float64Array_z = new Float64Array(1);
    var uInt8Float64Array_z = new Uint8Array(float64Array_z.buffer);
    float64Array_z[0] = +this.z;
    buff[offset + 0] = uInt8Float64Array_z[0];
    buff[offset + 1] = uInt8Float64Array_z[1];
    buff[offset + 2] = uInt8Float64Array_z[2];
    buff[offset + 3] = uInt8Float64Array_z[3];
    buff[offset + 4] = uInt8Float64Array_z[4];
    buff[offset + 5] = uInt8Float64Array_z[5];
    buff[offset + 6] = uInt8Float64Array_z[6];
    buff[offset + 7] = uInt8Float64Array_z[7];
    offset += 8;
    var float64Array_r = new Float64Array(1);
    var uInt8Float64Array_r = new Uint8Array(float64Array_r.buffer);
    float64Array_r[0] = +this.r;
    buff[offset + 0] = uInt8Float64Array_r[0];
    buff[offset + 1] = uInt8Float64Array_r[1];
    buff[offset + 2] = uInt8Float64Array_r[2];
    buff[offset + 3] = uInt8Float64Array_r[3];
    buff[offset + 4] = uInt8Float64Array_r[4];
    buff[offset + 5] = uInt8Float64Array_r[5];
    buff[offset + 6] = uInt8Float64Array_r[6];
    buff[offset + 7] = uInt8Float64Array_r[7];
    offset += 8;
    var float64Array_l_x = new Float64Array(1);
    var uInt8Float64Array_l_x = new Uint8Array(float64Array_l_x.buffer);
    float64Array_l_x[0] = +this.l_x;
    buff[offset + 0] = uInt8Float64Array_l_x[0];
    buff[offset + 1] = uInt8Float64Array_l_x[1];
    buff[offset + 2] = uInt8Float64Array_l_x[2];
    buff[offset + 3] = uInt8Float64Array_l_x[3];
    buff[offset + 4] = uInt8Float64Array_l_x[4];
    buff[offset + 5] = uInt8Float64Array_l_x[5];
    buff[offset + 6] = uInt8Float64Array_l_x[6];
    buff[offset + 7] = uInt8Float64Array_l_x[7];
    offset += 8;
    var float64Array_l_y = new Float64Array(1);
    var uInt8Float64Array_l_y = new Uint8Array(float64Array_l_y.buffer);
    float64Array_l_y[0] = +this.l_y;
    buff[offset + 0] = uInt8Float64Array_l_y[0];
    buff[offset + 1] = uInt8Float64Array_l_y[1];
    buff[offset + 2] = uInt8Float64Array_l_y[2];
    buff[offset + 3] = uInt8Float64Array_l_y[3];
    buff[offset + 4] = uInt8Float64Array_l_y[4];
    buff[offset + 5] = uInt8Float64Array_l_y[5];
    buff[offset + 6] = uInt8Float64Array_l_y[6];
    buff[offset + 7] = uInt8Float64Array_l_y[7];
    offset += 8;
    var float64Array_l_z = new Float64Array(1);
    var uInt8Float64Array_l_z = new Uint8Array(float64Array_l_z.buffer);
    float64Array_l_z[0] = +this.l_z;
    buff[offset + 0] = uInt8Float64Array_l_z[0];
    buff[offset + 1] = uInt8Float64Array_l_z[1];
    buff[offset + 2] = uInt8Float64Array_l_z[2];
    buff[offset + 3] = uInt8Float64Array_l_z[3];
    buff[offset + 4] = uInt8Float64Array_l_z[4];
    buff[offset + 5] = uInt8Float64Array_l_z[5];
    buff[offset + 6] = uInt8Float64Array_l_z[6];
    buff[offset + 7] = uInt8Float64Array_l_z[7];
    offset += 8;
    return offset;
};

GetPointMapROIRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var float64Array_x = new Float64Array(1);
    var uInt8Float64Array_x = new Uint8Array(float64Array_x.buffer);
    uInt8Float64Array_x[0] = buff[offset + 0];
    uInt8Float64Array_x[1] = buff[offset + 1];
    uInt8Float64Array_x[2] = buff[offset + 2];
    uInt8Float64Array_x[3] = buff[offset + 3];
    uInt8Float64Array_x[4] = buff[offset + 4];
    uInt8Float64Array_x[5] = buff[offset + 5];
    uInt8Float64Array_x[6] = buff[offset + 6];
    uInt8Float64Array_x[7] = buff[offset + 7];
    this.x = float64Array_x[0];
    offset += 8;
    var float64Array_y = new Float64Array(1);
    var uInt8Float64Array_y = new Uint8Array(float64Array_y.buffer);
    uInt8Float64Array_y[0] = buff[offset + 0];
    uInt8Float64Array_y[1] = buff[offset + 1];
    uInt8Float64Array_y[2] = buff[offset + 2];
    uInt8Float64Array_y[3] = buff[offset + 3];
    uInt8Float64Array_y[4] = buff[offset + 4];
    uInt8Float64Array_y[5] = buff[offset + 5];
    uInt8Float64Array_y[6] = buff[offset + 6];
    uInt8Float64Array_y[7] = buff[offset + 7];
    this.y = float64Array_y[0];
    offset += 8;
    var float64Array_z = new Float64Array(1);
    var uInt8Float64Array_z = new Uint8Array(float64Array_z.buffer);
    uInt8Float64Array_z[0] = buff[offset + 0];
    uInt8Float64Array_z[1] = buff[offset + 1];
    uInt8Float64Array_z[2] = buff[offset + 2];
    uInt8Float64Array_z[3] = buff[offset + 3];
    uInt8Float64Array_z[4] = buff[offset + 4];
    uInt8Float64Array_z[5] = buff[offset + 5];
    uInt8Float64Array_z[6] = buff[offset + 6];
    uInt8Float64Array_z[7] = buff[offset + 7];
    this.z = float64Array_z[0];
    offset += 8;
    var float64Array_r = new Float64Array(1);
    var uInt8Float64Array_r = new Uint8Array(float64Array_r.buffer);
    uInt8Float64Array_r[0] = buff[offset + 0];
    uInt8Float64Array_r[1] = buff[offset + 1];
    uInt8Float64Array_r[2] = buff[offset + 2];
    uInt8Float64Array_r[3] = buff[offset + 3];
    uInt8Float64Array_r[4] = buff[offset + 4];
    uInt8Float64Array_r[5] = buff[offset + 5];
    uInt8Float64Array_r[6] = buff[offset + 6];
    uInt8Float64Array_r[7] = buff[offset + 7];
    this.r = float64Array_r[0];
    offset += 8;
    var float64Array_l_x = new Float64Array(1);
    var uInt8Float64Array_l_x = new Uint8Array(float64Array_l_x.buffer);
    uInt8Float64Array_l_x[0] = buff[offset + 0];
    uInt8Float64Array_l_x[1] = buff[offset + 1];
    uInt8Float64Array_l_x[2] = buff[offset + 2];
    uInt8Float64Array_l_x[3] = buff[offset + 3];
    uInt8Float64Array_l_x[4] = buff[offset + 4];
    uInt8Float64Array_l_x[5] = buff[offset + 5];
    uInt8Float64Array_l_x[6] = buff[offset + 6];
    uInt8Float64Array_l_x[7] = buff[offset + 7];
    this.l_x = float64Array_l_x[0];
    offset += 8;
    var float64Array_l_y = new Float64Array(1);
    var uInt8Float64Array_l_y = new Uint8Array(float64Array_l_y.buffer);
    uInt8Float64Array_l_y[0] = buff[offset + 0];
    uInt8Float64Array_l_y[1] = buff[offset + 1];
    uInt8Float64Array_l_y[2] = buff[offset + 2];
    uInt8Float64Array_l_y[3] = buff[offset + 3];
    uInt8Float64Array_l_y[4] = buff[offset + 4];
    uInt8Float64Array_l_y[5] = buff[offset + 5];
    uInt8Float64Array_l_y[6] = buff[offset + 6];
    uInt8Float64Array_l_y[7] = buff[offset + 7];
    this.l_y = float64Array_l_y[0];
    offset += 8;
    var float64Array_l_z = new Float64Array(1);
    var uInt8Float64Array_l_z = new Uint8Array(float64Array_l_z.buffer);
    uInt8Float64Array_l_z[0] = buff[offset + 0];
    uInt8Float64Array_l_z[1] = buff[offset + 1];
    uInt8Float64Array_l_z[2] = buff[offset + 2];
    uInt8Float64Array_l_z[3] = buff[offset + 3];
    uInt8Float64Array_l_z[4] = buff[offset + 4];
    uInt8Float64Array_l_z[5] = buff[offset + 5];
    uInt8Float64Array_l_z[6] = buff[offset + 6];
    uInt8Float64Array_l_z[7] = buff[offset + 7];
    this.l_z = float64Array_l_z[0];
    offset += 8;
    return offset;
};

GetPointMapROIRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    return length;
};

GetPointMapROIRequest.prototype.echo = function() { return ""; };

GetPointMapROIRequest.prototype.getType = function() { return "map_msgs/GetPointMapROI"; };

GetPointMapROIRequest.prototype.getMD5 = function() { return "c338f5616930e00a72c38486f77e9810"; };

GetPointMapROIRequest.prototype.getID = function() { return this.__id__; };

GetPointMapROIRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetPointMapROIRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.GetPointMapROIResponse=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/sensor_msgs/PointCloud2.js'></script>");

function GetPointMapROIResponse() {
    this.__id__ = 0;
    this.sub_map = sensor_msgs.PointCloud2();
};

GetPointMapROIResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    offset += this.sub_map.serialize(buff, offset);
    return offset;
};

GetPointMapROIResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    offset += this.sub_map.deserialize(buff, offset);
    return offset;
};

GetPointMapROIResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += this.sub_map.serializedLength();
    return length;
};

GetPointMapROIResponse.prototype.echo = function() { return ""; };

GetPointMapROIResponse.prototype.getType = function() { return "map_msgs/GetPointMapROI"; };

GetPointMapROIResponse.prototype.getMD5 = function() { return "de10fb68f23987142a0ffbdb59b6e079"; };

GetPointMapROIResponse.prototype.getID = function() { return this.__id__; };

GetPointMapROIResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetPointMapROIResponse(); };

});
