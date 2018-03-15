package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSimpleComplex128 struct{ base.DescBase }

func (*DescSimpleComplex128) Encode(w io.Writer, v reflect.Value) error {
	return util.EncodeNum(w, v.Complex())
}

func (*DescSimpleComplex128) Decode(r io.Reader) (*reflect.Value, error) {
	var num complex128

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(num)

	return &res, nil
}
