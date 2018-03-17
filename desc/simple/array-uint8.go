package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayUint8 struct{ base.DescBase }

func (*DescArrayUint8) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	if count > 0 {
		if err := util.Write(&out, v.Bytes()); err != nil {
			return nil
		}
	}

	return util.Write(w, out.Bytes())
}

func (DescArrayUint8) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(uint8(0)))).Elem()
	if size == 0 {
		return &res, nil
	}

	if err := util.Read(r, res.Slice(0, size).Interface().([]byte)); err != nil {
		return nil, err
	}

	return &res, nil
}
