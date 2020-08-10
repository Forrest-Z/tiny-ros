(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.SetMapProjectionsRequest=f();}})(function(){var define,module,exports;'use strict';

function SetMapProjectionsRequest() {
    this.__id__ = 0;
};

SetMapProjectionsRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

SetMapProjectionsRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

SetMapProjectionsRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

SetMapProjectionsRequest.prototype.echo = function() { return ""; };

SetMapProjectionsRequest.prototype.getType = function() { return "map_msgs/SetMapProjections"; };

SetMapProjectionsRequest.prototype.getMD5 = function() { return "26dbba584adf9695d68b8667830be463"; };

SetMapProjectionsRequest.prototype.getID = function() { return this.__id__; };

SetMapProjectionsRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetMapProjectionsRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.map_msgs==="undefined"){g.map_msgs={};}g.map_msgs.SetMapProjectionsResponse=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/map_msgs/ProjectedMapInfo.js'></script>");

function SetMapProjectionsResponse() {
    this.__id__ = 0;
    this.projected_maps_info = new Array();
};

SetMapProjectionsResponse.prototype.serialize = function(buff, idx) {
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

SetMapProjectionsResponse.prototype.deserialize = function(buff, idx) {
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

SetMapProjectionsResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var length_projected_maps_info = this.projected_maps_info.length;
    length += 4;
    for (var i = 0; i < length_projected_maps_info; i++) {
        length += this.projected_maps_info[i].serializedLength();
    }
    return length;
};

SetMapProjectionsResponse.prototype.echo = function() { return ""; };

SetMapProjectionsResponse.prototype.getType = function() { return "map_msgs/SetMapProjections"; };

SetMapProjectionsResponse.prototype.getMD5 = function() { return "47b7931263dc316e5b0f0264428853e9"; };

SetMapProjectionsResponse.prototype.getID = function() { return this.__id__; };

SetMapProjectionsResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new SetMapProjectionsResponse(); };

});
