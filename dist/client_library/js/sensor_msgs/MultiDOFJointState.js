(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.MultiDOFJointState=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Transform.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Twist.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Wrench.js'></script>");

function MultiDOFJointState() {
    this.header = std_msgs.Header();
    this.joint_names = new Array();
    this.transforms = new Array();
    this.twist = new Array();
    this.wrench = new Array();
};

MultiDOFJointState.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
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
    var length_transforms = this.transforms.length;
    buff[offset + 0] = (length_transforms >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_transforms >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_transforms >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_transforms >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_transforms; i++) {
        offset += this.transforms[i].serialize(buff, offset);
    }
    var length_twist = this.twist.length;
    buff[offset + 0] = (length_twist >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_twist >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_twist >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_twist >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_twist; i++) {
        offset += this.twist[i].serialize(buff, offset);
    }
    var length_wrench = this.wrench.length;
    buff[offset + 0] = (length_wrench >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_wrench >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_wrench >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_wrench >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_wrench; i++) {
        offset += this.wrench[i].serialize(buff, offset);
    }
    return offset;
};

MultiDOFJointState.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
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
    var length_transforms = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_transforms |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_transforms |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_transforms |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.transforms = new Array(length_transforms);
    for (var i = 0; i < length_transforms; i++) {
        this.transforms[i] = geometry_msgs.Transform();
        offset += this.transforms[i].deserialize(buff, offset);
    }
    var length_twist = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_twist |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_twist |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_twist |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.twist = new Array(length_twist);
    for (var i = 0; i < length_twist; i++) {
        this.twist[i] = geometry_msgs.Twist();
        offset += this.twist[i].deserialize(buff, offset);
    }
    var length_wrench = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_wrench |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_wrench |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_wrench |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.wrench = new Array(length_wrench);
    for (var i = 0; i < length_wrench; i++) {
        this.wrench[i] = geometry_msgs.Wrench();
        offset += this.wrench[i].deserialize(buff, offset);
    }
    return offset;
};

MultiDOFJointState.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var length_joint_names = this.joint_names.length;
    length += 4;
    for (var i = 0; i < length_joint_names; i++) {
        var encoder_joint_namesi = new TextEncoder('utf8');
        var utf8array_joint_namesi = encoder_joint_namesi.encode(this.joint_names[i]);
        length += 4;
        length += utf8array_joint_namesi.length;
    }
    var length_transforms = this.transforms.length;
    length += 4;
    for (var i = 0; i < length_transforms; i++) {
        length += this.transforms[i].serializedLength();
    }
    var length_twist = this.twist.length;
    length += 4;
    for (var i = 0; i < length_twist; i++) {
        length += this.twist[i].serializedLength();
    }
    var length_wrench = this.wrench.length;
    length += 4;
    for (var i = 0; i < length_wrench; i++) {
        length += this.wrench[i].serializedLength();
    }
    return length;
};

MultiDOFJointState.prototype.echo = function() { return ""; };

MultiDOFJointState.prototype.getType = function() { return "sensor_msgs/MultiDOFJointState"; };

MultiDOFJointState.prototype.getMD5 = function() { return "c1b0232d8e5071b24db76eb143f286eb"; };

MultiDOFJointState.prototype.getID = function() { return 0; };

MultiDOFJointState.prototype.setID = function(id) { };

return function() { return new MultiDOFJointState(); };

});
