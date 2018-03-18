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

func (*DescArrayString) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	count := v.Len()

	var out bytes.Buffer

	if _, err := util.EncodeSize(&out, count); err != nil {
		return 0, err
	}

	for i := 0; i < count; i++ {
		if _, err := strCodec.Encode(spec, &out, v.Index(i)); err != nil {
			return 0, err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayString) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	count, cnt, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.New(reflect.ArrayOf(count, reflect.TypeOf(""))).Elem()
	if count == 0 {
		return &res, cnt, nil
	}

	for i := 0; i < count; i++ {
		val, n, err := strCodec.Decode(spec, r)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		res.Index(i).Set(*val)
	}

	return &res, cnt, nil
}
