(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.SaveMapRequest=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/String.js'></script>");

function SaveMapRequest() {
    this.__id__ = 0;
    this.filename = std_msgs.String();
};

SaveMapRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    offset += this.filename.serialize(buff, offset);
    return offset;
};

SaveMapRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    offset += this.filename.deserialize(buff, offset);
    return offset;
};

SaveMapRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += this.filename.serializedLength();
    return length;
};

SaveMapRequest.prototype.echo = function() { return ""; };

SaveMapRequest.prototype.getType = function() { return "map_msgs/SaveMap"; };

SaveMapRequest.prototype.getMD5 = function() { return "6643d8ede81a23998690e6a3ff657316"; };

SaveMapRequest.prototype.getID = function() { return this.__id__; };

SaveMapRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SaveMapRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.SaveMapResponse=f();}})(function(){var define,module,exports;'use strict';

function SaveMapResponse() {
    this.__id__ = 0;
};

SaveMapResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

SaveMapResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

SaveMapResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

SaveMapResponse.prototype.echo = function() { return ""; };

SaveMapResponse.prototype.getType = function() { return "map_msgs/SaveMap"; };

SaveMapResponse.prototype.getMD5 = function() { return "9cd07446fa1bd59b4758dadf19f196e9"; };

SaveMapResponse.prototype.getID = function() { return this.__id__; };

SaveMapResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SaveMapResponse(); };

});
