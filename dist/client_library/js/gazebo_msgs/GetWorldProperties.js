(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.GetWorldPropertiesRequest=f();}})(function(){var define,module,exports;'use strict';

function GetWorldPropertiesRequest() {
    this.__id__ = 0;
};

GetWorldPropertiesRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

GetWorldPropertiesRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

GetWorldPropertiesRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

GetWorldPropertiesRequest.prototype.echo = function() { return ""; };

GetWorldPropertiesRequest.prototype.getType = function() { return "gazebo_msgs/GetWorldProperties"; };

GetWorldPropertiesRequest.prototype.getMD5 = function() { return "3aa5de7106eec5dae41ad1c9ae681123"; };

GetWorldPropertiesRequest.prototype.getID = function() { return this.__id__; };

GetWorldPropertiesRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetWorldPropertiesRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.GetWorldPropertiesResponse=f();}})(function(){var define,module,exports;'use strict';

function GetWorldPropertiesResponse() {
    this.__id__ = 0;
    this.sim_time = 0.0;
    this.model_names = new Array();
    this.rendering_enabled = false;
    this.success = false;
    this.status_message = "";
};

GetWorldPropertiesResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var float64Array_sim_time = new Float64Array(1);
    var uInt8Float64Array_sim_time = new Uint8Array(float64Array_sim_time.buffer);
    float64Array_sim_time[0] = +this.sim_time;
    buff[offset + 0] = uInt8Float64Array_sim_time[0];
    buff[offset + 1] = uInt8Float64Array_sim_time[1];
    buff[offset + 2] = uInt8Float64Array_sim_time[2];
    buff[offset + 3] = uInt8Float64Array_sim_time[3];
    buff[offset + 4] = uInt8Float64Array_sim_time[4];
    buff[offset + 5] = uInt8Float64Array_sim_time[5];
    buff[offset + 6] = uInt8Float64Array_sim_time[6];
    buff[offset + 7] = uInt8Float64Array_sim_time[7];
    offset += 8;
    var length_model_names = this.model_names.length;
    buff[offset + 0] = (length_model_names >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_model_names >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_model_names >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_model_names >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_model_names; i++) {
        var encoder_model_namesi = new TextEncoder('utf8');
        var utf8array_model_namesi = encoder_model_namesi.encode(this.model_names[i]);
        buff[offset + 0] = (utf8array_model_namesi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_model_namesi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_model_namesi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_model_namesi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_model_namesi.length; i++) {
            buff[offset + i] = utf8array_model_namesi[i];
        }
        offset += utf8array_model_namesi.length;
    }
    buff[offset] = this.rendering_enabled === false ? 0 : 1;
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

GetWorldPropertiesResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var float64Array_sim_time = new Float64Array(1);
    var uInt8Float64Array_sim_time = new Uint8Array(float64Array_sim_time.buffer);
    uInt8Float64Array_sim_time[0] = buff[offset + 0];
    uInt8Float64Array_sim_time[1] = buff[offset + 1];
    uInt8Float64Array_sim_time[2] = buff[offset + 2];
    uInt8Float64Array_sim_time[3] = buff[offset + 3];
    uInt8Float64Array_sim_time[4] = buff[offset + 4];
    uInt8Float64Array_sim_time[5] = buff[offset + 5];
    uInt8Float64Array_sim_time[6] = buff[offset + 6];
    uInt8Float64Array_sim_time[7] = buff[offset + 7];
    this.sim_time = float64Array_sim_time[0];
    offset += 8;
    var length_model_names = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_model_names |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_model_names |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_model_names |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.model_names = new Array(length_model_names);
    for (var i = 0; i < length_model_names; i++) {
        var length_model_namesi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_model_namesi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_model_namesi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_model_namesi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_model_namesi = new TextDecoder('utf8');
        this.model_names[i] = decoder_model_namesi.decode(buff.slice(offset, offset + length_model_namesi));
        offset += length_model_namesi;
    }
    this.rendering_enabled = buff[offset] !== 0 ? true : false;
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

GetWorldPropertiesResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 8
    var length_model_names = this.model_names.length;
    length += 4;
    for (var i = 0; i < length_model_names; i++) {
        var encoder_model_namesi = new TextEncoder('utf8');
        var utf8array_model_namesi = encoder_model_namesi.encode(this.model_names[i]);
        length += 4;
        length += utf8array_model_namesi.length;
    }
    length += 1
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

GetWorldPropertiesResponse.prototype.echo = function() { return ""; };

GetWorldPropertiesResponse.prototype.getType = function() { return "gazebo_msgs/GetWorldProperties"; };

GetWorldPropertiesResponse.prototype.getMD5 = function() { return "fe944c1c210919291ad14bc43b6c10cf"; };

GetWorldPropertiesResponse.prototype.getID = function() { return this.__id__; };

GetWorldPropertiesResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetWorldPropertiesResponse(); };

});
