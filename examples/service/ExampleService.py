import os
import sys
import time
print('Tinyros modules: %s' % os.path.abspath(os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))))
sys.path.append(os.path.abspath(os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))))
import tinyros
import tinyros_hello.srv.Test

def service_cb(req, res):
    res.output = "Hello, tiny-ros ^_^"

def main():
    tinyros.init("PyExampleService", "127.0.0.1")
    tinyros.nh().advertiseService(tinyros.ServiceServer("test_srv", service_cb, \
            tinyros_hello.srv.Test.Request, tinyros_hello.srv.Test.Response))
    while True:
        time.sleep(10)

if __name__ == '__main__':
    main()

