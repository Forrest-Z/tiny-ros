package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)

func Go_POWER_SUPPLY_STATUS_UNKNOWN() (uint8) { return  0 }
func Go_POWER_SUPPLY_STATUS_CHARGING() (uint8) { return  1 }
func Go_POWER_SUPPLY_STATUS_DISCHARGING() (uint8) { return  2 }
func Go_POWER_SUPPLY_STATUS_NOT_CHARGING() (uint8) { return  3 }
func Go_POWER_SUPPLY_STATUS_FULL() (uint8) { return  4 }
func Go_POWER_SUPPLY_HEALTH_UNKNOWN() (uint8) { return  0 }
func Go_POWER_SUPPLY_HEALTH_GOOD() (uint8) { return  1 }
func Go_POWER_SUPPLY_HEALTH_OVERHEAT() (uint8) { return  2 }
func Go_POWER_SUPPLY_HEALTH_DEAD() (uint8) { return  3 }
func Go_POWER_SUPPLY_HEALTH_OVERVOLTAGE() (uint8) { return  4 }
func Go_POWER_SUPPLY_HEALTH_UNSPEC_FAILURE() (uint8) { return  5 }
func Go_POWER_SUPPLY_HEALTH_COLD() (uint8) { return  6 }
func Go_POWER_SUPPLY_HEALTH_WATCHDOG_TIMER_EXPIRE() (uint8) { return  7 }
func Go_POWER_SUPPLY_HEALTH_SAFETY_TIMER_EXPIRE() (uint8) { return  8 }
func Go_POWER_SUPPLY_TECHNOLOGY_UNKNOWN() (uint8) { return  0 }
func Go_POWER_SUPPLY_TECHNOLOGY_NIMH() (uint8) { return  1 }
func Go_POWER_SUPPLY_TECHNOLOGY_LION() (uint8) { return  2 }
func Go_POWER_SUPPLY_TECHNOLOGY_LIPO() (uint8) { return  3 }
func Go_POWER_SUPPLY_TECHNOLOGY_LIFE() (uint8) { return  4 }
func Go_POWER_SUPPLY_TECHNOLOGY_NICD() (uint8) { return  5 }
func Go_POWER_SUPPLY_TECHNOLOGY_LIMN() (uint8) { return  6 }

type BatteryState struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_voltage float32 `json:"voltage"`
    Go_current float32 `json:"current"`
    Go_charge float32 `json:"charge"`
    Go_capacity float32 `json:"capacity"`
    Go_design_capacity float32 `json:"design_capacity"`
    Go_percentage float32 `json:"percentage"`
    Go_power_supply_status uint8 `json:"power_supply_status"`
    Go_power_supply_health uint8 `json:"power_supply_health"`
    Go_power_supply_technology uint8 `json:"power_supply_technology"`
    Go_present bool `json:"present"`
    Go_cell_voltage []float32 `json:"cell_voltage"`
    Go_location string `json:"location"`
    Go_serial_number string `json:"serial_number"`
}

func NewBatteryState() (*BatteryState) {
    newBatteryState := new(BatteryState)
    newBatteryState.Go_header = std_msgs.NewHeader()
    newBatteryState.Go_voltage = 0.0
    newBatteryState.Go_current = 0.0
    newBatteryState.Go_charge = 0.0
    newBatteryState.Go_capacity = 0.0
    newBatteryState.Go_design_capacity = 0.0
    newBatteryState.Go_percentage = 0.0
    newBatteryState.Go_power_supply_status = 0
    newBatteryState.Go_power_supply_health = 0
    newBatteryState.Go_power_supply_technology = 0
    newBatteryState.Go_present = false
    newBatteryState.Go_cell_voltage = []float32{}
    newBatteryState.Go_location = ""
    newBatteryState.Go_serial_number = ""
    return newBatteryState
}

func (self *BatteryState) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_voltage = 0.0
    self.Go_current = 0.0
    self.Go_charge = 0.0
    self.Go_capacity = 0.0
    self.Go_design_capacity = 0.0
    self.Go_percentage = 0.0
    self.Go_power_supply_status = 0
    self.Go_power_supply_health = 0
    self.Go_power_supply_technology = 0
    self.Go_present = false
    self.Go_cell_voltage = []float32{}
    self.Go_location = ""
    self.Go_serial_number = ""
}

func (self *BatteryState) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    bits_voltage := math.Float32bits(self.Go_voltage)
    binary.LittleEndian.PutUint32(buff[offset:], bits_voltage)
    offset += 4
    bits_current := math.Float32bits(self.Go_current)
    binary.LittleEndian.PutUint32(buff[offset:], bits_current)
    offset += 4
    bits_charge := math.Float32bits(self.Go_charge)
    binary.LittleEndian.PutUint32(buff[offset:], bits_charge)
    offset += 4
    bits_capacity := math.Float32bits(self.Go_capacity)
    binary.LittleEndian.PutUint32(buff[offset:], bits_capacity)
    offset += 4
    bits_design_capacity := math.Float32bits(self.Go_design_capacity)
    binary.LittleEndian.PutUint32(buff[offset:], bits_design_capacity)
    offset += 4
    bits_percentage := math.Float32bits(self.Go_percentage)
    binary.LittleEndian.PutUint32(buff[offset:], bits_percentage)
    offset += 4
    buff[offset + 0] = byte((self.Go_power_supply_status >> (8 * 0)) & 0xFF)
    offset += 1
    buff[offset + 0] = byte((self.Go_power_supply_health >> (8 * 0)) & 0xFF)
    offset += 1
    buff[offset + 0] = byte((self.Go_power_supply_technology >> (8 * 0)) & 0xFF)
    offset += 1
    if self.Go_present {
        buff[offset] = byte(0x01)
    } else {
        buff[offset] = byte(0x00)
    }
    offset += 1
    length_cell_voltage := len(self.Go_cell_voltage)
    buff[offset + 0] = byte((length_cell_voltage >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_cell_voltage >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_cell_voltage >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_cell_voltage >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_cell_voltage; i++ {
        bits_cell_voltagei := math.Float32bits(self.Go_cell_voltage[i])
        binary.LittleEndian.PutUint32(buff[offset:], bits_cell_voltagei)
        offset += 4
    }
    length_location := len(self.Go_location)
    buff[offset + 0] = byte((length_location >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_location >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_location >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_location >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_location)], self.Go_location)
    offset += length_location
    length_serial_number := len(self.Go_serial_number)
    buff[offset + 0] = byte((length_serial_number >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_serial_number >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_serial_number >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_serial_number >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_serial_number)], self.Go_serial_number)
    offset += length_serial_number
    return offset
}

func (self *BatteryState) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    bits_voltage := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_voltage = math.Float32frombits(bits_voltage)
    offset += 4
    bits_current := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_current = math.Float32frombits(bits_current)
    offset += 4
    bits_charge := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_charge = math.Float32frombits(bits_charge)
    offset += 4
    bits_capacity := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_capacity = math.Float32frombits(bits_capacity)
    offset += 4
    bits_design_capacity := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_design_capacity = math.Float32frombits(bits_design_capacity)
    offset += 4
    bits_percentage := binary.LittleEndian.Uint32(buff[offset:])
    self.Go_percentage = math.Float32frombits(bits_percentage)
    offset += 4
    self.Go_power_supply_status = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    self.Go_power_supply_health = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    self.Go_power_supply_technology = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    if (buff[offset] & 0xFF) != 0 {
        self.Go_present = true
    } else {
        self.Go_present = false
    }
    offset += 1
    length_cell_voltage := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_cell_voltage |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_cell_voltage |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_cell_voltage |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_cell_voltage = make([]float32, length_cell_voltage)
    for i := 0; i < length_cell_voltage; i++ {
        bits_cell_voltagei := binary.LittleEndian.Uint32(buff[offset:])
        self.Go_cell_voltage[i] = math.Float32frombits(bits_cell_voltagei)
        offset += 4
    }
    length_location := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_location |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_location |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_location |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_location = string(buff[offset:(offset+length_location)])
    offset += length_location
    length_serial_number := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_serial_number |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_serial_number |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_serial_number |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_serial_number = string(buff[offset:(offset+length_serial_number)])
    offset += length_serial_number
    return offset
}

func (self *BatteryState) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
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
    length += 4
    length_cell_voltage := len(self.Go_cell_voltage)
    for i := 0; i < length_cell_voltage; i++ {
        length += 4
    }
    length_location := len(self.Go_location)
    length += 4
    length += length_location
    length_serial_number := len(self.Go_serial_number)
    length += 4
    length += length_serial_number
    return length
}

func (self *BatteryState) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *BatteryState) Go_getType() (string) { return "sensor_msgs/BatteryState" }
func (self *BatteryState) Go_getMD5() (string) { return "715c4769cacd76e4b679cc3ea4c347b4" }
func (self *BatteryState) Go_getID() (uint32) { return 0 }
func (self *BatteryState) Go_setID(id uint32) { }

