package composed

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceComplex128 struct{ base.DescBase }

func (*DescSliceComplex128) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := util.EncodeNum(&out, v.Index(i).Complex()); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescSliceComplex128) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]complex128, size)

	for i := 0; i < size; i++ {
		var v complex128

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		buf[i] = v
	}

	res := reflect.ValueOf(buf)

	return &res, nil
}
