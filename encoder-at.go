package gobincodec

import (
    "io"
    "reflect"
    "time"
    "unsafe"
)

type tEncAt struct {
    w io.WriterAt
}

func (ea *tEncAt) seek(p int64) Encoder {
    return &tSeekedEncoder{
        tSeekedCodec: tSeekedCodec{
            p: p,
        },

        w:  ea,
    }
}

func (ea *tEncAt) encode(v interface{}, at int64) (int, error)        { return 0, nil }
func (ea *tEncAt) encValue(v reflect.Value, at int64) (int, error)   { return 0, nil }
func (ea *tEncAt) encBool(v bool, at int64) (int, error)             { return 0, nil }
func (ea *tEncAt) encByte(v byte, at int64) (int, error)             { return 0, nil }
func (ea *tEncAt) encInt(v int, at int64) (int, error)               { return 0, nil }
func (ea *tEncAt) encInt8(v int8, at int64) (int, error)             { return 0, nil }
func (ea *tEncAt) encInt16(v int16, at int64) (int, error)           { return 0, nil }
func (ea *tEncAt) encInt32(v int32, at int64) (int, error)           { return 0, nil }
func (ea *tEncAt) encInt64(v int64, at int64) (int, error)           { return 0, nil }
func (ea *tEncAt) encUint(v uint, at int64) (int, error)             { return 0, nil }
func (ea *tEncAt) encUint8(v uint8, at int64) (int, error)           { return 0, nil }
func (ea *tEncAt) encUint16(v uint16, at int64) (int, error)         { return 0, nil }
func (ea *tEncAt) encUint32(v uint32, at int64) (int, error)         { return 0, nil }
func (ea *tEncAt) encUint64(v uint64, at int64) (int, error)         { return 0, nil }
func (ea *tEncAt) encUintptr(v uintptr, at int64) (int, error)       { return 0, nil }
func (ea *tEncAt) encPtr(v unsafe.Pointer, at int64) (int, error)    { return 0, nil }
func (ea *tEncAt) encFloat32(v float32, at int64) (int, error)       { return 0, nil }
func (ea *tEncAt) encFloat64(v float64, at int64) (int, error)       { return 0, nil }
func (ea *tEncAt) encCplx64(v complex64, at int64) (int, error)      { return 0, nil }
func (ea *tEncAt) encCplx128(v complex128, at int64) (int, error)    { return 0, nil }
func (ea *tEncAt) encString(v string, at int64) (int, error)         { return 0, nil }
func (ea *tEncAt) encTime(v time.Time, at int64) (int, error)        { return 0, nil }
func (ea *tEncAt) encBools(v []bool, at int64) (int, error)          { return 0, nil }
func (ea *tEncAt) encBytes(v []byte, at int64) (int, error)          { return 0, nil }
func (ea *tEncAt) encInts(v []int, at int64) (int, error)            { return 0, nil }
func (ea *tEncAt) encInt8s(v []int8, at int64) (int, error)          { return 0, nil }
func (ea *tEncAt) encInt16s(v []int16, at int64) (int, error)        { return 0, nil }
func (ea *tEncAt) encInt32s(v []int32, at int64) (int, error)        { return 0, nil }
func (ea *tEncAt) encInt64s(v []int64, at int64) (int, error)        { return 0, nil }
func (ea *tEncAt) encUints(v []uint, at int64) (int, error)          { return 0, nil }
func (ea *tEncAt) encUint8s(v []uint8, at int64) (int, error)        { return 0, nil }
func (ea *tEncAt) encUint16s(v []uint16, at int64) (int, error)      { return 0, nil }
func (ea *tEncAt) encUint32s(v []uint32, at int64) (int, error)      { return 0, nil }
func (ea *tEncAt) encUint64s(v []uint64, at int64) (int, error)      { return 0, nil }
func (ea *tEncAt) encUintptrs(v []uintptr, at int64) (int, error)    { return 0, nil }
func (ea *tEncAt) encFloat32s(v []float32, at int64) (int, error)    { return 0, nil }
func (ea *tEncAt) encFloat64s(v []float64, at int64) (int, error)    { return 0, nil }
func (ea *tEncAt) encCplx64s(v []complex64, at int64) (int, error)   { return 0, nil }
func (ea *tEncAt) encCplx128s(v []complex128, at int64) (int, error) { return 0, nil }
func (ea *tEncAt) encStrings(v []string, at int64) (int, error)      { return 0, nil }
func (ea *tEncAt) encPtrs(v []unsafe.Pointer, at int64) (int, error) { return 0, nil }
func (ea *tEncAt) encTimes(v []time.Time, at int64) (int, error)     { return 0, nil }
func (ea *tEncAt) encSlice(v []interface{}, at int64) (int, error)   { return 0, nil }
func (ea *tEncAt) encSerial(v Serializable, at int64) (int, error)   { return 0, nil }

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
    EncodePtrAt(unsafe.Pointer, int64) error
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
    EncodePtrSliceAt([]unsafe.Pointer, int64) error
    EncodeTimeSliceAt([]time.Time, int64) error
    EncodeSliceAt([]interface{}, int64) error
    EncodeSerializableAt(Serializable, int64) error
}

type tWriterAt struct {
    tIoBase

    enc *tEncAt
}

func (wa *tWriterAt) WriteHeaderAtBegin() error        { return wa.spec.writeAt(wa.enc.w, 0) }
func (wa *tWriterAt) WriteHeaderAt(offset int64) error { return wa.spec.writeAt(wa.enc.w, offset) }

func (wa *tWriterAt) CloneWriterAt(w io.WriterAt) WriterAt {
    return &tWriterAt{
        tIoBase: tIoBase{
            tCodecBase: tCodecBase{
                spec: wa.spec.clone(),
            },
        },

        enc: &tEncAt{
            w: w,
        },
    }
}

func (wa *tWriterAt) EncodeAt(v interface{}, p int64) error                 { return gerr(wa.enc.encode(v, p)) }
func (wa *tWriterAt) EncodeValueAt(v reflect.Value, p int64) error          { return gerr(wa.enc.encValue(v, p)) }
func (wa *tWriterAt) EncodeBoolAt(v bool, p int64) error                    { return gerr(wa.enc.encBool(v, p)) }
func (wa *tWriterAt) EncodeByteAt(v byte, p int64) error                    { return gerr(wa.enc.encByte(v, p)) }
func (wa *tWriterAt) EncodeIntAt(v int, p int64) error                      { return gerr(wa.enc.encInt(v, p)) }
func (wa *tWriterAt) EncodeInt8At(v int8, p int64) error                    { return gerr(wa.enc.encInt8(v, p)) }
func (wa *tWriterAt) EncodeInt16At(v int16, p int64) error                  { return gerr(wa.enc.encInt16(v, p)) }
func (wa *tWriterAt) EncodeInt32At(v int32, p int64) error                  { return gerr(wa.enc.encInt32(v, p)) }
func (wa *tWriterAt) EncodeInt64At(v int64, p int64) error                  { return gerr(wa.enc.encInt64(v, p)) }
func (wa *tWriterAt) EncodeUintAt(v uint, p int64) error                    { return gerr(wa.enc.encUint(v, p)) }
func (wa *tWriterAt) EncodeUint8At(v uint8, p int64) error                  { return gerr(wa.enc.encUint8(v, p)) }
func (wa *tWriterAt) EncodeUint16At(v uint16, p int64) error                { return gerr(wa.enc.encUint16(v, p)) }
func (wa *tWriterAt) EncodeUint32At(v uint32, p int64) error                { return gerr(wa.enc.encUint32(v, p)) }
func (wa *tWriterAt) EncodeUint64At(v uint64, p int64) error                { return gerr(wa.enc.encUint64(v, p)) }
func (wa *tWriterAt) EncodeUintptrAt(v uintptr, p int64) error              { return gerr(wa.enc.encUintptr(v, p)) }
func (wa *tWriterAt) EncodePtrAt(v unsafe.Pointer, p int64) error           { return gerr(wa.enc.encPtr(v, p)) }
func (wa *tWriterAt) EncodeFloat32At(v float32, p int64) error              { return gerr(wa.enc.encFloat32(v, p)) }
func (wa *tWriterAt) EncodeFloat64At(v float64, p int64) error              { return gerr(wa.enc.encFloat64(v, p)) }
func (wa *tWriterAt) EncodeComplex64At(v complex64, p int64) error          { return gerr(wa.enc.encCplx64(v, p)) }
func (wa *tWriterAt) EncodeComplex128At(v complex128, p int64) error        { return gerr(wa.enc.encCplx128(v, p)) }
func (wa *tWriterAt) EncodeStringAt(v string, p int64) error                { return gerr(wa.enc.encString(v, p)) }
func (wa *tWriterAt) EncodeTimeAt(v time.Time, p int64) error               { return gerr(wa.enc.encTime(v, p)) }
func (wa *tWriterAt) EncodeBoolSliceAt(v []bool, p int64) error             { return gerr(wa.enc.encBools(v, p)) }
func (wa *tWriterAt) EncodeByteSliceAt(v []byte, p int64) error             { return gerr(wa.enc.encBytes(v, p)) }
func (wa *tWriterAt) EncodeIntSliceAt(v []int, p int64) error               { return gerr(wa.enc.encInts(v, p)) }
func (wa *tWriterAt) EncodeInt8SliceAt(v []int8, p int64) error             { return gerr(wa.enc.encInt8s(v, p)) }
func (wa *tWriterAt) EncodeInt16SliceAt(v []int16, p int64) error           { return gerr(wa.enc.encInt16s(v, p)) }
func (wa *tWriterAt) EncodeInt32SliceAt(v []int32, p int64) error           { return gerr(wa.enc.encInt32s(v, p)) }
func (wa *tWriterAt) EncodeInt64SliceAt(v []int64, p int64) error           { return gerr(wa.enc.encInt64s(v, p)) }
func (wa *tWriterAt) EncodeUintSliceAt(v []uint, p int64) error             { return gerr(wa.enc.encUints(v, p)) }
func (wa *tWriterAt) EncodeUint8SliceAt(v []uint8, p int64) error           { return gerr(wa.enc.encUint8s(v, p)) }
func (wa *tWriterAt) EncodeUint16SliceAt(v []uint16, p int64) error         { return gerr(wa.enc.encUint16s(v, p)) }
func (wa *tWriterAt) EncodeUint32SliceAt(v []uint32, p int64) error         { return gerr(wa.enc.encUint32s(v, p)) }
func (wa *tWriterAt) EncodeUint64SliceAt(v []uint64, p int64) error         { return gerr(wa.enc.encUint64s(v, p)) }
func (wa *tWriterAt) EncodeUintptrSliceAt(v []uintptr, p int64) error       { return gerr(wa.enc.encUintptrs(v, p)) }
func (wa *tWriterAt) EncodeFloat32SliceAt(v []float32, p int64) error       { return gerr(wa.enc.encFloat32s(v, p)) }
func (wa *tWriterAt) EncodeFloat64SliceAt(v []float64, p int64) error       { return gerr(wa.enc.encFloat64s(v, p)) }
func (wa *tWriterAt) EncodeComplex64SliceAt(v []complex64, p int64) error   { return gerr(wa.enc.encCplx64s(v, p)) }
func (wa *tWriterAt) EncodeComplex128SliceAt(v []complex128, p int64) error { return gerr(wa.enc.encCplx128s(v, p)) }
func (wa *tWriterAt) EncodeStringSliceAt(v []string, p int64) error         { return gerr(wa.enc.encStrings(v, p)) }
func (wa *tWriterAt) EncodePtrSliceAt(v []unsafe.Pointer, p int64) error    { return gerr(wa.enc.encPtrs(v, p)) }
func (wa *tWriterAt) EncodeTimeSliceAt(v []time.Time, p int64) error        { return gerr(wa.enc.encTimes(v, p)) }
func (wa *tWriterAt) EncodeSliceAt(v []interface{}, p int64) error          { return gerr(wa.enc.encSlice(v, p)) }
func (wa *tWriterAt) EncodeSerializableAt(v Serializable, p int64) error    { return gerr(wa.enc.encSerial(v, p)) }
