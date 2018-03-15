package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveBool struct{ base.DescBase }

func (*DescPrimitiveBool) Encode(w io.Writer, v reflect.Value) error {
	return util.EncodeNum(w, v.Bool())
}

func (*DescPrimitiveBool) Decode(r io.Reader) (*reflect.Value, error) {
	var val bool

	if err := util.DecodeNum(r, &val); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(val)

	return &res, nil
}
