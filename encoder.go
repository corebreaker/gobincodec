package bincodec

import (
    "io"
    "reflect"
    "time"
    "unsafe"
)

type Encoder interface {
    Encode(...interface{}) error
    EncodeValue(reflect.Value) error
    EncodeBool(bool) error
    EncodeByte(byte) error
    EncodeInt(int) error
    EncodeInt8(int8) error
    EncodeInt16(int16) error
    EncodeInt32(int32) error
    EncodeInt64(int64) error
    EncodeUint(uint) error
    EncodeUint8(uint8) error
    EncodeUint16(uint16) error
    EncodeUint32(uint32) error
    EncodeUint64(uint64) error
    EncodeUintptr(uintptr) error
    EncodeUnsafePtr(unsafe.Pointer) error
    EncodeFloat32(float32) error
    EncodeFloat64(float64) error
    EncodeComplex64(complex64) error
    EncodeComplex128(complex128) error
    EncodeString(string) error
    EncodeTime(time.Time) error
    EncodeBoolSlice([]bool) error
    EncodeByteSlice([]byte) error
    EncodeIntSlice([]int) error
    EncodeInt8Slice([]int8) error
    EncodeInt16Slice([]int16) error
    EncodeInt32Slice([]int32) error
    EncodeInt64Slice([]int64) error
    EncodeUintSlice([]uint) error
    EncodeUint8Slice([]uint8) error
    EncodeUint16Slice([]uint16) error
    EncodeUint32Slice([]uint32) error
    EncodeUint64Slice([]uint64) error
    EncodeUintptrSlice([]uintptr) error
    EncodeFloat32Slice([]float32) error
    EncodeFloat64Slice([]float64) error
    EncodeComplex64Slice([]complex64) error
    EncodeComplex128Slice([]complex128) error
    EncodeStringSlice([]string) error
    EncodeSlice([]interface{}) error
    EncodeSerializable(Serializable) error
}

type Writer interface {
    CodecBase
    IoBase
    Encoder

    CloneWriter(io.Writer) Writer
    WriteHeader() error
}

type WriterAt interface {
    CodecBase
    IoBase

    WriteHeaderAtBegin() error
    WriteHeaderAt(int64) error
    CloneWriterAt(io.WriterAt) WriterAt

    EncodeAt(interface{}, int64) error
    EncodeValueAt(reflect.Value, int64) error
    EncodeBoolAt(bool, int64) error
    EncodeByteAt(byte, int64) error
    EncodeIntAt(int, int64) error
    EncodeInt8At(int8, int64) error
    EncodeInt16At(int16, int64) error
    EncodeInt32At(int32, int64) error
    EncodeInt64At(int64, int64) error
    EncodeUintAt(uint, int64) error
    EncodeUint8At(uint8, int64) error
    EncodeUint16At(uint16, int64) error
    EncodeUint32At(uint32, int64) error
    EncodeUint64At(uint64, int64) error
    EncodeUintptrAt(uintptr, int64) error
    EncodeUnsafePtrAt(unsafe.Pointer, int64) error
    EncodeFloat32At(float32, int64) error
    EncodeFloat64At(float64, int64) error
    EncodeComplex64At(complex64, int64) error
    EncodeComplex128At(complex128, int64) error
    EncodeStringAt(string, int64) error
    EncodeTimeAt(time.Time, int64) error
    EncodeBoolSliceAt([]bool, int64) error
    EncodeByteSliceAt([]byte, int64) error
    EncodeIntSliceAt([]int, int64) error
    EncodeInt8SliceAt([]int8, int64) error
    EncodeInt16SliceAt([]int16, int64) error
    EncodeInt32SliceAt([]int32, int64) error
    EncodeInt64SliceAt([]int64, int64) error
    EncodeUintSliceAt([]uint, int64) error
    EncodeUint8SliceAt([]uint8, int64) error
    EncodeUint16SliceAt([]uint16, int64) error
    EncodeUint32SliceAt([]uint32, int64) error
    EncodeUint64SliceAt([]uint64, int64) error
    EncodeUintptrSliceAt([]uintptr, int64) error
    EncodeFloat32SliceAt([]float32, int64) error
    EncodeFloat64SliceAt([]float64, int64) error
    EncodeComplex64SliceAt([]complex64, int64) error
    EncodeComplex128SliceAt([]complex128, int64) error
    EncodeStringSliceAt([]string, int64) error
    EncodeSliceAt([]interface{}, int64) error
    EncodeSerializableAt(Serializable, int64) error
}

type tWriter struct {
    tIoBase

    w   io.Writer
}

func (self *tWriter) CloneWriter(w io.Writer) Writer {
    return &tWriter{
        tIoBase: tIoBase{
            tCodecBase: tCodecBase{
                spec: self.spec.clone()
            },

            w: w,
        },
    }
}

func (self *tWriter) WriteHeader() error {
    return self.spec.write(self.w)
}

type tWriterAt struct {
    tIoBase

    w   io.WriterAt
}

func (self *tWriterAt) WriteHeaderAtBegin() error {
    return self.spec.write_at(self.w, 0)
}

func (self *tWriterAt) WriteHeaderAt(offset int64) error {
    return self.spec.write_at(self.w, offset)
}

func (self *tWriterAt) CloneWriterAt(w io.WriterAt) WriterAt {
    return &tWriterAt{
        tIoBase: tIoBase{
            tCodecBase: tCodecBase{
                spec: self.spec.clone()
            },

            w: w,
        },
    }
}
