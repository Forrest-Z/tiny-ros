(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.actionlib_msgs==="undefined"){g.actionlib_msgs={};}g.actionlib_msgs.GoalStatus=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/actionlib_msgs/GoalID.js'></script>");

function GoalStatus() {
    this.goal_id = actionlib_msgs.GoalID();
    this.status = 0;
    this.text = "";

    // ENUM{
    this.PENDING =  0   ;
    this.ACTIVE =  1   ;
    this.PREEMPTED =  2   ;
    this.SUCCEEDED =  3   ;
    this.ABORTED =  4   ;
    this.REJECTED =  5   ;
    this.PREEMPTING =  6   ;
    this.RECALLING =  7   ;
    this.RECALLED =  8   ;
    this.LOST =  9   ;
    // }ENUM
};

GoalStatus.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.goal_id.serialize(buff, offset);
    buff[offset + 0] = ((+this.status) >> (8 * 0)) & 0xFF;
    offset += 1;
    var encoder_text = new TextEncoder('utf8');
    var utf8array_text = encoder_text.encode(this.text);
    buff[offset + 0] = (utf8array_text.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_text.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_text.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_text.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_text.length; i++) {
        buff[offset + i] = utf8array_text[i];
    }
    offset += utf8array_text.length;
    return offset;
};

GoalStatus.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.goal_id.deserialize(buff, offset);
    this.status = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    var length_text = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_text |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_text |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_text |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_text = new TextDecoder('utf8');
    this.text = decoder_text.decode(buff.slice(offset, offset + length_text));
    offset += length_text;
    return offset;
};

GoalStatus.prototype.serializedLength = function() {
    var length = 0;
    length += this.goal_id.serializedLength();
    length += 1
    var encoder_text = new TextEncoder('utf8');
    var utf8array_text = encoder_text.encode(this.text);
    length += 4;
    length += utf8array_text.length;
    return length;
};

GoalStatus.prototype.echo = function() { return ""; };

GoalStatus.prototype.getType = function() { return "actionlib_msgs/GoalStatus"; };

GoalStatus.prototype.getMD5 = function() { return "086be35ea957e692de83fc3477e4ef0b"; };

GoalStatus.prototype.getID = function() { return 0; };

GoalStatus.prototype.setID = function(id) { };

return function() { return new GoalStatus(); };

});
