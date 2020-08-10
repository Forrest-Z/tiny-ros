(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.smach_msgs==="undefined"){g.smach_msgs={};}g.smach_msgs.SmachContainerStatus=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function SmachContainerStatus() {
    this.header = std_msgs.Header();
    this.path = "";
    this.initial_states = new Array();
    this.active_states = new Array();
    this.local_data = "";
    this.info = "";
};

SmachContainerStatus.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var encoder_path = new TextEncoder('utf8');
    var utf8array_path = encoder_path.encode(this.path);
    buff[offset + 0] = (utf8array_path.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_path.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_path.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_path.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_path.length; i++) {
        buff[offset + i] = utf8array_path[i];
    }
    offset += utf8array_path.length;
    var length_initial_states = this.initial_states.length;
    buff[offset + 0] = (length_initial_states >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_initial_states >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_initial_states >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_initial_states >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_initial_states; i++) {
        var encoder_initial_statesi = new TextEncoder('utf8');
        var utf8array_initial_statesi = encoder_initial_statesi.encode(this.initial_states[i]);
        buff[offset + 0] = (utf8array_initial_statesi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_initial_statesi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_initial_statesi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_initial_statesi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_initial_statesi.length; i++) {
            buff[offset + i] = utf8array_initial_statesi[i];
        }
        offset += utf8array_initial_statesi.length;
    }
    var length_active_states = this.active_states.length;
    buff[offset + 0] = (length_active_states >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_active_states >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_active_states >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_active_states >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_active_states; i++) {
        var encoder_active_statesi = new TextEncoder('utf8');
        var utf8array_active_statesi = encoder_active_statesi.encode(this.active_states[i]);
        buff[offset + 0] = (utf8array_active_statesi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_active_statesi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_active_statesi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_active_statesi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_active_statesi.length; i++) {
            buff[offset + i] = utf8array_active_statesi[i];
        }
        offset += utf8array_active_statesi.length;
    }
    var encoder_local_data = new TextEncoder('utf8');
    var utf8array_local_data = encoder_local_data.encode(this.local_data);
    buff[offset + 0] = (utf8array_local_data.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_local_data.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_local_data.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_local_data.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_local_data.length; i++) {
        buff[offset + i] = utf8array_local_data[i];
    }
    offset += utf8array_local_data.length;
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
    return offset;
};

SmachContainerStatus.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_path = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_path |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_path |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_path |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_path = new TextDecoder('utf8');
    this.path = decoder_path.decode(buff.slice(offset, offset + length_path));
    offset += length_path;
    var length_initial_states = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_initial_states |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_initial_states |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_initial_states |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.initial_states = new Array(length_initial_states);
    for (var i = 0; i < length_initial_states; i++) {
        var length_initial_statesi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_initial_statesi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_initial_statesi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_initial_statesi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_initial_statesi = new TextDecoder('utf8');
        this.initial_states[i] = decoder_initial_statesi.decode(buff.slice(offset, offset + length_initial_statesi));
        offset += length_initial_statesi;
    }
    var length_active_states = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_active_states |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_active_states |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_active_states |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.active_states = new Array(length_active_states);
    for (var i = 0; i < length_active_states; i++) {
        var length_active_statesi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_active_statesi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_active_statesi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_active_statesi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_active_statesi = new TextDecoder('utf8');
        this.active_states[i] = decoder_active_statesi.decode(buff.slice(offset, offset + length_active_statesi));
        offset += length_active_statesi;
    }
    var length_local_data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_local_data |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_local_data |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_local_data |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_local_data = new TextDecoder('utf8');
    this.local_data = decoder_local_data.decode(buff.slice(offset, offset + length_local_data));
    offset += length_local_data;
    var length_info = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_info |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_info |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_info |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_info = new TextDecoder('utf8');
    this.info = decoder_info.decode(buff.slice(offset, offset + length_info));
    offset += length_info;
    return offset;
};

SmachContainerStatus.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var encoder_path = new TextEncoder('utf8');
    var utf8array_path = encoder_path.encode(this.path);
    length += 4;
    length += utf8array_path.length;
    var length_initial_states = this.initial_states.length;
    length += 4;
    for (var i = 0; i < length_initial_states; i++) {
        var encoder_initial_statesi = new TextEncoder('utf8');
        var utf8array_initial_statesi = encoder_initial_statesi.encode(this.initial_states[i]);
        length += 4;
        length += utf8array_initial_statesi.length;
    }
    var length_active_states = this.active_states.length;
    length += 4;
    for (var i = 0; i < length_active_states; i++) {
        var encoder_active_statesi = new TextEncoder('utf8');
        var utf8array_active_statesi = encoder_active_statesi.encode(this.active_states[i]);
        length += 4;
        length += utf8array_active_statesi.length;
    }
    var encoder_local_data = new TextEncoder('utf8');
    var utf8array_local_data = encoder_local_data.encode(this.local_data);
    length += 4;
    length += utf8array_local_data.length;
    var encoder_info = new TextEncoder('utf8');
    var utf8array_info = encoder_info.encode(this.info);
    length += 4;
    length += utf8array_info.length;
    return length;
};

SmachContainerStatus.prototype.echo = function() { return ""; };

SmachContainerStatus.prototype.getType = function() { return "smach_msgs/SmachContainerStatus"; };

SmachContainerStatus.prototype.getMD5 = function() { return "3e74cf72da18311be27e7a5970ea6415"; };

SmachContainerStatus.prototype.getID = function() { return 0; };

SmachContainerStatus.prototype.setID = function(id) { };

return function() { return new SmachContainerStatus(); };

});
