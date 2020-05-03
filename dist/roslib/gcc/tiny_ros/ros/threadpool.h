/*
 * File      : threadpool.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2020-04-11     Pinkie.Fu    initial version
 */

#ifndef TINYROS_THREADPOOL_H_
#define TINYROS_THREADPOOL_H_
#include <stdint.h>
#include <vector>
#include <utility>
#include <queue>
#include <thread>
#include <functional>
#include <mutex>
#include <condition_variable>

namespace tinyros {
class ThreadPool{
public:
  typedef std::function<void()> Task;

  ThreadPool(int init_threads = 1) {
    started_ = true;
    threads_.reserve(init_threads);
    for (int i = 0; i < init_threads; i++) {
      threads_.push_back(new std::thread(std::bind(&ThreadPool::thread_loop, this)));
    }
  }
  
  ~ThreadPool() {
    shutdown();
  }

  void shutdown() {
    {
      started_ = false;
      std::unique_lock<std::mutex> lock(mutex_);
      tasks_.clear();
      cond_.notify_all();
    }

    for (Threads::iterator it = threads_.begin(); it != threads_.end() ; ++it) {
      (*it)->join();
      delete *it;
    }
    threads_.clear();
  }
  
  void schedule(const Task& task) {
    if (started_) {
      std::unique_lock<std::mutex> lock(mutex_);
      tasks_.push_back(task);
      cond_.notify_one();
    }
  }

private:
  void thread_loop() {
    while(started_) {
      Task task = take();
      if(started_ && task) { 
        task();
      }
    }
  }
  
  Task take() {
    std::unique_lock<std::mutex> lock(mutex_);
    while(tasks_.empty() && started_) {
      cond_.wait(lock);
    }

    Task task = nullptr;
    if(started_ && !tasks_.empty()) {
      task = tasks_.front();
      tasks_.pop_front();
    }
    return task;
  }

  typedef std::vector<std::thread*> Threads;
  typedef std::deque<Task> Tasks;

  Threads threads_;
  Tasks tasks_;

  std::mutex mutex_;
  std::condition_variable cond_;
  bool started_;
};

}

#endif //TINYROS_THREADPOOL_H_

