(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.tf2_msgs==="undefined"){g.tf2_msgs={};}g.tf2_msgs.TFMessage=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/TransformStamped.js'></script>");

function TFMessage() {
    this.transforms = new Array();
};

TFMessage.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var length_transforms = this.transforms.length;
    buff[offset + 0] = (length_transforms >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_transforms >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_transforms >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_transforms >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_transforms; i++) {
        offset += this.transforms[i].serialize(buff, offset);
    }
    return offset;
};

TFMessage.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_transforms = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_transforms |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_transforms |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_transforms |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.transforms = new Array(length_transforms);
    for (var i = 0; i < length_transforms; i++) {
        this.transforms[i] = geometry_msgs.TransformStamped();
        offset += this.transforms[i].deserialize(buff, offset);
    }
    return offset;
};

TFMessage.prototype.serializedLength = function() {
    var length = 0;
    var length_transforms = this.transforms.length;
    length += 4;
    for (var i = 0; i < length_transforms; i++) {
        length += this.transforms[i].serializedLength();
    }
    return length;
};

TFMessage.prototype.echo = function() { return ""; };

TFMessage.prototype.getType = function() { return "tf2_msgs/TFMessage"; };

TFMessage.prototype.getMD5 = function() { return "cb93cfe6a141f8d8af7cc34997ec99fe"; };

TFMessage.prototype.getID = function() { return 0; };

TFMessage.prototype.setID = function(id) { };

return function() { return new TFMessage(); };

});
