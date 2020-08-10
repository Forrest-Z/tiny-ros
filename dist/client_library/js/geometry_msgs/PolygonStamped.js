(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.PolygonStamped=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Polygon.js'></script>");

function PolygonStamped() {
    this.header = std_msgs.Header();
    this.polygon = geometry_msgs.Polygon();
};

PolygonStamped.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.polygon.serialize(buff, offset);
    return offset;
};

PolygonStamped.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.polygon.deserialize(buff, offset);
    return offset;
};

PolygonStamped.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.polygon.serializedLength();
    return length;
};

PolygonStamped.prototype.echo = function() { return ""; };

PolygonStamped.prototype.getType = function() { return "geometry_msgs/PolygonStamped"; };

PolygonStamped.prototype.getMD5 = function() { return "33bdf94066425e572879b25c9a51ed50"; };

PolygonStamped.prototype.getID = function() { return 0; };

PolygonStamped.prototype.setID = function(id) { };

return function() { return new PolygonStamped(); };

});
