#ifndef _TINYROS_STRING_H_
#define _TINYROS_STRING_H_

#include <string.h>

namespace tinyros
{
class string
{
public:
  string() {
    length = 0;
    data = new char[1];
    data[length] = '\0';
  }

  string(const char* str) {
    if (!str) {
      length = 0;
      data = new char[1];
    } else {
      length = strlen(str);
      data = new char[length + 1];
      strcpy(data, str);
    }
    data[length] = '\0';
  }

  string(const string &str) {
    length = str.size();
    data = new char[length + 1];
    strcpy(data, str.c_str());
    data[length] = '\0';
  }

  ~string() {
    delete[] data;
    length = 0;
  }

  string& operator=(const string &str) {
    if (this == &str) {
      return *this;
    }

    delete[] data;

    length = str.size();
    data = new char[length + 1];
    strcpy(data, str.c_str());
    data[length] = '\0';
    return *this;
  }  

  string& operator=(const char* str) {
    delete[] data;

    if (!str) {
      length = 0;
      data = new char[1];
    } else {
      length = strlen(str);
      data = new char[length + 1];
      strcpy(data, str);
    }
    data[length] = '\0';
    return *this;
  }

  string operator+(const string &str) const {
    string newString;
    newString.length = length + str.size();
    newString.data = new char[newString.length + 1];
    strcpy(newString.data, data);
    strcat(newString.data, str.data);
    data[length] = '\0';
    return newString;
  }

  string operator+(const char* str) const {
    string newString;
    if (!str) {
      newString.length = length;
      newString.data = new char[newString.length + 1];
      strcpy(newString.data, data);
    } else {
      newString.length = length + strlen(str);
      newString.data = new char[newString.length + 1];
      strcpy(newString.data, data);
      strcat(newString.data, str);
    }
    data[length] = '\0';
    return newString;
  }

  string& operator+=(const string &str) {
    length += str.size();
    char *newdata = new char[length + 1];
    strcpy(newdata, data);
    strcat(newdata, str.c_str());
    delete[] data;
    data = newdata;
    data[length] = '\0';
    return *this;
  }

  string& operator+=(const char* &str) {
    if (str) {
      length += strlen(str);
      char *newdata = new char[length + 1];
      strcpy(newdata, data);
      strcat(newdata, str);
      delete[] data;
      data = newdata;
    }
    data[length] = '\0';
    return *this;
  }

  char& operator[](int n) const {
    if (n >= length) {
      return data[length - 1];
    } else if (n < 0) {
      return data[0];
    } else {
      return data[n];
    }
  }

  bool operator==(const string &str) const {
    if (length != str.size()) {
      return false;
    }

    return strcmp(data, str.c_str()) ? false : true;
  }

  bool operator==(const char* &str) const {
    if (!str) {
      return false;
    }

    return strcmp(data, str) ? false : true;
  }

  int size() const {
    return length;
  }

  bool empty() const {
    return (length > 0 ? false : true);
  }

  const char *c_str() const {
    return data;
  }

private:
  char *data;
  int length;
};
}
#endif
