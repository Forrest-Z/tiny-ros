package gazebo_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type ContactsState struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_states []*ContactState `json:"states"`
}

func NewContactsState() (*ContactsState) {
    newContactsState := new(ContactsState)
    newContactsState.Go_header = std_msgs.NewHeader()
    newContactsState.Go_states = []*ContactState{}
    return newContactsState
}

func (self *ContactsState) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_states = []*ContactState{}
}

func (self *ContactsState) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
    length_states := len(self.Go_states)
    buff[offset + 0] = byte((length_states >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_states >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_states >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_states >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_states; i++ {
        offset += self.Go_states[i].Go_serialize(buff[offset:])
    }
    return offset
}

func (self *ContactsState) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_states := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_states |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_states |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_states |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_states = make([]*ContactState, length_states)
    for i := 0; i < length_states; i++ {
        self.Go_states[i] = NewContactState()
    }
    for i := 0; i < length_states; i++ {
        offset += self.Go_states[i].Go_deserialize(buff[offset:])
    }
    return offset
}

func (self *ContactsState) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length += 4
    length_states := len(self.Go_states)
    for i := 0; i < length_states; i++ {
        length += self.Go_states[i].Go_serializedLength()
    }
    return length
}

func (self *ContactsState) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *ContactsState) Go_getType() (string) { return "gazebo_msgs/ContactsState" }
func (self *ContactsState) Go_getMD5() (string) { return "d19cd2a086cbd43da4252eb8d5cc64f5" }
func (self *ContactsState) Go_getID() (uint32) { return 0 }
func (self *ContactsState) Go_setID(id uint32) { }

