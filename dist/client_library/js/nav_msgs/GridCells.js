(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.nav_msgs==="undefined"){g.nav_msgs={};}g.nav_msgs.GridCells=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");
document.write("<script language=javascript src='"+_ROOT_PATH_+"/geometry_msgs/Point.js'></script>");

function GridCells() {
    this.header = std_msgs.Header();
    this.cell_width = 0.0;
    this.cell_height = 0.0;
    this.cells = new Array();
};

GridCells.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var float32Array_cell_width = new Float32Array(1);
    var uInt8Float32Array_cell_width = new Uint8Array(float32Array_cell_width.buffer);
    float32Array_cell_width[0] = +this.cell_width;
    buff[offset + 0] = uInt8Float32Array_cell_width[0];
    buff[offset + 1] = uInt8Float32Array_cell_width[1];
    buff[offset + 2] = uInt8Float32Array_cell_width[2];
    buff[offset + 3] = uInt8Float32Array_cell_width[3];
    offset += 4;
    var float32Array_cell_height = new Float32Array(1);
    var uInt8Float32Array_cell_height = new Uint8Array(float32Array_cell_height.buffer);
    float32Array_cell_height[0] = +this.cell_height;
    buff[offset + 0] = uInt8Float32Array_cell_height[0];
    buff[offset + 1] = uInt8Float32Array_cell_height[1];
    buff[offset + 2] = uInt8Float32Array_cell_height[2];
    buff[offset + 3] = uInt8Float32Array_cell_height[3];
    offset += 4;
    var length_cells = this.cells.length;
    buff[offset + 0] = (length_cells >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_cells >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_cells >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_cells >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_cells; i++) {
        offset += this.cells[i].serialize(buff, offset);
    }
    return offset;
};

GridCells.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var float32Array_cell_width = new Float32Array(1);
    var uInt8Float32Array_cell_width = new Uint8Array(float32Array_cell_width.buffer);
    uInt8Float32Array_cell_width[0] = buff[offset + 0];
    uInt8Float32Array_cell_width[1] = buff[offset + 1];
    uInt8Float32Array_cell_width[2] = buff[offset + 2];
    uInt8Float32Array_cell_width[3] = buff[offset + 3];
    this.cell_width = float32Array_cell_width[0];
    offset += 4;
    var float32Array_cell_height = new Float32Array(1);
    var uInt8Float32Array_cell_height = new Uint8Array(float32Array_cell_height.buffer);
    uInt8Float32Array_cell_height[0] = buff[offset + 0];
    uInt8Float32Array_cell_height[1] = buff[offset + 1];
    uInt8Float32Array_cell_height[2] = buff[offset + 2];
    uInt8Float32Array_cell_height[3] = buff[offset + 3];
    this.cell_height = float32Array_cell_height[0];
    offset += 4;
    var length_cells = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_cells |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_cells |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_cells |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.cells = new Array(length_cells);
    for (var i = 0; i < length_cells; i++) {
        this.cells[i] = geometry_msgs.Point();
        offset += this.cells[i].deserialize(buff, offset);
    }
    return offset;
};

GridCells.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 4
    length += 4
    var length_cells = this.cells.length;
    length += 4;
    for (var i = 0; i < length_cells; i++) {
        length += this.cells[i].serializedLength();
    }
    return length;
};

GridCells.prototype.echo = function() { return ""; };

GridCells.prototype.getType = function() { return "nav_msgs/GridCells"; };

GridCells.prototype.getMD5 = function() { return "13ce9063aaf922c39d3a2207d3926427"; };

GridCells.prototype.getID = function() { return 0; };

GridCells.prototype.setID = function(id) { };

return function() { return new GridCells(); };

});
