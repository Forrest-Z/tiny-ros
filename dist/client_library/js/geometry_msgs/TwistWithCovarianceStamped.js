(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.geometry_msgs==="undefined"){g.geometry_msgs={};}g.geometry_msgs.TwistWithCovarianceStamped=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/TwistWithCovariance.js'></script>");

function TwistWithCovarianceStamped() {
    this.header = std_msgs.Header();
    this.twist = geometry_msgs.TwistWithCovariance();
};

TwistWithCovarianceStamped.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    offset += this.twist.serialize(buff, offset);
    return offset;
};

TwistWithCovarianceStamped.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    offset += this.twist.deserialize(buff, offset);
    return offset;
};

TwistWithCovarianceStamped.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += this.twist.serializedLength();
    return length;
};

TwistWithCovarianceStamped.prototype.echo = function() { return ""; };

TwistWithCovarianceStamped.prototype.getType = function() { return "geometry_msgs/TwistWithCovarianceStamped"; };

TwistWithCovarianceStamped.prototype.getMD5 = function() { return "2cbcab62cac39de1d1d01785b99ba778"; };

TwistWithCovarianceStamped.prototype.getID = function() { return 0; };

TwistWithCovarianceStamped.prototype.setID = function(id) { };

return function() { return new TwistWithCovarianceStamped(); };

});
