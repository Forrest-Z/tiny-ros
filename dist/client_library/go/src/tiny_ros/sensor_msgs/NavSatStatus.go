package sensor_msgs

import (
    "encoding/json"
)

func Go_STATUS_NO_FIX() (int8) { return   -1         }
func Go_STATUS_FIX() (int8) { return       0         }
func Go_STATUS_SBAS_FIX() (int8) { return  1         }
func Go_STATUS_GBAS_FIX() (int8) { return  2         }
func Go_SERVICE_GPS() (uint16) { return      1 }
func Go_SERVICE_GLONASS() (uint16) { return  2 }
func Go_SERVICE_COMPASS() (uint16) { return  4       }
func Go_SERVICE_GALILEO() (uint16) { return  8 }

type NavSatStatus struct {
    Go_status int8 `json:"status"`
    Go_service uint16 `json:"service"`
}

func NewNavSatStatus() (*NavSatStatus) {
    newNavSatStatus := new(NavSatStatus)
    newNavSatStatus.Go_status = 0
    newNavSatStatus.Go_service = 0
    return newNavSatStatus
}

func (self *NavSatStatus) Go_initialize() {
    self.Go_status = 0
    self.Go_service = 0
}

func (self *NavSatStatus) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte(uint8(self.Go_status >> (8 * 0)) & 0xFF)
    offset += 1
    buff[offset + 0] = byte((self.Go_service >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_service >> (8 * 1)) & 0xFF)
    offset += 2
    return offset
}

func (self *NavSatStatus) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_status = int8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    self.Go_service = uint16(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_service |= uint16(buff[offset + 1] & 0xFF) << (8 * 1)
    offset += 2
    return offset
}

func (self *NavSatStatus) Go_serializedLength() (int) {
    length := 0
    length += 1
    length += 2
    return length
}

func (self *NavSatStatus) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *NavSatStatus) Go_getType() (string) { return "sensor_msgs/NavSatStatus" }
func (self *NavSatStatus) Go_getMD5() (string) { return "85ed59cf80532c1c15a2cf735d06279b" }
func (self *NavSatStatus) Go_getID() (uint32) { return 0 }
func (self *NavSatStatus) Go_setID(id uint32) { }

