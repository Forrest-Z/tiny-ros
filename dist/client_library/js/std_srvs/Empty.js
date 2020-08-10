(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_srvs==="undefined"){g.std_srvs={};}g.std_srvs.EmptyRequest=f();}})(function(){var define,module,exports;'use strict';

function EmptyRequest() {
    this.__id__ = 0;
};

EmptyRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

EmptyRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

EmptyRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

EmptyRequest.prototype.echo = function() { return ""; };

EmptyRequest.prototype.getType = function() { return "std_srvs/Empty"; };

EmptyRequest.prototype.getMD5 = function() { return "6a9da448a5a2256d30e815f50a75bc57"; };

EmptyRequest.prototype.getID = function() { return this.__id__; };

EmptyRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new EmptyRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_srvs==="undefined"){g.std_srvs={};}g.std_srvs.EmptyResponse=f();}})(function(){var define,module,exports;'use strict';

function EmptyResponse() {
    this.__id__ = 0;
};

EmptyResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

EmptyResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

EmptyResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

EmptyResponse.prototype.echo = function() { return ""; };

EmptyResponse.prototype.getType = function() { return "std_srvs/Empty"; };

EmptyResponse.prototype.getMD5 = function() { return "1b9fedd0c70a6b7846f790471f388d15"; };

EmptyResponse.prototype.getID = function() { return this.__id__; };

EmptyResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new EmptyResponse(); };

});
