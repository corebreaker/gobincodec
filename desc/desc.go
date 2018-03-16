package desc

import (
	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/desc/composed"
	"github.com/corebreaker/gobincodec/desc/simple"
)

const (
	DESC_NIL base.DescId = iota
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

type (
	pb   = simple.DescPrimitiveBool
	pi   = simple.DescPrimitiveInt
	pi8  = simple.DescPrimitiveInt8
	pi16 = simple.DescPrimitiveInt16
	pi32 = simple.DescPrimitiveInt32
	pi64 = simple.DescPrimitiveInt64
	pu   = simple.DescPrimitiveUint
	pu8  = simple.DescPrimitiveUint8
	pu16 = simple.DescPrimitiveUint16
	pu32 = simple.DescPrimitiveUint32
	pu64 = simple.DescPrimitiveUint64
	pp   = simple.DescPrimitiveUintptr
	pus  = simple.DescPrimitiveUnsafePtr
	pf32 = simple.DescPrimitiveFloat32
	pf64 = simple.DescPrimitiveFloat64
	ps   = simple.DescPrimitiveString
	pc32 = composed.DescSimpleComplex64
	pc64 = composed.DescSimpleComplex128
	pt   = composed.DescSimpleTime

	sb   = simple.DescSliceBool
	si   = simple.DescSliceInt
	si8  = simple.DescSliceInt8
	si16 = simple.DescSliceInt16
	si32 = simple.DescSliceInt32
	si64 = simple.DescSliceInt64
	su   = simple.DescSliceUint
	su8  = simple.DescSliceUint8
	su16 = simple.DescSliceUint16
	su32 = simple.DescSliceUint32
	su64 = simple.DescSliceUint64
	sp   = simple.DescSliceUintptr
	sus  = simple.DescSliceUnsafePtr
	sf32 = simple.DescSliceFloat32
	sf64 = simple.DescSliceFloat64
	ss   = simple.DescSliceString
	sc32 = composed.DescSliceComplex64
	sc64 = composed.DescSliceComplex128
	st   = composed.DescSliceTime

	ab   = simple.DescArrayBool
	ai   = simple.DescArrayInt
	ai8  = simple.DescArrayInt8
	ai16 = simple.DescArrayInt16
	ai32 = simple.DescArrayInt32
	ai64 = simple.DescArrayInt64
	au   = simple.DescArrayUint
	au8  = simple.DescArrayUint8
	au16 = simple.DescArrayUint16
	au32 = simple.DescArrayUint32
	au64 = simple.DescArrayUint64
	ap   = simple.DescArrayUintptr
	aus  = simple.DescArrayUnsafePtr
	af32 = simple.DescArrayFloat32
	af64 = simple.DescArrayFloat64
	as   = simple.DescArrayString
	ac32 = composed.DescArrayComplex64
	ac64 = composed.DescArrayComplex128
	at   = composed.DescArrayTime
)

func nb(id base.DescId) base.DescBase { return base.NewBase(id) }

var descMap = map[base.DescId]base.IDesc{
	DESC_SIMPLE_BOOL:       &pb{DescBase: nb(DESC_SIMPLE_BOOL)},
	DESC_SIMPLE_INT:        &pi{DescPrimitiveInt64: pi64{DescBase: nb(DESC_SIMPLE_INT)}},
	DESC_SIMPLE_INT8:       &pi8{DescBase: nb(DESC_SIMPLE_INT8)},
	DESC_SIMPLE_INT16:      &pi16{DescBase: nb(DESC_SIMPLE_INT16)},
	DESC_SIMPLE_INT32:      &pi32{DescBase: nb(DESC_SIMPLE_INT32)},
	DESC_SIMPLE_INT64:      &pi64{DescBase: nb(DESC_SIMPLE_INT64)},
	DESC_SIMPLE_UINT:       &pu{DescPrimitiveUint64: pu64{DescBase: nb(DESC_SIMPLE_UINT)}},
	DESC_SIMPLE_UINT8:      &pu8{DescBase: nb(DESC_SIMPLE_UINT8)},
	DESC_SIMPLE_UINT16:     &pu16{DescBase: nb(DESC_SIMPLE_UINT16)},
	DESC_SIMPLE_UINT32:     &pu32{DescBase: nb(DESC_SIMPLE_UINT32)},
	DESC_SIMPLE_UINT64:     &pu64{DescBase: nb(DESC_SIMPLE_UINT64)},
	DESC_SIMPLE_UINTPTR:    &pp{DescPrimitiveUint64: pu64{DescBase: nb(DESC_SIMPLE_UINTPTR)}},
	DESC_SIMPLE_UNSAFEPTR:  &pus{DescBase: nb(DESC_SIMPLE_UNSAFEPTR)},
	DESC_SIMPLE_FLOAT32:    &pf32{DescBase: nb(DESC_SIMPLE_FLOAT32)},
	DESC_SIMPLE_FLOAT64:    &pf64{DescBase: nb(DESC_SIMPLE_FLOAT64)},
	DESC_SIMPLE_STRING:     &ps{DescBase: nb(DESC_SIMPLE_STRING)},
	DESC_SIMPLE_COMPLEX64:  &pc32{DescBase: nb(DESC_SIMPLE_COMPLEX64)},
	DESC_SIMPLE_COMPLEX128: &pc64{DescBase: nb(DESC_SIMPLE_COMPLEX128)},
	DESC_SIMPLE_TIME:       &pt{DescBase: nb(DESC_SIMPLE_TIME)},
	DESC_SLICE_BOOL:        &sb{DescBase: nb(DESC_SLICE_BOOL)},
	DESC_SLICE_INT:         &si{DescSliceInt64: si64{DescBase: nb(DESC_SLICE_INT)}},
	DESC_SLICE_INT8:        &si8{DescBase: nb(DESC_SLICE_INT8)},
	DESC_SLICE_INT16:       &si16{DescBase: nb(DESC_SLICE_INT16)},
	DESC_SLICE_INT32:       &si32{DescBase: nb(DESC_SLICE_INT32)},
	DESC_SLICE_INT64:       &si64{DescBase: nb(DESC_SLICE_INT64)},
	DESC_SLICE_UINT:        &su{DescSliceUint64: su64{DescBase: nb(DESC_SLICE_UINT)}},
	DESC_SLICE_UINT8:       &su8{DescBase: nb(DESC_SLICE_UINT8)},
	DESC_SLICE_UINT16:      &su16{DescBase: nb(DESC_SLICE_UINT16)},
	DESC_SLICE_UINT32:      &su32{DescBase: nb(DESC_SLICE_UINT32)},
	DESC_SLICE_UINT64:      &su64{DescBase: nb(DESC_SLICE_UINT64)},
	DESC_SLICE_UINTPTR:     &sp{DescSliceUint64: su64{DescBase: nb(DESC_SLICE_UINTPTR)}},
	DESC_SLICE_UNSAFEPTR:   &sus{DescBase: nb(DESC_SLICE_UNSAFEPTR)},
	DESC_SLICE_FLOAT32:     &sf32{DescBase: nb(DESC_SLICE_FLOAT32)},
	DESC_SLICE_FLOAT64:     &sf64{DescBase: nb(DESC_SLICE_FLOAT64)},
	DESC_SLICE_STRING:      &ss{DescBase: nb(DESC_SLICE_STRING)},
	DESC_SLICE_COMPLEX64:   &sc32{DescBase: nb(DESC_SLICE_COMPLEX64)},
	DESC_SLICE_COMPLEX128:  &sc64{DescBase: nb(DESC_SLICE_COMPLEX128)},
	DESC_SLICE_TIME:        &st{DescBase: nb(DESC_SLICE_TIME)},
	DESC_ARRAY_BOOL:        &ab{DescSliceBool: sb{DescBase: nb(DESC_ARRAY_BOOL)}},
	DESC_ARRAY_INT:         &ai{DescSliceInt: si{DescSliceInt64: si64{DescBase: nb(DESC_ARRAY_INT)}}},
	DESC_ARRAY_INT8:        &ai8{DescSliceInt8: si8{DescBase: nb(DESC_ARRAY_INT8)}},
	DESC_ARRAY_INT16:       &ai16{DescSliceInt16: si16{DescBase: nb(DESC_ARRAY_INT16)}},
	DESC_ARRAY_INT32:       &ai32{DescSliceInt32: si32{DescBase: nb(DESC_ARRAY_INT32)}},
	DESC_ARRAY_INT64:       &ai64{DescSliceInt64: si64{DescBase: nb(DESC_ARRAY_INT64)}},
	DESC_ARRAY_UINT:        &au{DescSliceUint: su{DescSliceUint64: su64{DescBase: nb(DESC_ARRAY_UINT)}}},
	DESC_ARRAY_UINT8:       &au8{DescSliceUint8: su8{DescBase: nb(DESC_ARRAY_UINT8)}},
	DESC_ARRAY_UINT16:      &au16{DescSliceUint16: su16{DescBase: nb(DESC_ARRAY_UINT16)}},
	DESC_ARRAY_UINT32:      &au32{DescSliceUint32: su32{DescBase: nb(DESC_ARRAY_UINT32)}},
	DESC_ARRAY_UINT64:      &au64{DescSliceUint64: su64{DescBase: nb(DESC_ARRAY_UINT64)}},
	DESC_ARRAY_UINTPTR:     &ap{DescSliceUintptr: sp{DescSliceUint64: su64{DescBase: nb(DESC_ARRAY_UINTPTR)}}},
	DESC_ARRAY_UNSAFEPTR:   &aus{DescSliceUnsafePtr: sus{DescBase: nb(DESC_ARRAY_UNSAFEPTR)}},
	DESC_ARRAY_FLOAT32:     &af32{DescSliceFloat32: sf32{DescBase: nb(DESC_ARRAY_FLOAT32)}},
	DESC_ARRAY_FLOAT64:     &af64{DescSliceFloat64: sf64{DescBase: nb(DESC_ARRAY_FLOAT64)}},
	DESC_ARRAY_STRING:      &as{DescSliceString: ss{DescBase: nb(DESC_ARRAY_STRING)}},
	DESC_ARRAY_COMPLEX64:   &ac32{DescSliceComplex64: sc32{DescBase: nb(DESC_ARRAY_COMPLEX64)}},
	DESC_ARRAY_COMPLEX128:  &ac64{DescSliceComplex128: sc64{DescBase: nb(DESC_ARRAY_COMPLEX128)}},
	DESC_ARRAY_TIME:        &at{DescSliceTime: st{DescBase: nb(DESC_ARRAY_TIME)}},
}

func GetBaseDesc(id base.DescId) base.IDesc {
	return descMap[id]
}
