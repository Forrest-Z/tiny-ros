(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GetMapAction=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/GetMapActionGoal.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/GetMapActionResult.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/GetMapActionFeedback.js'></script>");

function GetMapAction() {
    this.action_goal = nav_msgs.GetMapActionGoal();
    this.action_result = nav_msgs.GetMapActionResult();
    this.action_feedback = nav_msgs.GetMapActionFeedback();
};

GetMapAction.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.action_goal.serialize(buff, offset);
    offset += this.action_result.serialize(buff, offset);
    offset += this.action_feedback.serialize(buff, offset);
    return offset;
};

GetMapAction.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.action_goal.deserialize(buff, offset);
    offset += this.action_result.deserialize(buff, offset);
    offset += this.action_feedback.deserialize(buff, offset);
    return offset;
};

GetMapAction.prototype.serializedLength = function() {
    var length = 0;
    length += this.action_goal.serializedLength();
    length += this.action_result.serializedLength();
    length += this.action_feedback.serializedLength();
    return length;
};

GetMapAction.prototype.echo = function() { return ""; };

GetMapAction.prototype.getType = function() { return "nav_msgs/GetMapAction"; };

GetMapAction.prototype.getMD5 = function() { return "10a4e277d7b8e53bfc3df54d98b3edb1"; };

GetMapAction.prototype.getID = function() { return 0; };

GetMapAction.prototype.setID = function(id) { };

return function() { return new GetMapAction(); };

});
