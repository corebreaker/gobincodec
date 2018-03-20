package desc

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"time"
	"unsafe"

	"github.com/corebreaker/gobincodec/defs"
	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/desc/composed"
	"github.com/corebreaker/gobincodec/desc/simple"
	"github.com/corebreaker/gobincodec/util"
	"github.com/corebreaker/goerrors"
)

var (
	ordinals = []string{
		"1st",
		"2nd",
		"3rd",
	}
)

func ordinalToString(i int) string {
	if i < 4 {
		return ordinals[i]
	}

	return fmt.Sprintf("%dth", i+1)
}

type tEncoder struct {
	serType map[uintptr]bool
	spec    base.ISpec
	r       io.Reader
}

func (enc *tEncoder) Read(p []byte) (int, error) {
	return util.Read(enc.r, p)
}

func (enc *tEncoder) Decode(values ...interface{}) error {
	for i, value := range values {
		addr := reflect.ValueOf(value)
		if v.Kind() != reflect.Ptr {
			return goerrors.DecorateError(fmt.Errorf("%s value must a pointer", ordinalToString(i)))
		}

		ref := addr.Elem()

		t := ref.Type()
		desc := enc.spec.DescFromType(t)
		if desc == nil {
			return goerrors.DecorateError(fmt.Errorf("%s value has unknown type %s", ordinalToString(i), t))
		}

		val, err := desc.Decode(enc.spec, enc.r)
		if err != nil {
			return err
		}

		func() {
			defer goerrors.Catch(&err, nil, nil)

			ref.Set(*val)
		}()

		if err != nil {
			return err
		}
	}

	return nil
}

func (enc *tEncoder) DecodeNext() (interface{}, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeValue() (*reflect.Value, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeBool() (bool, error) {
	return false, nil
}

func (enc *tEncoder) DecodeByte() (byte, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeInt() (int, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeInt8() (int8, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeInt16() (int16, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeInt32() (int32, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeInt64() (int64, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeUint() (uint, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeUint8() (uint8, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeUint16() (uint16, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeUint32() (uint32, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeUint64() (uint64, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeUintptr() (uintptr, error) {
	return 0, nil
}

func (enc *tEncoder) DecodePtr() (unsafe.Pointer, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeFloat32() (float32, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeFloat64() (float64, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeComplex64() (complex64, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeComplex128() (complex128, error) {
	return 0, nil
}

func (enc *tEncoder) DecodeString() (string, error) {
	return "", nil
}

func (enc *tEncoder) DecodeTime() (*time.Time, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeBoolSlice() ([]bool, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeByteSlice() ([]byte, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeIntSlice() ([]int, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeInt8Slice() ([]int8, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeInt16Slice() ([]int16, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeInt32Slice() ([]int32, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeInt64Slice() ([]int64, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeUintSlice() ([]uint, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeUint8Slice() ([]uint8, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeUint16Slice() ([]uint16, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeUint32Slice() ([]uint32, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeUint64Slice() ([]uint64, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeUintptrSlice() ([]uintptr, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeFloat32Slice() ([]float32, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeFloat64Slice() ([]float64, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeComplex64Slice() ([]complex64, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeComplex128Slice() ([]complex128, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeStringSlice() ([]string, error) {
	return nil, nil
}

func (enc *tEncoder) DecodePtrSlice() ([]unsafe.Pointer, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeTimeSlice() ([]time.Time, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeSlice() ([]interface{}, error) {
	return nil, nil
}

func (enc *tEncoder) DecodeSerializable() (Serializable, error) {
	return nil, nil
}
