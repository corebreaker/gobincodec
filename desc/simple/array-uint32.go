package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/util"
)

type DescArrayUint32 struct{ DescSliceUint32 }

func (*DescArrayUint32) Decode(r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(uint32(0)))).Elem()

	for i := 0; i < size; i++ {
		var v uint32

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		res.Index(i).SetUint(uint64(v))
	}

	return &res, nil
}
