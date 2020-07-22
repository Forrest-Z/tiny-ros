package examples.service_client;

import com.roslib.ros.ServiceClient;
import com.roslib.ros.Tinyros;
import com.roslib.tinyros_hello.Test;

public class ExampleServiceClient {

    public static void main(String[] args) throws InterruptedException {
        Tinyros.init("JavaExampleServiceClient", "127.0.0.1");

        ServiceClient<Test.TestRequest, Test.TestResponse> client =
                new ServiceClient<Test.TestRequest, Test.TestResponse>(
                        "test_srv", new Test.TestRequest(), new Test.TestResponse());

        Tinyros.nh().serviceClient(client);

        while(true) {
            Test.TestRequest req = new Test.TestRequest();
            Test.TestResponse resp = new Test.TestResponse();
            if (client.call(req, resp, 3)) {
                System.out.println("service responsed with \"" + resp.output + "\"");
            } else {
                System.out.println("Service call failed.");
            }
            Thread.sleep(1000);
        }
    }
}
