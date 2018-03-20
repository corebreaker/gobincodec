package simple

import (
	"io"
	"reflect"
	"unsafe"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveUnsafePtr struct{ base.DescBase }

func (*DescPrimitiveUnsafePtr) TypeEquals(reflect.Type) bool {

}

func (*DescPrimitiveUnsafePtr) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescPrimitiveUnsafePtr) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	return util.EncodeNum(w, uint64(v.Pointer()))
}

func (*DescPrimitiveUnsafePtr) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num uint64

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(unsafe.Pointer(uintptr(num)))

	return &res, cnt, nil
}
