package smach_msgs

import (
    "encoding/json"
    "tiny_ros/std_msgs"
)


type SmachContainerStructure struct {
    Go_header *std_msgs.Header `json:"header"`
    Go_path string `json:"path"`
    Go_children []string `json:"children"`
    Go_internal_outcomes []string `json:"internal_outcomes"`
    Go_outcomes_from []string `json:"outcomes_from"`
    Go_outcomes_to []string `json:"outcomes_to"`
    Go_container_outcomes []string `json:"container_outcomes"`
}

func NewSmachContainerStructure() (*SmachContainerStructure) {
    newSmachContainerStructure := new(SmachContainerStructure)
    newSmachContainerStructure.Go_header = std_msgs.NewHeader()
    newSmachContainerStructure.Go_path = ""
    newSmachContainerStructure.Go_children = []string{}
    newSmachContainerStructure.Go_internal_outcomes = []string{}
    newSmachContainerStructure.Go_outcomes_from = []string{}
    newSmachContainerStructure.Go_outcomes_to = []string{}
    newSmachContainerStructure.Go_container_outcomes = []string{}
    return newSmachContainerStructure
}

func (self *SmachContainerStructure) Go_initialize() {
    self.Go_header = std_msgs.NewHeader()
    self.Go_path = ""
    self.Go_children = []string{}
    self.Go_internal_outcomes = []string{}
    self.Go_outcomes_from = []string{}
    self.Go_outcomes_to = []string{}
    self.Go_container_outcomes = []string{}
}

func (self *SmachContainerStructure) Go_serialize(buff []byte) (int) {
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
    length_children := len(self.Go_children)
    buff[offset + 0] = byte((length_children >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_children >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_children >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_children >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_children; i++ {
        length_childreni := len(self.Go_children[i])
        buff[offset + 0] = byte((length_childreni >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_childreni >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_childreni >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_childreni >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_childreni)], self.Go_children[i])
        offset += length_childreni
    }
    length_internal_outcomes := len(self.Go_internal_outcomes)
    buff[offset + 0] = byte((length_internal_outcomes >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_internal_outcomes >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_internal_outcomes >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_internal_outcomes >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_internal_outcomes; i++ {
        length_internal_outcomesi := len(self.Go_internal_outcomes[i])
        buff[offset + 0] = byte((length_internal_outcomesi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_internal_outcomesi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_internal_outcomesi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_internal_outcomesi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_internal_outcomesi)], self.Go_internal_outcomes[i])
        offset += length_internal_outcomesi
    }
    length_outcomes_from := len(self.Go_outcomes_from)
    buff[offset + 0] = byte((length_outcomes_from >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_outcomes_from >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_outcomes_from >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_outcomes_from >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_outcomes_from; i++ {
        length_outcomes_fromi := len(self.Go_outcomes_from[i])
        buff[offset + 0] = byte((length_outcomes_fromi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_outcomes_fromi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_outcomes_fromi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_outcomes_fromi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_outcomes_fromi)], self.Go_outcomes_from[i])
        offset += length_outcomes_fromi
    }
    length_outcomes_to := len(self.Go_outcomes_to)
    buff[offset + 0] = byte((length_outcomes_to >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_outcomes_to >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_outcomes_to >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_outcomes_to >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_outcomes_to; i++ {
        length_outcomes_toi := len(self.Go_outcomes_to[i])
        buff[offset + 0] = byte((length_outcomes_toi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_outcomes_toi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_outcomes_toi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_outcomes_toi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_outcomes_toi)], self.Go_outcomes_to[i])
        offset += length_outcomes_toi
    }
    length_container_outcomes := len(self.Go_container_outcomes)
    buff[offset + 0] = byte((length_container_outcomes >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_container_outcomes >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_container_outcomes >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_container_outcomes >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_container_outcomes; i++ {
        length_container_outcomesi := len(self.Go_container_outcomes[i])
        buff[offset + 0] = byte((length_container_outcomesi >> (8 * 0)) & 0xFF)
        buff[offset + 1] = byte((length_container_outcomesi >> (8 * 1)) & 0xFF)
        buff[offset + 2] = byte((length_container_outcomesi >> (8 * 2)) & 0xFF)
        buff[offset + 3] = byte((length_container_outcomesi >> (8 * 3)) & 0xFF)
        offset += 4
        copy(buff[offset:(offset+length_container_outcomesi)], self.Go_container_outcomes[i])
        offset += length_container_outcomesi
    }
    return offset
}

func (self *SmachContainerStructure) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_header.Go_deserialize(buff[offset:])
    length_path := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_path |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_path |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_path |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_path = string(buff[offset:(offset+length_path)])
    offset += length_path
    length_children := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_children |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_children |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_children |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_children = make([]string, length_children)
    for i := 0; i < length_children; i++ {
        length_childreni := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_childreni |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_childreni |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_childreni |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_children[i] = string(buff[offset:(offset+length_childreni)])
        offset += length_childreni
    }
    length_internal_outcomes := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_internal_outcomes |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_internal_outcomes |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_internal_outcomes |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_internal_outcomes = make([]string, length_internal_outcomes)
    for i := 0; i < length_internal_outcomes; i++ {
        length_internal_outcomesi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_internal_outcomesi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_internal_outcomesi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_internal_outcomesi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_internal_outcomes[i] = string(buff[offset:(offset+length_internal_outcomesi)])
        offset += length_internal_outcomesi
    }
    length_outcomes_from := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_outcomes_from |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_outcomes_from |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_outcomes_from |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_outcomes_from = make([]string, length_outcomes_from)
    for i := 0; i < length_outcomes_from; i++ {
        length_outcomes_fromi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_outcomes_fromi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_outcomes_fromi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_outcomes_fromi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_outcomes_from[i] = string(buff[offset:(offset+length_outcomes_fromi)])
        offset += length_outcomes_fromi
    }
    length_outcomes_to := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_outcomes_to |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_outcomes_to |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_outcomes_to |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_outcomes_to = make([]string, length_outcomes_to)
    for i := 0; i < length_outcomes_to; i++ {
        length_outcomes_toi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_outcomes_toi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_outcomes_toi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_outcomes_toi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_outcomes_to[i] = string(buff[offset:(offset+length_outcomes_toi)])
        offset += length_outcomes_toi
    }
    length_container_outcomes := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_container_outcomes |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_container_outcomes |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_container_outcomes |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_container_outcomes = make([]string, length_container_outcomes)
    for i := 0; i < length_container_outcomes; i++ {
        length_container_outcomesi := int(buff[offset + 0] & 0xFF) << (8 * 0)
        length_container_outcomesi |= int(buff[offset + 1] & 0xFF) << (8 * 1)
        length_container_outcomesi |= int(buff[offset + 2] & 0xFF) << (8 * 2)
        length_container_outcomesi |= int(buff[offset + 3] & 0xFF) << (8 * 3)
        offset += 4
        self.Go_container_outcomes[i] = string(buff[offset:(offset+length_container_outcomesi)])
        offset += length_container_outcomesi
    }
    return offset
}

func (self *SmachContainerStructure) Go_serializedLength() (int) {
    length := 0
    length += self.Go_header.Go_serializedLength()
    length_path := len(self.Go_path)
    length += 4
    length += length_path
    length += 4
    length_children := len(self.Go_children)
    for i := 0; i < length_children; i++ {
        length_childreni := len(self.Go_children[i])
        length += 4
        length += length_childreni
    }
    length += 4
    length_internal_outcomes := len(self.Go_internal_outcomes)
    for i := 0; i < length_internal_outcomes; i++ {
        length_internal_outcomesi := len(self.Go_internal_outcomes[i])
        length += 4
        length += length_internal_outcomesi
    }
    length += 4
    length_outcomes_from := len(self.Go_outcomes_from)
    for i := 0; i < length_outcomes_from; i++ {
        length_outcomes_fromi := len(self.Go_outcomes_from[i])
        length += 4
        length += length_outcomes_fromi
    }
    length += 4
    length_outcomes_to := len(self.Go_outcomes_to)
    for i := 0; i < length_outcomes_to; i++ {
        length_outcomes_toi := len(self.Go_outcomes_to[i])
        length += 4
        length += length_outcomes_toi
    }
    length += 4
    length_container_outcomes := len(self.Go_container_outcomes)
    for i := 0; i < length_container_outcomes; i++ {
        length_container_outcomesi := len(self.Go_container_outcomes[i])
        length += 4
        length += length_container_outcomesi
    }
    return length
}

func (self *SmachContainerStructure) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *SmachContainerStructure) Go_getType() (string) { return "smach_msgs/SmachContainerStructure" }
func (self *SmachContainerStructure) Go_getMD5() (string) { return "0209663ab13753a56b6ac1d094d07fbe" }
func (self *SmachContainerStructure) Go_getID() (uint32) { return 0 }
func (self *SmachContainerStructure) Go_setID(id uint32) { }

