(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf==="undefined"){g.tf={};}g.tf.FrameGraphRequest=f();}})(function(){var define,module,exports;'use strict';

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

FrameGraphRequest.prototype.getType = function() { return "tf/FrameGraph"; };

FrameGraphRequest.prototype.getMD5 = function() { return "d66179e20d21cee31291f40363976e1d"; };

FrameGraphRequest.prototype.getID = function() { return this.__id__; };

FrameGraphRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new FrameGraphRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf==="undefined"){g.tf={};}g.tf.FrameGraphResponse=f();}})(function(){var define,module,exports;'use strict';

function FrameGraphResponse() {
    this.__id__ = 0;
    this.dot_graph = "";
};

FrameGraphResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_dot_graph = new TextEncoder('utf8');
    var utf8array_dot_graph = encoder_dot_graph.encode(this.dot_graph);
    buff[offset + 0] = (utf8array_dot_graph.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_dot_graph.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_dot_graph.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_dot_graph.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_dot_graph.length; i++) {
        buff[offset + i] = utf8array_dot_graph[i];
    }
    offset += utf8array_dot_graph.length;
    return offset;
};

FrameGraphResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_dot_graph = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_dot_graph |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_dot_graph |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_dot_graph |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_dot_graph = new TextDecoder('utf8');
    this.dot_graph = decoder_dot_graph.decode(buff.slice(offset, offset + length_dot_graph));
    offset += length_dot_graph;
    return offset;
};

FrameGraphResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_dot_graph = new TextEncoder('utf8');
    var utf8array_dot_graph = encoder_dot_graph.encode(this.dot_graph);
    length += 4;
    length += utf8array_dot_graph.length;
    return length;
};

FrameGraphResponse.prototype.echo = function() { return ""; };

FrameGraphResponse.prototype.getType = function() { return "tf/FrameGraph"; };

FrameGraphResponse.prototype.getMD5 = function() { return "8528671f80a5c0815f9700a446efbc36"; };

FrameGraphResponse.prototype.getID = function() { return this.__id__; };

FrameGraphResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new FrameGraphResponse(); };

});
