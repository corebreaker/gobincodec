package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
)

type DescValueSlice struct {
	value base.IDesc
}

func (dv *DescValueSlice) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	return nil
}

func (dv *DescValueSlice) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	return nil, nil
}

func (dv *DescValueSlice) Read(spec base.ISpec, r io.Reader) error {
	return nil
}

func (dv *DescValueSlice) Write(spec base.ISpec, w io.Writer) error {
	return nil
}

func (dv *DescValueSlice) Make(spec base.ISpec, t reflect.Type) error {
	return nil
}
