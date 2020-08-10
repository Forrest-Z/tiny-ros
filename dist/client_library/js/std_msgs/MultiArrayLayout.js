(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.MultiArrayLayout=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/MultiArrayDimension.js'></script>");

function MultiArrayLayout() {
    this.dim = new Array();
    this.data_offset = 0;
};

MultiArrayLayout.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var length_dim = this.dim.length;
    buff[offset + 0] = (length_dim >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_dim >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_dim >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_dim >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_dim; i++) {
        offset += this.dim[i].serialize(buff, offset);
    }
    buff[offset + 0] = ((+this.data_offset) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.data_offset) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.data_offset) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.data_offset) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

MultiArrayLayout.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_dim = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_dim |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_dim |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_dim |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.dim = new Array(length_dim);
    for (var i = 0; i < length_dim; i++) {
        this.dim[i] = std_msgs.MultiArrayDimension();
        offset += this.dim[i].deserialize(buff, offset);
    }
    this.data_offset = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.data_offset |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.data_offset |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.data_offset |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

MultiArrayLayout.prototype.serializedLength = function() {
    var length = 0;
    var length_dim = this.dim.length;
    length += 4;
    for (var i = 0; i < length_dim; i++) {
        length += this.dim[i].serializedLength();
    }
    length += 4
    return length;
};

MultiArrayLayout.prototype.echo = function() { return ""; };

MultiArrayLayout.prototype.getType = function() { return "std_msgs/MultiArrayLayout"; };

MultiArrayLayout.prototype.getMD5 = function() { return "f40f0b5b285a93ca167c98c1012a989a"; };

MultiArrayLayout.prototype.getID = function() { return 0; };

MultiArrayLayout.prototype.setID = function(id) { };

return function() { return new MultiArrayLayout(); };

});
