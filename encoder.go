package bincodec

import (
	"io"
	"reflect"
	"time"
	"unsafe"
)

type Writer interface {
	CodecBase
	IoBase
	Encoder

	CloneWriter(io.Writer) Writer
	WriteHeader() error
}

type tWriter struct {
	tIoBase

	w io.Writer
}

func (self *tWriter) CloneWriter(w io.Writer) Writer {
	return &tWriter{
		tIoBase: tIoBase{
			tCodecBase: tCodecBase{
				spec: self.spec.clone(),
			},
		},

		w: w,
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
