package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

var strCodec = new(DescPrimitiveBool)

type DescArrayString struct{ base.DescBase }

func (*DescArrayString) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := strCodec.Encode(spec, &out, v.Index(i)); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayString) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	count, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(count, reflect.TypeOf(""))).Elem()
	if count == 0 {
		return &res, nil
	}

	for i := 0; i < count; i++ {
		val, err := strCodec.Decode(spec, r)
		if err != nil {
			return nil, err
		}

		res.Index(i).Set(*val)
	}

	return &res, nil
}
