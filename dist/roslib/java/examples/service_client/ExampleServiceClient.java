package examples.service_client;

import com.roslib.ros.NodeHandle;
import com.roslib.ros.ServiceClient;
import com.roslib.tinyros_hello.Test;

public class ExampleServiceClient {

	public static void main(String[] args) throws InterruptedException {

		NodeHandle nh = new NodeHandle();

        while (!nh.initNode("127.0.0.1")) {
			System.out.println("initNode failed.");
			Thread.sleep(500);
		}

		ServiceClient<Test.TestRequest, Test.TestResponse> client = new ServiceClient<Test.TestRequest, Test.TestResponse>(
				"test_srv", new Test.TestRequest(), new Test.TestResponse());

		nh.serviceClient(client);

		new Thread(new Runnable(){  
			ServiceClient<Test.TestRequest, Test.TestResponse> cl = client;
			public void run(){  
				while(true) {
					Test.TestRequest req = new Test.TestRequest();
					Test.TestResponse resp = new Test.TestResponse();
					if (cl.call(req, resp, 5)) {
						System.out.println("service responsed with \"" + resp.output + "\"");
					} else {
						System.out.println("Service call failed.");
					}
				}
			}}).start(); 

		while(nh.ok()) {
			nh.spinOnce();
			Thread.sleep(500);
		}

	}
}
