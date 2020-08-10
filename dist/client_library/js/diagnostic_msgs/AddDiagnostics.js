(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.diagnostic_msgs==="undefined"){g.diagnostic_msgs={};}g.diagnostic_msgs.AddDiagnosticsRequest=f();}})(function(){var define,module,exports;'use strict';

function AddDiagnosticsRequest() {
    this.__id__ = 0;
    this.load_namespace = "";
};

AddDiagnosticsRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_load_namespace = new TextEncoder('utf8');
    var utf8array_load_namespace = encoder_load_namespace.encode(this.load_namespace);
    buff[offset + 0] = (utf8array_load_namespace.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_load_namespace.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_load_namespace.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_load_namespace.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_load_namespace.length; i++) {
        buff[offset + i] = utf8array_load_namespace[i];
    }
    offset += utf8array_load_namespace.length;
    return offset;
};

AddDiagnosticsRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_load_namespace = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_load_namespace |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_load_namespace |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_load_namespace |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_load_namespace = new TextDecoder('utf8');
    this.load_namespace = decoder_load_namespace.decode(buff.slice(offset, offset + length_load_namespace));
    offset += length_load_namespace;
    return offset;
};

AddDiagnosticsRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_load_namespace = new TextEncoder('utf8');
    var utf8array_load_namespace = encoder_load_namespace.encode(this.load_namespace);
    length += 4;
    length += utf8array_load_namespace.length;
    return length;
};

AddDiagnosticsRequest.prototype.echo = function() { return ""; };

AddDiagnosticsRequest.prototype.getType = function() { return "diagnostic_msgs/AddDiagnostics"; };

AddDiagnosticsRequest.prototype.getMD5 = function() { return "005ba76b3cd04edebfe46acad928fbeb"; };

AddDiagnosticsRequest.prototype.getID = function() { return this.__id__; };

AddDiagnosticsRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new AddDiagnosticsRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.diagnostic_msgs==="undefined"){g.diagnostic_msgs={};}g.diagnostic_msgs.AddDiagnosticsResponse=f();}})(function(){var define,module,exports;'use strict';

function AddDiagnosticsResponse() {
    this.__id__ = 0;
    this.success = false;
    this.message = "";
};

AddDiagnosticsResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset] = this.success === false ? 0 : 1;
    offset += 1;
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
    return offset;
};

AddDiagnosticsResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.success = buff[offset] !== 0 ? true : false;
    offset += 1;
    var length_message = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_message |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_message |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_message |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_message = new TextDecoder('utf8');
    this.message = decoder_message.decode(buff.slice(offset, offset + length_message));
    offset += length_message;
    return offset;
};

AddDiagnosticsResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    var encoder_message = new TextEncoder('utf8');
    var utf8array_message = encoder_message.encode(this.message);
    length += 4;
    length += utf8array_message.length;
    return length;
};

AddDiagnosticsResponse.prototype.echo = function() { return ""; };

AddDiagnosticsResponse.prototype.getType = function() { return "diagnostic_msgs/AddDiagnostics"; };

AddDiagnosticsResponse.prototype.getMD5 = function() { return "9bd37b30a2340a31743d1e80a2c52ed0"; };

AddDiagnosticsResponse.prototype.getID = function() { return this.__id__; };

AddDiagnosticsResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new AddDiagnosticsResponse(); };

});
