import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import nav_msgs.msg
import std_msgs.msg
import geometry_msgs.msg

class GridCells(tinyros.Message):
    __slots__ = ['header','cell_width','cell_height','cells']
    _slot_types = ['std_msgs.msg.Header','float32','float32','geometry_msgs.msg.Point[]']

    def __init__(self):
        super(GridCells, self).__init__()
        self.header = std_msgs.msg.Header()
        self.cell_width = 0.
        self.cell_height = 0.
        self.cells = []

    def serialize(self, buff):
        offset = 0
        offset += self.header.serialize(buff)
        try:
            struct_cell_width = struct.Struct("<f")
            buff.write(struct_cell_width.pack(self.cell_width))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_cell_height = struct.Struct("<f")
            buff.write(struct_cell_height.pack(self.cell_height))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            cells_length = len(self.cells)
            buff.write(_struct_I.pack(cells_length))
            offset += 4
            for i in range(0, cells_length):
                offset += self.cells[i].serialize(buff)
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        offset += self.header.deserialize(buff[offset:])
        try:
            struct_cell_width = struct.Struct("<f")
            (self.cell_width,) = struct_cell_width.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_cell_height = struct.Struct("<f")
            (self.cell_height,) = struct_cell_height.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (cells_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.cells = [geometry_msgs.msg.Point() for x in range(0, cells_length)]
            offset += 4
            for i in range(0, cells_length):
                offset += self.cells[i].deserialize(buff[offset:])
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += self.header.serializedLength()
        length += 4
        length += 4
        cells_length = len(self.cells)
        length += 4
        for i in range(0, cells_length):
            length += self.cells[i].serializedLength()
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"header": {"'
        string_echo += self.header.echo()
        string_echo += '}, '
        string_echo += '"cell_width": %s'%cell_width
        string_echo += ', '
        string_echo += '"cell_height": %s'%cell_height
        string_echo += ', '
        string_echo += '"cells": ['
        cells_length = len(self.cells)
        for i in range(0, cells_length):
            if i == (cells_length - 1): 
                string_echo += '{"cells%s": {'%i
                string_echo += self.cells[i].echo()
                string_echo += '}}'
            else:
                string_echo += '{"cells%s": {'%i
                string_echo += self.cells[i].echo()
                string_echo += '}}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "nav_msgs/GridCells"

    def getMD5(self):
        return "13ce9063aaf922c39d3a2207d3926427"

_struct_I = struct.Struct('<I')

