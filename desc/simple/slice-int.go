package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceInt struct{ DescArrayInt }

func (ds *DescSliceInt) TypeEquals(reflect.Type) bool {

}

func (ds *DescSliceInt) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (ds *DescSliceInt) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	cnt1, err := util.WriteBool(w, false)
	if err != nil {
		return 0, err
	}

	cnt2, err := ds.DescArrayInt.Encode(spec, w, v)
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2, nil
}

func (*DescSliceInt) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	isNil, cnt1, err := util.ReadBool(r)
	if err != nil {
		return nil, 0, err
	}

	if isNil {
		var val []int

		res := reflect.ValueOf(&val).Elem()

		return &res, cnt1, nil
	}

	size, cnt2, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	cnt := cnt1 + cnt2
	buf := make([]int, size)

	for i := 0; i < size; i++ {
		var v int64

		n, err := util.DecodeNum(r, &v)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		buf[i] = int(v)
	}

	res := reflect.ValueOf(buf)

	return &res, cnt, nil
}
