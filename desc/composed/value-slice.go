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

func (dv *DescValueSlice) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	var out bytes.Buffer

	count := v.Len()
	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := dv.value.Encode(spec, &out, v.Index(i)); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (dv *DescValueSlice) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	if size == 0 {
		res := reflect.Zero(spec.GetType(dv))

		return &res, nil
	}

	res := reflect.MakeSlice(spec.GetType(dv), size, size)

	for i := 0; i < size; i++ {
		value, err := dv.value.Decode(spec, r)
		if err != nil {
			return nil, err
		}

		res.Index(i).Set(*value)
	}

	return &res, nil
}

func (dv *DescValueSlice) Read(spec base.ISpec, r io.Reader) error {
	var err error

	dv.value, err = spec.ReadDesc(r)

	return err
}

func (dv *DescValueSlice) Write(spec base.ISpec, w io.Writer) error {
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
