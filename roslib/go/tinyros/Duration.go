package tinyros

import (
    "math"
)

type Dutaion struct {
    Go_sec uint32 `json:"sec"`
    Go_nsec uint32 `json:"nsec"`
}

func NewDutaion() (*Dutaion) {
    newDutaion := new(Dutaion)
    newDutaion.Go_sec = 0
    newDutaion.Go_nsec = 0
    return newDutaion
}

func (self *Dutaion) Go_toMSec() (float64) { 
    return float64(self.Go_sec) * 1000 + 1e-6 * float64(self.Go_nsec)
}

func (self *Dutaion) Go_toSec() (float64) { 
    return float64(self.Go_sec)  + 1e-9 * float64(self.Go_nsec)
}

func (self *Dutaion) Go_fromSec(t float64) {
    self.Go_sec = uint32(math.Floor(t))
    self.Go_nsec = uint32(math.Floor((t - float64(self.Go_sec)) * 1e9 + 0.5))
}

func (self *Dutaion) Go_toNsec() (uint64) {
    return uint64(self.Go_sec * 1e9 + self.Go_nsec)
}
