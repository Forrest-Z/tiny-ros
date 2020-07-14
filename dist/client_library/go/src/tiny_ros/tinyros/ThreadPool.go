package tinyros
        
type ThreadPool struct {
    jobs_ chan SpinObject `json:"jobs"`
    working_ bool `json:"working"`
}

func NewThreadPool(workers int) (*ThreadPool) {
    newThreadPool := new(ThreadPool)
    newThreadPool.jobs_ = make(chan SpinObject, 100)
    newThreadPool.working_ = true
    for i := 0; i < workers; i++ {
        go newThreadPool.worker()
    }
    return newThreadPool
}

func (self *ThreadPool) worker() {
    for self.working_ {
        j := <- self.jobs_
        if j.sub_ != nil && j.message_in_ != nil {
            j.sub_.Go_callback(j.message_in_)
        }
    }
}

func (self *ThreadPool) schedule(job SpinObject) {
    self.jobs_ <- job
}

func (self *ThreadPool) shutdown() {
    self.working_ = false
    close(self.jobs_)
}
