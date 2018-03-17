package base

import (
	"io"
	"reflect"
)

type ISpec interface {
	ReadDesc(io.Reader) (IDesc, error)
	WriteDesc(io.Writer, IDesc) error

	DescFromType(reflect.Type) IDesc
	DescFromId(DescId) IDesc
	GetType(IDesc) reflect.Type
}
