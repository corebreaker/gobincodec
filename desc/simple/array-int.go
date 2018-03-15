package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/util"
)

type DescArrayInt struct{ DescSliceInt }

func (*DescArrayInt) Decode(r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(int(0)))).Elem()

	for i := 0; i < size; i++ {
		var v int64

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		res.Index(i).SetInt(v)
	}

	return &res, nil
}
