(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.PointStamped=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Point.js'></script>");

function PointStamped() {
    this.header = std_msgs.Header();
    this.point = geometry_msgs.Point();
};

PointStamped.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.point.serialize(buff, offset);
    return offset;
};

PointStamped.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.point.deserialize(buff, offset);
    return offset;
};

PointStamped.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.point.serializedLength();
    return length;
};

PointStamped.prototype.echo = function() { return ""; };

PointStamped.prototype.getType = function() { return "geometry_msgs/PointStamped"; };

PointStamped.prototype.getMD5 = function() { return "d34e83bdbef7bf4b617a6293aab8390e"; };

PointStamped.prototype.getID = function() { return 0; };

PointStamped.prototype.setID = function(id) { };

return function() { return new PointStamped(); };

});
