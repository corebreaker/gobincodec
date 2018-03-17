package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceInt16 struct{ DescArrayInt16 }

func (ds *DescSliceInt16) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	if err := util.WriteBool(w, false); err != nil {
		return nil
	}

	return ds.DescArrayInt16.Encode(spec, w, v)
}

func (*DescSliceInt16) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	isNil, err := util.ReadBool(r)
	if err != nil {
		return nil, err
	}

	if isNil {
		var val []int16

		res := reflect.ValueOf(&val).Elem()

		return &res, nil
	}

	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]int16, size)

	for i := 0; i < size; i++ {
		var v int16

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		buf[i] = v
	}

	res := reflect.ValueOf(buf)

	return &res, nil
}
