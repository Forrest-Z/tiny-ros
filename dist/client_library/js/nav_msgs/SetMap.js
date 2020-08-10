(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.SetMapRequest=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/OccupancyGrid.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/PoseWithCovarianceStamped.js'></script>");

function SetMapRequest() {
    this.__id__ = 0;
    this.map = nav_msgs.OccupancyGrid();
    this.initial_pose = geometry_msgs.PoseWithCovarianceStamped();
};

SetMapRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    offset += this.map.serialize(buff, offset);
    offset += this.initial_pose.serialize(buff, offset);
    return offset;
};

SetMapRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    offset += this.map.deserialize(buff, offset);
    offset += this.initial_pose.deserialize(buff, offset);
    return offset;
};

SetMapRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += this.map.serializedLength();
    length += this.initial_pose.serializedLength();
    return length;
};

SetMapRequest.prototype.echo = function() { return ""; };

SetMapRequest.prototype.getType = function() { return "nav_msgs/SetMap"; };

SetMapRequest.prototype.getMD5 = function() { return "946e1bd68c9db117a530a571e33d9e49"; };

SetMapRequest.prototype.getID = function() { return this.__id__; };

SetMapRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetMapRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.SetMapResponse=f();}})(function(){var define,module,exports;'use strict';

function SetMapResponse() {
    this.__id__ = 0;
    this.success = false;
};

SetMapResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset] = this.success === false ? 0 : 1;
    offset += 1;
    return offset;
};

SetMapResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.success = buff[offset] !== 0 ? true : false;
    offset += 1;
    return offset;
};

SetMapResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 1
    return length;
};

SetMapResponse.prototype.echo = function() { return ""; };

SetMapResponse.prototype.getType = function() { return "nav_msgs/SetMap"; };

SetMapResponse.prototype.getMD5 = function() { return "1e32607e79013262dafbbac9044e9cda"; };

SetMapResponse.prototype.getID = function() { return this.__id__; };

SetMapResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetMapResponse(); };

});
