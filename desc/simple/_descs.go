package simple

import (
	"github.com/corebreaker/gobincodec/desc/base"
)

var (
	boolCodec      = &DescPrimitiveBool{DescBase: base.NewBase(DESC_SIMPLE_BOOL)}
	intCodec       = &DescPrimitiveInt{DescPrimitiveInt64: pi64{DescBase: base.NewBase(DESC_SIMPLE_INT)}}
	int8Codec      = &DescPrimitiveInt8{DescBase: base.NewBase(DESC_SIMPLE_INT8)}
	int16Codec     = &DescPrimitiveInt16{DescBase: base.NewBase(DESC_SIMPLE_INT16)}
	int32Codec     = &DescPrimitiveInt32{DescBase: base.NewBase(DESC_SIMPLE_INT32)}
	int64Codec     = &DescPrimitiveInt64{DescBase: base.NewBase(DESC_SIMPLE_INT64)}
	uintCodec      = &DescPrimitiveUint{DescPrimitiveUint64: pu64{DescBase: base.NewBase(DESC_SIMPLE_UINT)}}
	uint8Codec     = &DescPrimitiveUint8{DescBase: base.NewBase(DESC_SIMPLE_UINT8)}
	uint16Codec    = &DescPrimitiveUint16{DescBase: base.NewBase(DESC_SIMPLE_UINT16)}
	uint32Codec    = &DescPrimitiveUint32{DescBase: base.NewBase(DESC_SIMPLE_UINT32)}
	uint64Codec    = &DescPrimitiveUint64{DescBase: base.NewBase(DESC_SIMPLE_UINT64)}
	uintPtrCodec   = &DescPrimitiveUintptr{DescPrimitiveUint64: pu64{DescBase: base.NewBase(DESC_SIMPLE_UINTPTR)}}
	unsafePtrCodec = &DescPrimitiveUnsafePtr{DescBase: base.NewBase(DESC_SIMPLE_UNSAFEPTR)}
	float32Codec   = &DescPrimitiveFloat32{DescBase: base.NewBase(DESC_SIMPLE_FLOAT32)}
	float64Codec   = &DescPrimitiveFloat64{DescBase: base.NewBase(DESC_SIMPLE_FLOAT64)}
	strCodec       = &DescPrimitiveString{DescBase: base.NewBase(DESC_SIMPLE_STRING)}
)
