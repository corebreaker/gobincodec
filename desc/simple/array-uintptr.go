package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayUintptr struct{ DescArrayUint64 }

func (DescArrayUintptr) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(uintptr(0)))).Elem()
	if size == 0 {
		return &res, nil
	}

	for i := 0; i < size; i++ {
		var v uint64

		if err := util.DecodeNum(r, &v); err != nil {
			return nil, err
		}

		res.Index(i).SetUint(v)
	}

	return &res, nil
}
