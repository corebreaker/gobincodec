package gobincodec

import (
	"reflect"
	"time"

	"github.com/corebreaker/gobincodec/desc"
	"github.com/corebreaker/gobincodec/desc/base"
)

type tCodecKind byte

const (
	_CODEC_KIND_SIMPLE tCodecKind = iota
	_CODEC_KIND_SLICE
	_CODEC_KIND_ARRAY
)

type tCodecRef struct {
	timeCodec    base.IDesc
	kindedCodecs map[reflect.Kind]base.IDesc
}

var (
	timeType = reflect.TypeOf((*time.Time)(nil)).Elem()

	baseCodecs = []tCodecRef{
		tCodecRef{
			timeCodec: desc.GetBaseDesc(desc.DESC_SIMPLE_TIME),
			kindedCodecs: map[reflect.Kind]base.IDesc{
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
			},
		},

		tCodecRef{
			timeCodec: desc.GetBaseDesc(desc.DESC_SLICE_TIME),
			kindedCodecs: map[reflect.Kind]base.IDesc{
				reflect.Bool:          desc.GetBaseDesc(desc.DESC_SLICE_BOOL),
				reflect.Int:           desc.GetBaseDesc(desc.DESC_SLICE_INT),
				reflect.Int8:          desc.GetBaseDesc(desc.DESC_SLICE_INT8),
				reflect.Int16:         desc.GetBaseDesc(desc.DESC_SLICE_INT16),
				reflect.Int32:         desc.GetBaseDesc(desc.DESC_SLICE_INT32),
				reflect.Int64:         desc.GetBaseDesc(desc.DESC_SLICE_INT64),
				reflect.Uint:          desc.GetBaseDesc(desc.DESC_SLICE_UINT),
				reflect.Uint8:         desc.GetBaseDesc(desc.DESC_SLICE_UINT8),
				reflect.Uint16:        desc.GetBaseDesc(desc.DESC_SLICE_UINT16),
				reflect.Uint32:        desc.GetBaseDesc(desc.DESC_SLICE_UINT32),
				reflect.Uint64:        desc.GetBaseDesc(desc.DESC_SLICE_UINT64),
				reflect.Uintptr:       desc.GetBaseDesc(desc.DESC_SLICE_UINTPTR),
				reflect.UnsafePointer: desc.GetBaseDesc(desc.DESC_SLICE_UNSAFEPTR),
				reflect.Float32:       desc.GetBaseDesc(desc.DESC_SLICE_FLOAT32),
				reflect.Float64:       desc.GetBaseDesc(desc.DESC_SLICE_FLOAT64),
				reflect.Complex64:     desc.GetBaseDesc(desc.DESC_SLICE_COMPLEX64),
				reflect.Complex128:    desc.GetBaseDesc(desc.DESC_SLICE_COMPLEX128),
				reflect.String:        desc.GetBaseDesc(desc.DESC_SLICE_STRING),
			},
		},

		tCodecRef{
			timeCodec: desc.GetBaseDesc(desc.DESC_ARRAY_TIME),
			kindedCodecs: map[reflect.Kind]base.IDesc{
				reflect.Bool:          desc.GetBaseDesc(desc.DESC_ARRAY_BOOL),
				reflect.Int:           desc.GetBaseDesc(desc.DESC_ARRAY_INT),
				reflect.Int8:          desc.GetBaseDesc(desc.DESC_ARRAY_INT8),
				reflect.Int16:         desc.GetBaseDesc(desc.DESC_ARRAY_INT16),
				reflect.Int32:         desc.GetBaseDesc(desc.DESC_ARRAY_INT32),
				reflect.Int64:         desc.GetBaseDesc(desc.DESC_ARRAY_INT64),
				reflect.Uint:          desc.GetBaseDesc(desc.DESC_ARRAY_UINT),
				reflect.Uint8:         desc.GetBaseDesc(desc.DESC_ARRAY_UINT8),
				reflect.Uint16:        desc.GetBaseDesc(desc.DESC_ARRAY_UINT16),
				reflect.Uint32:        desc.GetBaseDesc(desc.DESC_ARRAY_UINT32),
				reflect.Uint64:        desc.GetBaseDesc(desc.DESC_ARRAY_UINT64),
				reflect.Uintptr:       desc.GetBaseDesc(desc.DESC_ARRAY_UINTPTR),
				reflect.UnsafePointer: desc.GetBaseDesc(desc.DESC_ARRAY_UNSAFEPTR),
				reflect.Float32:       desc.GetBaseDesc(desc.DESC_ARRAY_FLOAT32),
				reflect.Float64:       desc.GetBaseDesc(desc.DESC_ARRAY_FLOAT64),
				reflect.Complex64:     desc.GetBaseDesc(desc.DESC_ARRAY_COMPLEX64),
				reflect.Complex128:    desc.GetBaseDesc(desc.DESC_ARRAY_COMPLEX128),
				reflect.String:        desc.GetBaseDesc(desc.DESC_ARRAY_STRING),
			},
		},
	}
)

func (self *tCodecBase) compileType(typ reflect.Type) base.IDesc {
	return nil
}
