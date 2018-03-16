package simple

import (
	"io"
	"reflect"
	"unsafe"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveUnsafePtr struct{ base.DescBase }

func (*DescPrimitiveUnsafePtr) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	return util.EncodeNum(w, uint64(v.Pointer()))
}

func (*DescPrimitiveUnsafePtr) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	var num uint64

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(unsafe.Pointer(uintptr(num)))

	return &res, nil
}
