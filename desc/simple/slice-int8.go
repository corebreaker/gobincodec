package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceInt8 struct{ DescArrayInt8 }

func (ds *DescSliceInt8) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	if err := util.WriteBool(w, false); err != nil {
		return nil
	}

	return ds.DescArrayInt8.Encode(spec, w, v)
}

func (*DescSliceInt8) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	isNil, err := util.ReadBool(r)
	if err != nil {
		return nil, err
	}

	if isNil {
		var val []int8

		res := reflect.ValueOf(&val).Elem()

		return &res, nil
	}

	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, size)
	if err := util.Read(r, buf); err != nil {
		return nil, err
	}

	in := make([]int8, size)
	for i, b := range buf {
		in[i] = int8(b)
	}

	res := reflect.ValueOf(in)

	return &res, nil
}
