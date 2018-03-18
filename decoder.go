package gobincodec

import (
	"io"
	"reflect"
	"time"
	"unsafe"
)

type Decoder interface {
	Decode(...interface{}) error
	DecodeNext() (interface{}, error)
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
	DecodePtrSlice(*[]unsafe.Pointer) error
	DecodeTimeSlice(*[]time.Time) error
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

	r io.Reader
}

func (rd *tReader) CloneReader(r io.Reader) Reader {
	return &tReader{
		tIoBase: tIoBase{
			tCodecBase: tCodecBase{
				spec: rd.spec.clone(),
			},
		},

		r: r,
	}
}

func (rd *tReader) ReadHeader() error {
	return rd.spec.read(rd.r)
}

func (rd *tReader) Decode(values ...interface{}) error          { return nil }
func (rd *tReader) DecodeNext() (interface{}, error)            { return nil }
func (rd *tReader) DecodeValue(v reflect.Value) error           { return nil }
func (rd *tReader) DecodeBool(v *bool) error                    { return nil }
func (rd *tReader) DecodeByte(v *byte) error                    { return nil }
func (rd *tReader) DecodeInt(v *int) error                      { return nil }
func (rd *tReader) DecodeInt8(v *int8) error                    { return nil }
func (rd *tReader) DecodeInt16(v *int16) error                  { return nil }
func (rd *tReader) DecodeInt32(v *int32) error                  { return nil }
func (rd *tReader) DecodeInt64(v *int64) error                  { return nil }
func (rd *tReader) DecodeUint(v *uint) error                    { return nil }
func (rd *tReader) DecodeUint8(v *uint8) error                  { return nil }
func (rd *tReader) DecodeUint16(v *uint16) error                { return nil }
func (rd *tReader) DecodeUint32(v *uint32) error                { return nil }
func (rd *tReader) DecodeUint64(v *uint64) error                { return nil }
func (rd *tReader) DecodeUintptr(v *uintptr) error              { return nil }
func (rd *tReader) DecodePtr(v *unsafe.Pointer) error           { return nil }
func (rd *tReader) DecodeFloat32(v *float32) error              { return nil }
func (rd *tReader) DecodeFloat64(v *float64) error              { return nil }
func (rd *tReader) DecodeComplex64(v *complex64) error          { return nil }
func (rd *tReader) DecodeComplex128(v *complex128) error        { return nil }
func (rd *tReader) DecodeString(v *string) error                { return nil }
func (rd *tReader) DecodeTime(v *time.Time) error               { return nil }
func (rd *tReader) DecodeBoolSlice(v *[]bool) error             { return nil }
func (rd *tReader) DecodeByteSlice(v *[]byte) error             { return nil }
func (rd *tReader) DecodeIntSlice(v *[]int) error               { return nil }
func (rd *tReader) DecodeInt8Slice(v *[]int8) error             { return nil }
func (rd *tReader) DecodeInt16Slice(v *[]int16) error           { return nil }
func (rd *tReader) DecodeInt32Slice(v *[]int32) error           { return nil }
func (rd *tReader) DecodeInt64Slice(v *[]int64) error           { return nil }
func (rd *tReader) DecodeUintSlice(v *[]uint) error             { return nil }
func (rd *tReader) DecodeUint8Slice(v *[]uint8) error           { return nil }
func (rd *tReader) DecodeUint16Slice(v *[]uint16) error         { return nil }
func (rd *tReader) DecodeUint32Slice(v *[]uint32) error         { return nil }
func (rd *tReader) DecodeUint64Slice(v *[]uint64) error         { return nil }
func (rd *tReader) DecodeUintptrSlice(v *[]uintptr) error       { return nil }
func (rd *tReader) DecodeFloat32Slice(v *[]float32) error       { return nil }
func (rd *tReader) DecodeFloat64Slice(v *[]float64) error       { return nil }
func (rd *tReader) DecodeComplex64Slice(v *[]complex64) error   { return nil }
func (rd *tReader) DecodeComplex128Slice(v *[]complex128) error { return nil }
func (rd *tReader) DecodeStringSlice(v *[]string) error         { return nil }
func (rd *tReader) DecodePtrSlice(v *[]unsafe.Pointer) error    { return nil }
func (rd *tReader) DecodeTimeSlice(v *[]time.Time) error        { return nil }
func (rd *tReader) DecodeSlice(v *[]interface{}) error          { return nil }
func (rd *tReader) DecodeSerializable(v Serializable) error     { return nil }
