package simple

import (
	"bytes"
	"io"
	"reflect"
	"unsafe"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceUnsafePtr struct{ base.DescBase }

func (*DescSliceUnsafePtr) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
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

func (*DescSliceUnsafePtr) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]unsafe.Pointer, size)

	for i := 0; i < size; i++ {
		var v uint64

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		buf[i] = unsafe.Pointer(uintptr(v))
	}

	res := reflect.ValueOf(buf)

	return &res, nil
}
