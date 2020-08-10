(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.GetLinkStateRequest=f();}})(function(){var define,module,exports;'use strict';

function GetLinkStateRequest() {
    this.__id__ = 0;
    this.link_name = "";
    this.reference_frame = "";
};

GetLinkStateRequest.prototype.serialize = function(buff, idx) {
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
    var encoder_reference_frame = new TextEncoder('utf8');
    var utf8array_reference_frame = encoder_reference_frame.encode(this.reference_frame);
    buff[offset + 0] = (utf8array_reference_frame.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_reference_frame.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_reference_frame.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_reference_frame.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_reference_frame.length; i++) {
        buff[offset + i] = utf8array_reference_frame[i];
    }
    offset += utf8array_reference_frame.length;
    return offset;
};

GetLinkStateRequest.prototype.deserialize = function(buff, idx) {
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
    var length_reference_frame = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_reference_frame |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_reference_frame |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_reference_frame |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_reference_frame = new TextDecoder('utf8');
    this.reference_frame = decoder_reference_frame.decode(buff.slice(offset, offset + length_reference_frame));
    offset += length_reference_frame;
    return offset;
};

GetLinkStateRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_link_name = new TextEncoder('utf8');
    var utf8array_link_name = encoder_link_name.encode(this.link_name);
    length += 4;
    length += utf8array_link_name.length;
    var encoder_reference_frame = new TextEncoder('utf8');
    var utf8array_reference_frame = encoder_reference_frame.encode(this.reference_frame);
    length += 4;
    length += utf8array_reference_frame.length;
    return length;
};

GetLinkStateRequest.prototype.echo = function() { return ""; };

GetLinkStateRequest.prototype.getType = function() { return "gazebo_msgs/GetLinkState"; };

GetLinkStateRequest.prototype.getMD5 = function() { return "b9de4ed1795bda93c873763a2e55e76b"; };

GetLinkStateRequest.prototype.getID = function() { return this.__id__; };

GetLinkStateRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetLinkStateRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.GetLinkStateResponse=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/gazebo_msgs/LinkState.js'></script>");

function GetLinkStateResponse() {
    this.__id__ = 0;
    this.link_state = gazebo_msgs.LinkState();
    this.success = false;
    this.status_message = "";
};

GetLinkStateResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    offset += this.link_state.serialize(buff, offset);
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

GetLinkStateResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    offset += this.link_state.deserialize(buff, offset);
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

GetLinkStateResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += this.link_state.serializedLength();
    length += 1
    var encoder_status_message = new TextEncoder('utf8');
    var utf8array_status_message = encoder_status_message.encode(this.status_message);
    length += 4;
    length += utf8array_status_message.length;
    return length;
};

GetLinkStateResponse.prototype.echo = function() { return ""; };

GetLinkStateResponse.prototype.getType = function() { return "gazebo_msgs/GetLinkState"; };

GetLinkStateResponse.prototype.getMD5 = function() { return "4d4305d53d97f8edc3b3ce04bcb94ed0"; };

GetLinkStateResponse.prototype.getID = function() { return this.__id__; };

GetLinkStateResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetLinkStateResponse(); };

});
