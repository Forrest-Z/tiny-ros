/*
 * File      : hardware.h
 * This file is part of tiny_ros
 *
 * Change Logs:
 * Date           Author       Notes
 * 2018-04-24     Pinkie.Fu    initial version
 */
#ifndef TINYROS_HARDWARE_H_
#define TINYROS_HARDWARE_H_

namespace tinyros {

class Hardware {
public:
  virtual bool init(std::string portName) = 0;

  virtual int read(uint8_t* data, int length) = 0;

  virtual bool write(uint8_t* data, int length) = 0;

  virtual bool connected() = 0;

  virtual void close() = 0;
};
}
#endif //TINYROS_HARDWARE_H_
