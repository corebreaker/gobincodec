package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
)

type DescValuePointer struct {
	elem base.IDesc
}

func (dv *DescValuePointer) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	return nil
}

func (dv *DescValuePointer) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	return nil, nil
}

func (dv *DescValuePointer) Read(spec base.ISpec, r io.Reader) error {
	return nil
}

func (dv *DescValuePointer) Write(spec base.ISpec, w io.Writer) error {
	return nil
}

func (dv *DescValuePointer) Make(spec base.ISpec, t reflect.Type) error {
	return nil
}
