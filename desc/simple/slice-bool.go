package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceBool struct{ base.DescBase }

func (*DescSliceBool) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer
	var current byte

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	pos := uint(7)

	for i := 0; i < count; i++ {
		if v.Index(i).Bool() {
			if pos == 0 {
				current |= byte(1)
				if err := util.WriteByte(&out, current); err != nil {
					return err
				}

				pos, current = 7, 0
			} else {
				current |= byte(1 << pos)
				pos--
			}
		}
	}

	if pos != 7 {
		if err := util.WriteByte(&out, current); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescSliceBool) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
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
