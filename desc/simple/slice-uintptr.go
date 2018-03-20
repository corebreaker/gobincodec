package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceUintptr struct{ DescArrayUintptr }

func (ds *DescSliceUintptr) TypeEquals(t reflect.Type) bool {
	return (t.Kind() == reflect.Slice) && (t.Elem().Kind() == reflect.Uintptr)
}

func (ds *DescSliceUintptr) Convert(v reflect.Value) *reflect.Value {
	if !v.IsValid() {
		return nil
	}

	t := reflect.SliceOf(reflect.TypeOf(uintptr(0)))
	if v.IsNil() {
		res := reflect.Zero(t)

		return &res
	}

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if ds.TypeEquals(v.Type()) {
		return &v
	}

	switch v.Kind() {
	case reflect.Array, reflect.Slice:
	default:
		res := reflect.MakeSlice(t, 1, 1)
		val := uintPtrCodec.Convert(v)
		if val == nil {
			return nil
		}

		res.Set(*val)

		return &res
	}

	sz := v.Len()
	res := reflect.MakeSlice(t, sz, v.Cap())
	for i := 0; i < sz; i++ {
		val := uintPtrCodec.Convert(v.Index(i))
		if val == nil {
			return nil
		}

		res.Index(i).Set(*val)
	}

	return &res
}

func (ds *DescSliceUintptr) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	cnt1, err := util.WriteBool(w, false)
	if err != nil {
		return 0, err
	}

	cnt2, err := ds.DescArrayUintptr.Encode(spec, w, v)
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2, nil
}

func (*DescSliceUintptr) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	isNil, cnt1, err := util.ReadBool(r)
	if err != nil {
		return nil, 0, err
	}

	if isNil {
		var val []uintptr

		res := reflect.ValueOf(&val).Elem()

		return &res, cnt1, nil
	}

	size, cnt2, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	cnt := cnt1 + cnt2
	buf := make([]uintptr, size)

	for i := 0; i < size; i++ {
		var v uint64

		n, err := util.DecodeNum(r, &v)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		buf[i] = uintptr(v)
	}

	res := reflect.ValueOf(buf)

	return &res, cnt, nil
}
