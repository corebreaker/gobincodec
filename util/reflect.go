package util

import (
	"reflect"
	"unicode"

	"github.com/corebreaker/goerrors"
)

func IsFieldExported(field *reflect.StructField) bool {
	name := []rune(field.Name)

	return unicode.IsUpper(name[0])
}

func IsNil(value reflect.Value) bool {
	defer goerrors.DiscardPanic()

	return value.IsNil()
}
