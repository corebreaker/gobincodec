package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayInt16 struct{ base.DescBase }

func (*DescArrayInt16) TypeEquals(reflect.Type) bool {

}

func (*DescArrayInt16) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescArrayInt16) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	count := v.Len()

	var out bytes.Buffer

	if _, err := util.EncodeSize(&out, count); err != nil {
		return 0, err
	}

	for i := 0; i < count; i++ {
		if _, err := util.EncodeNum(&out, int16(v.Index(i).Int())); err != nil {
			return 0, err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayInt16) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	size, cnt, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(int16(0)))).Elem()
	if size == 0 {
		return &res, cnt, nil
	}

	for i := 0; i < size; i++ {
		var v int16

		n, err := util.DecodeNum(r, &v)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		res.Index(i).SetInt(int64(v))
	}

	return &res, cnt, nil
}
