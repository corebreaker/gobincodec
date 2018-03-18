package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescArrayUint8 struct{ base.DescBase }

func (*DescArrayUint8) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	count := v.Len()

	var out bytes.Buffer

	if _, err := util.EncodeSize(&out, count); err != nil {
		return 0, err
	}

	if count > 0 {
		if _, err := util.Write(&out, v.Bytes()); err != nil {
			return 0, err
		}
	}

	return util.Write(w, out.Bytes())
}

func (DescArrayUint8) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	size, cnt1, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf(uint8(0)))).Elem()
	if size == 0 {
		return &res, cnt1, nil
	}

	cnt2, err := util.Read(r, res.Slice(0, size).Interface().([]byte))
	if err != nil {
		return nil, 0, err
	}

	return &res, cnt1 + cnt2, nil
}
