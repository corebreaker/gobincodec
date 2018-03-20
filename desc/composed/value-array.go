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

func (dv *DescValueArray) TypeEquals(reflect.Type) bool {

}

func (dv *DescValueArray) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (dv *DescValueArray) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	if dv.count == 0 {
		return 0, nil
	}

	var out bytes.Buffer

	for i := 0; i < dv.count; i++ {
		if _, err := dv.value.Encode(spec, &out, v.Index(i)); err != nil {
			return 0, err
		}
	}

	return util.Write(w, out.Bytes())
}

func (dv *DescValueArray) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	res := reflect.New(spec.GetType(dv)).Elem()
	cnt := 0

	for i := 0; i < dv.count; i++ {
		value, n, err := dv.value.Decode(spec, r)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		res.Field(i).Set(*value)
	}

	return &res, cnt, nil
}

func (dv *DescValueArray) Read(spec base.ISpec, r io.Reader) (int, error) {
	var err error
	var cnt1, cnt2 int

	dv.count, cnt1, err = util.DecodeSize(r)
	if err != nil {
		return 0, err
	}

	dv.value, cnt2, err = spec.ReadDesc(r)

	return cnt1 + cnt2, err
}

func (dv *DescValueArray) Write(spec base.ISpec, w io.Writer) (int, error) {
	cnt1, err := util.EncodeSize(w, dv.count)
	if err != nil {
		return 0, err
	}

	cnt2, err := spec.WriteDesc(w, dv.value)
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2, nil
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
