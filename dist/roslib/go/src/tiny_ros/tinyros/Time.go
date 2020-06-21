package tinyros

import (
    "math"
    "time"
)

type Time struct {
    Go_sec uint32 `json:"sec"`
    Go_nsec uint32 `json:"nsec"`
}

func NewTime() (*Time) {
    newTime := new(Time)
    newTime.Go_sec = 0
    newTime.Go_nsec = 0
    return newTime
}

func (self *Time) Go_toMSec() (float64) { 
    return float64(self.Go_sec) * 1000 + 1e-6 * float64(self.Go_nsec)
}

func (self *Time) Go_toSec() (float64) { 
    return float64(self.Go_sec)  + 1e-9 * float64(self.Go_nsec)
}

func (self *Time) Go_fromSec(t float64) {
    self.Go_sec = uint32(math.Floor(t))
    self.Go_nsec = uint32(math.Floor((t - float64(self.Go_sec)) * 1e9 + 0.5))
}

func (self *Time) Go_toNsec() (uint64) {
    return uint64(self.Go_sec * 1e9 + self.Go_nsec)
}

func TimeNow() (*Time) {
    now := time.Now().UnixNano()
    newTime := new(Time)
    newTime.Go_sec = uint32(now / 1e9)
    newTime.Go_nsec = uint32(now % 1e9)
    return newTime
}
