package base

import (
	"reflect"
)

type ISpec interface {
	DescFromType(reflect.Type) IDesc
	DescFromId(DescId) IDesc
}
