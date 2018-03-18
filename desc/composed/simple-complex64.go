package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSimpleComplex64 struct{ base.DescBase }

func (*DescSimpleComplex64) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	return util.EncodeNum(w, complex64(v.Complex()))
}

func (*DescSimpleComplex64) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num complex64

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(num)

	return &res, cnt, nil
}
