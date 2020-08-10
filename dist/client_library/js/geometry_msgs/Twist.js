(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.Twist=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Vector3.js'></script>");

function Twist() {
    this.linear = geometry_msgs.Vector3();
    this.angular = geometry_msgs.Vector3();
};

Twist.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.linear.serialize(buff, offset);
    offset += this.angular.serialize(buff, offset);
    return offset;
};

Twist.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.linear.deserialize(buff, offset);
    offset += this.angular.deserialize(buff, offset);
    return offset;
};

Twist.prototype.serializedLength = function() {
    var length = 0;
    length += this.linear.serializedLength();
    length += this.angular.serializedLength();
    return length;
};

Twist.prototype.echo = function() { return ""; };

Twist.prototype.getType = function() { return "geometry_msgs/Twist"; };

Twist.prototype.getMD5 = function() { return "29e7e4839b73f684ad08b19dc12c9c70"; };

Twist.prototype.getID = function() { return 0; };

Twist.prototype.setID = function(id) { };

return function() { return new Twist(); };

});
