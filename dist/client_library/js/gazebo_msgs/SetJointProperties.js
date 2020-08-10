(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SetJointPropertiesRequest=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/gazebo_msgs/ODEJointProperties.js'></script>");

function SetJointPropertiesRequest() {
    this.__id__ = 0;
    this.joint_name = "";
    this.ode_joint_config = gazebo_msgs.ODEJointProperties();
};

SetJointPropertiesRequest.prototype.serialize = function(buff, idx) {
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
    offset += this.ode_joint_config.serialize(buff, offset);
    return offset;
};

SetJointPropertiesRequest.prototype.deserialize = function(buff, idx) {
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
    offset += this.ode_joint_config.deserialize(buff, offset);
    return offset;
};

SetJointPropertiesRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_joint_name = new TextEncoder('utf8');
    var utf8array_joint_name = encoder_joint_name.encode(this.joint_name);
    length += 4;
    length += utf8array_joint_name.length;
    length += this.ode_joint_config.serializedLength();
    return length;
};

SetJointPropertiesRequest.prototype.echo = function() { return ""; };

SetJointPropertiesRequest.prototype.getType = function() { return "gazebo_msgs/SetJointProperties"; };

SetJointPropertiesRequest.prototype.getMD5 = function() { return "e9063603bda4e7bdd2b5530ad7705661"; };

SetJointPropertiesRequest.prototype.getID = function() { return this.__id__; };

SetJointPropertiesRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetJointPropertiesRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SetJointPropertiesResponse=f();}})(function(){var define,module,exports;'use strict';

function SetJointPropertiesResponse() {
    this.__id__ = 0;
    this.success = false;
    this.status_message = "";
};

SetJointPropertiesResponse.prototype.serialize = function(buff, idx) {
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

SetJointPropertiesResponse.prototype.deserialize = function(buff, idx) {
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

SetJointPropertiesResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

SetJointPropertiesResponse.prototype.echo = function() { return ""; };

SetJointPropertiesResponse.prototype.getType = function() { return "gazebo_msgs/SetJointProperties"; };

SetJointPropertiesResponse.prototype.getMD5 = function() { return "7e8650b70fd2dfc24b249dddf8f45cee"; };

SetJointPropertiesResponse.prototype.getID = function() { return this.__id__; };

SetJointPropertiesResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetJointPropertiesResponse(); };

});
