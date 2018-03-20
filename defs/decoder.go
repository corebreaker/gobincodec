package defs

import (
	"io"
	"reflect"
	"time"
	"unsafe"
)

type Decoder interface {
	io.Reader

	Decode(...interface{}) error
	DecodeNext() (interface{}, error)
	DecodeValue() (*reflect.Value, error)
	DecodeBool() (bool, error)
	DecodeByte() (byte, error)
	DecodeInt() (int, error)
	DecodeInt8() (int8, error)
	DecodeInt16() (int16, error)
	DecodeInt32() (int32, error)
	DecodeInt64() (int64, error)
	DecodeUint() (uint, error)
	DecodeUint8() (uint8, error)
	DecodeUint16() (uint16, error)
	DecodeUint32() (uint32, error)
	DecodeUint64() (uint64, error)
	DecodeUintptr() (uintptr, error)
	DecodePtr() (unsafe.Pointer, error)
	DecodeFloat32() (float32, error)
	DecodeFloat64() (float64, error)
	DecodeComplex64() (complex64, error)
	DecodeComplex128() (complex128, error)
	DecodeString() (string, error)
	DecodeTime() (*time.Time, error)
	DecodeBoolSlice() ([]bool, error)
	DecodeByteSlice() ([]byte, error)
	DecodeIntSlice() ([]int, error)
	DecodeInt8Slice() ([]int8, error)
	DecodeInt16Slice() ([]int16, error)
	DecodeInt32Slice() ([]int32, error)
	DecodeInt64Slice() ([]int64, error)
	DecodeUintSlice() ([]uint, error)
	DecodeUint8Slice() ([]uint8, error)
	DecodeUint16Slice() ([]uint16, error)
	DecodeUint32Slice() ([]uint32, error)
	DecodeUint64Slice() ([]uint64, error)
	DecodeUintptrSlice() ([]uintptr, error)
	DecodeFloat32Slice() ([]float32, error)
	DecodeFloat64Slice() ([]float64, error)
	DecodeComplex64Slice() ([]complex64, error)
	DecodeComplex128Slice() ([]complex128, error)
	DecodeStringSlice() ([]string, error)
	DecodePtrSlice() ([]unsafe.Pointer, error)
	DecodeTimeSlice() ([]time.Time, error)
	DecodeSlice() ([]interface{}, error)
	DecodeSerializable() (Serializable, error)
}
