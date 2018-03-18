package composed

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescValueSlice struct {
	base.DescBase

	value base.IDesc
}

func (dv *DescValueSlice) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	cnt1, err := util.WriteBool(w, false)
	if err != nil {
		return 0, err
	}

	var out bytes.Buffer

	count := v.Len()
	if _, err := util.EncodeSize(&out, count); err != nil {
		return 0, err
	}

	for i := 0; i < count; i++ {
		if _, err := dv.value.Encode(spec, &out, v.Index(i)); err != nil {
			return 0, err
		}
	}

	cnt2, err := util.Write(w, out.Bytes())
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2, nil
}

func (dv *DescValueSlice) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	size, cnt, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	if size == 0 {
		res := reflect.Zero(spec.GetType(dv))

		return &res, cnt, nil
	}

	res := reflect.MakeSlice(spec.GetType(dv), size, size)

	for i := 0; i < size; i++ {
		value, n, err := dv.value.Decode(spec, r)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		res.Index(i).Set(*value)
	}

	return &res, cnt, nil
}

func (dv *DescValueSlice) Read(spec base.ISpec, r io.Reader) (int, error) {
	var err error
	var cnt int

	dv.value, cnt, err = spec.ReadDesc(r)

	return cnt, err
}

func (dv *DescValueSlice) Write(spec base.ISpec, w io.Writer) (int, error) {
	return spec.WriteDesc(w, dv.value)
}

func (dv *DescValueSlice) Make(spec base.ISpec, t reflect.Type) error {
	dv.value = spec.DescFromType(t.Elem())

	return nil
}

func NewSliceDesc(id base.DescId) base.IDesc {
	return &DescValueSlice{
		DescBase: base.NewBase(id),
	}
}
