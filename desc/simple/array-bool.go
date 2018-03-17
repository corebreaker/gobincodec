package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayBool struct{ base.DescBase }

func (*DescArrayBool) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	if count > 0 {
		var current byte

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
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayBool) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(false))).Elem()
	if size == 0 {
		return &res, nil
	}

	buf := make([]byte, (size+7)/8)

	if err := util.Read(r, buf); err != nil {
		return nil, err
	}

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
