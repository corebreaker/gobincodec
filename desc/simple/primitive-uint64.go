package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveUint64 struct{ base.DescBase }

func (*DescPrimitiveUint64) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	return util.EncodeNum(w, v.Uint())
}

func (*DescPrimitiveUint64) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num uint64

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(num)

	return &res, cnt, nil
}
