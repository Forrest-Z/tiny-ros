(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.UInt16MultiArray=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/MultiArrayLayout.js'></script>");

function UInt16MultiArray() {
    this.layout = std_msgs.MultiArrayLayout();
    this.data = new Array();
};

UInt16MultiArray.prototype.serialize = function(buff, idx) {
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

UInt16MultiArray.prototype.deserialize = function(buff, idx) {
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

UInt16MultiArray.prototype.serializedLength = function() {
    var length = 0;
    length += this.layout.serializedLength();
    var length_data = this.data.length;
    length += 4;
    for (var i = 0; i < length_data; i++) {
        length += 2
    }
    return length;
};

UInt16MultiArray.prototype.echo = function() { return ""; };

UInt16MultiArray.prototype.getType = function() { return "std_msgs/UInt16MultiArray"; };

UInt16MultiArray.prototype.getMD5 = function() { return "a27fbcc6cf43df7ac433688c26323c7b"; };

UInt16MultiArray.prototype.getID = function() { return 0; };

UInt16MultiArray.prototype.setID = function(id) { };

return function() { return new UInt16MultiArray(); };

});
