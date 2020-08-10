(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.diagnostic_msgs==="undefined"){g.diagnostic_msgs={};}g.diagnostic_msgs.DiagnosticArray=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/diagnostic_msgs/DiagnosticStatus.js'></script>");

function DiagnosticArray() {
    this.header = std_msgs.Header();
    this.status = new Array();
};

DiagnosticArray.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var length_status = this.status.length;
    buff[offset + 0] = (length_status >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_status >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_status >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_status >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_status; i++) {
        offset += this.status[i].serialize(buff, offset);
    }
    return offset;
};

DiagnosticArray.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_status = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_status |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_status |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_status |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.status = new Array(length_status);
    for (var i = 0; i < length_status; i++) {
        this.status[i] = diagnostic_msgs.DiagnosticStatus();
        offset += this.status[i].deserialize(buff, offset);
    }
    return offset;
};

DiagnosticArray.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var length_status = this.status.length;
    length += 4;
    for (var i = 0; i < length_status; i++) {
        length += this.status[i].serializedLength();
    }
    return length;
};

DiagnosticArray.prototype.echo = function() { return ""; };

DiagnosticArray.prototype.getType = function() { return "diagnostic_msgs/DiagnosticArray"; };

DiagnosticArray.prototype.getMD5 = function() { return "79a87210f85eb6afbd600eb2ba49dd85"; };

DiagnosticArray.prototype.getID = function() { return 0; };

DiagnosticArray.prototype.setID = function(id) { };

return function() { return new DiagnosticArray(); };

});
