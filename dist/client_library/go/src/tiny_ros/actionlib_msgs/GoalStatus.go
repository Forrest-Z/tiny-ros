package actionlib_msgs

import (
    "encoding/json"
)

func Go_PENDING() (uint8) { return  0    }
func Go_ACTIVE() (uint8) { return  1    }
func Go_PREEMPTED() (uint8) { return  2    }
func Go_SUCCEEDED() (uint8) { return  3    }
func Go_ABORTED() (uint8) { return  4    }
func Go_REJECTED() (uint8) { return  5    }
func Go_PREEMPTING() (uint8) { return  6    }
func Go_RECALLING() (uint8) { return  7    }
func Go_RECALLED() (uint8) { return  8    }
func Go_LOST() (uint8) { return  9    }

type GoalStatus struct {
    Go_goal_id *GoalID `json:"goal_id"`
    Go_status uint8 `json:"status"`
    Go_text string `json:"text"`
}

func NewGoalStatus() (*GoalStatus) {
    newGoalStatus := new(GoalStatus)
    newGoalStatus.Go_goal_id = NewGoalID()
    newGoalStatus.Go_status = 0
    newGoalStatus.Go_text = ""
    return newGoalStatus
}

func (self *GoalStatus) Go_initialize() {
    self.Go_goal_id = NewGoalID()
    self.Go_status = 0
    self.Go_text = ""
}

func (self *GoalStatus) Go_serialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_goal_id.Go_serialize(buff[offset:])
    buff[offset + 0] = byte((self.Go_status >> (8 * 0)) & 0xFF)
    offset += 1
    length_text := len(self.Go_text)
    buff[offset + 0] = byte((length_text >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_text >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_text >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_text >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_text)], self.Go_text)
    offset += length_text
    return offset
}

func (self *GoalStatus) Go_deserialize(buff []byte) (int) {
    offset := 0
    offset += self.Go_goal_id.Go_deserialize(buff[offset:])
    self.Go_status = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    length_text := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_text |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_text |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_text |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_text = string(buff[offset:(offset+length_text)])
    offset += length_text
    return offset
}

func (self *GoalStatus) Go_serializedLength() (int) {
    length := 0
    length += self.Go_goal_id.Go_serializedLength()
    length += 1
    length_text := len(self.Go_text)
    length += 4
    length += length_text
    return length
}

func (self *GoalStatus) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *GoalStatus) Go_getType() (string) { return "actionlib_msgs/GoalStatus" }
func (self *GoalStatus) Go_getMD5() (string) { return "086be35ea957e692de83fc3477e4ef0b" }
func (self *GoalStatus) Go_getID() (uint32) { return 0 }
func (self *GoalStatus) Go_setID(id uint32) { }

