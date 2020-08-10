(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.actionlib_msgs==="undefined"){g.actionlib_msgs={};}g.actionlib_msgs.GoalID=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/tinyros/Time.js'></script>");

function GoalID() {
    this.stamp = tinyros.Time();
    this.id = "";
};

GoalID.prototype.serialize = function(buff, idx) {
    var offset = idx;
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
    var encoder_id = new TextEncoder('utf8');
    var utf8array_id = encoder_id.encode(this.id);
    buff[offset + 0] = (utf8array_id.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_id.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_id.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_id.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_id.length; i++) {
        buff[offset + i] = utf8array_id[i];
    }
    offset += utf8array_id.length;
    return offset;
};

GoalID.prototype.deserialize = function(buff, idx) {
    var offset = idx;
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
    var length_id = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_id |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_id |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_id |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_id = new TextDecoder('utf8');
    this.id = decoder_id.decode(buff.slice(offset, offset + length_id));
    offset += length_id;
    return offset;
};

GoalID.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += 4
    var encoder_id = new TextEncoder('utf8');
    var utf8array_id = encoder_id.encode(this.id);
    length += 4;
    length += utf8array_id.length;
    return length;
};

GoalID.prototype.echo = function() { return ""; };

GoalID.prototype.getType = function() { return "actionlib_msgs/GoalID"; };

GoalID.prototype.getMD5 = function() { return "a6cee90e5a185f4cb050de49bc4fa1f4"; };

GoalID.prototype.getID = function() { return 0; };

GoalID.prototype.setID = function(id) { };

return function() { return new GoalID(); };

});
