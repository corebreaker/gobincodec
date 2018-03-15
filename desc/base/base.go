package base

import (
	"io"
	"reflect"
)

type DescBase struct{}

func (*DescBase) Read(io.Reader) error  { return nil }
func (*DescBase) Write(io.Writer) error { return nil }
func (*DescBase) Make(reflect.Type)     {}
