package tf2_msgs

import (
    "encoding/json"
)

func Go_NO_ERROR() (uint8) { return  0 }
func Go_LOOKUP_ERROR() (uint8) { return  1 }
func Go_CONNECTIVITY_ERROR() (uint8) { return  2 }
func Go_EXTRAPOLATION_ERROR() (uint8) { return  3 }
func Go_INVALID_ARGUMENT_ERROR() (uint8) { return  4 }
func Go_TIMEOUT_ERROR() (uint8) { return  5 }
func Go_TRANSFORM_ERROR() (uint8) { return  6 }

type TF2Error struct {
    Go_error uint8 `json:"error"`
    Go_error_string string `json:"error_string"`
}

func NewTF2Error() (*TF2Error) {
    newTF2Error := new(TF2Error)
    newTF2Error.Go_error = 0
    newTF2Error.Go_error_string = ""
    return newTF2Error
}

func (self *TF2Error) Go_initialize() {
    self.Go_error = 0
    self.Go_error_string = ""
}

func (self *TF2Error) Go_serialize(buff []byte) (int) {
    offset := 0
    buff[offset + 0] = byte((self.Go_error >> (8 * 0)) & 0xFF)
    offset += 1
    length_error_string := len(self.Go_error_string)
    buff[offset + 0] = byte((length_error_string >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_error_string >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_error_string >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_error_string >> (8 * 3)) & 0xFF)
    offset += 4
    copy(buff[offset:(offset+length_error_string)], self.Go_error_string)
    offset += length_error_string
    return offset
}

func (self *TF2Error) Go_deserialize(buff []byte) (int) {
    offset := 0
    self.Go_error = uint8(buff[offset + 0] & 0xFF) << (8 * 0)
    offset += 1
    length_error_string := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_error_string |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_error_string |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_error_string |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_error_string = string(buff[offset:(offset+length_error_string)])
    offset += length_error_string
    return offset
}

func (self *TF2Error) Go_serializedLength() (int) {
    length := 0
    length += 1
    length_error_string := len(self.Go_error_string)
    length += 4
    length += length_error_string
    return length
}

func (self *TF2Error) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *TF2Error) Go_getType() (string) { return "tf2_msgs/TF2Error" }
func (self *TF2Error) Go_getMD5() (string) { return "ed32adf5a372962d977aea0e5630d1d6" }
func (self *TF2Error) Go_getID() (uint32) { return 0 }
func (self *TF2Error) Go_setID(id uint32) { }

