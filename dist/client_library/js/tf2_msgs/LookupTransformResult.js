(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf2_msgs==="undefined"){g.tf2_msgs={};}g.tf2_msgs.LookupTransformResult=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/TransformStamped.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tf2_msgs/TF2Error.js'></script>");

function LookupTransformResult() {
    this.transform = geometry_msgs.TransformStamped();
    this.error = tf2_msgs.TF2Error();
};

LookupTransformResult.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.transform.serialize(buff, offset);
    offset += this.error.serialize(buff, offset);
    return offset;
};

LookupTransformResult.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.transform.deserialize(buff, offset);
    offset += this.error.deserialize(buff, offset);
    return offset;
};

LookupTransformResult.prototype.serializedLength = function() {
    var length = 0;
    length += this.transform.serializedLength();
    length += this.error.serializedLength();
    return length;
};

LookupTransformResult.prototype.echo = function() { return ""; };

LookupTransformResult.prototype.getType = function() { return "tf2_msgs/LookupTransformResult"; };

LookupTransformResult.prototype.getMD5 = function() { return "7be4fc6719f512bc94491db1ccda6aee"; };

LookupTransformResult.prototype.getID = function() { return 0; };

LookupTransformResult.prototype.setID = function(id) { };

return function() { return new LookupTransformResult(); };

});
