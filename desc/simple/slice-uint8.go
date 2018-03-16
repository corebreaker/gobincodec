package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceUint8 struct{ base.DescBase }

func (*DescSliceUint8) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	if err := util.Write(&out, v.Bytes()); err != nil {
		return nil
	}

	return util.Write(w, out.Bytes())
}

func (*DescSliceUint8) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, size)
	if err := util.Read(r, buf); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(buf)

	return &res, nil
}
