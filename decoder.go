package bincodec

import (
    "io"
    "reflect"
    "time"
    "unsafe"
)

type Decoder interface {
    Decode(...interface{}) error
    DecodeValue(reflect.Value) error
    DecodeBool(*bool) error
    DecodeByte(*byte) error
    DecodeInt(*int) error
    DecodeInt8(*int8) error
    DecodeInt16(*int16) error
    DecodeInt32(*int32) error
    DecodeInt64(*int64) error
    DecodeUint(*uint) error
    DecodeUint8(*uint8) error
    DecodeUint16(*uint16) error
    DecodeUint32(*uint32) error
    DecodeUint64(*uint64) error
    DecodeUintptr(*uintptr) error
    DecodeUnsafePtr(*unsafe.Pointer) error
    DecodeFloat32(*float32) error
    DecodeFloat64(*float64) error
    DecodeComplex64(*complex64) error
    DecodeComplex128(*complex128) error
    DecodeString(*string) error
    DecodeTime(*time.Time) error
    DecodeBoolSlice(*[]bool) error
    DecodeByteSlice(*[]byte) error
    DecodeIntSlice(*[]int) error
    DecodeInt8Slice(*[]int8) error
    DecodeInt16Slice(*[]int16) error
    DecodeInt32Slice(*[]int32) error
    DecodeInt64Slice(*[]int64) error
    DecodeUintSlice(*[]uint) error
    DecodeUint8Slice(*[]uint8) error
    DecodeUint16Slice(*[]uint16) error
    DecodeUint32Slice(*[]uint32) error
    DecodeUint64Slice(*[]uint64) error
    DecodeUintptrSlice(*[]uintptr) error
    DecodeFloat32Slice(*[]float32) error
    DecodeFloat64Slice(*[]float64) error
    DecodeComplex64Slice(*[]complex64) error
    DecodeComplex128Slice(*[]complex128) error
    DecodeStringSlice(*[]string) error
    DecodeSlice(*[]interface{}) error
    DecodeSerializable(Serializable) error
}

type Reader interface {
    CodecBase
    IoBase
    Decoder

    CloneReader(io.Reader) Reader
    ReadHeader() error
}

type ReaderAt interface {
    CodecBase
    IoBase

    ReadHeaderAtBegin() error
    ReadHeaderAt(int64) error
    CloneReaderAt(io.ReaderAt) ReaderAt

    DecodeAt(interface{}, int64) error
    DecodeValueAt(reflect.Value, int64) error
    DecodeBoolAt(*bool, int64) error
    DecodeByteAt(*byte, int64) error
    DecodeIntAt(*int, int64) error
    DecodeInt8At(*int8, int64) error
    DecodeInt16At(*int16, int64) error
    DecodeInt32At(*int32, int64) error
    DecodeInt64At(*int64, int64) error
    DecodeUintAt(*uint, int64) error
    DecodeUint8At(*uint8, int64) error
    DecodeUint16At(*uint16, int64) error
    DecodeUint32At(*uint32, int64) error
    DecodeUint64At(*uint64, int64) error
    DecodeUintptrAt(*uintptr, int64) error
    DecodeUnsafePtrAt(*unsafe.Pointer, int64) error
    DecodeFloat32At(*float32, int64) error
    DecodeFloat64At(*float64, int64) error
    DecodeComplex64At(*complex64, int64) error
    DecodeComplex128At(*complex128, int64) error
    DecodeStringAt(*string, int64) error
    DecodeTimeAt(*time.Time, int64) error
    DecodeBoolSliceAt(*[]bool, int64) error
    DecodeByteSliceAt(*[]byte, int64) error
    DecodeIntSliceAt(*[]int, int64) error
    DecodeInt8SliceAt(*[]int8, int64) error
    DecodeInt16SliceAt(*[]int16, int64) error
    DecodeInt32SliceAt(*[]int32, int64) error
    DecodeInt64SliceAt(*[]int64, int64) error
    DecodeUintSliceAt(*[]uint, int64) error
    DecodeUint8SliceAt(*[]uint8, int64) error
    DecodeUint16SliceAt(*[]uint16, int64) error
    DecodeUint32SliceAt(*[]uint32, int64) error
    DecodeUint64SliceAt(*[]uint64, int64) error
    DecodeUintptrSliceAt(*[]uintptr, int64) error
    DecodeFloat32SliceAt(*[]float32, int64) error
    DecodeFloat64SliceAt(*[]float64, int64) error
    DecodeComplex64SliceAt(*[]complex64, int64) error
    DecodeComplex128SliceAt(*[]complex128, int64) error
    DecodeStringSliceAt(*[]string, int64) error
    DecodeSliceAt(*[]interface{}, int64) error
    DecodeSerializableAt(Serializable, int64) error
}

type tReader struct {
    tIoBase

    r   io.Reader
}

func (self *tReader) CloneReader(r io.Reader) Reader {
    return &tReader{
        tIoBase: tIoBase{
            tCodecBase: tCodecBase{
                spec: self.spec.clone()
            },

            r: r,
        },
    }
}

func (self *tReader) ReadHeader() error {
    return self.spec.read(self.r)
}

type tReaderAt struct {
    tIoBase

    r   io.ReaderAt
}

func (self *tReaderAt) ReadHeaderAtBegin() error {
    return self.spec.read_at(self.r, 0)
}

func (self *tReaderAt) ReadHeaderAt(offset int64) error {
    return self.spec.read_at(self.r, offset)
}

func (self *tReaderAt) CloneReaderAt(r io.ReaderAt) ReaderAt {
    return &tReaderAt{
        tIoBase: tIoBase{
            tCodecBase: tCodecBase{
                spec: self.spec.clone()
            },

            r: r,
        },
    }
}
