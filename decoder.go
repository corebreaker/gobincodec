package gobincodec

import (
	"io"
	"reflect"
	"time"
	"unsafe"

	"github.com/corebreaker/gobincodec/defs"
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

func (rd *tReader) Decode(values ...interface{}) error             { return nil }
func (rd *tReader) DecodeNext() (interface{}, error)               { return nil, nil }
func (rd *tReader) DecodeValue() (*reflect.Value, error)           { return nil, nil }
func (rd *tReader) DecodeBool() (bool, error)                      { return false, nil }
func (rd *tReader) DecodeByte() (byte, error)                      { return 0, nil }
func (rd *tReader) DecodeInt() (int, error)                        { return 0, nil }
func (rd *tReader) DecodeInt8() (int8, error)                      { return 0, nil }
func (rd *tReader) DecodeInt16() (int16, error)                    { return 0, nil }
func (rd *tReader) DecodeInt32() (int32, error)                    { return 0, nil }
func (rd *tReader) DecodeInt64() (int64, error)                    { return 0, nil }
func (rd *tReader) DecodeUint() (uint, error)                      { return 0, nil }
func (rd *tReader) DecodeUint8() (uint8, error)                    { return 0, nil }
func (rd *tReader) DecodeUint16() (uint16, error)                  { return 0, nil }
func (rd *tReader) DecodeUint32() (uint32, error)                  { return 0, nil }
func (rd *tReader) DecodeUint64() (uint64, error)                  { return 0, nil }
func (rd *tReader) DecodeUintptr() (uintptr, error)                { return 0, nil }
func (rd *tReader) DecodePtr() (unsafe.Pointer, error)             { return 0, nil }
func (rd *tReader) DecodeFloat32() (float32, error)                { return 0, nil }
func (rd *tReader) DecodeFloat64() (float64, error)                { return 0, nil }
func (rd *tReader) DecodeComplex64() (complex64, error)            { return 0, nil }
func (rd *tReader) DecodeComplex128() (complex128, error)          { return 0, nil }
func (rd *tReader) DecodeString() (string, error)                  { return "", nil }
func (rd *tReader) DecodeTime() (*time.Time, error)                { return nil, nil }
func (rd *tReader) DecodeBoolSlice() ([]bool, error)               { return nil, nil }
func (rd *tReader) DecodeByteSlice() ([]byte, error)               { return nil, nil }
func (rd *tReader) DecodeIntSlice() ([]int, error)                 { return nil, nil }
func (rd *tReader) DecodeInt8Slice() ([]int8, error)               { return nil, nil }
func (rd *tReader) DecodeInt16Slice() ([]int16, error)             { return nil, nil }
func (rd *tReader) DecodeInt32Slice() ([]int32, error)             { return nil, nil }
func (rd *tReader) DecodeInt64Slice() ([]int64, error)             { return nil, nil }
func (rd *tReader) DecodeUintSlice() ([]uint, error)               { return nil, nil }
func (rd *tReader) DecodeUint8Slice() ([]uint8, error)             { return nil, nil }
func (rd *tReader) DecodeUint16Slice() ([]uint16, error)           { return nil, nil }
func (rd *tReader) DecodeUint32Slice() ([]uint32, error)           { return nil, nil }
func (rd *tReader) DecodeUint64Slice() ([]uint64, error)           { return nil, nil }
func (rd *tReader) DecodeUintptrSlice() ([]uintptr, error)         { return nil, nil }
func (rd *tReader) DecodeFloat32Slice() ([]float32, error)         { return nil, nil }
func (rd *tReader) DecodeFloat64Slice() ([]float64, error)         { return nil, nil }
func (rd *tReader) DecodeComplex64Slice() ([]complex64, error)     { return nil, nil }
func (rd *tReader) DecodeComplex128Slice() ([]complex128, error)   { return nil, nil }
func (rd *tReader) DecodeStringSlice() ([]string, error)           { return nil, nil }
func (rd *tReader) DecodePtrSlice() ([]unsafe.Pointer, error)      { return nil, nil }
func (rd *tReader) DecodeTimeSlice() ([]time.Time, error)          { return nil, nil }
func (rd *tReader) DecodeSlice() ([]interface{}, error)            { return nil, nil }
func (rd *tReader) DecodeSerializable() (defs.Serializable, error) { return nil, nil }
