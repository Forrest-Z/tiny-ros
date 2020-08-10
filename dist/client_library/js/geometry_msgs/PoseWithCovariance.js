(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.PoseWithCovariance=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Pose.js'></script>");

function PoseWithCovariance() {
    this.pose = geometry_msgs.Pose();
    this.covariance = [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0];
};

PoseWithCovariance.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.pose.serialize(buff, offset);
    for (var i = 0; i < 36; i++) {
        var float64Array_covariancei = new Float64Array(1);
        var uInt8Float64Array_covariancei = new Uint8Array(float64Array_covariancei.buffer);
        float64Array_covariancei[0] = +this.covariance[i];
        buff[offset + 0] = uInt8Float64Array_covariancei[0];
        buff[offset + 1] = uInt8Float64Array_covariancei[1];
        buff[offset + 2] = uInt8Float64Array_covariancei[2];
        buff[offset + 3] = uInt8Float64Array_covariancei[3];
        buff[offset + 4] = uInt8Float64Array_covariancei[4];
        buff[offset + 5] = uInt8Float64Array_covariancei[5];
        buff[offset + 6] = uInt8Float64Array_covariancei[6];
        buff[offset + 7] = uInt8Float64Array_covariancei[7];
        offset += 8;
    }
    return offset;
};

PoseWithCovariance.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.pose.deserialize(buff, offset);
    for (var i = 0; i < 36; i++) {
        var float64Array_covariancei = new Float64Array(1);
        var uInt8Float64Array_covariancei = new Uint8Array(float64Array_covariancei.buffer);
        uInt8Float64Array_covariancei[0] = buff[offset + 0];
        uInt8Float64Array_covariancei[1] = buff[offset + 1];
        uInt8Float64Array_covariancei[2] = buff[offset + 2];
        uInt8Float64Array_covariancei[3] = buff[offset + 3];
        uInt8Float64Array_covariancei[4] = buff[offset + 4];
        uInt8Float64Array_covariancei[5] = buff[offset + 5];
        uInt8Float64Array_covariancei[6] = buff[offset + 6];
        uInt8Float64Array_covariancei[7] = buff[offset + 7];
        this.covariance[i] = float64Array_covariancei[0];
        offset += 8;
    }
    return offset;
};

PoseWithCovariance.prototype.serializedLength = function() {
    var length = 0;
    length += this.pose.serializedLength();
    for (var i = 0; i < 36; i++) {
        length += 8
    }
    return length;
};

PoseWithCovariance.prototype.echo = function() { return ""; };

PoseWithCovariance.prototype.getType = function() { return "geometry_msgs/PoseWithCovariance"; };

PoseWithCovariance.prototype.getMD5 = function() { return "054c6283d50e78f8d9358aaaee5f4c1b"; };

PoseWithCovariance.prototype.getID = function() { return 0; };

PoseWithCovariance.prototype.setID = function(id) { };

return function() { return new PoseWithCovariance(); };

});
