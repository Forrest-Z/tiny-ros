(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.JointState=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function JointState() {
    this.header = std_msgs.Header();
    this.name = new Array();
    this.position = new Array();
    this.velocity = new Array();
    this.effort = new Array();
};

JointState.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var length_name = this.name.length;
    buff[offset + 0] = (length_name >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_name >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_name >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_name >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_name; i++) {
        var encoder_namei = new TextEncoder('utf8');
        var utf8array_namei = encoder_namei.encode(this.name[i]);
        buff[offset + 0] = (utf8array_namei.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_namei.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_namei.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_namei.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_namei.length; i++) {
            buff[offset + i] = utf8array_namei[i];
        }
        offset += utf8array_namei.length;
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
    var length_velocity = this.velocity.length;
    buff[offset + 0] = (length_velocity >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_velocity >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_velocity >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_velocity >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_velocity; i++) {
        var float64Array_velocityi = new Float64Array(1);
        var uInt8Float64Array_velocityi = new Uint8Array(float64Array_velocityi.buffer);
        float64Array_velocityi[0] = +this.velocity[i];
        buff[offset + 0] = uInt8Float64Array_velocityi[0];
        buff[offset + 1] = uInt8Float64Array_velocityi[1];
        buff[offset + 2] = uInt8Float64Array_velocityi[2];
        buff[offset + 3] = uInt8Float64Array_velocityi[3];
        buff[offset + 4] = uInt8Float64Array_velocityi[4];
        buff[offset + 5] = uInt8Float64Array_velocityi[5];
        buff[offset + 6] = uInt8Float64Array_velocityi[6];
        buff[offset + 7] = uInt8Float64Array_velocityi[7];
        offset += 8;
    }
    var length_effort = this.effort.length;
    buff[offset + 0] = (length_effort >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_effort >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_effort >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_effort >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_effort; i++) {
        var float64Array_efforti = new Float64Array(1);
        var uInt8Float64Array_efforti = new Uint8Array(float64Array_efforti.buffer);
        float64Array_efforti[0] = +this.effort[i];
        buff[offset + 0] = uInt8Float64Array_efforti[0];
        buff[offset + 1] = uInt8Float64Array_efforti[1];
        buff[offset + 2] = uInt8Float64Array_efforti[2];
        buff[offset + 3] = uInt8Float64Array_efforti[3];
        buff[offset + 4] = uInt8Float64Array_efforti[4];
        buff[offset + 5] = uInt8Float64Array_efforti[5];
        buff[offset + 6] = uInt8Float64Array_efforti[6];
        buff[offset + 7] = uInt8Float64Array_efforti[7];
        offset += 8;
    }
    return offset;
};

JointState.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_name |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.name = new Array(length_name);
    for (var i = 0; i < length_name; i++) {
        var length_namei = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_namei |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_namei |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_namei |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_namei = new TextDecoder('utf8');
        this.name[i] = decoder_namei.decode(buff.slice(offset, offset + length_namei));
        offset += length_namei;
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
    var length_velocity = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_velocity |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_velocity |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_velocity |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.velocity = new Array(length_velocity);
    for (var i = 0; i < length_velocity; i++) {
        var float64Array_velocityi = new Float64Array(1);
        var uInt8Float64Array_velocityi = new Uint8Array(float64Array_velocityi.buffer);
        uInt8Float64Array_velocityi[0] = buff[offset + 0];
        uInt8Float64Array_velocityi[1] = buff[offset + 1];
        uInt8Float64Array_velocityi[2] = buff[offset + 2];
        uInt8Float64Array_velocityi[3] = buff[offset + 3];
        uInt8Float64Array_velocityi[4] = buff[offset + 4];
        uInt8Float64Array_velocityi[5] = buff[offset + 5];
        uInt8Float64Array_velocityi[6] = buff[offset + 6];
        uInt8Float64Array_velocityi[7] = buff[offset + 7];
        this.velocity[i] = float64Array_velocityi[0];
        offset += 8;
    }
    var length_effort = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_effort |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_effort |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_effort |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.effort = new Array(length_effort);
    for (var i = 0; i < length_effort; i++) {
        var float64Array_efforti = new Float64Array(1);
        var uInt8Float64Array_efforti = new Uint8Array(float64Array_efforti.buffer);
        uInt8Float64Array_efforti[0] = buff[offset + 0];
        uInt8Float64Array_efforti[1] = buff[offset + 1];
        uInt8Float64Array_efforti[2] = buff[offset + 2];
        uInt8Float64Array_efforti[3] = buff[offset + 3];
        uInt8Float64Array_efforti[4] = buff[offset + 4];
        uInt8Float64Array_efforti[5] = buff[offset + 5];
        uInt8Float64Array_efforti[6] = buff[offset + 6];
        uInt8Float64Array_efforti[7] = buff[offset + 7];
        this.effort[i] = float64Array_efforti[0];
        offset += 8;
    }
    return offset;
};

JointState.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var length_name = this.name.length;
    length += 4;
    for (var i = 0; i < length_name; i++) {
        var encoder_namei = new TextEncoder('utf8');
        var utf8array_namei = encoder_namei.encode(this.name[i]);
        length += 4;
        length += utf8array_namei.length;
    }
    var length_position = this.position.length;
    length += 4;
    for (var i = 0; i < length_position; i++) {
        length += 8
    }
    var length_velocity = this.velocity.length;
    length += 4;
    for (var i = 0; i < length_velocity; i++) {
        length += 8
    }
    var length_effort = this.effort.length;
    length += 4;
    for (var i = 0; i < length_effort; i++) {
        length += 8
    }
    return length;
};

JointState.prototype.echo = function() { return ""; };

JointState.prototype.getType = function() { return "sensor_msgs/JointState"; };

JointState.prototype.getMD5 = function() { return "6df7130a6d6a4c2f2037ce4a6e061fb9"; };

JointState.prototype.getID = function() { return 0; };

JointState.prototype.setID = function(id) { };

return function() { return new JointState(); };

});
