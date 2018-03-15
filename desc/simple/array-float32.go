package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/util"
)

type DescArrayFloat32 struct{ DescSliceFloat32 }

func (*DescArrayFloat32) Decode(r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(float32(0)))).Elem()

	for i := 0; i < size; i++ {
		var v float32

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		res.Index(i).SetFloat(float64(v))
	}

	return &res, nil
}
