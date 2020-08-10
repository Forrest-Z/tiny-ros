(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SpawnModelRequest=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Pose.js'></script>");

function SpawnModelRequest() {
    this.__id__ = 0;
    this.model_name = "";
    this.model_xml = "";
    this.robot_namespace = "";
    this.initial_pose = geometry_msgs.Pose();
    this.reference_frame = "";
};

SpawnModelRequest.prototype.serialize = function(buff, idx) {
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
    var encoder_model_xml = new TextEncoder('utf8');
    var utf8array_model_xml = encoder_model_xml.encode(this.model_xml);
    buff[offset + 0] = (utf8array_model_xml.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_model_xml.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_model_xml.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_model_xml.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_model_xml.length; i++) {
        buff[offset + i] = utf8array_model_xml[i];
    }
    offset += utf8array_model_xml.length;
    var encoder_robot_namespace = new TextEncoder('utf8');
    var utf8array_robot_namespace = encoder_robot_namespace.encode(this.robot_namespace);
    buff[offset + 0] = (utf8array_robot_namespace.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_robot_namespace.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_robot_namespace.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_robot_namespace.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_robot_namespace.length; i++) {
        buff[offset + i] = utf8array_robot_namespace[i];
    }
    offset += utf8array_robot_namespace.length;
    offset += this.initial_pose.serialize(buff, offset);
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
    return offset;
};

SpawnModelRequest.prototype.deserialize = function(buff, idx) {
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
    var length_model_xml = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_model_xml |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_model_xml |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_model_xml |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_model_xml = new TextDecoder('utf8');
    this.model_xml = decoder_model_xml.decode(buff.slice(offset, offset + length_model_xml));
    offset += length_model_xml;
    var length_robot_namespace = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_robot_namespace |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_robot_namespace |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_robot_namespace |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_robot_namespace = new TextDecoder('utf8');
    this.robot_namespace = decoder_robot_namespace.decode(buff.slice(offset, offset + length_robot_namespace));
    offset += length_robot_namespace;
    offset += this.initial_pose.deserialize(buff, offset);
    var length_reference_frame = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_reference_frame |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_reference_frame |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_reference_frame |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_reference_frame = new TextDecoder('utf8');
    this.reference_frame = decoder_reference_frame.decode(buff.slice(offset, offset + length_reference_frame));
    offset += length_reference_frame;
    return offset;
};

SpawnModelRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_model_name = new TextEncoder('utf8');
    var utf8array_model_name = encoder_model_name.encode(this.model_name);
    length += 4;
    length += utf8array_model_name.length;
    var encoder_model_xml = new TextEncoder('utf8');
    var utf8array_model_xml = encoder_model_xml.encode(this.model_xml);
    length += 4;
    length += utf8array_model_xml.length;
    var encoder_robot_namespace = new TextEncoder('utf8');
    var utf8array_robot_namespace = encoder_robot_namespace.encode(this.robot_namespace);
    length += 4;
    length += utf8array_robot_namespace.length;
    length += this.initial_pose.serializedLength();
    var encoder_reference_frame = new TextEncoder('utf8');
    var utf8array_reference_frame = encoder_reference_frame.encode(this.reference_frame);
    length += 4;
    length += utf8array_reference_frame.length;
    return length;
};

SpawnModelRequest.prototype.echo = function() { return ""; };

SpawnModelRequest.prototype.getType = function() { return "gazebo_msgs/SpawnModel"; };

SpawnModelRequest.prototype.getMD5 = function() { return "da34e61c8813e52ac159e5f31fbf55be"; };

SpawnModelRequest.prototype.getID = function() { return this.__id__; };

SpawnModelRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SpawnModelRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SpawnModelResponse=f();}})(function(){var define,module,exports;'use strict';

function SpawnModelResponse() {
    this.__id__ = 0;
    this.success = false;
    this.status_message = "";
};

SpawnModelResponse.prototype.serialize = function(buff, idx) {
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

SpawnModelResponse.prototype.deserialize = function(buff, idx) {
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

SpawnModelResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

SpawnModelResponse.prototype.echo = function() { return ""; };

SpawnModelResponse.prototype.getType = function() { return "gazebo_msgs/SpawnModel"; };

SpawnModelResponse.prototype.getMD5 = function() { return "d59d46cc4e5a64f978a429dd7c306d30"; };

SpawnModelResponse.prototype.getID = function() { return this.__id__; };

SpawnModelResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SpawnModelResponse(); };

});
