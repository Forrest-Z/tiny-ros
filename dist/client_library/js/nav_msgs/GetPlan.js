(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GetPlanRequest=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/PoseStamped.js'></script>");

function GetPlanRequest() {
    this.__id__ = 0;
    this.start = geometry_msgs.PoseStamped();
    this.goal = geometry_msgs.PoseStamped();
    this.tolerance = 0.0;
};

GetPlanRequest.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    offset += this.start.serialize(buff, offset);
    offset += this.goal.serialize(buff, offset);
    var float32Array_tolerance = new Float32Array(1);
    var uInt8Float32Array_tolerance = new Uint8Array(float32Array_tolerance.buffer);
    float32Array_tolerance[0] = +this.tolerance;
    buff[offset + 0] = uInt8Float32Array_tolerance[0];
    buff[offset + 1] = uInt8Float32Array_tolerance[1];
    buff[offset + 2] = uInt8Float32Array_tolerance[2];
    buff[offset + 3] = uInt8Float32Array_tolerance[3];
    offset += 4;
    return offset;
};

GetPlanRequest.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    offset += this.start.deserialize(buff, offset);
    offset += this.goal.deserialize(buff, offset);
    var float32Array_tolerance = new Float32Array(1);
    var uInt8Float32Array_tolerance = new Uint8Array(float32Array_tolerance.buffer);
    uInt8Float32Array_tolerance[0] = buff[offset + 0];
    uInt8Float32Array_tolerance[1] = buff[offset + 1];
    uInt8Float32Array_tolerance[2] = buff[offset + 2];
    uInt8Float32Array_tolerance[3] = buff[offset + 3];
    this.tolerance = float32Array_tolerance[0];
    offset += 4;
    return offset;
};

GetPlanRequest.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += this.start.serializedLength();
    length += this.goal.serializedLength();
    length += 4
    return length;
};

GetPlanRequest.prototype.echo = function() { return ""; };

GetPlanRequest.prototype.getType = function() { return "nav_msgs/GetPlan"; };

GetPlanRequest.prototype.getMD5 = function() { return "557d5ea947f7761284cf7abef1cd7227"; };

GetPlanRequest.prototype.getID = function() { return this.__id__; };

GetPlanRequest.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetPlanRequest(); };

});


////////////////////////////////////////////////////////////////////////////////////////////////////


(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GetPlanResponse=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/nav_msgs/Path.js'></script>");

function GetPlanResponse() {
    this.__id__ = 0;
    this.plan = nav_msgs.Path();
};

GetPlanResponse.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset + 0] = ((+this.__id__) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.__id__) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.__id__) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.__id__) >> (8 * 3)) & 0xFF;
    offset += 4;
    offset += this.plan.serialize(buff, offset);
    return offset;
};

GetPlanResponse.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.__id__ = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.__id__ |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.__id__ |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.__id__ |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    offset += this.plan.deserialize(buff, offset);
    return offset;
};

GetPlanResponse.prototype.serializedLength = function() {
    var length = 0;
    length += 4
    length += this.plan.serializedLength();
    return length;
};

GetPlanResponse.prototype.echo = function() { return ""; };

GetPlanResponse.prototype.getType = function() { return "nav_msgs/GetPlan"; };

GetPlanResponse.prototype.getMD5 = function() { return "67c62b8c931eabfe35c88aed4b8f1258"; };

GetPlanResponse.prototype.getID = function() { return this.__id__; };

GetPlanResponse.prototype.setID = function(id) { this.__id__ = id; };

return function() { return new GetPlanResponse(); };

});
