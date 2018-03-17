package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceUint struct{ DescArrayUint }

func (ds *DescSliceUint) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	if err := util.WriteBool(w, false); err != nil {
		return nil
	}

	return ds.DescArrayUint.Encode(spec, w, v)
}

func (*DescSliceUint) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	isNil, err := util.ReadBool(r)
	if err != nil {
		return nil, err
	}

	if isNil {
		var val []uint

		res := reflect.ValueOf(&val).Elem()

		return &res, nil
	}

	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]uint, size)

	for i := 0; i < size; i++ {
		var v uint64

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		buf[i] = uint(v)
	}

	res := reflect.ValueOf(buf)

	return &res, nil
}
