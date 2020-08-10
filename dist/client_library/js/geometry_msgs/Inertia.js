(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.Inertia=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Vector3.js'></script>");

function Inertia() {
    this.m = 0.0;
    this.com = geometry_msgs.Vector3();
    this.ixx = 0.0;
    this.ixy = 0.0;
    this.ixz = 0.0;
    this.iyy = 0.0;
    this.iyz = 0.0;
    this.izz = 0.0;
};

Inertia.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var float64Array_m = new Float64Array(1);
    var uInt8Float64Array_m = new Uint8Array(float64Array_m.buffer);
    float64Array_m[0] = +this.m;
    buff[offset + 0] = uInt8Float64Array_m[0];
    buff[offset + 1] = uInt8Float64Array_m[1];
    buff[offset + 2] = uInt8Float64Array_m[2];
    buff[offset + 3] = uInt8Float64Array_m[3];
    buff[offset + 4] = uInt8Float64Array_m[4];
    buff[offset + 5] = uInt8Float64Array_m[5];
    buff[offset + 6] = uInt8Float64Array_m[6];
    buff[offset + 7] = uInt8Float64Array_m[7];
    offset += 8;
    offset += this.com.serialize(buff, offset);
    var float64Array_ixx = new Float64Array(1);
    var uInt8Float64Array_ixx = new Uint8Array(float64Array_ixx.buffer);
    float64Array_ixx[0] = +this.ixx;
    buff[offset + 0] = uInt8Float64Array_ixx[0];
    buff[offset + 1] = uInt8Float64Array_ixx[1];
    buff[offset + 2] = uInt8Float64Array_ixx[2];
    buff[offset + 3] = uInt8Float64Array_ixx[3];
    buff[offset + 4] = uInt8Float64Array_ixx[4];
    buff[offset + 5] = uInt8Float64Array_ixx[5];
    buff[offset + 6] = uInt8Float64Array_ixx[6];
    buff[offset + 7] = uInt8Float64Array_ixx[7];
    offset += 8;
    var float64Array_ixy = new Float64Array(1);
    var uInt8Float64Array_ixy = new Uint8Array(float64Array_ixy.buffer);
    float64Array_ixy[0] = +this.ixy;
    buff[offset + 0] = uInt8Float64Array_ixy[0];
    buff[offset + 1] = uInt8Float64Array_ixy[1];
    buff[offset + 2] = uInt8Float64Array_ixy[2];
    buff[offset + 3] = uInt8Float64Array_ixy[3];
    buff[offset + 4] = uInt8Float64Array_ixy[4];
    buff[offset + 5] = uInt8Float64Array_ixy[5];
    buff[offset + 6] = uInt8Float64Array_ixy[6];
    buff[offset + 7] = uInt8Float64Array_ixy[7];
    offset += 8;
    var float64Array_ixz = new Float64Array(1);
    var uInt8Float64Array_ixz = new Uint8Array(float64Array_ixz.buffer);
    float64Array_ixz[0] = +this.ixz;
    buff[offset + 0] = uInt8Float64Array_ixz[0];
    buff[offset + 1] = uInt8Float64Array_ixz[1];
    buff[offset + 2] = uInt8Float64Array_ixz[2];
    buff[offset + 3] = uInt8Float64Array_ixz[3];
    buff[offset + 4] = uInt8Float64Array_ixz[4];
    buff[offset + 5] = uInt8Float64Array_ixz[5];
    buff[offset + 6] = uInt8Float64Array_ixz[6];
    buff[offset + 7] = uInt8Float64Array_ixz[7];
    offset += 8;
    var float64Array_iyy = new Float64Array(1);
    var uInt8Float64Array_iyy = new Uint8Array(float64Array_iyy.buffer);
    float64Array_iyy[0] = +this.iyy;
    buff[offset + 0] = uInt8Float64Array_iyy[0];
    buff[offset + 1] = uInt8Float64Array_iyy[1];
    buff[offset + 2] = uInt8Float64Array_iyy[2];
    buff[offset + 3] = uInt8Float64Array_iyy[3];
    buff[offset + 4] = uInt8Float64Array_iyy[4];
    buff[offset + 5] = uInt8Float64Array_iyy[5];
    buff[offset + 6] = uInt8Float64Array_iyy[6];
    buff[offset + 7] = uInt8Float64Array_iyy[7];
    offset += 8;
    var float64Array_iyz = new Float64Array(1);
    var uInt8Float64Array_iyz = new Uint8Array(float64Array_iyz.buffer);
    float64Array_iyz[0] = +this.iyz;
    buff[offset + 0] = uInt8Float64Array_iyz[0];
    buff[offset + 1] = uInt8Float64Array_iyz[1];
    buff[offset + 2] = uInt8Float64Array_iyz[2];
    buff[offset + 3] = uInt8Float64Array_iyz[3];
    buff[offset + 4] = uInt8Float64Array_iyz[4];
    buff[offset + 5] = uInt8Float64Array_iyz[5];
    buff[offset + 6] = uInt8Float64Array_iyz[6];
    buff[offset + 7] = uInt8Float64Array_iyz[7];
    offset += 8;
    var float64Array_izz = new Float64Array(1);
    var uInt8Float64Array_izz = new Uint8Array(float64Array_izz.buffer);
    float64Array_izz[0] = +this.izz;
    buff[offset + 0] = uInt8Float64Array_izz[0];
    buff[offset + 1] = uInt8Float64Array_izz[1];
    buff[offset + 2] = uInt8Float64Array_izz[2];
    buff[offset + 3] = uInt8Float64Array_izz[3];
    buff[offset + 4] = uInt8Float64Array_izz[4];
    buff[offset + 5] = uInt8Float64Array_izz[5];
    buff[offset + 6] = uInt8Float64Array_izz[6];
    buff[offset + 7] = uInt8Float64Array_izz[7];
    offset += 8;
    return offset;
};

Inertia.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var float64Array_m = new Float64Array(1);
    var uInt8Float64Array_m = new Uint8Array(float64Array_m.buffer);
    uInt8Float64Array_m[0] = buff[offset + 0];
    uInt8Float64Array_m[1] = buff[offset + 1];
    uInt8Float64Array_m[2] = buff[offset + 2];
    uInt8Float64Array_m[3] = buff[offset + 3];
    uInt8Float64Array_m[4] = buff[offset + 4];
    uInt8Float64Array_m[5] = buff[offset + 5];
    uInt8Float64Array_m[6] = buff[offset + 6];
    uInt8Float64Array_m[7] = buff[offset + 7];
    this.m = float64Array_m[0];
    offset += 8;
    offset += this.com.deserialize(buff, offset);
    var float64Array_ixx = new Float64Array(1);
    var uInt8Float64Array_ixx = new Uint8Array(float64Array_ixx.buffer);
    uInt8Float64Array_ixx[0] = buff[offset + 0];
    uInt8Float64Array_ixx[1] = buff[offset + 1];
    uInt8Float64Array_ixx[2] = buff[offset + 2];
    uInt8Float64Array_ixx[3] = buff[offset + 3];
    uInt8Float64Array_ixx[4] = buff[offset + 4];
    uInt8Float64Array_ixx[5] = buff[offset + 5];
    uInt8Float64Array_ixx[6] = buff[offset + 6];
    uInt8Float64Array_ixx[7] = buff[offset + 7];
    this.ixx = float64Array_ixx[0];
    offset += 8;
    var float64Array_ixy = new Float64Array(1);
    var uInt8Float64Array_ixy = new Uint8Array(float64Array_ixy.buffer);
    uInt8Float64Array_ixy[0] = buff[offset + 0];
    uInt8Float64Array_ixy[1] = buff[offset + 1];
    uInt8Float64Array_ixy[2] = buff[offset + 2];
    uInt8Float64Array_ixy[3] = buff[offset + 3];
    uInt8Float64Array_ixy[4] = buff[offset + 4];
    uInt8Float64Array_ixy[5] = buff[offset + 5];
    uInt8Float64Array_ixy[6] = buff[offset + 6];
    uInt8Float64Array_ixy[7] = buff[offset + 7];
    this.ixy = float64Array_ixy[0];
    offset += 8;
    var float64Array_ixz = new Float64Array(1);
    var uInt8Float64Array_ixz = new Uint8Array(float64Array_ixz.buffer);
    uInt8Float64Array_ixz[0] = buff[offset + 0];
    uInt8Float64Array_ixz[1] = buff[offset + 1];
    uInt8Float64Array_ixz[2] = buff[offset + 2];
    uInt8Float64Array_ixz[3] = buff[offset + 3];
    uInt8Float64Array_ixz[4] = buff[offset + 4];
    uInt8Float64Array_ixz[5] = buff[offset + 5];
    uInt8Float64Array_ixz[6] = buff[offset + 6];
    uInt8Float64Array_ixz[7] = buff[offset + 7];
    this.ixz = float64Array_ixz[0];
    offset += 8;
    var float64Array_iyy = new Float64Array(1);
    var uInt8Float64Array_iyy = new Uint8Array(float64Array_iyy.buffer);
    uInt8Float64Array_iyy[0] = buff[offset + 0];
    uInt8Float64Array_iyy[1] = buff[offset + 1];
    uInt8Float64Array_iyy[2] = buff[offset + 2];
    uInt8Float64Array_iyy[3] = buff[offset + 3];
    uInt8Float64Array_iyy[4] = buff[offset + 4];
    uInt8Float64Array_iyy[5] = buff[offset + 5];
    uInt8Float64Array_iyy[6] = buff[offset + 6];
    uInt8Float64Array_iyy[7] = buff[offset + 7];
    this.iyy = float64Array_iyy[0];
    offset += 8;
    var float64Array_iyz = new Float64Array(1);
    var uInt8Float64Array_iyz = new Uint8Array(float64Array_iyz.buffer);
    uInt8Float64Array_iyz[0] = buff[offset + 0];
    uInt8Float64Array_iyz[1] = buff[offset + 1];
    uInt8Float64Array_iyz[2] = buff[offset + 2];
    uInt8Float64Array_iyz[3] = buff[offset + 3];
    uInt8Float64Array_iyz[4] = buff[offset + 4];
    uInt8Float64Array_iyz[5] = buff[offset + 5];
    uInt8Float64Array_iyz[6] = buff[offset + 6];
    uInt8Float64Array_iyz[7] = buff[offset + 7];
    this.iyz = float64Array_iyz[0];
    offset += 8;
    var float64Array_izz = new Float64Array(1);
    var uInt8Float64Array_izz = new Uint8Array(float64Array_izz.buffer);
    uInt8Float64Array_izz[0] = buff[offset + 0];
    uInt8Float64Array_izz[1] = buff[offset + 1];
    uInt8Float64Array_izz[2] = buff[offset + 2];
    uInt8Float64Array_izz[3] = buff[offset + 3];
    uInt8Float64Array_izz[4] = buff[offset + 4];
    uInt8Float64Array_izz[5] = buff[offset + 5];
    uInt8Float64Array_izz[6] = buff[offset + 6];
    uInt8Float64Array_izz[7] = buff[offset + 7];
    this.izz = float64Array_izz[0];
    offset += 8;
    return offset;
};

Inertia.prototype.serializedLength = function() {
    var length = 0;
    length += 8
    length += this.com.serializedLength();
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    return length;
};

Inertia.prototype.echo = function() { return ""; };

Inertia.prototype.getType = function() { return "geometry_msgs/Inertia"; };

Inertia.prototype.getMD5 = function() { return "9116c935782bc29999dad1927624dff0"; };

Inertia.prototype.getID = function() { return 0; };

Inertia.prototype.setID = function(id) { };

return function() { return new Inertia(); };

});
