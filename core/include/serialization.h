#ifndef TINY_ROS_SERIALIZATION_H
#define TINY_ROS_SERIALIZATION_H
#include "exception.h"
#include <vector>
#include <map>
#include <stdint.h>
#include <cstring>

namespace tinyros
{
namespace serialization
{
class StreamOverrunException : public tinyros::Exception
{
public:
  StreamOverrunException(const std::string& what)
  : Exception(what)
  {}
};

void throwStreamOverrun();

template<typename T>
struct Serializer
{
  template<typename Stream>
  inline static void write(Stream& stream, T t)
  {
    stream.advance(t.serialize(stream.getData()));
  }

  template<typename Stream>
  inline static void read(Stream& stream, T& t)
  {
    stream.advance(t.deserialize(stream.getData()));
  }

  inline static uint32_t serializedLength(T t)
  {
    return t.serializedLength();
  }
};

template<typename T, typename Stream>
inline void serialize(Stream& stream, const T& t)
{
  Serializer<T>::write(stream, t);
}

template<typename T, typename Stream>
inline void deserialize(Stream& stream, T& t)
{
  Serializer<T>::read(stream, t);
}

template<typename T>
inline uint32_t serializationLength(const T& t)
{
  return Serializer<T>::serializedLength(t);
}

#define ROS_CREATE_SIMPLE_SERIALIZER(Type) \
  template<> struct Serializer<Type> \
  { \
    template<typename Stream> inline static void write(Stream& stream, const Type v) \
    { \
      *reinterpret_cast<Type*>(stream.advance(sizeof(v))) = v; \
    } \
    \
    template<typename Stream> inline static void read(Stream& stream, Type& v) \
    { \
      v = *reinterpret_cast<Type*>(stream.advance(sizeof(v))); \
    } \
    \
    inline static uint32_t serializedLength(const Type&) \
    { \
      return sizeof(Type); \
    } \
  };

#define ROS_CREATE_SIMPLE_SERIALIZER_ARM(Type) \
  template<> struct Serializer<Type> \
  { \
    template<typename Stream> inline static void write(Stream& stream, const Type v) \
    { \
      memcpy(stream.advance(sizeof(v)), &v, sizeof(v) ); \
    } \
    \
    template<typename Stream> inline static void read(Stream& stream, Type& v) \
    { \
      memcpy(&v, stream.advance(sizeof(v)), sizeof(v) ); \
    } \
    \
    inline static uint32_t serializedLength(const Type t) \
    { \
      return sizeof(Type); \
    } \
};

#if defined(__arm__) || defined(__arm)
    ROS_CREATE_SIMPLE_SERIALIZER_ARM(uint8_t);
    ROS_CREATE_SIMPLE_SERIALIZER_ARM(int8_t);
    ROS_CREATE_SIMPLE_SERIALIZER_ARM(uint16_t);
    ROS_CREATE_SIMPLE_SERIALIZER_ARM(int16_t);
    ROS_CREATE_SIMPLE_SERIALIZER_ARM(uint32_t);
    ROS_CREATE_SIMPLE_SERIALIZER_ARM(int32_t);
    ROS_CREATE_SIMPLE_SERIALIZER_ARM(uint64_t);
    ROS_CREATE_SIMPLE_SERIALIZER_ARM(int64_t);
    ROS_CREATE_SIMPLE_SERIALIZER_ARM(float);
    ROS_CREATE_SIMPLE_SERIALIZER_ARM(double);
#else
    ROS_CREATE_SIMPLE_SERIALIZER(uint8_t);
    ROS_CREATE_SIMPLE_SERIALIZER(int8_t);
    ROS_CREATE_SIMPLE_SERIALIZER(uint16_t);
    ROS_CREATE_SIMPLE_SERIALIZER(int16_t);
    ROS_CREATE_SIMPLE_SERIALIZER(uint32_t);
    ROS_CREATE_SIMPLE_SERIALIZER(int32_t);
    ROS_CREATE_SIMPLE_SERIALIZER(uint64_t);
    ROS_CREATE_SIMPLE_SERIALIZER(int64_t);
    ROS_CREATE_SIMPLE_SERIALIZER(float);
    ROS_CREATE_SIMPLE_SERIALIZER(double);
#endif

template<> struct Serializer<bool>
{
  template<typename Stream> inline static void write(Stream& stream, const bool v)
  {
    uint8_t b = (uint8_t)v;
#if defined(__arm__) || defined(__arm)
    memcpy(stream.advance(1), &b, 1 );
#else
    *reinterpret_cast<uint8_t*>(stream.advance(1)) = b;
#endif
  }

  template<typename Stream> inline static void read(Stream& stream, bool& v)
  {
    uint8_t b;
#if defined(__arm__) || defined(__arm)
    memcpy(&b, stream.advance(1), 1 );
#else
    b = *reinterpret_cast<uint8_t*>(stream.advance(1));
#endif
    v = (bool)b;
  }

  inline static uint32_t serializedLength(bool)
  {
    return 1;
  }
};

template<class ContainerAllocator>
struct Serializer<std::basic_string<char, std::char_traits<char>, ContainerAllocator> >
{
  typedef std::basic_string<char, std::char_traits<char>, ContainerAllocator> StringType;

  template<typename Stream>
  inline static void write(Stream& stream, const StringType& str)
  {
    size_t len = str.size();
    stream.next((uint32_t)len);

    if (len > 0)
    {
      memcpy(stream.advance((uint32_t)len), str.data(), len);
    }
  }

  template<typename Stream>
  inline static void read(Stream& stream, StringType& str)
  {
    uint32_t len;
    stream.next(len);
    if (len > 0)
    {
      str = StringType((char*)stream.advance(len), len);
    }
    else
    {
      str.clear();
    }
  }

  inline static uint32_t serializedLength(const StringType& str)
  {
    return 4 + (uint32_t)str.size();
  }
};

namespace stream_types
{
enum StreamType
{
  Input,
  Output,
  Length
};
}
typedef stream_types::StreamType StreamType;

struct Stream
{
  inline uint8_t* getData() { return data_; }
  uint8_t* advance(uint32_t len)
  {
    uint8_t* old_data = data_;
    data_ += len;
    if (data_ > end_)
    {
      throwStreamOverrun();
    }
    return old_data;
  }

  inline uint32_t getLength() { return (uint32_t)(end_ - data_); }

protected:
  Stream(uint8_t* _data, uint32_t _count)
  : data_(_data)
  , end_(_data + _count)
  {}

private:
  uint8_t* data_;
  uint8_t* end_;
};

struct IStream : public Stream
{
  static const StreamType stream_type = stream_types::Input;

  IStream(uint8_t* data, uint32_t count)
  : Stream(data, count)
  {}

  template<typename T>
  void next(T& t)
  {
    deserialize(*this, t);
  }

  template<typename T>
  IStream& operator>>(T& t)
  {
    deserialize(*this, t);
    return *this;
  }
};

struct OStream : public Stream
{
  static const StreamType stream_type = stream_types::Output;

  OStream(uint8_t* data, uint32_t count)
  : Stream(data, count)
  {}

  template<typename T>
  void next(const T& t)
  {
    serialize(*this, t);
  }

  template<typename T>
  OStream& operator<<(const T& t)
  {
    serialize(*this, t);
    return *this;
  }
};


struct LStream
{
  static const StreamType stream_type = stream_types::Length;

  LStream()
  : count_(0)
  {}

  template<typename T>
  void next(const T& t)
  {
    count_ += serializationLength(t);
  }

  uint32_t advance(uint32_t len)
  {
    uint32_t old = count_;
    count_ += len;
    return old;
  }

  inline uint32_t getLength() { return count_; }

private:
  uint32_t count_;
};
}
}

#endif // TINY_ROS_SERIALIZATION_H
