package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveInt32 struct{ base.DescBase }

func (*DescPrimitiveInt32) Encode(w io.Writer, v reflect.Value) error {
	return util.EncodeNum(w, int32(v.Int()))
}

func (*DescPrimitiveInt32) Decode(r io.Reader) (*reflect.Value, error) {
	var num int32

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(num)

	return &res, nil
}
