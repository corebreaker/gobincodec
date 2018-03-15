package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveUint16 struct{ base.DescBase }

func (*DescPrimitiveUint16) Encode(w io.Writer, v reflect.Value) error {
	return util.EncodeNum(w, uint16(v.Uint()))
}

func (*DescPrimitiveUint16) Decode(r io.Reader) (*reflect.Value, error) {
	var num uint16

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(num)

	return &res, nil
}
