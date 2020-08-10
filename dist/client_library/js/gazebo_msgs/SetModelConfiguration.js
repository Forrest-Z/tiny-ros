(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SetModelConfigurationRequest=f();}})(function(){var define,module,exports;'use strict';

function SetModelConfigurationRequest() {
    this.__id__ = 0;
    this.model_name = "";
    this.urdf_param_name = "";
    this.joint_names = new Array();
    this.joint_positions = new Array();
};

SetModelConfigurationRequest.prototype.serialize = function(buff, idx) {
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
    var encoder_urdf_param_name = new TextEncoder('utf8');
    var utf8array_urdf_param_name = encoder_urdf_param_name.encode(this.urdf_param_name);
    buff[offset + 0] = (utf8array_urdf_param_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_urdf_param_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_urdf_param_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_urdf_param_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_urdf_param_name.length; i++) {
        buff[offset + i] = utf8array_urdf_param_name[i];
    }
    offset += utf8array_urdf_param_name.length;
    var length_joint_names = this.joint_names.length;
    buff[offset + 0] = (length_joint_names >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_joint_names >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_joint_names >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_joint_names >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_joint_names; i++) {
        var encoder_joint_namesi = new TextEncoder('utf8');
        var utf8array_joint_namesi = encoder_joint_namesi.encode(this.joint_names[i]);
        buff[offset + 0] = (utf8array_joint_namesi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_joint_namesi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_joint_namesi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_joint_namesi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_joint_namesi.length; i++) {
            buff[offset + i] = utf8array_joint_namesi[i];
        }
        offset += utf8array_joint_namesi.length;
    }
    var length_joint_positions = this.joint_positions.length;
    buff[offset + 0] = (length_joint_positions >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_joint_positions >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_joint_positions >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_joint_positions >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_joint_positions; i++) {
        var float64Array_joint_positionsi = new Float64Array(1);
        var uInt8Float64Array_joint_positionsi = new Uint8Array(float64Array_joint_positionsi.buffer);
        float64Array_joint_positionsi[0] = +this.joint_positions[i];
        buff[offset + 0] = uInt8Float64Array_joint_positionsi[0];
        buff[offset + 1] = uInt8Float64Array_joint_positionsi[1];
        buff[offset + 2] = uInt8Float64Array_joint_positionsi[2];
        buff[offset + 3] = uInt8Float64Array_joint_positionsi[3];
        buff[offset + 4] = uInt8Float64Array_joint_positionsi[4];
        buff[offset + 5] = uInt8Float64Array_joint_positionsi[5];
        buff[offset + 6] = uInt8Float64Array_joint_positionsi[6];
        buff[offset + 7] = uInt8Float64Array_joint_positionsi[7];
        offset += 8;
    }
    return offset;
};

SetModelConfigurationRequest.prototype.deserialize = function(buff, idx) {
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
    var length_urdf_param_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_urdf_param_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_urdf_param_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_urdf_param_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_urdf_param_name = new TextDecoder('utf8');
    this.urdf_param_name = decoder_urdf_param_name.decode(buff.slice(offset, offset + length_urdf_param_name));
    offset += length_urdf_param_name;
    var length_joint_names = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_joint_names |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_joint_names |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_joint_names |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.joint_names = new Array(length_joint_names);
    for (var i = 0; i < length_joint_names; i++) {
        var length_joint_namesi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_joint_namesi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_joint_namesi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_joint_namesi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_joint_namesi = new TextDecoder('utf8');
        this.joint_names[i] = decoder_joint_namesi.decode(buff.slice(offset, offset + length_joint_namesi));
        offset += length_joint_namesi;
    }
    var length_joint_positions = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_joint_positions |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_joint_positions |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_joint_positions |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.joint_positions = new Array(length_joint_positions);
    for (var i = 0; i < length_joint_positions; i++) {
        var float64Array_joint_positionsi = new Float64Array(1);
        var uInt8Float64Array_joint_positionsi = new Uint8Array(float64Array_joint_positionsi.buffer);
        uInt8Float64Array_joint_positionsi[0] = buff[offset + 0];
        uInt8Float64Array_joint_positionsi[1] = buff[offset + 1];
        uInt8Float64Array_joint_positionsi[2] = buff[offset + 2];
        uInt8Float64Array_joint_positionsi[3] = buff[offset + 3];
        uInt8Float64Array_joint_positionsi[4] = buff[offset + 4];
        uInt8Float64Array_joint_positionsi[5] = buff[offset + 5];
        uInt8Float64Array_joint_positionsi[6] = buff[offset + 6];
        uInt8Float64Array_joint_positionsi[7] = buff[offset + 7];
        this.joint_positions[i] = float64Array_joint_positionsi[0];
        offset += 8;
    }
    return offset;
};

SetModelConfigurationRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_model_name = new TextEncoder('utf8');
    var utf8array_model_name = encoder_model_name.encode(this.model_name);
    length += 4;
    length += utf8array_model_name.length;
    var encoder_urdf_param_name = new TextEncoder('utf8');
    var utf8array_urdf_param_name = encoder_urdf_param_name.encode(this.urdf_param_name);
    length += 4;
    length += utf8array_urdf_param_name.length;
    var length_joint_names = this.joint_names.length;
    length += 4;
    for (var i = 0; i < length_joint_names; i++) {
        var encoder_joint_namesi = new TextEncoder('utf8');
        var utf8array_joint_namesi = encoder_joint_namesi.encode(this.joint_names[i]);
        length += 4;
        length += utf8array_joint_namesi.length;
    }
    var length_joint_positions = this.joint_positions.length;
    length += 4;
    for (var i = 0; i < length_joint_positions; i++) {
        length += 8
    }
    return length;
};

SetModelConfigurationRequest.prototype.echo = function() { return ""; };

SetModelConfigurationRequest.prototype.getType = function() { return "gazebo_msgs/SetModelConfiguration"; };

SetModelConfigurationRequest.prototype.getMD5 = function() { return "74db6184ae83468b540d4c02d244ada7"; };

SetModelConfigurationRequest.prototype.getID = function() { return this.__id__; };

SetModelConfigurationRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetModelConfigurationRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SetModelConfigurationResponse=f();}})(function(){var define,module,exports;'use strict';

function SetModelConfigurationResponse() {
    this.__id__ = 0;
    this.success = false;
    this.status_message = "";
};

SetModelConfigurationResponse.prototype.serialize = function(buff, idx) {
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

SetModelConfigurationResponse.prototype.deserialize = function(buff, idx) {
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

SetModelConfigurationResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

SetModelConfigurationResponse.prototype.echo = function() { return ""; };

SetModelConfigurationResponse.prototype.getType = function() { return "gazebo_msgs/SetModelConfiguration"; };

SetModelConfigurationResponse.prototype.getMD5 = function() { return "6f12aefa315c8b37040d5d47471e39ee"; };

SetModelConfigurationResponse.prototype.getID = function() { return this.__id__; };

SetModelConfigurationResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetModelConfigurationResponse(); };

});
