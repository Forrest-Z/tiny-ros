(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SetJointTrajectoryRequest=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/trajectory_msgs/JointTrajectory.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Pose.js'></script>");

function SetJointTrajectoryRequest() {
    this.__id__ = 0;
    this.model_name = "";
    this.joint_trajectory = trajectory_msgs.JointTrajectory();
    this.model_pose = geometry_msgs.Pose();
    this.set_model_pose = false;
    this.disable_physics_updates = false;
};

SetJointTrajectoryRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_model_name = new TextEncoder('utf8');
    var utf8array_model_name = encoder_model_name.encode(this.model_name);
    buff[offset + 0] = (utf8array_model_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_model_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_model_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_model_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_model_name.length; i++) {
        buff[offset + i] = utf8array_model_name[i];
    }
    offset += utf8array_model_name.length;
    offset += this.joint_trajectory.serialize(buff, offset);
    offset += this.model_pose.serialize(buff, offset);
    buff[offset] = this.set_model_pose === false ? 0 : 1;
    offset += 1;
    buff[offset] = this.disable_physics_updates === false ? 0 : 1;
    offset += 1;
    return offset;
};

SetJointTrajectoryRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_model_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_model_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_model_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_model_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_model_name = new TextDecoder('utf8');
    this.model_name = decoder_model_name.decode(buff.slice(offset, offset + length_model_name));
    offset += length_model_name;
    offset += this.joint_trajectory.deserialize(buff, offset);
    offset += this.model_pose.deserialize(buff, offset);
    this.set_model_pose = buff[offset] !== 0 ? true : false;
    offset += 1;
    this.disable_physics_updates = buff[offset] !== 0 ? true : false;
    offset += 1;
    return offset;
};

SetJointTrajectoryRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_model_name = new TextEncoder('utf8');
    var utf8array_model_name = encoder_model_name.encode(this.model_name);
    length += 4;
    length += utf8array_model_name.length;
    length += this.joint_trajectory.serializedLength();
    length += this.model_pose.serializedLength();
    length += 1
    length += 1
    return length;
};

SetJointTrajectoryRequest.prototype.echo = function() { return ""; };

SetJointTrajectoryRequest.prototype.getType = function() { return "gazebo_msgs/SetJointTrajectory"; };

SetJointTrajectoryRequest.prototype.getMD5 = function() { return "8230e1fcc0295d8d21fbd5df0ccb0af3"; };

SetJointTrajectoryRequest.prototype.getID = function() { return this.__id__; };

SetJointTrajectoryRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetJointTrajectoryRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SetJointTrajectoryResponse=f();}})(function(){var define,module,exports;'use strict';

function SetJointTrajectoryResponse() {
    this.__id__ = 0;
    this.success = false;
    this.status_message = "";
};

SetJointTrajectoryResponse.prototype.serialize = function(buff, idx) {
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

SetJointTrajectoryResponse.prototype.deserialize = function(buff, idx) {
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

SetJointTrajectoryResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

SetJointTrajectoryResponse.prototype.echo = function() { return ""; };

SetJointTrajectoryResponse.prototype.getType = function() { return "gazebo_msgs/SetJointTrajectory"; };

SetJointTrajectoryResponse.prototype.getMD5 = function() { return "2f5fe47400272efd54556969ffe63e7e"; };

SetJointTrajectoryResponse.prototype.getID = function() { return this.__id__; };

SetJointTrajectoryResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetJointTrajectoryResponse(); };

});
