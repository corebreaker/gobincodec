package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayInt64 struct{ base.DescBase }

func (*DescArrayInt64) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := util.EncodeNum(&out, v.Index(i).Int()); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayInt64) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(int64(0)))).Elem()
	if size == 0 {
		return &res, nil
	}

	for i := 0; i < size; i++ {
		var v int64

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		res.Index(i).SetInt(v)
	}

	return &res, nil
}
