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

func (dv *DescValueMap) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	cnt1, err := util.WriteBool(w, false)
	if err != nil {
		return 0, err
	}

	var out bytes.Buffer

	if _, err := util.EncodeSize(&out, v.Len()); err != nil {
		return 0, err
	}

	for _, key := range v.MapKeys() {
		if _, err := dv.index.Encode(spec, &out, key); err != nil {
			return 0, err
		}

		if _, err := dv.value.Encode(spec, &out, v.MapIndex(key)); err != nil {
			return 0, err
		}
	}

	cnt2, err := util.Write(w, out.Bytes())
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2, nil
}

func (dv *DescValueMap) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	isNil, cnt1, err := util.ReadBool(r)
	if err != nil {
		return nil, 0, err
	}

	if isNil {
		res := reflect.Zero(spec.GetType(dv))

		return &res, cnt1, nil
	}

	size, cnt2, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	cnt := cnt1 + cnt2
	res := reflect.MakeMap(spec.GetType(dv))

	for i := 0; i < size; i++ {
		index, n1, err := dv.index.Decode(spec, r)
		if err != nil {
			return nil, 0, err
		}

		value, n2, err := dv.value.Decode(spec, r)
		if err != nil {
			return nil, 0, err
		}

		cnt += n1 + n2
		res.SetMapIndex(*index, *value)
	}

	return &res, cnt, nil
}

func (dv *DescValueMap) Read(spec base.ISpec, r io.Reader) (int, error) {
	var err error
	var cnt1, cnt2 int

	dv.index, cnt1, err = spec.ReadDesc(r)
	if err != nil {
		return 0, err
	}

	dv.value, cnt2, err = spec.ReadDesc(r)

	return cnt1 + cnt2, err
}

func (dv *DescValueMap) Write(spec base.ISpec, w io.Writer) (int, error) {
	cnt1, err := spec.WriteDesc(w, dv.index)
	if err != nil {
		return 0, err
	}

	cnt2, err := spec.WriteDesc(w, dv.value)
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2, nil
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
