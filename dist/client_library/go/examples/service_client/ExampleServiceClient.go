package main

import (
    "fmt"
    "time"
    "tiny_ros/tinyros"
    "tiny_ros/tinyros_hello"
)

func main() {
    tinyros.Go_init("GoExampleServiceClient", "127.0.0.1")
    
    client := tinyros.NewServiceClient("test_srv", tinyros_hello.NewTestRequest(), tinyros_hello.NewTestResponse())

    tinyros.Go_nh().Go_serviceClient(client)

    for {
        req := tinyros_hello.NewTestRequest()
        resp := tinyros_hello.NewTestResponse()
        req.Go_input = "hello world!"
        if client.Go_call(req, resp) {
            fmt.Println("Service responsed with\"", resp.Go_output, "\"")
        } else {
            fmt.Println("Service call failed.")
        }
        time.Sleep(time.Second)
    }
}
