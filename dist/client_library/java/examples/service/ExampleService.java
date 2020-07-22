package examples.service;

import com.roslib.ros.CallbackSrvT;
import com.roslib.ros.Msg;
import com.roslib.ros.ServiceServer;
import com.roslib.ros.Tinyros;
import com.roslib.tinyros_hello.Test;

public class ExampleService {
    public static void main(String[] args) throws InterruptedException {
        Tinyros.init("JavaExampleService", "127.0.0.1");

        ServiceServer<Test.TestRequest, Test.TestResponse> srv = new ServiceServer<Test.TestRequest, Test.TestResponse>
        ("test_srv", new CallbackSrvT() {
            @Override
            public void callback(Msg req, Msg res) {
                ((Test.TestResponse)res).output = "Hello, tiny-ros ^_^";
            }
        }, new Test.TestRequest(), new Test.TestResponse());

        Tinyros.nh().advertiseService(srv);

        while(true) {
            Thread.sleep(10*1000);
        }
    }
}
