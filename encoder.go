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
	EncodeTimeSlice([]time.Time) error
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

	w io.Writer
}

func (wr *tWriter) CloneWriter(w io.Writer) Writer {
	return &tWriter{
		tIoBase: tIoBase{
			tCodecBase: tCodecBase{
				spec: wr.spec.clone(),
			},
		},

		w: w,
	}
}

func (wr *tWriter) WriteHeader() error {
	return wr.spec.write(wr.w)
}

func (wr *tWriter) Encode(values ...interface{}) error         { return nil }
func (wr *tWriter) EncodeValue(v reflect.Value) error          { return nil }
func (wr *tWriter) EncodeBool(v bool) error                    { return nil }
func (wr *tWriter) EncodeByte(v byte) error                    { return nil }
func (wr *tWriter) EncodeInt(v int) error                      { return nil }
func (wr *tWriter) EncodeInt8(v int8) error                    { return nil }
func (wr *tWriter) EncodeInt16(v int16) error                  { return nil }
func (wr *tWriter) EncodeInt32(v int32) error                  { return nil }
func (wr *tWriter) EncodeInt64(v int64) error                  { return nil }
func (wr *tWriter) EncodeUint(v uint) error                    { return nil }
func (wr *tWriter) EncodeUint8(v uint8) error                  { return nil }
func (wr *tWriter) EncodeUint16(v uint16) error                { return nil }
func (wr *tWriter) EncodeUint32(v uint32) error                { return nil }
func (wr *tWriter) EncodeUint64(v uint64) error                { return nil }
func (wr *tWriter) EncodeUintptr(v uintptr) error              { return nil }
func (wr *tWriter) EncodePtr(v unsafe.Pointer) error           { return nil }
func (wr *tWriter) EncodeFloat32(v float32) error              { return nil }
func (wr *tWriter) EncodeFloat64(v float64) error              { return nil }
func (wr *tWriter) EncodeComplex64(v complex64) error          { return nil }
func (wr *tWriter) EncodeComplex128(v complex128) error        { return nil }
func (wr *tWriter) EncodeString(v string) error                { return nil }
func (wr *tWriter) EncodeTime(v time.Time) error               { return nil }
func (wr *tWriter) EncodeBoolSlice(v []bool) error             { return nil }
func (wr *tWriter) EncodeByteSlice(v []byte) error             { return nil }
func (wr *tWriter) EncodeIntSlice(v []int) error               { return nil }
func (wr *tWriter) EncodeInt8Slice(v []int8) error             { return nil }
func (wr *tWriter) EncodeInt16Slice(v []int16) error           { return nil }
func (wr *tWriter) EncodeInt32Slice(v []int32) error           { return nil }
func (wr *tWriter) EncodeInt64Slice(v []int64) error           { return nil }
func (wr *tWriter) EncodeUintSlice(v []uint) error             { return nil }
func (wr *tWriter) EncodeUint8Slice(v []uint8) error           { return nil }
func (wr *tWriter) EncodeUint16Slice(v []uint16) error         { return nil }
func (wr *tWriter) EncodeUint32Slice(v []uint32) error         { return nil }
func (wr *tWriter) EncodeUint64Slice(v []uint64) error         { return nil }
func (wr *tWriter) EncodeUintptrSlice(v []uintptr) error       { return nil }
func (wr *tWriter) EncodeFloat32Slice(v []float32) error       { return nil }
func (wr *tWriter) EncodeFloat64Slice(v []float64) error       { return nil }
func (wr *tWriter) EncodeComplex64Slice(v []complex64) error   { return nil }
func (wr *tWriter) EncodeComplex128Slice(v []complex128) error { return nil }
func (wr *tWriter) EncodeStringSlice(v []string) error         { return nil }
func (wr *tWriter) EncodeSlice(v []interface{}) error          { return nil }
func (wr *tWriter) EncodeSerializable(v Serializable) error    { return nil }
