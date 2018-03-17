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

func (*DescArrayUnsafePtr) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := util.EncodeNum(&out, uint64(v.Index(i).Pointer())); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (DescArrayUnsafePtr) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(unsafe.Pointer(nil)))).Elem()
	if size == 0 {
		return &res, nil
	}

	for i := 0; i < size; i++ {
		var v uint64

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		res.Index(i).SetPointer(unsafe.Pointer(uintptr(v)))
	}

	return &res, nil
}
