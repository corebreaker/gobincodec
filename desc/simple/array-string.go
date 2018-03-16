package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayString struct{ DescSliceString }

func (*DescArrayString) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	code := make([]byte, size)
	if err := util.Read(r, code); err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(code)

	count, err := util.DecodeSize(buf)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(count, reflect.TypeOf(""))).Elem()

	for i := 0; i < count; i++ {
		val, err := strCodec.Decode(spec, buf)
		if err != nil {
			return nil, err
		}

		res.Index(i).Set(*val)
	}

	return &res, nil
}
