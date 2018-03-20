package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceInt8 struct{ DescArrayInt8 }

func (ds *DescSliceInt8) TypeEquals(reflect.Type) bool {

}

func (ds *DescSliceInt8) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (ds *DescSliceInt8) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	cnt1, err := util.WriteBool(w, false)
	if err != nil {
		return 0, err
	}

	cnt2, err := ds.DescArrayInt8.Encode(spec, w, v)
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2, nil
}

func (*DescSliceInt8) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	isNil, cnt1, err := util.ReadBool(r)
	if err != nil {
		return nil, 0, err
	}

	if isNil {
		var val []int8

		res := reflect.ValueOf(&val).Elem()

		return &res, cnt1, nil
	}

	size, cnt2, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	buf := make([]byte, size)

	cnt3, err := util.Read(r, buf)
	if err != nil {
		return nil, 0, err
	}

	in := make([]int8, size)
	for i, b := range buf {
		in[i] = int8(b)
	}

	res := reflect.ValueOf(in)

	return &res, cnt1 + cnt2 + cnt3, nil
}
