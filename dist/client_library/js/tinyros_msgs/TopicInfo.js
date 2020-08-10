(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tinyros_msgs==="undefined"){g.tinyros_msgs={};}g.tinyros_msgs.TopicInfo=f();}})(function(){var define,module,exports;'use strict';

function TopicInfo() {
    this.topic_id = 0;
    this.topic_name = "";
    this.message_type = "";
    this.md5sum = "";
    this.buffer_size = 0;
    this.negotiated = false;
    this.node = "";

    // ENUM{
    this.ID_PUBLISHER = 0;
    this.ID_SUBSCRIBER = 1;
    this.ID_SERVICE_SERVER = 2;
    this.ID_SERVICE_CLIENT = 4;
    this.ID_ROSTOPIC_REQUEST = 6;
    this.ID_ROSSERVICE_REQUEST = 7;
    this.ID_LOG = 8;
    this.ID_TIME = 9;
    this.ID_NEGOTIATED = 10;
    this.ID_SESSION_ID = 11;
    // }ENUM
};

TopicInfo.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.topic_id) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.topic_id) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.topic_id) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.topic_id) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_topic_name = new TextEncoder('utf8');
    var utf8array_topic_name = encoder_topic_name.encode(this.topic_name);
    buff[offset + 0] = (utf8array_topic_name.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_topic_name.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_topic_name.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_topic_name.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_topic_name.length; i++) {
        buff[offset + i] = utf8array_topic_name[i];
    }
    offset += utf8array_topic_name.length;
    var encoder_message_type = new TextEncoder('utf8');
    var utf8array_message_type = encoder_message_type.encode(this.message_type);
    buff[offset + 0] = (utf8array_message_type.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_message_type.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_message_type.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_message_type.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_message_type.length; i++) {
        buff[offset + i] = utf8array_message_type[i];
    }
    offset += utf8array_message_type.length;
    var encoder_md5sum = new TextEncoder('utf8');
    var utf8array_md5sum = encoder_md5sum.encode(this.md5sum);
    buff[offset + 0] = (utf8array_md5sum.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_md5sum.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_md5sum.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_md5sum.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_md5sum.length; i++) {
        buff[offset + i] = utf8array_md5sum[i];
    }
    offset += utf8array_md5sum.length;
    buff[offset + 0] = ((+this.buffer_size) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.buffer_size) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.buffer_size) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.buffer_size) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset] = this.negotiated === false ? 0 : 1;
    offset += 1;
    var encoder_node = new TextEncoder('utf8');
    var utf8array_node = encoder_node.encode(this.node);
    buff[offset + 0] = (utf8array_node.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_node.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_node.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_node.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_node.length; i++) {
        buff[offset + i] = utf8array_node[i];
    }
    offset += utf8array_node.length;
    return offset;
};

TopicInfo.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.topic_id = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.topic_id |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.topic_id |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.topic_id |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_topic_name = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_topic_name |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_topic_name |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_topic_name |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_topic_name = new TextDecoder('utf8');
    this.topic_name = decoder_topic_name.decode(buff.slice(offset, offset + length_topic_name));
    offset += length_topic_name;
    var length_message_type = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_message_type |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_message_type |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_message_type |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_message_type = new TextDecoder('utf8');
    this.message_type = decoder_message_type.decode(buff.slice(offset, offset + length_message_type));
    offset += length_message_type;
    var length_md5sum = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_md5sum |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_md5sum |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_md5sum |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_md5sum = new TextDecoder('utf8');
    this.md5sum = decoder_md5sum.decode(buff.slice(offset, offset + length_md5sum));
    offset += length_md5sum;
    this.buffer_size = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.buffer_size |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.buffer_size |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.buffer_size |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.negotiated = buff[offset] !== 0 ? true : false;
    offset += 1;
    var length_node = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_node |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_node |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_node |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_node = new TextDecoder('utf8');
    this.node = decoder_node.decode(buff.slice(offset, offset + length_node));
    offset += length_node;
    return offset;
};

TopicInfo.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    var encoder_topic_name = new TextEncoder('utf8');
    var utf8array_topic_name = encoder_topic_name.encode(this.topic_name);
    length += 4;
    length += utf8array_topic_name.length;
    var encoder_message_type = new TextEncoder('utf8');
    var utf8array_message_type = encoder_message_type.encode(this.message_type);
    length += 4;
    length += utf8array_message_type.length;
    var encoder_md5sum = new TextEncoder('utf8');
    var utf8array_md5sum = encoder_md5sum.encode(this.md5sum);
    length += 4;
    length += utf8array_md5sum.length;
    length += 4
    length += 1
    var encoder_node = new TextEncoder('utf8');
    var utf8array_node = encoder_node.encode(this.node);
    length += 4;
    length += utf8array_node.length;
    return length;
};

TopicInfo.prototype.echo = function() { return ""; };

TopicInfo.prototype.getType = function() { return "tinyros_msgs/TopicInfo"; };

TopicInfo.prototype.getMD5 = function() { return "76d40676946fcde66f228def7575451a"; };

TopicInfo.prototype.getID = function() { return 0; };

TopicInfo.prototype.setID = function(id) { };

return function() { return new TopicInfo(); };

});
