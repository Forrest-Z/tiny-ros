(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.LinkState=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Pose.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Twist.js'></script>");

function LinkState() {
    this.link_name = "";
    this.pose = geometry_msgs.Pose();
    this.twist = geometry_msgs.Twist();
    this.reference_frame = "";
};

LinkState.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var encoder_link_name = new TextEncoder('utf8');
    var utf8array_link_name = encoder_link_name.encode(this.link_name);
    buff[offset + 0] = (utf8array_link_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_link_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_link_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_link_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_link_name.length; i++) {
        buff[offset + i] = utf8array_link_name[i];
    }
    offset += utf8array_link_name.length;
    offset += this.pose.serialize(buff, offset);
    offset += this.twist.serialize(buff, offset);
    var encoder_reference_frame = new TextEncoder('utf8');
    var utf8array_reference_frame = encoder_reference_frame.encode(this.reference_frame);
    buff[offset + 0] = (utf8array_reference_frame.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_reference_frame.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_reference_frame.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_reference_frame.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_reference_frame.length; i++) {
        buff[offset + i] = utf8array_reference_frame[i];
    }
    offset += utf8array_reference_frame.length;
    return offset;
};

LinkState.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_link_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_link_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_link_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_link_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_link_name = new TextDecoder('utf8');
    this.link_name = decoder_link_name.decode(buff.slice(offset, offset + length_link_name));
    offset += length_link_name;
    offset += this.pose.deserialize(buff, offset);
    offset += this.twist.deserialize(buff, offset);
    var length_reference_frame = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_reference_frame |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_reference_frame |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_reference_frame |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_reference_frame = new TextDecoder('utf8');
    this.reference_frame = decoder_reference_frame.decode(buff.slice(offset, offset + length_reference_frame));
    offset += length_reference_frame;
    return offset;
};

LinkState.prototype.serializedLength = function() {
    var length = 0;
    var encoder_link_name = new TextEncoder('utf8');
    var utf8array_link_name = encoder_link_name.encode(this.link_name);
    length += 4;
    length += utf8array_link_name.length;
    length += this.pose.serializedLength();
    length += this.twist.serializedLength();
    var encoder_reference_frame = new TextEncoder('utf8');
    var utf8array_reference_frame = encoder_reference_frame.encode(this.reference_frame);
    length += 4;
    length += utf8array_reference_frame.length;
    return length;
};

LinkState.prototype.echo = function() { return ""; };

LinkState.prototype.getType = function() { return "gazebo_msgs/LinkState"; };

LinkState.prototype.getMD5 = function() { return "eb3584856a5c068877b12eba5fc9372d"; };

LinkState.prototype.getID = function() { return 0; };

LinkState.prototype.setID = function(id) { };

return function() { return new LinkState(); };

});
