package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveUintptr struct{ DescPrimitiveUint64 }

func (*DescPrimitiveUintptr) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num uint64

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(uintptr(num))

	return &res, cnt, nil
}
