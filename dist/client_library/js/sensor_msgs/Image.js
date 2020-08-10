(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.Image=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function Image() {
    this.header = std_msgs.Header();
    this.height = 0;
    this.width = 0;
    this.encoding = "";
    this.is_bigendian = 0;
    this.step = 0;
    this.data = new Array();
};

Image.prototype.serialize = function(buff, idx) {
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
    var encoder_encoding = new TextEncoder('utf8');
    var utf8array_encoding = encoder_encoding.encode(this.encoding);
    buff[offset + 0] = (utf8array_encoding.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_encoding.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_encoding.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_encoding.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_encoding.length; i++) {
        buff[offset + i] = utf8array_encoding[i];
    }
    offset += utf8array_encoding.length;
    buff[offset + 0] = ((+this.is_bigendian) >> (8 * 0)) & 0xFF;
    offset += 1;
    buff[offset + 0] = ((+this.step) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.step) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.step) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.step) >> (8 * 3)) & 0xFF;
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
    return offset;
};

Image.prototype.deserialize = function(buff, idx) {
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
    var length_encoding = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_encoding |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_encoding |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_encoding |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_encoding = new TextDecoder('utf8');
    this.encoding = decoder_encoding.decode(buff.slice(offset, offset + length_encoding));
    offset += length_encoding;
    this.is_bigendian = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    this.step = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.step |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.step |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.step |= +((buff[offset + 3] & 0xFF) << (8 * 3));
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
    return offset;
};

Image.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 4
    length += 4
    var encoder_encoding = new TextEncoder('utf8');
    var utf8array_encoding = encoder_encoding.encode(this.encoding);
    length += 4;
    length += utf8array_encoding.length;
    length += 1
    length += 4
    var length_data = this.data.length;
    length += 4;
    for (var i = 0; i < length_data; i++) {
        length += 1
    }
    return length;
};

Image.prototype.echo = function() { return ""; };

Image.prototype.getType = function() { return "sensor_msgs/Image"; };

Image.prototype.getMD5 = function() { return "886f928dc81bf7f1496a8b452057c5b2"; };

Image.prototype.getID = function() { return 0; };

Image.prototype.setID = function(id) { };

return function() { return new Image(); };

});
