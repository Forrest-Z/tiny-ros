(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.ProjectedMapInfo=f();}})(function(){var define,module,exports;'use strict';

function ProjectedMapInfo() {
    this.frame_id = "";
    this.x = 0.0;
    this.y = 0.0;
    this.width = 0.0;
    this.height = 0.0;
    this.min_z = 0.0;
    this.max_z = 0.0;
};

ProjectedMapInfo.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var encoder_frame_id = new TextEncoder('utf8');
    var utf8array_frame_id = encoder_frame_id.encode(this.frame_id);
    buff[offset + 0] = (utf8array_frame_id.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_frame_id.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_frame_id.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_frame_id.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_frame_id.length; i++) {
        buff[offset + i] = utf8array_frame_id[i];
    }
    offset += utf8array_frame_id.length;
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
    var float64Array_width = new Float64Array(1);
    var uInt8Float64Array_width = new Uint8Array(float64Array_width.buffer);
    float64Array_width[0] = +this.width;
    buff[offset + 0] = uInt8Float64Array_width[0];
    buff[offset + 1] = uInt8Float64Array_width[1];
    buff[offset + 2] = uInt8Float64Array_width[2];
    buff[offset + 3] = uInt8Float64Array_width[3];
    buff[offset + 4] = uInt8Float64Array_width[4];
    buff[offset + 5] = uInt8Float64Array_width[5];
    buff[offset + 6] = uInt8Float64Array_width[6];
    buff[offset + 7] = uInt8Float64Array_width[7];
    offset += 8;
    var float64Array_height = new Float64Array(1);
    var uInt8Float64Array_height = new Uint8Array(float64Array_height.buffer);
    float64Array_height[0] = +this.height;
    buff[offset + 0] = uInt8Float64Array_height[0];
    buff[offset + 1] = uInt8Float64Array_height[1];
    buff[offset + 2] = uInt8Float64Array_height[2];
    buff[offset + 3] = uInt8Float64Array_height[3];
    buff[offset + 4] = uInt8Float64Array_height[4];
    buff[offset + 5] = uInt8Float64Array_height[5];
    buff[offset + 6] = uInt8Float64Array_height[6];
    buff[offset + 7] = uInt8Float64Array_height[7];
    offset += 8;
    var float64Array_min_z = new Float64Array(1);
    var uInt8Float64Array_min_z = new Uint8Array(float64Array_min_z.buffer);
    float64Array_min_z[0] = +this.min_z;
    buff[offset + 0] = uInt8Float64Array_min_z[0];
    buff[offset + 1] = uInt8Float64Array_min_z[1];
    buff[offset + 2] = uInt8Float64Array_min_z[2];
    buff[offset + 3] = uInt8Float64Array_min_z[3];
    buff[offset + 4] = uInt8Float64Array_min_z[4];
    buff[offset + 5] = uInt8Float64Array_min_z[5];
    buff[offset + 6] = uInt8Float64Array_min_z[6];
    buff[offset + 7] = uInt8Float64Array_min_z[7];
    offset += 8;
    var float64Array_max_z = new Float64Array(1);
    var uInt8Float64Array_max_z = new Uint8Array(float64Array_max_z.buffer);
    float64Array_max_z[0] = +this.max_z;
    buff[offset + 0] = uInt8Float64Array_max_z[0];
    buff[offset + 1] = uInt8Float64Array_max_z[1];
    buff[offset + 2] = uInt8Float64Array_max_z[2];
    buff[offset + 3] = uInt8Float64Array_max_z[3];
    buff[offset + 4] = uInt8Float64Array_max_z[4];
    buff[offset + 5] = uInt8Float64Array_max_z[5];
    buff[offset + 6] = uInt8Float64Array_max_z[6];
    buff[offset + 7] = uInt8Float64Array_max_z[7];
    offset += 8;
    return offset;
};

ProjectedMapInfo.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_frame_id = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_frame_id |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_frame_id |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_frame_id |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_frame_id = new TextDecoder('utf8');
    this.frame_id = decoder_frame_id.decode(buff.slice(offset, offset + length_frame_id));
    offset += length_frame_id;
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
    var float64Array_width = new Float64Array(1);
    var uInt8Float64Array_width = new Uint8Array(float64Array_width.buffer);
    uInt8Float64Array_width[0] = buff[offset + 0];
    uInt8Float64Array_width[1] = buff[offset + 1];
    uInt8Float64Array_width[2] = buff[offset + 2];
    uInt8Float64Array_width[3] = buff[offset + 3];
    uInt8Float64Array_width[4] = buff[offset + 4];
    uInt8Float64Array_width[5] = buff[offset + 5];
    uInt8Float64Array_width[6] = buff[offset + 6];
    uInt8Float64Array_width[7] = buff[offset + 7];
    this.width = float64Array_width[0];
    offset += 8;
    var float64Array_height = new Float64Array(1);
    var uInt8Float64Array_height = new Uint8Array(float64Array_height.buffer);
    uInt8Float64Array_height[0] = buff[offset + 0];
    uInt8Float64Array_height[1] = buff[offset + 1];
    uInt8Float64Array_height[2] = buff[offset + 2];
    uInt8Float64Array_height[3] = buff[offset + 3];
    uInt8Float64Array_height[4] = buff[offset + 4];
    uInt8Float64Array_height[5] = buff[offset + 5];
    uInt8Float64Array_height[6] = buff[offset + 6];
    uInt8Float64Array_height[7] = buff[offset + 7];
    this.height = float64Array_height[0];
    offset += 8;
    var float64Array_min_z = new Float64Array(1);
    var uInt8Float64Array_min_z = new Uint8Array(float64Array_min_z.buffer);
    uInt8Float64Array_min_z[0] = buff[offset + 0];
    uInt8Float64Array_min_z[1] = buff[offset + 1];
    uInt8Float64Array_min_z[2] = buff[offset + 2];
    uInt8Float64Array_min_z[3] = buff[offset + 3];
    uInt8Float64Array_min_z[4] = buff[offset + 4];
    uInt8Float64Array_min_z[5] = buff[offset + 5];
    uInt8Float64Array_min_z[6] = buff[offset + 6];
    uInt8Float64Array_min_z[7] = buff[offset + 7];
    this.min_z = float64Array_min_z[0];
    offset += 8;
    var float64Array_max_z = new Float64Array(1);
    var uInt8Float64Array_max_z = new Uint8Array(float64Array_max_z.buffer);
    uInt8Float64Array_max_z[0] = buff[offset + 0];
    uInt8Float64Array_max_z[1] = buff[offset + 1];
    uInt8Float64Array_max_z[2] = buff[offset + 2];
    uInt8Float64Array_max_z[3] = buff[offset + 3];
    uInt8Float64Array_max_z[4] = buff[offset + 4];
    uInt8Float64Array_max_z[5] = buff[offset + 5];
    uInt8Float64Array_max_z[6] = buff[offset + 6];
    uInt8Float64Array_max_z[7] = buff[offset + 7];
    this.max_z = float64Array_max_z[0];
    offset += 8;
    return offset;
};

ProjectedMapInfo.prototype.serializedLength = function() {
    var length = 0;
    var encoder_frame_id = new TextEncoder('utf8');
    var utf8array_frame_id = encoder_frame_id.encode(this.frame_id);
    length += 4;
    length += utf8array_frame_id.length;
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    return length;
};

ProjectedMapInfo.prototype.echo = function() { return ""; };

ProjectedMapInfo.prototype.getType = function() { return "map_msgs/ProjectedMapInfo"; };

ProjectedMapInfo.prototype.getMD5 = function() { return "f661365637fb759e63cb5d179a4461e1"; };

ProjectedMapInfo.prototype.getID = function() { return 0; };

ProjectedMapInfo.prototype.setID = function(id) { };

return function() { return new ProjectedMapInfo(); };

});
