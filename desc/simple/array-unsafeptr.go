package simple

import (
	"bytes"
	"io"
	"reflect"
	"unsafe"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayUnsafePtr struct{ base.DescBase }

func (*DescArrayUnsafePtr) TypeEquals(reflect.Type) bool {

}

func (*DescArrayUnsafePtr) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescArrayUnsafePtr) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	count := v.Len()

	var out bytes.Buffer

	if _, err := util.EncodeSize(&out, count); err != nil {
		return 0, err
	}

	for i := 0; i < count; i++ {
		if _, err := util.EncodeNum(&out, uint64(v.Index(i).Pointer())); err != nil {
			return 0, err
		}
	}

	return util.Write(w, out.Bytes())
}

func (DescArrayUnsafePtr) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	size, cnt, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(unsafe.Pointer(nil)))).Elem()
	if size == 0 {
		return &res, cnt, nil
	}

	for i := 0; i < size; i++ {
		var v uint64

		n, err := util.DecodeNum(r, &v)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		res.Index(i).SetPointer(unsafe.Pointer(uintptr(v)))
	}

	return &res, cnt, nil
}
