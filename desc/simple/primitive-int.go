package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveInt struct{ DescPrimitiveInt64 }

func (*DescPrimitiveInt) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num int64

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(int(num))

	return &res, cnt, nil
}
