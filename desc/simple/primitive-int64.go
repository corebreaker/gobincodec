package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveInt64 struct{ base.DescBase }

func (*DescPrimitiveInt64) TypeEquals(reflect.Type) bool {

}

func (*DescPrimitiveInt64) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescPrimitiveInt64) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	return util.EncodeNum(w, v.Int())
}

func (*DescPrimitiveInt64) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num int64

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(num)

	return &res, cnt, nil
}
