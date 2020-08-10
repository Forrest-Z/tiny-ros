(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.Transform=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Vector3.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Quaternion.js'></script>");

function Transform() {
    this.translation = geometry_msgs.Vector3();
    this.rotation = geometry_msgs.Quaternion();
};

Transform.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.translation.serialize(buff, offset);
    offset += this.rotation.serialize(buff, offset);
    return offset;
};

Transform.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.translation.deserialize(buff, offset);
    offset += this.rotation.deserialize(buff, offset);
    return offset;
};

Transform.prototype.serializedLength = function() {
    var length = 0;
    length += this.translation.serializedLength();
    length += this.rotation.serializedLength();
    return length;
};

Transform.prototype.echo = function() { return ""; };

Transform.prototype.getType = function() { return "geometry_msgs/Transform"; };

Transform.prototype.getMD5 = function() { return "2526ee1b1cc2e723e386c3c1b048ba72"; };

Transform.prototype.getID = function() { return 0; };

Transform.prototype.setID = function(id) { };

return function() { return new Transform(); };

});
