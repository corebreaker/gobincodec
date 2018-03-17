package composed

import (
	"io"
	"reflect"
	"time"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceTime struct{ DescArrayTime }

func (ds *DescSliceTime) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	if err := util.WriteBool(w, false); err != nil {
		return nil
	}

	return ds.DescArrayTime.Encode(spec, w, v)
}

func (*DescSliceTime) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	isNil, err := util.ReadBool(r)
	if err != nil {
		return nil, err
	}

	if isNil {
		var val []time.Time

		res := reflect.ValueOf(&val).Elem()

		return &res, nil
	}

	count, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	in := make([]time.Time, count)

	for i := 0; i < count; i++ {
		val, err := timeCodec.Decode(spec, r)
		if err != nil {
			return nil, err
		}

		in[i] = val.Interface().(time.Time)
	}

	res := reflect.ValueOf(in)

	return &res, nil
}
