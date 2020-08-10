(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.rosgraph_msgs==="undefined"){g.rosgraph_msgs={};}g.rosgraph_msgs.Log=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function Log() {
    this.header = std_msgs.Header();
    this.level = 0;
    this.name = "";
    this.msg = "";
    this.file = "";
    this.function = "";
    this.line = 0;
    this.topics = new Array();

    // ENUM{
    this.DEBUG = 1 ;
    this.INFO = 2  ;
    this.WARN = 4  ;
    this.ERROR = 8 ;
    this.FATAL = 16 ;
    // }ENUM
};

Log.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    buff[offset + 0] = ((+this.level) >> (8 * 0)) & 0xFF;
    offset += 1;
    var encoder_name = new TextEncoder('utf8');
    var utf8array_name = encoder_name.encode(this.name);
    buff[offset + 0] = (utf8array_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_name.length; i++) {
        buff[offset + i] = utf8array_name[i];
    }
    offset += utf8array_name.length;
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
    var encoder_file = new TextEncoder('utf8');
    var utf8array_file = encoder_file.encode(this.file);
    buff[offset + 0] = (utf8array_file.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_file.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_file.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_file.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_file.length; i++) {
        buff[offset + i] = utf8array_file[i];
    }
    offset += utf8array_file.length;
    var encoder_function = new TextEncoder('utf8');
    var utf8array_function = encoder_function.encode(this.function);
    buff[offset + 0] = (utf8array_function.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_function.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_function.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_function.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_function.length; i++) {
        buff[offset + i] = utf8array_function[i];
    }
    offset += utf8array_function.length;
    buff[offset + 0] = ((+this.line) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.line) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.line) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.line) >> (8 * 3)) & 0xFF;
    offset += 4;
    var length_topics = this.topics.length;
    buff[offset + 0] = (length_topics >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_topics >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_topics >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_topics >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_topics; i++) {
        var encoder_topicsi = new TextEncoder('utf8');
        var utf8array_topicsi = encoder_topicsi.encode(this.topics[i]);
        buff[offset + 0] = (utf8array_topicsi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_topicsi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_topicsi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_topicsi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_topicsi.length; i++) {
            buff[offset + i] = utf8array_topicsi[i];
        }
        offset += utf8array_topicsi.length;
    }
    return offset;
};

Log.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    this.level = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    var length_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_name = new TextDecoder('utf8');
    this.name = decoder_name.decode(buff.slice(offset, offset + length_name));
    offset += length_name;
    var length_msg = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_msg |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_msg |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_msg |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_msg = new TextDecoder('utf8');
    this.msg = decoder_msg.decode(buff.slice(offset, offset + length_msg));
    offset += length_msg;
    var length_file = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_file |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_file |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_file |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_file = new TextDecoder('utf8');
    this.file = decoder_file.decode(buff.slice(offset, offset + length_file));
    offset += length_file;
    var length_function = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_function |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_function |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_function |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_function = new TextDecoder('utf8');
    this.function = decoder_function.decode(buff.slice(offset, offset + length_function));
    offset += length_function;
    this.line = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.line |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.line |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.line |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_topics = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_topics |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_topics |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_topics |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.topics = new Array(length_topics);
    for (var i = 0; i < length_topics; i++) {
        var length_topicsi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_topicsi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_topicsi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_topicsi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_topicsi = new TextDecoder('utf8');
        this.topics[i] = decoder_topicsi.decode(buff.slice(offset, offset + length_topicsi));
        offset += length_topicsi;
    }
    return offset;
};

Log.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 1
    var encoder_name = new TextEncoder('utf8');
    var utf8array_name = encoder_name.encode(this.name);
    length += 4;
    length += utf8array_name.length;
    var encoder_msg = new TextEncoder('utf8');
    var utf8array_msg = encoder_msg.encode(this.msg);
    length += 4;
    length += utf8array_msg.length;
    var encoder_file = new TextEncoder('utf8');
    var utf8array_file = encoder_file.encode(this.file);
    length += 4;
    length += utf8array_file.length;
    var encoder_function = new TextEncoder('utf8');
    var utf8array_function = encoder_function.encode(this.function);
    length += 4;
    length += utf8array_function.length;
    length += 4
    var length_topics = this.topics.length;
    length += 4;
    for (var i = 0; i < length_topics; i++) {
        var encoder_topicsi = new TextEncoder('utf8');
        var utf8array_topicsi = encoder_topicsi.encode(this.topics[i]);
        length += 4;
        length += utf8array_topicsi.length;
    }
    return length;
};

Log.prototype.echo = function() { return ""; };

Log.prototype.getType = function() { return "rosgraph_msgs/Log"; };

Log.prototype.getMD5 = function() { return "2de9daf47e984009074d74dbdd492d49"; };

Log.prototype.getID = function() { return 0; };

Log.prototype.setID = function(id) { };

return function() { return new Log(); };

});
