package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceBool struct{ DescArrayBool }

func (ds *DescSliceBool) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	if err := util.WriteBool(w, false); err != nil {
		return nil
	}

	return ds.DescArrayBool.Encode(spec, w, v)
}

func (*DescSliceBool) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	isNil, err := util.ReadBool(r)
	if err != nil {
		return nil, err
	}

	if isNil {
		var val []bool

		res := reflect.ValueOf(&val).Elem()

		return &res, nil
	}

	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, (size+7)/8)

	if err := util.Read(r, buf); err != nil {
		return nil, err
	}

	in := make([]bool, size)

	var current byte
	var pos uint
	var j int

	for i := 0; i < size; i++ {
		if pos == 0 {
			current, pos = buf[j], 7
			j++
		}

		in[i] = ((current >> pos) & 1) != 0
		pos--
	}

	res := reflect.ValueOf(in)

	return &res, nil
}
