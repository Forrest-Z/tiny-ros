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
                     'char', 'uint8', 'uint16', 'uint32', 'uint64', 'float64', 'string']:
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
    elif field_type == 'char':
        return "' '"
    elif field_type == 'string':
        return '""'
    elif field_type.endswith(']'): # array type
        base_type, is_array, array_len = parse_type(field_type)
        if array_len is None: #var-length
            return 'new Array()'
        else: # fixed-length, fill values
            def_val = primitive_default_value(base_type)
            return '[' + ','.join(itertools.repeat(def_val, array_len)) + ']'
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

    def make_initializer(self, f, header):
        f.write('%sthis.%s = %s;\n' % (header, self.name, self.value))

class PrimitiveDataType:
    """ Our datatype is a C/C++ primitive. """

    def __init__(self, name, ty, bytes):
        self.name = name
        self.type = ty
        self.bytes = bytes
        
    def make_initializer(self, f, header):
        f.write('%sthis.%s = %s;\n' % (header, self.name, primitive_default_value(self.type)))

    def serialize(self, f, header):
        cn = self.name.replace("[","").replace("]","").split(".")[-1]
        if self.type == 'char':
            f.write('%sbuff[offset] = this.%s.charCodeAt() & 0xFF;\n' % (header, self.name))
        elif self.type == 'bool':
            f.write('%sbuff[offset] = this.%s === false ? 0 : 1;\n' % (header, self.name))
        elif self.type == 'float32':
            f.write('%svar float32Array_%s = new Float32Array(1);\n' % (header, cn))
            f.write('%svar uInt8Float32Array_%s = new Uint8Array(float32Array_%s.buffer);\n' % (header, cn, cn))
            f.write('%sfloat32Array_%s[0] = +this.%s;\n' % (header, cn, self.name))
            for i in range(self.bytes):
                f.write('%sbuff[offset + %d] = uInt8Float32Array_%s[%s];\n' % (header, i, cn, i))
        elif self.type == 'float64':
            f.write('%svar float64Array_%s = new Float64Array(1);\n' % (header, cn))
            f.write('%svar uInt8Float64Array_%s = new Uint8Array(float64Array_%s.buffer);\n' % (header, cn, cn))
            f.write('%sfloat64Array_%s[0] = +this.%s;\n' % (header, cn, self.name))
            for i in range(self.bytes):
                f.write('%sbuff[offset + %d] = uInt8Float64Array_%s[%s];\n' % (header, i, cn, i))
        else:
            for i in range(self.bytes):
                f.write('%sbuff[offset + %d] = ((+this.%s) >> (8 * %d)) & 0xFF;\n' % (header, i, self.name, i))
        f.write('%soffset += %s;\n' % (header, self.bytes) )
                
    def deserialize(self, f, header):
        cn = self.name.replace("[","").replace("]","").split(".")[-1]
        if self.type == 'char':
            f.write('%sthis.%s = String.fromCharCode(buff[offset]);\n' % (header, self.name))
        elif self.type == 'bool':
            f.write('%sthis.%s = buff[offset] !== 0 ? true : false;\n' % (header, self.name))
        elif self.type == 'float32':
            f.write('%svar float32Array_%s = new Float32Array(1);\n' % (header, cn))
            f.write('%svar uInt8Float32Array_%s = new Uint8Array(float32Array_%s.buffer);\n' % (header, cn, cn))
            for i in range(self.bytes):
                f.write('%suInt8Float32Array_%s[%s] = buff[offset + %d];\n' % (header, cn, i, i))
            f.write('%sthis.%s = float32Array_%s[0];\n' % (header, self.name, cn))
        elif self.type == 'float64':
            f.write('%svar float64Array_%s = new Float64Array(1);\n' % (header, cn))
            f.write('%svar uInt8Float64Array_%s = new Uint8Array(float64Array_%s.buffer);\n' % (header, cn, cn))
            for i in range(self.bytes):
                f.write('%suInt8Float64Array_%s[%s] = buff[offset + %d];\n' % (header, cn, i, i))
            f.write('%sthis.%s = float64Array_%s[0];\n' % (header, self.name, cn))
        else:
            f.write('%sthis.%s = +((buff[offset + 0] & 0xFF) << (8 * 0));\n' % (header, self.name))
            for i in range(self.bytes-1):
                f.write('%sthis.%s |= +((buff[offset + %d] & 0xFF) << (8 * %d));\n' % (header, self.name, i+1, i+1) )
        f.write('%soffset += %s;\n' % (header, self.bytes) )

    def serializedLength(self, f, header):
        f.write('%slength += %s\n' % (header, self.bytes))

    def echo(self, f, header, trailer):
        pass

class MessageDataType(PrimitiveDataType):
    """ For when our data type is another message. """

    def make_initializer(self, f, header):
        f.write('%sthis.%s = %s();\n' % (header, self.name, self.type))

    def serialize(self, f, header):
        f.write('%soffset += this.%s.serialize(buff, offset);\n' % (header, self.name))

    def deserialize(self, f, header):
        f.write('%soffset += this.%s.deserialize(buff, offset);\n' % (header, self.name))

    def serializedLength(self, f, header):
        f.write('%slength += this.%s.serializedLength();\n' % (header, self.name))

    def echo(self, f, header, trailer):
        pass

class StringDataType(PrimitiveDataType):
    def make_initializer(self, f, header):
        f.write('%sthis.%s = "";\n' % (header, self.name))

    def serialize(self, f, header):
        cn = self.name.replace("[","").replace("]","")
        f.write('%svar encoder_%s = new TextEncoder(\'utf8\');\n' % (header, cn))
        f.write('%svar utf8array_%s = encoder_%s.encode(this.%s);\n' % (header, cn, cn, self.name))
        f.write('%sbuff[offset + 0] = (utf8array_%s.length >> (8 * 0)) & 0xFF;\n' % (header, cn))
        f.write('%sbuff[offset + 1] = (utf8array_%s.length >> (8 * 1)) & 0xFF;\n' % (header, cn))
        f.write('%sbuff[offset + 2] = (utf8array_%s.length >> (8 * 2)) & 0xFF;\n' % (header, cn))
        f.write('%sbuff[offset + 3] = (utf8array_%s.length >> (8 * 3)) & 0xFF;\n' % (header, cn))
        f.write('%soffset += 4;\n' % header)
        f.write('%sfor (var i = 0; i < utf8array_%s.length; i++) {\n' % (header, cn))
        f.write('%s    buff[offset + i] = utf8array_%s[i];\n' % (header, cn))
        f.write('%s}\n' % header)
        f.write('%soffset += utf8array_%s.length;\n' % (header, cn))
 
    def deserialize(self, f, header):
        cn = self.name.replace("[","").replace("]","")
        f.write('%svar length_%s = +((buff[offset + 0] & 0xFF) << (8 * 0));\n' % (header, cn))
        f.write('%slength_%s |= +((buff[offset + 1] & 0xFF) << (8 * 1));\n' % (header, cn))
        f.write('%slength_%s |= +((buff[offset + 2] & 0xFF) << (8 * 2));\n' % (header, cn))
        f.write('%slength_%s |= +((buff[offset + 3] & 0xFF) << (8 * 3))\n' % (header, cn))
        f.write('%soffset += 4;\n' % header)
        f.write('%svar decoder_%s = new TextDecoder(\'utf8\');\n' % (header, cn))
        f.write('%sthis.%s = decoder_%s.decode(buff.slice(offset, offset + length_%s));\n' % (header, self.name, cn, cn))
        f.write('%soffset += length_%s;\n' % (header, cn))

    def serializedLength(self, f, header):
        cn = self.name.replace("[","").replace("]","")
        f.write('%svar encoder_%s = new TextEncoder(\'utf8\');\n' % (header, cn))
        f.write('%svar utf8array_%s = encoder_%s.encode(this.%s);\n' % (header, cn, cn, self.name))
        f.write('%slength += 4;\n' % header)
        f.write('%slength += utf8array_%s.length;\n' % (header, cn))

    def echo(self, f, header, trailer):
        pass

class TimeDataType(PrimitiveDataType):

    def __init__(self, name, ty, bytes):
        self.name = name
        self.type = ty
        self.sec = PrimitiveDataType(name+'.sec','uint32',4)
        self.nsec = PrimitiveDataType(name+'.nsec','uint32',4)

    def make_initializer(self, f, header):
        f.write('%sthis.%s = %s();\n' % (header, self.name, self.type))

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

    def __init__(self, name, ty, bytes, cls, array_size=None):
        self.name = name
        self.type = ty
        self.bytes = bytes
        self.size = array_size
        self.cls = cls

    def make_initializer(self, f, header):
        if self.size == None:
            f.write('%sthis.%s = %s;\n' % (header, self.name, primitive_default_value(self.type + "[]")))
        else:
            f.write('%sthis.%s = %s;\n' % (header, self.name, primitive_default_value(self.type + "[" + str(self.size) + "]")))

    def serialize(self, f, header):
        c = self.cls(self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            f.write('%svar length_%s = this.%s.length;\n' % (header, self.name, self.name))
            f.write('%sbuff[offset + 0] = (length_%s >> (8 * 0)) & 0xFF;\n' % (header, self.name))
            f.write('%sbuff[offset + 1] = (length_%s >> (8 * 1)) & 0xFF;\n' % (header, self.name))
            f.write('%sbuff[offset + 2] = (length_%s >> (8 * 2)) & 0xFF;\n' % (header, self.name))
            f.write('%sbuff[offset + 3] = (length_%s >> (8 * 3)) & 0xFF;\n' % (header, self.name))
            f.write('%soffset += 4;\n' % header)
            f.write('%sfor (var i = 0; i < length_%s; i++) {\n' % (header, self.name))
            c.serialize(f, header + '    ')
            f.write('%s}\n' % header)
        else:
            f.write('%sfor (var i = 0; i < %d; i++) {\n' % (header, self.size) )
            c.serialize(f ,header + '    ')
            f.write('%s}\n' % header)

    def deserialize(self, f, header):
        c = self.cls(self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            # deserialize length
            f.write('%svar length_%s = +((buff[offset + 0] & 0xFF) << (8 * 0));\n' % (header, self.name))
            f.write('%slength_%s |= +((buff[offset + 1] & 0xFF) << (8 * 1));\n' % (header, self.name))
            f.write('%slength_%s |= +((buff[offset + 2] & 0xFF) << (8 * 2));\n' % (header, self.name))
            f.write('%slength_%s |= +((buff[offset + 3] & 0xFF) << (8 * 3));\n' % (header, self.name))
            f.write('%soffset += 4;\n' % header)
            f.write('%sthis.%s = new Array(length_%s);\n' % (header, self.name, self.name))
            f.write('%sfor (var i = 0; i < length_%s; i++) {\n' % (header, self.name))
            if not primitive_type(self.type):
                f.write('%s    this.%s[i] = %s;\n' % (header, self.name, primitive_default_value(self.type)))
            c.deserialize(f, header + '    ')
            f.write('%s}\n' % header)
        else:
            f.write('%sfor (var i = 0; i < %d; i++) {\n' % (header, self.size) )
            c.deserialize(f, header + '    ')
            f.write('%s}\n' % header)

    def serializedLength(self, f, header):
        c = self.cls(self.name+"[i]", self.type, self.bytes)
        if self.size == None:
            f.write('%svar length_%s = this.%s.length;\n' % (header, self.name, self.name))
            f.write('%slength += 4;\n' % header)
            f.write('%sfor (var i = 0; i < length_%s; i++) {\n' % (header, self.name))
            c.serializedLength(f, header + '    ')
            f.write('%s}\n' % header)
        else:
            f.write('%sfor (var i = 0; i < %d; i++) {\n' % (header, self.size) )
            c.serializedLength(f ,header + '    ')
            f.write('%s}\n' % header)

    def echo(self, f, header, trailer):
        pass

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
    'time'    :   ('ros.Time',         8, TimeDataType, ['ros/Time']),
    'duration':   ('ros.Duration',     8, TimeDataType, ['ros/Duration']),
    'string'  :   ('string',           0, StringDataType, []),
    'Header'  :   ('std_msgs.Header',   0, MessageDataType, ['std_msgs/Header'])
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
                code_type = type_package + "." + type_name
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
        f.write('%s.prototype.getID = function() { return 0; };\n\n' % self.name)

    def _write_setID(self, f):
        f.write('%s.prototype.setID = function(id) { };\n\n' % self.name)

    def _write_serializer(self, f):
        # serializer
        f.write('%s.prototype.serialize = function(buff, idx) {\n' % self.name)
        f.write('    var offset = idx;\n')
        self._write_id_serializer(f)
        for d in self.data:
            d.serialize(f, "    ")
        f.write('    return offset;\n');
        f.write('};\n\n')

    def _write_deserializer(self, f):
        # deserializer
        f.write('%s.prototype.deserialize = function(buff, idx) {\n' % self.name)
        f.write('    var offset = idx;\n')
        self._write_id_deserializer(f)
        for d in self.data:
            d.deserialize(f, "    ")
        f.write('    return offset;\n');
        f.write('};\n\n')

    def _write_serializedLength(self, f):
        # serializedLength
        f.write('%s.prototype.serializedLength = function() {\n' % self.name)
        f.write('    var length = 0;\n')
        for d in self.data:
            d.serializedLength(f, "    ")
        f.write('    return length;\n');
        f.write('};\n\n')

    def _write_echo(self, f):
        f.write('%s.prototype.echo = function() { return ""; };\n\n' % self.name)

    def _write_constructor(self, f):
        f.write('function %s() {\n' % self.name)
        if self.data:
            for d in self.data:
                d.make_initializer(f, "    ")
        self._write_id_initializer(f)
        if self.enums:
            f.write('\n    // ENUM{\n')
            for e in self.enums:
                e.make_initializer(f, "    ")
            f.write('    // }ENUM\n')
        f.write('};\n\n')

    def _write_data(self, f):
        pass

    def _write_getType(self, f):
        f.write('%s.prototype.getType = function() { return "%s/%s"; };\n\n' % (self.name, self.package, self.name))

    def _write_getMD5(self, f):
        f.write('%s.prototype.getMD5 = function() { return "%s"; };\n\n' % (self.name, self.md5))
    def _write_impl(self, f):
        self._write_constructor(f)
        self._write_id(f)
        self._write_data(f)
        self._write_serializer(f)
        self._write_deserializer(f)
        self._write_serializedLength(f)
        self._write_echo(f)
        self._write_getType(f)
        self._write_getMD5(f)
        self._write_getID(f)
        self._write_setID(f)

    def make_header(self, f):
        
        f.write('(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}')
        f.write('else if(typeof define==="function"&&define.amd){define([],f);}')
        f.write('else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}')
        f.write('if(typeof g.%s==="undefined"){g.%s={};}g.%s.%s=f();}})(function(){' %(self.package, self.package, self.package, self.name))
        f.write('var define,module,exports;\'use strict\';\n')
        if self.includes:
            f.write('const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));\n')
            f.write('const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));\n')
        for inc in self.includes:
            f.write('document.write("<script language=javascript src=\'"+_ROOT_PATH_+"/%s.js\'></script>");\n' % inc)
        f.write('\n')
        self._write_impl(f)
        f.write('return function() { return new %s(); };\n\n' % self.name)
        f.write('});\n')

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

        def write_getIDRequest(out):
            out.write('%sRequest.prototype.getID = function() { return this.__id__; };\n\n' % self.name)
        def write_getIDResponse(out):
            out.write('%sResponse.prototype.getID = function() { return this.__id__; };\n\n' % self.name)
        _write_getIDRequest = lambda out: write_getIDRequest(out)
        _write_getIDResponse = lambda out: write_getIDResponse(out)
        self.req._write_getID = _write_getIDRequest
        self.resp._write_getID = _write_getIDResponse

        def write_setIDRequest(out):
            out.write('%sRequest.prototype.setID = function(id) { this.__id__ = id; };\n\n' % self.name)
        def write_setIDResponse(out):
            out.write('%sResponse.prototype.setID = function(id) { this.__id__ = id; };\n\n' % self.name)
        _write_setIDRequest = lambda out: write_setIDRequest(out)
        _write_setIDResponse = lambda out: write_setIDResponse(out)
        self.req._write_setID = _write_setIDRequest
        self.resp._write_setID = _write_setIDResponse

        def write_typeRequest(out, name):
            out.write('%sRequest.prototype.getType = function() { return "%s/%s"; };\n\n' % (self.name, self.package, self.name))
        def write_typeResponse(out, name):
            out.write('%sResponse.prototype.getType = function() { return "%s/%s"; };\n\n' % (self.name, self.package, self.name))
        _write_getTypeRequest = lambda out: write_typeRequest(out, self.name.upper())
        _write_getTypeResponse = lambda out: write_typeResponse(out, self.name.upper())
        self.req._write_getType = _write_getTypeRequest
        self.resp._write_getType = _write_getTypeResponse
        
        f.write('(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}')
        f.write('else if(typeof define==="function"&&define.amd){define([],f);}')
        f.write('else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}')
        f.write('if(typeof g.%s==="undefined"){g.%s={};}g.%s.%sRequest=f();}})(function(){' %(self.package, self.package, self.package, self.name))
        f.write('var define,module,exports;\'use strict\';\n')
        if self.req.includes:
            f.write('const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));\n')
            f.write('const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));\n')
        for inc in self.req.includes:
            f.write('document.write("<script language=javascript src=\'"+_ROOT_PATH_+"/%s.js\'></script>");\n' % inc)
        f.write('\n')
        self.req._write_impl(f)
        f.write('return function() { return new %sRequest(); };\n\n' % self.name)
        f.write('});\n\n\n')
        
        for i in range(100):
            f.write('/')
        f.write('\n\n\n')
        
        f.write('(function(f){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=f();}')
        f.write('else if(typeof define==="function"&&define.amd){define([],f);}')
        f.write('else{var g;if(typeof window!=="undefined"){g=window;}else if(typeof global!=="undefined"){g=global;}else if(typeof self!=="undefined"){g=self;}else{g=this;}')
        f.write('if(typeof g.%s==="undefined"){g.%s={};}g.%s.%sResponse=f();}})(function(){' %(self.package, self.package, self.package, self.name))
        f.write('var define,module,exports;\'use strict\';\n')
        if self.resp.includes:
            f.write('const _CURRENT_PATH_ = document.currentScript.src.substring(0, document.currentScript.src.lastIndexOf("/"));\n')
            f.write('const _ROOT_PATH_ = _CURRENT_PATH_.substring(0, _CURRENT_PATH_.lastIndexOf("/"));\n')
        for inc in self.resp.includes:
            f.write('document.write("<script language=javascript src=\'"+_ROOT_PATH_+"/%s.js\'></script>");\n' % inc)
        f.write('\n')
        self.resp._write_impl(f)
        f.write('return function() { return new %sResponse(); };\n\n' % self.name)
        f.write('});\n')

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
    messages_output_path = output_path + "/" + package
    if not os.path.exists(messages_output_path):
        os.makedirs(messages_output_path)
    for msg in messages:
        msg_header = open(messages_output_path + "/" + msg.name + ".js", "w")
        msg.make_header(msg_header)
        msg_header.close()
  
    # generate for each service
    services_output_path = output_path + "/" + package
    if not os.path.exists(services_output_path):
        os.makedirs(services_output_path)
    for srv in services:
        srv_header = open(services_output_path + "/" + srv.name + ".js", "w")
        srv.make_header(srv_header)
        srv_header.close()

def messages_generate(path):
    # gimme messages
    failed = []
    mydir = sys.argv[1] + "/msgs"
    builddir = sys.argv[1] + "/build/CMake/javascript_msgs"
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
    if not os.path.exists(path+"ros"):
        os.makedirs(path+"ros")
    files = ['ros/Time.js',
             'ros/Duration.js',
             'ros/NodeHandle.js',
             'ros/Publisher.js',
             'ros/ros.js',
             'ros/Subscriber.js']

    mydir = sys.argv[3] + "/roslib/javascript/"
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
messages_generate(path+"/")
if os.path.exists(sys.argv[1] + "/build/CMake/javascript_msgs"):
    shutil.rmtree(sys.argv[1] + "/build/CMake/javascript_msgs")
shutil.copytree(sys.argv[1] + "/msgs", sys.argv[1] + "/build/CMake/javascript_msgs")

