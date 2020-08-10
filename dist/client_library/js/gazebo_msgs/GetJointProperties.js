(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.GetJointPropertiesRequest=f();}})(function(){var define,module,exports;'use strict';

function GetJointPropertiesRequest() {
    this.__id__ = 0;
    this.joint_name = "";
};

GetJointPropertiesRequest.prototype.serialize = function(buff, idx) {
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
    return offset;
};

GetJointPropertiesRequest.prototype.deserialize = function(buff, idx) {
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
    return offset;
};

GetJointPropertiesRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_joint_name = new TextEncoder('utf8');
    var utf8array_joint_name = encoder_joint_name.encode(this.joint_name);
    length += 4;
    length += utf8array_joint_name.length;
    return length;
};

GetJointPropertiesRequest.prototype.echo = function() { return ""; };

GetJointPropertiesRequest.prototype.getType = function() { return "gazebo_msgs/GetJointProperties"; };

GetJointPropertiesRequest.prototype.getMD5 = function() { return "b07a3ce5fb5aba1cfc56577c9cb3ecc6"; };

GetJointPropertiesRequest.prototype.getID = function() { return this.__id__; };

GetJointPropertiesRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetJointPropertiesRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.GetJointPropertiesResponse=f();}})(function(){var define,module,exports;'use strict';

function GetJointPropertiesResponse() {
    this.__id__ = 0;
    this.type = 0;
    this.damping = new Array();
    this.position = new Array();
    this.rate = new Array();
    this.success = false;
    this.status_message = "";

    // ENUM{
    this.REVOLUTE =  0                ;
    this.CONTINUOUS =  1                ;
    this.PRISMATIC =  2                ;
    this.FIXED =  3                ;
    this.BALL =  4                ;
    this.UNIVERSAL =  5                ;
    // }ENUM
};

GetJointPropertiesResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.type) >> (8 * 0)) & 0xFF;
    offset += 1;
    var length_damping = this.damping.length;
    buff[offset + 0] = (length_damping >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_damping >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_damping >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_damping >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_damping; i++) {
        var float64Array_dampingi = new Float64Array(1);
        var uInt8Float64Array_dampingi = new Uint8Array(float64Array_dampingi.buffer);
        float64Array_dampingi[0] = +this.damping[i];
        buff[offset + 0] = uInt8Float64Array_dampingi[0];
        buff[offset + 1] = uInt8Float64Array_dampingi[1];
        buff[offset + 2] = uInt8Float64Array_dampingi[2];
        buff[offset + 3] = uInt8Float64Array_dampingi[3];
        buff[offset + 4] = uInt8Float64Array_dampingi[4];
        buff[offset + 5] = uInt8Float64Array_dampingi[5];
        buff[offset + 6] = uInt8Float64Array_dampingi[6];
        buff[offset + 7] = uInt8Float64Array_dampingi[7];
        offset += 8;
    }
    var length_position = this.position.length;
    buff[offset + 0] = (length_position >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_position >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_position >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_position >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_position; i++) {
        var float64Array_positioni = new Float64Array(1);
        var uInt8Float64Array_positioni = new Uint8Array(float64Array_positioni.buffer);
        float64Array_positioni[0] = +this.position[i];
        buff[offset + 0] = uInt8Float64Array_positioni[0];
        buff[offset + 1] = uInt8Float64Array_positioni[1];
        buff[offset + 2] = uInt8Float64Array_positioni[2];
        buff[offset + 3] = uInt8Float64Array_positioni[3];
        buff[offset + 4] = uInt8Float64Array_positioni[4];
        buff[offset + 5] = uInt8Float64Array_positioni[5];
        buff[offset + 6] = uInt8Float64Array_positioni[6];
        buff[offset + 7] = uInt8Float64Array_positioni[7];
        offset += 8;
    }
    var length_rate = this.rate.length;
    buff[offset + 0] = (length_rate >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_rate >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_rate >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_rate >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_rate; i++) {
        var float64Array_ratei = new Float64Array(1);
        var uInt8Float64Array_ratei = new Uint8Array(float64Array_ratei.buffer);
        float64Array_ratei[0] = +this.rate[i];
        buff[offset + 0] = uInt8Float64Array_ratei[0];
        buff[offset + 1] = uInt8Float64Array_ratei[1];
        buff[offset + 2] = uInt8Float64Array_ratei[2];
        buff[offset + 3] = uInt8Float64Array_ratei[3];
        buff[offset + 4] = uInt8Float64Array_ratei[4];
        buff[offset + 5] = uInt8Float64Array_ratei[5];
        buff[offset + 6] = uInt8Float64Array_ratei[6];
        buff[offset + 7] = uInt8Float64Array_ratei[7];
        offset += 8;
    }
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

GetJointPropertiesResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.type = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    var length_damping = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_damping |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_damping |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_damping |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.damping = new Array(length_damping);
    for (var i = 0; i < length_damping; i++) {
        var float64Array_dampingi = new Float64Array(1);
        var uInt8Float64Array_dampingi = new Uint8Array(float64Array_dampingi.buffer);
        uInt8Float64Array_dampingi[0] = buff[offset + 0];
        uInt8Float64Array_dampingi[1] = buff[offset + 1];
        uInt8Float64Array_dampingi[2] = buff[offset + 2];
        uInt8Float64Array_dampingi[3] = buff[offset + 3];
        uInt8Float64Array_dampingi[4] = buff[offset + 4];
        uInt8Float64Array_dampingi[5] = buff[offset + 5];
        uInt8Float64Array_dampingi[6] = buff[offset + 6];
        uInt8Float64Array_dampingi[7] = buff[offset + 7];
        this.damping[i] = float64Array_dampingi[0];
        offset += 8;
    }
    var length_position = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_position |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_position |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_position |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.position = new Array(length_position);
    for (var i = 0; i < length_position; i++) {
        var float64Array_positioni = new Float64Array(1);
        var uInt8Float64Array_positioni = new Uint8Array(float64Array_positioni.buffer);
        uInt8Float64Array_positioni[0] = buff[offset + 0];
        uInt8Float64Array_positioni[1] = buff[offset + 1];
        uInt8Float64Array_positioni[2] = buff[offset + 2];
        uInt8Float64Array_positioni[3] = buff[offset + 3];
        uInt8Float64Array_positioni[4] = buff[offset + 4];
        uInt8Float64Array_positioni[5] = buff[offset + 5];
        uInt8Float64Array_positioni[6] = buff[offset + 6];
        uInt8Float64Array_positioni[7] = buff[offset + 7];
        this.position[i] = float64Array_positioni[0];
        offset += 8;
    }
    var length_rate = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_rate |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_rate |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_rate |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.rate = new Array(length_rate);
    for (var i = 0; i < length_rate; i++) {
        var float64Array_ratei = new Float64Array(1);
        var uInt8Float64Array_ratei = new Uint8Array(float64Array_ratei.buffer);
        uInt8Float64Array_ratei[0] = buff[offset + 0];
        uInt8Float64Array_ratei[1] = buff[offset + 1];
        uInt8Float64Array_ratei[2] = buff[offset + 2];
        uInt8Float64Array_ratei[3] = buff[offset + 3];
        uInt8Float64Array_ratei[4] = buff[offset + 4];
        uInt8Float64Array_ratei[5] = buff[offset + 5];
        uInt8Float64Array_ratei[6] = buff[offset + 6];
        uInt8Float64Array_ratei[7] = buff[offset + 7];
        this.rate[i] = float64Array_ratei[0];
        offset += 8;
    }
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

GetJointPropertiesResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    var length_damping = this.damping.length;
    length += 4;
    for (var i = 0; i < length_damping; i++) {
        length += 8
    }
    var length_position = this.position.length;
    length += 4;
    for (var i = 0; i < length_position; i++) {
        length += 8
    }
    var length_rate = this.rate.length;
    length += 4;
    for (var i = 0; i < length_rate; i++) {
        length += 8
    }
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

GetJointPropertiesResponse.prototype.echo = function() { return ""; };

GetJointPropertiesResponse.prototype.getType = function() { return "gazebo_msgs/GetJointProperties"; };

GetJointPropertiesResponse.prototype.getMD5 = function() { return "a60fbf691ac539e1355c979ca09b4573"; };

GetJointPropertiesResponse.prototype.getID = function() { return this.__id__; };

GetJointPropertiesResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetJointPropertiesResponse(); };

});
