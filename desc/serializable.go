package base

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
)

type tDescSerializable struct {
	base.DescBase
}

func (d *tDescSerializable) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	return 0, nil
}

func (d *tDescSerializable) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	return nil, 0, nil
}

func (d *tDescSerializable) Read(spec base.ISpec, r io.Reader) (int, error) {
	return 0, nil
}

func (d *tDescSerializable) Write(spec base.ISpec, w io.Writer) (int, error) {
	return 0, nil
}

func (d *tDescSerializable) Make(spec base.ISpec, t reflect.Type) error {
	return nil
}
