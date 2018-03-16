package base

import (
	"io"
	"reflect"
)

type DescBase struct {
	id DescId
}

func (d *DescBase) GetId() DescId                { return d.id }
func (*DescBase) Read(ISpec, io.Reader) error    { return nil }
func (*DescBase) Write(ISpec, io.Writer) error   { return nil }
func (*DescBase) Make(ISpec, reflect.Type) error { return nil }

func NewBase(id DescId) DescBase {
	return DescBase{
		id: id,
	}
}
