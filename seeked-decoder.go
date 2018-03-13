package bincodec

import (
    "reflect"
    "time"
    "unsafe"
)

type tSeekedDecoder struct {
    tSeekedCodec

    r   *tDecAt
}

func (self *tSeekedDecoder) Decode(values ...interface{}) error {
    for v := range values {
        if err := self.next(self.r.decode(v, self.p)); err != nil {
            return err
        }
    }

    return nil
}

func (self *tSeekedDecoder) DecodeValue(v reflect.Value) error           { return self.next(self.r.dec_value(v, self.p)) }
func (self *tSeekedDecoder) DecodeBool(v *bool) error                    { return self.next(self.r.dec_bool(v, self.p)) }
func (self *tSeekedDecoder) DecodeByte(v *byte) error                    { return self.next(self.r.dec_byte(v, self.p)) }
func (self *tSeekedDecoder) DecodeInt(v *int) error                      { return self.next(self.r.dec_int(v, self.p)) }
func (self *tSeekedDecoder) DecodeInt8(v *int8) error                    { return self.next(self.r.dec_int8(v, self.p)) }
func (self *tSeekedDecoder) DecodeInt16(v *int16) error                  { return self.next(self.r.dec_int16(v, self.p)) }
func (self *tSeekedDecoder) DecodeInt32(v *int32) error                  { return self.next(self.r.dec_int32(v, self.p)) }
func (self *tSeekedDecoder) DecodeInt64(v *int64) error                  { return self.next(self.r.dec_int64(v, self.p)) }
func (self *tSeekedDecoder) DecodeUint(v *uint) error                    { return self.next(self.r.dec_uint(v, self.p)) }
func (self *tSeekedDecoder) DecodeUint8(v *uint8) error                  { return self.next(self.r.dec_uint8(v, self.p)) }
func (self *tSeekedDecoder) DecodeUint16(v *uint16) error                { return self.next(self.r.dec_uint16(v, self.p)) }
func (self *tSeekedDecoder) DecodeUint32(v *uint32) error                { return self.next(self.r.dec_uint32(v, self.p)) }
func (self *tSeekedDecoder) DecodeUint64(v *uint64) error                { return self.next(self.r.dec_uint64(v, self.p)) }
func (self *tSeekedDecoder) DecodeUintptr(v *uintptr) error              { return self.next(self.r.dec_uintptr(v, self.p)) }
func (self *tSeekedDecoder) DecodePtr(v *unsafe.Pointer) error           { return self.next(self.r.dec_ptr(v, self.p)) }
func (self *tSeekedDecoder) DecodeFloat32(v *float32) error              { return self.next(self.r.dec_float32(v, self.p)) }
func (self *tSeekedDecoder) DecodeFloat64(v *float64) error              { return self.next(self.r.dec_float64(v, self.p)) }
func (self *tSeekedDecoder) DecodeComplex64(v *complex64) error          { return self.next(self.r.dec_cplx64(v, self.p)) }
func (self *tSeekedDecoder) DecodeComplex128(v *complex128) error        { return self.next(self.r.dec_cplx128(v, self.p)) }
func (self *tSeekedDecoder) DecodeString(v *string) error                { return self.next(self.r.dec_string(v, self.p)) }
func (self *tSeekedDecoder) DecodeTime(v *time.Time) error               { return self.next(self.r.dec_time(v, self.p)) }
func (self *tSeekedDecoder) DecodeBoolSlice(v *[]bool) error             { return self.next(self.r.dec_bools(v, self.p)) }
func (self *tSeekedDecoder) DecodeByteSlice(v *[]byte) error             { return self.next(self.r.dec_bytes(v, self.p)) }
func (self *tSeekedDecoder) DecodeIntSlice(v *[]int) error               { return self.next(self.r.dec_ints(v, self.p)) }
func (self *tSeekedDecoder) DecodeInt8Slice(v *[]int8) error             { return self.next(self.r.dec_int8s(v, self.p)) }
func (self *tSeekedDecoder) DecodeInt16Slice(v *[]int16) error           { return self.next(self.r.dec_int16s(v, self.p)) }
func (self *tSeekedDecoder) DecodeInt32Slice(v *[]int32) error           { return self.next(self.r.dec_int32s(v, self.p)) }
func (self *tSeekedDecoder) DecodeInt64Slice(v *[]int64) error           { return self.next(self.r.dec_int64s(v, self.p)) }
func (self *tSeekedDecoder) DecodeUintSlice(v *[]uint) error             { return self.next(self.r.dec_uints(v, self.p)) }
func (self *tSeekedDecoder) DecodeUint8Slice(v *[]uint8) error           { return self.next(self.r.dec_uint8s(v, self.p)) }
func (self *tSeekedDecoder) DecodeUint16Slice(v *[]uint16) error         { return self.next(self.r.dec_uint16s(v, self.p)) }
func (self *tSeekedDecoder) DecodeUint32Slice(v *[]uint32) error         { return self.next(self.r.dec_uint32s(v, self.p)) }
func (self *tSeekedDecoder) DecodeUint64Slice(v *[]uint64) error         { return self.next(self.r.dec_uint64s(v, self.p)) }
func (self *tSeekedDecoder) DecodeUintptrSlice(v *[]uintptr) error       { return self.next(self.r.dec_uintptrs(v, self.p)) }
func (self *tSeekedDecoder) DecodeFloat32Slice(v *[]float32) error       { return self.next(self.r.dec_float32s(v, self.p)) }
func (self *tSeekedDecoder) DecodeFloat64Slice(v *[]float64) error       { return self.next(self.r.dec_float64s(v, self.p)) }
func (self *tSeekedDecoder) DecodeComplex64Slice(v *[]complex64) error   { return self.next(self.r.dec_cplx64s(v, self.p)) }
func (self *tSeekedDecoder) DecodeComplex128Slice(v *[]complex128) error { return self.next(self.r.dec_cplx128s(v, self.p)) }
func (self *tSeekedDecoder) DecodeStringSlice(v *[]string) error         { return self.next(self.r.dec_strings(v, self.p)) }
func (self *tSeekedDecoder) DecodeSlice(v *[]interface{}) error          { return self.next(self.r.dec_slice(v, self.p)) }
func (self *tSeekedDecoder) DecodeSerializable(v Serializable) error     { return self.next(self.r.dec_serial(v, self.p)) }
