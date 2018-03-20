package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSimpleComplex128 struct{ base.DescBase }

func (*DescSimpleComplex128) TypeEquals(reflect.Type) bool {

}

func (*DescSimpleComplex128) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescSimpleComplex128) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	return util.EncodeNum(w, v.Complex())
}

func (*DescSimpleComplex128) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num complex128

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(num)

	return &res, cnt, nil
}
