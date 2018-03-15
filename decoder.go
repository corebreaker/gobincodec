package bincodec

import (
	"io"
	"reflect"
	"time"
	"unsafe"
)

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

func (self *tReader) CloneReader(r io.Reader) Reader {
	return &tReader{
		tIoBase: tIoBase{
			tCodecBase: tCodecBase{
				spec: self.spec.clone(),
			},
		},

		r: r,
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
