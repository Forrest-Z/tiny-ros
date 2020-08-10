(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.WrenchStamped=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Wrench.js'></script>");

function WrenchStamped() {
    this.header = std_msgs.Header();
    this.wrench = geometry_msgs.Wrench();
};

WrenchStamped.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.wrench.serialize(buff, offset);
    return offset;
};

WrenchStamped.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.wrench.deserialize(buff, offset);
    return offset;
};

WrenchStamped.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.wrench.serializedLength();
    return length;
};

WrenchStamped.prototype.echo = function() { return ""; };

WrenchStamped.prototype.getType = function() { return "geometry_msgs/WrenchStamped"; };

WrenchStamped.prototype.getMD5 = function() { return "cf53874aa63609de4155ec8e9cf2c540"; };

WrenchStamped.prototype.getID = function() { return 0; };

WrenchStamped.prototype.setID = function(id) { };

return function() { return new WrenchStamped(); };

});
