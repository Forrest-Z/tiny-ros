package smach_msgs

import (
    "encoding/json"
)


type SmachContainerInitialStatusCmd struct {
    Go_path string `json:"path"`
    Go_initial_states []string `json:"initial_states"`
    Go_local_data string `json:"local_data"`
}

func NewSmachContainerInitialStatusCmd() (*SmachContainerInitialStatusCmd) {
    newSmachContainerInitialStatusCmd := new(SmachContainerInitialStatusCmd)
    newSmachContainerInitialStatusCmd.Go_path = ""
    newSmachContainerInitialStatusCmd.Go_initial_states = []string{}
    newSmachContainerInitialStatusCmd.Go_local_data = ""
    return newSmachContainerInitialStatusCmd
}

func (self *SmachContainerInitialStatusCmd) Go_initialize() {
    self.Go_path = ""
    self.Go_initial_states = []string{}
    self.Go_local_data = ""
}

func (self *SmachContainerInitialStatusCmd) Go_serialize(buff []byte) (int) {
    offset := 0
    length_path := len(self.Go_path)
    buff[offset + 0] = byte((length_path >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_path >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_path >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_path >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_path)], self.Go_path)
    offset += length_path
    length_initial_states := len(self.Go_initial_states)
    buff[offset + 0] = byte((length_initial_states >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_initial_states >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_initial_states >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_initial_states >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_initial_states; i++ {
        length_initial_statesi := len(self.Go_initial_states[i])
        buff[offset + 0] = byte((length_initial_statesi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_initial_statesi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_initial_statesi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_initial_statesi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_initial_statesi)], self.Go_initial_states[i])
        offset += length_initial_statesi
    }
    length_local_data := len(self.Go_local_data)
    buff[offset + 0] = byte((length_local_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_local_data >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_local_data >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_local_data >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_local_data)], self.Go_local_data)
    offset += length_local_data
    return offset
}

func (self *SmachContainerInitialStatusCmd) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_path := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_path |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_path |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_path |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_path = string(buff[offset:(offset+length_path)])
    offset += length_path
    length_initial_states := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_initial_states |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_initial_states |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_initial_states |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_initial_states = make([]string, length_initial_states)
    for i := 0; i < length_initial_states; i++ {
        length_initial_statesi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_initial_statesi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_initial_statesi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_initial_statesi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_initial_states[i] = string(buff[offset:(offset+length_initial_statesi)])
        offset += length_initial_statesi
    }
    length_local_data := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_local_data |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_local_data |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_local_data |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_local_data = string(buff[offset:(offset+length_local_data)])
    offset += length_local_data
    return offset
}

func (self *SmachContainerInitialStatusCmd) Go_serializedLength() (int) {
    length := 0
    length_path := len(self.Go_path)
    length += 4
    length += length_path
    length += 4
    length_initial_states := len(self.Go_initial_states)
    for i := 0; i < length_initial_states; i++ {
        length_initial_statesi := len(self.Go_initial_states[i])
        length += 4
        length += length_initial_statesi
    }
    length_local_data := len(self.Go_local_data)
    length += 4
    length += length_local_data
    return length
}

func (self *SmachContainerInitialStatusCmd) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SmachContainerInitialStatusCmd) Go_getType() (string) { return "smach_msgs/SmachContainerInitialStatusCmd" }
func (self *SmachContainerInitialStatusCmd) Go_getMD5() (string) { return "b7c8a2bbd4d7c89f80561645ea1f4f13" }
func (self *SmachContainerInitialStatusCmd) Go_getID() (uint32) { return 0 }
func (self *SmachContainerInitialStatusCmd) Go_setID(id uint32) { }

