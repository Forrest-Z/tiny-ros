(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SetLinkPropertiesRequest=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Pose.js'></script>");

function SetLinkPropertiesRequest() {
    this.__id__ = 0;
    this.link_name = "";
    this.com = geometry_msgs.Pose();
    this.gravity_mode = false;
    this.mass = 0.0;
    this.ixx = 0.0;
    this.ixy = 0.0;
    this.ixz = 0.0;
    this.iyy = 0.0;
    this.iyz = 0.0;
    this.izz = 0.0;
};

SetLinkPropertiesRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_link_name = new TextEncoder('utf8');
    var utf8array_link_name = encoder_link_name.encode(this.link_name);
    buff[offset + 0] = (utf8array_link_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_link_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_link_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_link_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_link_name.length; i++) {
        buff[offset + i] = utf8array_link_name[i];
    }
    offset += utf8array_link_name.length;
    offset += this.com.serialize(buff, offset);
    buff[offset] = this.gravity_mode === false ? 0 : 1;
    offset += 1;
    var float64Array_mass = new Float64Array(1);
    var uInt8Float64Array_mass = new Uint8Array(float64Array_mass.buffer);
    float64Array_mass[0] = +this.mass;
    buff[offset + 0] = uInt8Float64Array_mass[0];
    buff[offset + 1] = uInt8Float64Array_mass[1];
    buff[offset + 2] = uInt8Float64Array_mass[2];
    buff[offset + 3] = uInt8Float64Array_mass[3];
    buff[offset + 4] = uInt8Float64Array_mass[4];
    buff[offset + 5] = uInt8Float64Array_mass[5];
    buff[offset + 6] = uInt8Float64Array_mass[6];
    buff[offset + 7] = uInt8Float64Array_mass[7];
    offset += 8;
    var float64Array_ixx = new Float64Array(1);
    var uInt8Float64Array_ixx = new Uint8Array(float64Array_ixx.buffer);
    float64Array_ixx[0] = +this.ixx;
    buff[offset + 0] = uInt8Float64Array_ixx[0];
    buff[offset + 1] = uInt8Float64Array_ixx[1];
    buff[offset + 2] = uInt8Float64Array_ixx[2];
    buff[offset + 3] = uInt8Float64Array_ixx[3];
    buff[offset + 4] = uInt8Float64Array_ixx[4];
    buff[offset + 5] = uInt8Float64Array_ixx[5];
    buff[offset + 6] = uInt8Float64Array_ixx[6];
    buff[offset + 7] = uInt8Float64Array_ixx[7];
    offset += 8;
    var float64Array_ixy = new Float64Array(1);
    var uInt8Float64Array_ixy = new Uint8Array(float64Array_ixy.buffer);
    float64Array_ixy[0] = +this.ixy;
    buff[offset + 0] = uInt8Float64Array_ixy[0];
    buff[offset + 1] = uInt8Float64Array_ixy[1];
    buff[offset + 2] = uInt8Float64Array_ixy[2];
    buff[offset + 3] = uInt8Float64Array_ixy[3];
    buff[offset + 4] = uInt8Float64Array_ixy[4];
    buff[offset + 5] = uInt8Float64Array_ixy[5];
    buff[offset + 6] = uInt8Float64Array_ixy[6];
    buff[offset + 7] = uInt8Float64Array_ixy[7];
    offset += 8;
    var float64Array_ixz = new Float64Array(1);
    var uInt8Float64Array_ixz = new Uint8Array(float64Array_ixz.buffer);
    float64Array_ixz[0] = +this.ixz;
    buff[offset + 0] = uInt8Float64Array_ixz[0];
    buff[offset + 1] = uInt8Float64Array_ixz[1];
    buff[offset + 2] = uInt8Float64Array_ixz[2];
    buff[offset + 3] = uInt8Float64Array_ixz[3];
    buff[offset + 4] = uInt8Float64Array_ixz[4];
    buff[offset + 5] = uInt8Float64Array_ixz[5];
    buff[offset + 6] = uInt8Float64Array_ixz[6];
    buff[offset + 7] = uInt8Float64Array_ixz[7];
    offset += 8;
    var float64Array_iyy = new Float64Array(1);
    var uInt8Float64Array_iyy = new Uint8Array(float64Array_iyy.buffer);
    float64Array_iyy[0] = +this.iyy;
    buff[offset + 0] = uInt8Float64Array_iyy[0];
    buff[offset + 1] = uInt8Float64Array_iyy[1];
    buff[offset + 2] = uInt8Float64Array_iyy[2];
    buff[offset + 3] = uInt8Float64Array_iyy[3];
    buff[offset + 4] = uInt8Float64Array_iyy[4];
    buff[offset + 5] = uInt8Float64Array_iyy[5];
    buff[offset + 6] = uInt8Float64Array_iyy[6];
    buff[offset + 7] = uInt8Float64Array_iyy[7];
    offset += 8;
    var float64Array_iyz = new Float64Array(1);
    var uInt8Float64Array_iyz = new Uint8Array(float64Array_iyz.buffer);
    float64Array_iyz[0] = +this.iyz;
    buff[offset + 0] = uInt8Float64Array_iyz[0];
    buff[offset + 1] = uInt8Float64Array_iyz[1];
    buff[offset + 2] = uInt8Float64Array_iyz[2];
    buff[offset + 3] = uInt8Float64Array_iyz[3];
    buff[offset + 4] = uInt8Float64Array_iyz[4];
    buff[offset + 5] = uInt8Float64Array_iyz[5];
    buff[offset + 6] = uInt8Float64Array_iyz[6];
    buff[offset + 7] = uInt8Float64Array_iyz[7];
    offset += 8;
    var float64Array_izz = new Float64Array(1);
    var uInt8Float64Array_izz = new Uint8Array(float64Array_izz.buffer);
    float64Array_izz[0] = +this.izz;
    buff[offset + 0] = uInt8Float64Array_izz[0];
    buff[offset + 1] = uInt8Float64Array_izz[1];
    buff[offset + 2] = uInt8Float64Array_izz[2];
    buff[offset + 3] = uInt8Float64Array_izz[3];
    buff[offset + 4] = uInt8Float64Array_izz[4];
    buff[offset + 5] = uInt8Float64Array_izz[5];
    buff[offset + 6] = uInt8Float64Array_izz[6];
    buff[offset + 7] = uInt8Float64Array_izz[7];
    offset += 8;
    return offset;
};

SetLinkPropertiesRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_link_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_link_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_link_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_link_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_link_name = new TextDecoder('utf8');
    this.link_name = decoder_link_name.decode(buff.slice(offset, offset + length_link_name));
    offset += length_link_name;
    offset += this.com.deserialize(buff, offset);
    this.gravity_mode = buff[offset] !== 0 ? true : false;
    offset += 1;
    var float64Array_mass = new Float64Array(1);
    var uInt8Float64Array_mass = new Uint8Array(float64Array_mass.buffer);
    uInt8Float64Array_mass[0] = buff[offset + 0];
    uInt8Float64Array_mass[1] = buff[offset + 1];
    uInt8Float64Array_mass[2] = buff[offset + 2];
    uInt8Float64Array_mass[3] = buff[offset + 3];
    uInt8Float64Array_mass[4] = buff[offset + 4];
    uInt8Float64Array_mass[5] = buff[offset + 5];
    uInt8Float64Array_mass[6] = buff[offset + 6];
    uInt8Float64Array_mass[7] = buff[offset + 7];
    this.mass = float64Array_mass[0];
    offset += 8;
    var float64Array_ixx = new Float64Array(1);
    var uInt8Float64Array_ixx = new Uint8Array(float64Array_ixx.buffer);
    uInt8Float64Array_ixx[0] = buff[offset + 0];
    uInt8Float64Array_ixx[1] = buff[offset + 1];
    uInt8Float64Array_ixx[2] = buff[offset + 2];
    uInt8Float64Array_ixx[3] = buff[offset + 3];
    uInt8Float64Array_ixx[4] = buff[offset + 4];
    uInt8Float64Array_ixx[5] = buff[offset + 5];
    uInt8Float64Array_ixx[6] = buff[offset + 6];
    uInt8Float64Array_ixx[7] = buff[offset + 7];
    this.ixx = float64Array_ixx[0];
    offset += 8;
    var float64Array_ixy = new Float64Array(1);
    var uInt8Float64Array_ixy = new Uint8Array(float64Array_ixy.buffer);
    uInt8Float64Array_ixy[0] = buff[offset + 0];
    uInt8Float64Array_ixy[1] = buff[offset + 1];
    uInt8Float64Array_ixy[2] = buff[offset + 2];
    uInt8Float64Array_ixy[3] = buff[offset + 3];
    uInt8Float64Array_ixy[4] = buff[offset + 4];
    uInt8Float64Array_ixy[5] = buff[offset + 5];
    uInt8Float64Array_ixy[6] = buff[offset + 6];
    uInt8Float64Array_ixy[7] = buff[offset + 7];
    this.ixy = float64Array_ixy[0];
    offset += 8;
    var float64Array_ixz = new Float64Array(1);
    var uInt8Float64Array_ixz = new Uint8Array(float64Array_ixz.buffer);
    uInt8Float64Array_ixz[0] = buff[offset + 0];
    uInt8Float64Array_ixz[1] = buff[offset + 1];
    uInt8Float64Array_ixz[2] = buff[offset + 2];
    uInt8Float64Array_ixz[3] = buff[offset + 3];
    uInt8Float64Array_ixz[4] = buff[offset + 4];
    uInt8Float64Array_ixz[5] = buff[offset + 5];
    uInt8Float64Array_ixz[6] = buff[offset + 6];
    uInt8Float64Array_ixz[7] = buff[offset + 7];
    this.ixz = float64Array_ixz[0];
    offset += 8;
    var float64Array_iyy = new Float64Array(1);
    var uInt8Float64Array_iyy = new Uint8Array(float64Array_iyy.buffer);
    uInt8Float64Array_iyy[0] = buff[offset + 0];
    uInt8Float64Array_iyy[1] = buff[offset + 1];
    uInt8Float64Array_iyy[2] = buff[offset + 2];
    uInt8Float64Array_iyy[3] = buff[offset + 3];
    uInt8Float64Array_iyy[4] = buff[offset + 4];
    uInt8Float64Array_iyy[5] = buff[offset + 5];
    uInt8Float64Array_iyy[6] = buff[offset + 6];
    uInt8Float64Array_iyy[7] = buff[offset + 7];
    this.iyy = float64Array_iyy[0];
    offset += 8;
    var float64Array_iyz = new Float64Array(1);
    var uInt8Float64Array_iyz = new Uint8Array(float64Array_iyz.buffer);
    uInt8Float64Array_iyz[0] = buff[offset + 0];
    uInt8Float64Array_iyz[1] = buff[offset + 1];
    uInt8Float64Array_iyz[2] = buff[offset + 2];
    uInt8Float64Array_iyz[3] = buff[offset + 3];
    uInt8Float64Array_iyz[4] = buff[offset + 4];
    uInt8Float64Array_iyz[5] = buff[offset + 5];
    uInt8Float64Array_iyz[6] = buff[offset + 6];
    uInt8Float64Array_iyz[7] = buff[offset + 7];
    this.iyz = float64Array_iyz[0];
    offset += 8;
    var float64Array_izz = new Float64Array(1);
    var uInt8Float64Array_izz = new Uint8Array(float64Array_izz.buffer);
    uInt8Float64Array_izz[0] = buff[offset + 0];
    uInt8Float64Array_izz[1] = buff[offset + 1];
    uInt8Float64Array_izz[2] = buff[offset + 2];
    uInt8Float64Array_izz[3] = buff[offset + 3];
    uInt8Float64Array_izz[4] = buff[offset + 4];
    uInt8Float64Array_izz[5] = buff[offset + 5];
    uInt8Float64Array_izz[6] = buff[offset + 6];
    uInt8Float64Array_izz[7] = buff[offset + 7];
    this.izz = float64Array_izz[0];
    offset += 8;
    return offset;
};

SetLinkPropertiesRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_link_name = new TextEncoder('utf8');
    var utf8array_link_name = encoder_link_name.encode(this.link_name);
    length += 4;
    length += utf8array_link_name.length;
    length += this.com.serializedLength();
    length += 1
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    return length;
};

SetLinkPropertiesRequest.prototype.echo = function() { return ""; };

SetLinkPropertiesRequest.prototype.getType = function() { return "gazebo_msgs/SetLinkProperties"; };

SetLinkPropertiesRequest.prototype.getMD5 = function() { return "03d7e308476601832a9e1ce9d7eab722"; };

SetLinkPropertiesRequest.prototype.getID = function() { return this.__id__; };

SetLinkPropertiesRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetLinkPropertiesRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.SetLinkPropertiesResponse=f();}})(function(){var define,module,exports;'use strict';

function SetLinkPropertiesResponse() {
    this.__id__ = 0;
    this.success = false;
    this.status_message = "";
};

SetLinkPropertiesResponse.prototype.serialize = function(buff, idx) {
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

SetLinkPropertiesResponse.prototype.deserialize = function(buff, idx) {
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

SetLinkPropertiesResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

SetLinkPropertiesResponse.prototype.echo = function() { return ""; };

SetLinkPropertiesResponse.prototype.getType = function() { return "gazebo_msgs/SetLinkProperties"; };

SetLinkPropertiesResponse.prototype.getMD5 = function() { return "777dea05607e1bca37e3206f97801d89"; };

SetLinkPropertiesResponse.prototype.getID = function() { return this.__id__; };

SetLinkPropertiesResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetLinkPropertiesResponse(); };

});
