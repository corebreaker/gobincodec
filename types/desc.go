package types

import (
	"io"
	"reflect"
)

type TypeId byte

type IDesc interface {
	Encode(Encoder, interface{}) error
	Decode(Decoder) (reflect.Value, error)

	Read(io.Reader) error
	Write(io.Writer) error
	Make(reflect.Type)
}
