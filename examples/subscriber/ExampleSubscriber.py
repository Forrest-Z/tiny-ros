import os
import sys
import time
print('Tinyros modules: %s' % os.path.abspath(os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))))
sys.path.append(os.path.abspath(os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))))
import tinyros
import tinyros_hello.msg.TinyrosHello

def subscriber_cb(received_msg):
    print('%s' % received_msg.hello)

def main():
    tinyros.init("PyExampleSubscriber", "127.0.0.1")
    if 1:
        tinyros.nh().subscribe(tinyros.Subscriber("tinyros_hello", subscriber_cb, tinyros_hello.msg.TinyrosHello))
    else:
        tinyros.udp().subscribe(tinyros.Subscriber("tinyros_hello", subscriber_cb, tinyros_hello.msg.TinyrosHello))
    while True:
        time.sleep(10)

if __name__ == '__main__':
    main()

