(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf2_msgs==="undefined"){g.tf2_msgs={};}g.tf2_msgs.LookupTransformActionGoal=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/actionlib_msgs/GoalID.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tf2_msgs/LookupTransformGoal.js'></script>");

function LookupTransformActionGoal() {
    this.header = std_msgs.Header();
    this.goal_id = actionlib_msgs.GoalID();
    this.goal = tf2_msgs.LookupTransformGoal();
};

LookupTransformActionGoal.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.goal_id.serialize(buff, offset);
    offset += this.goal.serialize(buff, offset);
    return offset;
};

LookupTransformActionGoal.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.goal_id.deserialize(buff, offset);
    offset += this.goal.deserialize(buff, offset);
    return offset;
};

LookupTransformActionGoal.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.goal_id.serializedLength();
    length += this.goal.serializedLength();
    return length;
};

LookupTransformActionGoal.prototype.echo = function() { return ""; };

LookupTransformActionGoal.prototype.getType = function() { return "tf2_msgs/LookupTransformActionGoal"; };

LookupTransformActionGoal.prototype.getMD5 = function() { return "b8a7d4ffa64f063b4df7b1dd3fc2bf79"; };

LookupTransformActionGoal.prototype.getID = function() { return 0; };

LookupTransformActionGoal.prototype.setID = function(id) { };

return function() { return new LookupTransformActionGoal(); };

});
