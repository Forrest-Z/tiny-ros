(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.ModelState=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Pose.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Twist.js'></script>");

function ModelState() {
    this.model_name = "";
    this.pose = geometry_msgs.Pose();
    this.twist = geometry_msgs.Twist();
    this.reference_frame = "";
};

ModelState.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var encoder_model_name = new TextEncoder('utf8');
    var utf8array_model_name = encoder_model_name.encode(this.model_name);
    buff[offset + 0] = (utf8array_model_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_model_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_model_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_model_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_model_name.length; i++) {
        buff[offset + i] = utf8array_model_name[i];
    }
    offset += utf8array_model_name.length;
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

ModelState.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_model_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_model_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_model_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_model_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_model_name = new TextDecoder('utf8');
    this.model_name = decoder_model_name.decode(buff.slice(offset, offset + length_model_name));
    offset += length_model_name;
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

ModelState.prototype.serializedLength = function() {
    var length = 0;
    var encoder_model_name = new TextEncoder('utf8');
    var utf8array_model_name = encoder_model_name.encode(this.model_name);
    length += 4;
    length += utf8array_model_name.length;
    length += this.pose.serializedLength();
    length += this.twist.serializedLength();
    var encoder_reference_frame = new TextEncoder('utf8');
    var utf8array_reference_frame = encoder_reference_frame.encode(this.reference_frame);
    length += 4;
    length += utf8array_reference_frame.length;
    return length;
};

ModelState.prototype.echo = function() { return ""; };

ModelState.prototype.getType = function() { return "gazebo_msgs/ModelState"; };

ModelState.prototype.getMD5 = function() { return "dee4d802363b4d6bd1ed61e20c2c4635"; };

ModelState.prototype.getID = function() { return 0; };

ModelState.prototype.setID = function(id) { };

return function() { return new ModelState(); };

});
