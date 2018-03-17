package composed

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescValueMap struct {
	base.DescBase

	index base.IDesc
	value base.IDesc
}

func (dv *DescValueMap) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	if err := util.WriteBool(w, false); err != nil {
		return nil
	}

	var out bytes.Buffer

	if err := util.EncodeSize(&out, v.Len()); err != nil {
		return err
	}

	for _, key := range v.MapKeys() {
		if err := dv.index.Encode(spec, &out, key); err != nil {
			return err
		}

		if err := dv.value.Encode(spec, &out, v.MapIndex(key)); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (dv *DescValueMap) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	isNil, err := util.ReadBool(r)
	if err != nil {
		return nil, err
	}

	if isNil {
		res := reflect.Zero(spec.GetType(dv))

		return &res, nil
	}

	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.MakeMap(spec.GetType(dv))

	for i := 0; i < size; i++ {
		index, err := dv.index.Decode(spec, r)
		if err != nil {
			return nil, err
		}

		value, err := dv.value.Decode(spec, r)
		if err != nil {
			return nil, err
		}

		res.SetMapIndex(*index, *value)
	}

	return &res, nil
}

func (dv *DescValueMap) Read(spec base.ISpec, r io.Reader) error {
	var err error

	dv.index, err = spec.ReadDesc(r)
	if err != nil {
		return err
	}

	dv.value, err = spec.ReadDesc(r)

	return err
}

func (dv *DescValueMap) Write(spec base.ISpec, w io.Writer) error {
	if err := spec.WriteDesc(w, dv.index); err != nil {
		return err
	}

	return spec.WriteDesc(w, dv.value)
}

func (dv *DescValueMap) Make(spec base.ISpec, t reflect.Type) error {
	dv.index, dv.value = spec.DescFromType(t.Key()), spec.DescFromType(t.Elem())

	return nil
}

func NewMapDesc(id base.DescId) base.IDesc {
	return &DescValueMap{
		DescBase: base.NewBase(id),
	}
}
