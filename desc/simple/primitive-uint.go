package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveUint struct{ DescPrimitiveUint64 }

func (*DescPrimitiveUint) Decode(r io.Reader) (*reflect.Value, error) {
	var num uint64

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(uint(num))

	return &res, nil
}
