(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.PoseWithCovarianceStamped=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/PoseWithCovariance.js'></script>");

function PoseWithCovarianceStamped() {
    this.header = std_msgs.Header();
    this.pose = geometry_msgs.PoseWithCovariance();
};

PoseWithCovarianceStamped.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.pose.serialize(buff, offset);
    return offset;
};

PoseWithCovarianceStamped.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.pose.deserialize(buff, offset);
    return offset;
};

PoseWithCovarianceStamped.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.pose.serializedLength();
    return length;
};

PoseWithCovarianceStamped.prototype.echo = function() { return ""; };

PoseWithCovarianceStamped.prototype.getType = function() { return "geometry_msgs/PoseWithCovarianceStamped"; };

PoseWithCovarianceStamped.prototype.getMD5 = function() { return "14ff1431078f35103bf1b202333b4704"; };

PoseWithCovarianceStamped.prototype.getID = function() { return 0; };

PoseWithCovarianceStamped.prototype.setID = function(id) { };

return function() { return new PoseWithCovarianceStamped(); };

});
