package trajectory_msgs

import (
    "encoding/json"
    "encoding/binary"
    "math"
    "tiny_ros/tinyros/time"
)


type JointTrajectoryPoint struct {
    Go_positions []float64 `json:"positions"`
    Go_velocities []float64 `json:"velocities"`
    Go_accelerations []float64 `json:"accelerations"`
    Go_effort []float64 `json:"effort"`
    Go_time_from_start *rostime.Duration `json:"time_from_start"`
}

func NewJointTrajectoryPoint() (*JointTrajectoryPoint) {
    newJointTrajectoryPoint := new(JointTrajectoryPoint)
    newJointTrajectoryPoint.Go_positions = []float64{}
    newJointTrajectoryPoint.Go_velocities = []float64{}
    newJointTrajectoryPoint.Go_accelerations = []float64{}
    newJointTrajectoryPoint.Go_effort = []float64{}
    newJointTrajectoryPoint.Go_time_from_start = rostime.NewDuration()
    return newJointTrajectoryPoint
}

func (self *JointTrajectoryPoint) Go_initialize() {
    self.Go_positions = []float64{}
    self.Go_velocities = []float64{}
    self.Go_accelerations = []float64{}
    self.Go_effort = []float64{}
    self.Go_time_from_start = rostime.NewDuration()
}

func (self *JointTrajectoryPoint) Go_serialize(buff []byte) (int) {
    offset := 0
    length_positions := len(self.Go_positions)
    buff[offset + 0] = byte((length_positions >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_positions >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_positions >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_positions >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_positions; i++ {
        bits_positionsi := math.Float64bits(self.Go_positions[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_positionsi)
        offset += 8
    }
    length_velocities := len(self.Go_velocities)
    buff[offset + 0] = byte((length_velocities >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_velocities >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_velocities >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_velocities >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_velocities; i++ {
        bits_velocitiesi := math.Float64bits(self.Go_velocities[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_velocitiesi)
        offset += 8
    }
    length_accelerations := len(self.Go_accelerations)
    buff[offset + 0] = byte((length_accelerations >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_accelerations >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_accelerations >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_accelerations >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_accelerations; i++ {
        bits_accelerationsi := math.Float64bits(self.Go_accelerations[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_accelerationsi)
        offset += 8
    }
    length_effort := len(self.Go_effort)
    buff[offset + 0] = byte((length_effort >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((length_effort >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((length_effort >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((length_effort >> (8 * 3)) & 0xFF)
    offset += 4
    for i := 0; i < length_effort; i++ {
        bits_efforti := math.Float64bits(self.Go_effort[i])
        binary.LittleEndian.PutUint64(buff[offset:], bits_efforti)
        offset += 8
    }
    buff[offset + 0] = byte((self.Go_time_from_start.Go_sec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_time_from_start.Go_sec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_time_from_start.Go_sec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_time_from_start.Go_sec >> (8 * 3)) & 0xFF)
    offset += 4
    buff[offset + 0] = byte((self.Go_time_from_start.Go_nsec >> (8 * 0)) & 0xFF)
    buff[offset + 1] = byte((self.Go_time_from_start.Go_nsec >> (8 * 1)) & 0xFF)
    buff[offset + 2] = byte((self.Go_time_from_start.Go_nsec >> (8 * 2)) & 0xFF)
    buff[offset + 3] = byte((self.Go_time_from_start.Go_nsec >> (8 * 3)) & 0xFF)
    offset += 4
    return offset
}

func (self *JointTrajectoryPoint) Go_deserialize(buff []byte) (int) {
    offset := 0
    length_positions := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_positions |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_positions |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_positions |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_positions = make([]float64, length_positions)
    for i := 0; i < length_positions; i++ {
        bits_positionsi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_positions[i] = math.Float64frombits(bits_positionsi)
        offset += 8
    }
    length_velocities := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_velocities |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_velocities |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_velocities |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_velocities = make([]float64, length_velocities)
    for i := 0; i < length_velocities; i++ {
        bits_velocitiesi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_velocities[i] = math.Float64frombits(bits_velocitiesi)
        offset += 8
    }
    length_accelerations := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_accelerations |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_accelerations |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_accelerations |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_accelerations = make([]float64, length_accelerations)
    for i := 0; i < length_accelerations; i++ {
        bits_accelerationsi := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_accelerations[i] = math.Float64frombits(bits_accelerationsi)
        offset += 8
    }
    length_effort := int(buff[offset + 0] & 0xFF) << (8 * 0)
    length_effort |= int(buff[offset + 1] & 0xFF) << (8 * 1)
    length_effort |= int(buff[offset + 2] & 0xFF) << (8 * 2)
    length_effort |= int(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_effort = make([]float64, length_effort)
    for i := 0; i < length_effort; i++ {
        bits_efforti := binary.LittleEndian.Uint64(buff[offset:])
        self.Go_effort[i] = math.Float64frombits(bits_efforti)
        offset += 8
    }
    self.Go_time_from_start.Go_sec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_time_from_start.Go_sec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_time_from_start.Go_sec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_time_from_start.Go_sec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    self.Go_time_from_start.Go_nsec = uint32(buff[offset + 0] & 0xFF) << (8 * 0)
    self.Go_time_from_start.Go_nsec |= uint32(buff[offset + 1] & 0xFF) << (8 * 1)
    self.Go_time_from_start.Go_nsec |= uint32(buff[offset + 2] & 0xFF) << (8 * 2)
    self.Go_time_from_start.Go_nsec |= uint32(buff[offset + 3] & 0xFF) << (8 * 3)
    offset += 4
    return offset
}

func (self *JointTrajectoryPoint) Go_serializedLength() (int) {
    length := 0
    length += 4
    length_positions := len(self.Go_positions)
    for i := 0; i < length_positions; i++ {
        length += 8
    }
    length += 4
    length_velocities := len(self.Go_velocities)
    for i := 0; i < length_velocities; i++ {
        length += 8
    }
    length += 4
    length_accelerations := len(self.Go_accelerations)
    for i := 0; i < length_accelerations; i++ {
        length += 8
    }
    length += 4
    length_effort := len(self.Go_effort)
    for i := 0; i < length_effort; i++ {
        length += 8
    }
    length += 4
    length += 4
    return length
}

func (self *JointTrajectoryPoint) Go_echo() (string) { 
    data, _ := json.Marshal(self)
    return string(data)
}

func (self *JointTrajectoryPoint) Go_getType() (string) { return "trajectory_msgs/JointTrajectoryPoint" }
func (self *JointTrajectoryPoint) Go_getMD5() (string) { return "38679319321341510f6fde7d7f745eb0" }
func (self *JointTrajectoryPoint) Go_getID() (uint32) { return 0 }
func (self *JointTrajectoryPoint) Go_setID(id uint32) { }

