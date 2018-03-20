package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveFloat64 struct{ base.DescBase }

func (*DescPrimitiveFloat64) TypeEquals(reflect.Type) bool {

}

func (*DescPrimitiveFloat64) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescPrimitiveFloat64) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	return util.EncodeNum(w, v.Float())
}

func (*DescPrimitiveFloat64) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num float64

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(num)

	return &res, cnt, nil
}
