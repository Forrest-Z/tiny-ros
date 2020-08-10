(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.ContactsState=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/gazebo_msgs/ContactState.js'></script>");

function ContactsState() {
    this.header = std_msgs.Header();
    this.states = new Array();
};

ContactsState.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var length_states = this.states.length;
    buff[offset + 0] = (length_states >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_states >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_states >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_states >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_states; i++) {
        offset += this.states[i].serialize(buff, offset);
    }
    return offset;
};

ContactsState.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_states = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_states |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_states |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_states |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.states = new Array(length_states);
    for (var i = 0; i < length_states; i++) {
        this.states[i] = gazebo_msgs.ContactState();
        offset += this.states[i].deserialize(buff, offset);
    }
    return offset;
};

ContactsState.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var length_states = this.states.length;
    length += 4;
    for (var i = 0; i < length_states; i++) {
        length += this.states[i].serializedLength();
    }
    return length;
};

ContactsState.prototype.echo = function() { return ""; };

ContactsState.prototype.getType = function() { return "gazebo_msgs/ContactsState"; };

ContactsState.prototype.getMD5 = function() { return "d19cd2a086cbd43da4252eb8d5cc64f5"; };

ContactsState.prototype.getID = function() { return 0; };

ContactsState.prototype.setID = function(id) { };

return function() { return new ContactsState(); };

});
