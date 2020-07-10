#!/usr/bin/python

from __future__ import print_function

__usage__ = """
python make_library.py <output_path>
"""
import traceback

import os, sys, re, hashlib

# for copying files
import shutil

def type_to_var(ty):
    lookup = {
        1 : 'uint8_t',
        2 : 'uint16_t',
        4 : 'uint32_t',
        8 : 'uint64_t',
    }
    return lookup[ty]

def hashlib_md5sum(name, package, definition):
    str = package + "/" + name
    # parse definition
    for line in definition:
        # prep work
        line = line.strip().rstrip()
        value = None
        if line.find("#") > -1:
            line = line[0:line.find("#")]
        if line.find("=") > -1:
            try:
                value = line[line.find("=")+1:]
            except:
                value = '"' + line[line.find("=")+1:] + '"'
            line = line[0:line.find("=")]

        # find package/class name
        line = line.replace("\t", " ")
        l = line.split(" ")
        while "" in l:
            l.remove("")
        if len(l) < 2:
            continue
        ty, name = l[0:2]
        str = str + "|" + ty + " " + name
        if value != None:
            value = value.strip().rstrip()
            str = str + "=" + value
    hl = hashlib.md5()
    hl.update(str)
    return hl.hexdigest()

def hashlib_md5sum_definition(definition):
    str = "";
    # parse definition
    for line in definition:
        str += line;
    hl = hashlib.md5()
    hl.update(str)
    return hl.hexdigest()

#####################################################################
# Data Types

class EnumerationType:
    """ For data values. """

    def __init__(self, name, ty, value):
        self.name = name
        self.type = ty
        self.value = value

    def make_declaration(self, f):
        f.write('      enum { %s = %s };\n' % (self.name, self.value))

class PrimitiveDataType:
    """ Our datatype is a C/C++ primitive. """

    def __init__(self, name, ty, bytes):
        self.name = name
        self.type = ty
        self.bytes = bytes

    def make_initializer(self, f, trailer):
        f.write('      %s(0)%s\n' % (self.name, trailer))

    def make_declaration(self, f):
        f.write('      typedef %s _%s_type;\n      _%s_type %s;\n' % (self.type, self.name, self.name, self.name) )

    def serialize(self, f, header):
        cn = self.name.replace("[","").replace("]","").split(".")[-1]
        if self.type != type_to_var(self.bytes):
            f.write('%s      union {\n' % header)
            f.write('%s        %s real;\n' % (header, self.type))
            f.write('%s        %s base;\n' % (header, type_to_var(self.bytes)))
            f.write('%s      } u_%s;\n' % (header, cn))
            f.write('%s      u_%s.real = this->%s;\n' % (header, cn, self.name))
            for i in range(self.bytes):
                f.write('%s      *(outbuffer + offset + %d) = (u_%s.base >> (8 * %d)) & 0xFF;\n' % (header, i, cn, i) )
        else:
            for i in range(self.bytes):
                f.write('%s      *(outbuffer + offset + %d) = (this->%s >> (8 * %d)) & 0xFF;\n' % (header, i, self.name, i) )
        f.write('%s      offset += sizeof(this->%s);\n' % (header, self.name))

    def deserialize(self, f, header):
        cn = self.name.replace("[","").replace("]","").split(".")[-1]
        if self.type != type_to_var(self.bytes):
            f.write('%s      union {\n' % header)
            f.write('%s        %s real;\n' % (header, self.type))
            f.write('%s        %s base;\n' % (header, type_to_var(self.bytes)))
            f.write('%s      } u_%s;\n' % (header, cn))
            f.write('%s      u_%s.base = 0;\n' % (header, cn))
            for i in range(self.bytes):
                f.write('%s      u_%s.base |= ((%s) (*(inbuffer + offset + %d))) << (8 * %d);\n' % (header, cn,type_to_var(self.bytes),i,i) )
            f.write('%s      this->%s = u_%s.real;\n' % (header, self.name, cn) )
        else:
            f.write('%s      this->%s =  ((%s) (*(inbuffer + offset)));\n' % (header, self.name,self.type) )
            for i in range(self.bytes-1):
                f.write('%s      this->%s |= ((%s) (*(inbuffer + offset + %d))) << (8 * %d);\n' % (header, self.name,self.type,i+1,i+1) )
        f.write('%s      offset += sizeof(this->%s);\n' % (header, self.name))

    def serializedLength(self, f, header):
        f.write('%s      length += sizeof(this->%s);\n' % (header, self.name))

    def echo(self, f, header, trailer):
        if self.name.find('[') > 0:
            cn = self.name.replace("[","").replace("]","").split(".")[-1]
            if self.type == 'int8_t':
                f.write('%s      std::stringstream ss_%s; ss_%s <<(int16_t)%s <<"%s";\n' % (header, cn, cn, self.name, trailer))
            elif self.type == 'uint8_t':
                f.write('%s      std::stringstream ss_%s; ss_%s << (uint16_t)%s <<"%s";\n' % (header, cn, cn, self.name, trailer))
            elif self.type == 'char':
                f.write('%s      std::stringstream ss_%s; ss_%s <<"\\"" << %s <<"\\"%s";\n' % (header, cn, cn, self.name, trailer))
            else:
                f.write('%s      std::stringstream ss_%s; ss_%s << %s <<"%s";\n' % (header, cn, cn, self.name, trailer))
            f.write('%s      string_echo += ss_%s.str();\n' % (header, cn))
        else:
            if self.type == 'int8_t':
                f.write('%s      std::stringstream ss_%s; ss_%s << "\\"%s\\":" << (int16_t)%s <<"%s";\n' % (header, self.name, self.name, self.name, self.name, trailer))
            elif self.type == 'uint8_t':
                f.write('%s      std::stringstream ss_%s; ss_%s << "\\"%s\\":" << (uint16_t)%s <<"%s";\n' % (header, self.name, self.name, self.name, self.name, trailer))
            elif self.type == 'char':
                f.write('%s      std::stringstream ss_%s; ss_%s << "\\"%s\\":\\"" << %s <<"\\"%s";\n' % (header, self.name, self.name, self.name, self.name, trailer))
            else:
                f.write('%s      std::stringstream ss_%s; ss_%s << "\\"%s\\":" << %s <<"%s";\n' % (header, self.name, self.name, self.name, self.name, trailer))
            f.write('%s      string_echo += ss_%s.str();\n' % (header, self.name)) 

class MessageDataType(PrimitiveDataType):
    """ For when our data type is another message. """

    def make_initializer(self, f, trailer):
        f.write('      %s()%s\n' % (self.name, trailer))

    def serialize(self, f, header):
        f.write('%s      offset += this->%s.serialize(outbuffer + offset);\n' % (header, self.name))

    def deserialize(self, f, header):
        f.write('%s      offset += this->%s.deserialize(inbuffer + offset);\n' % (header, self.name))

    def serializedLength(self, f, header):
        f.write('%s      length += this->%s.serializedLength();\n' % (header, self.name))

    def echo(self, f, header, trailer):
        if self.name.find('[') > 0:
            f.write('%s      string_echo += this->%s.echo();\n' % (header, self.name))
            f.write('%s      string_echo += "%s";\n' % (header, trailer))
        else:
            f.write('%s      string_echo += "\\"%s\\":";\n' % (header, self.name))
            f.write('%s      string_echo += this->%s.echo();\n' % (header, self.name))
            f.write('%s      string_echo += "%s";\n' % (header, trailer))

class StringDataType(PrimitiveDataType):
    def make_initializer(self, f, trailer):
        f.write('      %s("")%s\n' % (self.name, trailer))

    def make_declaration(self, f):
        f.write('      typedef std::string _%s_type;\n      _%s_type %s;\n' % (self.name, self.name, self.name) )

    def serialize(self, f, header):
        cn = self.name.replace("[","").replace("]","")
        f.write('%s      uint32_t length_%s = this->%s.size();\n' % (header, cn,self.name))
        f.write('%s      varToArr(outbuffer + offset, length_%s);\n' % (header, cn))
        f.write('%s      offset += 4;\n' % header)
        f.write('%s      memcpy(outbuffer + offset, this->%s.c_str(), length_%s);\n' % (header, self.name,cn))
        f.write('%s      offset += length_%s;\n' % (header, cn))

    def deserialize(self, f, header):
        cn = self.name.replace("[","").replace("]","")
        f.write('%s      uint32_t length_%s;\n' % (header, cn))
        f.write('%s      arrToVar(length_%s, (inbuffer + offset));\n' % (header, cn))
        f.write('%s      offset += 4;\n' % header)
        f.write('%s      for(unsigned int k= offset; k< offset+length_%s; ++k){\n' % (header, cn))
        f.write('%s        inbuffer[k-1]=inbuffer[k];\n' % header)
        f.write('%s      }\n' % header)
        f.write('%s      inbuffer[offset+length_%s-1]=0;\n' % (header, cn))
        f.write('%s      this->%s = (char *)(inbuffer + offset-1);\n' % (header, self.name))
        f.write('%s      offset += length_%s;\n' % (header, cn))

    def serializedLength(self, f, header):
        cn = self.name.replace("[","").replace("]","")
        f.write('%s      uint32_t length_%s = this->%s.size();\n' % (header, cn, self.name))
        f.write('%s      length += 4;\n' % header)
        f.write('%s      length += length_%s;\n' % (header, cn))

    def echo(self, f, header, trailer):
        if self.name.find('[') > 0:
            cn = self.name.replace("[","").replace("]","").split(".")[-1]
            f.write('%s      std::stringstream ss_%s; ss_%s << "\\"";\n' % (header, cn, cn))
            f.write('%s      string_echo += ss_%s.str();\n' % (header, cn))
            f.write('%s      std::size_t %s_pos = %s.find("\\"");\n' % (header, cn, self.name))
            f.write('%s      while(%s_pos != std::string::npos){\n' % (header, cn))
            f.write('%s        %s.replace(%s_pos, 1,"\\\\\\"");\n' % (header, self.name, cn))
            f.write('%s        %s_pos = %s.find("\\"", %s_pos+2);\n' % (header, cn, self.name, cn))
            f.write('%s      }\n' % header)
            f.write('%s      string_echo += %s;\n' % (header, self.name))
            f.write('%s      string_echo += "\\"%s";\n' % (header, trailer))
        else:
            f.write('%s      std::size_t %s_pos = %s.find("\\"");\n' % (header, self.name, self.name))
            f.write('%s      while(%s_pos != std::string::npos){\n' % (header, self.name))
            f.write('%s        %s.replace(%s_pos, 1,"\\\\\\"");\n' % (header, self.name, self.name))
            f.write('%s        %s_pos = %s.find("\\"", %s_pos+2);\n' % (header, self.name, self.name, self.name))
            f.write('%s      }\n' % header)
            f.write('%s      string_echo += "\\"%s\\":\\"";\n' % (header, self.name))
            f.write('%s      string_echo += %s;\n' % (header, self.name))
            f.write('%s      string_echo += "\\"%s";\n' % (header, trailer))

class TimeDataType(PrimitiveDataType):

    def __init__(self, name, ty, bytes):
        self.name = name
        self.type = ty
        self.sec = PrimitiveDataType(name+'.sec','uint32_t',4)
        self.nsec = PrimitiveDataType(name+'.nsec','uint32_t',4)

    def make_initializer(self, f, trailer):
        f.write('      %s()%s\n' % (self.name, trailer))

    def make_declaration(self, f):
        f.write('      typedef %s _%s_type;\n      _%s_type %s;\n' % (self.type, self.name, self.name, self.name) )

    def serialize(self, f, header):
        self.sec.serialize(f, header)
        self.nsec.serialize(f, header)

    def deserialize(self, f, header):
        self.sec.deserialize(f, header)
        self.nsec.deserialize(f, header)

    def serializedLength(self, f, header):
        self.sec.serializedLength(f, header)
        self.nsec.serializedLength(f, header)

    def echo(self, f, header, trailer):
        if self.name.find('[') > 0:
            cn = self.name.replace("[","").replace("]","").split(".")[-1]
            f.write('%s      std::stringstream ss_%s; ss_%s << "{";\n' % (header, cn, cn))
            f.write('%s      ss_%s << "\\"sec\\":" << %s.sec;\n' % (header, cn, self.name))
            f.write('%s      ss_%s << ",\\"nsec\\":" << %s.nsec << "}%s";\n' % (header, cn, self.name, trailer))
            f.write('%s      string_echo += ss_%s.str();\n' % (header, cn))
        else:
            f.write('%s      std::stringstream ss_%s;\n' % (header, self.name))
            f.write('%s      ss_%s << "\\"%s\\":{\\"sec\\":" << %s.sec;\n' % (header, self.name, self.name, self.name))
            f.write('%s      ss_%s << ",\\"nsec\\":" << %s.nsec << "}%s";\n' % (header, self.name, self.name, trailer))
            f.write('%s      string_echo += ss_%s.str();\n' % (header, self.name))

class ArrayDataType(PrimitiveDataType):

    def __init__(self, name, ty, bytes, cls, array_size=None):
        self.name = name
        self.type = ty
        self.bytes = bytes
        self.size = array_size
        self.cls = cls

    def make_initializer(self, f, trailer):
        if self.size == None:
            f.write('      %s_length(0), %s(NULL)%s\n' % (self.name, self.name, trailer))
        else:
            f.write('      %s()%s\n' % (self.name, trailer))

    def make_declaration(self, f):
        if self.size == None:
            f.write('      uint32_t %s_length;\n' % self.name)
            f.write('      typedef %s _%s_type;\n' % (self.type, self.name))
            f.write('      _%s_type st_%s;\n' % (self.name, self.name)) # static instance for copy
            f.write('      _%s_type * %s;\n' % (self.name, self.name))
        else:
            f.write('      %s %s[%d];\n' % (self.type, self.name, self.size))

    def serialize(self, f, header):
        c = self.cls(self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            # serialize length
            f.write('%s      *(outbuffer + offset + 0) = (this->%s_length >> (8 * 0)) & 0xFF;\n' % (header, self.name))
            f.write('%s      *(outbuffer + offset + 1) = (this->%s_length >> (8 * 1)) & 0xFF;\n' % (header, self.name))
            f.write('%s      *(outbuffer + offset + 2) = (this->%s_length >> (8 * 2)) & 0xFF;\n' % (header, self.name))
            f.write('%s      *(outbuffer + offset + 3) = (this->%s_length >> (8 * 3)) & 0xFF;\n' % (header, self.name))
            f.write('%s      offset += sizeof(this->%s_length);\n' % (header, self.name))
            f.write('%s      for( uint32_t i = 0; i < %s_length; i++) {\n' % (header, self.name))
            c.serialize(f, header + "  ")
            f.write('%s      }\n' % header)
        else:
            f.write('%s      for( uint32_t i = 0; i < %d; i++) {\n' % (header, self.size) )
            c.serialize(f, header + "  ")
            f.write('%s      }\n' % header)

    def deserialize(self, f, header):
        if self.size == None:
            c = self.cls("st_"+self.name, self.type, self.bytes)
            # deserialize length
            f.write('%s      uint32_t %s_lengthT = ((uint32_t) (*(inbuffer + offset))); \n' % (header, self.name))
            f.write('%s      %s_lengthT |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1); \n' % (header, self.name))
            f.write('%s      %s_lengthT |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2); \n' % (header, self.name))
            f.write('%s      %s_lengthT |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3); \n' % (header, self.name))
            f.write('%s      offset += sizeof(this->%s_length);\n' % (header, self.name))
            f.write('%s      if(%s_lengthT > %s_length)\n' % (header, self.name, self.name))
            f.write('%s        this->%s = (%s*)realloc(this->%s, %s_lengthT * sizeof(%s));\n' % (header, self.name, self.type, self.name, self.name, self.type))
            f.write('%s      %s_length = %s_lengthT;\n' % (header, self.name, self.name))
            f.write('%s      for( uint32_t i = 0; i < %s_length; i++) {\n' % (header, self.name))
            c.deserialize(f, header + "  ")
            f.write('%s        memcpy( &(this->%s[i]), &(this->st_%s), sizeof(%s));\n' % (header, self.name, self.name, self.type))
            f.write('%s      }\n' % header)
        else:
            c = self.cls(self.name+"[i]", self.type, self.bytes)
            f.write('%s      for( uint32_t i = 0; i < %d; i++){\n' % (header, self.size) )
            c.deserialize(f, header + "  ")
            f.write('%s      }\n' % header)

    def serializedLength(self, f, header):
        c = self.cls(self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            # serialize length
            f.write('%s      length += sizeof(this->%s_length);\n' % (header, self.name))
            f.write('%s      for( uint32_t i = 0; i < %s_length; i++) {\n' % (header, self.name))
            c.serializedLength(f, header + "  ")
            f.write('%s      }\n' % header)
        else:
            f.write('%s      for( uint32_t i = 0; i < %d; i++) {\n' % (header, self.size) )
            c.serializedLength(f, header + "  ")
            f.write('%s      }\n' % header)

    def echo(self, f, header, trailer):
        c = self.cls(self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            f.write('%s      string_echo += "%s:[";\n' % (header, self.name))
            f.write('%s      for( uint32_t i = 0; i < %s_length; i++) {\n' % (header, self.name))
            f.write('%s        if( i == (%s_length - 1)) {\n' % (header, self.name))
            c.echo(f, header + "    ", "")
            f.write('%s        } else {\n' % header)
            c.echo(f, header + "    ", ",")
            f.write('%s        }\n' % header)
            f.write('%s      }\n' % header)
            f.write('%s      string_echo += "]%s";\n' % (header, trailer))
        else:
            f.write('%s      string_echo += "%s:[";\n' % (header, self.name))
            f.write('%s      for( uint32_t i = 0; i < %d; i++) {\n' % (header, self.size) )
            f.write('%s        if( i == (%d - 1)) {\n' % (header, self.size))
            c.echo(f, header + "    ", "")
            f.write('%s        } else {\n' % header)
            c.echo(f, header + "    ", ",")
            f.write('%s        }\n' % header)
            f.write('%s      }\n' % header)
            f.write('%s      string_echo += "]%s";\n' % (header, trailer))

ROS_TO_EMBEDDED_TYPES = {
    'bool'    :   ('bool',              1, PrimitiveDataType, []),
    'byte'    :   ('int8_t',            1, PrimitiveDataType, []),
    'int8'    :   ('int8_t',            1, PrimitiveDataType, []),
    'char'    :   ('char',              1, PrimitiveDataType, []),
    'uint8'   :   ('uint8_t',           1, PrimitiveDataType, []),
    'int16'   :   ('int16_t',           2, PrimitiveDataType, []),
    'uint16'  :   ('uint16_t',          2, PrimitiveDataType, []),
    'int32'   :   ('int32_t',           4, PrimitiveDataType, []),
    'uint32'  :   ('uint32_t',          4, PrimitiveDataType, []),
    'int64'   :   ('int64_t',           8, PrimitiveDataType, []),
    'uint64'  :   ('uint64_t',          8, PrimitiveDataType, []),
    'float32' :   ('float',             4, PrimitiveDataType, []),
    'float64' :   ('double',            8, PrimitiveDataType, []),
    'time'    :   ('tinyros::Time',         8, TimeDataType, ['ros/time']),
    'duration':   ('tinyros::Duration',     8, TimeDataType, ['ros/duration']),
    'string'  :   ('std::string',             0, StringDataType, []),
    'Header'  :   ('std_msgs::Header',  0, MessageDataType, ['std_msgs/Header'])
}

#####################################################################
# Messages

class Message:
    """ Parses message definitions into something we can export. """
    def __init__(self, name, package, definition):

        self.name = name            # name of message/class
        self.package = package      # package we reside in
        self.includes = list()      # other files we must include

        self.data = list()          # data types for code generation
        self.enums = list()
        self.md5 = hashlib_md5sum(name, package, definition)

        # parse definition
        for line in definition:
            # prep work
            line = line.strip().rstrip()
            value = None
            if line.find("#") > -1:
                line = line[0:line.find("#")]
            if line.find("=") > -1:
                try:
                    value = line[line.find("=")+1:]
                except:
                    value = '"' + line[line.find("=")+1:] + '"';
                line = line[0:line.find("=")]

            # find package/class name
            line = line.replace("\t", " ")
            l = line.split(" ")
            while "" in l:
                l.remove("")
            if len(l) < 2:
                continue
            ty, name = l[0:2]
            if value != None:
                self.enums.append( EnumerationType(name, ty, value))
                continue

            try:
                type_package, type_name = ty.split("/")
            except:
                type_package = None
                type_name = ty
            type_array = False
            if type_name.find('[') > 0:
                type_array = True
                try:
                    type_array_size = int(type_name[type_name.find('[')+1:type_name.find(']')])
                except:
                    type_array_size = None
                type_name = type_name[0:type_name.find('[')]

            # convert to C type if primitive, expand name otherwise
            try:
                code_type = ROS_TO_EMBEDDED_TYPES[type_name][0]
                size = ROS_TO_EMBEDDED_TYPES[type_name][1]
                cls = ROS_TO_EMBEDDED_TYPES[type_name][2]
                for include in ROS_TO_EMBEDDED_TYPES[type_name][3]:
                    if include not in self.includes:
                        self.includes.append(include)
            except:
                if type_package == None:
                    type_package = self.package
                if type_package+"/"+type_name not in self.includes:
                    self.includes.append(type_package+"/"+type_name)
                cls = MessageDataType
                code_type = type_package + "::" + type_name
                size = 0
            if type_array:
                self.data.append( ArrayDataType(name, code_type, size, cls, type_array_size ) )
            else:
                self.data.append( cls(name, code_type, size) )

    def _write_id(self, f):
        f.write('')

    def _write_id_initializer(self, f):
        f.write('')

    def _write_id_serializer(self, f):
        f.write('')

    def _write_id_deserializer(self, f):
        f.write('')

    def _write_getID(self, f):
        f.write('')

    def _write_setID(self, f):
        f.write('')

    def _write_serializer(self, f):
                # serializer
        f.write('    virtual int serialize(unsigned char *outbuffer) const\n')
        f.write('    {\n')
        f.write('      int offset = 0;\n')
        self._write_id_serializer(f)
        for d in self.data:
            d.serialize(f, "")
        f.write('      return offset;\n');
        f.write('    }\n')
        f.write('\n')

    def _write_deserializer(self, f):
        # deserializer
        f.write('    virtual int deserialize(unsigned char *inbuffer)\n')
        f.write('    {\n')
        f.write('      int offset = 0;\n')
        self._write_id_deserializer(f)
        for d in self.data:
            d.deserialize(f, "")
        f.write('      return offset;\n');
        f.write('    }\n')
        f.write('\n')

    def _write_serializedLength(self, f):
        # serializedLength
        f.write('    virtual int serializedLength() const\n')
        f.write('    {\n')
        f.write('      int length = 0;\n')
        for d in self.data:
            d.serializedLength(f, "")
        f.write('      return length;\n');
        f.write('    }\n')
        f.write('\n')

    def _write_echo(self, f):
        f.write('    virtual std::string echo()\n')
        f.write('    {\n')
        f.write('      std::string string_echo = "{";\n');
        if self.data:
            for d in self.data[:-1]:
                d.echo(f, "", ",")
            self.data[-1].echo(f, "", "")
        f.write('      string_echo += "}";\n');
        f.write('      return string_echo;\n');
        f.write('    }\n')
        f.write('\n')

    def _write_std_includes(self, f):
        f.write('#include <stdint.h>\n')
        f.write('#include <string>\n')
        f.write('#include <stdio.h>\n')
        f.write('#include <string.h>\n')
        f.write('#include <stdlib.h>\n')
        f.write('#include "tiny_ros/ros/msg.h"\n')

    def _write_msg_includes(self,f):
        for i in self.includes:
            f.write('#include "tiny_ros/%s.h"\n' % i)

    def _write_constructor(self, f):
        f.write('    %s()%s\n' % (self.name, ':' if self.data else ''))
        if self.data:
            for d in self.data[:-1]:
                d.make_initializer(f, ',')
            self.data[-1].make_initializer(f, '')
        f.write('    {\n')
        self._write_id_initializer(f)
        f.write('    }\n\n')

    def _write_data(self, f):
        for d in self.data:
            d.make_declaration(f)
        for e in self.enums:
            e.make_declaration(f)
        f.write('\n')

    def _write_getType(self, f):
        f.write('    virtual std::string getType(){ return "%s/%s"; }\n'%(self.package, self.name))

    def _write_getMD5(self, f):
        f.write('    virtual std::string getMD5(){ return "%s"; }\n'%self.md5)

    def _write_impl(self, f):
        f.write('  class %s : public tinyros::Msg\n' % self.name)
        f.write('  {\n')
        self._write_id(f)
        f.write('    public:\n')
        self._write_data(f)
        self._write_constructor(f)
        self._write_serializer(f)
        self._write_deserializer(f)
        self._write_serializedLength(f)
        self._write_echo(f)
        self._write_getType(f)
        self._write_getMD5(f)
        self._write_getID(f)
        self._write_setID(f)
        f.write('\n')
        f.write('  };\n')

    def make_header(self, f):
        f.write('#ifndef _TINYROS_%s_%s_h\n'%(self.package, self.name))
        f.write('#define _TINYROS_%s_%s_h\n'%(self.package, self.name))
        f.write('\n')
        self._write_std_includes(f)
        self._write_msg_includes(f)

        f.write('\n')
        f.write('namespace %s\n' % self.package)
        f.write('{\n')
        f.write('\n')
        self._write_impl(f)
        f.write('\n')
        f.write('}\n')

        f.write('#endif\n')

class Service:
    def __init__(self, name, package, definition):
        """
        @param name -  name of service
        @param package - name of service package
        @param definition - list of lines of  definition
        """

        self.name = name
        self.package = package

        sep_line = len(definition)
        sep = re.compile('---*')
        for i in range(0, len(definition)):
            if (None!= re.match(sep, definition[i]) ):
                sep_line = i
                break
        self.req_def = definition[0:sep_line]
        self.resp_def = definition[sep_line+1:]

        self.req = Message(name+"Request", package, self.req_def)
        self.resp = Message(name+"Response", package, self.resp_def)

    def make_header(self, f):
        f.write('#ifndef _TINYROS_SERVICE_%s_h\n' % self.name)
        f.write('#define _TINYROS_SERVICE_%s_h\n' % self.name)

        self.req._write_std_includes(f)
        includes = self.req.includes
        includes.extend(self.resp.includes)
        includes = list(set(includes))
        for inc in includes:
            f.write('#include "tiny_ros/%s.h"\n' % inc)

        f.write('\n')
        f.write('namespace %s\n' % self.package)
        f.write('{\n')
        f.write('\n')
        f.write('static const char %s[] = "%s/%s";\n'%(self.name.upper(), self.package, self.name))

        def write_id(out):
            out.write('    private:\n')
            out.write('      typedef uint32_t ___id___type;\n')
            out.write('      ___id___type __id__;\n\n')
        _write_id = lambda out: write_id(out)
        self.req._write_id = _write_id
        self.resp._write_id = _write_id

        def write_id_initializer(out):
            out.write('      this->__id__ = 0;\n')
        _write_id_initializer = lambda out: write_id_initializer(out)
        self.req._write_id_initializer = _write_id_initializer
        self.resp._write_id_initializer = _write_id_initializer

        def write_id_serializer(out):
            out.write('      *(outbuffer + offset + 0) = (this->__id__ >> (8 * 0)) & 0xFF;\n')
            out.write('      *(outbuffer + offset + 1) = (this->__id__ >> (8 * 1)) & 0xFF;\n')
            out.write('      *(outbuffer + offset + 2) = (this->__id__ >> (8 * 2)) & 0xFF;\n')
            out.write('      *(outbuffer + offset + 3) = (this->__id__ >> (8 * 3)) & 0xFF;\n')
            out.write('      offset += sizeof(this->__id__);\n')
        _write_id_serializer = lambda out: write_id_serializer(out)
        self.req._write_id_serializer = _write_id_serializer
        self.resp._write_id_serializer = _write_id_serializer

        def write_id_deserializer(out):
            out.write('      this->__id__ =  ((uint32_t) (*(inbuffer + offset)));\n')
            out.write('      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 1))) << (8 * 1);\n')
            out.write('      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 2))) << (8 * 2);\n')
            out.write('      this->__id__ |= ((uint32_t) (*(inbuffer + offset + 3))) << (8 * 3);\n')
            out.write('      offset += sizeof(this->__id__);\n')
        _write_id_deserializer = lambda out: write_id_deserializer(out)
        self.req._write_id_deserializer = _write_id_deserializer
        self.resp._write_id_deserializer = _write_id_deserializer

        def write_getID(out):
            out.write('    uint32_t getID() const { return this->__id__; }\n')
        _write_getID = lambda out: write_getID(out)
        self.req._write_getID = _write_getID
        self.resp._write_getID = _write_getID

        def write_setID(out):
            out.write('    void setID(uint32_t id){ this->__id__ = id; }\n')
        _write_setID = lambda out: write_setID(out)
        self.req._write_setID = _write_setID
        self.resp._write_setID = _write_setID

        def write_type(out, name):
            out.write('    virtual std::string getType(){ return %s; }\n'%(name))
        _write_getType = lambda out: write_type(out, self.name.upper())
        self.req._write_getType = _write_getType
        self.resp._write_getType = _write_getType

        f.write('\n')
        self.req._write_impl(f)
        f.write('\n')
        self.resp._write_impl(f)
        f.write('\n')
        f.write('  class %s {\n' % self.name )
        f.write('    public:\n')
        f.write('    typedef %s Request;\n' % self.req.name )
        f.write('    typedef %s Response;\n' % self.resp.name )
        f.write('  };\n')
        f.write('\n')

        f.write('}\n')

        f.write('#endif\n')

#####################################################################
# Make a Library

def MakeLibrary(pkg_dir, package, output_path, build_dir):
    # find the messages in this package
    messages = list()
    if os.path.exists(pkg_dir+"/msg"):
        print('Exporting %s\n'%package)
        sys.stdout.write('  Messages:')
        sys.stdout.write('\n    ')
        for f in os.listdir(pkg_dir+"/msg"):
            if f.endswith(".msg"):
                tmp_msg_file =  build_dir + "/msg/" + f
                msg_file = pkg_dir + "/msg/" + f
                msg_definition = open(msg_file).readlines()
                # add to list of messages
                if not os.path.exists(tmp_msg_file):
                    print('%s,'%f[0:-4], end='')
                    messages.append( Message(f[0:-4], package, msg_definition) )
                else:
                    tmp_msg_definition = open(tmp_msg_file).readlines()
                    tmp_msg_md5 = hashlib_md5sum(f[0:-4], package, tmp_msg_definition)
                    msg_md5 = hashlib_md5sum(f[0:-4], package, msg_definition)
                    if msg_md5 != tmp_msg_md5:
                       print('%s,'%f[0:-4], end='')
                       messages.append( Message(f[0:-4], package, msg_definition) )

    # find the services in this package
    if (os.path.exists(pkg_dir+"/srv/")):
        if messages == list():
            print('Exporting %s\n'%package)
        else:
            print('\n')
        sys.stdout.write('  Services:')
        sys.stdout.write('\n    ')
        for f in os.listdir(pkg_dir+"/srv"):
            if f.endswith(".srv"):
                tmp_srv_file = build_dir + "/srv/" + f
                srv_file = pkg_dir + "/srv/" + f
                # add to list of messages
                srv_definition = open(srv_file).readlines()
                if not os.path.exists(tmp_srv_file):
                    print('%s,'%f[0:-4], end='')
                    messages.append( Service(f[0:-4], package, srv_definition) )
                else:
                    tmp_srv_definition = open(tmp_srv_file).readlines()
                    tmp_srv_md5 = hashlib_md5sum(f[0:-4], package, tmp_srv_definition)
                    srv_md5 = hashlib_md5sum(f[0:-4], package, srv_definition)
                    if srv_md5 != tmp_srv_md5:
                       print('%s,'%f[0:-4], end='')
                       messages.append( Service(f[0:-4], package, srv_definition) )
        print('\n')
    elif messages != list():
        print('\n')

    # generate for each message
    output_path = output_path + "/" + package
    for msg in messages:
        if not os.path.exists(output_path):
            os.makedirs(output_path)
        header = open(output_path + "/" + msg.name + ".h", "w")
        msg.make_header(header)
        header.close()

def messages_generate(path):
    # gimme messages
    failed = []
    mydir = sys.argv[1] + "/msgs"
    builddir = sys.argv[1] + "/build/CMake/gcc_msgs"
    for d in sorted(os.listdir(mydir)):
	    try:
	        MakeLibrary(mydir + "/" + d, d, path, builddir + "/" + d)
	    except Exception as e:
	        failed.append(d + " ("+str(e)+")")
	        print('[%s]: Unable to build messages: %s\n' % (d, str(e)))
	        print(traceback.format_exc())

    print('\n')
    if len(failed) > 0:
        print('*** Warning, failed to generate libraries for the following packages: ***')
        for f in failed:
            print('    %s'%f)
        raise Exception("Failed to generate libraries for: " + str(failed))
    print('\n')

def MakeSubscribers(messages, output_path):
    # generate for each message
    for msg in messages:
        if not os.path.exists(output_path):
            os.makedirs(output_path)
        f = open(output_path + "/rostopic_subscribers.cpp", "w")
        f.write('#ifndef _TINYROS_ROSTOPIC_SUBSCRIBERS_h\n')
        f.write('#define _TINYROS_ROSTOPIC_SUBSCRIBERS_h\n')
        f.write('\n')
        f.write('#include "tiny_ros/ros.h"\n')
        for i in messages:
            f.write('#include "tiny_ros/%s.h"\n' % i)
        f.write('\n')
        f.write('namespace tinyros\n')
        f.write('{\n')
        f.write('template<typename MsgT>\n')
        f.write('class EchoSubscriber: public tinyros::Subscriber_\n')
        f.write('{\n')
        f.write('public:\n')
        f.write('MsgT msg;\n')
        f.write('  virtual void callback(unsigned char* data)\n')
        f.write('  {\n')
        f.write('    MsgT tmsg;\n')
        f.write('    tmsg.deserialize(data);\n')
        f.write('    spdlog::get("logger")->info("{0}[{1}]->>{2}", topic_, getMsgType(), tmsg.echo());\n')
        f.write('  }\n')
        f.write('  virtual std::string getMsgType()\n')
        f.write('  {\n')
        f.write('    return this->msg.getType();\n')
        f.write('  }\n')
        f.write('  virtual std::string getMsgMD5()\n')
        f.write('  {\n')
        f.write('    return this->msg.getMD5();\n')
        f.write('  }\n')
        f.write('  virtual int getEndpointType()\n')
        f.write('  {\n')
        f.write('    return tinyros_msgs::TopicInfo::ID_SUBSCRIBER;\n')
        f.write('  }\n')
        f.write('};\n\n')
        f.write('static std::map<std::string, tinyros::Subscriber_*> rostopic_subscribers = {\n')
        for s in messages:
	        pkg = s[0:s.find('/')]
	        name = s[s.find('/')+1:]
	        f.write('    {"%s", new EchoSubscriber<%s::%s>()},\n' % (s, pkg, name))
        f.write('};\n\n')
        f.write('}\n')
        f.write('#endif\n\n')
        f.close()
        
def subscribers_generate(path):
    # gimme messages
    failed = []
    messages = list()
    mydir = sys.argv[1] + "/msgs"
    for d in sorted(os.listdir(mydir)):
	    try:
		    if os.path.exists(mydir + "/" + d +"/msg"):
		        for f in os.listdir(mydir + "/" + d +"/msg"):
		            messages.append(d + "/" + f[0:-4])
	    except Exception as e:
	        failed.append(d + " ("+str(e)+")")
	        print('[%s]: Unable to append messages: %s\n' % (d, str(e)))
	        print(traceback.format_exc())

    MakeSubscribers(messages, sys.argv[3] + "/tools/rostopic")

    print('\n')
    if len(failed) > 0:
        print('*** Warning, failed to generate subscribers for the following packages: ***')
        for f in failed:
            print('    %s'%f)
        raise Exception("Failed to generate subscribers for: " + str(failed))
    print('\n')

def roslib_copy_roslib_files(path):
    if not os.path.exists(path+"ros"):
        os.makedirs(path+"ros")
    files = ['duration.cpp',
             'time.cpp',
             'ros.cpp',
             'log.cpp',
             'ros.h',
             'ros/duration.h',
             'ros/msg.h',
             'ros/log.h',
             'ros/node_handle_base.h',
             'ros/node_handle_udp.h',
             'ros/node_handle.h',
             'ros/publisher.h',
             'ros/subscriber.h',
             'ros/service_server.h',
             'ros/service_client.h',
             'ros/threadpool.h',
             'ros/hardware.h',
             'ros/hardware_udp.h',
             'ros/hardware_tcp.h',
             'ros/time.h']

    mydir = sys.argv[3] + "/roslib/gcc/"
    for f in files:
        if not os.path.exists(path+f):
            shutil.copy(mydir+f, path+f)
        else:
            tmp_file_definition = open(path+f).readlines()
            file_definition = open(mydir+f).readlines()
            tmp_file_md5 = hashlib_md5sum_definition(tmp_file_definition)
            file_md5 = hashlib_md5sum_definition(file_definition)
            if file_md5 != tmp_file_md5:
                shutil.copy(mydir+f, path+f)

def roslib_copy_examples_files(path):
    if not os.path.exists(path+"examples/publisher"):
        os.makedirs(path+"examples/publisher")
    if not os.path.exists(path+"examples/subscriber"):
        os.makedirs(path+"examples/subscriber")
    if not os.path.exists(path+"examples/service"):
        os.makedirs(path+"examples/service")
    if not os.path.exists(path+"examples/service_client"):
        os.makedirs(path+"examples/service_client")
    files = ['examples/publisher/ExamplePublisher.cpp',
             'examples/subscriber/ExampleSubscriber.cpp',
             'examples/service/ExampleService.cpp',
             'examples/service_client/ExampleServiceClient.cpp']

    mydir = sys.argv[3] + "/"
    for f in files:
        shutil.copy(mydir+f, path+f)

# Enforce correct inputs
if (len(sys.argv) < 3):
    print(__usage__)
    exit(1)

# Sanitize output path
path = sys.argv[2]
if path[-1] == "/":
    path = path[0:-1]
print("\nExporting to %s" % path)

roslib_copy_roslib_files(path+"/tiny_ros/")
roslib_copy_examples_files(path+"/tiny_ros/")
messages_generate(path+"/tiny_ros/")
subscribers_generate(path+"/tiny_ros/")
if os.path.exists(sys.argv[1] + "/build/CMake/gcc_msgs"):
    shutil.rmtree(sys.argv[1] + "/build/CMake/gcc_msgs")
shutil.copytree(sys.argv[1] + "/msgs", sys.argv[1] + "/build/CMake/gcc_msgs")
