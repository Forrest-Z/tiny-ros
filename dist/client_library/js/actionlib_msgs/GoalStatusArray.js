(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.actionlib_msgs==="undefined"){g.actionlib_msgs={};}g.actionlib_msgs.GoalStatusArray=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/actionlib_msgs/GoalStatus.js'></script>");

function GoalStatusArray() {
    this.header = std_msgs.Header();
    this.status_list = new Array();
};

GoalStatusArray.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var length_status_list = this.status_list.length;
    buff[offset + 0] = (length_status_list >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_status_list >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_status_list >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_status_list >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_status_list; i++) {
        offset += this.status_list[i].serialize(buff, offset);
    }
    return offset;
};

GoalStatusArray.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_status_list = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_status_list |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_status_list |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_status_list |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.status_list = new Array(length_status_list);
    for (var i = 0; i < length_status_list; i++) {
        this.status_list[i] = actionlib_msgs.GoalStatus();
        offset += this.status_list[i].deserialize(buff, offset);
    }
    return offset;
};

GoalStatusArray.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var length_status_list = this.status_list.length;
    length += 4;
    for (var i = 0; i < length_status_list; i++) {
        length += this.status_list[i].serializedLength();
    }
    return length;
};

GoalStatusArray.prototype.echo = function() { return ""; };

GoalStatusArray.prototype.getType = function() { return "actionlib_msgs/GoalStatusArray"; };

GoalStatusArray.prototype.getMD5 = function() { return "53f6501f7c14f5f3963638de4bbe3a71"; };

GoalStatusArray.prototype.getID = function() { return 0; };

GoalStatusArray.prototype.setID = function(id) { };

return function() { return new GoalStatusArray(); };

});
