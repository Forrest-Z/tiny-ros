package main

import (
    "fmt"
    "time"
    "tiny_ros/tinyros"
    "tiny_ros/tinyros_hello"
)

func subscriber_cb(msg tinyros.Msg) {
    tmsg := msg.(*tinyros_hello.TinyrosHello)
    fmt.Println(tmsg.Go_hello)
}

func main() {
    tinyros.Go_init("GoExampleSubscriber", "127.0.0.1")

    sub := tinyros.NewSubscriber("tinyros_hello", subscriber_cb, tinyros_hello.NewTinyrosHello())

    if true {
        tinyros.Go_nh().Go_subscribe(sub)
    } else {
        tinyros.Go_udp().Go_subscribe(sub)
    }
    
    for {
        time.Sleep(10 * time.Second)
    }
}
