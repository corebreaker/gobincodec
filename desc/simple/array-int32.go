package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/util"
)

type DescArrayInt32 struct{ DescSliceInt32 }

func (*DescArrayInt32) Decode(r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(int32(0)))).Elem()

	for i := 0; i < size; i++ {
		var v int32

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		res.Index(i).SetInt(int64(v))
	}

	return &res, nil
}
