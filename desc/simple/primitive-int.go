package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveInt struct{ DescPrimitiveInt64 }

func (*DescPrimitiveInt) Decode(r io.Reader) (*reflect.Value, error) {
	var num int64

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(int(num))

	return &res, nil
}
