package base

import (
	"io"
	"reflect"
)

type DescBase struct {
	id DescId
}

func (d *DescBase) GetId() DescId                     { return d.id }
func (*DescBase) IsNil() bool                         { return false }
func (*DescBase) Read(ISpec, io.Reader) (int, error)  { return 0, nil }
func (*DescBase) Write(ISpec, io.Writer) (int, error) { return 0, nil }
func (*DescBase) Make(ISpec, reflect.Type) error      { return nil }

func NewBase(id DescId) DescBase {
	return DescBase{
		id: id,
	}
}
