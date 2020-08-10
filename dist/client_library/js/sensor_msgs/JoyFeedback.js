(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.JoyFeedback=f();}})(function(){var define,module,exports;'use strict';

function JoyFeedback() {
    this.type = 0;
    this.id = 0;
    this.intensity = 0.0;

    // ENUM{
    this.TYPE_LED =  0;
    this.TYPE_RUMBLE =  1;
    this.TYPE_BUZZER =  2;
    // }ENUM
};

JoyFeedback.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.type) >> (8 * 0)) & 0xFF;
    offset += 1;
    buff[offset + 0] = ((+this.id) >> (8 * 0)) & 0xFF;
    offset += 1;
    var float32Array_intensity = new Float32Array(1);
    var uInt8Float32Array_intensity = new Uint8Array(float32Array_intensity.buffer);
    float32Array_intensity[0] = +this.intensity;
    buff[offset + 0] = uInt8Float32Array_intensity[0];
    buff[offset + 1] = uInt8Float32Array_intensity[1];
    buff[offset + 2] = uInt8Float32Array_intensity[2];
    buff[offset + 3] = uInt8Float32Array_intensity[3];
    offset += 4;
    return offset;
};

JoyFeedback.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.type = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    this.id = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    var float32Array_intensity = new Float32Array(1);
    var uInt8Float32Array_intensity = new Uint8Array(float32Array_intensity.buffer);
    uInt8Float32Array_intensity[0] = buff[offset + 0];
    uInt8Float32Array_intensity[1] = buff[offset + 1];
    uInt8Float32Array_intensity[2] = buff[offset + 2];
    uInt8Float32Array_intensity[3] = buff[offset + 3];
    this.intensity = float32Array_intensity[0];
    offset += 4;
    return offset;
};

JoyFeedback.prototype.serializedLength = function() {
    var length = 0;
    length += 1
    length += 1
    length += 4
    return length;
};

JoyFeedback.prototype.echo = function() { return ""; };

JoyFeedback.prototype.getType = function() { return "sensor_msgs/JoyFeedback"; };

JoyFeedback.prototype.getMD5 = function() { return "206b65e86c8b195f816ccbe40b3568a2"; };

JoyFeedback.prototype.getID = function() { return 0; };

JoyFeedback.prototype.setID = function(id) { };

return function() { return new JoyFeedback(); };

});
