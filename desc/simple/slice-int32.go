package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceInt32 struct{ base.DescBase }

func (*DescSliceInt32) Encode(w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := util.EncodeNum(&out, int32(v.Index(i).Int())); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescSliceInt32) Decode(r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]int32, size)

	for i := 0; i < size; i++ {
		var v int32

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		buf[i] = v
	}

	res := reflect.ValueOf(buf)

	return &res, nil
}
