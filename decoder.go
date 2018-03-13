package bincodec

import (
    "io"
    "reflect"
    "time"
    "unsafe"
)

type Decoder interface {
    Decode(...interface{}) error
    DecodeValue(reflect.Value) error
    DecodeBool(*bool) error
    DecodeByte(*byte) error
    DecodeInt(*int) error
    DecodeInt8(*int8) error
    DecodeInt16(*int16) error
    DecodeInt32(*int32) error
    DecodeInt64(*int64) error
    DecodeUint(*uint) error
    DecodeUint8(*uint8) error
    DecodeUint16(*uint16) error
    DecodeUint32(*uint32) error
    DecodeUint64(*uint64) error
    DecodeUintptr(*uintptr) error
    DecodePtr(*unsafe.Pointer) error
    DecodeFloat32(*float32) error
    DecodeFloat64(*float64) error
    DecodeComplex64(*complex64) error
    DecodeComplex128(*complex128) error
    DecodeString(*string) error
    DecodeTime(*time.Time) error
    DecodeBoolSlice(*[]bool) error
    DecodeByteSlice(*[]byte) error
    DecodeIntSlice(*[]int) error
    DecodeInt8Slice(*[]int8) error
    DecodeInt16Slice(*[]int16) error
    DecodeInt32Slice(*[]int32) error
    DecodeInt64Slice(*[]int64) error
    DecodeUintSlice(*[]uint) error
    DecodeUint8Slice(*[]uint8) error
    DecodeUint16Slice(*[]uint16) error
    DecodeUint32Slice(*[]uint32) error
    DecodeUint64Slice(*[]uint64) error
    DecodeUintptrSlice(*[]uintptr) error
    DecodeFloat32Slice(*[]float32) error
    DecodeFloat64Slice(*[]float64) error
    DecodeComplex64Slice(*[]complex64) error
    DecodeComplex128Slice(*[]complex128) error
    DecodeStringSlice(*[]string) error
    DecodeSlice(*[]interface{}) error
    DecodeSerializable(Serializable) error
}

type Reader interface {
    CodecBase
    IoBase
    Decoder

    CloneReader(io.Reader) Reader
    ReadHeader() error
}

type tReader struct {
    tIoBase

    r   io.Reader
}

func (self *tReader) CloneReader(r io.Reader) Reader {
    return &tReader{
        tIoBase: tIoBase{
            tCodecBase: tCodecBase{
                spec: self.spec.clone(),
            },
        },

        r:  r,
    }
}

func (self *tReader) ReadHeader() error {
    return self.spec.read(self.r)
}

func (self *tReader) Decode(values ...interface{}) error          { return nil }
func (self *tReader) DecodeValue(v reflect.Value) error           { return nil }
func (self *tReader) DecodeBool(v *bool) error                    { return nil }
func (self *tReader) DecodeByte(v *byte) error                    { return nil }
func (self *tReader) DecodeInt(v *int) error                      { return nil }
func (self *tReader) DecodeInt8(v *int8) error                    { return nil }
func (self *tReader) DecodeInt16(v *int16) error                  { return nil }
func (self *tReader) DecodeInt32(v *int32) error                  { return nil }
func (self *tReader) DecodeInt64(v *int64) error                  { return nil }
func (self *tReader) DecodeUint(v *uint) error                    { return nil }
func (self *tReader) DecodeUint8(v *uint8) error                  { return nil }
func (self *tReader) DecodeUint16(v *uint16) error                { return nil }
func (self *tReader) DecodeUint32(v *uint32) error                { return nil }
func (self *tReader) DecodeUint64(v *uint64) error                { return nil }
func (self *tReader) DecodeUintptr(v *uintptr) error              { return nil }
func (self *tReader) DecodePtr(v *unsafe.Pointer) error           { return nil }
func (self *tReader) DecodeFloat32(v *float32) error              { return nil }
func (self *tReader) DecodeFloat64(v *float64) error              { return nil }
func (self *tReader) DecodeComplex64(v *complex64) error          { return nil }
func (self *tReader) DecodeComplex128(v *complex128) error        { return nil }
func (self *tReader) DecodeString(v *string) error                { return nil }
func (self *tReader) DecodeTime(v *time.Time) error               { return nil }
func (self *tReader) DecodeBoolSlice(v *[]bool) error             { return nil }
func (self *tReader) DecodeByteSlice(v *[]byte) error             { return nil }
func (self *tReader) DecodeIntSlice(v *[]int) error               { return nil }
func (self *tReader) DecodeInt8Slice(v *[]int8) error             { return nil }
func (self *tReader) DecodeInt16Slice(v *[]int16) error           { return nil }
func (self *tReader) DecodeInt32Slice(v *[]int32) error           { return nil }
func (self *tReader) DecodeInt64Slice(v *[]int64) error           { return nil }
func (self *tReader) DecodeUintSlice(v *[]uint) error             { return nil }
func (self *tReader) DecodeUint8Slice(v *[]uint8) error           { return nil }
func (self *tReader) DecodeUint16Slice(v *[]uint16) error         { return nil }
func (self *tReader) DecodeUint32Slice(v *[]uint32) error         { return nil }
func (self *tReader) DecodeUint64Slice(v *[]uint64) error         { return nil }
func (self *tReader) DecodeUintptrSlice(v *[]uintptr) error       { return nil }
func (self *tReader) DecodeFloat32Slice(v *[]float32) error       { return nil }
func (self *tReader) DecodeFloat64Slice(v *[]float64) error       { return nil }
func (self *tReader) DecodeComplex64Slice(v *[]complex64) error   { return nil }
func (self *tReader) DecodeComplex128Slice(v *[]complex128) error { return nil }
func (self *tReader) DecodeStringSlice(v *[]string) error         { return nil }
func (self *tReader) DecodeSlice(v *[]interface{}) error          { return nil }
func (self *tReader) DecodeSerializable(v Serializable) error     { return nil }
