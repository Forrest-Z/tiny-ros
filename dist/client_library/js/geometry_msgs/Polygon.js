(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.Polygon=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Point32.js'></script>");

function Polygon() {
    this.points = new Array();
};

Polygon.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var length_points = this.points.length;
    buff[offset + 0] = (length_points >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_points >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_points >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_points >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_points; i++) {
        offset += this.points[i].serialize(buff, offset);
    }
    return offset;
};

Polygon.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_points = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_points |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_points |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_points |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.points = new Array(length_points);
    for (var i = 0; i < length_points; i++) {
        this.points[i] = geometry_msgs.Point32();
        offset += this.points[i].deserialize(buff, offset);
    }
    return offset;
};

Polygon.prototype.serializedLength = function() {
    var length = 0;
    var length_points = this.points.length;
    length += 4;
    for (var i = 0; i < length_points; i++) {
        length += this.points[i].serializedLength();
    }
    return length;
};

Polygon.prototype.echo = function() { return ""; };

Polygon.prototype.getType = function() { return "geometry_msgs/Polygon"; };

Polygon.prototype.getMD5 = function() { return "f94a78a947b7879954bd14397db4bc9d"; };

Polygon.prototype.getID = function() { return 0; };

Polygon.prototype.setID = function(id) { };

return function() { return new Polygon(); };

});
