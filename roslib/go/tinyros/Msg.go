package tinyros

type Msg interface {
    Go_initialize()
    Go_serialize(buff []byte) (int)
    Go_deserialize(buff []byte) (int)
    Go_serializedLength() (int)
    Go_getType() (string)
    Go_getMD5() (string)
    Go_echo() (string)
    Go_getID() (uint32)
    Go_setID(id uint32)
}

