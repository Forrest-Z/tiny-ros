(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.Vector3Stamped=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Vector3.js'></script>");

function Vector3Stamped() {
    this.header = std_msgs.Header();
    this.vector = geometry_msgs.Vector3();
};

Vector3Stamped.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.vector.serialize(buff, offset);
    return offset;
};

Vector3Stamped.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.vector.deserialize(buff, offset);
    return offset;
};

Vector3Stamped.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.vector.serializedLength();
    return length;
};

Vector3Stamped.prototype.echo = function() { return ""; };

Vector3Stamped.prototype.getType = function() { return "geometry_msgs/Vector3Stamped"; };

Vector3Stamped.prototype.getMD5 = function() { return "4b85025eb6f70f6b1e0cefbb75f69ac2"; };

Vector3Stamped.prototype.getID = function() { return 0; };

Vector3Stamped.prototype.setID = function(id) { };

return function() { return new Vector3Stamped(); };

});
