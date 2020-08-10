(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GetMapActionGoal=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/actionlib_msgs/GoalID.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/GetMapGoal.js'></script>");

function GetMapActionGoal() {
    this.header = std_msgs.Header();
    this.goal_id = actionlib_msgs.GoalID();
    this.goal = nav_msgs.GetMapGoal();
};

GetMapActionGoal.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.goal_id.serialize(buff, offset);
    offset += this.goal.serialize(buff, offset);
    return offset;
};

GetMapActionGoal.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.goal_id.deserialize(buff, offset);
    offset += this.goal.deserialize(buff, offset);
    return offset;
};

GetMapActionGoal.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.goal_id.serializedLength();
    length += this.goal.serializedLength();
    return length;
};

GetMapActionGoal.prototype.echo = function() { return ""; };

GetMapActionGoal.prototype.getType = function() { return "nav_msgs/GetMapActionGoal"; };

GetMapActionGoal.prototype.getMD5 = function() { return "8aea83336b4ee626241742bb14b14d90"; };

GetMapActionGoal.prototype.getID = function() { return 0; };

GetMapActionGoal.prototype.setID = function(id) { };

return function() { return new GetMapActionGoal(); };

});
