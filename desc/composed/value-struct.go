package composed

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescValueStruct struct {
	base.DescBase

	names  []string
	fields map[string]base.IDesc
}

func (dv *DescValueStruct) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	var out bytes.Buffer

	for _, name := range dv.names {
		if _, err := dv.fields[name].Encode(spec, &out, v.FieldByName(name)); err != nil {
			return 0, err
		}
	}

	return util.Write(w, out.Bytes())
}

func (dv *DescValueStruct) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	cnt := 0

	res := reflect.New(spec.GetType(dv)).Elem()
	for _, name := range dv.names {
		field, n, err := dv.fields[name].Decode(spec, r)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		res.FieldByName(name).Set(*field)
	}

	return &res, cnt, nil
}

func (dv *DescValueStruct) Read(spec base.ISpec, r io.Reader) (int, error) {
	size, cnt, err := util.DecodeSize(r)
	if err != nil {
		return 0, err
	}

	for i := 0; i < size; i++ {
		strlen, n1, err := util.DecodeSize(r)
		if err != nil {
			return 0, err
		}

		name_buf := make([]byte, strlen)
		n2, err := util.Read(r, name_buf)
		if err != nil {
			return 0, err
		}

		name := string(name_buf)
		ftype, n3, err := spec.ReadDesc(r)
		if err != nil {
			return 0, err
		}

		cnt += n1 + n2 + n3
		dv.names = append(dv.names, name)
		dv.fields[name] = ftype
	}

	return cnt, nil
}

func (dv *DescValueStruct) Write(spec base.ISpec, w io.Writer) (int, error) {
	var out bytes.Buffer

	if _, err := util.EncodeSize(&out, len(dv.names)); err != nil {
		return 0, err
	}

	for _, name := range dv.names {
		name_buf := []byte(name)
		if _, err := util.EncodeSize(&out, len(name_buf)); err != nil {
			return 0, err
		}

		if _, err := util.Write(&out, name_buf); err != nil {
			return 0, err
		}

		if _, err := spec.WriteDesc(&out, dv.fields[name]); err != nil {
			return 0, err
		}
	}

	return util.Write(w, out.Bytes())
}

func (dv *DescValueStruct) Make(spec base.ISpec, t reflect.Type) error {
	count := t.NumField()

	for i := 0; i < count; i++ {
		field := t.Field(i)

		if !util.IsFieldExported(&field) {
			continue
		}

		dv.names = append(dv.names, field.Name)
		dv.fields[field.Name] = spec.DescFromType(field.Type)
	}

	return nil
}

func NewStructDesc(id base.DescId) base.IDesc {
	return &DescValueStruct{
		DescBase: base.NewBase(id),
		fields:   make(map[string]base.IDesc),
	}
}
