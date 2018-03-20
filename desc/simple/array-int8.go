package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayInt8 struct{ base.DescBase }

func (*DescArrayInt8) TypeEquals(reflect.Type) bool {

}

func (*DescArrayInt8) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescArrayInt8) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	count := v.Len()

	var out bytes.Buffer

	if _, err := util.EncodeSize(&out, count); err != nil {
		return 0, err
	}

	for i := 0; i < count; i++ {
		if err := util.WriteByte(&out, byte(v.Index(i).Int())); err != nil {
			return 0, err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayInt8) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	size, cnt1, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(int8(0)))).Elem()
	if size == 0 {
		return &res, cnt1, nil
	}

	buf := make([]byte, size)

	cnt2, err := util.Read(r, buf)
	if err != nil {
		return nil, 0, err
	}

	for i, b := range buf {
		res.Index(i).SetInt(int64(b))
	}

	return &res, cnt1 + cnt2, nil
}
