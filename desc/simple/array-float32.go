package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayFloat32 struct{ base.DescBase }

func (*DescArrayFloat32) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := util.EncodeNum(&out, float32(v.Index(i).Float())); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayFloat32) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(float32(0)))).Elem()
	if size == 0 {
		return &res, nil
	}

	for i := 0; i < size; i++ {
		var v float32

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		res.Index(i).SetFloat(float64(v))
	}

	return &res, nil
}
