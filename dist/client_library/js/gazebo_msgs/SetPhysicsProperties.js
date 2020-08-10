(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SetPhysicsPropertiesRequest=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Vector3.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/gazebo_msgs/ODEPhysics.js'></script>");

function SetPhysicsPropertiesRequest() {
    this.__id__ = 0;
    this.time_step = 0.0;
    this.max_update_rate = 0.0;
    this.gravity = geometry_msgs.Vector3();
    this.ode_config = gazebo_msgs.ODEPhysics();
};

SetPhysicsPropertiesRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var float64Array_time_step = new Float64Array(1);
    var uInt8Float64Array_time_step = new Uint8Array(float64Array_time_step.buffer);
    float64Array_time_step[0] = +this.time_step;
    buff[offset + 0] = uInt8Float64Array_time_step[0];
    buff[offset + 1] = uInt8Float64Array_time_step[1];
    buff[offset + 2] = uInt8Float64Array_time_step[2];
    buff[offset + 3] = uInt8Float64Array_time_step[3];
    buff[offset + 4] = uInt8Float64Array_time_step[4];
    buff[offset + 5] = uInt8Float64Array_time_step[5];
    buff[offset + 6] = uInt8Float64Array_time_step[6];
    buff[offset + 7] = uInt8Float64Array_time_step[7];
    offset += 8;
    var float64Array_max_update_rate = new Float64Array(1);
    var uInt8Float64Array_max_update_rate = new Uint8Array(float64Array_max_update_rate.buffer);
    float64Array_max_update_rate[0] = +this.max_update_rate;
    buff[offset + 0] = uInt8Float64Array_max_update_rate[0];
    buff[offset + 1] = uInt8Float64Array_max_update_rate[1];
    buff[offset + 2] = uInt8Float64Array_max_update_rate[2];
    buff[offset + 3] = uInt8Float64Array_max_update_rate[3];
    buff[offset + 4] = uInt8Float64Array_max_update_rate[4];
    buff[offset + 5] = uInt8Float64Array_max_update_rate[5];
    buff[offset + 6] = uInt8Float64Array_max_update_rate[6];
    buff[offset + 7] = uInt8Float64Array_max_update_rate[7];
    offset += 8;
    offset += this.gravity.serialize(buff, offset);
    offset += this.ode_config.serialize(buff, offset);
    return offset;
};

SetPhysicsPropertiesRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var float64Array_time_step = new Float64Array(1);
    var uInt8Float64Array_time_step = new Uint8Array(float64Array_time_step.buffer);
    uInt8Float64Array_time_step[0] = buff[offset + 0];
    uInt8Float64Array_time_step[1] = buff[offset + 1];
    uInt8Float64Array_time_step[2] = buff[offset + 2];
    uInt8Float64Array_time_step[3] = buff[offset + 3];
    uInt8Float64Array_time_step[4] = buff[offset + 4];
    uInt8Float64Array_time_step[5] = buff[offset + 5];
    uInt8Float64Array_time_step[6] = buff[offset + 6];
    uInt8Float64Array_time_step[7] = buff[offset + 7];
    this.time_step = float64Array_time_step[0];
    offset += 8;
    var float64Array_max_update_rate = new Float64Array(1);
    var uInt8Float64Array_max_update_rate = new Uint8Array(float64Array_max_update_rate.buffer);
    uInt8Float64Array_max_update_rate[0] = buff[offset + 0];
    uInt8Float64Array_max_update_rate[1] = buff[offset + 1];
    uInt8Float64Array_max_update_rate[2] = buff[offset + 2];
    uInt8Float64Array_max_update_rate[3] = buff[offset + 3];
    uInt8Float64Array_max_update_rate[4] = buff[offset + 4];
    uInt8Float64Array_max_update_rate[5] = buff[offset + 5];
    uInt8Float64Array_max_update_rate[6] = buff[offset + 6];
    uInt8Float64Array_max_update_rate[7] = buff[offset + 7];
    this.max_update_rate = float64Array_max_update_rate[0];
    offset += 8;
    offset += this.gravity.deserialize(buff, offset);
    offset += this.ode_config.deserialize(buff, offset);
    return offset;
};

SetPhysicsPropertiesRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 8
    length += 8
    length += this.gravity.serializedLength();
    length += this.ode_config.serializedLength();
    return length;
};

SetPhysicsPropertiesRequest.prototype.echo = function() { return ""; };

SetPhysicsPropertiesRequest.prototype.getType = function() { return "gazebo_msgs/SetPhysicsProperties"; };

SetPhysicsPropertiesRequest.prototype.getMD5 = function() { return "373e5371b10119be0a008429a9660679"; };

SetPhysicsPropertiesRequest.prototype.getID = function() { return this.__id__; };

SetPhysicsPropertiesRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetPhysicsPropertiesRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SetPhysicsPropertiesResponse=f();}})(function(){var define,module,exports;'use strict';

function SetPhysicsPropertiesResponse() {
    this.__id__ = 0;
    this.success = false;
    this.status_message = "";
};

SetPhysicsPropertiesResponse.prototype.serialize = function(buff, idx) {
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

SetPhysicsPropertiesResponse.prototype.deserialize = function(buff, idx) {
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

SetPhysicsPropertiesResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

SetPhysicsPropertiesResponse.prototype.echo = function() { return ""; };

SetPhysicsPropertiesResponse.prototype.getType = function() { return "gazebo_msgs/SetPhysicsProperties"; };

SetPhysicsPropertiesResponse.prototype.getMD5 = function() { return "5b1d14bf828ba119319cc03e2bb3473a"; };

SetPhysicsPropertiesResponse.prototype.getID = function() { return this.__id__; };

SetPhysicsPropertiesResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetPhysicsPropertiesResponse(); };

});
