(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.QuaternionStamped=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Quaternion.js'></script>");

function QuaternionStamped() {
    this.header = std_msgs.Header();
    this.quaternion = geometry_msgs.Quaternion();
};

QuaternionStamped.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.quaternion.serialize(buff, offset);
    return offset;
};

QuaternionStamped.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.quaternion.deserialize(buff, offset);
    return offset;
};

QuaternionStamped.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.quaternion.serializedLength();
    return length;
};

QuaternionStamped.prototype.echo = function() { return ""; };

QuaternionStamped.prototype.getType = function() { return "geometry_msgs/QuaternionStamped"; };

QuaternionStamped.prototype.getMD5 = function() { return "69e39922feb9ec6eaf93755f93fce2cf"; };

QuaternionStamped.prototype.getID = function() { return 0; };

QuaternionStamped.prototype.setID = function(id) { };

return function() { return new QuaternionStamped(); };

});
