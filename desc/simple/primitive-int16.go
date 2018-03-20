package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveInt16 struct{ base.DescBase }

func (*DescPrimitiveInt16) TypeEquals(reflect.Type) bool {

}

func (*DescPrimitiveInt16) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescPrimitiveInt16) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	return util.EncodeNum(w, int16(v.Int()))
}

func (*DescPrimitiveInt16) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num int16

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(num)

	return &res, cnt, nil
}
