package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayInt8 struct{ base.DescBase }

func (*DescArrayInt8) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := util.WriteByte(&out, byte(v.Index(i).Int())); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayInt8) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(int8(0)))).Elem()
	if size == 0 {
		return &res, nil
	}

	buf := make([]byte, size)
	if err := util.Read(r, buf); err != nil {
		return nil, err
	}

	for i, b := range buf {
		res.Index(i).SetInt(int64(b))
	}

	return &res, nil
}
