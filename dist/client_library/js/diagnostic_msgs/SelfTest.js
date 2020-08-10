(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.diagnostic_msgs==="undefined"){g.diagnostic_msgs={};}g.diagnostic_msgs.SelfTestRequest=f();}})(function(){var define,module,exports;'use strict';

function SelfTestRequest() {
    this.__id__ = 0;
};

SelfTestRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

SelfTestRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

SelfTestRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

SelfTestRequest.prototype.echo = function() { return ""; };

SelfTestRequest.prototype.getType = function() { return "diagnostic_msgs/SelfTest"; };

SelfTestRequest.prototype.getMD5 = function() { return "049f87742408b36b8ef5f7dd71e3ef5a"; };

SelfTestRequest.prototype.getID = function() { return this.__id__; };

SelfTestRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SelfTestRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.diagnostic_msgs==="undefined"){g.diagnostic_msgs={};}g.diagnostic_msgs.SelfTestResponse=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/diagnostic_msgs/DiagnosticStatus.js'></script>");

function SelfTestResponse() {
    this.__id__ = 0;
    this.id = "";
    this.passed = 0;
    this.status = new Array();
};

SelfTestResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_id = new TextEncoder('utf8');
    var utf8array_id = encoder_id.encode(this.id);
    buff[offset + 0] = (utf8array_id.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_id.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_id.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_id.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_id.length; i++) {
        buff[offset + i] = utf8array_id[i];
    }
    offset += utf8array_id.length;
    buff[offset + 0] = ((+this.passed) >> (8 * 0)) & 0xFF;
    offset += 1;
    var length_status = this.status.length;
    buff[offset + 0] = (length_status >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_status >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_status >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_status >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_status; i++) {
        offset += this.status[i].serialize(buff, offset);
    }
    return offset;
};

SelfTestResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_id = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_id |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_id |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_id |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_id = new TextDecoder('utf8');
    this.id = decoder_id.decode(buff.slice(offset, offset + length_id));
    offset += length_id;
    this.passed = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    var length_status = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_status |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_status |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_status |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.status = new Array(length_status);
    for (var i = 0; i < length_status; i++) {
        this.status[i] = diagnostic_msgs.DiagnosticStatus();
        offset += this.status[i].deserialize(buff, offset);
    }
    return offset;
};

SelfTestResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_id = new TextEncoder('utf8');
    var utf8array_id = encoder_id.encode(this.id);
    length += 4;
    length += utf8array_id.length;
    length += 1
    var length_status = this.status.length;
    length += 4;
    for (var i = 0; i < length_status; i++) {
        length += this.status[i].serializedLength();
    }
    return length;
};

SelfTestResponse.prototype.echo = function() { return ""; };

SelfTestResponse.prototype.getType = function() { return "diagnostic_msgs/SelfTest"; };

SelfTestResponse.prototype.getMD5 = function() { return "70aaf2a851ccb5e946b2d112ea26f7b9"; };

SelfTestResponse.prototype.getID = function() { return this.__id__; };

SelfTestResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SelfTestResponse(); };

});
