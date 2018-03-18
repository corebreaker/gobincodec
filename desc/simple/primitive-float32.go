package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveFloat32 struct{ base.DescBase }

func (*DescPrimitiveFloat32) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	return util.EncodeNum(w, float32(v.Float()))
}

func (*DescPrimitiveFloat32) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num float32

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(num)

	return &res, cnt, nil
}
