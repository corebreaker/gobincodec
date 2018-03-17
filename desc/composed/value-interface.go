package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescValueInterface struct{ base.DescBase }

func (dv *DescValueInterface) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	if err := util.WriteBool(w, false); err != nil {
		return nil
	}

	desc := spec.DescFromType(v.Type())
	if err := util.EncodeNum(w, desc.GetId()); err != nil {
		return err
	}

	if v.IsNil() {
		return nil
	}

	return desc.Encode(spec, w, v.Elem())
}

func (dv *DescValueInterface) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	isNil, err := util.ReadBool(r)
	if err != nil {
		return nil, err
	}

	if isNil {
		res := reflect.Zero(spec.GetType(dv))

		return &res, nil
	}

	var id base.DescId

	if err := util.DecodeNum(r, &id); err != nil {
		return nil, err
	}

	desc := spec.DescFromId(id)
	res := reflect.New(spec.GetType(dv)).Elem()

	if desc.IsNil() {
		return &res, nil
	}

	value, err := desc.Decode(spec, r)
	if err != nil {
		return nil, err
	}

	res.Set(*value)

	return &res, nil
}

func NewInterfaceDesc(id base.DescId) base.IDesc {
	return &DescValueInterface{
		DescBase: base.NewBase(id),
	}
}
