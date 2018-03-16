package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
)

type DescValueMap struct {
	index base.IDesc
	value base.IDesc
}

func (dv *DescValueMap) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	return nil
}

func (dv *DescValueMap) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	return nil, nil
}

func (dv *DescValueMap) Read(spec base.ISpec, r io.Reader) error {
	return nil
}

func (dv *DescValueMap) Write(spec base.ISpec, w io.Writer) error {
	return nil
}

func (dv *DescValueMap) Make(spec base.ISpec, t reflect.Type) error {
	return nil
}
