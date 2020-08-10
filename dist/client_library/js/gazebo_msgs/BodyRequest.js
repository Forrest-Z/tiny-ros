(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.BodyRequestRequest=f();}})(function(){var define,module,exports;'use strict';

function BodyRequestRequest() {
    this.__id__ = 0;
    this.body_name = "";
};

BodyRequestRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_body_name = new TextEncoder('utf8');
    var utf8array_body_name = encoder_body_name.encode(this.body_name);
    buff[offset + 0] = (utf8array_body_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_body_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_body_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_body_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_body_name.length; i++) {
        buff[offset + i] = utf8array_body_name[i];
    }
    offset += utf8array_body_name.length;
    return offset;
};

BodyRequestRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_body_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_body_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_body_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_body_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_body_name = new TextDecoder('utf8');
    this.body_name = decoder_body_name.decode(buff.slice(offset, offset + length_body_name));
    offset += length_body_name;
    return offset;
};

BodyRequestRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_body_name = new TextEncoder('utf8');
    var utf8array_body_name = encoder_body_name.encode(this.body_name);
    length += 4;
    length += utf8array_body_name.length;
    return length;
};

BodyRequestRequest.prototype.echo = function() { return ""; };

BodyRequestRequest.prototype.getType = function() { return "gazebo_msgs/BodyRequest"; };

BodyRequestRequest.prototype.getMD5 = function() { return "d1c66fbceb0ee1b51e3b09ac030dedec"; };

BodyRequestRequest.prototype.getID = function() { return this.__id__; };

BodyRequestRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new BodyRequestRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.BodyRequestResponse=f();}})(function(){var define,module,exports;'use strict';

function BodyRequestResponse() {
    this.__id__ = 0;
};

BodyRequestResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

BodyRequestResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

BodyRequestResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    return length;
};

BodyRequestResponse.prototype.echo = function() { return ""; };

BodyRequestResponse.prototype.getType = function() { return "gazebo_msgs/BodyRequest"; };

BodyRequestResponse.prototype.getMD5 = function() { return "e0caf2eb9951542b962f95924c6eeb34"; };

BodyRequestResponse.prototype.getID = function() { return this.__id__; };

BodyRequestResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new BodyRequestResponse(); };

});
