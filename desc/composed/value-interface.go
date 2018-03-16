package composed

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
)

type DescValueInterface struct{ base.DescBase }

func (*DescValueInterface) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	return nil
}

func (*DescValueInterface) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	return nil, nil
}
