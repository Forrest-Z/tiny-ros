/*
 * File      : tinyros_entries.h
 * This file is part of tinyros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-01-22     Pinkie.Fu    initial version
 */

#ifndef _TINYROS_ENTRIES_H_
#define _TINYROS_ENTRIES_H_
#include <stdint.h>
#include <stdlib.h>
#ifdef __cplusplus
extern "C" {
#endif
void tinyros_example_publisher(void* parameter);
void tinyros_example_subscriber(void* parameter);
void tinyros_example_service(void* parameter);
void tinyros_example_service_client(void* parameter);

#ifdef __cplusplus
}
#endif
#endif

