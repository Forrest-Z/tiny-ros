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
                      'char', 'uint8', 'uint16', 'uint32', 'uint64']:
        return '0'
    elif field_type in ['float32', 'float64']:
        return '0.'
    elif field_type == 'bool':
        return 'False'
    elif field_type == 'string':
        return "''"
    elif field_type.endswith(']'): # array type
        base_type, is_array, array_len = parse_type(field_type)
        if base_type in ['char', 'uint8']:
            # strings, char[], and uint8s are all optimized to be strings
            if array_len is not None:
                return "chr(0)*%s"%array_len
            else:
                return "''"
        elif array_len is None: #var-length
            return '[]'
        else: # fixed-length, fill values
            def_val = primitive_default_value(base_type)
            #return '[' + ','.join(itertools.repeat(def_val, array_len)) + ']'
            return ('[%s for x in range(0, %s)]' % (def_val, array_len))
    else:
        return '%s()'%field_type


#####################################################################
# Data Types

class EnumerationType:
    """ For data values. """

    def __init__(self, name, ty, value):
        self.name = name
        self.type = ty
        self.value = value

    def make_declaration(self, f, header):
        f.write('%s%s = %s\n' % (header, self.name, self.value))

class PrimitiveDataType:
    """ Our datatype is a C/C++ primitive. """

    def __init__(self, name, ty, bytes):
        self.name = name
        self.type = ty
        self.bytes = bytes

    def make_declaration(self, f, trailer):
        f.write('\'%s\'%s' % (self.name, trailer))

    def make_declaration_type(self, f, trailer):
        f.write('\'%s\'%s' % (self.type, trailer))
        
    def make_initializer(self, f, header):
        f.write('%sself.%s = %s\n' % (header, self.name, primitive_default_value(self.type)))

    def serialize(self, f, header):
        cn = self.name.replace("[","").replace("]","").split(".")[-1]
        f.write('%stry:\n' % header)
        if self.type == 'byte' or self.type == 'int8':
             f.write('%s    struct_%s = struct.Struct("<b")\n' % (header, cn))
             f.write('%s    buff.write(struct_%s.pack(self.%s))\n' % (header, cn, self.name))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'uint8' or self.type == 'bool' or self.type == 'char':
             f.write('%s    struct_%s = struct.Struct("<B")\n' % (header, cn))
             f.write('%s    buff.write(struct_%s.pack(self.%s))\n' % (header, cn, self.name))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'int16':
             f.write('%s    struct_%s = struct.Struct("<h")\n' % (header, cn))
             f.write('%s    buff.write(struct_%s.pack(self.%s))\n' % (header, cn, self.name))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'uint16':
             f.write('%s    struct_%s = struct.Struct("<H")\n' % (header, cn))
             f.write('%s    buff.write(struct_%s.pack(self.%s))\n' % (header, cn, self.name))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'int32':
             f.write('%s    struct_%s = struct.Struct("<i")\n' % (header, cn))
             f.write('%s    buff.write(struct_%s.pack(self.%s))\n' % (header, cn, self.name))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'uint32':
             f.write('%s    buff.write(_struct_I.pack(self.%s))\n' % (header, self.name))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'int64':
             f.write('%s    struct_%s = struct.Struct("<q")\n' % (header, cn))
             f.write('%s    buff.write(struct_%s.pack(self.%s))\n' % (header, cn, self.name))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'uint64':
             f.write('%s    struct_%s = struct.Struct("<Q")\n' % (header, cn))
             f.write('%s    buff.write(struct_%s.pack(self.%s))\n' % (header, cn, self.name))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'float32':
             f.write('%s    struct_%s = struct.Struct("<f")\n' % (header, cn))
             f.write('%s    buff.write(struct_%s.pack(self.%s))\n' % (header, cn, self.name))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'float64':
             f.write('%s    struct_%s = struct.Struct("<d")\n' % (header, cn))
             f.write('%s    buff.write(struct_%s.pack(self.%s))\n' % (header, cn, self.name))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        f.write('%sexcept struct.error as ex:\n' % header)
        f.write('%s    print(\'Unable to serialize messages: ' % (header))
        f.write('%s\'%str(ex))\n')
                
    def deserialize(self, f, header):
        cn = self.name.replace("[","").replace("]","").split(".")[-1]
        f.write('%stry:\n' % header)
        if self.type == 'byte' or self.type == 'int8':
             f.write('%s    struct_%s = struct.Struct("<b")\n' % (header, cn))
             f.write('%s    (self.%s,) = struct_%s.unpack(buff[offset:(offset + %s)])\n' % (header, self.name, cn, self.bytes))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'uint8' or self.type == 'bool' or self.type == 'char':
             f.write('%s    struct_%s = struct.Struct("<B")\n' % (header, cn))
             f.write('%s    (self.%s,) = struct_%s.unpack(buff[offset:(offset + %s)])\n' % (header, self.name, cn, self.bytes))
             if self.type == 'bool':
                 f.write('%s    self.%s = bool(self.%s)\n' % (header, self.name, self.name))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'int16':
             f.write('%s    struct_%s = struct.Struct("<h")\n' % (header, cn))
             f.write('%s    (self.%s,) = struct_%s.unpack(buff[offset:(offset + %s)])\n' % (header, self.name, cn, self.bytes))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'uint16':
             f.write('%s    struct_%s = struct.Struct("<H")\n' % (header, cn))
             f.write('%s    (self.%s,) = struct_%s.unpack(buff[offset:(offset + %s)])\n' % (header, self.name, cn, self.bytes))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'int32':
             f.write('%s    struct_%s = struct.Struct("<i")\n' % (header, cn))
             f.write('%s    (self.%s,) = struct_%s.unpack(buff[offset:(offset + %s)])\n' % (header, self.name, cn, self.bytes))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'uint32':
             f.write('%s    (self.%s,) = _struct_I.unpack(buff[offset:(offset + %s)])\n' % (header, self.name, self.bytes))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'int64':
             f.write('%s    struct_%s = struct.Struct("<q")\n' % (header, cn))
             f.write('%s    (self.%s,) = struct_%s.unpack(buff[offset:(offset + %s)])\n' % (header, self.name, cn, self.bytes))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'uint64':
             f.write('%s    struct_%s = struct.Struct("<Q")\n' % (header, cn))
             f.write('%s    (self.%s,) = struct_%s.unpack(buff[offset:(offset + %s)])\n' % (header, self.name, cn, self.bytes))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'float32':
             f.write('%s    struct_%s = struct.Struct("<f")\n' % (header, cn))
             f.write('%s    (self.%s,) = struct_%s.unpack(buff[offset:(offset + %s)])\n' % (header, self.name, cn, self.bytes))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        elif self.type == 'float64':
             f.write('%s    struct_%s = struct.Struct("<d")\n' % (header, cn))
             f.write('%s    (self.%s,) = struct_%s.unpack(buff[offset:(offset + %s)])\n' % (header, self.name, cn, self.bytes))
             f.write('%s    offset += %s\n' % (header, self.bytes))
        f.write('%sexcept struct.error as ex:\n' % header)
        f.write('%s    print(\'Unable to deserialize messages: ' % (header))
        f.write('%s\'%str(ex))\n')

    def serializedLength(self, f, header):
        f.write('%slength += %s\n' % (header, self.bytes))

    def echo(self, f, header, trailer):
        if self.name.find('[') > 0:
            tn = self.name[0:self.name.find('[')]
            cn = self.name.replace("[","").replace("]","").split(".")[-1]
            idx = self.name[self.name.find('[')+1:self.name.find(']')]
            if self.type == 'char':
                f.write('%sstring_echo += \'{"%s' % (header, tn))
                f.write('%s": "%s\'%')
                f.write('(%s,%s)\n' % (idx, self.name))
                f.write('%sstring_echo += \'"}%s\'\n' % (header, trailer))
            else:
                f.write('%sstring_echo += \'{"%s' % (header, tn))
                f.write('%s": %s\'%')
                f.write('(%s,%s)\n' % (idx, self.name))
                f.write('%sstring_echo += \'}%s\'\n' % (header, trailer))
        else:
            if self.type == 'char':
                f.write('%sstring_echo += \'"%s": "' % (header, self.name))
                f.write('%s\'%')
                f.write('%s\n'%self.name)
                f.write('%sstring_echo += \'"%s\'\n' % (header, trailer))
            else:
                f.write('%sstring_echo += \'"%s": ' % (header, self.name))
                f.write('%s\'%')
                f.write('%s\n'%self.name)
                f.write('%sstring_echo += \'%s\'\n' % (header, trailer))

class MessageDataType(PrimitiveDataType):
    """ For when our data type is another message. """

    def make_declaration(self, f, trailer):
        f.write('\'%s\'%s' % (self.name, trailer))

    def make_declaration_type(self, f, trailer):
        f.write('\'%s\'%s' % (self.type, trailer))

    def make_initializer(self, f, header):
        f.write('%sself.%s = %s()\n' % (header, self.name, self.type))

    def serialize(self, f, header):
        f.write('%soffset += self.%s.serialize(buff)\n' % (header, self.name))

    def deserialize(self, f, header):
        f.write('%soffset += self.%s.deserialize(buff[offset:])\n' % (header, self.name))

    def serializedLength(self, f, header):
        f.write('%slength += self.%s.serializedLength()\n' % (header, self.name))

    def echo(self, f, header, trailer):
        if self.name.find('[') > 0:
            tn = self.name[0:self.name.find('[')]
            cn = self.name.replace("[","").replace("]","").split(".")[-1]
            idx = self.name[self.name.find('[')+1:self.name.find(']')]
            f.write('%sstring_echo += \'{"%s' % (header, tn))
            f.write('%s": {\'%')
            f.write('%s\n' % idx)
            f.write('%sstring_echo += self.%s.echo()\n' % (header, self.name))
            f.write('%sstring_echo += \'}}%s\'\n' % (header, trailer))
        else:
            f.write('%sstring_echo += \'"%s": {"\'\n' % (header, self.name))
            f.write('%sstring_echo += self.%s.echo()\n' % (header, self.name))
            f.write('%sstring_echo += \'}%s\'\n' % (header, trailer))

class StringDataType(PrimitiveDataType):
    def make_declaration(self, f, trailer):
        f.write('\'%s\'%s' % (self.name, trailer))

    def make_declaration_type(self, f, trailer):
        f.write('\'%s\'%s' % (self.type, trailer))

    def make_initializer(self, f, header):
        f.write('%sself.%s = \'\'\n' % (header, self.name))

    def serialize(self, f, header):
        f.write('%stry:\n' % header)
        f.write('%s    _x = self.%s\n' % (header, self.name))
        f.write('%s    length = len(_x)\n' % (header))
        f.write('%s    if python3 or type(_x) == unicode:\n' % (header))
        f.write('%s        _x = _x.encode(\'utf-8\')\n' % (header))
        f.write('%s        length = len(_x)\n' % (header))
        f.write('%s    if python3:\n' % (header))
        f.write('%s' % (header))
        f.write('        buff.write(struct.pack(\'<I%sB\'%length, length, *_x))\n')
        f.write('%s    else:\n' % (header))
        f.write('%s' % (header))
        f.write('        buff.write(struct.pack(\'<I%ss\'%length, length, _x))\n')
        f.write('%s    offset += 4\n' % (header))
        f.write('%s    offset += length\n' % (header))
        f.write('%sexcept struct.error as ex:\n' % header)
        f.write('%s    print(\'Unable to serialize messages: ' % (header))
        f.write('%s\'%str(ex))\n')
 
    def deserialize(self, f, header):
        cn = self.name.replace("[","").replace("]","")
        f.write('%stry:\n' % header)
        f.write('%s    (length_%s,) = _struct_I.unpack(buff[offset:(offset + 4)])\n' % (header, cn))
        f.write('%s    offset += 4\n' % (header))
        f.write('%s    if python3:\n' % (header))
        f.write('%s        self.%s = buff[offset:(offset + length_%s)].decode(\'utf-8\')\n' % (header, self.name, cn))
        f.write('%s    else:\n' % (header))
        f.write('%s        self.%s = buff[offset:(offset + length_%s)]\n' % (header, self.name, cn))
        f.write('%s    offset += length_%s\n' % (header, cn))
        f.write('%sexcept struct.error as ex:\n' % header)
        f.write('%s    print(\'Unable to deserialize messages: ' % (header))
        f.write('%s\'%str(ex))\n')

    def serializedLength(self, f, header):
        cn = self.name.replace("[","").replace("]","")
        f.write('%s%s_x = self.%s\n' % (header, cn, self.name))
        f.write('%s%s_length = len(%s_x)\n' % (header, cn, cn))
        f.write('%sif python3 or type(%s_x) == unicode:\n' % (header, cn))
        f.write('%s    %s_x = %s_x.encode(\'utf-8\')\n' % (header,cn, cn))
        f.write('%s    %s_length = len(%s_x)\n' % (header,cn, cn))
        f.write('%slength += 4\n' % (header))
        f.write('%slength += %s_length\n' % (header, cn))

    def echo(self, f, header, trailer):
        f.write('%sstring_echo += \'"%s": "' % (header, self.name))
        f.write('%s"\'%')
        f.write('%s\n' % self.name)
        f.write('%sstring_echo += \'"%s\'\n' % (header, trailer))

class TimeDataType(PrimitiveDataType):

    def __init__(self, name, ty, bytes):
        self.name = name
        self.type = ty
        self.sec = PrimitiveDataType(name+'.sec','uint32',4)
        self.nsec = PrimitiveDataType(name+'.nsec','uint32',4)

    def make_declaration(self, f, trailer):
        f.write('\'%s\'%s' % (self.name, trailer))

    def make_declaration_type(self, f, trailer):
        f.write('\'%s\'%s' % (self.type, trailer))

    def make_initializer(self, f, header):
        f.write('%sself.%s = %s()\n' % (header, self.name, self.type))

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
        f.write('%sstring_echo += \'"%s.sec": ' % (header, self.name))
        f.write('%s, \'%')
        f.write('%s.sec\n' % self.name)
        f.write('%sstring_echo += \'"%s.nsec": ' % (header, self.name))
        f.write('%s\'%')
        f.write('%s.nsec\n' % self.name)
        f.write('%sstring_echo += \'%s\'\n' % (header, trailer))

class ArrayDataType(PrimitiveDataType):

    def __init__(self, name, ty, bytes, cls, array_size=None):
        self.name = name
        self.type = ty
        self.bytes = bytes
        self.size = array_size
        self.cls = cls
 
    def make_declaration(self, f, trailer):
        f.write('\'%s\'%s' % (self.name, trailer))

    def make_declaration_type(self, f, trailer):
        if self.size == None:
            f.write('\'%s[]\'%s' % (self.type, trailer))
        else:
            f.write('\'%s[%s]\'%s' % (self.type, self.size, trailer))

    def make_initializer(self, f, header):
        if self.size == None:
            f.write('%sself.%s = %s\n' % (header, self.name, primitive_default_value(self.type + "[]")))
        else:
            f.write('%sself.%s = %s\n' % (header, self.name, primitive_default_value(self.type + "[" + str(self.size) + "]")))

    def serialize(self, f, header):
        c = self.cls(self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            f.write('%stry:\n' % header)
            f.write('%s    %s_length = len(self.%s)\n' % (header, self.name, self.name))
            f.write('%s    buff.write(_struct_I.pack(%s_length))\n' % (header, self.name))
            f.write('%s    offset += 4\n' % header)
            f.write('%s    for i in range(0, %s_length):\n' % (header, self.name))
            c.serialize(f, header + "        ")
            f.write('%sexcept struct.error as ex:\n' % header)
            f.write('%s    print(\'Unable to serialize messages: ' % (header))
            f.write('%s\'%str(ex))\n')
        else:
            f.write('%stry:\n' % header)
            f.write('%s    for i in range(0, %s):\n' % (header, self.size))
            c.serialize(f, header + "        ")
            f.write('%sexcept struct.error as ex:\n' % header)
            f.write('%s    print(\'Unable to serialize messages: ' % (header))
            f.write('%s\'%str(ex))\n')

    def deserialize(self, f, header):
        c = self.cls(self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            f.write('%stry:\n' % header)
            f.write('%s    (%s_length,) = _struct_I.unpack(buff[offset:(offset + 4)])\n' % (header, self.name))
            f.write('%s    self.%s = [%s for x in range(0, %s_length)]\n' % (header, self.name, primitive_default_value(self.type), self.name))
            f.write('%s    offset += %s\n' % (header, 4))
            f.write('%s    for i in range(0, %s_length):\n' % (header, self.name))
            c.deserialize(f, header + "        ")
            f.write('%sexcept struct.error as ex:\n' % header)
            f.write('%s    print(\'Unable to deserialize messages: ' % (header))
            f.write('%s\'%str(ex))\n')
        else:
            f.write('%stry:\n' % header)
            f.write('%s    self.%s = [%s for x in range(0, %s)]\n' % (header, self.name, primitive_default_value(self.type), self.size))
            f.write('%s    for i in range(0, %s):\n' % (header, self.size))
            c.deserialize(f, header + "        ")
            f.write('%sexcept struct.error as ex:\n' % header)
            f.write('%s    print(\'Unable to deserialize messages: ' % (header))
            f.write('%s\'%str(ex))\n')

    def serializedLength(self, f, header):
        c = self.cls(self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            f.write('%s%s_length = len(self.%s)\n' % (header, self.name, self.name))
            f.write('%slength += 4\n' % header)
            f.write('%sfor i in range(0, %s_length):\n' % (header, self.name))
            c.serializedLength(f, header + "    ")
        else:
            f.write('%sfor i in range(0, %s):\n' % (header, self.size))
            c.serializedLength(f, header + "    ")

    def echo(self, f, header, trailer):
        c = self.cls(self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            f.write('%sstring_echo += \'"%s": [\'\n' % (header, self.name))
            f.write('%s%s_length = len(self.%s)\n' % (header, self.name, self.name))
            f.write('%sfor i in range(0, %s_length):\n' % (header, self.name))
            f.write('%s    if i == (%s_length - 1): \n' % (header, self.name))
            c.echo(f, header + '          ', '')
            f.write('%s    else:\n' % (header))
            c.echo(f, header + '        ', ', ')
            f.write('%sstring_echo += \']%s\'\n' % (header, trailer))
        else:
            f.write('%sstring_echo += \'"%s": [\'\n' % (header, self.name))
            f.write('%sfor i in range(0, %s):\n' % (header, self.size))
            f.write('%s    if i == (%s - 1): \n' % (header, self.size))
            c.echo(f, header + '        ', '')
            f.write('%s    else:\n' % (header))
            c.echo(f, header + '        ', ', ')
            f.write('%sstring_echo += \']%s\'\n' % (header, trailer))

ROS_TO_EMBEDDED_TYPES = {
    'bool'    :   ('bool',            1, PrimitiveDataType, []),
    'byte'    :   ('byte',            1, PrimitiveDataType, []),
    'int8'    :   ('int8',            1, PrimitiveDataType, []),
    'char'    :   ('char',            1, PrimitiveDataType, []),
    'uint8'   :   ('uint8',           1, PrimitiveDataType, []),
    'int16'   :   ('int16',           2, PrimitiveDataType, []),
    'uint16'  :   ('uint16',          2, PrimitiveDataType, []),
    'int32'   :   ('int32',           4, PrimitiveDataType, []),
    'uint32'  :   ('uint32',          4, PrimitiveDataType, []),
    'int64'   :   ('int64',           8, PrimitiveDataType, []),
    'uint64'  :   ('uint64',          8, PrimitiveDataType, []),
    'float32' :   ('float32',         4, PrimitiveDataType, []),
    'float64' :   ('float64',         8, PrimitiveDataType, []),
    'time'    :   ('tinyros.Time',         8, TimeDataType, ['tinyros']),
    'duration':   ('tinyros.Duration',     8, TimeDataType, ['tinyros']),
    'string'  :   ('string',           0, StringDataType, []),
    'Header'  :   ('roslib_msgs.msg.Header',   0, MessageDataType, ['roslib_msgs.msg'])
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
        
        self.includes.append(self.package + ".msg")

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
                if type_package + ".msg" not in self.includes:
                    self.includes.append(type_package + ".msg")
                cls = MessageDataType
                code_type = type_package + ".msg." + type_name
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
        f.write('    def serialize(self, buff):\n')
        f.write('        offset = 0\n')
        self._write_id_serializer(f)
        for d in self.data:
            d.serialize(f, "        ")
        f.write('        return offset\n');
        f.write('\n')

    def _write_deserializer(self, f):
        # deserializer
        f.write('    def deserialize(self, buff):\n')
        f.write('        offset = 0\n')
        self._write_id_deserializer(f)
        for d in self.data:
            d.deserialize(f, "        ")
        f.write('        return offset\n');
        f.write('\n')

    def _write_serializedLength(self, f):
        # serializedLength
        f.write('    def serializedLength(self):\n')
        f.write('        length = 0\n')
        for d in self.data:
            d.serializedLength(f, "        ")
        f.write('        return length\n');
        f.write('\n')

    def _write_echo(self, f):
        f.write('    def echo(self):\n')
        f.write('        string_echo = \'{\'\n');
        if self.data:
            for d in self.data[:-1]:
                d.echo(f, '        ', ', ')
            self.data[-1].echo(f, '        ', '')
        f.write('        string_echo += \'}\'\n');
        f.write('        return string_echo\n');
        f.write('\n')

    def _write_std_includes(self, f):
        f.write('import sys\n')
        f.write('python3 = True if sys.hexversion > 0x03000000 else False\n')
        f.write('import struct\n')
        f.write('import tinyros\n')

    def _write_msg_includes(self,f):
        for i in self.includes:
            f.write('import %s\n' % i)

    def _write_constructor(self, f):
        f.write('    def __init__(self):\n')
        f.write('        super(%s, self).__init__()\n' % self.name)
        if self.data:
            for d in self.data[:-1]:
                d.make_initializer(f, "        ")
            self.data[-1].make_initializer(f, "        ")
        self._write_id_initializer(f)
        f.write('\n')

    def _write_data(self, f):
        if self.data:
            f.write('    __slots__ = [')
            for d in self.data[:-1]:
                d.make_declaration(f, ",")
            self.data[-1].make_declaration(f, "")
            f.write(']\n')
            f.write('    _slot_types = [')
            for d in self.data[:-1]:
                d.make_declaration_type(f, ",")
            self.data[-1].make_declaration_type(f, "")
            f.write(']\n\n')
        if self.enums:
            for e in self.enums:
                e.make_declaration(f, "    ")
            f.write('\n')

    def _write_getType(self, f):
        f.write('    def getType(self):\n')
        f.write('        return "%s/%s"\n'%(self.package, self.name))
        f.write('\n')

    def _write_getMD5(self, f):
        f.write('    def getMD5(self):\n')
        f.write('        return "%s"\n'%self.md5)

    def _write_impl(self, f):
        f.write('class %s(tinyros.Message):\n' % self.name)
        self._write_id(f)
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

    def make_header(self, f):
        self._write_std_includes(f)
        self._write_msg_includes(f)
        f.write('\n')
        self._write_impl(f)
        f.write('\n')
        f.write('_struct_I = struct.Struct(\'<I\')')
        f.write('\n\n')

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
        cls = PrimitiveDataType
        self.req.data.insert(0, cls("__id__", 'uint32', 4) )
        self.resp.data.insert(0, cls("__id__", 'uint32', 4) )
        
        self.req._write_std_includes(f)
        includes = self.req.includes
        includes.extend(self.resp.includes)
        includes = list(set(includes))
        for inc in includes:
            f.write('import %s\n' % inc)

        def write_getID(out):
            out.write('\n    def getID(self):\n        return self.__id__\n\n')
        _write_getID = lambda out: write_getID(out)
        self.req._write_getID = _write_getID
        self.resp._write_getID = _write_getID

        def write_setID(out):
            out.write('    def setID(self, id):\n        self.__id__ = id\n\n')
        _write_setID = lambda out: write_setID(out)
        self.req._write_setID = _write_setID
        self.resp._write_setID = _write_setID

        def write_type(out, name):
            out.write('    def getType(self):\n        return "%s/%s"\n\n'%(self.package, self.name))
        _write_getType = lambda out: write_type(out, self.name.upper())
        self.req._write_getType = _write_getType
        self.resp._write_getType = _write_getType

        f.write('\n')
        self.req._write_impl(f)
        self.resp._write_impl(f)
        f.write('class %s(object):\n' % self.name )
        f.write('    Request = %s\n' % self.req.name )
        f.write('    Response = %s\n' % self.resp.name )
        f.write('\n')
        f.write('_struct_I = struct.Struct(\'<I\')')
        f.write('\n\n')

#####################################################################
# Make a Library

def MakeLibrary(pkg_dir, package, output_path, build_dir):
    # find the messages in this package
    messages = list()
    services = list()
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
        if services == list():
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
                    services.append( Service(f[0:-4], package, srv_definition) )
                else:
                    tmp_srv_definition = open(tmp_srv_file).readlines()
                    tmp_srv_md5 = hashlib_md5sum(f[0:-4], package, tmp_srv_definition)
                    srv_md5 = hashlib_md5sum(f[0:-4], package, srv_definition)
                    if srv_md5 != tmp_srv_md5:
                       print('%s,'%f[0:-4], end='')
                       services.append( Service(f[0:-4], package, srv_definition) )
        print('\n')
    elif services != list():
        print('\n')

    # generate for each message
    messages_output_path = output_path + "/" + package + "/msg"
    if not os.path.exists(messages_output_path):
        os.makedirs(messages_output_path)
    for msg in messages:
        msg_header = open(messages_output_path + "/" + msg.name + ".py", "w")
        msg.make_header(msg_header)
        msg_header.close()

    msg_import = open(messages_output_path + "/__init__.py", "w")
    for f in os.listdir(messages_output_path):
        suffix = f[f.rfind('.'):]
        if f != "__init__.py" and suffix == ".py":
            msg_import.write('from .%s import *\n' % f[0:-3])
    msg_import.close()
  
    # generate for each service
    services_output_path = output_path + "/" + package + "/srv"
    if not os.path.exists(services_output_path):
        os.makedirs(services_output_path)
    for srv in services:
        srv_header = open(services_output_path + "/" + srv.name + ".py", "w")
        srv.make_header(srv_header)
        srv_header.close()

    srv_import = open(services_output_path + "/__init__.py", "w")
    for f in os.listdir(services_output_path):
        suffix = f[f.rfind('.'):]
        if f != "__init__.py" and suffix == ".py":
            srv_import.write('from .%s import *\n' % f[0:-3])
    srv_import.close()
    
    package_import = open(output_path + "/" + package + "/__init__.py", "w")
    package_import.write('\n')
    package_import.close()

def messages_generate(path):
    # gimme messages
    failed = []
    mydir = sys.argv[1] + "/msgs"
    builddir = sys.argv[1] + "/build/CMake/python_msgs"
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
    files = ['tinyros/__init__.py',
             'tinyros/Duration.py',
             'tinyros/Hardware.py',
             'tinyros/HardwareUdp.py',
             'tinyros/Message.py',
             'tinyros/NodeHandle.py',
             'tinyros/Publisher.py',
             'tinyros/ServiceClient.py',
             'tinyros/ServiceServer.py',
             'tinyros/Subscriber.py',
             'tinyros/ThreadPool.py',
             'tinyros/Time.py']

    mydir = sys.argv[3] + "/roslib/python/"
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
    files = ['examples/publisher/ExamplePublisher.py',
             'examples/subscriber/ExampleSubscriber.py',
             'examples/service/ExampleService.py',
             'examples/service_client/ExampleServiceClient.py']

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

roslib_copy_roslib_files(path+"/")
roslib_copy_examples_files(path+"/")
messages_generate(path+"/")
if os.path.exists(sys.argv[1] + "/build/CMake/python_msgs"):
    shutil.rmtree(sys.argv[1] + "/build/CMake/python_msgs")
shutil.copytree(sys.argv[1] + "/msgs", sys.argv[1] + "/build/CMake/python_msgs")

