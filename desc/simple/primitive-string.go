package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveString struct{ base.DescBase }

func (*DescPrimitiveString) TypeEquals(reflect.Type) bool {

}

func (*DescPrimitiveString) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescPrimitiveString) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	var out bytes.Buffer

	value := []byte(v.String())
	if _, err := util.EncodeSize(&out, len(value)); err != nil {
		return 0, err
	}

	if _, err := util.Write(&out, value); err != nil {
		return 0, err
	}

	return util.Write(w, out.Bytes())
}

func (*DescPrimitiveString) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	size, cnt1, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	buf := make([]byte, size)

	cnt2, err := util.Read(r, buf)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(string(buf))

	return &res, cnt1 + cnt2, nil
}
