package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
)

type DescValueArray struct {
	value base.IDesc
}

func (dv *DescValueArray) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	return nil
}

func (dv *DescValueArray) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	return nil, nil
}

func (dv *DescValueArray) Read(spec base.ISpec, r io.Reader) error {
	return nil
}

func (dv *DescValueArray) Write(spec base.ISpec, w io.Writer) error {
	return nil
}

func (dv *DescValueArray) Make(spec base.ISpec, t reflect.Type) error {
	return nil
}
