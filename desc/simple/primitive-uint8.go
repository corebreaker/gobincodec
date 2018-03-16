package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveUint8 struct{ base.DescBase }

func (*DescPrimitiveUint8) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	return util.EncodeNum(w, uint8(v.Uint()))
}

func (*DescPrimitiveUint8) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	var num uint8

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(num)

	return &res, nil
}
