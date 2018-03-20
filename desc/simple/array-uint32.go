package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayUint32 struct{ base.DescBase }

func (*DescArrayUint32) TypeEquals(reflect.Type) bool {

}

func (*DescArrayUint32) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (*DescArrayUint32) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	count := v.Len()

	var out bytes.Buffer

	if _, err := util.EncodeSize(&out, count); err != nil {
		return 0, err
	}

	for i := 0; i < count; i++ {
		if _, err := util.EncodeNum(&out, uint32(v.Index(i).Uint())); err != nil {
			return 0, err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayUint32) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	size, cnt, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(uint32(0)))).Elem()
	if size == 0 {
		return &res, cnt, nil
	}

	for i := 0; i < size; i++ {
		var v uint32

		n, err := util.DecodeNum(r, &v)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		res.Index(i).SetUint(uint64(v))
	}

	return &res, cnt, nil
}
