(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.AccelWithCovarianceStamped=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/AccelWithCovariance.js'></script>");

function AccelWithCovarianceStamped() {
    this.header = std_msgs.Header();
    this.accel = geometry_msgs.AccelWithCovariance();
};

AccelWithCovarianceStamped.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.accel.serialize(buff, offset);
    return offset;
};

AccelWithCovarianceStamped.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.accel.deserialize(buff, offset);
    return offset;
};

AccelWithCovarianceStamped.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.accel.serializedLength();
    return length;
};

AccelWithCovarianceStamped.prototype.echo = function() { return ""; };

AccelWithCovarianceStamped.prototype.getType = function() { return "geometry_msgs/AccelWithCovarianceStamped"; };

AccelWithCovarianceStamped.prototype.getMD5 = function() { return "efd9e7d0b5ca262cc8b05aa8e97c984f"; };

AccelWithCovarianceStamped.prototype.getID = function() { return 0; };

AccelWithCovarianceStamped.prototype.setID = function(id) { };

return function() { return new AccelWithCovarianceStamped(); };

});
