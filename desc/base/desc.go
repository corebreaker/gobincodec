package base

import (
	"io"
	"reflect"
)

type DescId uint16

type IDesc interface {
	GetId() DescId
	IsNil() bool
	TypeEquals(reflect.Type) bool
	Convert(reflect.Value, reflect.Type) *reflect.Value

	Encode(ISpec, io.Writer, reflect.Value) (int, error)
	Decode(ISpec, io.Reader) (*reflect.Value, int, error)

	Read(ISpec, io.Reader) (int, error)
	Write(ISpec, io.Writer) (int, error)
	Make(ISpec, reflect.Type) error
}
