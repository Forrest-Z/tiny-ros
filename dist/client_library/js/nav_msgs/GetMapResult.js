(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GetMapResult=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/OccupancyGrid.js'></script>");

function GetMapResult() {
    this.map = nav_msgs.OccupancyGrid();
};

GetMapResult.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.map.serialize(buff, offset);
    return offset;
};

GetMapResult.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.map.deserialize(buff, offset);
    return offset;
};

GetMapResult.prototype.serializedLength = function() {
    var length = 0;
    length += this.map.serializedLength();
    return length;
};

GetMapResult.prototype.echo = function() { return ""; };

GetMapResult.prototype.getType = function() { return "nav_msgs/GetMapResult"; };

GetMapResult.prototype.getMD5 = function() { return "dd8eb0759b1a400b141d7f3238732c4d"; };

GetMapResult.prototype.getID = function() { return 0; };

GetMapResult.prototype.setID = function(id) { };

return function() { return new GetMapResult(); };

});
