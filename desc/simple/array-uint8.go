package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/util"
)

type DescArrayUint8 struct{ DescSliceUint8 }

func (DescArrayUint8) Decode(r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(uint8(0)))).Elem()

	if err := util.Read(r, res.Slice(0, size).Interface().([]byte)); err != nil {
		return nil, err
	}

	return &res, nil
}
