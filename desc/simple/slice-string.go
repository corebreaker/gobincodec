package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceString struct{ DescArrayString }

func (ds *DescSliceString) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	if err := util.WriteBool(w, false); err != nil {
		return nil
	}

	return ds.DescArrayString.Encode(spec, w, v)
}

func (*DescSliceString) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	isNil, err := util.ReadBool(r)
	if err != nil {
		return nil, err
	}

	if isNil {
		var val []string

		res := reflect.ValueOf(&val).Elem()

		return &res, nil
	}

	count, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	in := make([]string, count)

	for i := 0; i < count; i++ {
		val, err := strCodec.Decode(spec, r)
		if err != nil {
			return nil, err
		}

		in[i] = val.String()
	}

	res := reflect.ValueOf(in)

	return &res, nil
}
