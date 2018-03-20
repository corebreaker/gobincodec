package gobincodec

import (
	"io"
	"reflect"
	"time"
	"unsafe"

	"github.com/corebreaker/gobincodec/defs"
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

func (wr *tWriter) Encode(values ...interface{}) error           { return nil }
func (wr *tWriter) EncodeValue(v reflect.Value) error            { return nil }
func (wr *tWriter) EncodeBool(v bool) error                      { return nil }
func (wr *tWriter) EncodeByte(v byte) error                      { return nil }
func (wr *tWriter) EncodeInt(v int) error                        { return nil }
func (wr *tWriter) EncodeInt8(v int8) error                      { return nil }
func (wr *tWriter) EncodeInt16(v int16) error                    { return nil }
func (wr *tWriter) EncodeInt32(v int32) error                    { return nil }
func (wr *tWriter) EncodeInt64(v int64) error                    { return nil }
func (wr *tWriter) EncodeUint(v uint) error                      { return nil }
func (wr *tWriter) EncodeUint8(v uint8) error                    { return nil }
func (wr *tWriter) EncodeUint16(v uint16) error                  { return nil }
func (wr *tWriter) EncodeUint32(v uint32) error                  { return nil }
func (wr *tWriter) EncodeUint64(v uint64) error                  { return nil }
func (wr *tWriter) EncodeUintptr(v uintptr) error                { return nil }
func (wr *tWriter) EncodePtr(v unsafe.Pointer) error             { return nil }
func (wr *tWriter) EncodeFloat32(v float32) error                { return nil }
func (wr *tWriter) EncodeFloat64(v float64) error                { return nil }
func (wr *tWriter) EncodeComplex64(v complex64) error            { return nil }
func (wr *tWriter) EncodeComplex128(v complex128) error          { return nil }
func (wr *tWriter) EncodeString(v string) error                  { return nil }
func (wr *tWriter) EncodeTime(v time.Time) error                 { return nil }
func (wr *tWriter) EncodeBoolSlice(v []bool) error               { return nil }
func (wr *tWriter) EncodeByteSlice(v []byte) error               { return nil }
func (wr *tWriter) EncodeIntSlice(v []int) error                 { return nil }
func (wr *tWriter) EncodeInt8Slice(v []int8) error               { return nil }
func (wr *tWriter) EncodeInt16Slice(v []int16) error             { return nil }
func (wr *tWriter) EncodeInt32Slice(v []int32) error             { return nil }
func (wr *tWriter) EncodeInt64Slice(v []int64) error             { return nil }
func (wr *tWriter) EncodeUintSlice(v []uint) error               { return nil }
func (wr *tWriter) EncodeUint8Slice(v []uint8) error             { return nil }
func (wr *tWriter) EncodeUint16Slice(v []uint16) error           { return nil }
func (wr *tWriter) EncodeUint32Slice(v []uint32) error           { return nil }
func (wr *tWriter) EncodeUint64Slice(v []uint64) error           { return nil }
func (wr *tWriter) EncodeUintptrSlice(v []uintptr) error         { return nil }
func (wr *tWriter) EncodeFloat32Slice(v []float32) error         { return nil }
func (wr *tWriter) EncodeFloat64Slice(v []float64) error         { return nil }
func (wr *tWriter) EncodeComplex64Slice(v []complex64) error     { return nil }
func (wr *tWriter) EncodeComplex128Slice(v []complex128) error   { return nil }
func (wr *tWriter) EncodeStringSlice(v []string) error           { return nil }
func (wr *tWriter) EncodePtrSlice(v []unsafe.Pointer) error      { return nil }
func (wr *tWriter) EncodeTimeSlice(v []time.Time) error          { return nil }
func (wr *tWriter) EncodeSlice(v []interface{}) error            { return nil }
func (wr *tWriter) EncodeSerializable(v defs.Serializable) error { return nil }
