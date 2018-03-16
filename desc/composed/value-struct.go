package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
)

type DescValueStruct struct {
	fields map[string]base.IDesc
}

func (dv *DescValueStruct) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	return nil
}

func (dv *DescValueStruct) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	return nil, nil
}

func (dv *DescValueStruct) Read(spec base.ISpec, r io.Reader) error {
	return nil
}

func (dv *DescValueStruct) Write(spec base.ISpec, w io.Writer) error {
	return nil
}

func (dv *DescValueStruct) Make(spec base.ISpec, t reflect.Type) error {
	return nil
}
