package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/util"
)

type DescArrayBool struct{ DescSliceBool }

func (*DescArrayBool) Decode(r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, (size+7)/8)

	if err := util.Read(r, buf); err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(false))).Elem()

	var current byte
	var pos uint
	var j int

	for i := 0; i < size; i++ {
		if pos == 0 {
			current, pos = buf[j], 7
			j++
		}

		res.Index(i).SetBool(((current >> pos) & 1) != 0)
		pos--
	}

	return &res, nil
}
