package bincodec

import (
    "io"
    "reflect"
    "time"
    "unsafe"
)

type Encoder interface {
    Encode(...interface{}) error
    EncodeValue(reflect.Value) error
    EncodeBool(bool) error
    EncodeByte(byte) error
    EncodeInt(int) error
    EncodeInt8(int8) error
    EncodeInt16(int16) error
    EncodeInt32(int32) error
    EncodeInt64(int64) error
    EncodeUint(uint) error
    EncodeUint8(uint8) error
    EncodeUint16(uint16) error
    EncodeUint32(uint32) error
    EncodeUint64(uint64) error
    EncodeUintptr(uintptr) error
    EncodePtr(unsafe.Pointer) error
    EncodeFloat32(float32) error
    EncodeFloat64(float64) error
    EncodeComplex64(complex64) error
    EncodeComplex128(complex128) error
    EncodeString(string) error
    EncodeTime(time.Time) error
    EncodeBoolSlice([]bool) error
    EncodeByteSlice([]byte) error
    EncodeIntSlice([]int) error
    EncodeInt8Slice([]int8) error
    EncodeInt16Slice([]int16) error
    EncodeInt32Slice([]int32) error
    EncodeInt64Slice([]int64) error
    EncodeUintSlice([]uint) error
    EncodeUint8Slice([]uint8) error
    EncodeUint16Slice([]uint16) error
    EncodeUint32Slice([]uint32) error
    EncodeUint64Slice([]uint64) error
    EncodeUintptrSlice([]uintptr) error
    EncodeFloat32Slice([]float32) error
    EncodeFloat64Slice([]float64) error
    EncodeComplex64Slice([]complex64) error
    EncodeComplex128Slice([]complex128) error
    EncodeStringSlice([]string) error
    EncodeSlice([]interface{}) error
    EncodeSerializable(Serializable) error
}

type Writer interface {
    CodecBase
    IoBase
    Encoder

    CloneWriter(io.Writer) Writer
    WriteHeader() error
}

type tWriter struct {
    tIoBase

    w   io.Writer
}

func (self *tWriter) CloneWriter(w io.Writer) Writer {
    return &tWriter{
        tIoBase: tIoBase{
            tCodecBase: tCodecBase{
                spec: self.spec.clone(),
            },
        },

        w:  w,
    }
}

func (self *tWriter) WriteHeader() error {
    return self.spec.write(self.w)
}

func (self *tWriter) Encode(values ...interface{}) error         { return nil }
func (self *tWriter) EncodeValue(v reflect.Value) error          { return nil }
func (self *tWriter) EncodeBool(v bool) error                    { return nil }
func (self *tWriter) EncodeByte(v byte) error                    { return nil }
func (self *tWriter) EncodeInt(v int) error                      { return nil }
func (self *tWriter) EncodeInt8(v int8) error                    { return nil }
func (self *tWriter) EncodeInt16(v int16) error                  { return nil }
func (self *tWriter) EncodeInt32(v int32) error                  { return nil }
func (self *tWriter) EncodeInt64(v int64) error                  { return nil }
func (self *tWriter) EncodeUint(v uint) error                    { return nil }
func (self *tWriter) EncodeUint8(v uint8) error                  { return nil }
func (self *tWriter) EncodeUint16(v uint16) error                { return nil }
func (self *tWriter) EncodeUint32(v uint32) error                { return nil }
func (self *tWriter) EncodeUint64(v uint64) error                { return nil }
func (self *tWriter) EncodeUintptr(v uintptr) error              { return nil }
func (self *tWriter) EncodePtr(v unsafe.Pointer) error           { return nil }
func (self *tWriter) EncodeFloat32(v float32) error              { return nil }
func (self *tWriter) EncodeFloat64(v float64) error              { return nil }
func (self *tWriter) EncodeComplex64(v complex64) error          { return nil }
func (self *tWriter) EncodeComplex128(v complex128) error        { return nil }
func (self *tWriter) EncodeString(v string) error                { return nil }
func (self *tWriter) EncodeTime(v time.Time) error               { return nil }
func (self *tWriter) EncodeBoolSlice(v []bool) error             { return nil }
func (self *tWriter) EncodeByteSlice(v []byte) error             { return nil }
func (self *tWriter) EncodeIntSlice(v []int) error               { return nil }
func (self *tWriter) EncodeInt8Slice(v []int8) error             { return nil }
func (self *tWriter) EncodeInt16Slice(v []int16) error           { return nil }
func (self *tWriter) EncodeInt32Slice(v []int32) error           { return nil }
func (self *tWriter) EncodeInt64Slice(v []int64) error           { return nil }
func (self *tWriter) EncodeUintSlice(v []uint) error             { return nil }
func (self *tWriter) EncodeUint8Slice(v []uint8) error           { return nil }
func (self *tWriter) EncodeUint16Slice(v []uint16) error         { return nil }
func (self *tWriter) EncodeUint32Slice(v []uint32) error         { return nil }
func (self *tWriter) EncodeUint64Slice(v []uint64) error         { return nil }
func (self *tWriter) EncodeUintptrSlice(v []uintptr) error       { return nil }
func (self *tWriter) EncodeFloat32Slice(v []float32) error       { return nil }
func (self *tWriter) EncodeFloat64Slice(v []float64) error       { return nil }
func (self *tWriter) EncodeComplex64Slice(v []complex64) error   { return nil }
func (self *tWriter) EncodeComplex128Slice(v []complex128) error { return nil }
func (self *tWriter) EncodeStringSlice(v []string) error         { return nil }
func (self *tWriter) EncodeSlice(v []interface{}) error          { return nil }
func (self *tWriter) EncodeSerializable(v Serializable) error    { return nil }
