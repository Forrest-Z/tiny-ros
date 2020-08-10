(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.shape_msgs==="undefined"){g.shape_msgs={};}g.shape_msgs.Mesh=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/shape_msgs/MeshTriangle.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Point.js'></script>");

function Mesh() {
    this.triangles = new Array();
    this.vertices = new Array();
};

Mesh.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var length_triangles = this.triangles.length;
    buff[offset + 0] = (length_triangles >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_triangles >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_triangles >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_triangles >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_triangles; i++) {
        offset += this.triangles[i].serialize(buff, offset);
    }
    var length_vertices = this.vertices.length;
    buff[offset + 0] = (length_vertices >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_vertices >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_vertices >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_vertices >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_vertices; i++) {
        offset += this.vertices[i].serialize(buff, offset);
    }
    return offset;
};

Mesh.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_triangles = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_triangles |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_triangles |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_triangles |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.triangles = new Array(length_triangles);
    for (var i = 0; i < length_triangles; i++) {
        this.triangles[i] = shape_msgs.MeshTriangle();
        offset += this.triangles[i].deserialize(buff, offset);
    }
    var length_vertices = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_vertices |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_vertices |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_vertices |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.vertices = new Array(length_vertices);
    for (var i = 0; i < length_vertices; i++) {
        this.vertices[i] = geometry_msgs.Point();
        offset += this.vertices[i].deserialize(buff, offset);
    }
    return offset;
};

Mesh.prototype.serializedLength = function() {
    var length = 0;
    var length_triangles = this.triangles.length;
    length += 4;
    for (var i = 0; i < length_triangles; i++) {
        length += this.triangles[i].serializedLength();
    }
    var length_vertices = this.vertices.length;
    length += 4;
    for (var i = 0; i < length_vertices; i++) {
        length += this.vertices[i].serializedLength();
    }
    return length;
};

Mesh.prototype.echo = function() { return ""; };

Mesh.prototype.getType = function() { return "shape_msgs/Mesh"; };

Mesh.prototype.getMD5 = function() { return "1579670b316f622b47d6700cd4f7e18d"; };

Mesh.prototype.getID = function() { return 0; };

Mesh.prototype.setID = function(id) { };

return function() { return new Mesh(); };

});
