import sys
import threading
import tinyros

class ThreadPool(object):
    __slots__ = ['threads_','cond_','started_','tasks_']

    def __init__(self, init_threads = 1):
        self.tasks_ = []
        self.threads_ = []
        self.cond_ = threading.Condition()
        self.started_ = True
        for i in range(0, init_threads):
            _thread = threading.Thread(target=self.thread_loop)
            _thread.start()
            self.threads_.append(_thread)

    def thread_loop(self): 
        while self.started_:
            self.cond_.acquire()
            while len(self.tasks_) <= 0 and self.started_:
                self.cond_.wait()
            _task = None
            if len(self.tasks_) > 0 and self.started_:
                _task = self.tasks_.pop(0)
            self.cond_.release()
            if _task != None:
                _task.sub.callback(_task.message_in)
    
    def shutdown(self):
        self.started_ = False
        self.cond_.acquire()
        self.tasks_[:] = []
        self.cond_.notify_all()
        self.cond_.release()
        for x in self.threads_:
            x.join()
        self.threads_[:] = []
  
    def schedule(self, task):
        if self.started_:
            self.cond_.acquire()
            self.tasks_.append(task)
            self.cond_.notify_all()
            self.cond_.release()

