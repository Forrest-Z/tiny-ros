package examples.subscriber;

import com.roslib.ros.CallbackSubT;
import com.roslib.ros.Msg;
import com.roslib.ros.Subscriber;
import com.roslib.ros.Tinyros;
import com.roslib.tinyros_hello.TinyrosHello;

public class ExampleSubscriber {

    public static void main(String[] args) throws InterruptedException {
        Tinyros.init("JavaExampleSubscriber", "127.0.0.1");

        Subscriber<TinyrosHello> sub = new Subscriber<TinyrosHello>
        ("tinyros_hello", new CallbackSubT() {
            @Override
            public void callback(Msg msg) {
                TinyrosHello m = (TinyrosHello)msg;
                System.out.println(m.hello);
            }
        }, new TinyrosHello());

        Tinyros.nh().subscribe(sub);
        /*Tinyros.udp().subscribe(sub);*/

        while(true) {
            Thread.sleep(10*1000);
        }
    }
}
