package sensor_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
    "encoding/binary"
    "math"
)


type FluidPressure struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_fluid_pressure float64 `json:"fluid_pressure"`
    Go_variance float64 `json:"variance"`
}

func NewFluidPressure() (*FluidPressure) {
    newFluidPressure := new(FluidPressure)
    newFluidPressure.Go_header = std_msgs.NewHeader()
    newFluidPressure.Go_fluid_pressure = 0.0
    newFluidPressure.Go_variance = 0.0
    return newFluidPressure
}

func (self *FluidPressure) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_fluid_pressure = 0.0
    self.Go_variance = 0.0
}

func (self *FluidPressure) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    bits_fluid_pressure := math.Float64bits(self.Go_fluid_pressure)
    binary.LittleEndian.PutUint64(buff[offset:], bits_fluid_pressure)
    offset += 8
    bits_variance := math.Float64bits(self.Go_variance)
    binary.LittleEndian.PutUint64(buff[offset:], bits_variance)
    offset += 8
    return offset
}

func (self *FluidPressure) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    bits_fluid_pressure := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_fluid_pressure = math.Float64frombits(bits_fluid_pressure)
    offset += 8
    bits_variance := binary.LittleEndian.Uint64(buff[offset:])
    self.Go_variance = math.Float64frombits(bits_variance)
    offset += 8
    return offset
}

func (self *FluidPressure) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 8
    length += 8
    return length
}

func (self *FluidPressure) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *FluidPressure) Go_getType() (string) { return "sensor_msgs/FluidPressure" }
func (self *FluidPressure) Go_getMD5() (string) { return "0fdea137019d78ebf8c2cb91c31a458a" }
func (self *FluidPressure) Go_getID() (uint32) { return 0 }
func (self *FluidPressure) Go_setID(id uint32) { }

