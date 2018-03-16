package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveFloat64 struct{ base.DescBase }

func (*DescPrimitiveFloat64) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	return util.EncodeNum(w, v.Float())
}

func (*DescPrimitiveFloat64) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	var num float64

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(num)

	return &res, nil
}
