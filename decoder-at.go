package gobincodec

import (
    "io"
    "reflect"
    "time"
    "unsafe"
)

type tDecAt struct {
    r io.ReaderAt
}

func (da *tDecAt) seek(p int64) Decoder {
    return &tSeekedDecoder{
        tSeekedCodec: tSeekedCodec{
            p: p,
        },

        r:  da,
    }
}

func (da *tDecAt) decode(v interface{}, at int64) (int, error)        { return 0, nil }
func (da *tDecAt) decNext(at int64) (interface{}, int, error)         { return nil, 0, nil }
func (da *tDecAt) decValue(v reflect.Value, at int64) (int, error)    { return 0, nil }
func (da *tDecAt) decBool(v *bool, at int64) (int, error)             { return 0, nil }
func (da *tDecAt) decByte(v *byte, at int64) (int, error)             { return 0, nil }
func (da *tDecAt) decInt(v *int, at int64) (int, error)               { return 0, nil }
func (da *tDecAt) decInt8(v *int8, at int64) (int, error)             { return 0, nil }
func (da *tDecAt) decInt16(v *int16, at int64) (int, error)           { return 0, nil }
func (da *tDecAt) decInt32(v *int32, at int64) (int, error)           { return 0, nil }
func (da *tDecAt) decInt64(v *int64, at int64) (int, error)           { return 0, nil }
func (da *tDecAt) decUint(v *uint, at int64) (int, error)             { return 0, nil }
func (da *tDecAt) decUint8(v *uint8, at int64) (int, error)           { return 0, nil }
func (da *tDecAt) decUint16(v *uint16, at int64) (int, error)         { return 0, nil }
func (da *tDecAt) decUint32(v *uint32, at int64) (int, error)         { return 0, nil }
func (da *tDecAt) decUint64(v *uint64, at int64) (int, error)         { return 0, nil }
func (da *tDecAt) decUintptr(v *uintptr, at int64) (int, error)       { return 0, nil }
func (da *tDecAt) decPtr(v *unsafe.Pointer, at int64) (int, error)    { return 0, nil }
func (da *tDecAt) decFloat32(v *float32, at int64) (int, error)       { return 0, nil }
func (da *tDecAt) decFloat64(v *float64, at int64) (int, error)       { return 0, nil }
func (da *tDecAt) decCplx64(v *complex64, at int64) (int, error)      { return 0, nil }
func (da *tDecAt) decCplx128(v *complex128, at int64) (int, error)    { return 0, nil }
func (da *tDecAt) decString(v *string, at int64) (int, error)         { return 0, nil }
func (da *tDecAt) decTime(v *time.Time, at int64) (int, error)        { return 0, nil }
func (da *tDecAt) decBools(v *[]bool, at int64) (int, error)          { return 0, nil }
func (da *tDecAt) decBytes(v *[]byte, at int64) (int, error)          { return 0, nil }
func (da *tDecAt) decInts(v *[]int, at int64) (int, error)            { return 0, nil }
func (da *tDecAt) decInt8s(v *[]int8, at int64) (int, error)          { return 0, nil }
func (da *tDecAt) decInt16s(v *[]int16, at int64) (int, error)        { return 0, nil }
func (da *tDecAt) decInt32s(v *[]int32, at int64) (int, error)        { return 0, nil }
func (da *tDecAt) decInt64s(v *[]int64, at int64) (int, error)        { return 0, nil }
func (da *tDecAt) decUints(v *[]uint, at int64) (int, error)          { return 0, nil }
func (da *tDecAt) decUint8s(v *[]uint8, at int64) (int, error)        { return 0, nil }
func (da *tDecAt) decUint16s(v *[]uint16, at int64) (int, error)      { return 0, nil }
func (da *tDecAt) decUint32s(v *[]uint32, at int64) (int, error)      { return 0, nil }
func (da *tDecAt) decUint64s(v *[]uint64, at int64) (int, error)      { return 0, nil }
func (da *tDecAt) decUintptrs(v *[]uintptr, at int64) (int, error)    { return 0, nil }
func (da *tDecAt) decFloat32s(v *[]float32, at int64) (int, error)    { return 0, nil }
func (da *tDecAt) decFloat64s(v *[]float64, at int64) (int, error)    { return 0, nil }
func (da *tDecAt) decCplx64s(v *[]complex64, at int64) (int, error)   { return 0, nil }
func (da *tDecAt) decCplx128s(v *[]complex128, at int64) (int, error) { return 0, nil }
func (da *tDecAt) decStrings(v *[]string, at int64) (int, error)      { return 0, nil }
func (da *tDecAt) decPtrs(v *[]unsafe.Pointer, at int64) (int, error) { return 0, nil }
func (da *tDecAt) decTimes(v *[]time.Time, at int64) (int, error)     { return 0, nil }
func (da *tDecAt) decSlice(v *[]interface{}, at int64) (int, error)   { return 0, nil }
func (da *tDecAt) decSerial(v Serializable, at int64) (int, error)    { return 0, nil }

type ReaderAt interface {
    CodecBase
    IoBase

    ReadHeaderAtBegin() error
    ReadHeaderAt(int64) error
    CloneReaderAt(io.ReaderAt) ReaderAt

    DecodeAt(interface{}, int64) error
    DecodeNextAt(int64) (interface{}, error)
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
    DecodePtrAt(*unsafe.Pointer, int64) error
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
    DecodePtrSliceAt(*[]unsafe.Pointer, int64) error
    DecodeTimeSliceAt(*[]time.Time, int64) error
    DecodeSliceAt(*[]interface{}, int64) error
    DecodeSerializableAt(Serializable, int64) error
}

type tReaderAt struct {
    tIoBase

    dec *tDecAt
}

func (ra *tReaderAt) ReadHeaderAtBegin() error        { return ra.spec.readAt(ra.dec.r, 0) }
func (ra *tReaderAt) ReadHeaderAt(offset int64) error { return ra.spec.readAt(ra.dec.r, offset) }

func (ra *tReaderAt) CloneReaderAt(r io.ReaderAt) ReaderAt {
    return &tReaderAt{
        tIoBase: tIoBase{
            tCodecBase: tCodecBase{
                spec: ra.spec.clone(),
            },
        },

        dec: &tDecAt{
            r: r,
        },
    }
}

func (ra *tReaderAt) DecodeAt(v interface{}, p int64) error                  { return gerr(ra.dec.decode(v, p)) }
func (ra *tReaderAt) DecodeNextAt(p int64) (interface{}, error)              { return gerr2(ra.dec.decode(v, p)) }
func (ra *tReaderAt) DecodeValueAt(v reflect.Value, p int64) error           { return gerr(ra.dec.decValue(v, p)) }
func (ra *tReaderAt) DecodeBoolAt(v *bool, p int64) error                    { return gerr(ra.dec.decBool(v, p)) }
func (ra *tReaderAt) DecodeByteAt(v *byte, p int64) error                    { return gerr(ra.dec.decByte(v, p)) }
func (ra *tReaderAt) DecodeIntAt(v *int, p int64) error                      { return gerr(ra.dec.decInt(v, p)) }
func (ra *tReaderAt) DecodeInt8At(v *int8, p int64) error                    { return gerr(ra.dec.decInt8(v, p)) }
func (ra *tReaderAt) DecodeInt16At(v *int16, p int64) error                  { return gerr(ra.dec.decInt16(v, p)) }
func (ra *tReaderAt) DecodeInt32At(v *int32, p int64) error                  { return gerr(ra.dec.decInt32(v, p)) }
func (ra *tReaderAt) DecodeInt64At(v *int64, p int64) error                  { return gerr(ra.dec.decInt64(v, p)) }
func (ra *tReaderAt) DecodeUintAt(v *uint, p int64) error                    { return gerr(ra.dec.decUint(v, p)) }
func (ra *tReaderAt) DecodeUint8At(v *uint8, p int64) error                  { return gerr(ra.dec.decUint8(v, p)) }
func (ra *tReaderAt) DecodeUint16At(v *uint16, p int64) error                { return gerr(ra.dec.decUint16(v, p)) }
func (ra *tReaderAt) DecodeUint32At(v *uint32, p int64) error                { return gerr(ra.dec.decUint32(v, p)) }
func (ra *tReaderAt) DecodeUint64At(v *uint64, p int64) error                { return gerr(ra.dec.decUint64(v, p)) }
func (ra *tReaderAt) DecodeUintptrAt(v *uintptr, p int64) error              { return gerr(ra.dec.decUintptr(v, p)) }
func (ra *tReaderAt) DecodePtrAt(v *unsafe.Pointer, p int64) error           { return gerr(ra.dec.decPtr(v, p)) }
func (ra *tReaderAt) DecodeFloat32At(v *float32, p int64) error              { return gerr(ra.dec.decFloat32(v, p)) }
func (ra *tReaderAt) DecodeFloat64At(v *float64, p int64) error              { return gerr(ra.dec.decFloat64(v, p)) }
func (ra *tReaderAt) DecodeComplex64At(v *complex64, p int64) error          { return gerr(ra.dec.decCplx64(v, p)) }
func (ra *tReaderAt) DecodeComplex128At(v *complex128, p int64) error        { return gerr(ra.dec.decCplx128(v, p)) }
func (ra *tReaderAt) DecodeStringAt(v *string, p int64) error                { return gerr(ra.dec.decString(v, p)) }
func (ra *tReaderAt) DecodeTimeAt(v *time.Time, p int64) error               { return gerr(ra.dec.decTime(v, p)) }
func (ra *tReaderAt) DecodeBoolSliceAt(v *[]bool, p int64) error             { return gerr(ra.dec.decBools(v, p)) }
func (ra *tReaderAt) DecodeByteSliceAt(v *[]byte, p int64) error             { return gerr(ra.dec.decBytes(v, p)) }
func (ra *tReaderAt) DecodeIntSliceAt(v *[]int, p int64) error               { return gerr(ra.dec.decInts(v, p)) }
func (ra *tReaderAt) DecodeInt8SliceAt(v *[]int8, p int64) error             { return gerr(ra.dec.decInt8s(v, p)) }
func (ra *tReaderAt) DecodeInt16SliceAt(v *[]int16, p int64) error           { return gerr(ra.dec.decInt16s(v, p)) }
func (ra *tReaderAt) DecodeInt32SliceAt(v *[]int32, p int64) error           { return gerr(ra.dec.decInt32s(v, p)) }
func (ra *tReaderAt) DecodeInt64SliceAt(v *[]int64, p int64) error           { return gerr(ra.dec.decInt64s(v, p)) }
func (ra *tReaderAt) DecodeUintSliceAt(v *[]uint, p int64) error             { return gerr(ra.dec.decUints(v, p)) }
func (ra *tReaderAt) DecodeUint8SliceAt(v *[]uint8, p int64) error           { return gerr(ra.dec.decUint8s(v, p)) }
func (ra *tReaderAt) DecodeUint16SliceAt(v *[]uint16, p int64) error         { return gerr(ra.dec.decUint16s(v, p)) }
func (ra *tReaderAt) DecodeUint32SliceAt(v *[]uint32, p int64) error         { return gerr(ra.dec.decUint32s(v, p)) }
func (ra *tReaderAt) DecodeUint64SliceAt(v *[]uint64, p int64) error         { return gerr(ra.dec.decUint64s(v, p)) }
func (ra *tReaderAt) DecodeUintptrSliceAt(v *[]uintptr, p int64) error       { return gerr(ra.dec.decUintptrs(v, p)) }
func (ra *tReaderAt) DecodeFloat32SliceAt(v *[]float32, p int64) error       { return gerr(ra.dec.decFloat32s(v, p)) }
func (ra *tReaderAt) DecodeFloat64SliceAt(v *[]float64, p int64) error       { return gerr(ra.dec.decFloat64s(v, p)) }
func (ra *tReaderAt) DecodeComplex64SliceAt(v *[]complex64, p int64) error   { return gerr(ra.dec.decCplx64s(v, p)) }
func (ra *tReaderAt) DecodeComplex128SliceAt(v *[]complex128, p int64) error { return gerr(ra.dec.decCplx128s(v, p)) }
func (ra *tReaderAt) DecodeStringSliceAt(v *[]string, p int64) error         { return gerr(ra.dec.decStrings(v, p)) }
func (ra *tReaderAt) DecodePtrSliceAt(v *[]unsafe.Pointer, p int64) error    { return gerr(ra.dec.decPtrs(v, p)) }
func (ra *tReaderAt) DecodeTimeSliceAt(v *[]time.Time, p int64) error        { return gerr(ra.dec.decTimes(v, p)) }
func (ra *tReaderAt) DecodeSliceAt(v *[]interface{}, p int64) error          { return gerr(ra.dec.decSlice(v, p)) }
func (ra *tReaderAt) DecodeSerializableAt(v Serializable, p int64) error     { return gerr(ra.dec.decSerial(v, p)) }
