(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.JoyFeedbackArray=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/sensor_msgs/JoyFeedback.js'></script>");

function JoyFeedbackArray() {
    this.array = new Array();
};

JoyFeedbackArray.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var length_array = this.array.length;
    buff[offset + 0] = (length_array >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_array >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_array >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_array >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_array; i++) {
        offset += this.array[i].serialize(buff, offset);
    }
    return offset;
};

JoyFeedbackArray.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_array = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_array |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_array |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_array |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.array = new Array(length_array);
    for (var i = 0; i < length_array; i++) {
        this.array[i] = sensor_msgs.JoyFeedback();
        offset += this.array[i].deserialize(buff, offset);
    }
    return offset;
};

JoyFeedbackArray.prototype.serializedLength = function() {
    var length = 0;
    var length_array = this.array.length;
    length += 4;
    for (var i = 0; i < length_array; i++) {
        length += this.array[i].serializedLength();
    }
    return length;
};

JoyFeedbackArray.prototype.echo = function() { return ""; };

JoyFeedbackArray.prototype.getType = function() { return "sensor_msgs/JoyFeedbackArray"; };

JoyFeedbackArray.prototype.getMD5 = function() { return "45361e458d526d5670706a9f083819b6"; };

JoyFeedbackArray.prototype.getID = function() { return 0; };

JoyFeedbackArray.prototype.setID = function(id) { };

return function() { return new JoyFeedbackArray(); };

});
