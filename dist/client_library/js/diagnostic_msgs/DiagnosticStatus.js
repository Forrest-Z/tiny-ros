(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.diagnostic_msgs==="undefined"){g.diagnostic_msgs={};}g.diagnostic_msgs.DiagnosticStatus=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/diagnostic_msgs/KeyValue.js'></script>");

function DiagnosticStatus() {
    this.level = 0;
    this.name = "";
    this.message = "";
    this.hardware_id = "";
    this.values = new Array();

    // ENUM{
    this.OK = 0;
    this.WARN = 1;
    this.ERROR = 2;
    this.STALE = 3;
    // }ENUM
};

DiagnosticStatus.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.level) >> (8 * 0)) & 0xFF;
    offset += 1;
    var encoder_name = new TextEncoder('utf8');
    var utf8array_name = encoder_name.encode(this.name);
    buff[offset + 0] = (utf8array_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_name.length; i++) {
        buff[offset + i] = utf8array_name[i];
    }
    offset += utf8array_name.length;
    var encoder_message = new TextEncoder('utf8');
    var utf8array_message = encoder_message.encode(this.message);
    buff[offset + 0] = (utf8array_message.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_message.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_message.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_message.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_message.length; i++) {
        buff[offset + i] = utf8array_message[i];
    }
    offset += utf8array_message.length;
    var encoder_hardware_id = new TextEncoder('utf8');
    var utf8array_hardware_id = encoder_hardware_id.encode(this.hardware_id);
    buff[offset + 0] = (utf8array_hardware_id.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_hardware_id.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_hardware_id.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_hardware_id.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_hardware_id.length; i++) {
        buff[offset + i] = utf8array_hardware_id[i];
    }
    offset += utf8array_hardware_id.length;
    var length_values = this.values.length;
    buff[offset + 0] = (length_values >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_values >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_values >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_values >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_values; i++) {
        offset += this.values[i].serialize(buff, offset);
    }
    return offset;
};

DiagnosticStatus.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.level = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    var length_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_name = new TextDecoder('utf8');
    this.name = decoder_name.decode(buff.slice(offset, offset + length_name));
    offset += length_name;
    var length_message = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_message |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_message |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_message |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_message = new TextDecoder('utf8');
    this.message = decoder_message.decode(buff.slice(offset, offset + length_message));
    offset += length_message;
    var length_hardware_id = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_hardware_id |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_hardware_id |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_hardware_id |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_hardware_id = new TextDecoder('utf8');
    this.hardware_id = decoder_hardware_id.decode(buff.slice(offset, offset + length_hardware_id));
    offset += length_hardware_id;
    var length_values = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_values |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_values |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_values |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.values = new Array(length_values);
    for (var i = 0; i < length_values; i++) {
        this.values[i] = diagnostic_msgs.KeyValue();
        offset += this.values[i].deserialize(buff, offset);
    }
    return offset;
};

DiagnosticStatus.prototype.serializedLength = function() {
    var length = 0;
    length += 1
    var encoder_name = new TextEncoder('utf8');
    var utf8array_name = encoder_name.encode(this.name);
    length += 4;
    length += utf8array_name.length;
    var encoder_message = new TextEncoder('utf8');
    var utf8array_message = encoder_message.encode(this.message);
    length += 4;
    length += utf8array_message.length;
    var encoder_hardware_id = new TextEncoder('utf8');
    var utf8array_hardware_id = encoder_hardware_id.encode(this.hardware_id);
    length += 4;
    length += utf8array_hardware_id.length;
    var length_values = this.values.length;
    length += 4;
    for (var i = 0; i < length_values; i++) {
        length += this.values[i].serializedLength();
    }
    return length;
};

DiagnosticStatus.prototype.echo = function() { return ""; };

DiagnosticStatus.prototype.getType = function() { return "diagnostic_msgs/DiagnosticStatus"; };

DiagnosticStatus.prototype.getMD5 = function() { return "9ec892d2145f478061efd60bb1762361"; };

DiagnosticStatus.prototype.getID = function() { return 0; };

DiagnosticStatus.prototype.setID = function(id) { };

return function() { return new DiagnosticStatus(); };

});
