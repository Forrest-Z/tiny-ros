(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Header=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Time.js'></script>");

function Header() {
    this.seq = 0;
    this.stamp = tinyros.Time();
    this.frame_id = "";
};

Header.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.seq) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.seq) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.seq) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.seq) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.stamp.sec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.stamp.sec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.stamp.sec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.stamp.sec) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.stamp.nsec) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.stamp.nsec) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.stamp.nsec) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.stamp.nsec) >> (8 * 3)) & 0xFF;
    offset += 4;
    var encoder_frame_id = new TextEncoder('utf8');
    var utf8array_frame_id = encoder_frame_id.encode(this.frame_id);
    buff[offset + 0] = (utf8array_frame_id.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_frame_id.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_frame_id.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_frame_id.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_frame_id.length; i++) {
        buff[offset + i] = utf8array_frame_id[i];
    }
    offset += utf8array_frame_id.length;
    return offset;
};

Header.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.seq = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.seq |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.seq |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.seq |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.stamp.sec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.stamp.sec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.stamp.sec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.stamp.sec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.stamp.nsec = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.stamp.nsec |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.stamp.nsec |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.stamp.nsec |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var length_frame_id = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_frame_id |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_frame_id |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_frame_id |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_frame_id = new TextDecoder('utf8');
    this.frame_id = decoder_frame_id.decode(buff.slice(offset, offset + length_frame_id));
    offset += length_frame_id;
    return offset;
};

Header.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 4
    length += 4
    var encoder_frame_id = new TextEncoder('utf8');
    var utf8array_frame_id = encoder_frame_id.encode(this.frame_id);
    length += 4;
    length += utf8array_frame_id.length;
    return length;
};

Header.prototype.echo = function() { return ""; };

Header.prototype.getType = function() { return "std_msgs/Header"; };

Header.prototype.getMD5 = function() { return "d33440e88be7b5b8255fc61ebbca06ad"; };

Header.prototype.getID = function() { return 0; };

Header.prototype.setID = function(id) { };

return function() { return new Header(); };

});
