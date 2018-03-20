package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveInt8 struct{ base.DescBase }

func (*DescPrimitiveInt8) TypeEquals(reflect.Type) bool {

}

func (*DescPrimitiveInt8) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescPrimitiveInt8) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	return util.EncodeNum(w, int8(v.Int()))
}

func (*DescPrimitiveInt8) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num int8

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(num)

	return &res, cnt, nil
}
