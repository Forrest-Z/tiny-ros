(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.ProjectedMapsInfoRequest=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/map_msgs/ProjectedMapInfo.js'></script>");

function ProjectedMapsInfoRequest() {
    this.__id__ = 0;
    this.projected_maps_info = new Array();
};

ProjectedMapsInfoRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var length_projected_maps_info = this.projected_maps_info.length;
    buff[offset + 0] = (length_projected_maps_info >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_projected_maps_info >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_projected_maps_info >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_projected_maps_info >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_projected_maps_info; i++) {
        offset += this.projected_maps_info[i].serialize(buff, offset);
    }
    return offset;
};

ProjectedMapsInfoRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_projected_maps_info = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_projected_maps_info |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_projected_maps_info |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_projected_maps_info |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.projected_maps_info = new Array(length_projected_maps_info);
    for (var i = 0; i < length_projected_maps_info; i++) {
        this.projected_maps_info[i] = map_msgs.ProjectedMapInfo();
        offset += this.projected_maps_info[i].deserialize(buff, offset);
    }
    return offset;
};

ProjectedMapsInfoRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var length_projected_maps_info = this.projected_maps_info.length;
    length += 4;
    for (var i = 0; i < length_projected_maps_info; i++) {
        length += this.projected_maps_info[i].serializedLength();
    }
    return length;
};

ProjectedMapsInfoRequest.prototype.echo = function() { return ""; };

ProjectedMapsInfoRequest.prototype.getType = function() { return "map_msgs/ProjectedMapsInfo"; };

ProjectedMapsInfoRequest.prototype.getMD5 = function() { return "59778fc7286f314a408be52b4611a8b4"; };

ProjectedMapsInfoRequest.prototype.getID = function() { return this.__id__; };

ProjectedMapsInfoRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new ProjectedMapsInfoRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.ProjectedMapsInfoResponse=f();}})(function(){var define,module,exports;'use strict';

function ProjectedMapsInfoResponse() {
    this.__id__ = 0;
};

ProjectedMapsInfoResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

ProjectedMapsInfoResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

ProjectedMapsInfoResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

ProjectedMapsInfoResponse.prototype.echo = function() { return ""; };

ProjectedMapsInfoResponse.prototype.getType = function() { return "map_msgs/ProjectedMapsInfo"; };

ProjectedMapsInfoResponse.prototype.getMD5 = function() { return "223a7c48f052c5181dd525823dcc67fc"; };

ProjectedMapsInfoResponse.prototype.getID = function() { return this.__id__; };

ProjectedMapsInfoResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new ProjectedMapsInfoResponse(); };

});
