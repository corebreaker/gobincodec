package base

import (
	"io"
	"reflect"
)

type DescId uint16

type IDesc interface {
	GetId() DescId

	Encode(ISpec, io.Writer, reflect.Value) error
	Decode(ISpec, io.Reader) (*reflect.Value, error)

	Read(ISpec, io.Reader) error
	Write(ISpec, io.Writer) error
	Make(ISpec, reflect.Type) error
}
