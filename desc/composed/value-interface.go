package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescValueInterface struct{ base.DescBase }

func (dv *DescValueInterface) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	cnt1, err := util.WriteBool(w, false)
	if err != nil {
		return 0, err
	}

	desc := spec.DescFromType(v.Type())

	cnt2, err := util.EncodeNum(w, desc.GetId())
	if err != nil {
		return 0, err
	}

	if v.IsNil() {
		return cnt1 + cnt2, nil
	}

	cnt3, err := desc.Encode(spec, w, v.Elem())
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2 + cnt3, nil
}

func (dv *DescValueInterface) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	isNil, cnt1, err := util.ReadBool(r)
	if err != nil {
		return nil, 0, err
	}

	if isNil {
		res := reflect.Zero(spec.GetType(dv))

		return &res, cnt1, nil
	}

	var id base.DescId

	cnt2, err := util.DecodeNum(r, &id)
	if err != nil {
		return nil, 0, err
	}

	desc := spec.DescFromId(id)
	res := reflect.New(spec.GetType(dv)).Elem()

	if desc.IsNil() {
		return &res, cnt1 + cnt2, nil
	}

	value, cnt3, err := desc.Decode(spec, r)
	if err != nil {
		return nil, 0, err
	}

	res.Set(*value)

	return &res, cnt1 + cnt2 + cnt3, nil
}

func NewInterfaceDesc(id base.DescId) base.IDesc {
	return &DescValueInterface{
		DescBase: base.NewBase(id),
	}
}
