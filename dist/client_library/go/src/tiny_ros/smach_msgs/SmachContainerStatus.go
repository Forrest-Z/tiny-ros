package smach_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type SmachContainerStatus struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_path string `json:"path"`
    Go_initial_states []string `json:"initial_states"`
    Go_active_states []string `json:"active_states"`
    Go_local_data string `json:"local_data"`
    Go_info string `json:"info"`
}

func NewSmachContainerStatus() (*SmachContainerStatus) {
    newSmachContainerStatus := new(SmachContainerStatus)
    newSmachContainerStatus.Go_header = std_msgs.NewHeader()
    newSmachContainerStatus.Go_path = ""
    newSmachContainerStatus.Go_initial_states = []string{}
    newSmachContainerStatus.Go_active_states = []string{}
    newSmachContainerStatus.Go_local_data = ""
    newSmachContainerStatus.Go_info = ""
    return newSmachContainerStatus
}

func (self *SmachContainerStatus) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_path = ""
    self.Go_initial_states = []string{}
    self.Go_active_states = []string{}
    self.Go_local_data = ""
    self.Go_info = ""
}

func (self *SmachContainerStatus) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_serialize(buff[offset:])
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
    length_active_states := len(self.Go_active_states)
    buff[offset + 0] = byte((length_active_states >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_active_states >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_active_states >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_active_states >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_active_states; i++ {
        length_active_statesi := len(self.Go_active_states[i])
        buff[offset + 0] = byte((length_active_statesi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_active_statesi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_active_statesi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_active_statesi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_active_statesi)], self.Go_active_states[i])
        offset += length_active_statesi
    }
    length_local_data := len(self.Go_local_data)
    buff[offset + 0] = byte((length_local_data >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_local_data >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_local_data >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_local_data >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_local_data)], self.Go_local_data)
    offset += length_local_data
    length_info := len(self.Go_info)
    buff[offset + 0] = byte((length_info >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_info >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_info >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_info >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_info)], self.Go_info)
    offset += length_info
    return offset
}

func (self *SmachContainerStatus) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
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
    length_active_states := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_active_states |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_active_states |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_active_states |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_active_states = make([]string, length_active_states)
    for i := 0; i < length_active_states; i++ {
        length_active_statesi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_active_statesi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_active_statesi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_active_statesi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_active_states[i] = string(buff[offset:(offset+length_active_statesi)])
        offset += length_active_statesi
    }
    length_local_data := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_local_data |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_local_data |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_local_data |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_local_data = string(buff[offset:(offset+length_local_data)])
    offset += length_local_data
    length_info := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_info |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_info |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_info |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_info = string(buff[offset:(offset+length_info)])
    offset += length_info
    return offset
}

func (self *SmachContainerStatus) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
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
    length += 4
    length_active_states := len(self.Go_active_states)
    for i := 0; i < length_active_states; i++ {
        length_active_statesi := len(self.Go_active_states[i])
        length += 4
        length += length_active_statesi
    }
    length_local_data := len(self.Go_local_data)
    length += 4
    length += length_local_data
    length_info := len(self.Go_info)
    length += 4
    length += length_info
    return length
}

func (self *SmachContainerStatus) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SmachContainerStatus) Go_getType() (string) { return "smach_msgs/SmachContainerStatus" }
func (self *SmachContainerStatus) Go_getMD5() (string) { return "3e74cf72da18311be27e7a5970ea6415" }
func (self *SmachContainerStatus) Go_getID() (uint32) { return 0 }
func (self *SmachContainerStatus) Go_setID(id uint32) { }

