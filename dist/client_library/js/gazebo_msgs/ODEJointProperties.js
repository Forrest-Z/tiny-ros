(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.gazebo_msgs==="undefined"){g.gazebo_msgs={};}g.gazebo_msgs.ODEJointProperties=f();}})(function(){var define,module,exports;'use strict';

function ODEJointProperties() {
    this.damping = new Array();
    this.hiStop = new Array();
    this.loStop = new Array();
    this.erp = new Array();
    this.cfm = new Array();
    this.stop_erp = new Array();
    this.stop_cfm = new Array();
    this.fudge_factor = new Array();
    this.fmax = new Array();
    this.vel = new Array();
};

ODEJointProperties.prototype.serialize = function(buff, idx) {
    var offset = idx;
    var length_damping = this.damping.length;
    buff[offset + 0] = (length_damping >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_damping >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_damping >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_damping >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_damping; i++) {
        var float64Array_dampingi = new Float64Array(1);
        var uInt8Float64Array_dampingi = new Uint8Array(float64Array_dampingi.buffer);
        float64Array_dampingi[0] = +this.damping[i];
        buff[offset + 0] = uInt8Float64Array_dampingi[0];
        buff[offset + 1] = uInt8Float64Array_dampingi[1];
        buff[offset + 2] = uInt8Float64Array_dampingi[2];
        buff[offset + 3] = uInt8Float64Array_dampingi[3];
        buff[offset + 4] = uInt8Float64Array_dampingi[4];
        buff[offset + 5] = uInt8Float64Array_dampingi[5];
        buff[offset + 6] = uInt8Float64Array_dampingi[6];
        buff[offset + 7] = uInt8Float64Array_dampingi[7];
        offset += 8;
    }
    var length_hiStop = this.hiStop.length;
    buff[offset + 0] = (length_hiStop >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_hiStop >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_hiStop >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_hiStop >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_hiStop; i++) {
        var float64Array_hiStopi = new Float64Array(1);
        var uInt8Float64Array_hiStopi = new Uint8Array(float64Array_hiStopi.buffer);
        float64Array_hiStopi[0] = +this.hiStop[i];
        buff[offset + 0] = uInt8Float64Array_hiStopi[0];
        buff[offset + 1] = uInt8Float64Array_hiStopi[1];
        buff[offset + 2] = uInt8Float64Array_hiStopi[2];
        buff[offset + 3] = uInt8Float64Array_hiStopi[3];
        buff[offset + 4] = uInt8Float64Array_hiStopi[4];
        buff[offset + 5] = uInt8Float64Array_hiStopi[5];
        buff[offset + 6] = uInt8Float64Array_hiStopi[6];
        buff[offset + 7] = uInt8Float64Array_hiStopi[7];
        offset += 8;
    }
    var length_loStop = this.loStop.length;
    buff[offset + 0] = (length_loStop >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_loStop >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_loStop >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_loStop >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_loStop; i++) {
        var float64Array_loStopi = new Float64Array(1);
        var uInt8Float64Array_loStopi = new Uint8Array(float64Array_loStopi.buffer);
        float64Array_loStopi[0] = +this.loStop[i];
        buff[offset + 0] = uInt8Float64Array_loStopi[0];
        buff[offset + 1] = uInt8Float64Array_loStopi[1];
        buff[offset + 2] = uInt8Float64Array_loStopi[2];
        buff[offset + 3] = uInt8Float64Array_loStopi[3];
        buff[offset + 4] = uInt8Float64Array_loStopi[4];
        buff[offset + 5] = uInt8Float64Array_loStopi[5];
        buff[offset + 6] = uInt8Float64Array_loStopi[6];
        buff[offset + 7] = uInt8Float64Array_loStopi[7];
        offset += 8;
    }
    var length_erp = this.erp.length;
    buff[offset + 0] = (length_erp >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_erp >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_erp >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_erp >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_erp; i++) {
        var float64Array_erpi = new Float64Array(1);
        var uInt8Float64Array_erpi = new Uint8Array(float64Array_erpi.buffer);
        float64Array_erpi[0] = +this.erp[i];
        buff[offset + 0] = uInt8Float64Array_erpi[0];
        buff[offset + 1] = uInt8Float64Array_erpi[1];
        buff[offset + 2] = uInt8Float64Array_erpi[2];
        buff[offset + 3] = uInt8Float64Array_erpi[3];
        buff[offset + 4] = uInt8Float64Array_erpi[4];
        buff[offset + 5] = uInt8Float64Array_erpi[5];
        buff[offset + 6] = uInt8Float64Array_erpi[6];
        buff[offset + 7] = uInt8Float64Array_erpi[7];
        offset += 8;
    }
    var length_cfm = this.cfm.length;
    buff[offset + 0] = (length_cfm >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_cfm >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_cfm >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_cfm >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_cfm; i++) {
        var float64Array_cfmi = new Float64Array(1);
        var uInt8Float64Array_cfmi = new Uint8Array(float64Array_cfmi.buffer);
        float64Array_cfmi[0] = +this.cfm[i];
        buff[offset + 0] = uInt8Float64Array_cfmi[0];
        buff[offset + 1] = uInt8Float64Array_cfmi[1];
        buff[offset + 2] = uInt8Float64Array_cfmi[2];
        buff[offset + 3] = uInt8Float64Array_cfmi[3];
        buff[offset + 4] = uInt8Float64Array_cfmi[4];
        buff[offset + 5] = uInt8Float64Array_cfmi[5];
        buff[offset + 6] = uInt8Float64Array_cfmi[6];
        buff[offset + 7] = uInt8Float64Array_cfmi[7];
        offset += 8;
    }
    var length_stop_erp = this.stop_erp.length;
    buff[offset + 0] = (length_stop_erp >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_stop_erp >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_stop_erp >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_stop_erp >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_stop_erp; i++) {
        var float64Array_stop_erpi = new Float64Array(1);
        var uInt8Float64Array_stop_erpi = new Uint8Array(float64Array_stop_erpi.buffer);
        float64Array_stop_erpi[0] = +this.stop_erp[i];
        buff[offset + 0] = uInt8Float64Array_stop_erpi[0];
        buff[offset + 1] = uInt8Float64Array_stop_erpi[1];
        buff[offset + 2] = uInt8Float64Array_stop_erpi[2];
        buff[offset + 3] = uInt8Float64Array_stop_erpi[3];
        buff[offset + 4] = uInt8Float64Array_stop_erpi[4];
        buff[offset + 5] = uInt8Float64Array_stop_erpi[5];
        buff[offset + 6] = uInt8Float64Array_stop_erpi[6];
        buff[offset + 7] = uInt8Float64Array_stop_erpi[7];
        offset += 8;
    }
    var length_stop_cfm = this.stop_cfm.length;
    buff[offset + 0] = (length_stop_cfm >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_stop_cfm >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_stop_cfm >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_stop_cfm >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_stop_cfm; i++) {
        var float64Array_stop_cfmi = new Float64Array(1);
        var uInt8Float64Array_stop_cfmi = new Uint8Array(float64Array_stop_cfmi.buffer);
        float64Array_stop_cfmi[0] = +this.stop_cfm[i];
        buff[offset + 0] = uInt8Float64Array_stop_cfmi[0];
        buff[offset + 1] = uInt8Float64Array_stop_cfmi[1];
        buff[offset + 2] = uInt8Float64Array_stop_cfmi[2];
        buff[offset + 3] = uInt8Float64Array_stop_cfmi[3];
        buff[offset + 4] = uInt8Float64Array_stop_cfmi[4];
        buff[offset + 5] = uInt8Float64Array_stop_cfmi[5];
        buff[offset + 6] = uInt8Float64Array_stop_cfmi[6];
        buff[offset + 7] = uInt8Float64Array_stop_cfmi[7];
        offset += 8;
    }
    var length_fudge_factor = this.fudge_factor.length;
    buff[offset + 0] = (length_fudge_factor >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_fudge_factor >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_fudge_factor >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_fudge_factor >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_fudge_factor; i++) {
        var float64Array_fudge_factori = new Float64Array(1);
        var uInt8Float64Array_fudge_factori = new Uint8Array(float64Array_fudge_factori.buffer);
        float64Array_fudge_factori[0] = +this.fudge_factor[i];
        buff[offset + 0] = uInt8Float64Array_fudge_factori[0];
        buff[offset + 1] = uInt8Float64Array_fudge_factori[1];
        buff[offset + 2] = uInt8Float64Array_fudge_factori[2];
        buff[offset + 3] = uInt8Float64Array_fudge_factori[3];
        buff[offset + 4] = uInt8Float64Array_fudge_factori[4];
        buff[offset + 5] = uInt8Float64Array_fudge_factori[5];
        buff[offset + 6] = uInt8Float64Array_fudge_factori[6];
        buff[offset + 7] = uInt8Float64Array_fudge_factori[7];
        offset += 8;
    }
    var length_fmax = this.fmax.length;
    buff[offset + 0] = (length_fmax >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_fmax >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_fmax >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_fmax >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_fmax; i++) {
        var float64Array_fmaxi = new Float64Array(1);
        var uInt8Float64Array_fmaxi = new Uint8Array(float64Array_fmaxi.buffer);
        float64Array_fmaxi[0] = +this.fmax[i];
        buff[offset + 0] = uInt8Float64Array_fmaxi[0];
        buff[offset + 1] = uInt8Float64Array_fmaxi[1];
        buff[offset + 2] = uInt8Float64Array_fmaxi[2];
        buff[offset + 3] = uInt8Float64Array_fmaxi[3];
        buff[offset + 4] = uInt8Float64Array_fmaxi[4];
        buff[offset + 5] = uInt8Float64Array_fmaxi[5];
        buff[offset + 6] = uInt8Float64Array_fmaxi[6];
        buff[offset + 7] = uInt8Float64Array_fmaxi[7];
        offset += 8;
    }
    var length_vel = this.vel.length;
    buff[offset + 0] = (length_vel >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_vel >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_vel >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_vel >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_vel; i++) {
        var float64Array_veli = new Float64Array(1);
        var uInt8Float64Array_veli = new Uint8Array(float64Array_veli.buffer);
        float64Array_veli[0] = +this.vel[i];
        buff[offset + 0] = uInt8Float64Array_veli[0];
        buff[offset + 1] = uInt8Float64Array_veli[1];
        buff[offset + 2] = uInt8Float64Array_veli[2];
        buff[offset + 3] = uInt8Float64Array_veli[3];
        buff[offset + 4] = uInt8Float64Array_veli[4];
        buff[offset + 5] = uInt8Float64Array_veli[5];
        buff[offset + 6] = uInt8Float64Array_veli[6];
        buff[offset + 7] = uInt8Float64Array_veli[7];
        offset += 8;
    }
    return offset;
};

ODEJointProperties.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    var length_damping = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_damping |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_damping |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_damping |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.damping = new Array(length_damping);
    for (var i = 0; i < length_damping; i++) {
        var float64Array_dampingi = new Float64Array(1);
        var uInt8Float64Array_dampingi = new Uint8Array(float64Array_dampingi.buffer);
        uInt8Float64Array_dampingi[0] = buff[offset + 0];
        uInt8Float64Array_dampingi[1] = buff[offset + 1];
        uInt8Float64Array_dampingi[2] = buff[offset + 2];
        uInt8Float64Array_dampingi[3] = buff[offset + 3];
        uInt8Float64Array_dampingi[4] = buff[offset + 4];
        uInt8Float64Array_dampingi[5] = buff[offset + 5];
        uInt8Float64Array_dampingi[6] = buff[offset + 6];
        uInt8Float64Array_dampingi[7] = buff[offset + 7];
        this.damping[i] = float64Array_dampingi[0];
        offset += 8;
    }
    var length_hiStop = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_hiStop |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_hiStop |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_hiStop |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.hiStop = new Array(length_hiStop);
    for (var i = 0; i < length_hiStop; i++) {
        var float64Array_hiStopi = new Float64Array(1);
        var uInt8Float64Array_hiStopi = new Uint8Array(float64Array_hiStopi.buffer);
        uInt8Float64Array_hiStopi[0] = buff[offset + 0];
        uInt8Float64Array_hiStopi[1] = buff[offset + 1];
        uInt8Float64Array_hiStopi[2] = buff[offset + 2];
        uInt8Float64Array_hiStopi[3] = buff[offset + 3];
        uInt8Float64Array_hiStopi[4] = buff[offset + 4];
        uInt8Float64Array_hiStopi[5] = buff[offset + 5];
        uInt8Float64Array_hiStopi[6] = buff[offset + 6];
        uInt8Float64Array_hiStopi[7] = buff[offset + 7];
        this.hiStop[i] = float64Array_hiStopi[0];
        offset += 8;
    }
    var length_loStop = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_loStop |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_loStop |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_loStop |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.loStop = new Array(length_loStop);
    for (var i = 0; i < length_loStop; i++) {
        var float64Array_loStopi = new Float64Array(1);
        var uInt8Float64Array_loStopi = new Uint8Array(float64Array_loStopi.buffer);
        uInt8Float64Array_loStopi[0] = buff[offset + 0];
        uInt8Float64Array_loStopi[1] = buff[offset + 1];
        uInt8Float64Array_loStopi[2] = buff[offset + 2];
        uInt8Float64Array_loStopi[3] = buff[offset + 3];
        uInt8Float64Array_loStopi[4] = buff[offset + 4];
        uInt8Float64Array_loStopi[5] = buff[offset + 5];
        uInt8Float64Array_loStopi[6] = buff[offset + 6];
        uInt8Float64Array_loStopi[7] = buff[offset + 7];
        this.loStop[i] = float64Array_loStopi[0];
        offset += 8;
    }
    var length_erp = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_erp |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_erp |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_erp |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.erp = new Array(length_erp);
    for (var i = 0; i < length_erp; i++) {
        var float64Array_erpi = new Float64Array(1);
        var uInt8Float64Array_erpi = new Uint8Array(float64Array_erpi.buffer);
        uInt8Float64Array_erpi[0] = buff[offset + 0];
        uInt8Float64Array_erpi[1] = buff[offset + 1];
        uInt8Float64Array_erpi[2] = buff[offset + 2];
        uInt8Float64Array_erpi[3] = buff[offset + 3];
        uInt8Float64Array_erpi[4] = buff[offset + 4];
        uInt8Float64Array_erpi[5] = buff[offset + 5];
        uInt8Float64Array_erpi[6] = buff[offset + 6];
        uInt8Float64Array_erpi[7] = buff[offset + 7];
        this.erp[i] = float64Array_erpi[0];
        offset += 8;
    }
    var length_cfm = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_cfm |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_cfm |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_cfm |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.cfm = new Array(length_cfm);
    for (var i = 0; i < length_cfm; i++) {
        var float64Array_cfmi = new Float64Array(1);
        var uInt8Float64Array_cfmi = new Uint8Array(float64Array_cfmi.buffer);
        uInt8Float64Array_cfmi[0] = buff[offset + 0];
        uInt8Float64Array_cfmi[1] = buff[offset + 1];
        uInt8Float64Array_cfmi[2] = buff[offset + 2];
        uInt8Float64Array_cfmi[3] = buff[offset + 3];
        uInt8Float64Array_cfmi[4] = buff[offset + 4];
        uInt8Float64Array_cfmi[5] = buff[offset + 5];
        uInt8Float64Array_cfmi[6] = buff[offset + 6];
        uInt8Float64Array_cfmi[7] = buff[offset + 7];
        this.cfm[i] = float64Array_cfmi[0];
        offset += 8;
    }
    var length_stop_erp = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_stop_erp |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_stop_erp |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_stop_erp |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.stop_erp = new Array(length_stop_erp);
    for (var i = 0; i < length_stop_erp; i++) {
        var float64Array_stop_erpi = new Float64Array(1);
        var uInt8Float64Array_stop_erpi = new Uint8Array(float64Array_stop_erpi.buffer);
        uInt8Float64Array_stop_erpi[0] = buff[offset + 0];
        uInt8Float64Array_stop_erpi[1] = buff[offset + 1];
        uInt8Float64Array_stop_erpi[2] = buff[offset + 2];
        uInt8Float64Array_stop_erpi[3] = buff[offset + 3];
        uInt8Float64Array_stop_erpi[4] = buff[offset + 4];
        uInt8Float64Array_stop_erpi[5] = buff[offset + 5];
        uInt8Float64Array_stop_erpi[6] = buff[offset + 6];
        uInt8Float64Array_stop_erpi[7] = buff[offset + 7];
        this.stop_erp[i] = float64Array_stop_erpi[0];
        offset += 8;
    }
    var length_stop_cfm = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_stop_cfm |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_stop_cfm |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_stop_cfm |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.stop_cfm = new Array(length_stop_cfm);
    for (var i = 0; i < length_stop_cfm; i++) {
        var float64Array_stop_cfmi = new Float64Array(1);
        var uInt8Float64Array_stop_cfmi = new Uint8Array(float64Array_stop_cfmi.buffer);
        uInt8Float64Array_stop_cfmi[0] = buff[offset + 0];
        uInt8Float64Array_stop_cfmi[1] = buff[offset + 1];
        uInt8Float64Array_stop_cfmi[2] = buff[offset + 2];
        uInt8Float64Array_stop_cfmi[3] = buff[offset + 3];
        uInt8Float64Array_stop_cfmi[4] = buff[offset + 4];
        uInt8Float64Array_stop_cfmi[5] = buff[offset + 5];
        uInt8Float64Array_stop_cfmi[6] = buff[offset + 6];
        uInt8Float64Array_stop_cfmi[7] = buff[offset + 7];
        this.stop_cfm[i] = float64Array_stop_cfmi[0];
        offset += 8;
    }
    var length_fudge_factor = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_fudge_factor |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_fudge_factor |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_fudge_factor |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.fudge_factor = new Array(length_fudge_factor);
    for (var i = 0; i < length_fudge_factor; i++) {
        var float64Array_fudge_factori = new Float64Array(1);
        var uInt8Float64Array_fudge_factori = new Uint8Array(float64Array_fudge_factori.buffer);
        uInt8Float64Array_fudge_factori[0] = buff[offset + 0];
        uInt8Float64Array_fudge_factori[1] = buff[offset + 1];
        uInt8Float64Array_fudge_factori[2] = buff[offset + 2];
        uInt8Float64Array_fudge_factori[3] = buff[offset + 3];
        uInt8Float64Array_fudge_factori[4] = buff[offset + 4];
        uInt8Float64Array_fudge_factori[5] = buff[offset + 5];
        uInt8Float64Array_fudge_factori[6] = buff[offset + 6];
        uInt8Float64Array_fudge_factori[7] = buff[offset + 7];
        this.fudge_factor[i] = float64Array_fudge_factori[0];
        offset += 8;
    }
    var length_fmax = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_fmax |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_fmax |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_fmax |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.fmax = new Array(length_fmax);
    for (var i = 0; i < length_fmax; i++) {
        var float64Array_fmaxi = new Float64Array(1);
        var uInt8Float64Array_fmaxi = new Uint8Array(float64Array_fmaxi.buffer);
        uInt8Float64Array_fmaxi[0] = buff[offset + 0];
        uInt8Float64Array_fmaxi[1] = buff[offset + 1];
        uInt8Float64Array_fmaxi[2] = buff[offset + 2];
        uInt8Float64Array_fmaxi[3] = buff[offset + 3];
        uInt8Float64Array_fmaxi[4] = buff[offset + 4];
        uInt8Float64Array_fmaxi[5] = buff[offset + 5];
        uInt8Float64Array_fmaxi[6] = buff[offset + 6];
        uInt8Float64Array_fmaxi[7] = buff[offset + 7];
        this.fmax[i] = float64Array_fmaxi[0];
        offset += 8;
    }
    var length_vel = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_vel |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_vel |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_vel |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.vel = new Array(length_vel);
    for (var i = 0; i < length_vel; i++) {
        var float64Array_veli = new Float64Array(1);
        var uInt8Float64Array_veli = new Uint8Array(float64Array_veli.buffer);
        uInt8Float64Array_veli[0] = buff[offset + 0];
        uInt8Float64Array_veli[1] = buff[offset + 1];
        uInt8Float64Array_veli[2] = buff[offset + 2];
        uInt8Float64Array_veli[3] = buff[offset + 3];
        uInt8Float64Array_veli[4] = buff[offset + 4];
        uInt8Float64Array_veli[5] = buff[offset + 5];
        uInt8Float64Array_veli[6] = buff[offset + 6];
        uInt8Float64Array_veli[7] = buff[offset + 7];
        this.vel[i] = float64Array_veli[0];
        offset += 8;
    }
    return offset;
};

ODEJointProperties.prototype.serializedLength = function() {
    var length = 0;
    var length_damping = this.damping.length;
    length += 4;
    for (var i = 0; i < length_damping; i++) {
        length += 8
    }
    var length_hiStop = this.hiStop.length;
    length += 4;
    for (var i = 0; i < length_hiStop; i++) {
        length += 8
    }
    var length_loStop = this.loStop.length;
    length += 4;
    for (var i = 0; i < length_loStop; i++) {
        length += 8
    }
    var length_erp = this.erp.length;
    length += 4;
    for (var i = 0; i < length_erp; i++) {
        length += 8
    }
    var length_cfm = this.cfm.length;
    length += 4;
    for (var i = 0; i < length_cfm; i++) {
        length += 8
    }
    var length_stop_erp = this.stop_erp.length;
    length += 4;
    for (var i = 0; i < length_stop_erp; i++) {
        length += 8
    }
    var length_stop_cfm = this.stop_cfm.length;
    length += 4;
    for (var i = 0; i < length_stop_cfm; i++) {
        length += 8
    }
    var length_fudge_factor = this.fudge_factor.length;
    length += 4;
    for (var i = 0; i < length_fudge_factor; i++) {
        length += 8
    }
    var length_fmax = this.fmax.length;
    length += 4;
    for (var i = 0; i < length_fmax; i++) {
        length += 8
    }
    var length_vel = this.vel.length;
    length += 4;
    for (var i = 0; i < length_vel; i++) {
        length += 8
    }
    return length;
};

ODEJointProperties.prototype.echo = function() { return ""; };

ODEJointProperties.prototype.getType = function() { return "gazebo_msgs/ODEJointProperties"; };

ODEJointProperties.prototype.getMD5 = function() { return "a9e264dbf3eff8e202d2bebecf081639"; };

ODEJointProperties.prototype.getID = function() { return 0; };

ODEJointProperties.prototype.setID = function(id) { };

return function() { return new ODEJointProperties(); };

});
