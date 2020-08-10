(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf2_msgs==="undefined"){g.tf2_msgs={};}g.tf2_msgs.LookupTransformAction=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tf2_msgs/LookupTransformActionGoal.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tf2_msgs/LookupTransformActionResult.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tf2_msgs/LookupTransformActionFeedback.js'></script>");

function LookupTransformAction() {
    this.action_goal = tf2_msgs.LookupTransformActionGoal();
    this.action_result = tf2_msgs.LookupTransformActionResult();
    this.action_feedback = tf2_msgs.LookupTransformActionFeedback();
};

LookupTransformAction.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.action_goal.serialize(buff, offset);
    offset += this.action_result.serialize(buff, offset);
    offset += this.action_feedback.serialize(buff, offset);
    return offset;
};

LookupTransformAction.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.action_goal.deserialize(buff, offset);
    offset += this.action_result.deserialize(buff, offset);
    offset += this.action_feedback.deserialize(buff, offset);
    return offset;
};

LookupTransformAction.prototype.serializedLength = function() {
    var length = 0;
    length += this.action_goal.serializedLength();
    length += this.action_result.serializedLength();
    length += this.action_feedback.serializedLength();
    return length;
};

LookupTransformAction.prototype.echo = function() { return ""; };

LookupTransformAction.prototype.getType = function() { return "tf2_msgs/LookupTransformAction"; };

LookupTransformAction.prototype.getMD5 = function() { return "49655e848adf6c64870a8eb763b94484"; };

LookupTransformAction.prototype.getID = function() { return 0; };

LookupTransformAction.prototype.setID = function(id) { };

return function() { return new LookupTransformAction(); };

});
