package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveUintptr struct{ DescPrimitiveUint64 }

func (*DescPrimitiveUintptr) Decode(r io.Reader) (*reflect.Value, error) {
	var num uint64

	if err := util.DecodeNum(r, &num); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(uintptr(num))

	return &res, nil
}
