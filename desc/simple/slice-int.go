package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/util"
)

type DescSliceInt struct{ DescSliceInt64 }

func (*DescSliceInt) Decode(r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]int, size)

	for i := 0; i < size; i++ {
		var v int64

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		buf[i] = int(v)
	}

	res := reflect.ValueOf(buf)

	return &res, nil
}
