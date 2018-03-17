package composed

import (
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

func (dv *DescValueStruct) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	for _, name := range dv.names {
		if err := dv.fields[name].Encode(spec, w, v.FieldByName(name)); err != nil {
			return err
		}
	}

	return nil
}

func (dv *DescValueStruct) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	res := reflect.New(spec.GetType(dv)).Elem()
	for _, name := range dv.names {
		field, err := dv.fields[name].Decode(spec, r)
		if err != nil {
			return nil, err
		}

		res.FieldByName(name).Set(*field)
	}

	return &res, nil
}

func (dv *DescValueStruct) Read(spec base.ISpec, r io.Reader) error {
	count, err := util.DecodeSize(r)
	if err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		strlen, err := util.DecodeSize(r)
		if err != nil {
			return err
		}

		name_buf := make([]byte, strlen)
		if err := util.Read(r, name_buf); err != nil {
			return err
		}

		name := string(name_buf)
		ftype, err := spec.ReadDesc(r)
		if err != nil {
			return err
		}

		dv.names = append(dv.names, name)
		dv.fields[name] = ftype
	}

	return nil
}

func (dv *DescValueStruct) Write(spec base.ISpec, w io.Writer) error {
	if err := util.EncodeSize(w, len(dv.names)); err != nil {
		return err
	}

	for _, name := range dv.names {
		name_buf := []byte(name)
		if err := util.EncodeSize(w, len(name_buf)); err != nil {
			return err
		}

		if err := util.Write(w, name_buf); err != nil {
			return err
		}

		if err := spec.WriteDesc(w, dv.fields[name]); err != nil {
			return err
		}
	}

	return nil
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
