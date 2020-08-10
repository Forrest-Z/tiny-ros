(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.GetPointMapRequest=f();}})(function(){var define,module,exports;'use strict';

function GetPointMapRequest() {
    this.__id__ = 0;
};

GetPointMapRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

GetPointMapRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

GetPointMapRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

GetPointMapRequest.prototype.echo = function() { return ""; };

GetPointMapRequest.prototype.getType = function() { return "map_msgs/GetPointMap"; };

GetPointMapRequest.prototype.getMD5 = function() { return "418d4ee8c9d758b7ef1aab3e068c7568"; };

GetPointMapRequest.prototype.getID = function() { return this.__id__; };

GetPointMapRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetPointMapRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.GetPointMapResponse=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/sensor_msgs/PointCloud2.js'></script>");

function GetPointMapResponse() {
    this.__id__ = 0;
    this.map = sensor_msgs.PointCloud2();
};

GetPointMapResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    offset += this.map.serialize(buff, offset);
    return offset;
};

GetPointMapResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    offset += this.map.deserialize(buff, offset);
    return offset;
};

GetPointMapResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += this.map.serializedLength();
    return length;
};

GetPointMapResponse.prototype.echo = function() { return ""; };

GetPointMapResponse.prototype.getType = function() { return "map_msgs/GetPointMap"; };

GetPointMapResponse.prototype.getMD5 = function() { return "abc97e6c5b25f536248da280bdf271d7"; };

GetPointMapResponse.prototype.getID = function() { return this.__id__; };

GetPointMapResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetPointMapResponse(); };

});
