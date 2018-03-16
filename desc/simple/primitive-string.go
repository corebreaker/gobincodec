package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveString struct{ base.DescBase }

func (*DescPrimitiveString) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	var out bytes.Buffer

	value := []byte(v.String())
	if err := util.EncodeSize(&out, len(value)); err != nil {
		return err
	}

	if err := util.Write(&out, value); err != nil {
		return err
	}

	return util.Write(w, out.Bytes())
}

func (*DescPrimitiveString) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, size)

	if err := util.Read(r, buf); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(string(buf))

	return &res, nil
}
