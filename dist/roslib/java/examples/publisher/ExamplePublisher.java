package examples.publisher;

import com.roslib.ros.Publisher;
import com.roslib.ros.Tinyros;
import com.roslib.tinyros_hello.TinyrosHello;

public class ExamplePublisher {

    public static void main(String[] args) throws InterruptedException {
        Tinyros.init("127.0.0.1");
        
        Publisher<TinyrosHello> pub = new Publisher<TinyrosHello>("tinyros_hello", new TinyrosHello());
        
        Tinyros.nh().advertise(pub);
        /*Tinyros.udp().advertise(pub);*/
        
        while(true) {
            TinyrosHello msg = new TinyrosHello();
            msg.hello = "Hello, tiny-ros ^_^";
            pub.publish(msg);
            Thread.sleep(1000);
        }
    }
}
