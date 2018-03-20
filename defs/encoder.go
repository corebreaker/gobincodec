package defs

import (
	"io"
	"reflect"
	"time"
	"unsafe"
)

type Encoder interface {
	io.Writer

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
	EncodePtrSlice([]unsafe.Pointer) error
	EncodeTimeSlice([]time.Time) error
	EncodeSlice([]interface{}) error
	EncodeSerializable(Serializable) error
}
