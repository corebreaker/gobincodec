package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSimpleComplex64 struct{ base.DescBase }

func (*DescSimpleComplex64) Encode(w io.Writer, v reflect.Value) error {
	return util.EncodeNum(w, complex64(v.Complex()))
}

func (*DescSimpleComplex64) Decode(r io.Reader) (*reflect.Value, error) {
	var num complex64

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(num)

	return &res, nil
}
