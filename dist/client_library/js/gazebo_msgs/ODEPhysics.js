(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.ODEPhysics=f();}})(function(){var define,module,exports;'use strict';

function ODEPhysics() {
    this.auto_disable_bodies = false;
    this.sor_pgs_precon_iters = 0;
    this.sor_pgs_iters = 0;
    this.sor_pgs_w = 0.0;
    this.sor_pgs_rms_error_tol = 0.0;
    this.contact_surface_layer = 0.0;
    this.contact_max_correcting_vel = 0.0;
    this.cfm = 0.0;
    this.erp = 0.0;
    this.max_contacts = 0;
};

ODEPhysics.prototype.serialize = function(buff, idx) {
    var offset = idx;
    buff[offset] = this.auto_disable_bodies === false ? 0 : 1;
    offset += 1;
    buff[offset + 0] = ((+this.sor_pgs_precon_iters) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.sor_pgs_precon_iters) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.sor_pgs_precon_iters) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.sor_pgs_precon_iters) >> (8 * 3)) & 0xFF;
    offset += 4;
    buff[offset + 0] = ((+this.sor_pgs_iters) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.sor_pgs_iters) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.sor_pgs_iters) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.sor_pgs_iters) >> (8 * 3)) & 0xFF;
    offset += 4;
    var float64Array_sor_pgs_w = new Float64Array(1);
    var uInt8Float64Array_sor_pgs_w = new Uint8Array(float64Array_sor_pgs_w.buffer);
    float64Array_sor_pgs_w[0] = +this.sor_pgs_w;
    buff[offset + 0] = uInt8Float64Array_sor_pgs_w[0];
    buff[offset + 1] = uInt8Float64Array_sor_pgs_w[1];
    buff[offset + 2] = uInt8Float64Array_sor_pgs_w[2];
    buff[offset + 3] = uInt8Float64Array_sor_pgs_w[3];
    buff[offset + 4] = uInt8Float64Array_sor_pgs_w[4];
    buff[offset + 5] = uInt8Float64Array_sor_pgs_w[5];
    buff[offset + 6] = uInt8Float64Array_sor_pgs_w[6];
    buff[offset + 7] = uInt8Float64Array_sor_pgs_w[7];
    offset += 8;
    var float64Array_sor_pgs_rms_error_tol = new Float64Array(1);
    var uInt8Float64Array_sor_pgs_rms_error_tol = new Uint8Array(float64Array_sor_pgs_rms_error_tol.buffer);
    float64Array_sor_pgs_rms_error_tol[0] = +this.sor_pgs_rms_error_tol;
    buff[offset + 0] = uInt8Float64Array_sor_pgs_rms_error_tol[0];
    buff[offset + 1] = uInt8Float64Array_sor_pgs_rms_error_tol[1];
    buff[offset + 2] = uInt8Float64Array_sor_pgs_rms_error_tol[2];
    buff[offset + 3] = uInt8Float64Array_sor_pgs_rms_error_tol[3];
    buff[offset + 4] = uInt8Float64Array_sor_pgs_rms_error_tol[4];
    buff[offset + 5] = uInt8Float64Array_sor_pgs_rms_error_tol[5];
    buff[offset + 6] = uInt8Float64Array_sor_pgs_rms_error_tol[6];
    buff[offset + 7] = uInt8Float64Array_sor_pgs_rms_error_tol[7];
    offset += 8;
    var float64Array_contact_surface_layer = new Float64Array(1);
    var uInt8Float64Array_contact_surface_layer = new Uint8Array(float64Array_contact_surface_layer.buffer);
    float64Array_contact_surface_layer[0] = +this.contact_surface_layer;
    buff[offset + 0] = uInt8Float64Array_contact_surface_layer[0];
    buff[offset + 1] = uInt8Float64Array_contact_surface_layer[1];
    buff[offset + 2] = uInt8Float64Array_contact_surface_layer[2];
    buff[offset + 3] = uInt8Float64Array_contact_surface_layer[3];
    buff[offset + 4] = uInt8Float64Array_contact_surface_layer[4];
    buff[offset + 5] = uInt8Float64Array_contact_surface_layer[5];
    buff[offset + 6] = uInt8Float64Array_contact_surface_layer[6];
    buff[offset + 7] = uInt8Float64Array_contact_surface_layer[7];
    offset += 8;
    var float64Array_contact_max_correcting_vel = new Float64Array(1);
    var uInt8Float64Array_contact_max_correcting_vel = new Uint8Array(float64Array_contact_max_correcting_vel.buffer);
    float64Array_contact_max_correcting_vel[0] = +this.contact_max_correcting_vel;
    buff[offset + 0] = uInt8Float64Array_contact_max_correcting_vel[0];
    buff[offset + 1] = uInt8Float64Array_contact_max_correcting_vel[1];
    buff[offset + 2] = uInt8Float64Array_contact_max_correcting_vel[2];
    buff[offset + 3] = uInt8Float64Array_contact_max_correcting_vel[3];
    buff[offset + 4] = uInt8Float64Array_contact_max_correcting_vel[4];
    buff[offset + 5] = uInt8Float64Array_contact_max_correcting_vel[5];
    buff[offset + 6] = uInt8Float64Array_contact_max_correcting_vel[6];
    buff[offset + 7] = uInt8Float64Array_contact_max_correcting_vel[7];
    offset += 8;
    var float64Array_cfm = new Float64Array(1);
    var uInt8Float64Array_cfm = new Uint8Array(float64Array_cfm.buffer);
    float64Array_cfm[0] = +this.cfm;
    buff[offset + 0] = uInt8Float64Array_cfm[0];
    buff[offset + 1] = uInt8Float64Array_cfm[1];
    buff[offset + 2] = uInt8Float64Array_cfm[2];
    buff[offset + 3] = uInt8Float64Array_cfm[3];
    buff[offset + 4] = uInt8Float64Array_cfm[4];
    buff[offset + 5] = uInt8Float64Array_cfm[5];
    buff[offset + 6] = uInt8Float64Array_cfm[6];
    buff[offset + 7] = uInt8Float64Array_cfm[7];
    offset += 8;
    var float64Array_erp = new Float64Array(1);
    var uInt8Float64Array_erp = new Uint8Array(float64Array_erp.buffer);
    float64Array_erp[0] = +this.erp;
    buff[offset + 0] = uInt8Float64Array_erp[0];
    buff[offset + 1] = uInt8Float64Array_erp[1];
    buff[offset + 2] = uInt8Float64Array_erp[2];
    buff[offset + 3] = uInt8Float64Array_erp[3];
    buff[offset + 4] = uInt8Float64Array_erp[4];
    buff[offset + 5] = uInt8Float64Array_erp[5];
    buff[offset + 6] = uInt8Float64Array_erp[6];
    buff[offset + 7] = uInt8Float64Array_erp[7];
    offset += 8;
    buff[offset + 0] = ((+this.max_contacts) >> (8 * 0)) & 0xFF;
    buff[offset + 1] = ((+this.max_contacts) >> (8 * 1)) & 0xFF;
    buff[offset + 2] = ((+this.max_contacts) >> (8 * 2)) & 0xFF;
    buff[offset + 3] = ((+this.max_contacts) >> (8 * 3)) & 0xFF;
    offset += 4;
    return offset;
};

ODEPhysics.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    this.auto_disable_bodies = buff[offset] !== 0 ? true : false;
    offset += 1;
    this.sor_pgs_precon_iters = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.sor_pgs_precon_iters |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.sor_pgs_precon_iters |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.sor_pgs_precon_iters |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.sor_pgs_iters = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.sor_pgs_iters |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.sor_pgs_iters |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.sor_pgs_iters |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    var float64Array_sor_pgs_w = new Float64Array(1);
    var uInt8Float64Array_sor_pgs_w = new Uint8Array(float64Array_sor_pgs_w.buffer);
    uInt8Float64Array_sor_pgs_w[0] = buff[offset + 0];
    uInt8Float64Array_sor_pgs_w[1] = buff[offset + 1];
    uInt8Float64Array_sor_pgs_w[2] = buff[offset + 2];
    uInt8Float64Array_sor_pgs_w[3] = buff[offset + 3];
    uInt8Float64Array_sor_pgs_w[4] = buff[offset + 4];
    uInt8Float64Array_sor_pgs_w[5] = buff[offset + 5];
    uInt8Float64Array_sor_pgs_w[6] = buff[offset + 6];
    uInt8Float64Array_sor_pgs_w[7] = buff[offset + 7];
    this.sor_pgs_w = float64Array_sor_pgs_w[0];
    offset += 8;
    var float64Array_sor_pgs_rms_error_tol = new Float64Array(1);
    var uInt8Float64Array_sor_pgs_rms_error_tol = new Uint8Array(float64Array_sor_pgs_rms_error_tol.buffer);
    uInt8Float64Array_sor_pgs_rms_error_tol[0] = buff[offset + 0];
    uInt8Float64Array_sor_pgs_rms_error_tol[1] = buff[offset + 1];
    uInt8Float64Array_sor_pgs_rms_error_tol[2] = buff[offset + 2];
    uInt8Float64Array_sor_pgs_rms_error_tol[3] = buff[offset + 3];
    uInt8Float64Array_sor_pgs_rms_error_tol[4] = buff[offset + 4];
    uInt8Float64Array_sor_pgs_rms_error_tol[5] = buff[offset + 5];
    uInt8Float64Array_sor_pgs_rms_error_tol[6] = buff[offset + 6];
    uInt8Float64Array_sor_pgs_rms_error_tol[7] = buff[offset + 7];
    this.sor_pgs_rms_error_tol = float64Array_sor_pgs_rms_error_tol[0];
    offset += 8;
    var float64Array_contact_surface_layer = new Float64Array(1);
    var uInt8Float64Array_contact_surface_layer = new Uint8Array(float64Array_contact_surface_layer.buffer);
    uInt8Float64Array_contact_surface_layer[0] = buff[offset + 0];
    uInt8Float64Array_contact_surface_layer[1] = buff[offset + 1];
    uInt8Float64Array_contact_surface_layer[2] = buff[offset + 2];
    uInt8Float64Array_contact_surface_layer[3] = buff[offset + 3];
    uInt8Float64Array_contact_surface_layer[4] = buff[offset + 4];
    uInt8Float64Array_contact_surface_layer[5] = buff[offset + 5];
    uInt8Float64Array_contact_surface_layer[6] = buff[offset + 6];
    uInt8Float64Array_contact_surface_layer[7] = buff[offset + 7];
    this.contact_surface_layer = float64Array_contact_surface_layer[0];
    offset += 8;
    var float64Array_contact_max_correcting_vel = new Float64Array(1);
    var uInt8Float64Array_contact_max_correcting_vel = new Uint8Array(float64Array_contact_max_correcting_vel.buffer);
    uInt8Float64Array_contact_max_correcting_vel[0] = buff[offset + 0];
    uInt8Float64Array_contact_max_correcting_vel[1] = buff[offset + 1];
    uInt8Float64Array_contact_max_correcting_vel[2] = buff[offset + 2];
    uInt8Float64Array_contact_max_correcting_vel[3] = buff[offset + 3];
    uInt8Float64Array_contact_max_correcting_vel[4] = buff[offset + 4];
    uInt8Float64Array_contact_max_correcting_vel[5] = buff[offset + 5];
    uInt8Float64Array_contact_max_correcting_vel[6] = buff[offset + 6];
    uInt8Float64Array_contact_max_correcting_vel[7] = buff[offset + 7];
    this.contact_max_correcting_vel = float64Array_contact_max_correcting_vel[0];
    offset += 8;
    var float64Array_cfm = new Float64Array(1);
    var uInt8Float64Array_cfm = new Uint8Array(float64Array_cfm.buffer);
    uInt8Float64Array_cfm[0] = buff[offset + 0];
    uInt8Float64Array_cfm[1] = buff[offset + 1];
    uInt8Float64Array_cfm[2] = buff[offset + 2];
    uInt8Float64Array_cfm[3] = buff[offset + 3];
    uInt8Float64Array_cfm[4] = buff[offset + 4];
    uInt8Float64Array_cfm[5] = buff[offset + 5];
    uInt8Float64Array_cfm[6] = buff[offset + 6];
    uInt8Float64Array_cfm[7] = buff[offset + 7];
    this.cfm = float64Array_cfm[0];
    offset += 8;
    var float64Array_erp = new Float64Array(1);
    var uInt8Float64Array_erp = new Uint8Array(float64Array_erp.buffer);
    uInt8Float64Array_erp[0] = buff[offset + 0];
    uInt8Float64Array_erp[1] = buff[offset + 1];
    uInt8Float64Array_erp[2] = buff[offset + 2];
    uInt8Float64Array_erp[3] = buff[offset + 3];
    uInt8Float64Array_erp[4] = buff[offset + 4];
    uInt8Float64Array_erp[5] = buff[offset + 5];
    uInt8Float64Array_erp[6] = buff[offset + 6];
    uInt8Float64Array_erp[7] = buff[offset + 7];
    this.erp = float64Array_erp[0];
    offset += 8;
    this.max_contacts = +((buff[offset + 0] & 0xFF) << (8 * 0));
    this.max_contacts |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    this.max_contacts |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    this.max_contacts |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    return offset;
};

ODEPhysics.prototype.serializedLength = function() {
    var length = 0;
    length += 1
    length += 4
    length += 4
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 8
    length += 4
    return length;
};

ODEPhysics.prototype.echo = function() { return ""; };

ODEPhysics.prototype.getType = function() { return "gazebo_msgs/ODEPhysics"; };

ODEPhysics.prototype.getMD5 = function() { return "67a077e58362b50f63dc189c25d01418"; };

ODEPhysics.prototype.getID = function() { return 0; };

ODEPhysics.prototype.setID = function(id) { };

return function() { return new ODEPhysics(); };

});
