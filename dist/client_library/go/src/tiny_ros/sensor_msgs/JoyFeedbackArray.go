package sensor_msgs

import (
    "encoding/json"
)


type JoyFeedbackArray struct {
    Go_array []*JoyFeedback `json:"array"`
}

func NewJoyFeedbackArray() (*JoyFeedbackArray) {
    newJoyFeedbackArray := new(JoyFeedbackArray)
    newJoyFeedbackArray.Go_array = []*JoyFeedback{}
    return newJoyFeedbackArray
}

func (self *JoyFeedbackArray) Go_initialize() {
    self.Go_array = []*JoyFeedback{}
}

func (self *JoyFeedbackArray) Go_serialize(buff []byte) (int) {
    offset := 0
    length_array := len(self.Go_array)
    buff[offset + 0] = byte((length_array >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_array >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_array >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_array >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_array; i++ {
        offset += self.Go_array[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *JoyFeedbackArray) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_array := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_array |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_array |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_array |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_array = make([]*JoyFeedback, length_array)
    for i := 0; i < length_array; i++ {
        self.Go_array[i] = NewJoyFeedback()
    }
    for i := 0; i < length_array; i++ {
        offset += self.Go_array[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *JoyFeedbackArray) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_array := len(self.Go_array)
    for i := 0; i < length_array; i++ {
        length += self.Go_array[i].Go_serializedLength()
    }
    return length
}

func (self *JoyFeedbackArray) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *JoyFeedbackArray) Go_getType() (string) { return "sensor_msgs/JoyFeedbackArray" }
func (self *JoyFeedbackArray) Go_getMD5() (string) { return "45361e458d526d5670706a9f083819b6" }
func (self *JoyFeedbackArray) Go_getID() (uint32) { return 0 }
func (self *JoyFeedbackArray) Go_setID(id uint32) { }

