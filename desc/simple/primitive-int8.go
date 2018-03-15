package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveInt8 struct{ base.DescBase }

func (*DescPrimitiveInt8) Encode(w io.Writer, v reflect.Value) error {
	return util.EncodeNum(w, int8(v.Int()))
}

func (*DescPrimitiveInt8) Decode(r io.Reader) (*reflect.Value, error) {
	var num int8

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(num)

	return &res, nil
}
