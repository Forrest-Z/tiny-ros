#!/usr/bin/python

from __future__ import print_function

__usage__ = """
python make_library.py <output_path>
"""
import traceback

import os, sys, re, hashlib, itertools

# for copying files
import shutil

def type_to_var(ty):
    lookup = {
        1 : 'uint8',
        2 : 'uint16',
        4 : 'uint32',
        8 : 'uint64',
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

def primitive_type(base_type):
    if base_type in ['bool', 'byte', 'int8', 'int16', 'int32', 'int64', 'float32',\
                     'uint8', 'uint16', 'uint32', 'uint64', 'float64', 'string']:
        return True
    else:
        return False

def parse_type(msg_type):
    if not msg_type:
        raise ValueError("Invalid empty type")
    if '[' in msg_type:
        var_length = msg_type.endswith('[]')
        splits = msg_type.split('[')
        if len(splits) > 2:
            raise ValueError("Currently only support 1-dimensional array types: %s"%msg_type)
        if var_length:
            return msg_type[:-2], True, None
        else:
            try:
                length = int(splits[1][:-1])
                return splits[0], True, length
            except ValueError:
                raise ValueError("Invalid array dimension: [%s]"%splits[1][:-1])
    else:
        return msg_type, False, None

def primitive_default_value(field_type):
    if field_type in ['byte', 'int8', 'int16', 'int32', 'int64',\
                      'uint8', 'uint16', 'uint32', 'uint64']:
        return '0'
    elif field_type in ['float32', 'float64']:
        return '0.0'
    elif field_type == 'bool':
        return 'false'
    elif field_type == 'string':
        return '""'
    elif field_type.endswith(']'): # array type
        base_type, is_array, array_len = parse_type(field_type)
        if array_len is None: #var-length
            if primitive_type(base_type):
                return ('[]%s{}' % base_type)
            else:
                return ('[]*%s{}' % base_type)
        else: # fixed-length, fill values
            def_val = primitive_default_value(base_type)
            if primitive_type(base_type):
                def_h = ('[%s]%s' % (array_len, base_type))
            else:
                def_h = ('[%s]*%s' % (array_len, base_type))
            return (def_h + '{' + ', '.join(itertools.repeat(def_val, array_len)) + '}')
    else:
        try:
            type_package, type_name = field_type.split(".")
        except:
            type_package = None
            type_name = field_type
        if type_package != None:
            return ('%s.New%s()' % (type_package, type_name))
        else:
            return ('New%s()\n' % (type_name))
 
#####################################################################
# Data Types

class EnumerationType:
    """ For data values. """

    def __init__(self, cls, name, ty, value):
        self.cls = cls
        self.name = name
        self.type = ty
        self.value = value

    def make_declaration(self, f):
        f.write('func Go_%s() (%s) { return %s }\n' % (self.name, self.type, self.value))

class PrimitiveDataType:
    def __init__(self, cls, name, ty, bytes):
        self.cls = cls
        self.name = name
        self.type = ty
        self.bytes = bytes

    def make_constructor(self, f, trailer):
        f.write('    new%s.Go_%s = %s\n' % (self.cls, self.name, primitive_default_value(self.type)))

    def make_initializer(self, f, trailer):
        f.write('    self.Go_%s = %s\n' % (self.name, primitive_default_value(self.type)))

    def make_declaration(self, f):
        f.write('    Go_%s %s `json:"%s"`\n' % (self.name, self.type, self.name) )

    def serialize(self, f, header):
        cn = self.name.replace("[","").replace("]","").split(".")[-1]
        if self.type == 'float32':
            f.write('%s    bits_%s := math.Float32bits(self.Go_%s)\n' % (header, cn, self.name) )
            f.write('%s    binary.LittleEndian.PutUint32(buff[offset:], bits_%s)\n' % (header, cn) )
        elif self.type == 'float64':
            f.write('%s    bits_%s := math.Float64bits(self.Go_%s)\n' % (header, cn, self.name) )
            f.write('%s    binary.LittleEndian.PutUint64(buff[offset:], bits_%s)\n' % (header, cn) )
        elif self.type == 'bool':
            f.write('%s    if self.Go_%s {\n' % (header, self.name) )
            f.write('%s        buff[offset] = byte(0x01)\n' % (header) )
            f.write('%s    } else {\n' % (header) )
            f.write('%s        buff[offset] = byte(0x00)\n' % (header) )
            f.write('%s    }\n' % (header) )
        elif self.type == 'int8':
            for i in range(self.bytes):
                f.write('%s    buff[offset + %d] = byte(uint8(self.Go_%s >> (8 * %d)) & 0xFF)\n' % (header, i, self.name, i))
        else:
            for i in range(self.bytes):
                f.write('%s    buff[offset + %d] = byte((self.Go_%s >> (8 * %d)) & 0xFF)\n' % (header, i, self.name, i))
        f.write('%s    offset += %s\n' % (header, self.bytes) )
        
    def deserialize(self, f, header):
        cn = self.name.replace("[","").replace("]","").split(".")[-1]
        if self.type == 'float32':
            f.write('%s    bits_%s := binary.LittleEndian.Uint32(buff[offset:])\n' % (header, cn))
            f.write('%s    self.Go_%s = math.Float32frombits(bits_%s)\n' % (header, self.name, cn)) 
        elif self.type == 'float64':
            f.write('%s    bits_%s := binary.LittleEndian.Uint64(buff[offset:])\n' % (header, cn))
            f.write('%s    self.Go_%s = math.Float64frombits(bits_%s)\n' % (header, self.name, cn)) 
        elif self.type == 'bool':
            f.write('%s    if (buff[offset] & 0xFF) != 0 {\n' % (header) )
            f.write('%s        self.Go_%s = true\n' % (header, self.name) )
            f.write('%s    } else {\n' % (header) )
            f.write('%s        self.Go_%s = false\n' % (header, self.name) )
            f.write('%s    }\n' % (header) )
        else:
            f.write('%s    self.Go_%s = %s(buff[offset + 0] & 0xFF) << (8 * 0)\n' % (header, self.name, self.type) )
            for i in range(self.bytes-1):
                f.write('%s    self.Go_%s |= %s(buff[offset + %d] & 0xFF) << (8 * %d)\n' % (header, self.name, self.type, i+1, i+1) )
        f.write('%s    offset += %s\n' % (header, self.bytes) )

    def serializedLength(self, f, header):
        f.write('%s    length += %s\n' % (header, self.bytes))

    def echo(self, f, header, trailer):
        pass

class MessageDataType(PrimitiveDataType):
    def make_declaration(self, f):
        f.write('    Go_%s *%s `json:"%s"`\n' % (self.name, self.type, self.name) )
    
    def make_constructor(self, f, trailer):
        try:
            type_package, type_name = self.type.split(".")
        except:
            type_package = None
            type_name = self.type
        if type_package != None:
            f.write('    new%s.Go_%s = %s.New%s()\n' % (self.cls, self.name, type_package, type_name))
        else:
            f.write('    new%s.Go_%s = New%s()\n' % (self.cls, self.name, type_name))
    
    def make_initializer(self, f, trailer):
        try:
            type_package, type_name = self.type.split(".")
        except:
            type_package = None
            type_name = self.type
        if type_package != None:
            f.write('    self.Go_%s = %s.New%s()\n' % (self.name, type_package, type_name))
        else:
            f.write('    self.Go_%s = New%s()\n' % (self.name, type_name))

    def serialize(self, f, header):
        f.write('%s    offset += self.Go_%s.Go_serialize(buff[offset:])\n' % (header, self.name))

    def deserialize(self, f, header):
        f.write('%s    offset += self.Go_%s.Go_deserialize(buff[offset:])\n' % (header, self.name))

    def serializedLength(self, f, header):
        f.write('%s    length += self.Go_%s.Go_serializedLength()\n' % (header, self.name))

    def echo(self, f, header, trailer):
        pass

class StringDataType(PrimitiveDataType):
    def make_constructor(self, f, trailer):
        f.write('    new%s.Go_%s = ""\n' % (self.cls, self.name))
 
    def make_initializer(self, f, trailer):
        f.write('    self.Go_%s = ""\n' % (self.name))

    def make_declaration(self, f):
        f.write('    Go_%s %s `json:"%s"`\n' % (self.name, self.type, self.name) )

    def serialize(self, f, header):
        cn = self.name.replace("[","").replace("]","")
        f.write('%s    length_%s := len(self.Go_%s)\n' % (header, cn, self.name))
        f.write('%s    buff[offset + 0] = byte((length_%s >> (8 * 0)) & 0xFF)\n' % (header, cn))
        f.write('%s    buff[offset + 1] = byte((length_%s >> (8 * 1)) & 0xFF)\n' % (header, cn))
        f.write('%s    buff[offset + 2] = byte((length_%s >> (8 * 2)) & 0xFF)\n' % (header, cn))
        f.write('%s    buff[offset + 3] = byte((length_%s >> (8 * 3)) & 0xFF)\n' % (header, cn))
        f.write('%s    offset += 4\n' % header)
        f.write('%s    copy(buff[offset:(offset+length_%s)], self.Go_%s)\n' % (header, cn, self.name))
        f.write('%s    offset += length_%s\n' % (header, cn))

    def deserialize(self, f, header):
        cn = self.name.replace("[","").replace("]","")
        f.write('%s    length_%s := int(buff[offset + 0] & 0xFF) << (8 * 0)\n' % (header, cn))
        f.write('%s    length_%s |= int(buff[offset + 1] & 0xFF) << (8 * 1)\n' % (header, cn))
        f.write('%s    length_%s |= int(buff[offset + 2] & 0xFF) << (8 * 2)\n' % (header, cn))
        f.write('%s    length_%s |= int(buff[offset + 3] & 0xFF) << (8 * 3)\n' % (header, cn))
        f.write('%s    offset += 4\n' % header)
        f.write('%s    self.Go_%s = string(buff[offset:(offset+length_%s)])\n' % (header, self.name, cn))
        f.write('%s    offset += length_%s\n' % (header, cn))

    def serializedLength(self, f, header):
        cn = self.name.replace("[","").replace("]","")
        f.write('%s    length_%s := len(self.Go_%s)\n' % (header, cn, self.name))
        f.write('%s    length += 4\n' % header)
        f.write('%s    length += length_%s\n' % (header,cn))

    def echo(self, f, header, trailer):
        pass

class TimeDataType(PrimitiveDataType):

    def __init__(self, cls, name, ty, bytes):
        self.cls = cls
        self.name = name
        self.type = ty
        self.sec = PrimitiveDataType(self.cls, name+'.Go_sec','uint32',4)
        self.nsec = PrimitiveDataType(self.cls, name+'.Go_nsec','uint32',4)

    def make_constructor(self, f, trailer):
        if self.type == 'rostime.Duration':
            f.write('    new%s.Go_%s = rostime.NewDuration()\n' % (self.cls, self.name))
        else:
            f.write('    new%s.Go_%s = rostime.NewTime()\n' % (self.cls, self.name))

    def make_initializer(self, f, trailer):
        if self.type == 'rostime.Duration':
            f.write('    self.Go_%s = rostime.NewDuration()\n' % (self.name))
        else:
            f.write('    self.Go_%s = rostime.NewTime()\n' % (self.name))

    def make_declaration(self, f):
        f.write('    Go_%s *%s `json:"%s"`\n' % (self.name, self.type, self.name) )

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
        pass

class ArrayDataType(PrimitiveDataType):

    def __init__(self, cls, name, ty, bytes, cls1, array_size=None):
        self.name = name
        self.type = ty
        self.bytes = bytes
        self.size = array_size
        self.cls = cls
        self.cls1 = cls1

    def make_constructor(self, f, trailer):
        if self.size == None:
            f.write('    new%s.Go_%s = %s\n' % (self.cls, self.name, primitive_default_value(self.type + "[]")))
        else:
            f.write('    new%s.Go_%s = %s\n' % (self.cls, self.name, primitive_default_value(self.type + "[" + str(self.size) + "]")))

    def make_initializer(self, f, trailer):
        if self.size == None:
            f.write('    self.Go_%s = %s\n' % (self.name, primitive_default_value(self.type + "[]")))
        else:
            f.write('    self.Go_%s = %s\n' % (self.name, primitive_default_value(self.type + "[" + str(self.size) + "]")))

    def make_declaration(self, f):
        if self.size == None:
            if primitive_type(self.type):
                f.write('    Go_%s []%s `json:"%s"`\n' % (self.name, self.type, self.name))
            else:
                f.write('    Go_%s []*%s `json:"%s"`\n' % (self.name, self.type, self.name))
        else:
            if primitive_type(self.type):
                f.write('    Go_%s [%s]%s `json:"%s"`\n' % (self.name, self.size, self.type, self.name))
            else:
                f.write('    Go_%s [%s]*%s `json:"%s"`\n' % (self.name, self.size, self.type, self.name))

    def serialize(self, f, header):
        c = self.cls1(self.cls, self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            # serialize length
            f.write('%s    length_%s := len(self.Go_%s)\n' % (header, self.name, self.name) )
            f.write('%s    buff[offset + 0] = byte((length_%s >> (8 * 0)) & 0xFF)\n' % (header, self.name))
            f.write('%s    buff[offset + 1] = byte((length_%s >> (8 * 1)) & 0xFF)\n' % (header, self.name))
            f.write('%s    buff[offset + 2] = byte((length_%s >> (8 * 2)) & 0xFF)\n' % (header, self.name))
            f.write('%s    buff[offset + 3] = byte((length_%s >> (8 * 3)) & 0xFF)\n' % (header, self.name))
            f.write('%s    offset += 4\n' % header)
            f.write('%s    for i := 0; i < length_%s; i++ {\n' % (header, self.name))
            c.serialize(f, header + '    ')
            f.write('%s    }\n' % header)
        else:
            f.write('%s    for i := 0; i < %d; i++ {\n' % (header, self.size) )
            c.serialize(f ,header + '    ')
            f.write('%s    }\n' % header)

    def deserialize(self, f, header):
        c = self.cls1(self.cls, self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            # deserialize length
            f.write('%s    length_%s := int(buff[offset + 0] & 0xFF) << (8 * 0)\n' % (header, self.name))
            f.write('%s    length_%s |= int(buff[offset + 1] & 0xFF) << (8 * 1)\n' % (header, self.name))
            f.write('%s    length_%s |= int(buff[offset + 2] & 0xFF) << (8 * 2)\n' % (header, self.name))
            f.write('%s    length_%s |= int(buff[offset + 3] & 0xFF) << (8 * 3)\n' % (header, self.name))
            f.write('%s    offset += 4\n' % header)
            if primitive_type(self.type):
                f.write('%s    self.Go_%s = make([]%s, length_%s)\n' % (header, self.name, self.type, self.name))
            else:
                package = None
                name = None
                try:
                    package, name = self.type.split(".")
                except:
                    package = None
                    name = self.type
                f.write('%s    self.Go_%s = make([]*%s, length_%s)\n' % (header, self.name, self.type, self.name))
                f.write('%s    for i := 0; i < length_%s; i++ {\n' % (header, self.name))
                if package != None:
                    f.write('%s        self.Go_%s[i] = %s.New%s()\n' % (header, self.name, package, name))
                else:
                    f.write('%s        self.Go_%s[i] = New%s()\n' % (header, self.name, name))
                f.write('%s    }\n' % header)
            f.write('%s    for i := 0; i < length_%s; i++ {\n' % (header, self.name))
            c.deserialize(f, header + '    ')
            f.write('%s    }\n' % header)
        else:
            f.write('%s    for i := 0; i < %d; i++ {\n' % (header, self.size) )
            c.deserialize(f, header + '    ')
            f.write('%s    }\n' % header)

    def serializedLength(self, f, header):
        c = self.cls1(self.cls, self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            # serialize length
            f.write('%s    length += 4\n' % header)
            f.write('%s    length_%s := len(self.Go_%s)\n' % (header, self.name, self.name) )
            f.write('%s    for i := 0; i < length_%s; i++ {\n' % (header, self.name))
            c.serializedLength(f, header + '    ')
            f.write('%s    }\n' % header)
        else:
            f.write('%s    for i := 0; i < %d; i++ {\n' % (header, self.size) )
            c.serializedLength(f, header + '    ')
            f.write('%s    }\n' % header)

    def echo(self, f, header, trailer):
        pass

ROS_TO_EMBEDDED_TYPES = {
    'bool'    :   ('bool',              1, PrimitiveDataType, []),
    'byte'    :   ('byte',            1, PrimitiveDataType, []),
    'int8'    :   ('int8',            1, PrimitiveDataType, []),
    'char'    :   ('byte',              1, PrimitiveDataType, []),
    'uint8'   :   ('uint8',           1, PrimitiveDataType, []),
    'int16'   :   ('int16',           2, PrimitiveDataType, []),
    'uint16'  :   ('uint16',          2, PrimitiveDataType, []),
    'int32'   :   ('int32',           4, PrimitiveDataType, []),
    'uint32'  :   ('uint32',          4, PrimitiveDataType, []),
    'int64'   :   ('int64',           8, PrimitiveDataType, []),
    'uint64'  :   ('uint64',          8, PrimitiveDataType, []),
    'float32' :   ('float32',             4, PrimitiveDataType, ['encoding/binary', 'math']),
    'float64' :   ('float64',            8, PrimitiveDataType, ['encoding/binary', 'math']),
    'time'    :   ('rostime.Time',         8, TimeDataType, ['tiny_ros/tinyros/time']),
    'duration':   ('rostime.Duration',     8, TimeDataType, ['tiny_ros/tinyros/time']),
    'string'  :   ('string',             0, StringDataType, []),
    'Header'  :   ('std_msgs.Header',  0, MessageDataType, ['tiny_ros/std_msgs'])
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
                self.enums.append( EnumerationType(self.name, name, ty, value))
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
                if "tiny_ros/" + type_package not in self.includes:
                    self.includes.append("tiny_ros/" + type_package)
                cls = MessageDataType
                if type_package == self.package:
                    code_type = type_name
                else:
                    code_type = type_package + "." + type_name
                size = 0
            if type_array:
                self.data.append( ArrayDataType(self.name, name, code_type, size, cls, type_array_size ) )
            else:
                self.data.append( cls(self.name, name, code_type, size) )

    def _write_id(self, f):
        f.write('')

    def _write_id_constructor(self, f):
        f.write('')

    def _write_id_initializer(self, f):
        f.write('')

    def _write_id_serializer(self, f):
        f.write('')

    def _write_id_deserializer(self, f):
        f.write('')

    def _write_getID(self, f):
        f.write('func (self *%s) Go_getID() (uint32) { return 0 }\n' % (self.name))

    def _write_setID(self, f):
        f.write('func (self *%s) Go_setID(id uint32) { }\n' % (self.name))

    def _write_serializer(self, f):
        f.write('func (self *%s) Go_serialize(buff []byte) (int) {\n' % (self.name))
        f.write('    offset := 0\n')
        self._write_id_serializer(f)
        for d in self.data:
            d.serialize(f, '')
        f.write('    return offset\n');
        f.write('}\n')
        f.write('\n')

    def _write_deserializer(self, f):
        # deserializer
        f.write('func (self *%s) Go_deserialize(buff []byte) (int) {\n' % (self.name))
        f.write('    offset := 0\n')
        self._write_id_deserializer(f)
        for d in self.data:
            d.deserialize(f, '')
        f.write('    return offset\n');
        f.write('}\n')
        f.write('\n')

    def _write_serializedLength(self, f):
        # serializedLength
        f.write('func (self *%s) Go_serializedLength() (int) {\n' % (self.name))
        f.write('    length := 0\n')
        for d in self.data:
            d.serializedLength(f, '')
        f.write('    return length\n');
        f.write('}\n')
        f.write('\n')

    def _write_echo(self, f):
        f.write('func (self *%s) Go_echo() (string) { \n' % (self.name))
        f.write('    data, _ := json.Marshal(self)\n')
        f.write('    return string(data)\n')
        f.write('}\n\n')

    def _write_std_includes(self, f):
        f.write('    "encoding/json"\n')

    def _write_msg_includes(self,f):
        for i in self.includes:
            if i != "tiny_ros/" + self.package:
                f.write('    "%s"\n' % i)

    def _write_constructor(self, f):
        f.write('func New%s() (*%s) {\n' % (self.name, self.name))
        f.write('    new%s := new(%s)\n' % (self.name, self.name))
        if self.data:
            for d in self.data[:-1]:
                d.make_constructor(f, ',')
            self.data[-1].make_constructor(f, '')
        self._write_id_constructor(f)
        f.write('    return new%s\n' % self.name)
        f.write('}\n\n')

    def _write_initializer(self, f):
        f.write('func (self *%s) Go_initialize() {\n' % (self.name))
        if self.data:
            for d in self.data[:-1]:
                d.make_initializer(f, ',')
            self.data[-1].make_initializer(f, '')
        self._write_id_initializer(f)
        f.write('}\n\n')

    def _write_data(self, f):
        for e in self.enums:
            e.make_declaration(f)
        f.write('\ntype %s struct {\n' % self.name)
        self._write_id(f)
        for d in self.data:
            d.make_declaration(f)
        f.write('}\n\n')
   
    def _write_getType(self, f):
        f.write('func (self *%s) Go_getType() (string) { return "%s/%s" }\n'%(self.name, self.package, self.name))

    def _write_getMD5(self, f):
        f.write('func (self *%s) Go_getMD5() (string) { return "%s" }\n'% (self.name, self.md5))

    def _write_impl(self, f):
        self._write_data(f)
        self._write_constructor(f)
        self._write_initializer(f)
        self._write_serializer(f)
        self._write_deserializer(f)
        self._write_serializedLength(f)
        self._write_echo(f)
        self._write_getType(f)
        self._write_getMD5(f)
        self._write_getID(f)
        self._write_setID(f)
        f.write('\n')

    def make_header(self, f):
        f.write('package %s\n\n' % self.package)
        f.write('import (\n')
        self._write_std_includes(f)
        self._write_msg_includes(f)
        f.write(')\n\n')
        self._write_impl(f)

class Service:
    def __init__(self, name, package, definition):
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
        f.write('package %s\n\n' % self.package)
        f.write('import (\n')
        self.req._write_std_includes(f)
        includes = self.req.includes
        includes.extend(self.resp.includes)
        includes = list(set(includes))
        for inc in includes:
            if inc != "tiny_ros/" + self.package:
                f.write('    "%s"\n' % inc)
        f.write(')\n\n')

        def write_id(out):
            out.write('    __id__ uint32 `json:"__id__"`\n')
        _write_id = lambda out: write_id(out)
        self.req._write_id = _write_id
        self.resp._write_id = _write_id

        def write_id_constructor_req(out):
            out.write('    new%sRequest.__id__ = 0\n' % (self.name))
        _write_id_constructor_req = lambda out: write_id_constructor_req(out)
        self.req._write_id_constructor = _write_id_constructor_req
        
        def write_id_constructor_resp(out):
            out.write('    new%sResponse.__id__ = 0\n' % (self.name))
        _write_id_constructor_resp = lambda out: write_id_constructor_resp(out)
        self.resp._write_id_constructor = _write_id_constructor_resp

        def write_id_initializer_req(out):
            out.write('    self.__id__ = 0\n')
        _write_id_initializer_req = lambda out: write_id_initializer_req(out)
        self.req._write_id_initializer = _write_id_initializer_req
        
        def write_id_initializer_resp(out):
            out.write('    self.__id__ = 0\n')
        _write_id_initializer_resp = lambda out: write_id_initializer_resp(out)
        self.resp._write_id_initializer = _write_id_initializer_resp

        def write_id_serializer(out):
            out.write('    buff[offset + 0] = byte((self.__id__ >> (8 * 0)) & 0xFF)\n')
            out.write('    buff[offset + 1] = byte((self.__id__ >> (8 * 1)) & 0xFF)\n')
            out.write('    buff[offset + 2] = byte((self.__id__ >> (8 * 2)) & 0xFF)\n')
            out.write('    buff[offset + 3] = byte((self.__id__ >> (8 * 3)) & 0xFF)\n')
            out.write('    offset += 4\n')
        _write_id_serializer = lambda out: write_id_serializer(out)
        self.req._write_id_serializer = _write_id_serializer
        self.resp._write_id_serializer = _write_id_serializer

        def write_id_deserializer(out):
            out.write('    self.__id__ =  uint32(buff[offset + 0] & 0xFF) << (8 * 0)\n')
            out.write('    self.__id__ |=  uint32(buff[offset + 1] & 0xFF) << (8 * 1)\n')
            out.write('    self.__id__ |=  uint32(buff[offset + 2] & 0xFF) << (8 * 2)\n')
            out.write('    self.__id__ |=  uint32(buff[offset + 3] & 0xFF) << (8 * 3)\n')
            out.write('    offset += 4\n')
        _write_id_deserializer = lambda out: write_id_deserializer(out)
        self.req._write_id_deserializer = _write_id_deserializer
        self.resp._write_id_deserializer = _write_id_deserializer

        def write_getID_req(out):
            out.write('func (self *%sRequest) Go_getID() (uint32) { return self.__id__ }\n' % (self.name))
        _write_getID_req = lambda out: write_getID_req(out)
        self.req._write_getID = _write_getID_req

        def write_getID_resp(out):
            out.write('func (self *%sResponse) Go_getID() (uint32) { return self.__id__ }\n' % (self.name))
        _write_getID_resp = lambda out: write_getID_resp(out)
        self.resp._write_getID = _write_getID_resp

        def write_setID_req(out):
            out.write('func (self *%sRequest) Go_setID(id uint32) { self.__id__ = id }\n' % (self.name))
        _write_setID_req = lambda out: write_setID_req(out)
        self.req._write_setID = _write_setID_req

        def write_setID_resp(out):
            out.write('func (self *%sResponse) Go_setID(id uint32) { self.__id__ = id }\n' % (self.name))
        _write_setID_resp = lambda out: write_setID_resp(out)
        self.resp._write_setID = _write_setID_resp

        def write_type_req(out):
            out.write('func (self *%sRequest) Go_getType() (string) { return "%s/%s" }\n' % (self.name, self.package, self.name))
        _write_type_req = lambda out: write_type_req(out)
        self.req._write_getType = _write_type_req

        def write_type_resp(out):
            out.write('func (self *%sResponse) Go_getType() (string) { return "%s/%s" }\n' % (self.name, self.package, self.name))
        _write_type_resp = lambda out: write_type_resp(out)
        self.resp._write_getType = _write_type_resp

        f.write('\n')
        self.req._write_impl(f)
        f.write('\n///////////////////////////////////////////////////////////////////////////\n\n')
        self.resp._write_impl(f)
        f.write('\n')

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
        header = open(output_path + "/" + msg.name + ".go", "w")
        msg.make_header(header)
        header.close()

def messages_generate(path):
    # gimme messages
    failed = []
    mydir = sys.argv[1] + "/msgs"
    builddir = sys.argv[1] + "/build/CMake/go_msgs"
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

def roslib_copy_roslib_files(path):
    if not os.path.exists(path+"tinyros"):
        os.makedirs(path+"tinyros")
    if not os.path.exists(path+"tinyros/time"):
        os.makedirs(path+"tinyros/time")
    files = ['tinyros/Msg.go',
             'tinyros/time/Time.go',
             'tinyros/time/Duration.go',
             'tinyros/Hardware.go',
             'tinyros/NodeHandle.go',
             'tinyros/Publisher.go',
             'tinyros/Subscriber.go',
             'tinyros/ServiceClient.go',
             'tinyros/ThreadPool.go',
             'tinyros/ServiceServer.go']

    mydir = sys.argv[3] + "/roslib/go/"
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
    files = ['examples/publisher/ExamplePublisher.go',
             'examples/subscriber/ExampleSubscriber.go',
             'examples/service/ExampleService.go',
             'examples/service_client/ExampleServiceClient.go']

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

roslib_copy_roslib_files(path+"/src/tiny_ros/")
roslib_copy_examples_files(path+"/")
messages_generate(path+"/src/tiny_ros/")
if os.path.exists(sys.argv[1] + "/build/CMake/go_msgs"):
    shutil.rmtree(sys.argv[1] + "/build/CMake/go_msgs")
shutil.copytree(sys.argv[1] + "/msgs", sys.argv[1] + "/build/CMake/go_msgs")
