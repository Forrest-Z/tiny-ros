package diagnostic_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type DiagnosticArray struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_status []*DiagnosticStatus `json:"status"`
}

func NewDiagnosticArray() (*DiagnosticArray) {
    newDiagnosticArray := new(DiagnosticArray)
    newDiagnosticArray.Go_header = std_msgs.NewHeader()
    newDiagnosticArray.Go_status = []*DiagnosticStatus{}
    return newDiagnosticArray
}

func (self *DiagnosticArray) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_status = []*DiagnosticStatus{}
}

func (self *DiagnosticArray) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    length_status := len(self.Go_status)
    buff[offset + 0] = byte((length_status >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_status >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_status >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_status >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_status; i++ {
        offset += self.Go_status[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *DiagnosticArray) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_status := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_status |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_status |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_status |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_status = make([]*DiagnosticStatus, length_status)
    for i := 0; i < length_status; i++ {
        self.Go_status[i] = NewDiagnosticStatus()
    }
    for i := 0; i < length_status; i++ {
        offset += self.Go_status[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *DiagnosticArray) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length_status := len(self.Go_status)
    for i := 0; i < length_status; i++ {
        length += self.Go_status[i].Go_serializedLength()
    }
    return length
}

func (self *DiagnosticArray) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *DiagnosticArray) Go_getType() (string) { return "diagnostic_msgs/DiagnosticArray" }
func (self *DiagnosticArray) Go_getMD5() (string) { return "79a87210f85eb6afbd600eb2ba49dd85" }
func (self *DiagnosticArray) Go_getID() (uint32) { return 0 }
func (self *DiagnosticArray) Go_setID(id uint32) { }

