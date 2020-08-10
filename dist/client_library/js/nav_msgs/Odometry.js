(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.Odometry=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/PoseWithCovariance.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/TwistWithCovariance.js'></script>");

function Odometry() {
    this.header = std_msgs.Header();
    this.child_frame_id = "";
    this.pose = geometry_msgs.PoseWithCovariance();
    this.twist = geometry_msgs.TwistWithCovariance();
};

Odometry.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var encoder_child_frame_id = new TextEncoder('utf8');
    var utf8array_child_frame_id = encoder_child_frame_id.encode(this.child_frame_id);
    buff[offset + 0] = (utf8array_child_frame_id.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_child_frame_id.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_child_frame_id.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_child_frame_id.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_child_frame_id.length; i++) {
        buff[offset + i] = utf8array_child_frame_id[i];
    }
    offset += utf8array_child_frame_id.length;
    offset += this.pose.serialize(buff, offset);
    offset += this.twist.serialize(buff, offset);
    return offset;
};

Odometry.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_child_frame_id = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_child_frame_id |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_child_frame_id |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_child_frame_id |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_child_frame_id = new TextDecoder('utf8');
    this.child_frame_id = decoder_child_frame_id.decode(buff.slice(offset, offset + length_child_frame_id));
    offset += length_child_frame_id;
    offset += this.pose.deserialize(buff, offset);
    offset += this.twist.deserialize(buff, offset);
    return offset;
};

Odometry.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var encoder_child_frame_id = new TextEncoder('utf8');
    var utf8array_child_frame_id = encoder_child_frame_id.encode(this.child_frame_id);
    length += 4;
    length += utf8array_child_frame_id.length;
    length += this.pose.serializedLength();
    length += this.twist.serializedLength();
    return length;
};

Odometry.prototype.echo = function() { return ""; };

Odometry.prototype.getType = function() { return "nav_msgs/Odometry"; };

Odometry.prototype.getMD5 = function() { return "8fbd8c2e0caeb7be9b30b66a3e735193"; };

Odometry.prototype.getID = function() { return 0; };

Odometry.prototype.setID = function(id) { };

return function() { return new Odometry(); };

});
