(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf2_msgs==="undefined"){g.tf2_msgs={};}g.tf2_msgs.LookupTransformActionResult=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/actionlib_msgs/GoalStatus.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tf2_msgs/LookupTransformResult.js'></script>");

function LookupTransformActionResult() {
    this.header = std_msgs.Header();
    this.status = actionlib_msgs.GoalStatus();
    this.result = tf2_msgs.LookupTransformResult();
};

LookupTransformActionResult.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.status.serialize(buff, offset);
    offset += this.result.serialize(buff, offset);
    return offset;
};

LookupTransformActionResult.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.status.deserialize(buff, offset);
    offset += this.result.deserialize(buff, offset);
    return offset;
};

LookupTransformActionResult.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.status.serializedLength();
    length += this.result.serializedLength();
    return length;
};

LookupTransformActionResult.prototype.echo = function() { return ""; };

LookupTransformActionResult.prototype.getType = function() { return "tf2_msgs/LookupTransformActionResult"; };

LookupTransformActionResult.prototype.getMD5 = function() { return "5a8abe079c2126ea9966563c5cae6d29"; };

LookupTransformActionResult.prototype.getID = function() { return 0; };

LookupTransformActionResult.prototype.setID = function(id) { };

return function() { return new LookupTransformActionResult(); };

});
