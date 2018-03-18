package base

import (
	"io"
	"reflect"
)

type ISpec interface {
	ReadDesc(io.Reader) (IDesc, int, error)
	WriteDesc(io.Writer, IDesc) (int, error)

	DescFromType(reflect.Type) IDesc
	DescFromId(DescId) IDesc
	GetType(IDesc) reflect.Type
}
