(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.ApplyBodyWrenchRequest=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Point.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Wrench.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Time.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Duration.js'></script>");

function ApplyBodyWrenchRequest() {
    this.__id__ = 0;
    this.body_name = "";
    this.reference_frame = "";
    this.reference_point = geometry_msgs.Point();
    this.wrench = geometry_msgs.Wrench();
    this.start_time = tinyros.Time();
    this.duration = tinyros.Duration();
};

ApplyBodyWrenchRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_body_name = new TextEncoder('utf8');
    var utf8array_body_name = encoder_body_name.encode(this.body_name);
    buff[offset + 0] = (utf8array_body_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_body_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_body_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_body_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_body_name.length; i++) {
        buff[offset + i] = utf8array_body_name[i];
    }
    offset += utf8array_body_name.length;
    var encoder_reference_frame = new TextEncoder('utf8');
    var utf8array_reference_frame = encoder_reference_frame.encode(this.reference_frame);
    buff[offset + 0] = (utf8array_reference_frame.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_reference_frame.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_reference_frame.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_reference_frame.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_reference_frame.length; i++) {
        buff[offset + i] = utf8array_reference_frame[i];
    }
    offset += utf8array_reference_frame.length;
    offset += this.reference_point.serialize(buff, offset);
    offset += this.wrench.serialize(buff, offset);
    buff[offset + 0] = ((+this.start_time.sec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.start_time.sec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.start_time.sec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.start_time.sec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.start_time.nsec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.start_time.nsec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.start_time.nsec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.start_time.nsec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.duration.sec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.duration.sec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.duration.sec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.duration.sec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.duration.nsec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.duration.nsec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.duration.nsec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.duration.nsec) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

ApplyBodyWrenchRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_body_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_body_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_body_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_body_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_body_name = new TextDecoder('utf8');
    this.body_name = decoder_body_name.decode(buff.slice(offset, offset + length_body_name));
    offset += length_body_name;
    var length_reference_frame = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_reference_frame |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_reference_frame |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_reference_frame |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_reference_frame = new TextDecoder('utf8');
    this.reference_frame = decoder_reference_frame.decode(buff.slice(offset, offset + length_reference_frame));
    offset += length_reference_frame;
    offset += this.reference_point.deserialize(buff, offset);
    offset += this.wrench.deserialize(buff, offset);
    this.start_time.sec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.start_time.sec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.start_time.sec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.start_time.sec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.start_time.nsec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.start_time.nsec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.start_time.nsec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.start_time.nsec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.duration.sec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.duration.sec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.duration.sec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.duration.sec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.duration.nsec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.duration.nsec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.duration.nsec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.duration.nsec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

ApplyBodyWrenchRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_body_name = new TextEncoder('utf8');
    var utf8array_body_name = encoder_body_name.encode(this.body_name);
    length += 4;
    length += utf8array_body_name.length;
    var encoder_reference_frame = new TextEncoder('utf8');
    var utf8array_reference_frame = encoder_reference_frame.encode(this.reference_frame);
    length += 4;
    length += utf8array_reference_frame.length;
    length += this.reference_point.serializedLength();
    length += this.wrench.serializedLength();
    length += 4
    length += 4
    length += 4
    length += 4
    return length;
};

ApplyBodyWrenchRequest.prototype.echo = function() { return ""; };

ApplyBodyWrenchRequest.prototype.getType = function() { return "gazebo_msgs/ApplyBodyWrench"; };

ApplyBodyWrenchRequest.prototype.getMD5 = function() { return "434adb4bdbb64c5610c7fadb31f0fa7d"; };

ApplyBodyWrenchRequest.prototype.getID = function() { return this.__id__; };

ApplyBodyWrenchRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new ApplyBodyWrenchRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.ApplyBodyWrenchResponse=f();}})(function(){var define,module,exports;'use strict';

function ApplyBodyWrenchResponse() {
    this.__id__ = 0;
    this.success = false;
    this.status_message = "";
};

ApplyBodyWrenchResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset] = this.success === false ? 0 : 1;
    offset += 1;
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    buff[offset + 0] = (utf8array_status_message.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_status_message.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_status_message.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_status_message.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_status_message.length; i++) {
        buff[offset + i] = utf8array_status_message[i];
    }
    offset += utf8array_status_message.length;
    return offset;
};

ApplyBodyWrenchResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.success = buff[offset] !== 0 ? true : false;
    offset += 1;
    var length_status_message = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_status_message |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_status_message |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_status_message |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_status_message = new TextDecoder('utf8');
    this.status_message = decoder_status_message.decode(buff.slice(offset, offset + length_status_message));
    offset += length_status_message;
    return offset;
};

ApplyBodyWrenchResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

ApplyBodyWrenchResponse.prototype.echo = function() { return ""; };

ApplyBodyWrenchResponse.prototype.getType = function() { return "gazebo_msgs/ApplyBodyWrench"; };

ApplyBodyWrenchResponse.prototype.getMD5 = function() { return "f29b3c75e7d692065eda02aae6d3a3a0"; };

ApplyBodyWrenchResponse.prototype.getID = function() { return this.__id__; };

ApplyBodyWrenchResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new ApplyBodyWrenchResponse(); };

});
