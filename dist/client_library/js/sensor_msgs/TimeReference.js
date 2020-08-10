(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.TimeReference=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Time.js'></script>");

function TimeReference() {
    this.header = std_msgs.Header();
    this.time_ref = tinyros.Time();
    this.source = "";
};

TimeReference.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    buff[offset + 0] = ((+this.time_ref.sec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.time_ref.sec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.time_ref.sec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.time_ref.sec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.time_ref.nsec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.time_ref.nsec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.time_ref.nsec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.time_ref.nsec) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_source = new TextEncoder('utf8');
    var utf8array_source = encoder_source.encode(this.source);
    buff[offset + 0] = (utf8array_source.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_source.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_source.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_source.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_source.length; i++) {
        buff[offset + i] = utf8array_source[i];
    }
    offset += utf8array_source.length;
    return offset;
};

TimeReference.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    this.time_ref.sec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.time_ref.sec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.time_ref.sec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.time_ref.sec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.time_ref.nsec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.time_ref.nsec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.time_ref.nsec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.time_ref.nsec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_source = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_source |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_source |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_source |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_source = new TextDecoder('utf8');
    this.source = decoder_source.decode(buff.slice(offset, offset + length_source));
    offset += length_source;
    return offset;
};

TimeReference.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 4
    length += 4
    var encoder_source = new TextEncoder('utf8');
    var utf8array_source = encoder_source.encode(this.source);
    length += 4;
    length += utf8array_source.length;
    return length;
};

TimeReference.prototype.echo = function() { return ""; };

TimeReference.prototype.getType = function() { return "sensor_msgs/TimeReference"; };

TimeReference.prototype.getMD5 = function() { return "8e1576e01de57cd0d55758112f0e84ec"; };

TimeReference.prototype.getID = function() { return 0; };

TimeReference.prototype.setID = function(id) { };

return function() { return new TimeReference(); };

});
