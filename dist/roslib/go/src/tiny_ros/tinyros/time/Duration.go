package rostime

import (
    "math"
)

type Duration struct {
    Go_sec uint32 `json:"sec"`
    Go_nsec uint32 `json:"nsec"`
}

func NewDuration() (*Duration) {
    newDuration := new(Duration)
    newDuration.Go_sec = 0
    newDuration.Go_nsec = 0
    return newDuration
}

func (self *Duration) Go_toMSec() (float64) { 
    return float64(self.Go_sec) * 1000 + 1e-6 * float64(self.Go_nsec)
}

func (self *Duration) Go_toSec() (float64) { 
    return float64(self.Go_sec)  + 1e-9 * float64(self.Go_nsec)
}

func (self *Duration) Go_fromSec(t float64) {
    self.Go_sec = uint32(math.Floor(t))
    self.Go_nsec = uint32(math.Floor((t - float64(self.Go_sec)) * 1e9 + 0.5))
}

func (self *Duration) Go_toNsec() (uint64) {
    return uint64(self.Go_sec * 1e9 + self.Go_nsec)
}
