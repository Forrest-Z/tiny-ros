/*
 * File      : main.c
 * This file is part of tinyros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-01-22     Pinkie.Fu    initial version
 */
#include <rtthread.h>
#include <lwip/sys.h>
#include <lwip/api.h>
#include <netif/ethernetif.h>
#include "tinyros_entries.h"

extern void lwip_sys_init(void);

int main(void) {
  rt_thread_t publisher_thread_ = RT_NULL;
  rt_thread_t subscriber_thread_ = RT_NULL;
  rt_thread_t service_thread_ = RT_NULL;
  rt_thread_t service_client_thread_ = RT_NULL;

  //{ init lwip
  eth_system_device_init();
  rt_device_init_all();
  lwip_sys_init();
  // }

  // {tinyros_example_publisher
  publisher_thread_ = rt_thread_create("pub", tinyros_example_publisher, RT_NULL, 1024, 5, 20);
  rt_thread_startup(publisher_thread_);
  // }

  // {tinyros_example_subscriber
  subscriber_thread_ = rt_thread_create("sub", tinyros_example_subscriber, RT_NULL, 1024, 5, 20);
  rt_thread_startup(subscriber_thread_);
  // }

  // {tinyros_example_service
  service_thread_ = rt_thread_create("svc", tinyros_example_service, RT_NULL, 1024, 5, 20);
  rt_thread_startup(service_thread_);
  // }

  // {tinyros_example_service_client
  service_client_thread_ = rt_thread_create("svcclient", tinyros_example_service_client, RT_NULL, 1024, 5, 20);
  rt_thread_startup(service_client_thread_);
  // }
  return 0;
}

