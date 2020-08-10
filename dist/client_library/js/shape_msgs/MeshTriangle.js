(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.shape_msgs==="undefined"){g.shape_msgs={};}g.shape_msgs.MeshTriangle=f();}})(function(){var define,module,exports;'use strict';

function MeshTriangle() {
    this.vertex_indices = [0,0,0];
};

MeshTriangle.prototype.serialize = function(buff, idx) {
    var offset = idx;
    for (var i = 0; i < 3; i++) {
        buff[offset + 0] = ((+this.vertex_indices[i]) >> (8 * 0)) & 0xFF;
        buff[offset + 1] = ((+this.vertex_indices[i]) >> (8 * 1)) & 0xFF;
        buff[offset + 2] = ((+this.vertex_indices[i]) >> (8 * 2)) & 0xFF;
        buff[offset + 3] = ((+this.vertex_indices[i]) >> (8 * 3)) & 0xFF;
        offset += 4;
    }
    return offset;
};

MeshTriangle.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    for (var i = 0; i < 3; i++) {
        this.vertex_indices[i] = +((buff[offset + 0] & 0xFF) << (8 * 0));
        this.vertex_indices[i] |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        this.vertex_indices[i] |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        this.vertex_indices[i] |= +((buff[offset + 3] & 0xFF) << (8 * 3));
        offset += 4;
    }
    return offset;
};

MeshTriangle.prototype.serializedLength = function() {
    var length = 0;
    for (var i = 0; i < 3; i++) {
        length += 4
    }
    return length;
};

MeshTriangle.prototype.echo = function() { return ""; };

MeshTriangle.prototype.getType = function() { return "shape_msgs/MeshTriangle"; };

MeshTriangle.prototype.getMD5 = function() { return "01020cfeb9ad7679dd18bbd7149962ba"; };

MeshTriangle.prototype.getID = function() { return 0; };

MeshTriangle.prototype.setID = function(id) { };

return function() { return new MeshTriangle(); };

});
