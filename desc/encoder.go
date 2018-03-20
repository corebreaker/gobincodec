package desc

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"time"
	"unsafe"

	"github.com/corebreaker/gobincodec/defs"
	"github.com/corebreaker/gobincodec/desc/composed"
	"github.com/corebreaker/gobincodec/desc/simple"
	"github.com/corebreaker/gobincodec/util"
	"github.com/corebreaker/goerrors"
)

type tEncoder struct {
	serType map[uintptr]bool
	spec    base.ISpec
	w       bytes.Buffer
}

func (enc *tEncoder) Write(p []byte) (int, error) {
	return util.Write(p)
}

func (enc *tEncoder) Encode(...interface{}) error {
	return nil
}

func (enc *tEncoder) EncodeValue(reflect.Value) error {
	return nil
}

func (enc *tEncoder) EncodeBool(bool) error {
	return nil
}

func (enc *tEncoder) EncodeByte(byte) error {
	return nil
}

func (enc *tEncoder) EncodeInt(int) error {
	return nil
}

func (enc *tEncoder) EncodeInt8(int8) error {
	return nil
}

func (enc *tEncoder) EncodeInt16(int16) error {
	return nil
}

func (enc *tEncoder) EncodeInt32(int32) error {
	return nil
}

func (enc *tEncoder) EncodeInt64(int64) error {
	return nil
}

func (enc *tEncoder) EncodeUint(uint) error {
	return nil
}

func (enc *tEncoder) EncodeUint8(uint8) error {
	return nil
}

func (enc *tEncoder) EncodeUint16(uint16) error {
	return nil
}

func (enc *tEncoder) EncodeUint32(uint32) error {
	return nil
}

func (enc *tEncoder) EncodeUint64(uint64) error {
	return nil
}

func (enc *tEncoder) EncodeUintptr(uintptr) error {
	return nil
}

func (enc *tEncoder) EncodePtr(unsafe.Pointer) error {
	return nil
}

func (enc *tEncoder) EncodeFloat32(float32) error {
	return nil
}

func (enc *tEncoder) EncodeFloat64(float64) error {
	return nil
}

func (enc *tEncoder) EncodeComplex64(complex64) error {
	return nil
}

func (enc *tEncoder) EncodeComplex128(complex128) error {
	return nil
}

func (enc *tEncoder) EncodeString(string) error {
	return nil
}

func (enc *tEncoder) EncodeTime(time.Time) error {
	return nil
}

func (enc *tEncoder) EncodeBoolSlice([]bool) error {
	return nil
}

func (enc *tEncoder) EncodeByteSlice([]byte) error {
	return nil
}

func (enc *tEncoder) EncodeIntSlice([]int) error {
	return nil
}

func (enc *tEncoder) EncodeInt8Slice([]int8) error {
	return nil
}

func (enc *tEncoder) EncodeInt16Slice([]int16) error {
	return nil
}

func (enc *tEncoder) EncodeInt32Slice([]int32) error {
	return nil
}

func (enc *tEncoder) EncodeInt64Slice([]int64) error {
	return nil
}

func (enc *tEncoder) EncodeUintSlice([]uint) error {
	return nil
}

func (enc *tEncoder) EncodeUint8Slice([]uint8) error {
	return nil
}

func (enc *tEncoder) EncodeUint16Slice([]uint16) error {
	return nil
}

func (enc *tEncoder) EncodeUint32Slice([]uint32) error {
	return nil
}

func (enc *tEncoder) EncodeUint64Slice([]uint64) error {
	return nil
}

func (enc *tEncoder) EncodeUintptrSlice([]uintptr) error {
	return nil
}

func (enc *tEncoder) EncodeFloat32Slice([]float32) error {
	return nil
}

func (enc *tEncoder) EncodeFloat64Slice([]float64) error {
	return nil
}

func (enc *tEncoder) EncodeComplex64Slice([]complex64) error {
	return nil
}

func (enc *tEncoder) EncodeComplex128Slice([]complex128) error {
	return nil
}

func (enc *tEncoder) EncodeStringSlice([]string) error {
	return nil
}

func (enc *tEncoder) EncodePtrSlice([]unsafe.Pointer) error {
	return nil
}

func (enc *tEncoder) EncodeTimeSlice([]time.Time) error {
	return nil
}

func (enc *tEncoder) EncodeSlice([]interface{}) error {
	return nil
}

func (enc *tEncoder) EncodeSerializable(defs.Serializable) error {
	return nil
}
