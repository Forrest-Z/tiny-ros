(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.PointCloud2=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/sensor_msgs/PointField.js'></script>");

function PointCloud2() {
    this.header = std_msgs.Header();
    this.height = 0;
    this.width = 0;
    this.fields = new Array();
    this.is_bigendian = false;
    this.point_step = 0;
    this.row_step = 0;
    this.data = new Array();
    this.is_dense = false;
};

PointCloud2.prototype.serialize = function(buff, idx) {
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
    var length_fields = this.fields.length;
    buff[offset + 0] = (length_fields >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_fields >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_fields >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_fields >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_fields; i++) {
        offset += this.fields[i].serialize(buff, offset);
    }
    buff[offset] = this.is_bigendian === false ? 0 : 1;
    offset += 1;
    buff[offset + 0] = ((+this.point_step) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.point_step) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.point_step) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.point_step) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.row_step) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.row_step) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.row_step) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.row_step) >> (8 * 3)) & 0xFF;
    offset += 4;
    var length_data = this.data.length;
    buff[offset + 0] = (length_data >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_data >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_data >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_data >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_data; i++) {
        buff[offset + 0] = ((+this.data[i]) >> (8 * 0)) & 0xFF;
        offset += 1;
    }
    buff[offset] = this.is_dense === false ? 0 : 1;
    offset += 1;
    return offset;
};

PointCloud2.prototype.deserialize = function(buff, idx) {
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
    var length_fields = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_fields |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_fields |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_fields |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.fields = new Array(length_fields);
    for (var i = 0; i < length_fields; i++) {
        this.fields[i] = sensor_msgs.PointField();
        offset += this.fields[i].deserialize(buff, offset);
    }
    this.is_bigendian = buff[offset] !== 0 ? true : false;
    offset += 1;
    this.point_step = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.point_step |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.point_step |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.point_step |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.row_step = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.row_step |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.row_step |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.row_step |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_data |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_data |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_data |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.data = new Array(length_data);
    for (var i = 0; i < length_data; i++) {
        this.data[i] = +((buff[offset + 0] & 0xFF) << (8 * 0));
        offset += 1;
    }
    this.is_dense = buff[offset] !== 0 ? true : false;
    offset += 1;
    return offset;
};

PointCloud2.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 4
    length += 4
    var length_fields = this.fields.length;
    length += 4;
    for (var i = 0; i < length_fields; i++) {
        length += this.fields[i].serializedLength();
    }
    length += 1
    length += 4
    length += 4
    var length_data = this.data.length;
    length += 4;
    for (var i = 0; i < length_data; i++) {
        length += 1
    }
    length += 1
    return length;
};

PointCloud2.prototype.echo = function() { return ""; };

PointCloud2.prototype.getType = function() { return "sensor_msgs/PointCloud2"; };

PointCloud2.prototype.getMD5 = function() { return "6aa926339b282463281af40546b3b3cf"; };

PointCloud2.prototype.getID = function() { return 0; };

PointCloud2.prototype.setID = function(id) { };

return function() { return new PointCloud2(); };

});
