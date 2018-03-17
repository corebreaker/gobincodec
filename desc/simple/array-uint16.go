package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayUint16 struct{ base.DescBase }

func (*DescArrayUint16) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := util.EncodeNum(&out, uint16(v.Index(i).Uint())); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayUint16) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(uint16(0)))).Elem()
	if size == 0 {
		return &res, nil
	}

	for i := 0; i < size; i++ {
		var v uint16

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		res.Index(i).SetUint(uint64(v))
	}

	return &res, nil
}
