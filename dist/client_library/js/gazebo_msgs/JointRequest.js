(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.JointRequestRequest=f();}})(function(){var define,module,exports;'use strict';

function JointRequestRequest() {
    this.__id__ = 0;
    this.joint_name = "";
};

JointRequestRequest.prototype.serialize = function(buff, idx) {
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

JointRequestRequest.prototype.deserialize = function(buff, idx) {
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

JointRequestRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_joint_name = new TextEncoder('utf8');
    var utf8array_joint_name = encoder_joint_name.encode(this.joint_name);
    length += 4;
    length += utf8array_joint_name.length;
    return length;
};

JointRequestRequest.prototype.echo = function() { return ""; };

JointRequestRequest.prototype.getType = function() { return "gazebo_msgs/JointRequest"; };

JointRequestRequest.prototype.getMD5 = function() { return "e0bdc37edb92be07f3069573364a169f"; };

JointRequestRequest.prototype.getID = function() { return this.__id__; };

JointRequestRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new JointRequestRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.JointRequestResponse=f();}})(function(){var define,module,exports;'use strict';

function JointRequestResponse() {
    this.__id__ = 0;
};

JointRequestResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

JointRequestResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

JointRequestResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

JointRequestResponse.prototype.echo = function() { return ""; };

JointRequestResponse.prototype.getType = function() { return "gazebo_msgs/JointRequest"; };

JointRequestResponse.prototype.getMD5 = function() { return "ac5876a62df51a76c2828bb62894779d"; };

JointRequestResponse.prototype.getID = function() { return this.__id__; };

JointRequestResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new JointRequestResponse(); };

});
