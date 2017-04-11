package bincodec

import (
    "reflect"
)

type iCodec interface {
    EncodeValue(Encoder, interface{}) error
    DecodeValue(Decoder) (reflect.Value, error)
}

type tTypeId byte

const (
    _TID_NIL tTypeId = iota
    _TID_BOOL
    _TID_INT
    _TID_INT8
    _TID_INT16
    _TID_INT32
    _TID_INT64
    _TID_UINT
    _TID_UINT8
    _TID_UINT16
    _TID_UINT32
    _TID_UINT64
    _TID_UINTPTR
    _TID_UNSAFEPTR
    _TID_FLOAT32
    _TID_FLOAT64
    _TID_COMPLEX64
    _TID_COMPLEX128
    _TID_STRING
    _TID_ARRAY
    _TID_SLICE
    _TID_MAP
    _TID_PTR
    _TID_TIME
    _TID_STRUCT
    _TID_INTERFACE
    _TID_SERIALIZABLE
)

var (
    base_codecs = map[reflect.Kind]iCodec{
        reflect.Bool:          tCodecBool(_TID_BOOL),
        reflect.Int:           tCodecInt(_TID_INT),
        reflect.Int8:          tCodecInt8(_TID_INT8),
        reflect.Int16:         tCodecInt16(_TID_INT16),
        reflect.Int32:         tCodecInt32(_TID_INT32),
        reflect.Int64:         tCodecInt64(_TID_INT64),
        reflect.Uint:          tCodecUint(_TID_UINT),
        reflect.Uint8:         tCodecUint8(_TID_UINT8),
        reflect.Uint16:        tCodecUint16(_TID_UINT16),
        reflect.Uint32:        tCodecUint32(_TID_UINT32),
        reflect.Uint64:        tCodecUint64(_TID_UINT64),
        reflect.Uintptr:       tCodecUintptr(_TID_UINTPTR),
        reflect.UnsafePointer: tCodecUnsafePtr(_TID_UNSAFEPTR),
        reflect.Float32:       tCodecFloat32(_TID_FLOAT32),
        reflect.Float64:       tCodecFloat64(_TID_FLOAT64),
        reflect.Complex64:     tCodecComplex64(_TID_COMPLEX64),
        reflect.Complex128:    tCodecComplex128(_TID_COMPLEX128),
        reflect.String:        tCodecString(_TID_STRING),
    }
)

func (self *tCodecBase) compile_type(typ reflect.Type) iCodec {
    return nil
}
