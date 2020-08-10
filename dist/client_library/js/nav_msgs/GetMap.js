(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GetMapRequest=f();}})(function(){var define,module,exports;'use strict';

function GetMapRequest() {
    this.__id__ = 0;
};

GetMapRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

GetMapRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

GetMapRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

GetMapRequest.prototype.echo = function() { return ""; };

GetMapRequest.prototype.getType = function() { return "nav_msgs/GetMap"; };

GetMapRequest.prototype.getMD5 = function() { return "8ea512c109a0b3a9eca8de407dd02d2a"; };

GetMapRequest.prototype.getID = function() { return this.__id__; };

GetMapRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetMapRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GetMapResponse=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/OccupancyGrid.js'></script>");

function GetMapResponse() {
    this.__id__ = 0;
    this.map = nav_msgs.OccupancyGrid();
};

GetMapResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    offset += this.map.serialize(buff, offset);
    return offset;
};

GetMapResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    offset += this.map.deserialize(buff, offset);
    return offset;
};

GetMapResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += this.map.serializedLength();
    return length;
};

GetMapResponse.prototype.echo = function() { return ""; };

GetMapResponse.prototype.getType = function() { return "nav_msgs/GetMap"; };

GetMapResponse.prototype.getMD5 = function() { return "5bf895a6cdaff312c69ca1cef20fed64"; };

GetMapResponse.prototype.getID = function() { return this.__id__; };

GetMapResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetMapResponse(); };

});
