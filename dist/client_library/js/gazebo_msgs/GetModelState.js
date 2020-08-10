(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.GetModelStateRequest=f();}})(function(){var define,module,exports;'use strict';

function GetModelStateRequest() {
    this.__id__ = 0;
    this.model_name = "";
    this.relative_entity_name = "";
};

GetModelStateRequest.prototype.serialize = function(buff, idx) {
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
    var encoder_relative_entity_name = new TextEncoder('utf8');
    var utf8array_relative_entity_name = encoder_relative_entity_name.encode(this.relative_entity_name);
    buff[offset + 0] = (utf8array_relative_entity_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_relative_entity_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_relative_entity_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_relative_entity_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_relative_entity_name.length; i++) {
        buff[offset + i] = utf8array_relative_entity_name[i];
    }
    offset += utf8array_relative_entity_name.length;
    return offset;
};

GetModelStateRequest.prototype.deserialize = function(buff, idx) {
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
    var length_relative_entity_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_relative_entity_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_relative_entity_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_relative_entity_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_relative_entity_name = new TextDecoder('utf8');
    this.relative_entity_name = decoder_relative_entity_name.decode(buff.slice(offset, offset + length_relative_entity_name));
    offset += length_relative_entity_name;
    return offset;
};

GetModelStateRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_model_name = new TextEncoder('utf8');
    var utf8array_model_name = encoder_model_name.encode(this.model_name);
    length += 4;
    length += utf8array_model_name.length;
    var encoder_relative_entity_name = new TextEncoder('utf8');
    var utf8array_relative_entity_name = encoder_relative_entity_name.encode(this.relative_entity_name);
    length += 4;
    length += utf8array_relative_entity_name.length;
    return length;
};

GetModelStateRequest.prototype.echo = function() { return ""; };

GetModelStateRequest.prototype.getType = function() { return "gazebo_msgs/GetModelState"; };

GetModelStateRequest.prototype.getMD5 = function() { return "92a8c6443abf59a40e396c81c0e29d40"; };

GetModelStateRequest.prototype.getID = function() { return this.__id__; };

GetModelStateRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetModelStateRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.GetModelStateResponse=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Pose.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Twist.js'></script>");

function GetModelStateResponse() {
    this.__id__ = 0;
    this.pose = geometry_msgs.Pose();
    this.twist = geometry_msgs.Twist();
    this.success = false;
    this.status_message = "";
};

GetModelStateResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    offset += this.pose.serialize(buff, offset);
    offset += this.twist.serialize(buff, offset);
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

GetModelStateResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    offset += this.pose.deserialize(buff, offset);
    offset += this.twist.deserialize(buff, offset);
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

GetModelStateResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += this.pose.serializedLength();
    length += this.twist.serializedLength();
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

GetModelStateResponse.prototype.echo = function() { return ""; };

GetModelStateResponse.prototype.getType = function() { return "gazebo_msgs/GetModelState"; };

GetModelStateResponse.prototype.getMD5 = function() { return "3fd873975bc823929b01f7473704b974"; };

GetModelStateResponse.prototype.getID = function() { return this.__id__; };

GetModelStateResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetModelStateResponse(); };

});
