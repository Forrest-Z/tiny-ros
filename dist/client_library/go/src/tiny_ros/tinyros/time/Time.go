package rostime

import (
    "sync"
    "math"
    "time"
)

type Time struct {
    Go_sec uint32 `json:"sec"`
    Go_nsec uint32 `json:"nsec"`
}

var (
    Go_time_start int64 = 0
    Go_time_dds int64 = 0
    Go_time_last int64 = 0
    Go_time_mutex sync.Mutex
)

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

func TimeDDS() (*Time) {
    Go_time_mutex.Lock()
    t := TimeNow()
    var offset int64 = int64(t.Go_toMSec())
    if (offset > Go_time_start) && (Go_time_start > 0) {
        offset = offset - Go_time_start
    } else {
        offset = 0
    }
    t.Go_sec = uint32(offset / 1000)
    t.Go_nsec = uint32((offset % 1000) * 1000000)
    t.Go_sec += uint32(Go_time_dds / 1000)
    t.Go_nsec += uint32((Go_time_dds % 1000) * 1000000)
    t.Go_sec += uint32(t.Go_nsec / 1e9)
    t.Go_nsec = uint32(t.Go_nsec % 1e9)
    Go_time_mutex.Unlock()
    return t
}

func TimeNow() (*Time) {
    now := time.Now().UnixNano()
    newTime := new(Time)
    newTime.Go_sec = uint32(now / 1e9)
    newTime.Go_nsec = uint32(now % 1e9)
    return newTime
}

