package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceUint32 struct{ DescArrayUint32 }

func (ds *DescSliceUint32) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	if err := util.WriteBool(w, false); err != nil {
		return nil
	}

	return ds.DescArrayUint32.Encode(spec, w, v)
}

func (*DescSliceUint32) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	isNil, err := util.ReadBool(r)
	if err != nil {
		return nil, err
	}

	if isNil {
		var val []uint32

		res := reflect.ValueOf(&val).Elem()

		return &res, nil
	}

	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]uint32, size)

	for i := 0; i < size; i++ {
		var v uint32

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		buf[i] = v
	}

	res := reflect.ValueOf(buf)

	return &res, nil
}
