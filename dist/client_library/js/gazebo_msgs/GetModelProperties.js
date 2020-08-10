(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.GetModelPropertiesRequest=f();}})(function(){var define,module,exports;'use strict';

function GetModelPropertiesRequest() {
    this.__id__ = 0;
    this.model_name = "";
};

GetModelPropertiesRequest.prototype.serialize = function(buff, idx) {
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
    return offset;
};

GetModelPropertiesRequest.prototype.deserialize = function(buff, idx) {
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
    return offset;
};

GetModelPropertiesRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_model_name = new TextEncoder('utf8');
    var utf8array_model_name = encoder_model_name.encode(this.model_name);
    length += 4;
    length += utf8array_model_name.length;
    return length;
};

GetModelPropertiesRequest.prototype.echo = function() { return ""; };

GetModelPropertiesRequest.prototype.getType = function() { return "gazebo_msgs/GetModelProperties"; };

GetModelPropertiesRequest.prototype.getMD5 = function() { return "fe0194bf75c917c89b820b09c12fe6c1"; };

GetModelPropertiesRequest.prototype.getID = function() { return this.__id__; };

GetModelPropertiesRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetModelPropertiesRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.GetModelPropertiesResponse=f();}})(function(){var define,module,exports;'use strict';

function GetModelPropertiesResponse() {
    this.__id__ = 0;
    this.parent_model_name = "";
    this.canonical_body_name = "";
    this.body_names = new Array();
    this.geom_names = new Array();
    this.joint_names = new Array();
    this.child_model_names = new Array();
    this.is_static = false;
    this.success = false;
    this.status_message = "";
};

GetModelPropertiesResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_parent_model_name = new TextEncoder('utf8');
    var utf8array_parent_model_name = encoder_parent_model_name.encode(this.parent_model_name);
    buff[offset + 0] = (utf8array_parent_model_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_parent_model_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_parent_model_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_parent_model_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_parent_model_name.length; i++) {
        buff[offset + i] = utf8array_parent_model_name[i];
    }
    offset += utf8array_parent_model_name.length;
    var encoder_canonical_body_name = new TextEncoder('utf8');
    var utf8array_canonical_body_name = encoder_canonical_body_name.encode(this.canonical_body_name);
    buff[offset + 0] = (utf8array_canonical_body_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_canonical_body_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_canonical_body_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_canonical_body_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_canonical_body_name.length; i++) {
        buff[offset + i] = utf8array_canonical_body_name[i];
    }
    offset += utf8array_canonical_body_name.length;
    var length_body_names = this.body_names.length;
    buff[offset + 0] = (length_body_names >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_body_names >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_body_names >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_body_names >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_body_names; i++) {
        var encoder_body_namesi = new TextEncoder('utf8');
        var utf8array_body_namesi = encoder_body_namesi.encode(this.body_names[i]);
        buff[offset + 0] = (utf8array_body_namesi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_body_namesi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_body_namesi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_body_namesi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_body_namesi.length; i++) {
            buff[offset + i] = utf8array_body_namesi[i];
        }
        offset += utf8array_body_namesi.length;
    }
    var length_geom_names = this.geom_names.length;
    buff[offset + 0] = (length_geom_names >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_geom_names >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_geom_names >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_geom_names >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_geom_names; i++) {
        var encoder_geom_namesi = new TextEncoder('utf8');
        var utf8array_geom_namesi = encoder_geom_namesi.encode(this.geom_names[i]);
        buff[offset + 0] = (utf8array_geom_namesi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_geom_namesi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_geom_namesi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_geom_namesi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_geom_namesi.length; i++) {
            buff[offset + i] = utf8array_geom_namesi[i];
        }
        offset += utf8array_geom_namesi.length;
    }
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
    var length_child_model_names = this.child_model_names.length;
    buff[offset + 0] = (length_child_model_names >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_child_model_names >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_child_model_names >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_child_model_names >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_child_model_names; i++) {
        var encoder_child_model_namesi = new TextEncoder('utf8');
        var utf8array_child_model_namesi = encoder_child_model_namesi.encode(this.child_model_names[i]);
        buff[offset + 0] = (utf8array_child_model_namesi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_child_model_namesi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_child_model_namesi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_child_model_namesi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_child_model_namesi.length; i++) {
            buff[offset + i] = utf8array_child_model_namesi[i];
        }
        offset += utf8array_child_model_namesi.length;
    }
    buff[offset] = this.is_static === false ? 0 : 1;
    offset += 1;
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

GetModelPropertiesResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_parent_model_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_parent_model_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_parent_model_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_parent_model_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_parent_model_name = new TextDecoder('utf8');
    this.parent_model_name = decoder_parent_model_name.decode(buff.slice(offset, offset + length_parent_model_name));
    offset += length_parent_model_name;
    var length_canonical_body_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_canonical_body_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_canonical_body_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_canonical_body_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_canonical_body_name = new TextDecoder('utf8');
    this.canonical_body_name = decoder_canonical_body_name.decode(buff.slice(offset, offset + length_canonical_body_name));
    offset += length_canonical_body_name;
    var length_body_names = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_body_names |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_body_names |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_body_names |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.body_names = new Array(length_body_names);
    for (var i = 0; i < length_body_names; i++) {
        var length_body_namesi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_body_namesi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_body_namesi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_body_namesi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_body_namesi = new TextDecoder('utf8');
        this.body_names[i] = decoder_body_namesi.decode(buff.slice(offset, offset + length_body_namesi));
        offset += length_body_namesi;
    }
    var length_geom_names = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_geom_names |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_geom_names |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_geom_names |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.geom_names = new Array(length_geom_names);
    for (var i = 0; i < length_geom_names; i++) {
        var length_geom_namesi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_geom_namesi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_geom_namesi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_geom_namesi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_geom_namesi = new TextDecoder('utf8');
        this.geom_names[i] = decoder_geom_namesi.decode(buff.slice(offset, offset + length_geom_namesi));
        offset += length_geom_namesi;
    }
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
    var length_child_model_names = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_child_model_names |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_child_model_names |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_child_model_names |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.child_model_names = new Array(length_child_model_names);
    for (var i = 0; i < length_child_model_names; i++) {
        var length_child_model_namesi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_child_model_namesi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_child_model_namesi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_child_model_namesi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_child_model_namesi = new TextDecoder('utf8');
        this.child_model_names[i] = decoder_child_model_namesi.decode(buff.slice(offset, offset + length_child_model_namesi));
        offset += length_child_model_namesi;
    }
    this.is_static = buff[offset] !== 0 ? true : false;
    offset += 1;
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

GetModelPropertiesResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_parent_model_name = new TextEncoder('utf8');
    var utf8array_parent_model_name = encoder_parent_model_name.encode(this.parent_model_name);
    length += 4;
    length += utf8array_parent_model_name.length;
    var encoder_canonical_body_name = new TextEncoder('utf8');
    var utf8array_canonical_body_name = encoder_canonical_body_name.encode(this.canonical_body_name);
    length += 4;
    length += utf8array_canonical_body_name.length;
    var length_body_names = this.body_names.length;
    length += 4;
    for (var i = 0; i < length_body_names; i++) {
        var encoder_body_namesi = new TextEncoder('utf8');
        var utf8array_body_namesi = encoder_body_namesi.encode(this.body_names[i]);
        length += 4;
        length += utf8array_body_namesi.length;
    }
    var length_geom_names = this.geom_names.length;
    length += 4;
    for (var i = 0; i < length_geom_names; i++) {
        var encoder_geom_namesi = new TextEncoder('utf8');
        var utf8array_geom_namesi = encoder_geom_namesi.encode(this.geom_names[i]);
        length += 4;
        length += utf8array_geom_namesi.length;
    }
    var length_joint_names = this.joint_names.length;
    length += 4;
    for (var i = 0; i < length_joint_names; i++) {
        var encoder_joint_namesi = new TextEncoder('utf8');
        var utf8array_joint_namesi = encoder_joint_namesi.encode(this.joint_names[i]);
        length += 4;
        length += utf8array_joint_namesi.length;
    }
    var length_child_model_names = this.child_model_names.length;
    length += 4;
    for (var i = 0; i < length_child_model_names; i++) {
        var encoder_child_model_namesi = new TextEncoder('utf8');
        var utf8array_child_model_namesi = encoder_child_model_namesi.encode(this.child_model_names[i]);
        length += 4;
        length += utf8array_child_model_namesi.length;
    }
    length += 1
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

GetModelPropertiesResponse.prototype.echo = function() { return ""; };

GetModelPropertiesResponse.prototype.getType = function() { return "gazebo_msgs/GetModelProperties"; };

GetModelPropertiesResponse.prototype.getMD5 = function() { return "d8f16b08abaf0220a551cf9025748602"; };

GetModelPropertiesResponse.prototype.getID = function() { return this.__id__; };

GetModelPropertiesResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetModelPropertiesResponse(); };

});
