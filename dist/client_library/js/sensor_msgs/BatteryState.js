(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}else if(typeof define==="function"&&define.amd){define([],f);}else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}if(typeof g.sensor_msgs==="undefined"){g.sensor_msgs={};}g.sensor_msgs.BatteryState=f();}})(function(){var define,module,exports;'use strict';
const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));
const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));
document.write("<script language=javascript src='"+_ROOT_PATH_+"/std_msgs/Header.js'></script>");

function BatteryState() {
    this.header = std_msgs.Header();
    this.voltage = 0.0;
    this.current = 0.0;
    this.charge = 0.0;
    this.capacity = 0.0;
    this.design_capacity = 0.0;
    this.percentage = 0.0;
    this.power_supply_status = 0;
    this.power_supply_health = 0;
    this.power_supply_technology = 0;
    this.present = false;
    this.cell_voltage = new Array();
    this.location = "";
    this.serial_number = "";

    // ENUM{
    this.POWER_SUPPLY_STATUS_UNKNOWN =  0;
    this.POWER_SUPPLY_STATUS_CHARGING =  1;
    this.POWER_SUPPLY_STATUS_DISCHARGING =  2;
    this.POWER_SUPPLY_STATUS_NOT_CHARGING =  3;
    this.POWER_SUPPLY_STATUS_FULL =  4;
    this.POWER_SUPPLY_HEALTH_UNKNOWN =  0;
    this.POWER_SUPPLY_HEALTH_GOOD =  1;
    this.POWER_SUPPLY_HEALTH_OVERHEAT =  2;
    this.POWER_SUPPLY_HEALTH_DEAD =  3;
    this.POWER_SUPPLY_HEALTH_OVERVOLTAGE =  4;
    this.POWER_SUPPLY_HEALTH_UNSPEC_FAILURE =  5;
    this.POWER_SUPPLY_HEALTH_COLD =  6;
    this.POWER_SUPPLY_HEALTH_WATCHDOG_TIMER_EXPIRE =  7;
    this.POWER_SUPPLY_HEALTH_SAFETY_TIMER_EXPIRE =  8;
    this.POWER_SUPPLY_TECHNOLOGY_UNKNOWN =  0;
    this.POWER_SUPPLY_TECHNOLOGY_NIMH =  1;
    this.POWER_SUPPLY_TECHNOLOGY_LION =  2;
    this.POWER_SUPPLY_TECHNOLOGY_LIPO =  3;
    this.POWER_SUPPLY_TECHNOLOGY_LIFE =  4;
    this.POWER_SUPPLY_TECHNOLOGY_NICD =  5;
    this.POWER_SUPPLY_TECHNOLOGY_LIMN =  6;
    // }ENUM
};

BatteryState.prototype.serialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.serialize(buff, offset);
    var float32Array_voltage = new Float32Array(1);
    var uInt8Float32Array_voltage = new Uint8Array(float32Array_voltage.buffer);
    float32Array_voltage[0] = +this.voltage;
    buff[offset + 0] = uInt8Float32Array_voltage[0];
    buff[offset + 1] = uInt8Float32Array_voltage[1];
    buff[offset + 2] = uInt8Float32Array_voltage[2];
    buff[offset + 3] = uInt8Float32Array_voltage[3];
    offset += 4;
    var float32Array_current = new Float32Array(1);
    var uInt8Float32Array_current = new Uint8Array(float32Array_current.buffer);
    float32Array_current[0] = +this.current;
    buff[offset + 0] = uInt8Float32Array_current[0];
    buff[offset + 1] = uInt8Float32Array_current[1];
    buff[offset + 2] = uInt8Float32Array_current[2];
    buff[offset + 3] = uInt8Float32Array_current[3];
    offset += 4;
    var float32Array_charge = new Float32Array(1);
    var uInt8Float32Array_charge = new Uint8Array(float32Array_charge.buffer);
    float32Array_charge[0] = +this.charge;
    buff[offset + 0] = uInt8Float32Array_charge[0];
    buff[offset + 1] = uInt8Float32Array_charge[1];
    buff[offset + 2] = uInt8Float32Array_charge[2];
    buff[offset + 3] = uInt8Float32Array_charge[3];
    offset += 4;
    var float32Array_capacity = new Float32Array(1);
    var uInt8Float32Array_capacity = new Uint8Array(float32Array_capacity.buffer);
    float32Array_capacity[0] = +this.capacity;
    buff[offset + 0] = uInt8Float32Array_capacity[0];
    buff[offset + 1] = uInt8Float32Array_capacity[1];
    buff[offset + 2] = uInt8Float32Array_capacity[2];
    buff[offset + 3] = uInt8Float32Array_capacity[3];
    offset += 4;
    var float32Array_design_capacity = new Float32Array(1);
    var uInt8Float32Array_design_capacity = new Uint8Array(float32Array_design_capacity.buffer);
    float32Array_design_capacity[0] = +this.design_capacity;
    buff[offset + 0] = uInt8Float32Array_design_capacity[0];
    buff[offset + 1] = uInt8Float32Array_design_capacity[1];
    buff[offset + 2] = uInt8Float32Array_design_capacity[2];
    buff[offset + 3] = uInt8Float32Array_design_capacity[3];
    offset += 4;
    var float32Array_percentage = new Float32Array(1);
    var uInt8Float32Array_percentage = new Uint8Array(float32Array_percentage.buffer);
    float32Array_percentage[0] = +this.percentage;
    buff[offset + 0] = uInt8Float32Array_percentage[0];
    buff[offset + 1] = uInt8Float32Array_percentage[1];
    buff[offset + 2] = uInt8Float32Array_percentage[2];
    buff[offset + 3] = uInt8Float32Array_percentage[3];
    offset += 4;
    buff[offset + 0] = ((+this.power_supply_status) >> (8 * 0)) & 0xFF;
    offset += 1;
    buff[offset + 0] = ((+this.power_supply_health) >> (8 * 0)) & 0xFF;
    offset += 1;
    buff[offset + 0] = ((+this.power_supply_technology) >> (8 * 0)) & 0xFF;
    offset += 1;
    buff[offset] = this.present === false ? 0 : 1;
    offset += 1;
    var length_cell_voltage = this.cell_voltage.length;
    buff[offset + 0] = (length_cell_voltage >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (length_cell_voltage >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (length_cell_voltage >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (length_cell_voltage >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < length_cell_voltage; i++) {
        var float32Array_cell_voltagei = new Float32Array(1);
        var uInt8Float32Array_cell_voltagei = new Uint8Array(float32Array_cell_voltagei.buffer);
        float32Array_cell_voltagei[0] = +this.cell_voltage[i];
        buff[offset + 0] = uInt8Float32Array_cell_voltagei[0];
        buff[offset + 1] = uInt8Float32Array_cell_voltagei[1];
        buff[offset + 2] = uInt8Float32Array_cell_voltagei[2];
        buff[offset + 3] = uInt8Float32Array_cell_voltagei[3];
        offset += 4;
    }
    var encoder_location = new TextEncoder('utf8');
    var utf8array_location = encoder_location.encode(this.location);
    buff[offset + 0] = (utf8array_location.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_location.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_location.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_location.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_location.length; i++) {
        buff[offset + i] = utf8array_location[i];
    }
    offset += utf8array_location.length;
    var encoder_serial_number = new TextEncoder('utf8');
    var utf8array_serial_number = encoder_serial_number.encode(this.serial_number);
    buff[offset + 0] = (utf8array_serial_number.length >> (8 * 0)) & 0xFF;
    buff[offset + 1] = (utf8array_serial_number.length >> (8 * 1)) & 0xFF;
    buff[offset + 2] = (utf8array_serial_number.length >> (8 * 2)) & 0xFF;
    buff[offset + 3] = (utf8array_serial_number.length >> (8 * 3)) & 0xFF;
    offset += 4;
    for (var i = 0; i < utf8array_serial_number.length; i++) {
        buff[offset + i] = utf8array_serial_number[i];
    }
    offset += utf8array_serial_number.length;
    return offset;
};

BatteryState.prototype.deserialize = function(buff, idx) {
    var offset = idx;
    offset += this.header.deserialize(buff, offset);
    var float32Array_voltage = new Float32Array(1);
    var uInt8Float32Array_voltage = new Uint8Array(float32Array_voltage.buffer);
    uInt8Float32Array_voltage[0] = buff[offset + 0];
    uInt8Float32Array_voltage[1] = buff[offset + 1];
    uInt8Float32Array_voltage[2] = buff[offset + 2];
    uInt8Float32Array_voltage[3] = buff[offset + 3];
    this.voltage = float32Array_voltage[0];
    offset += 4;
    var float32Array_current = new Float32Array(1);
    var uInt8Float32Array_current = new Uint8Array(float32Array_current.buffer);
    uInt8Float32Array_current[0] = buff[offset + 0];
    uInt8Float32Array_current[1] = buff[offset + 1];
    uInt8Float32Array_current[2] = buff[offset + 2];
    uInt8Float32Array_current[3] = buff[offset + 3];
    this.current = float32Array_current[0];
    offset += 4;
    var float32Array_charge = new Float32Array(1);
    var uInt8Float32Array_charge = new Uint8Array(float32Array_charge.buffer);
    uInt8Float32Array_charge[0] = buff[offset + 0];
    uInt8Float32Array_charge[1] = buff[offset + 1];
    uInt8Float32Array_charge[2] = buff[offset + 2];
    uInt8Float32Array_charge[3] = buff[offset + 3];
    this.charge = float32Array_charge[0];
    offset += 4;
    var float32Array_capacity = new Float32Array(1);
    var uInt8Float32Array_capacity = new Uint8Array(float32Array_capacity.buffer);
    uInt8Float32Array_capacity[0] = buff[offset + 0];
    uInt8Float32Array_capacity[1] = buff[offset + 1];
    uInt8Float32Array_capacity[2] = buff[offset + 2];
    uInt8Float32Array_capacity[3] = buff[offset + 3];
    this.capacity = float32Array_capacity[0];
    offset += 4;
    var float32Array_design_capacity = new Float32Array(1);
    var uInt8Float32Array_design_capacity = new Uint8Array(float32Array_design_capacity.buffer);
    uInt8Float32Array_design_capacity[0] = buff[offset + 0];
    uInt8Float32Array_design_capacity[1] = buff[offset + 1];
    uInt8Float32Array_design_capacity[2] = buff[offset + 2];
    uInt8Float32Array_design_capacity[3] = buff[offset + 3];
    this.design_capacity = float32Array_design_capacity[0];
    offset += 4;
    var float32Array_percentage = new Float32Array(1);
    var uInt8Float32Array_percentage = new Uint8Array(float32Array_percentage.buffer);
    uInt8Float32Array_percentage[0] = buff[offset + 0];
    uInt8Float32Array_percentage[1] = buff[offset + 1];
    uInt8Float32Array_percentage[2] = buff[offset + 2];
    uInt8Float32Array_percentage[3] = buff[offset + 3];
    this.percentage = float32Array_percentage[0];
    offset += 4;
    this.power_supply_status = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    this.power_supply_health = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    this.power_supply_technology = +((buff[offset + 0] & 0xFF) << (8 * 0));
    offset += 1;
    this.present = buff[offset] !== 0 ? true : false;
    offset += 1;
    var length_cell_voltage = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_cell_voltage |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_cell_voltage |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_cell_voltage |= +((buff[offset + 3] & 0xFF) << (8 * 3));
    offset += 4;
    this.cell_voltage = new Array(length_cell_voltage);
    for (var i = 0; i < length_cell_voltage; i++) {
        var float32Array_cell_voltagei = new Float32Array(1);
        var uInt8Float32Array_cell_voltagei = new Uint8Array(float32Array_cell_voltagei.buffer);
        uInt8Float32Array_cell_voltagei[0] = buff[offset + 0];
        uInt8Float32Array_cell_voltagei[1] = buff[offset + 1];
        uInt8Float32Array_cell_voltagei[2] = buff[offset + 2];
        uInt8Float32Array_cell_voltagei[3] = buff[offset + 3];
        this.cell_voltage[i] = float32Array_cell_voltagei[0];
        offset += 4;
    }
    var length_location = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_location |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_location |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_location |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_location = new TextDecoder('utf8');
    this.location = decoder_location.decode(buff.slice(offset, offset + length_location));
    offset += length_location;
    var length_serial_number = +((buff[offset + 0] & 0xFF) << (8 * 0));
    length_serial_number |= +((buff[offset + 1] & 0xFF) << (8 * 1));
    length_serial_number |= +((buff[offset + 2] & 0xFF) << (8 * 2));
    length_serial_number |= +((buff[offset + 3] & 0xFF) << (8 * 3))
    offset += 4;
    var decoder_serial_number = new TextDecoder('utf8');
    this.serial_number = decoder_serial_number.decode(buff.slice(offset, offset + length_serial_number));
    offset += length_serial_number;
    return offset;
};

BatteryState.prototype.serializedLength = function() {
    var length = 0;
    length += this.header.serializedLength();
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 4
    length += 1
    length += 1
    length += 1
    length += 1
    var length_cell_voltage = this.cell_voltage.length;
    length += 4;
    for (var i = 0; i < length_cell_voltage; i++) {
        length += 4
    }
    var encoder_location = new TextEncoder('utf8');
    var utf8array_location = encoder_location.encode(this.location);
    length += 4;
    length += utf8array_location.length;
    var encoder_serial_number = new TextEncoder('utf8');
    var utf8array_serial_number = encoder_serial_number.encode(this.serial_number);
    length += 4;
    length += utf8array_serial_number.length;
    return length;
};

BatteryState.prototype.echo = function() { return ""; };

BatteryState.prototype.getType = function() { return "sensor_msgs/BatteryState"; };

BatteryState.prototype.getMD5 = function() { return "715c4769cacd76e4b679cc3ea4c347b4"; };

BatteryState.prototype.getID = function() { return 0; };

BatteryState.prototype.setID = function(id) { };

return function() { return new BatteryState(); };

});
