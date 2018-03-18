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

func (dv *DescValuePointer) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	cnt1, err := util.WriteBool(w, false)
	if err != nil {
		return 0, err
	}

	cnt2, err := dv.elem.Encode(spec, w, v.Elem())
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2, nil
}

func (dv *DescValuePointer) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	isNil, cnt1, err := util.ReadBool(r)
	if err != nil {
		return nil, 0, err
	}

	if isNil {
		res := reflect.Zero(spec.GetType(dv))

		return &res, cnt1, nil
	}

	value, cnt2, err := dv.elem.Decode(spec, r)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.New(spec.GetType(dv.elem))
	res.Elem().Set(*value)

	return &res, cnt1 + cnt2, nil
}

func (dv *DescValuePointer) Read(spec base.ISpec, r io.Reader) (int, error) {
	var err error
    var cnt int

	dv.elem, cnt, err = spec.ReadDesc(r)

	return cnt, err
}

func (dv *DescValuePointer) Write(spec base.ISpec, w io.Writer) (int, error) {
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
