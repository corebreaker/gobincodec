package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayInt8 struct{ DescSliceInt8 }

func (*DescArrayInt8) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, size)
	if err := util.Read(r, buf); err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(int8(0)))).Elem()

	for i, b := range buf {
		res.Index(i).SetInt(int64(b))
	}

	return &res, nil
}
