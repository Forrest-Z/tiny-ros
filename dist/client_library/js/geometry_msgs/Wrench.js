(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.Wrench=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Vector3.js'></script>");

function Wrench() {
    this.force = geometry_msgs.Vector3();
    this.torque = geometry_msgs.Vector3();
};

Wrench.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.force.serialize(buff, offset);
    offset += this.torque.serialize(buff, offset);
    return offset;
};

Wrench.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.force.deserialize(buff, offset);
    offset += this.torque.deserialize(buff, offset);
    return offset;
};

Wrench.prototype.serializedLength = function() {
    var length = 0;
    length += this.force.serializedLength();
    length += this.torque.serializedLength();
    return length;
};

Wrench.prototype.echo = function() { return ""; };

Wrench.prototype.getType = function() { return "geometry_msgs/Wrench"; };

Wrench.prototype.getMD5 = function() { return "02d01d4a8dc253c7b42d4c9866201aee"; };

Wrench.prototype.getID = function() { return 0; };

Wrench.prototype.setID = function(id) { };

return function() { return new Wrench(); };

});
