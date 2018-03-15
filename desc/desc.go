package desc

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/composed"
	"github.com/corebreaker/gobincodec/desc/simple"
)

type DescId byte

const (
	DESC_NIL DescId = iota
	DESC_SIMPLE_BOOL
	DESC_SIMPLE_INT
	DESC_SIMPLE_INT8
	DESC_SIMPLE_INT16
	DESC_SIMPLE_INT32
	DESC_SIMPLE_INT64
	DESC_SIMPLE_UINT
	DESC_SIMPLE_UINT8
	DESC_SIMPLE_UINT16
	DESC_SIMPLE_UINT32
	DESC_SIMPLE_UINT64
	DESC_SIMPLE_UINTPTR
	DESC_SIMPLE_UNSAFEPTR
	DESC_SIMPLE_FLOAT32
	DESC_SIMPLE_FLOAT64
	DESC_SIMPLE_COMPLEX64
	DESC_SIMPLE_COMPLEX128
	DESC_SIMPLE_STRING
	DESC_SIMPLE_TIME
	DESC_SLICE_BOOL
	DESC_SLICE_INT
	DESC_SLICE_INT8
	DESC_SLICE_INT16
	DESC_SLICE_INT32
	DESC_SLICE_INT64
	DESC_SLICE_UINT
	DESC_SLICE_UINT8
	DESC_SLICE_UINT16
	DESC_SLICE_UINT32
	DESC_SLICE_UINT64
	DESC_SLICE_UINTPTR
	DESC_SLICE_UNSAFEPTR
	DESC_SLICE_FLOAT32
	DESC_SLICE_FLOAT64
	DESC_SLICE_COMPLEX64
	DESC_SLICE_COMPLEX128
	DESC_SLICE_STRING
	DESC_SLICE_TIME
	DESC_ARRAY_BOOL
	DESC_ARRAY_INT
	DESC_ARRAY_INT8
	DESC_ARRAY_INT16
	DESC_ARRAY_INT32
	DESC_ARRAY_INT64
	DESC_ARRAY_UINT
	DESC_ARRAY_UINT8
	DESC_ARRAY_UINT16
	DESC_ARRAY_UINT32
	DESC_ARRAY_UINT64
	DESC_ARRAY_UINTPTR
	DESC_ARRAY_UNSAFEPTR
	DESC_ARRAY_FLOAT32
	DESC_ARRAY_FLOAT64
	DESC_ARRAY_COMPLEX64
	DESC_ARRAY_COMPLEX128
	DESC_ARRAY_STRING
	DESC_ARRAY_TIME
	DESC_OTHER
)

type IDesc interface {
	Encode(io.Writer, reflect.Value) error
	Decode(io.Reader) (*reflect.Value, error)

	Read(io.Reader) error
	Write(io.Writer) error
	Make(reflect.Type)
}

var descMap = map[DescId]IDesc{
	DESC_SIMPLE_BOOL:       new(simple.DescPrimitiveBool),
	DESC_SIMPLE_INT:        new(simple.DescPrimitiveInt),
	DESC_SIMPLE_INT8:       new(simple.DescPrimitiveInt8),
	DESC_SIMPLE_INT16:      new(simple.DescPrimitiveInt16),
	DESC_SIMPLE_INT32:      new(simple.DescPrimitiveInt32),
	DESC_SIMPLE_INT64:      new(simple.DescPrimitiveInt64),
	DESC_SIMPLE_UINT:       new(simple.DescPrimitiveUint),
	DESC_SIMPLE_UINT8:      new(simple.DescPrimitiveUint8),
	DESC_SIMPLE_UINT16:     new(simple.DescPrimitiveUint16),
	DESC_SIMPLE_UINT32:     new(simple.DescPrimitiveUint32),
	DESC_SIMPLE_UINT64:     new(simple.DescPrimitiveUint64),
	DESC_SIMPLE_UINTPTR:    new(simple.DescPrimitiveUintptr),
	DESC_SIMPLE_UNSAFEPTR:  new(simple.DescPrimitiveUnsafePtr),
	DESC_SIMPLE_FLOAT32:    new(simple.DescPrimitiveFloat32),
	DESC_SIMPLE_FLOAT64:    new(simple.DescPrimitiveFloat64),
	DESC_SIMPLE_STRING:     new(simple.DescPrimitiveString),
	DESC_SIMPLE_COMPLEX64:  new(composed.DescSimpleComplex64),
	DESC_SIMPLE_COMPLEX128: new(composed.DescSimpleComplex128),
	DESC_SIMPLE_TIME:       new(composed.DescSimpleTime),
	DESC_SLICE_BOOL:        new(simple.DescSliceBool),
	DESC_SLICE_INT:         new(simple.DescSliceInt),
	DESC_SLICE_INT8:        new(simple.DescSliceInt8),
	DESC_SLICE_INT16:       new(simple.DescSliceInt16),
	DESC_SLICE_INT32:       new(simple.DescSliceInt32),
	DESC_SLICE_INT64:       new(simple.DescSliceInt64),
	DESC_SLICE_UINT:        new(simple.DescSliceUint),
	DESC_SLICE_UINT8:       new(simple.DescSliceUint8),
	DESC_SLICE_UINT16:      new(simple.DescSliceUint16),
	DESC_SLICE_UINT32:      new(simple.DescSliceUint32),
	DESC_SLICE_UINT64:      new(simple.DescSliceUint64),
	DESC_SLICE_UINTPTR:     new(simple.DescSliceUintptr),
	DESC_SLICE_UNSAFEPTR:   new(simple.DescSliceUnsafePtr),
	DESC_SLICE_FLOAT32:     new(simple.DescSliceFloat32),
	DESC_SLICE_FLOAT64:     new(simple.DescSliceFloat64),
	DESC_SLICE_STRING:      new(simple.DescSliceString),
	DESC_SLICE_COMPLEX64:   new(composed.DescSliceComplex64),
	DESC_SLICE_COMPLEX128:  new(composed.DescSliceComplex128),
	DESC_SLICE_TIME:        new(composed.DescSliceTime),
	DESC_ARRAY_BOOL:        new(simple.DescArrayBool),
	DESC_ARRAY_INT:         new(simple.DescArrayInt),
	DESC_ARRAY_INT8:        new(simple.DescArrayInt8),
	DESC_ARRAY_INT16:       new(simple.DescArrayInt16),
	DESC_ARRAY_INT32:       new(simple.DescArrayInt32),
	DESC_ARRAY_INT64:       new(simple.DescArrayInt64),
	DESC_ARRAY_UINT:        new(simple.DescArrayUint),
	DESC_ARRAY_UINT8:       new(simple.DescArrayUint8),
	DESC_ARRAY_UINT16:      new(simple.DescArrayUint16),
	DESC_ARRAY_UINT32:      new(simple.DescArrayUint32),
	DESC_ARRAY_UINT64:      new(simple.DescArrayUint64),
	DESC_ARRAY_UINTPTR:     new(simple.DescArrayUintptr),
	DESC_ARRAY_UNSAFEPTR:   new(simple.DescArrayUnsafePtr),
	DESC_ARRAY_FLOAT32:     new(simple.DescArrayFloat32),
	DESC_ARRAY_FLOAT64:     new(simple.DescArrayFloat64),
	DESC_ARRAY_STRING:      new(simple.DescArrayString),
	DESC_ARRAY_COMPLEX64:   new(composed.DescArrayComplex64),
	DESC_ARRAY_COMPLEX128:  new(composed.DescArrayComplex128),
	DESC_ARRAY_TIME:        new(composed.DescArrayTime),
}

func GetBaseDesc(id DescId) IDesc {
	return descMap[id]
}
