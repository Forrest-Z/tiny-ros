(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.Joy=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function Joy() {
    this.header = std_msgs.Header();
    this.axes = new Array();
    this.buttons = new Array();
};

Joy.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var length_axes = this.axes.length;
    buff[offset + 0] = (length_axes >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_axes >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_axes >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_axes >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_axes; i++) {
        var float32Array_axesi = new Float32Array(1);
        var uInt8Float32Array_axesi = new Uint8Array(float32Array_axesi.buffer);
        float32Array_axesi[0] = +this.axes[i];
        buff[offset + 0] = uInt8Float32Array_axesi[0];
        buff[offset + 1] = uInt8Float32Array_axesi[1];
        buff[offset + 2] = uInt8Float32Array_axesi[2];
        buff[offset + 3] = uInt8Float32Array_axesi[3];
        offset += 4;
    }
    var length_buttons = this.buttons.length;
    buff[offset + 0] = (length_buttons >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_buttons >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_buttons >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_buttons >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_buttons; i++) {
        buff[offset + 0] = ((+this.buttons[i]) >> (8 * 0)) & 0xFF;
        buff[offset + 1] = ((+this.buttons[i]) >> (8 * 1)) & 0xFF;
        buff[offset + 2] = ((+this.buttons[i]) >> (8 * 2)) & 0xFF;
        buff[offset + 3] = ((+this.buttons[i]) >> (8 * 3)) & 0xFF;
        offset += 4;
    }
    return offset;
};

Joy.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_axes = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_axes |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_axes |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_axes |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.axes = new Array(length_axes);
    for (var i = 0; i < length_axes; i++) {
        var float32Array_axesi = new Float32Array(1);
        var uInt8Float32Array_axesi = new Uint8Array(float32Array_axesi.buffer);
        uInt8Float32Array_axesi[0] = buff[offset + 0];
        uInt8Float32Array_axesi[1] = buff[offset + 1];
        uInt8Float32Array_axesi[2] = buff[offset + 2];
        uInt8Float32Array_axesi[3] = buff[offset + 3];
        this.axes[i] = float32Array_axesi[0];
        offset += 4;
    }
    var length_buttons = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_buttons |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_buttons |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_buttons |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.buttons = new Array(length_buttons);
    for (var i = 0; i < length_buttons; i++) {
        this.buttons[i] = +((buff[offset + 0] & 0xFF) << (8 * 0));
        this.buttons[i] |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        this.buttons[i] |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        this.buttons[i] |= +((buff[offset + 3] & 0xFF) << (8 * 3));
        offset += 4;
    }
    return offset;
};

Joy.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var length_axes = this.axes.length;
    length += 4;
    for (var i = 0; i < length_axes; i++) {
        length += 4
    }
    var length_buttons = this.buttons.length;
    length += 4;
    for (var i = 0; i < length_buttons; i++) {
        length += 4
    }
    return length;
};

Joy.prototype.echo = function() { return ""; };

Joy.prototype.getType = function() { return "sensor_msgs/Joy"; };

Joy.prototype.getMD5 = function() { return "da3438323dbe92a4d6e4658e06bf8da1"; };

Joy.prototype.getID = function() { return 0; };

Joy.prototype.setID = function(id) { };

return function() { return new Joy(); };

});
