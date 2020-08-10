(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf2_msgs==="undefined"){g.tf2_msgs={};}g.tf2_msgs.FrameGraphRequest=f();}})(function(){var define,module,exports;'use strict';

function FrameGraphRequest() {
    this.__id__ = 0;
};

FrameGraphRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

FrameGraphRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

FrameGraphRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

FrameGraphRequest.prototype.echo = function() { return ""; };

FrameGraphRequest.prototype.getType = function() { return "tf2_msgs/FrameGraph"; };

FrameGraphRequest.prototype.getMD5 = function() { return "aa023d3c31410363e0583979223258c8"; };

FrameGraphRequest.prototype.getID = function() { return this.__id__; };

FrameGraphRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new FrameGraphRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf2_msgs==="undefined"){g.tf2_msgs={};}g.tf2_msgs.FrameGraphResponse=f();}})(function(){var define,module,exports;'use strict';

function FrameGraphResponse() {
    this.__id__ = 0;
    this.frame_yaml = "";
};

FrameGraphResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_frame_yaml = new TextEncoder('utf8');
    var utf8array_frame_yaml = encoder_frame_yaml.encode(this.frame_yaml);
    buff[offset + 0] = (utf8array_frame_yaml.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_frame_yaml.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_frame_yaml.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_frame_yaml.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_frame_yaml.length; i++) {
        buff[offset + i] = utf8array_frame_yaml[i];
    }
    offset += utf8array_frame_yaml.length;
    return offset;
};

FrameGraphResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_frame_yaml = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_frame_yaml |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_frame_yaml |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_frame_yaml |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_frame_yaml = new TextDecoder('utf8');
    this.frame_yaml = decoder_frame_yaml.decode(buff.slice(offset, offset + length_frame_yaml));
    offset += length_frame_yaml;
    return offset;
};

FrameGraphResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_frame_yaml = new TextEncoder('utf8');
    var utf8array_frame_yaml = encoder_frame_yaml.encode(this.frame_yaml);
    length += 4;
    length += utf8array_frame_yaml.length;
    return length;
};

FrameGraphResponse.prototype.echo = function() { return ""; };

FrameGraphResponse.prototype.getType = function() { return "tf2_msgs/FrameGraph"; };

FrameGraphResponse.prototype.getMD5 = function() { return "97e361486f8cc8fb1a460cf17555126b"; };

FrameGraphResponse.prototype.getID = function() { return this.__id__; };

FrameGraphResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new FrameGraphResponse(); };

});
