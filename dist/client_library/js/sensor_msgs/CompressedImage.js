(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.CompressedImage=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function CompressedImage() {
    this.header = std_msgs.Header();
    this.format = "";
    this.data = new Array();
};

CompressedImage.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var encoder_format = new TextEncoder('utf8');
    var utf8array_format = encoder_format.encode(this.format);
    buff[offset + 0] = (utf8array_format.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_format.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_format.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_format.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_format.length; i++) {
        buff[offset + i] = utf8array_format[i];
    }
    offset += utf8array_format.length;
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

CompressedImage.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_format = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_format |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_format |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_format |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_format = new TextDecoder('utf8');
    this.format = decoder_format.decode(buff.slice(offset, offset + length_format));
    offset += length_format;
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

CompressedImage.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var encoder_format = new TextEncoder('utf8');
    var utf8array_format = encoder_format.encode(this.format);
    length += 4;
    length += utf8array_format.length;
    var length_data = this.data.length;
    length += 4;
    for (var i = 0; i < length_data; i++) {
        length += 1
    }
    return length;
};

CompressedImage.prototype.echo = function() { return ""; };

CompressedImage.prototype.getType = function() { return "sensor_msgs/CompressedImage"; };

CompressedImage.prototype.getMD5 = function() { return "eed57d856457441995644e6294152301"; };

CompressedImage.prototype.getID = function() { return 0; };

CompressedImage.prototype.setID = function(id) { };

return function() { return new CompressedImage(); };

});
