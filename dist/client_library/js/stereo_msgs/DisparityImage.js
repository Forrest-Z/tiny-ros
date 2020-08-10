(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.stereo_msgs==="undefined"){g.stereo_msgs={};}g.stereo_msgs.DisparityImage=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/sensor_msgs/Image.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/sensor_msgs/RegionOfInterest.js'></script>");

function DisparityImage() {
    this.header = std_msgs.Header();
    this.image = sensor_msgs.Image();
    this.f = 0.0;
    this.T = 0.0;
    this.valid_window = sensor_msgs.RegionOfInterest();
    this.min_disparity = 0.0;
    this.max_disparity = 0.0;
    this.delta_d = 0.0;
};

DisparityImage.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.image.serialize(buff, offset);
    var float32Array_f = new Float32Array(1);
    var uInt8Float32Array_f = new Uint8Array(float32Array_f.buffer);
    float32Array_f[0] = +this.f;
    buff[offset + 0] = uInt8Float32Array_f[0];
    buff[offset + 1] = uInt8Float32Array_f[1];
    buff[offset + 2] = uInt8Float32Array_f[2];
    buff[offset + 3] = uInt8Float32Array_f[3];
    offset += 4;
    var float32Array_T = new Float32Array(1);
    var uInt8Float32Array_T = new Uint8Array(float32Array_T.buffer);
    float32Array_T[0] = +this.T;
    buff[offset + 0] = uInt8Float32Array_T[0];
    buff[offset + 1] = uInt8Float32Array_T[1];
    buff[offset + 2] = uInt8Float32Array_T[2];
    buff[offset + 3] = uInt8Float32Array_T[3];
    offset += 4;
    offset += this.valid_window.serialize(buff, offset);
    var float32Array_min_disparity = new Float32Array(1);
    var uInt8Float32Array_min_disparity = new Uint8Array(float32Array_min_disparity.buffer);
    float32Array_min_disparity[0] = +this.min_disparity;
    buff[offset + 0] = uInt8Float32Array_min_disparity[0];
    buff[offset + 1] = uInt8Float32Array_min_disparity[1];
    buff[offset + 2] = uInt8Float32Array_min_disparity[2];
    buff[offset + 3] = uInt8Float32Array_min_disparity[3];
    offset += 4;
    var float32Array_max_disparity = new Float32Array(1);
    var uInt8Float32Array_max_disparity = new Uint8Array(float32Array_max_disparity.buffer);
    float32Array_max_disparity[0] = +this.max_disparity;
    buff[offset + 0] = uInt8Float32Array_max_disparity[0];
    buff[offset + 1] = uInt8Float32Array_max_disparity[1];
    buff[offset + 2] = uInt8Float32Array_max_disparity[2];
    buff[offset + 3] = uInt8Float32Array_max_disparity[3];
    offset += 4;
    var float32Array_delta_d = new Float32Array(1);
    var uInt8Float32Array_delta_d = new Uint8Array(float32Array_delta_d.buffer);
    float32Array_delta_d[0] = +this.delta_d;
    buff[offset + 0] = uInt8Float32Array_delta_d[0];
    buff[offset + 1] = uInt8Float32Array_delta_d[1];
    buff[offset + 2] = uInt8Float32Array_delta_d[2];
    buff[offset + 3] = uInt8Float32Array_delta_d[3];
    offset += 4;
    return offset;
};

DisparityImage.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.image.deserialize(buff, offset);
    var float32Array_f = new Float32Array(1);
    var uInt8Float32Array_f = new Uint8Array(float32Array_f.buffer);
    uInt8Float32Array_f[0] = buff[offset + 0];
    uInt8Float32Array_f[1] = buff[offset + 1];
    uInt8Float32Array_f[2] = buff[offset + 2];
    uInt8Float32Array_f[3] = buff[offset + 3];
    this.f = float32Array_f[0];
    offset += 4;
    var float32Array_T = new Float32Array(1);
    var uInt8Float32Array_T = new Uint8Array(float32Array_T.buffer);
    uInt8Float32Array_T[0] = buff[offset + 0];
    uInt8Float32Array_T[1] = buff[offset + 1];
    uInt8Float32Array_T[2] = buff[offset + 2];
    uInt8Float32Array_T[3] = buff[offset + 3];
    this.T = float32Array_T[0];
    offset += 4;
    offset += this.valid_window.deserialize(buff, offset);
    var float32Array_min_disparity = new Float32Array(1);
    var uInt8Float32Array_min_disparity = new Uint8Array(float32Array_min_disparity.buffer);
    uInt8Float32Array_min_disparity[0] = buff[offset + 0];
    uInt8Float32Array_min_disparity[1] = buff[offset + 1];
    uInt8Float32Array_min_disparity[2] = buff[offset + 2];
    uInt8Float32Array_min_disparity[3] = buff[offset + 3];
    this.min_disparity = float32Array_min_disparity[0];
    offset += 4;
    var float32Array_max_disparity = new Float32Array(1);
    var uInt8Float32Array_max_disparity = new Uint8Array(float32Array_max_disparity.buffer);
    uInt8Float32Array_max_disparity[0] = buff[offset + 0];
    uInt8Float32Array_max_disparity[1] = buff[offset + 1];
    uInt8Float32Array_max_disparity[2] = buff[offset + 2];
    uInt8Float32Array_max_disparity[3] = buff[offset + 3];
    this.max_disparity = float32Array_max_disparity[0];
    offset += 4;
    var float32Array_delta_d = new Float32Array(1);
    var uInt8Float32Array_delta_d = new Uint8Array(float32Array_delta_d.buffer);
    uInt8Float32Array_delta_d[0] = buff[offset + 0];
    uInt8Float32Array_delta_d[1] = buff[offset + 1];
    uInt8Float32Array_delta_d[2] = buff[offset + 2];
    uInt8Float32Array_delta_d[3] = buff[offset + 3];
    this.delta_d = float32Array_delta_d[0];
    offset += 4;
    return offset;
};

DisparityImage.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.image.serializedLength();
    length += 4
    length += 4
    length += this.valid_window.serializedLength();
    length += 4
    length += 4
    length += 4
    return length;
};

DisparityImage.prototype.echo = function() { return ""; };

DisparityImage.prototype.getType = function() { return "stereo_msgs/DisparityImage"; };

DisparityImage.prototype.getMD5 = function() { return "03545cef8df8d20bea21fdbbf9482b4b"; };

DisparityImage.prototype.getID = function() { return 0; };

DisparityImage.prototype.setID = function(id) { };

return function() { return new DisparityImage(); };

});
