import sys
python3 = True if sys.hexversion > 0x03000000 else False
import struct
import tinyros
import gazebo_msgs.msg

class ODEJointProperties(tinyros.Message):
    __slots__ = ['damping','hiStop','loStop','erp','cfm','stop_erp','stop_cfm','fudge_factor','fmax','vel']
    _slot_types = ['float64[]','float64[]','float64[]','float64[]','float64[]','float64[]','float64[]','float64[]','float64[]','float64[]']

    def __init__(self):
        super(ODEJointProperties, self).__init__()
        self.damping = []
        self.hiStop = []
        self.loStop = []
        self.erp = []
        self.cfm = []
        self.stop_erp = []
        self.stop_cfm = []
        self.fudge_factor = []
        self.fmax = []
        self.vel = []

    def serialize(self, buff):
        offset = 0
        try:
            damping_length = len(self.damping)
            buff.write(_struct_I.pack(damping_length))
            offset += 4
            for i in range(0, damping_length):
                try:
                    struct_dampingi = struct.Struct("<d")
                    buff.write(struct_dampingi.pack(self.damping[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            hiStop_length = len(self.hiStop)
            buff.write(_struct_I.pack(hiStop_length))
            offset += 4
            for i in range(0, hiStop_length):
                try:
                    struct_hiStopi = struct.Struct("<d")
                    buff.write(struct_hiStopi.pack(self.hiStop[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            loStop_length = len(self.loStop)
            buff.write(_struct_I.pack(loStop_length))
            offset += 4
            for i in range(0, loStop_length):
                try:
                    struct_loStopi = struct.Struct("<d")
                    buff.write(struct_loStopi.pack(self.loStop[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            erp_length = len(self.erp)
            buff.write(_struct_I.pack(erp_length))
            offset += 4
            for i in range(0, erp_length):
                try:
                    struct_erpi = struct.Struct("<d")
                    buff.write(struct_erpi.pack(self.erp[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            cfm_length = len(self.cfm)
            buff.write(_struct_I.pack(cfm_length))
            offset += 4
            for i in range(0, cfm_length):
                try:
                    struct_cfmi = struct.Struct("<d")
                    buff.write(struct_cfmi.pack(self.cfm[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            stop_erp_length = len(self.stop_erp)
            buff.write(_struct_I.pack(stop_erp_length))
            offset += 4
            for i in range(0, stop_erp_length):
                try:
                    struct_stop_erpi = struct.Struct("<d")
                    buff.write(struct_stop_erpi.pack(self.stop_erp[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            stop_cfm_length = len(self.stop_cfm)
            buff.write(_struct_I.pack(stop_cfm_length))
            offset += 4
            for i in range(0, stop_cfm_length):
                try:
                    struct_stop_cfmi = struct.Struct("<d")
                    buff.write(struct_stop_cfmi.pack(self.stop_cfm[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            fudge_factor_length = len(self.fudge_factor)
            buff.write(_struct_I.pack(fudge_factor_length))
            offset += 4
            for i in range(0, fudge_factor_length):
                try:
                    struct_fudge_factori = struct.Struct("<d")
                    buff.write(struct_fudge_factori.pack(self.fudge_factor[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            fmax_length = len(self.fmax)
            buff.write(_struct_I.pack(fmax_length))
            offset += 4
            for i in range(0, fmax_length):
                try:
                    struct_fmaxi = struct.Struct("<d")
                    buff.write(struct_fmaxi.pack(self.fmax[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        try:
            vel_length = len(self.vel)
            buff.write(_struct_I.pack(vel_length))
            offset += 4
            for i in range(0, vel_length):
                try:
                    struct_veli = struct.Struct("<d")
                    buff.write(struct_veli.pack(self.vel[i]))
                    offset += 8
                except struct.error as ex:
                    print('Unable to serialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to serialize messages: %s'%str(ex))
        return offset

    def deserialize(self, buff):
        offset = 0
        try:
            (damping_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.damping = [0. for x in range(0, damping_length)]
            offset += 4
            for i in range(0, damping_length):
                try:
                    struct_dampingi = struct.Struct("<d")
                    (self.damping[i],) = struct_dampingi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (hiStop_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.hiStop = [0. for x in range(0, hiStop_length)]
            offset += 4
            for i in range(0, hiStop_length):
                try:
                    struct_hiStopi = struct.Struct("<d")
                    (self.hiStop[i],) = struct_hiStopi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (loStop_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.loStop = [0. for x in range(0, loStop_length)]
            offset += 4
            for i in range(0, loStop_length):
                try:
                    struct_loStopi = struct.Struct("<d")
                    (self.loStop[i],) = struct_loStopi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (erp_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.erp = [0. for x in range(0, erp_length)]
            offset += 4
            for i in range(0, erp_length):
                try:
                    struct_erpi = struct.Struct("<d")
                    (self.erp[i],) = struct_erpi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (cfm_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.cfm = [0. for x in range(0, cfm_length)]
            offset += 4
            for i in range(0, cfm_length):
                try:
                    struct_cfmi = struct.Struct("<d")
                    (self.cfm[i],) = struct_cfmi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (stop_erp_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.stop_erp = [0. for x in range(0, stop_erp_length)]
            offset += 4
            for i in range(0, stop_erp_length):
                try:
                    struct_stop_erpi = struct.Struct("<d")
                    (self.stop_erp[i],) = struct_stop_erpi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (stop_cfm_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.stop_cfm = [0. for x in range(0, stop_cfm_length)]
            offset += 4
            for i in range(0, stop_cfm_length):
                try:
                    struct_stop_cfmi = struct.Struct("<d")
                    (self.stop_cfm[i],) = struct_stop_cfmi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (fudge_factor_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.fudge_factor = [0. for x in range(0, fudge_factor_length)]
            offset += 4
            for i in range(0, fudge_factor_length):
                try:
                    struct_fudge_factori = struct.Struct("<d")
                    (self.fudge_factor[i],) = struct_fudge_factori.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (fmax_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.fmax = [0. for x in range(0, fmax_length)]
            offset += 4
            for i in range(0, fmax_length):
                try:
                    struct_fmaxi = struct.Struct("<d")
                    (self.fmax[i],) = struct_fmaxi.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        try:
            (vel_length,) = _struct_I.unpack(buff[offset:(offset + 4)])
            self.vel = [0. for x in range(0, vel_length)]
            offset += 4
            for i in range(0, vel_length):
                try:
                    struct_veli = struct.Struct("<d")
                    (self.vel[i],) = struct_veli.unpack(buff[offset:(offset + 8)])
                    offset += 8
                except struct.error as ex:
                    print('Unable to deserialize messages: %s'%str(ex))
        except struct.error as ex:
            print('Unable to deserialize messages: %s'%str(ex))
        return offset

    def serializedLength(self):
        length = 0
        damping_length = len(self.damping)
        length += 4
        for i in range(0, damping_length):
            length += 8
        hiStop_length = len(self.hiStop)
        length += 4
        for i in range(0, hiStop_length):
            length += 8
        loStop_length = len(self.loStop)
        length += 4
        for i in range(0, loStop_length):
            length += 8
        erp_length = len(self.erp)
        length += 4
        for i in range(0, erp_length):
            length += 8
        cfm_length = len(self.cfm)
        length += 4
        for i in range(0, cfm_length):
            length += 8
        stop_erp_length = len(self.stop_erp)
        length += 4
        for i in range(0, stop_erp_length):
            length += 8
        stop_cfm_length = len(self.stop_cfm)
        length += 4
        for i in range(0, stop_cfm_length):
            length += 8
        fudge_factor_length = len(self.fudge_factor)
        length += 4
        for i in range(0, fudge_factor_length):
            length += 8
        fmax_length = len(self.fmax)
        length += 4
        for i in range(0, fmax_length):
            length += 8
        vel_length = len(self.vel)
        length += 4
        for i in range(0, vel_length):
            length += 8
        return length

    def echo(self):
        string_echo = '{'
        string_echo += '"damping": ['
        damping_length = len(self.damping)
        for i in range(0, damping_length):
            if i == (damping_length - 1): 
                string_echo += '{"damping%s": %s'%(i,damping[i])
                string_echo += '}'
            else:
                string_echo += '{"damping%s": %s'%(i,damping[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"hiStop": ['
        hiStop_length = len(self.hiStop)
        for i in range(0, hiStop_length):
            if i == (hiStop_length - 1): 
                string_echo += '{"hiStop%s": %s'%(i,hiStop[i])
                string_echo += '}'
            else:
                string_echo += '{"hiStop%s": %s'%(i,hiStop[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"loStop": ['
        loStop_length = len(self.loStop)
        for i in range(0, loStop_length):
            if i == (loStop_length - 1): 
                string_echo += '{"loStop%s": %s'%(i,loStop[i])
                string_echo += '}'
            else:
                string_echo += '{"loStop%s": %s'%(i,loStop[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"erp": ['
        erp_length = len(self.erp)
        for i in range(0, erp_length):
            if i == (erp_length - 1): 
                string_echo += '{"erp%s": %s'%(i,erp[i])
                string_echo += '}'
            else:
                string_echo += '{"erp%s": %s'%(i,erp[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"cfm": ['
        cfm_length = len(self.cfm)
        for i in range(0, cfm_length):
            if i == (cfm_length - 1): 
                string_echo += '{"cfm%s": %s'%(i,cfm[i])
                string_echo += '}'
            else:
                string_echo += '{"cfm%s": %s'%(i,cfm[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"stop_erp": ['
        stop_erp_length = len(self.stop_erp)
        for i in range(0, stop_erp_length):
            if i == (stop_erp_length - 1): 
                string_echo += '{"stop_erp%s": %s'%(i,stop_erp[i])
                string_echo += '}'
            else:
                string_echo += '{"stop_erp%s": %s'%(i,stop_erp[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"stop_cfm": ['
        stop_cfm_length = len(self.stop_cfm)
        for i in range(0, stop_cfm_length):
            if i == (stop_cfm_length - 1): 
                string_echo += '{"stop_cfm%s": %s'%(i,stop_cfm[i])
                string_echo += '}'
            else:
                string_echo += '{"stop_cfm%s": %s'%(i,stop_cfm[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"fudge_factor": ['
        fudge_factor_length = len(self.fudge_factor)
        for i in range(0, fudge_factor_length):
            if i == (fudge_factor_length - 1): 
                string_echo += '{"fudge_factor%s": %s'%(i,fudge_factor[i])
                string_echo += '}'
            else:
                string_echo += '{"fudge_factor%s": %s'%(i,fudge_factor[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"fmax": ['
        fmax_length = len(self.fmax)
        for i in range(0, fmax_length):
            if i == (fmax_length - 1): 
                string_echo += '{"fmax%s": %s'%(i,fmax[i])
                string_echo += '}'
            else:
                string_echo += '{"fmax%s": %s'%(i,fmax[i])
                string_echo += '}, '
        string_echo += '], '
        string_echo += '"vel": ['
        vel_length = len(self.vel)
        for i in range(0, vel_length):
            if i == (vel_length - 1): 
                string_echo += '{"vel%s": %s'%(i,vel[i])
                string_echo += '}'
            else:
                string_echo += '{"vel%s": %s'%(i,vel[i])
                string_echo += '}, '
        string_echo += ']'
        string_echo += '}'
        return string_echo

    def getType(self):
        return "gazebo_msgs/ODEJointProperties"

    def getMD5(self):
        return "a9e264dbf3eff8e202d2bebecf081639"

_struct_I = struct.Struct('<I')

