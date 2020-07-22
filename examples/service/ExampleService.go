package main

import (
    "time"
    "tiny_ros/tinyros"
    "tiny_ros/tinyros_hello"
)

func service_cb(req tinyros.Msg, resp tinyros.Msg) {
    tresp := resp.(*tinyros_hello.TestResponse)
    tresp.Go_output = "Hello, tiny-ros ^_^"
}

func main() {
    tinyros.Go_init("GoExampleService", "127.0.0.1")
    
    server := tinyros.NewServiceServer("test_srv", service_cb, tinyros_hello.NewTestRequest(), tinyros_hello.NewTestResponse())

    tinyros.Go_nh().Go_advertiseService(server)
    
    for {
        time.Sleep(10 * time.Second)
    }
}
