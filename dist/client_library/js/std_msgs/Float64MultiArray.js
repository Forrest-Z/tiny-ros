(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.std_msgs==="undefined"){g.std_msgs={};}g.std_msgs.Float64MultiArray=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/MultiArrayLayout.js'></script>");

function Float64MultiArray() {
    this.layout = std_msgs.MultiArrayLayout();
    this.data = new Array();
};

Float64MultiArray.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.layout.serialize(buff, offset);
    var length_data = this.data.length;
    buff[offset + 0] = (length_data >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_data >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_data >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_data >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_data; i++) {
        var float64Array_datai = new Float64Array(1);
        var uInt8Float64Array_datai = new Uint8Array(float64Array_datai.buffer);
        float64Array_datai[0] = +this.data[i];
        buff[offset + 0] = uInt8Float64Array_datai[0];
        buff[offset + 1] = uInt8Float64Array_datai[1];
        buff[offset + 2] = uInt8Float64Array_datai[2];
        buff[offset + 3] = uInt8Float64Array_datai[3];
        buff[offset + 4] = uInt8Float64Array_datai[4];
        buff[offset + 5] = uInt8Float64Array_datai[5];
        buff[offset + 6] = uInt8Float64Array_datai[6];
        buff[offset + 7] = uInt8Float64Array_datai[7];
        offset += 8;
    }
    return offset;
};

Float64MultiArray.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.layout.deserialize(buff, offset);
    var length_data = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_data |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_data |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_data |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.data = new Array(length_data);
    for (var i = 0; i < length_data; i++) {
        var float64Array_datai = new Float64Array(1);
        var uInt8Float64Array_datai = new Uint8Array(float64Array_datai.buffer);
        uInt8Float64Array_datai[0] = buff[offset + 0];
        uInt8Float64Array_datai[1] = buff[offset + 1];
        uInt8Float64Array_datai[2] = buff[offset + 2];
        uInt8Float64Array_datai[3] = buff[offset + 3];
        uInt8Float64Array_datai[4] = buff[offset + 4];
        uInt8Float64Array_datai[5] = buff[offset + 5];
        uInt8Float64Array_datai[6] = buff[offset + 6];
        uInt8Float64Array_datai[7] = buff[offset + 7];
        this.data[i] = float64Array_datai[0];
        offset += 8;
    }
    return offset;
};

Float64MultiArray.prototype.serializedLength = function() {
    var length = 0;
    length += this.layout.serializedLength();
    var length_data = this.data.length;
    length += 4;
    for (var i = 0; i < length_data; i++) {
        length += 8
    }
    return length;
};

Float64MultiArray.prototype.echo = function() { return ""; };

Float64MultiArray.prototype.getType = function() { return "std_msgs/Float64MultiArray"; };

Float64MultiArray.prototype.getMD5 = function() { return "e3061da26924f3790a70f9dbf06fc1a5"; };

Float64MultiArray.prototype.getID = function() { return 0; };

Float64MultiArray.prototype.setID = function(id) { };

return function() { return new Float64MultiArray(); };

});
