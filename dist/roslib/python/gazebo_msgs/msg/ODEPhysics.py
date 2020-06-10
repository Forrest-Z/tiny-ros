import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg

class ODEPhysics(tinyros.Message):
    __slots__ = ['auto_disable_bodies','sor_pgs_precon_iters','sor_pgs_iters','sor_pgs_w','sor_pgs_rms_error_tol','contact_surface_layer','contact_max_correcting_vel','cfm','erp','max_contacts']
    _slot_types = ['bool','uint32','uint32','float64','float64','float64','float64','float64','float64','uint32']

    def __init__(self):
        super(ODEPhysics, self).__init__()
        self.auto_disable_bodies = False
        self.sor_pgs_precon_iters = 0
        self.sor_pgs_iters = 0
        self.sor_pgs_w = 0.
        self.sor_pgs_rms_error_tol = 0.
        self.contact_surface_layer = 0.
        self.contact_max_correcting_vel = 0.
        self.cfm = 0.
        self.erp = 0.
        self.max_contacts = 0

    def serialize(self, buff):
        offset = 0
        try:
            struct_auto_disable_bodies = struct.Struct("<B")
            buff.write(struct_auto_disable_bodies.pack(self.auto_disable_bodies))
            offset += 1
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.sor_pgs_precon_iters))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.sor_pgs_iters))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_sor_pgs_w = struct.Struct("<d")
            buff.write(struct_sor_pgs_w.pack(self.sor_pgs_w))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_sor_pgs_rms_error_tol = struct.Struct("<d")
            buff.write(struct_sor_pgs_rms_error_tol.pack(self.sor_pgs_rms_error_tol))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_contact_surface_layer = struct.Struct("<d")
            buff.write(struct_contact_surface_layer.pack(self.contact_surface_layer))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_contact_max_correcting_vel = struct.Struct("<d")
            buff.write(struct_contact_max_correcting_vel.pack(self.contact_max_correcting_vel))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_cfm = struct.Struct("<d")
            buff.write(struct_cfm.pack(self.cfm))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            struct_erp = struct.Struct("<d")
            buff.write(struct_erp.pack(self.erp))
            offset += 8
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            buff.write(_struct_I.pack(self.max_contacts))
            offset += 4
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            struct_auto_disable_bodies = struct.Struct("<B")
            (self.auto_disable_bodies,) = struct_auto_disable_bodies.unpack(buff[offset:(offset + 1)])
            self.auto_disable_bodies = bool(self.auto_disable_bodies)
            offset += 1
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.sor_pgs_precon_iters,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.sor_pgs_iters,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_sor_pgs_w = struct.Struct("<d")
            (self.sor_pgs_w,) = struct_sor_pgs_w.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_sor_pgs_rms_error_tol = struct.Struct("<d")
            (self.sor_pgs_rms_error_tol,) = struct_sor_pgs_rms_error_tol.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_contact_surface_layer = struct.Struct("<d")
            (self.contact_surface_layer,) = struct_contact_surface_layer.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_contact_max_correcting_vel = struct.Struct("<d")
            (self.contact_max_correcting_vel,) = struct_contact_max_correcting_vel.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_cfm = struct.Struct("<d")
            (self.cfm,) = struct_cfm.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            struct_erp = struct.Struct("<d")
            (self.erp,) = struct_erp.unpack(buff[offset:(offset + 8)])
            offset += 8
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (self.max_contacts,) = _struct_I.unpack(buff[offset:(offset + 4)])
            offset += 4
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        length += 1
        length += 4
        length += 4
        length += 8
        length += 8
        length += 8
        length += 8
        length += 8
        length += 8
        length += 4
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"auto_disable_bodies": %s'%auto_disable_bodies
        string_echo += ', '
        string_echo += '"sor_pgs_precon_iters": %s'%sor_pgs_precon_iters
        string_echo += ', '
        string_echo += '"sor_pgs_iters": %s'%sor_pgs_iters
        string_echo += ', '
        string_echo += '"sor_pgs_w": %s'%sor_pgs_w
        string_echo += ', '
        string_echo += '"sor_pgs_rms_error_tol": %s'%sor_pgs_rms_error_tol
        string_echo += ', '
        string_echo += '"contact_surface_layer": %s'%contact_surface_layer
        string_echo += ', '
        string_echo += '"contact_max_correcting_vel": %s'%contact_max_correcting_vel
        string_echo += ', '
        string_echo += '"cfm": %s'%cfm
        string_echo += ', '
        string_echo += '"erp": %s'%erp
        string_echo += ', '
        string_echo += '"max_contacts": %s'%max_contacts
        string_echo += ''
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/ODEPhysics"

    def getMD5(self):
        return "67a077e58362b50f63dc189c25d01418"

_struct_I = struct.Struct('<I')

