(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf2_msgs==="undefined"){g.tf2_msgs={};}g.tf2_msgs.LookupTransformGoal=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Time.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Duration.js'></script>");

function LookupTransformGoal() {
    this.target_frame = "";
    this.source_frame = "";
    this.source_time = tinyros.Time();
    this.timeout = tinyros.Duration();
    this.target_time = tinyros.Time();
    this.fixed_frame = "";
    this.advanced = false;
};

LookupTransformGoal.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var encoder_target_frame = new TextEncoder('utf8');
    var utf8array_target_frame = encoder_target_frame.encode(this.target_frame);
    buff[offset + 0] = (utf8array_target_frame.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_target_frame.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_target_frame.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_target_frame.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_target_frame.length; i++) {
        buff[offset + i] = utf8array_target_frame[i];
    }
    offset += utf8array_target_frame.length;
    var encoder_source_frame = new TextEncoder('utf8');
    var utf8array_source_frame = encoder_source_frame.encode(this.source_frame);
    buff[offset + 0] = (utf8array_source_frame.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_source_frame.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_source_frame.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_source_frame.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_source_frame.length; i++) {
        buff[offset + i] = utf8array_source_frame[i];
    }
    offset += utf8array_source_frame.length;
    buff[offset + 0] = ((+this.source_time.sec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.source_time.sec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.source_time.sec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.source_time.sec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.source_time.nsec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.source_time.nsec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.source_time.nsec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.source_time.nsec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.timeout.sec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.timeout.sec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.timeout.sec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.timeout.sec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.timeout.nsec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.timeout.nsec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.timeout.nsec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.timeout.nsec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.target_time.sec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.target_time.sec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.target_time.sec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.target_time.sec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.target_time.nsec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.target_time.nsec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.target_time.nsec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.target_time.nsec) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_fixed_frame = new TextEncoder('utf8');
    var utf8array_fixed_frame = encoder_fixed_frame.encode(this.fixed_frame);
    buff[offset + 0] = (utf8array_fixed_frame.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_fixed_frame.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_fixed_frame.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_fixed_frame.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_fixed_frame.length; i++) {
        buff[offset + i] = utf8array_fixed_frame[i];
    }
    offset += utf8array_fixed_frame.length;
    buff[offset] = this.advanced === false ? 0 : 1;
    offset += 1;
    return offset;
};

LookupTransformGoal.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_target_frame = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_target_frame |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_target_frame |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_target_frame |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_target_frame = new TextDecoder('utf8');
    this.target_frame = decoder_target_frame.decode(buff.slice(offset, offset + length_target_frame));
    offset += length_target_frame;
    var length_source_frame = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_source_frame |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_source_frame |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_source_frame |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_source_frame = new TextDecoder('utf8');
    this.source_frame = decoder_source_frame.decode(buff.slice(offset, offset + length_source_frame));
    offset += length_source_frame;
    this.source_time.sec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.source_time.sec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.source_time.sec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.source_time.sec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.source_time.nsec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.source_time.nsec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.source_time.nsec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.source_time.nsec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.timeout.sec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.timeout.sec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.timeout.sec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.timeout.sec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.timeout.nsec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.timeout.nsec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.timeout.nsec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.timeout.nsec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.target_time.sec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.target_time.sec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.target_time.sec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.target_time.sec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.target_time.nsec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.target_time.nsec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.target_time.nsec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.target_time.nsec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_fixed_frame = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_fixed_frame |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_fixed_frame |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_fixed_frame |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_fixed_frame = new TextDecoder('utf8');
    this.fixed_frame = decoder_fixed_frame.decode(buff.slice(offset, offset + length_fixed_frame));
    offset += length_fixed_frame;
    this.advanced = buff[offset] !== 0 ? true : false;
    offset += 1;
    return offset;
};

LookupTransformGoal.prototype.serializedLength = function() {
    var length = 0;
    var encoder_target_frame = new TextEncoder('utf8');
    var utf8array_target_frame = encoder_target_frame.encode(this.target_frame);
    length += 4;
    length += utf8array_target_frame.length;
    var encoder_source_frame = new TextEncoder('utf8');
    var utf8array_source_frame = encoder_source_frame.encode(this.source_frame);
    length += 4;
    length += utf8array_source_frame.length;
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    var encoder_fixed_frame = new TextEncoder('utf8');
    var utf8array_fixed_frame = encoder_fixed_frame.encode(this.fixed_frame);
    length += 4;
    length += utf8array_fixed_frame.length;
    length += 1
    return length;
};

LookupTransformGoal.prototype.echo = function() { return ""; };

LookupTransformGoal.prototype.getType = function() { return "tf2_msgs/LookupTransformGoal"; };

LookupTransformGoal.prototype.getMD5 = function() { return "677c84a912b788ccaaea5fbc38570182"; };

LookupTransformGoal.prototype.getID = function() { return 0; };

LookupTransformGoal.prototype.setID = function(id) { };

return function() { return new LookupTransformGoal(); };

});
