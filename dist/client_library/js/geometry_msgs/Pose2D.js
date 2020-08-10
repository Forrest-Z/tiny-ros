(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.Pose2D=f();}})(function(){var define,module,exports;'use strict';

function Pose2D() {
    this.x = 0.0;
    this.y = 0.0;
    this.theta = 0.0;
};

Pose2D.prototype.serialize = function(buff, idx) {
    var offset = idx;
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
    var float64Array_theta = new Float64Array(1);
    var uInt8Float64Array_theta = new Uint8Array(float64Array_theta.buffer);
    float64Array_theta[0] = +this.theta;
    buff[offset + 0] = uInt8Float64Array_theta[0];
    buff[offset + 1] = uInt8Float64Array_theta[1];
    buff[offset + 2] = uInt8Float64Array_theta[2];
    buff[offset + 3] = uInt8Float64Array_theta[3];
    buff[offset + 4] = uInt8Float64Array_theta[4];
    buff[offset + 5] = uInt8Float64Array_theta[5];
    buff[offset + 6] = uInt8Float64Array_theta[6];
    buff[offset + 7] = uInt8Float64Array_theta[7];
    offset += 8;
    return offset;
};

Pose2D.prototype.deserialize = function(buff, idx) {
    var offset = idx;
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
    var float64Array_theta = new Float64Array(1);
    var uInt8Float64Array_theta = new Uint8Array(float64Array_theta.buffer);
    uInt8Float64Array_theta[0] = buff[offset + 0];
    uInt8Float64Array_theta[1] = buff[offset + 1];
    uInt8Float64Array_theta[2] = buff[offset + 2];
    uInt8Float64Array_theta[3] = buff[offset + 3];
    uInt8Float64Array_theta[4] = buff[offset + 4];
    uInt8Float64Array_theta[5] = buff[offset + 5];
    uInt8Float64Array_theta[6] = buff[offset + 6];
    uInt8Float64Array_theta[7] = buff[offset + 7];
    this.theta = float64Array_theta[0];
    offset += 8;
    return offset;
};

Pose2D.prototype.serializedLength = function() {
    var length = 0;
    length += 8
    length += 8
    length += 8
    return length;
};

Pose2D.prototype.echo = function() { return ""; };

Pose2D.prototype.getType = function() { return "geometry_msgs/Pose2D"; };

Pose2D.prototype.getMD5 = function() { return "509f362ff66c4d3df21020fa7c01f8c6"; };

Pose2D.prototype.getID = function() { return 0; };

Pose2D.prototype.setID = function(id) { };

return function() { return new Pose2D(); };

});
