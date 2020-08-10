(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.PointCloud=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Point32.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/sensor_msgs/ChannelFloat32.js'></script>");

function PointCloud() {
    this.header = std_msgs.Header();
    this.points = new Array();
    this.channels = new Array();
};

PointCloud.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var length_points = this.points.length;
    buff[offset + 0] = (length_points >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_points >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_points >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_points >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_points; i++) {
        offset += this.points[i].serialize(buff, offset);
    }
    var length_channels = this.channels.length;
    buff[offset + 0] = (length_channels >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_channels >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_channels >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_channels >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_channels; i++) {
        offset += this.channels[i].serialize(buff, offset);
    }
    return offset;
};

PointCloud.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_points = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_points |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_points |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_points |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.points = new Array(length_points);
    for (var i = 0; i < length_points; i++) {
        this.points[i] = geometry_msgs.Point32();
        offset += this.points[i].deserialize(buff, offset);
    }
    var length_channels = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_channels |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_channels |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_channels |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.channels = new Array(length_channels);
    for (var i = 0; i < length_channels; i++) {
        this.channels[i] = sensor_msgs.ChannelFloat32();
        offset += this.channels[i].deserialize(buff, offset);
    }
    return offset;
};

PointCloud.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var length_points = this.points.length;
    length += 4;
    for (var i = 0; i < length_points; i++) {
        length += this.points[i].serializedLength();
    }
    var length_channels = this.channels.length;
    length += 4;
    for (var i = 0; i < length_channels; i++) {
        length += this.channels[i].serializedLength();
    }
    return length;
};

PointCloud.prototype.echo = function() { return ""; };

PointCloud.prototype.getType = function() { return "sensor_msgs/PointCloud"; };

PointCloud.prototype.getMD5 = function() { return "b01249148cae0106a561ab36cd1e48a8"; };

PointCloud.prototype.getID = function() { return 0; };

PointCloud.prototype.setID = function(id) { };

return function() { return new PointCloud(); };

});
