package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescValuePointer struct {
	base.DescBase

	elem base.IDesc
}

func (dv *DescValuePointer) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	if err := util.WriteBool(w, false); err != nil {
		return nil
	}

	return dv.elem.Encode(spec, w, v.Elem())
}

func (dv *DescValuePointer) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	isNil, err := util.ReadBool(r)
	if err != nil {
		return nil, err
	}

	if isNil {
		res := reflect.Zero(spec.GetType(dv))

		return &res, nil
	}

	value, err := dv.elem.Decode(spec, r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(spec.GetType(dv.elem))
	res.Elem().Set(*value)

	return &res, nil
}

func (dv *DescValuePointer) Read(spec base.ISpec, r io.Reader) error {
	var err error

	dv.elem, err = spec.ReadDesc(r)

	return err
}

func (dv *DescValuePointer) Write(spec base.ISpec, w io.Writer) error {
	return spec.WriteDesc(w, dv.elem)
}

func (dv *DescValuePointer) Make(spec base.ISpec, t reflect.Type) error {
	dv.elem = spec.DescFromType(t.Elem())

	return nil
}

func NewPointerDesc(id base.DescId) base.IDesc {
	return &DescValuePointer{
		DescBase: base.NewBase(id),
	}
}
