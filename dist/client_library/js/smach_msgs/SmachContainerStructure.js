(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.smach_msgs==="undefined"){g.smach_msgs={};}g.smach_msgs.SmachContainerStructure=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function SmachContainerStructure() {
    this.header = std_msgs.Header();
    this.path = "";
    this.children = new Array();
    this.internal_outcomes = new Array();
    this.outcomes_from = new Array();
    this.outcomes_to = new Array();
    this.container_outcomes = new Array();
};

SmachContainerStructure.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var encoder_path = new TextEncoder('utf8');
    var utf8array_path = encoder_path.encode(this.path);
    buff[offset + 0] = (utf8array_path.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_path.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_path.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_path.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_path.length; i++) {
        buff[offset + i] = utf8array_path[i];
    }
    offset += utf8array_path.length;
    var length_children = this.children.length;
    buff[offset + 0] = (length_children >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_children >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_children >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_children >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_children; i++) {
        var encoder_childreni = new TextEncoder('utf8');
        var utf8array_childreni = encoder_childreni.encode(this.children[i]);
        buff[offset + 0] = (utf8array_childreni.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_childreni.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_childreni.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_childreni.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_childreni.length; i++) {
            buff[offset + i] = utf8array_childreni[i];
        }
        offset += utf8array_childreni.length;
    }
    var length_internal_outcomes = this.internal_outcomes.length;
    buff[offset + 0] = (length_internal_outcomes >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_internal_outcomes >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_internal_outcomes >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_internal_outcomes >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_internal_outcomes; i++) {
        var encoder_internal_outcomesi = new TextEncoder('utf8');
        var utf8array_internal_outcomesi = encoder_internal_outcomesi.encode(this.internal_outcomes[i]);
        buff[offset + 0] = (utf8array_internal_outcomesi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_internal_outcomesi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_internal_outcomesi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_internal_outcomesi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_internal_outcomesi.length; i++) {
            buff[offset + i] = utf8array_internal_outcomesi[i];
        }
        offset += utf8array_internal_outcomesi.length;
    }
    var length_outcomes_from = this.outcomes_from.length;
    buff[offset + 0] = (length_outcomes_from >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_outcomes_from >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_outcomes_from >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_outcomes_from >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_outcomes_from; i++) {
        var encoder_outcomes_fromi = new TextEncoder('utf8');
        var utf8array_outcomes_fromi = encoder_outcomes_fromi.encode(this.outcomes_from[i]);
        buff[offset + 0] = (utf8array_outcomes_fromi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_outcomes_fromi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_outcomes_fromi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_outcomes_fromi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_outcomes_fromi.length; i++) {
            buff[offset + i] = utf8array_outcomes_fromi[i];
        }
        offset += utf8array_outcomes_fromi.length;
    }
    var length_outcomes_to = this.outcomes_to.length;
    buff[offset + 0] = (length_outcomes_to >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_outcomes_to >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_outcomes_to >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_outcomes_to >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_outcomes_to; i++) {
        var encoder_outcomes_toi = new TextEncoder('utf8');
        var utf8array_outcomes_toi = encoder_outcomes_toi.encode(this.outcomes_to[i]);
        buff[offset + 0] = (utf8array_outcomes_toi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_outcomes_toi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_outcomes_toi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_outcomes_toi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_outcomes_toi.length; i++) {
            buff[offset + i] = utf8array_outcomes_toi[i];
        }
        offset += utf8array_outcomes_toi.length;
    }
    var length_container_outcomes = this.container_outcomes.length;
    buff[offset + 0] = (length_container_outcomes >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_container_outcomes >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_container_outcomes >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_container_outcomes >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_container_outcomes; i++) {
        var encoder_container_outcomesi = new TextEncoder('utf8');
        var utf8array_container_outcomesi = encoder_container_outcomesi.encode(this.container_outcomes[i]);
        buff[offset + 0] = (utf8array_container_outcomesi.length >> (8 * 0)) & 0xFF;
        buff[offset + 1] = (utf8array_container_outcomesi.length >> (8 * 1)) & 0xFF;
        buff[offset + 2] = (utf8array_container_outcomesi.length >> (8 * 2)) & 0xFF;
        buff[offset + 3] = (utf8array_container_outcomesi.length >> (8 * 3)) & 0xFF;
        offset += 4;
        for (var i = 0; i < utf8array_container_outcomesi.length; i++) {
            buff[offset + i] = utf8array_container_outcomesi[i];
        }
        offset += utf8array_container_outcomesi.length;
    }
    return offset;
};

SmachContainerStructure.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var length_path = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_path |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_path |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_path |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_path = new TextDecoder('utf8');
    this.path = decoder_path.decode(buff.slice(offset, offset + length_path));
    offset += length_path;
    var length_children = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_children |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_children |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_children |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.children = new Array(length_children);
    for (var i = 0; i < length_children; i++) {
        var length_childreni = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_childreni |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_childreni |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_childreni |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_childreni = new TextDecoder('utf8');
        this.children[i] = decoder_childreni.decode(buff.slice(offset, offset + length_childreni));
        offset += length_childreni;
    }
    var length_internal_outcomes = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_internal_outcomes |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_internal_outcomes |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_internal_outcomes |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.internal_outcomes = new Array(length_internal_outcomes);
    for (var i = 0; i < length_internal_outcomes; i++) {
        var length_internal_outcomesi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_internal_outcomesi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_internal_outcomesi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_internal_outcomesi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_internal_outcomesi = new TextDecoder('utf8');
        this.internal_outcomes[i] = decoder_internal_outcomesi.decode(buff.slice(offset, offset + length_internal_outcomesi));
        offset += length_internal_outcomesi;
    }
    var length_outcomes_from = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_outcomes_from |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_outcomes_from |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_outcomes_from |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.outcomes_from = new Array(length_outcomes_from);
    for (var i = 0; i < length_outcomes_from; i++) {
        var length_outcomes_fromi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_outcomes_fromi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_outcomes_fromi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_outcomes_fromi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_outcomes_fromi = new TextDecoder('utf8');
        this.outcomes_from[i] = decoder_outcomes_fromi.decode(buff.slice(offset, offset + length_outcomes_fromi));
        offset += length_outcomes_fromi;
    }
    var length_outcomes_to = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_outcomes_to |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_outcomes_to |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_outcomes_to |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.outcomes_to = new Array(length_outcomes_to);
    for (var i = 0; i < length_outcomes_to; i++) {
        var length_outcomes_toi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_outcomes_toi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_outcomes_toi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_outcomes_toi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_outcomes_toi = new TextDecoder('utf8');
        this.outcomes_to[i] = decoder_outcomes_toi.decode(buff.slice(offset, offset + length_outcomes_toi));
        offset += length_outcomes_toi;
    }
    var length_container_outcomes = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_container_outcomes |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_container_outcomes |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_container_outcomes |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.container_outcomes = new Array(length_container_outcomes);
    for (var i = 0; i < length_container_outcomes; i++) {
        var length_container_outcomesi = +((buff[offset + 0] & 0xFF) << (8 * 0));
        length_container_outcomesi |= +((buff[offset + 1] & 0xFF) << (8 * 1));
        length_container_outcomesi |= +((buff[offset + 2] & 0xFF) << (8 * 2));
        length_container_outcomesi |= +((buff[offset + 3] & 0xFF) << (8 * 3))
        offset += 4;
        var decoder_container_outcomesi = new TextDecoder('utf8');
        this.container_outcomes[i] = decoder_container_outcomesi.decode(buff.slice(offset, offset + length_container_outcomesi));
        offset += length_container_outcomesi;
    }
    return offset;
};

SmachContainerStructure.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    var encoder_path = new TextEncoder('utf8');
    var utf8array_path = encoder_path.encode(this.path);
    length += 4;
    length += utf8array_path.length;
    var length_children = this.children.length;
    length += 4;
    for (var i = 0; i < length_children; i++) {
        var encoder_childreni = new TextEncoder('utf8');
        var utf8array_childreni = encoder_childreni.encode(this.children[i]);
        length += 4;
        length += utf8array_childreni.length;
    }
    var length_internal_outcomes = this.internal_outcomes.length;
    length += 4;
    for (var i = 0; i < length_internal_outcomes; i++) {
        var encoder_internal_outcomesi = new TextEncoder('utf8');
        var utf8array_internal_outcomesi = encoder_internal_outcomesi.encode(this.internal_outcomes[i]);
        length += 4;
        length += utf8array_internal_outcomesi.length;
    }
    var length_outcomes_from = this.outcomes_from.length;
    length += 4;
    for (var i = 0; i < length_outcomes_from; i++) {
        var encoder_outcomes_fromi = new TextEncoder('utf8');
        var utf8array_outcomes_fromi = encoder_outcomes_fromi.encode(this.outcomes_from[i]);
        length += 4;
        length += utf8array_outcomes_fromi.length;
    }
    var length_outcomes_to = this.outcomes_to.length;
    length += 4;
    for (var i = 0; i < length_outcomes_to; i++) {
        var encoder_outcomes_toi = new TextEncoder('utf8');
        var utf8array_outcomes_toi = encoder_outcomes_toi.encode(this.outcomes_to[i]);
        length += 4;
        length += utf8array_outcomes_toi.length;
    }
    var length_container_outcomes = this.container_outcomes.length;
    length += 4;
    for (var i = 0; i < length_container_outcomes; i++) {
        var encoder_container_outcomesi = new TextEncoder('utf8');
        var utf8array_container_outcomesi = encoder_container_outcomesi.encode(this.container_outcomes[i]);
        length += 4;
        length += utf8array_container_outcomesi.length;
    }
    return length;
};

SmachContainerStructure.prototype.echo = function() { return ""; };

SmachContainerStructure.prototype.getType = function() { return "smach_msgs/SmachContainerStructure"; };

SmachContainerStructure.prototype.getMD5 = function() { return "0209663ab13753a56b6ac1d094d07fbe"; };

SmachContainerStructure.prototype.getID = function() { return 0; };

SmachContainerStructure.prototype.setID = function(id) { };

return function() { return new SmachContainerStructure(); };

});
