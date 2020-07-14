package diagnostic_msgs

import (
    "encoding/json"
)

func Go_OK() (byte) { return 0 }
func Go_WARN() (byte) { return 1 }
func Go_ERROR() (byte) { return 2 }
func Go_STALE() (byte) { return 3 }

type DiagnosticStatus struct {
    Go_level byte `json:"level"`
    Go_name string `json:"name"`
    Go_message string `json:"message"`
    Go_hardware_id string `json:"hardware_id"`
    Go_values []*KeyValue `json:"values"`
}

func NewDiagnosticStatus() (*DiagnosticStatus) {
    newDiagnosticStatus := new(DiagnosticStatus)
    newDiagnosticStatus.Go_level = 0
    newDiagnosticStatus.Go_name = ""
    newDiagnosticStatus.Go_message = ""
    newDiagnosticStatus.Go_hardware_id = ""
    newDiagnosticStatus.Go_values = []*KeyValue{}
    return newDiagnosticStatus
}

func (self *DiagnosticStatus) Go_initialize() {
    self.Go_level = 0
    self.Go_name = ""
    self.Go_message = ""
    self.Go_hardware_id = ""
    self.Go_values = []*KeyValue{}
}

func (self *DiagnosticStatus) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_level >> (8 * 0)) & 0xFF)
    offset += 1
    length_name := len(self.Go_name)
    buff[offset + 0] = byte((length_name >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_name >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_name >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_name >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_name)], self.Go_name)
    offset += length_name
    length_message := len(self.Go_message)
    buff[offset + 0] = byte((length_message >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_message >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_message >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_message >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_message)], self.Go_message)
    offset += length_message
    length_hardware_id := len(self.Go_hardware_id)
    buff[offset + 0] = byte((length_hardware_id >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_hardware_id >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_hardware_id >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_hardware_id >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_hardware_id)], self.Go_hardware_id)
    offset += length_hardware_id
    length_values := len(self.Go_values)
    buff[offset + 0] = byte((length_values >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_values >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_values >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_values >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_values; i++ {
        offset += self.Go_values[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *DiagnosticStatus) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_level = byte(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    length_name := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_name |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_name |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_name |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_name = string(buff[offset:(offset+length_name)])
    offset += length_name
    length_message := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_message |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_message |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_message |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_message = string(buff[offset:(offset+length_message)])
    offset += length_message
    length_hardware_id := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_hardware_id |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_hardware_id |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_hardware_id |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_hardware_id = string(buff[offset:(offset+length_hardware_id)])
    offset += length_hardware_id
    length_values := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_values |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_values |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_values |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_values = make([]*KeyValue, length_values)
    for i := 0; i < length_values; i++ {
        self.Go_values[i] = NewKeyValue()
    }
    for i := 0; i < length_values; i++ {
        offset += self.Go_values[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *DiagnosticStatus) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_name := len(self.Go_name)
    length += 4
    length += length_name
    length_message := len(self.Go_message)
    length += 4
    length += length_message
    length_hardware_id := len(self.Go_hardware_id)
    length += 4
    length += length_hardware_id
    length += 4
    length_values := len(self.Go_values)
    for i := 0; i < length_values; i++ {
        length += self.Go_values[i].Go_serializedLength()
    }
    return length
}

func (self *DiagnosticStatus) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *DiagnosticStatus) Go_getType() (string) { return "diagnostic_msgs/DiagnosticStatus" }
func (self *DiagnosticStatus) Go_getMD5() (string) { return "9ec892d2145f478061efd60bb1762361" }
func (self *DiagnosticStatus) Go_getID() (uint32) { return 0 }
func (self *DiagnosticStatus) Go_setID(id uint32) { }

