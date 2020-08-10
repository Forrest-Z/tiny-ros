(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tinyros_msgs==="undefined"){g.tinyros_msgs={};}g.tinyros_msgs.Log=f();}})(function(){var define,module,exports;'use strict';

function Log() {
    this.level = 0;
    this.msg = "";

    // ENUM{
    this.ROSDEBUG = 0;
    this.ROSINFO = 1;
    this.ROSWARN = 2;
    this.ROSERROR = 3;
    this.ROSFATAL = 4;
    // }ENUM
};

Log.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.level) >> (8 * 0)) & 0xFF;
    offset += 1;
    var encoder_msg = new TextEncoder('utf8');
    var utf8array_msg = encoder_msg.encode(this.msg);
    buff[offset + 0] = (utf8array_msg.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_msg.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_msg.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_msg.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_msg.length; i++) {
        buff[offset + i] = utf8array_msg[i];
    }
    offset += utf8array_msg.length;
    return offset;
};

Log.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.level = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    var length_msg = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_msg |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_msg |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_msg |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_msg = new TextDecoder('utf8');
    this.msg = decoder_msg.decode(buff.slice(offset, offset + length_msg));
    offset += length_msg;
    return offset;
};

Log.prototype.serializedLength = function() {
    var length = 0;
    length += 1
    var encoder_msg = new TextEncoder('utf8');
    var utf8array_msg = encoder_msg.encode(this.msg);
    length += 4;
    length += utf8array_msg.length;
    return length;
};

Log.prototype.echo = function() { return ""; };

Log.prototype.getType = function() { return "tinyros_msgs/Log"; };

Log.prototype.getMD5 = function() { return "0bd74339b4d77cb15766d831a3d15eeb"; };

Log.prototype.getID = function() { return 0; };

Log.prototype.setID = function(id) { };

return function() { return new Log(); };

});
