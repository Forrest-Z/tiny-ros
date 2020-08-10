(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GetMapActionResult=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/actionlib_msgs/GoalStatus.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/GetMapResult.js'></script>");

function GetMapActionResult() {
    this.header = std_msgs.Header();
    this.status = actionlib_msgs.GoalStatus();
    this.result = nav_msgs.GetMapResult();
};

GetMapActionResult.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.status.serialize(buff, offset);
    offset += this.result.serialize(buff, offset);
    return offset;
};

GetMapActionResult.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.status.deserialize(buff, offset);
    offset += this.result.deserialize(buff, offset);
    return offset;
};

GetMapActionResult.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.status.serializedLength();
    length += this.result.serializedLength();
    return length;
};

GetMapActionResult.prototype.echo = function() { return ""; };

GetMapActionResult.prototype.getType = function() { return "nav_msgs/GetMapActionResult"; };

GetMapActionResult.prototype.getMD5 = function() { return "9c9f64758f2627a010c16b17ea745028"; };

GetMapActionResult.prototype.getID = function() { return 0; };

GetMapActionResult.prototype.setID = function(id) { };

return function() { return new GetMapActionResult(); };

});
