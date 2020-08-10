(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf2_msgs==="undefined"){g.tf2_msgs={};}g.tf2_msgs.LookupTransformActionFeedback=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/actionlib_msgs/GoalStatus.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tf2_msgs/LookupTransformFeedback.js'></script>");

function LookupTransformActionFeedback() {
    this.header = std_msgs.Header();
    this.status = actionlib_msgs.GoalStatus();
    this.feedback = tf2_msgs.LookupTransformFeedback();
};

LookupTransformActionFeedback.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.status.serialize(buff, offset);
    offset += this.feedback.serialize(buff, offset);
    return offset;
};

LookupTransformActionFeedback.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.status.deserialize(buff, offset);
    offset += this.feedback.deserialize(buff, offset);
    return offset;
};

LookupTransformActionFeedback.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.status.serializedLength();
    length += this.feedback.serializedLength();
    return length;
};

LookupTransformActionFeedback.prototype.echo = function() { return ""; };

LookupTransformActionFeedback.prototype.getType = function() { return "tf2_msgs/LookupTransformActionFeedback"; };

LookupTransformActionFeedback.prototype.getMD5 = function() { return "f432c3fa33a3ac9586ace248b606de3a"; };

LookupTransformActionFeedback.prototype.getID = function() { return 0; };

LookupTransformActionFeedback.prototype.setID = function(id) { };

return function() { return new LookupTransformActionFeedback(); };

});
