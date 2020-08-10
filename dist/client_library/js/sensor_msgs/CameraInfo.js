(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.CameraInfo=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/sensor_msgs/RegionOfInterest.js'></script>");

function CameraInfo() {
    this.header = std_msgs.Header();
    this.height = 0;
    this.width = 0;
    this.distortion_model = "";
    this.D = new Array();
    this.K = [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0];
    this.R = [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0];
    this.P = [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0];
    this.binning_x = 0;
    this.binning_y = 0;
    this.roi = sensor_msgs.RegionOfInterest();
};

CameraInfo.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
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
    var encoder_distortion_model = new TextEncoder('utf8');
    var utf8array_distortion_model = encoder_distortion_model.encode(this.distortion_model);
    buff[offset + 0] = (utf8array_distortion_model.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_distortion_model.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_distortion_model.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_distortion_model.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_distortion_model.length; i++) {
        buff[offset + i] = utf8array_distortion_model[i];
    }
    offset += utf8array_distortion_model.length;
    var length_D = this.D.length;
    buff[offset + 0] = (length_D >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_D >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_D >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_D >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_D; i++) {
        var float64Array_Di = new Float64Array(1);
        var uInt8Float64Array_Di = new Uint8Array(float64Array_Di.buffer);
        float64Array_Di[0] = +this.D[i];
        buff[offset + 0] = uInt8Float64Array_Di[0];
        buff[offset + 1] = uInt8Float64Array_Di[1];
        buff[offset + 2] = uInt8Float64Array_Di[2];
        buff[offset + 3] = uInt8Float64Array_Di[3];
        buff[offset + 4] = uInt8Float64Array_Di[4];
        buff[offset + 5] = uInt8Float64Array_Di[5];
        buff[offset + 6] = uInt8Float64Array_Di[6];
        buff[offset + 7] = uInt8Float64Array_Di[7];
        offset += 8;
    }
    for (var i = 0; i < 9; i++) {
        var float64Array_Ki = new Float64Array(1);
        var uInt8Float64Array_Ki = new Uint8Array(float64Array_Ki.buffer);
        float64Array_Ki[0] = +this.K[i];
        buff[offset + 0] = uInt8Float64Array_Ki[0];
        buff[offset + 1] = uInt8Float64Array_Ki[1];
        buff[offset + 2] = uInt8Float64Array_Ki[2];
        buff[offset + 3] = uInt8Float64Array_Ki[3];
        buff[offset + 4] = uInt8Float64Array_Ki[4];
        buff[offset + 5] = uInt8Float64Array_Ki[5];
        buff[offset + 6] = uInt8Float64Array_Ki[6];
        buff[offset + 7] = uInt8Float64Array_Ki[7];
        offset += 8;
    }
    for (var i = 0; i < 9; i++) {
        var float64Array_Ri = new Float64Array(1);
        var uInt8Float64Array_Ri = new Uint8Array(float64Array_Ri.buffer);
        float64Array_Ri[0] = +this.R[i];
        buff[offset + 0] = uInt8Float64Array_Ri[0];
        buff[offset + 1] = uInt8Float64Array_Ri[1];
        buff[offset + 2] = uInt8Float64Array_Ri[2];
        buff[offset + 3] = uInt8Float64Array_Ri[3];
        buff[offset + 4] = uInt8Float64Array_Ri[4];
        buff[offset + 5] = uInt8Float64Array_Ri[5];
        buff[offset + 6] = uInt8Float64Array_Ri[6];
        buff[offset + 7] = uInt8Float64Array_Ri[7];
        offset += 8;
    }
    for (var i = 0; i < 12; i++) {
        var float64Array_Pi = new Float64Array(1);
        var uInt8Float64Array_Pi = new Uint8Array(float64Array_Pi.buffer);
        float64Array_Pi[0] = +this.P[i];
        buff[offset + 0] = uInt8Float64Array_Pi[0];
        buff[offset + 1] = uInt8Float64Array_Pi[1];
        buff[offset + 2] = uInt8Float64Array_Pi[2];
        buff[offset + 3] = uInt8Float64Array_Pi[3];
        buff[offset + 4] = uInt8Float64Array_Pi[4];
        buff[offset + 5] = uInt8Float64Array_Pi[5];
        buff[offset + 6] = uInt8Float64Array_Pi[6];
        buff[offset + 7] = uInt8Float64Array_Pi[7];
        offset += 8;
    }
    buff[offset + 0] = ((+this.binning_x) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.binning_x) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.binning_x) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.binning_x) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.binning_y) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.binning_y) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.binning_y) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.binning_y) >> (8 * 3)) & 0xFF;
    offset += 4;
    offset += this.roi.serialize(buff, offset);
    return offset;
};

CameraInfo.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
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
    var length_distortion_model = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_distortion_model |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_distortion_model |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_distortion_model |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_distortion_model = new TextDecoder('utf8');
    this.distortion_model = decoder_distortion_model.decode(buff.slice(offset, offset + length_distortion_model));
    offset += length_distortion_model;
    var length_D = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_D |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_D |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_D |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.D = new Array(length_D);
    for (var i = 0; i < length_D; i++) {
        var float64Array_Di = new Float64Array(1);
        var uInt8Float64Array_Di = new Uint8Array(float64Array_Di.buffer);
        uInt8Float64Array_Di[0] = buff[offset + 0];
        uInt8Float64Array_Di[1] = buff[offset + 1];
        uInt8Float64Array_Di[2] = buff[offset + 2];
        uInt8Float64Array_Di[3] = buff[offset + 3];
        uInt8Float64Array_Di[4] = buff[offset + 4];
        uInt8Float64Array_Di[5] = buff[offset + 5];
        uInt8Float64Array_Di[6] = buff[offset + 6];
        uInt8Float64Array_Di[7] = buff[offset + 7];
        this.D[i] = float64Array_Di[0];
        offset += 8;
    }
    for (var i = 0; i < 9; i++) {
        var float64Array_Ki = new Float64Array(1);
        var uInt8Float64Array_Ki = new Uint8Array(float64Array_Ki.buffer);
        uInt8Float64Array_Ki[0] = buff[offset + 0];
        uInt8Float64Array_Ki[1] = buff[offset + 1];
        uInt8Float64Array_Ki[2] = buff[offset + 2];
        uInt8Float64Array_Ki[3] = buff[offset + 3];
        uInt8Float64Array_Ki[4] = buff[offset + 4];
        uInt8Float64Array_Ki[5] = buff[offset + 5];
        uInt8Float64Array_Ki[6] = buff[offset + 6];
        uInt8Float64Array_Ki[7] = buff[offset + 7];
        this.K[i] = float64Array_Ki[0];
        offset += 8;
    }
    for (var i = 0; i < 9; i++) {
        var float64Array_Ri = new Float64Array(1);
        var uInt8Float64Array_Ri = new Uint8Array(float64Array_Ri.buffer);
        uInt8Float64Array_Ri[0] = buff[offset + 0];
        uInt8Float64Array_Ri[1] = buff[offset + 1];
        uInt8Float64Array_Ri[2] = buff[offset + 2];
        uInt8Float64Array_Ri[3] = buff[offset + 3];
        uInt8Float64Array_Ri[4] = buff[offset + 4];
        uInt8Float64Array_Ri[5] = buff[offset + 5];
        uInt8Float64Array_Ri[6] = buff[offset + 6];
        uInt8Float64Array_Ri[7] = buff[offset + 7];
        this.R[i] = float64Array_Ri[0];
        offset += 8;
    }
    for (var i = 0; i < 12; i++) {
        var float64Array_Pi = new Float64Array(1);
        var uInt8Float64Array_Pi = new Uint8Array(float64Array_Pi.buffer);
        uInt8Float64Array_Pi[0] = buff[offset + 0];
        uInt8Float64Array_Pi[1] = buff[offset + 1];
        uInt8Float64Array_Pi[2] = buff[offset + 2];
        uInt8Float64Array_Pi[3] = buff[offset + 3];
        uInt8Float64Array_Pi[4] = buff[offset + 4];
        uInt8Float64Array_Pi[5] = buff[offset + 5];
        uInt8Float64Array_Pi[6] = buff[offset + 6];
        uInt8Float64Array_Pi[7] = buff[offset + 7];
        this.P[i] = float64Array_Pi[0];
        offset += 8;
    }
    this.binning_x = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.binning_x |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.binning_x |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.binning_x |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.binning_y = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.binning_y |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.binning_y |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.binning_y |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    offset += this.roi.deserialize(buff, offset);
    return offset;
};

CameraInfo.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 4
    length += 4
    var encoder_distortion_model = new TextEncoder('utf8');
    var utf8array_distortion_model = encoder_distortion_model.encode(this.distortion_model);
    length += 4;
    length += utf8array_distortion_model.length;
    var length_D = this.D.length;
    length += 4;
    for (var i = 0; i < length_D; i++) {
        length += 8
    }
    for (var i = 0; i < 9; i++) {
        length += 8
    }
    for (var i = 0; i < 9; i++) {
        length += 8
    }
    for (var i = 0; i < 12; i++) {
        length += 8
    }
    length += 4
    length += 4
    length += this.roi.serializedLength();
    return length;
};

CameraInfo.prototype.echo = function() { return ""; };

CameraInfo.prototype.getType = function() { return "sensor_msgs/CameraInfo"; };

CameraInfo.prototype.getMD5 = function() { return "57d2553deec0a7842f00837f40032798"; };

CameraInfo.prototype.getID = function() { return 0; };

CameraInfo.prototype.setID = function(id) { };

return function() { return new CameraInfo(); };

});
