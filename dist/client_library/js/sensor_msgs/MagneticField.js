(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.MagneticField=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Vector3.js'></script>");

function MagneticField() {
    this.header = std_msgs.Header();
    this.magnetic_field = geometry_msgs.Vector3();
    this.magnetic_field_covariance = [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0];
};

MagneticField.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.magnetic_field.serialize(buff, offset);
    for (var i = 0; i < 9; i++) {
        var float64Array_magnetic_field_covariancei = new Float64Array(1);
        var uInt8Float64Array_magnetic_field_covariancei = new Uint8Array(float64Array_magnetic_field_covariancei.buffer);
        float64Array_magnetic_field_covariancei[0] = +this.magnetic_field_covariance[i];
        buff[offset + 0] = uInt8Float64Array_magnetic_field_covariancei[0];
        buff[offset + 1] = uInt8Float64Array_magnetic_field_covariancei[1];
        buff[offset + 2] = uInt8Float64Array_magnetic_field_covariancei[2];
        buff[offset + 3] = uInt8Float64Array_magnetic_field_covariancei[3];
        buff[offset + 4] = uInt8Float64Array_magnetic_field_covariancei[4];
        buff[offset + 5] = uInt8Float64Array_magnetic_field_covariancei[5];
        buff[offset + 6] = uInt8Float64Array_magnetic_field_covariancei[6];
        buff[offset + 7] = uInt8Float64Array_magnetic_field_covariancei[7];
        offset += 8;
    }
    return offset;
};

MagneticField.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.magnetic_field.deserialize(buff, offset);
    for (var i = 0; i < 9; i++) {
        var float64Array_magnetic_field_covariancei = new Float64Array(1);
        var uInt8Float64Array_magnetic_field_covariancei = new Uint8Array(float64Array_magnetic_field_covariancei.buffer);
        uInt8Float64Array_magnetic_field_covariancei[0] = buff[offset + 0];
        uInt8Float64Array_magnetic_field_covariancei[1] = buff[offset + 1];
        uInt8Float64Array_magnetic_field_covariancei[2] = buff[offset + 2];
        uInt8Float64Array_magnetic_field_covariancei[3] = buff[offset + 3];
        uInt8Float64Array_magnetic_field_covariancei[4] = buff[offset + 4];
        uInt8Float64Array_magnetic_field_covariancei[5] = buff[offset + 5];
        uInt8Float64Array_magnetic_field_covariancei[6] = buff[offset + 6];
        uInt8Float64Array_magnetic_field_covariancei[7] = buff[offset + 7];
        this.magnetic_field_covariance[i] = float64Array_magnetic_field_covariancei[0];
        offset += 8;
    }
    return offset;
};

MagneticField.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.magnetic_field.serializedLength();
    for (var i = 0; i < 9; i++) {
        length += 8
    }
    return length;
};

MagneticField.prototype.echo = function() { return ""; };

MagneticField.prototype.getType = function() { return "sensor_msgs/MagneticField"; };

MagneticField.prototype.getMD5 = function() { return "f8e051d776de1349146122759c66db92"; };

MagneticField.prototype.getID = function() { return 0; };

MagneticField.prototype.setID = function(id) { };

return function() { return new MagneticField(); };

});
