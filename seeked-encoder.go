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

func (enc *tSeekedEncoder) Encode(values ...interface{}) error {
    for v := range values {
        if err := enc.next(enc.w.encode(v, enc.p)); err != nil {
            return err
        }
    }

    return nil
}

func (enc *tSeekedEncoder) EncodeValue(v reflect.Value) error          { return enc.next(enc.w.encValue(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeBool(v bool) error                    { return enc.next(enc.w.encBool(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeByte(v byte) error                    { return enc.next(enc.w.encByte(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeInt(v int) error                      { return enc.next(enc.w.encInt(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeInt8(v int8) error                    { return enc.next(enc.w.encInt8(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeInt16(v int16) error                  { return enc.next(enc.w.encInt16(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeInt32(v int32) error                  { return enc.next(enc.w.encInt32(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeInt64(v int64) error                  { return enc.next(enc.w.encInt64(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUint(v uint) error                    { return enc.next(enc.w.encUint(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUint8(v uint8) error                  { return enc.next(enc.w.encUint8(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUint16(v uint16) error                { return enc.next(enc.w.encUint16(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUint32(v uint32) error                { return enc.next(enc.w.encUint32(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUint64(v uint64) error                { return enc.next(enc.w.encUint64(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUintptr(v uintptr) error              { return enc.next(enc.w.encUintptr(v, enc.p)) }
func (enc *tSeekedEncoder) EncodePtr(v unsafe.Pointer) error           { return enc.next(enc.w.encPtr(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeFloat32(v float32) error              { return enc.next(enc.w.encFloat32(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeFloat64(v float64) error              { return enc.next(enc.w.encFloat64(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeComplex64(v complex64) error          { return enc.next(enc.w.encCplx64(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeComplex128(v complex128) error        { return enc.next(enc.w.encCplx128(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeString(v string) error                { return enc.next(enc.w.encString(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeTime(v time.Time) error               { return enc.next(enc.w.encTime(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeBoolSlice(v []bool) error             { return enc.next(enc.w.encBools(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeByteSlice(v []byte) error             { return enc.next(enc.w.encBytes(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeIntSlice(v []int) error               { return enc.next(enc.w.encInts(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeInt8Slice(v []int8) error             { return enc.next(enc.w.encInt8s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeInt16Slice(v []int16) error           { return enc.next(enc.w.encInt16s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeInt32Slice(v []int32) error           { return enc.next(enc.w.encInt32s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeInt64Slice(v []int64) error           { return enc.next(enc.w.encInt64s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUintSlice(v []uint) error             { return enc.next(enc.w.encUints(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUint8Slice(v []uint8) error           { return enc.next(enc.w.encUint8s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUint16Slice(v []uint16) error         { return enc.next(enc.w.encUint16s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUint32Slice(v []uint32) error         { return enc.next(enc.w.encUint32s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUint64Slice(v []uint64) error         { return enc.next(enc.w.encUint64s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeUintptrSlice(v []uintptr) error       { return enc.next(enc.w.encUintptrs(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeFloat32Slice(v []float32) error       { return enc.next(enc.w.encFloat32s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeFloat64Slice(v []float64) error       { return enc.next(enc.w.encFloat64s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeComplex64Slice(v []complex64) error   { return enc.next(enc.w.encCplx64s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeComplex128Slice(v []complex128) error { return enc.next(enc.w.encCplx128s(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeStringSlice(v []string) error         { return enc.next(enc.w.encStrings(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeSlice(v []interface{}) error          { return enc.next(enc.w.encSlice(v, enc.p)) }
func (enc *tSeekedEncoder) EncodeSerializable(v Serializable) error    { return enc.next(enc.w.encSerial(v, enc.p)) }
