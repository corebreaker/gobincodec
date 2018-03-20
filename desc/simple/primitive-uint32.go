package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveUint32 struct{ base.DescBase }

func (*DescPrimitiveUint32) TypeEquals(reflect.Type) bool {

}

func (*DescPrimitiveUint32) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescPrimitiveUint32) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	return util.EncodeNum(w, uint32(v.Uint()))
}

func (*DescPrimitiveUint32) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num uint32

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(num)

	return &res, cnt, nil
}
