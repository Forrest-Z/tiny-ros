import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import sensor_msgs.msg
import std_msgs.msg

class BatteryState(tinyros.Message):
    __slots__ = ['header','voltage','current','charge','capacity','design_capacity','percentage','power_supply_status','power_supply_health','power_supply_technology','present','cell_voltage','location','serial_number']
    _slot_types = ['std_msgs.msg.Header','float32','float32','float32','float32','float32','float32','uint8','uint8','uint8','bool','float32[]','string','string']

    POWER_SUPPLY_STATUS_UNKNOWN =  0
    POWER_SUPPLY_STATUS_CHARGING =  1
    POWER_SUPPLY_STATUS_DISCHARGING =  2
    POWER_SUPPLY_STATUS_NOT_CHARGING =  3
    POWER_SUPPLY_STATUS_FULL =  4
    POWER_SUPPLY_HEALTH_UNKNOWN =  0
    POWER_SUPPLY_HEALTH_GOOD =  1
    POWER_SUPPLY_HEALTH_OVERHEAT =  2
    POWER_SUPPLY_HEALTH_DEAD =  3
    POWER_SUPPLY_HEALTH_OVERVOLTAGE =  4
    POWER_SUPPLY_HEALTH_UNSPEC_FAILURE =  5
    POWER_SUPPLY_HEALTH_COLD =  6
    POWER_SUPPLY_HEALTH_WATCHDOG_TIMER_EXPIRE =  7
    POWER_SUPPLY_HEALTH_SAFETY_TIMER_EXPIRE =  8
    POWER_SUPPLY_TECHNOLOGY_UNKNOWN =  0
    POWER_SUPPLY_TECHNOLOGY_NIMH =  1
    POWER_SUPPLY_TECHNOLOGY_LION =  2
    POWER_SUPPLY_TECHNOLOGY_LIPO =  3
    POWER_SUPPLY_TECHNOLOGY_LIFE =  4
    POWER_SUPPLY_TECHNOLOGY_NICD =  5
    POWER_SUPPLY_TECHNOLOGY_LIMN =  6

    def __init__(self):
        super(BatteryState, self).__init__()
        self.header = std_msgs.msg.Header()
        self.voltage = 0.
        self.current = 0.
        self.charge = 0.
        self.capacity = 0.
        self.design_capacity = 0.
        self.percentage = 0.
        self.power_supply_status = 0
        self.power_supply_health = 0
        self.power_supply_technology = 0
        self.present = False
        self.cell_voltage = []
        self.location = ''
        self.serial_number = ''

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            struct_voltage = struct.Struct("<f")
            buff.write(struct_voltage.pack(self.voltage))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_current = struct.Struct("<f")
            buff.write(struct_current.pack(self.current))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_charge = struct.Struct("<f")
            buff.write(struct_charge.pack(self.charge))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_capacity = struct.Struct("<f")
            buff.write(struct_capacity.pack(self.capacity))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_design_capacity = struct.Struct("<f")
            buff.write(struct_design_capacity.pack(self.design_capacity))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_percentage = struct.Struct("<f")
            buff.write(struct_percentage.pack(self.percentage))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_power_supply_status = struct.Struct("<B")
            buff.write(struct_power_supply_status.pack(self.power_supply_status))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_power_supply_health = struct.Struct("<B")
            buff.write(struct_power_supply_health.pack(self.power_supply_health))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_power_supply_technology = struct.Struct("<B")
            buff.write(struct_power_supply_technology.pack(self.power_supply_technology))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_present = struct.Struct("<B")
            buff.write(struct_present.pack(self.present))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            cell_voltage_length = len(self.cell_voltage)
            buff.write(_struct_I.pack(cell_voltage_length))
            offset += 4
            for i in range(0, cell_voltage_length):
                try:
                    struct_cell_voltagei = struct.Struct("<f")
                    buff.write(struct_cell_voltagei.pack(self.cell_voltage[i]))
                    offset += 4
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.location
            length = len(_x)
            if python3 or type(_x) == unicode:
                _x = _x.encode('utf-8')
                length = len(_x)
            if python3:
                buff.write(struct.pack('<I%sB'%length, length, *_x))
            else:
                buff.write(struct.pack('<I%ss'%length, length, _x))
            offset += 4
            offset += length
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            _x = self.serial_number
            length = len(_x)
            if python3 or type(_x) == unicode:
                _x = _x.encode('utf-8')
                length = len(_x)
            if python3:
                buff.write(struct.pack('<I%sB'%length, length, *_x))
            else:
                buff.write(struct.pack('<I%ss'%length, length, _x))
            offset += 4
            offset += length
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            struct_voltage = struct.Struct("<f")
            (self.voltage,) = struct_voltage.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_current = struct.Struct("<f")
            (self.current,) = struct_current.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_charge = struct.Struct("<f")
            (self.charge,) = struct_charge.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_capacity = struct.Struct("<f")
            (self.capacity,) = struct_capacity.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_design_capacity = struct.Struct("<f")
            (self.design_capacity,) = struct_design_capacity.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_percentage = struct.Struct("<f")
            (self.percentage,) = struct_percentage.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_power_supply_status = struct.Struct("<B")
            (self.power_supply_status,) = struct_power_supply_status.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_power_supply_health = struct.Struct("<B")
            (self.power_supply_health,) = struct_power_supply_health.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_power_supply_technology = struct.Struct("<B")
            (self.power_supply_technology,) = struct_power_supply_technology.unpack(buff[offset:(offset + 1)])
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_present = struct.Struct("<B")
            (self.present,) = struct_present.unpack(buff[offset:(offset + 1)])
            self.present = bool(self.present)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (cell_voltage_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.cell_voltage = [0. for x in range(0, cell_voltage_length)]
            offset += 4
            for i in range(0, cell_voltage_length):
                try:
                    struct_cell_voltagei = struct.Struct("<f")
                    (self.cell_voltage[i],) = struct_cell_voltagei.unpack(buff[offset:(offset + 4)])
                    offset += 4
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_location,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.location = buff[offset:(offset + length_location)].decode('utf-8')
            else:
                self.location = buff[offset:(offset + length_location)]
            offset += length_location
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (length_serial_number,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
            if python3:
                self.serial_number = buff[offset:(offset + length_serial_number)].decode('utf-8')
            else:
                self.serial_number = buff[offset:(offset + length_serial_number)]
            offset += length_serial_number
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
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
        cell_voltage_length = len(self.cell_voltage)
        length += 4
        for i in range(0, cell_voltage_length):
            length += 4
        location_x = self.location
        location_length = len(location_x)
        if python3 or type(location_x) == unicode:
            location_x = location_x.encode('utf-8')
            location_length = len(location_x)
        length += 4
        length += location_length
        serial_number_x = self.serial_number
        serial_number_length = len(serial_number_x)
        if python3 or type(serial_number_x) == unicode:
            serial_number_x = serial_number_x.encode('utf-8')
            serial_number_length = len(serial_number_x)
        length += 4
        length += serial_number_length
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"voltage": %s'%voltage
        string_echo += ', '
        string_echo += '"current": %s'%current
        string_echo += ', '
        string_echo += '"charge": %s'%charge
        string_echo += ', '
        string_echo += '"capacity": %s'%capacity
        string_echo += ', '
        string_echo += '"design_capacity": %s'%design_capacity
        string_echo += ', '
        string_echo += '"percentage": %s'%percentage
        string_echo += ', '
        string_echo += '"power_supply_status": %s'%power_supply_status
        string_echo += ', '
        string_echo += '"power_supply_health": %s'%power_supply_health
        string_echo += ', '
        string_echo += '"power_supply_technology": %s'%power_supply_technology
        string_echo += ', '
        string_echo += '"present": %s'%present
        string_echo += ', '
        string_echo += '"cell_voltage": ['
        cell_voltage_length = len(self.cell_voltage)
        for i in range(0, cell_voltage_length):
            if i == (cell_voltage_length - 1): 
                string_echo += '{"cell_voltage%s": %s'%(i,cell_voltage[i])
                string_echo += '}'
            else:
                string_echo += '{"cell_voltage%s": %s'%(i,cell_voltage[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"location": "%s"'%location
        string_echo += '", '
        string_echo += '"serial_number": "%s"'%serial_number
        string_echo += '"'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "sensor_msgs/BatteryState"

    def getMD5(self):
        return "715c4769cacd76e4b679cc3ea4c347b4"

_struct_I = struct.Struct('<I')

