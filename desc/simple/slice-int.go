package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceInt struct{ DescArrayInt }

func (ds *DescSliceInt) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	if err := util.WriteBool(w, false); err != nil {
		return nil
	}

	return ds.DescArrayInt.Encode(spec, w, v)
}

func (*DescSliceInt) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	isNil, err := util.ReadBool(r)
	if err != nil {
		return nil, err
	}

	if isNil {
		var val []int

		res := reflect.ValueOf(&val).Elem()

		return &res, nil
	}

	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]int, size)

	for i := 0; i < size; i++ {
		var v int64

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		buf[i] = int(v)
	}

	res := reflect.ValueOf(buf)

	return &res, nil
}
