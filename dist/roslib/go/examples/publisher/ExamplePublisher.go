package main

import (
    "time"
    "tiny_ros/tinyros"
    "tiny_ros/tinyros_hello"
)

func main() {
   nh := tinyros.Go_udp()
   for {
       msg := new(tinyros_hello.TinyrosHello)
       msg.Go_hello = "Go Hello, tiny-ros ^_^ "
       nh.Go_publish(3364715186, msg)
       time.Sleep(time.Second)
   }
}

