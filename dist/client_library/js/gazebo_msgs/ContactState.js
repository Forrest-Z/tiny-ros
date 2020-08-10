(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.ContactState=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Wrench.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Vector3.js'></script>");

function ContactState() {
    this.info = "";
    this.collision1_name = "";
    this.collision2_name = "";
    this.wrenches = new Array();
    this.total_wrench = geometry_msgs.Wrench();
    this.contact_positions = new Array();
    this.contact_normals = new Array();
    this.depths = new Array();
};

ContactState.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var encoder_info = new TextEncoder('utf8');
    var utf8array_info = encoder_info.encode(this.info);
    buff[offset + 0] = (utf8array_info.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_info.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_info.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_info.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_info.length; i++) {
        buff[offset + i] = utf8array_info[i];
    }
    offset += utf8array_info.length;
    var encoder_collision1_name = new TextEncoder('utf8');
    var utf8array_collision1_name = encoder_collision1_name.encode(this.collision1_name);
    buff[offset + 0] = (utf8array_collision1_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_collision1_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_collision1_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_collision1_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_collision1_name.length; i++) {
        buff[offset + i] = utf8array_collision1_name[i];
    }
    offset += utf8array_collision1_name.length;
    var encoder_collision2_name = new TextEncoder('utf8');
    var utf8array_collision2_name = encoder_collision2_name.encode(this.collision2_name);
    buff[offset + 0] = (utf8array_collision2_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_collision2_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_collision2_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_collision2_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_collision2_name.length; i++) {
        buff[offset + i] = utf8array_collision2_name[i];
    }
    offset += utf8array_collision2_name.length;
    var length_wrenches = this.wrenches.length;
    buff[offset + 0] = (length_wrenches >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_wrenches >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_wrenches >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_wrenches >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_wrenches; i++) {
        offset += this.wrenches[i].serialize(buff, offset);
    }
    offset += this.total_wrench.serialize(buff, offset);
    var length_contact_positions = this.contact_positions.length;
    buff[offset + 0] = (length_contact_positions >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_contact_positions >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_contact_positions >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_contact_positions >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_contact_positions; i++) {
        offset += this.contact_positions[i].serialize(buff, offset);
    }
    var length_contact_normals = this.contact_normals.length;
    buff[offset + 0] = (length_contact_normals >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_contact_normals >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_contact_normals >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_contact_normals >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_contact_normals; i++) {
        offset += this.contact_normals[i].serialize(buff, offset);
    }
    var length_depths = this.depths.length;
    buff[offset + 0] = (length_depths >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_depths >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_depths >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_depths >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_depths; i++) {
        var float64Array_depthsi = new Float64Array(1);
        var uInt8Float64Array_depthsi = new Uint8Array(float64Array_depthsi.buffer);
        float64Array_depthsi[0] = +this.depths[i];
        buff[offset + 0] = uInt8Float64Array_depthsi[0];
        buff[offset + 1] = uInt8Float64Array_depthsi[1];
        buff[offset + 2] = uInt8Float64Array_depthsi[2];
        buff[offset + 3] = uInt8Float64Array_depthsi[3];
        buff[offset + 4] = uInt8Float64Array_depthsi[4];
        buff[offset + 5] = uInt8Float64Array_depthsi[5];
        buff[offset + 6] = uInt8Float64Array_depthsi[6];
        buff[offset + 7] = uInt8Float64Array_depthsi[7];
        offset += 8;
    }
    return offset;
};

ContactState.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_info = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_info |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_info |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_info |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_info = new TextDecoder('utf8');
    this.info = decoder_info.decode(buff.slice(offset, offset + length_info));
    offset += length_info;
    var length_collision1_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_collision1_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_collision1_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_collision1_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_collision1_name = new TextDecoder('utf8');
    this.collision1_name = decoder_collision1_name.decode(buff.slice(offset, offset + length_collision1_name));
    offset += length_collision1_name;
    var length_collision2_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_collision2_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_collision2_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_collision2_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_collision2_name = new TextDecoder('utf8');
    this.collision2_name = decoder_collision2_name.decode(buff.slice(offset, offset + length_collision2_name));
    offset += length_collision2_name;
    var length_wrenches = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_wrenches |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_wrenches |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_wrenches |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.wrenches = new Array(length_wrenches);
    for (var i = 0; i < length_wrenches; i++) {
        this.wrenches[i] = geometry_msgs.Wrench();
        offset += this.wrenches[i].deserialize(buff, offset);
    }
    offset += this.total_wrench.deserialize(buff, offset);
    var length_contact_positions = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_contact_positions |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_contact_positions |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_contact_positions |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.contact_positions = new Array(length_contact_positions);
    for (var i = 0; i < length_contact_positions; i++) {
        this.contact_positions[i] = geometry_msgs.Vector3();
        offset += this.contact_positions[i].deserialize(buff, offset);
    }
    var length_contact_normals = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_contact_normals |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_contact_normals |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_contact_normals |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.contact_normals = new Array(length_contact_normals);
    for (var i = 0; i < length_contact_normals; i++) {
        this.contact_normals[i] = geometry_msgs.Vector3();
        offset += this.contact_normals[i].deserialize(buff, offset);
    }
    var length_depths = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_depths |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_depths |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_depths |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.depths = new Array(length_depths);
    for (var i = 0; i < length_depths; i++) {
        var float64Array_depthsi = new Float64Array(1);
        var uInt8Float64Array_depthsi = new Uint8Array(float64Array_depthsi.buffer);
        uInt8Float64Array_depthsi[0] = buff[offset + 0];
        uInt8Float64Array_depthsi[1] = buff[offset + 1];
        uInt8Float64Array_depthsi[2] = buff[offset + 2];
        uInt8Float64Array_depthsi[3] = buff[offset + 3];
        uInt8Float64Array_depthsi[4] = buff[offset + 4];
        uInt8Float64Array_depthsi[5] = buff[offset + 5];
        uInt8Float64Array_depthsi[6] = buff[offset + 6];
        uInt8Float64Array_depthsi[7] = buff[offset + 7];
        this.depths[i] = float64Array_depthsi[0];
        offset += 8;
    }
    return offset;
};

ContactState.prototype.serializedLength = function() {
    var length = 0;
    var encoder_info = new TextEncoder('utf8');
    var utf8array_info = encoder_info.encode(this.info);
    length += 4;
    length += utf8array_info.length;
    var encoder_collision1_name = new TextEncoder('utf8');
    var utf8array_collision1_name = encoder_collision1_name.encode(this.collision1_name);
    length += 4;
    length += utf8array_collision1_name.length;
    var encoder_collision2_name = new TextEncoder('utf8');
    var utf8array_collision2_name = encoder_collision2_name.encode(this.collision2_name);
    length += 4;
    length += utf8array_collision2_name.length;
    var length_wrenches = this.wrenches.length;
    length += 4;
    for (var i = 0; i < length_wrenches; i++) {
        length += this.wrenches[i].serializedLength();
    }
    length += this.total_wrench.serializedLength();
    var length_contact_positions = this.contact_positions.length;
    length += 4;
    for (var i = 0; i < length_contact_positions; i++) {
        length += this.contact_positions[i].serializedLength();
    }
    var length_contact_normals = this.contact_normals.length;
    length += 4;
    for (var i = 0; i < length_contact_normals; i++) {
        length += this.contact_normals[i].serializedLength();
    }
    var length_depths = this.depths.length;
    length += 4;
    for (var i = 0; i < length_depths; i++) {
        length += 8
    }
    return length;
};

ContactState.prototype.echo = function() { return ""; };

ContactState.prototype.getType = function() { return "gazebo_msgs/ContactState"; };

ContactState.prototype.getMD5 = function() { return "d82d0f0cae88aebf6b2cc86caea33a2b"; };

ContactState.prototype.getID = function() { return 0; };

ContactState.prototype.setID = function(id) { };

return function() { return new ContactState(); };

});
