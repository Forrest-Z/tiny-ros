package main

import (
    "time"
    "tiny_ros/tinyros"
    "tiny_ros/tinyros_hello"
)

func main() {
    tinyros.Go_init("GoExamplePublisher", "127.0.0.1")
    
    pub := tinyros.NewPublisher("tinyros_hello", tinyros_hello.NewTinyrosHello())
    
    if true {
        tinyros.Go_nh().Go_advertise(pub)
    } else {
        tinyros.Go_udp().Go_advertise(pub)
    }
    
    for {
        msg := tinyros_hello.NewTinyrosHello()
        msg.Go_hello = "Hello, tiny-ros ^_^ "
        pub.Go_publish(msg)
        time.Sleep(1 * time.Second)
    }
}
