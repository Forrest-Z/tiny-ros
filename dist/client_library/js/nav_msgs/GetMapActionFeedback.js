(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GetMapActionFeedback=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/actionlib_msgs/GoalStatus.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/GetMapFeedback.js'></script>");

function GetMapActionFeedback() {
    this.header = std_msgs.Header();
    this.status = actionlib_msgs.GoalStatus();
    this.feedback = nav_msgs.GetMapFeedback();
};

GetMapActionFeedback.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.status.serialize(buff, offset);
    offset += this.feedback.serialize(buff, offset);
    return offset;
};

GetMapActionFeedback.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.status.deserialize(buff, offset);
    offset += this.feedback.deserialize(buff, offset);
    return offset;
};

GetMapActionFeedback.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.status.serializedLength();
    length += this.feedback.serializedLength();
    return length;
};

GetMapActionFeedback.prototype.echo = function() { return ""; };

GetMapActionFeedback.prototype.getType = function() { return "nav_msgs/GetMapActionFeedback"; };

GetMapActionFeedback.prototype.getMD5 = function() { return "9ebb88ff2cf2120160bf2197071a69b6"; };

GetMapActionFeedback.prototype.getID = function() { return 0; };

GetMapActionFeedback.prototype.setID = function(id) { };

return function() { return new GetMapActionFeedback(); };

});
