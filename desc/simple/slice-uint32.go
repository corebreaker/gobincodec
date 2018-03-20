package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceUint32 struct{ DescArrayUint32 }

func (ds *DescSliceUint32) TypeEquals(reflect.Type) bool {

}

func (ds *DescSliceUint32) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (ds *DescSliceUint32) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	cnt1, err := util.WriteBool(w, false)
	if err != nil {
		return 0, err
	}

	cnt2, err := ds.DescArrayUint32.Encode(spec, w, v)
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2, nil
}

func (*DescSliceUint32) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	isNil, cnt1, err := util.ReadBool(r)
	if err != nil {
		return nil, 0, err
	}

	if isNil {
		var val []uint32

		res := reflect.ValueOf(&val).Elem()

		return &res, cnt1, nil
	}

	size, cnt2, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	cnt := cnt1 + cnt2
	buf := make([]uint32, size)

	for i := 0; i < size; i++ {
		var v uint32

		n, err := util.DecodeNum(r, &v)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		buf[i] = v
	}

	res := reflect.ValueOf(buf)

	return &res, cnt, nil
}
