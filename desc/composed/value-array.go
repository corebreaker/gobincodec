package composed

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescValueArray struct {
	base.DescBase

	count int
	value base.IDesc
}

func (dv *DescValueArray) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if dv.count == 0 {
		return nil
	}

	var out bytes.Buffer

	for i := 0; i < dv.count; i++ {
		if err := dv.value.Encode(spec, &out, v.Index(i)); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (dv *DescValueArray) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	res := reflect.New(spec.GetType(dv)).Elem()

	for i := 0; i < dv.count; i++ {
		value, err := dv.value.Decode(spec, r)
		if err != nil {
			return nil, err
		}

		res.Field(i).Set(*value)
	}

	return &res, nil
}

func (dv *DescValueArray) Read(spec base.ISpec, r io.Reader) error {
	var err error

	dv.count, err = util.DecodeSize(r)
	if err != nil {
		return err
	}

	dv.value, err = spec.ReadDesc(r)

	return err
}

func (dv *DescValueArray) Write(spec base.ISpec, w io.Writer) error {
	if err := util.EncodeSize(w, dv.count); err != nil {
		return err
	}

	return spec.WriteDesc(w, dv.value)
}

func (dv *DescValueArray) Make(spec base.ISpec, t reflect.Type) error {
	dv.count = t.Len()
	dv.value = spec.DescFromType(t.Elem())

	return nil
}

func NewArrayDesc(id base.DescId) base.IDesc {
	return &DescValueArray{
		DescBase: base.NewBase(id),
	}
}
