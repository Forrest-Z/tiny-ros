(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.ApplyJointEffortRequest=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Time.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Duration.js'></script>");

function ApplyJointEffortRequest() {
    this.__id__ = 0;
    this.joint_name = "";
    this.effort = 0.0;
    this.start_time = tinyros.Time();
    this.duration = tinyros.Duration();
};

ApplyJointEffortRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_joint_name = new TextEncoder('utf8');
    var utf8array_joint_name = encoder_joint_name.encode(this.joint_name);
    buff[offset + 0] = (utf8array_joint_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_joint_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_joint_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_joint_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_joint_name.length; i++) {
        buff[offset + i] = utf8array_joint_name[i];
    }
    offset += utf8array_joint_name.length;
    var float64Array_effort = new Float64Array(1);
    var uInt8Float64Array_effort = new Uint8Array(float64Array_effort.buffer);
    float64Array_effort[0] = +this.effort;
    buff[offset + 0] = uInt8Float64Array_effort[0];
    buff[offset + 1] = uInt8Float64Array_effort[1];
    buff[offset + 2] = uInt8Float64Array_effort[2];
    buff[offset + 3] = uInt8Float64Array_effort[3];
    buff[offset + 4] = uInt8Float64Array_effort[4];
    buff[offset + 5] = uInt8Float64Array_effort[5];
    buff[offset + 6] = uInt8Float64Array_effort[6];
    buff[offset + 7] = uInt8Float64Array_effort[7];
    offset += 8;
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

ApplyJointEffortRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_joint_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_joint_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_joint_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_joint_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_joint_name = new TextDecoder('utf8');
    this.joint_name = decoder_joint_name.decode(buff.slice(offset, offset + length_joint_name));
    offset += length_joint_name;
    var float64Array_effort = new Float64Array(1);
    var uInt8Float64Array_effort = new Uint8Array(float64Array_effort.buffer);
    uInt8Float64Array_effort[0] = buff[offset + 0];
    uInt8Float64Array_effort[1] = buff[offset + 1];
    uInt8Float64Array_effort[2] = buff[offset + 2];
    uInt8Float64Array_effort[3] = buff[offset + 3];
    uInt8Float64Array_effort[4] = buff[offset + 4];
    uInt8Float64Array_effort[5] = buff[offset + 5];
    uInt8Float64Array_effort[6] = buff[offset + 6];
    uInt8Float64Array_effort[7] = buff[offset + 7];
    this.effort = float64Array_effort[0];
    offset += 8;
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

ApplyJointEffortRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_joint_name = new TextEncoder('utf8');
    var utf8array_joint_name = encoder_joint_name.encode(this.joint_name);
    length += 4;
    length += utf8array_joint_name.length;
    length += 8
    length += 4
    length += 4
    length += 4
    length += 4
    return length;
};

ApplyJointEffortRequest.prototype.echo = function() { return ""; };

ApplyJointEffortRequest.prototype.getType = function() { return "gazebo_msgs/ApplyJointEffort"; };

ApplyJointEffortRequest.prototype.getMD5 = function() { return "90595a46cf1fb4ee17158e2f7034a0eb"; };

ApplyJointEffortRequest.prototype.getID = function() { return this.__id__; };

ApplyJointEffortRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new ApplyJointEffortRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.ApplyJointEffortResponse=f();}})(function(){var define,module,exports;'use strict';

function ApplyJointEffortResponse() {
    this.__id__ = 0;
    this.success = false;
    this.status_message = "";
};

ApplyJointEffortResponse.prototype.serialize = function(buff, idx) {
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

ApplyJointEffortResponse.prototype.deserialize = function(buff, idx) {
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

ApplyJointEffortResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

ApplyJointEffortResponse.prototype.echo = function() { return ""; };

ApplyJointEffortResponse.prototype.getType = function() { return "gazebo_msgs/ApplyJointEffort"; };

ApplyJointEffortResponse.prototype.getMD5 = function() { return "953194fc8ca726693bef2876cebb0438"; };

ApplyJointEffortResponse.prototype.getID = function() { return this.__id__; };

ApplyJointEffortResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new ApplyJointEffortResponse(); };

});
