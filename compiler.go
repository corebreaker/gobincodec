package gobincodec

import (
	"reflect"

	"github.com/corebreaker/gobincodec/desc"
	"github.com/corebreaker/gobincodec/desc/base"
)

var (
	base_codecs = map[reflect.Kind]base.IDesc{
		reflect.Bool:          desc.GetBaseDesc(desc.DESC_SIMPLE_BOOL),
		reflect.Int:           desc.GetBaseDesc(desc.DESC_SIMPLE_INT),
		reflect.Int8:          desc.GetBaseDesc(desc.DESC_SIMPLE_INT8),
		reflect.Int16:         desc.GetBaseDesc(desc.DESC_SIMPLE_INT16),
		reflect.Int32:         desc.GetBaseDesc(desc.DESC_SIMPLE_INT32),
		reflect.Int64:         desc.GetBaseDesc(desc.DESC_SIMPLE_INT64),
		reflect.Uint:          desc.GetBaseDesc(desc.DESC_SIMPLE_UINT),
		reflect.Uint8:         desc.GetBaseDesc(desc.DESC_SIMPLE_UINT8),
		reflect.Uint16:        desc.GetBaseDesc(desc.DESC_SIMPLE_UINT16),
		reflect.Uint32:        desc.GetBaseDesc(desc.DESC_SIMPLE_UINT32),
		reflect.Uint64:        desc.GetBaseDesc(desc.DESC_SIMPLE_UINT64),
		reflect.Uintptr:       desc.GetBaseDesc(desc.DESC_SIMPLE_UINTPTR),
		reflect.UnsafePointer: desc.GetBaseDesc(desc.DESC_SIMPLE_UNSAFEPTR),
		reflect.Float32:       desc.GetBaseDesc(desc.DESC_SIMPLE_FLOAT32),
		reflect.Float64:       desc.GetBaseDesc(desc.DESC_SIMPLE_FLOAT64),
		reflect.Complex64:     desc.GetBaseDesc(desc.DESC_SIMPLE_COMPLEX64),
		reflect.Complex128:    desc.GetBaseDesc(desc.DESC_SIMPLE_COMPLEX128),
		reflect.String:        desc.GetBaseDesc(desc.DESC_SIMPLE_STRING),
	}
)

func (self *tCodecBase) compileType(typ reflect.Type) base.IDesc {
	return nil
}
