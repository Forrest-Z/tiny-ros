(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Int16MultiArray=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/MultiArrayLayout.js'></script>");

function Int16MultiArray() {
    this.layout = std_msgs.MultiArrayLayout();
    this.data = new Array();
};

Int16MultiArray.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.layout.serialize(buff, offset);
    var length_data = this.data.length;
    buff[offset + 0] = (length_data >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_data >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_data >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_data >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_data; i++) {
        buff[offset + 0] = ((+this.data[i]) >> (8 * 0)) & 0xFF;
        buff[offset + 1] = ((+this.data[i]) >> (8 * 1)) & 0xFF;
        offset += 2;
    }
    return offset;
};

Int16MultiArray.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.layout.deserialize(buff, offset);
    var length_data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_data |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_data |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_data |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.data = new Array(length_data);
    for (var i = 0; i < length_data; i++) {
        this.data[i] = +((buff[offset + 0] & 0xFF) << (8 * 0));
        this.data[i] |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        offset += 2;
    }
    return offset;
};

Int16MultiArray.prototype.serializedLength = function() {
    var length = 0;
    length += this.layout.serializedLength();
    var length_data = this.data.length;
    length += 4;
    for (var i = 0; i < length_data; i++) {
        length += 2
    }
    return length;
};

Int16MultiArray.prototype.echo = function() { return ""; };

Int16MultiArray.prototype.getType = function() { return "std_msgs/Int16MultiArray"; };

Int16MultiArray.prototype.getMD5 = function() { return "5344ae4ffc5e4576421880b487e96c93"; };

Int16MultiArray.prototype.getID = function() { return 0; };

Int16MultiArray.prototype.setID = function(id) { };

return function() { return new Int16MultiArray(); };

});
