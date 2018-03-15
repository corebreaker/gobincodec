package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/util"
)

type DescSliceUintptr struct{ DescSliceUint64 }

func (*DescSliceUintptr) Decode(r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]uintptr, size)

	for i := 0; i < size; i++ {
		var v uint64

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		buf[i] = uintptr(v)
	}

	res := reflect.ValueOf(buf)

	return &res, nil
}
