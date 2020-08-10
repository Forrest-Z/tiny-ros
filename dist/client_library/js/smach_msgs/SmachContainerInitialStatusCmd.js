(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.smach_msgs==="undefined"){g.smach_msgs={};}g.smach_msgs.SmachContainerInitialStatusCmd=f();}})(function(){var define,module,exports;'use strict';

function SmachContainerInitialStatusCmd() {
    this.path = "";
    this.initial_states = new Array();
    this.local_data = "";
};

SmachContainerInitialStatusCmd.prototype.serialize = function(buff, idx) {
    var offset = idx;
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
    return offset;
};

SmachContainerInitialStatusCmd.prototype.deserialize = function(buff, idx) {
    var offset = idx;
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
    var length_local_data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_local_data |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_local_data |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_local_data |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_local_data = new TextDecoder('utf8');
    this.local_data = decoder_local_data.decode(buff.slice(offset, offset + length_local_data));
    offset += length_local_data;
    return offset;
};

SmachContainerInitialStatusCmd.prototype.serializedLength = function() {
    var length = 0;
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
    var encoder_local_data = new TextEncoder('utf8');
    var utf8array_local_data = encoder_local_data.encode(this.local_data);
    length += 4;
    length += utf8array_local_data.length;
    return length;
};

SmachContainerInitialStatusCmd.prototype.echo = function() { return ""; };

SmachContainerInitialStatusCmd.prototype.getType = function() { return "smach_msgs/SmachContainerInitialStatusCmd"; };

SmachContainerInitialStatusCmd.prototype.getMD5 = function() { return "b7c8a2bbd4d7c89f80561645ea1f4f13"; };

SmachContainerInitialStatusCmd.prototype.getID = function() { return 0; };

SmachContainerInitialStatusCmd.prototype.setID = function(id) { };

return function() { return new SmachContainerInitialStatusCmd(); };

});
