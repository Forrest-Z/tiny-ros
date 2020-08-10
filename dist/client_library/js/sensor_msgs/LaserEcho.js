(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.LaserEcho=f();}})(function(){var define,module,exports;'use strict';

function LaserEcho() {
    this.echoes = new Array();
};

LaserEcho.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var length_echoes = this.echoes.length;
    buff[offset + 0] = (length_echoes >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_echoes >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_echoes >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_echoes >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_echoes; i++) {
        var float32Array_echoesi = new Float32Array(1);
        var uInt8Float32Array_echoesi = new Uint8Array(float32Array_echoesi.buffer);
        float32Array_echoesi[0] = +this.echoes[i];
        buff[offset + 0] = uInt8Float32Array_echoesi[0];
        buff[offset + 1] = uInt8Float32Array_echoesi[1];
        buff[offset + 2] = uInt8Float32Array_echoesi[2];
        buff[offset + 3] = uInt8Float32Array_echoesi[3];
        offset += 4;
    }
    return offset;
};

LaserEcho.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_echoes = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_echoes |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_echoes |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_echoes |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.echoes = new Array(length_echoes);
    for (var i = 0; i < length_echoes; i++) {
        var float32Array_echoesi = new Float32Array(1);
        var uInt8Float32Array_echoesi = new Uint8Array(float32Array_echoesi.buffer);
        uInt8Float32Array_echoesi[0] = buff[offset + 0];
        uInt8Float32Array_echoesi[1] = buff[offset + 1];
        uInt8Float32Array_echoesi[2] = buff[offset + 2];
        uInt8Float32Array_echoesi[3] = buff[offset + 3];
        this.echoes[i] = float32Array_echoesi[0];
        offset += 4;
    }
    return offset;
};

LaserEcho.prototype.serializedLength = function() {
    var length = 0;
    var length_echoes = this.echoes.length;
    length += 4;
    for (var i = 0; i < length_echoes; i++) {
        length += 4
    }
    return length;
};

LaserEcho.prototype.echo = function() { return ""; };

LaserEcho.prototype.getType = function() { return "sensor_msgs/LaserEcho"; };

LaserEcho.prototype.getMD5 = function() { return "a8537b388573845b3240b44db5bc4905"; };

LaserEcho.prototype.getID = function() { return 0; };

LaserEcho.prototype.setID = function(id) { };

return function() { return new LaserEcho(); };

});
