package bincodec

import (
    "reflect"
    "time"
    "unsafe"
)

type tSeekedEncoder struct {
    tSeekedCodec

    w   *tEncAt
}

func (self *tSeekedEncoder) Encode(values ...interface{}) error {
    for v := range values {
        if err := self.next(self.w.encode(v, self.p)); err != nil {
            return err
        }
    }

    return nil
}

func (self *tSeekedEncoder) EncodeValue(v reflect.Value) error          { return self.next(self.w.enc_value(v, self.p)) }
func (self *tSeekedEncoder) EncodeBool(v bool) error                    { return self.next(self.w.enc_bool(v, self.p)) }
func (self *tSeekedEncoder) EncodeByte(v byte) error                    { return self.next(self.w.enc_byte(v, self.p)) }
func (self *tSeekedEncoder) EncodeInt(v int) error                      { return self.next(self.w.enc_int(v, self.p)) }
func (self *tSeekedEncoder) EncodeInt8(v int8) error                    { return self.next(self.w.enc_int8(v, self.p)) }
func (self *tSeekedEncoder) EncodeInt16(v int16) error                  { return self.next(self.w.enc_int16(v, self.p)) }
func (self *tSeekedEncoder) EncodeInt32(v int32) error                  { return self.next(self.w.enc_int32(v, self.p)) }
func (self *tSeekedEncoder) EncodeInt64(v int64) error                  { return self.next(self.w.enc_int64(v, self.p)) }
func (self *tSeekedEncoder) EncodeUint(v uint) error                    { return self.next(self.w.enc_uint(v, self.p)) }
func (self *tSeekedEncoder) EncodeUint8(v uint8) error                  { return self.next(self.w.enc_uint8(v, self.p)) }
func (self *tSeekedEncoder) EncodeUint16(v uint16) error                { return self.next(self.w.enc_uint16(v, self.p)) }
func (self *tSeekedEncoder) EncodeUint32(v uint32) error                { return self.next(self.w.enc_uint32(v, self.p)) }
func (self *tSeekedEncoder) EncodeUint64(v uint64) error                { return self.next(self.w.enc_uint64(v, self.p)) }
func (self *tSeekedEncoder) EncodeUintptr(v uintptr) error              { return self.next(self.w.enc_uintptr(v, self.p)) }
func (self *tSeekedEncoder) EncodePtr(v unsafe.Pointer) error           { return self.next(self.w.enc_ptr(v, self.p)) }
func (self *tSeekedEncoder) EncodeFloat32(v float32) error              { return self.next(self.w.enc_float32(v, self.p)) }
func (self *tSeekedEncoder) EncodeFloat64(v float64) error              { return self.next(self.w.enc_float64(v, self.p)) }
func (self *tSeekedEncoder) EncodeComplex64(v complex64) error          { return self.next(self.w.enc_cplx64(v, self.p)) }
func (self *tSeekedEncoder) EncodeComplex128(v complex128) error        { return self.next(self.w.enc_cplx128(v, self.p)) }
func (self *tSeekedEncoder) EncodeString(v string) error                { return self.next(self.w.enc_string(v, self.p)) }
func (self *tSeekedEncoder) EncodeTime(v time.Time) error               { return self.next(self.w.enc_time(v, self.p)) }
func (self *tSeekedEncoder) EncodeBoolSlice(v []bool) error             { return self.next(self.w.enc_bools(v, self.p)) }
func (self *tSeekedEncoder) EncodeByteSlice(v []byte) error             { return self.next(self.w.enc_bytes(v, self.p)) }
func (self *tSeekedEncoder) EncodeIntSlice(v []int) error               { return self.next(self.w.enc_ints(v, self.p)) }
func (self *tSeekedEncoder) EncodeInt8Slice(v []int8) error             { return self.next(self.w.enc_int8s(v, self.p)) }
func (self *tSeekedEncoder) EncodeInt16Slice(v []int16) error           { return self.next(self.w.enc_int16s(v, self.p)) }
func (self *tSeekedEncoder) EncodeInt32Slice(v []int32) error           { return self.next(self.w.enc_int32s(v, self.p)) }
func (self *tSeekedEncoder) EncodeInt64Slice(v []int64) error           { return self.next(self.w.enc_int64s(v, self.p)) }
func (self *tSeekedEncoder) EncodeUintSlice(v []uint) error             { return self.next(self.w.enc_uints(v, self.p)) }
func (self *tSeekedEncoder) EncodeUint8Slice(v []uint8) error           { return self.next(self.w.enc_uint8s(v, self.p)) }
func (self *tSeekedEncoder) EncodeUint16Slice(v []uint16) error         { return self.next(self.w.enc_uint16s(v, self.p)) }
func (self *tSeekedEncoder) EncodeUint32Slice(v []uint32) error         { return self.next(self.w.enc_uint32s(v, self.p)) }
func (self *tSeekedEncoder) EncodeUint64Slice(v []uint64) error         { return self.next(self.w.enc_uint64s(v, self.p)) }
func (self *tSeekedEncoder) EncodeUintptrSlice(v []uintptr) error       { return self.next(self.w.enc_uintptrs(v, self.p)) }
func (self *tSeekedEncoder) EncodeFloat32Slice(v []float32) error       { return self.next(self.w.enc_float32s(v, self.p)) }
func (self *tSeekedEncoder) EncodeFloat64Slice(v []float64) error       { return self.next(self.w.enc_float64s(v, self.p)) }
func (self *tSeekedEncoder) EncodeComplex64Slice(v []complex64) error   { return self.next(self.w.enc_cplx64s(v, self.p)) }
func (self *tSeekedEncoder) EncodeComplex128Slice(v []complex128) error { return self.next(self.w.enc_cplx128s(v, self.p)) }
func (self *tSeekedEncoder) EncodeStringSlice(v []string) error         { return self.next(self.w.enc_strings(v, self.p)) }
func (self *tSeekedEncoder) EncodeSlice(v []interface{}) error          { return self.next(self.w.enc_slice(v, self.p)) }
func (self *tSeekedEncoder) EncodeSerializable(v Serializable) error    { return self.next(self.w.enc_serial(v, self.p)) }
