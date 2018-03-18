package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveUint struct{ DescPrimitiveUint64 }

func (*DescPrimitiveUint) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num uint64

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(uint(num))

	return &res, cnt, nil
}
