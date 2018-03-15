package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveFloat32 struct{ base.DescBase }

func (*DescPrimitiveFloat32) Encode(w io.Writer, v reflect.Value) error {
	return util.EncodeNum(w, float32(v.Float()))
}

func (*DescPrimitiveFloat32) Decode(r io.Reader) (*reflect.Value, error) {
	var num float32

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(num)

	return &res, nil
}
