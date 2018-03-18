package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceBool struct{ DescArrayBool }

func (ds *DescSliceBool) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	cnt1, err := util.WriteBool(w, false)
	if err != nil {
		return 0, err
	}

	cnt2, err := ds.DescArrayBool.Encode(spec, w, v)
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2, nil
}

func (*DescSliceBool) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	isNil, cnt1, err := util.ReadBool(r)
	if err != nil {
		return nil, 0, err
	}

	if isNil {
		var val []uint64

		res := reflect.ValueOf(&val).Elem()

		return &res, cnt1, nil
	}

	size, cnt2, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	buf := make([]byte, (size+7)/8)

	cnt3, err := util.Read(r, buf)
	if err != nil {
		return nil, 0, err
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

	return &res, cnt1 + cnt2 + cnt3, nil
}
