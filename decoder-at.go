package bincodec

import (
    "io"
    "reflect"
    "time"
    "unsafe"
)

type tDecAt struct {
    r io.ReaderAt
}

func (self *tDecAt) seek(p int64) Decoder {
    return &tSeekedDecoder{
        tSeekedCodec: tSeekedCodec{
            p: p,
        },

        r:  self,
    }
}

func (self *tDecAt) decode(v interface{}, at int64) (int, error)         { return 0, nil }
func (self *tDecAt) dec_value(v reflect.Value, at int64) (int, error)    { return 0, nil }
func (self *tDecAt) dec_bool(v *bool, at int64) (int, error)             { return 0, nil }
func (self *tDecAt) dec_byte(v *byte, at int64) (int, error)             { return 0, nil }
func (self *tDecAt) dec_int(v *int, at int64) (int, error)               { return 0, nil }
func (self *tDecAt) dec_int8(v *int8, at int64) (int, error)             { return 0, nil }
func (self *tDecAt) dec_int16(v *int16, at int64) (int, error)           { return 0, nil }
func (self *tDecAt) dec_int32(v *int32, at int64) (int, error)           { return 0, nil }
func (self *tDecAt) dec_int64(v *int64, at int64) (int, error)           { return 0, nil }
func (self *tDecAt) dec_uint(v *uint, at int64) (int, error)             { return 0, nil }
func (self *tDecAt) dec_uint8(v *uint8, at int64) (int, error)           { return 0, nil }
func (self *tDecAt) dec_uint16(v *uint16, at int64) (int, error)         { return 0, nil }
func (self *tDecAt) dec_uint32(v *uint32, at int64) (int, error)         { return 0, nil }
func (self *tDecAt) dec_uint64(v *uint64, at int64) (int, error)         { return 0, nil }
func (self *tDecAt) dec_uintptr(v *uintptr, at int64) (int, error)       { return 0, nil }
func (self *tDecAt) dec_ptr(v *unsafe.Pointer, at int64) (int, error)    { return 0, nil }
func (self *tDecAt) dec_float32(v *float32, at int64) (int, error)       { return 0, nil }
func (self *tDecAt) dec_float64(v *float64, at int64) (int, error)       { return 0, nil }
func (self *tDecAt) dec_cplx64(v *complex64, at int64) (int, error)      { return 0, nil }
func (self *tDecAt) dec_cplx128(v *complex128, at int64) (int, error)    { return 0, nil }
func (self *tDecAt) dec_string(v *string, at int64) (int, error)         { return 0, nil }
func (self *tDecAt) dec_time(v *time.Time, at int64) (int, error)        { return 0, nil }
func (self *tDecAt) dec_bools(v *[]bool, at int64) (int, error)          { return 0, nil }
func (self *tDecAt) dec_bytes(v *[]byte, at int64) (int, error)          { return 0, nil }
func (self *tDecAt) dec_ints(v *[]int, at int64) (int, error)            { return 0, nil }
func (self *tDecAt) dec_int8s(v *[]int8, at int64) (int, error)          { return 0, nil }
func (self *tDecAt) dec_int16s(v *[]int16, at int64) (int, error)        { return 0, nil }
func (self *tDecAt) dec_int32s(v *[]int32, at int64) (int, error)        { return 0, nil }
func (self *tDecAt) dec_int64s(v *[]int64, at int64) (int, error)        { return 0, nil }
func (self *tDecAt) dec_uints(v *[]uint, at int64) (int, error)          { return 0, nil }
func (self *tDecAt) dec_uint8s(v *[]uint8, at int64) (int, error)        { return 0, nil }
func (self *tDecAt) dec_uint16s(v *[]uint16, at int64) (int, error)      { return 0, nil }
func (self *tDecAt) dec_uint32s(v *[]uint32, at int64) (int, error)      { return 0, nil }
func (self *tDecAt) dec_uint64s(v *[]uint64, at int64) (int, error)      { return 0, nil }
func (self *tDecAt) dec_uintptrs(v *[]uintptr, at int64) (int, error)    { return 0, nil }
func (self *tDecAt) dec_float32s(v *[]float32, at int64) (int, error)    { return 0, nil }
func (self *tDecAt) dec_float64s(v *[]float64, at int64) (int, error)    { return 0, nil }
func (self *tDecAt) dec_cplx64s(v *[]complex64, at int64) (int, error)   { return 0, nil }
func (self *tDecAt) dec_cplx128s(v *[]complex128, at int64) (int, error) { return 0, nil }
func (self *tDecAt) dec_strings(v *[]string, at int64) (int, error)      { return 0, nil }
func (self *tDecAt) dec_slice(v *[]interface{}, at int64) (int, error)   { return 0, nil }
func (self *tDecAt) dec_serial(v Serializable, at int64) (int, error)    { return 0, nil }

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
    DecodeSliceAt(*[]interface{}, int64) error
    DecodeSerializableAt(Serializable, int64) error
}

type tReaderAt struct {
    tIoBase

    dec *tDecAt
}

func (self *tReaderAt) ReadHeaderAtBegin() error        { return self.spec.read_at(self.dec.r, 0) }
func (self *tReaderAt) ReadHeaderAt(offset int64) error { return self.spec.read_at(self.dec.r, offset) }

func (self *tReaderAt) CloneReaderAt(r io.ReaderAt) ReaderAt {
    return &tReaderAt{
        tIoBase: tIoBase{
            tCodecBase: tCodecBase{
                spec: self.spec.clone(),
            },
        },

        dec: &tDecAt{
            r: r,
        },
    }
}

func (self *tReaderAt) DecodeAt(v interface{}, p int64) error                  { return gerr(self.dec.decode(v, p)) }
func (self *tReaderAt) DecodeValueAt(v reflect.Value, p int64) error           { return gerr(self.dec.dec_value(v, p)) }
func (self *tReaderAt) DecodeBoolAt(v *bool, p int64) error                    { return gerr(self.dec.dec_bool(v, p)) }
func (self *tReaderAt) DecodeByteAt(v *byte, p int64) error                    { return gerr(self.dec.dec_byte(v, p)) }
func (self *tReaderAt) DecodeIntAt(v *int, p int64) error                      { return gerr(self.dec.dec_int(v, p)) }
func (self *tReaderAt) DecodeInt8At(v *int8, p int64) error                    { return gerr(self.dec.dec_int8(v, p)) }
func (self *tReaderAt) DecodeInt16At(v *int16, p int64) error                  { return gerr(self.dec.dec_int16(v, p)) }
func (self *tReaderAt) DecodeInt32At(v *int32, p int64) error                  { return gerr(self.dec.dec_int32(v, p)) }
func (self *tReaderAt) DecodeInt64At(v *int64, p int64) error                  { return gerr(self.dec.dec_int64(v, p)) }
func (self *tReaderAt) DecodeUintAt(v *uint, p int64) error                    { return gerr(self.dec.dec_uint(v, p)) }
func (self *tReaderAt) DecodeUint8At(v *uint8, p int64) error                  { return gerr(self.dec.dec_uint8(v, p)) }
func (self *tReaderAt) DecodeUint16At(v *uint16, p int64) error                { return gerr(self.dec.dec_uint16(v, p)) }
func (self *tReaderAt) DecodeUint32At(v *uint32, p int64) error                { return gerr(self.dec.dec_uint32(v, p)) }
func (self *tReaderAt) DecodeUint64At(v *uint64, p int64) error                { return gerr(self.dec.dec_uint64(v, p)) }
func (self *tReaderAt) DecodeUintptrAt(v *uintptr, p int64) error              { return gerr(self.dec.dec_uintptr(v, p)) }
func (self *tReaderAt) DecodePtrAt(v *unsafe.Pointer, p int64) error           { return gerr(self.dec.dec_ptr(v, p)) }
func (self *tReaderAt) DecodeFloat32At(v *float32, p int64) error              { return gerr(self.dec.dec_float32(v, p)) }
func (self *tReaderAt) DecodeFloat64At(v *float64, p int64) error              { return gerr(self.dec.dec_float64(v, p)) }
func (self *tReaderAt) DecodeComplex64At(v *complex64, p int64) error          { return gerr(self.dec.dec_cplx64(v, p)) }
func (self *tReaderAt) DecodeComplex128At(v *complex128, p int64) error        { return gerr(self.dec.dec_cplx128(v, p)) }
func (self *tReaderAt) DecodeStringAt(v *string, p int64) error                { return gerr(self.dec.dec_string(v, p)) }
func (self *tReaderAt) DecodeTimeAt(v *time.Time, p int64) error               { return gerr(self.dec.dec_time(v, p)) }
func (self *tReaderAt) DecodeBoolSliceAt(v *[]bool, p int64) error             { return gerr(self.dec.dec_bools(v, p)) }
func (self *tReaderAt) DecodeByteSliceAt(v *[]byte, p int64) error             { return gerr(self.dec.dec_bytes(v, p)) }
func (self *tReaderAt) DecodeIntSliceAt(v *[]int, p int64) error               { return gerr(self.dec.dec_ints(v, p)) }
func (self *tReaderAt) DecodeInt8SliceAt(v *[]int8, p int64) error             { return gerr(self.dec.dec_int8s(v, p)) }
func (self *tReaderAt) DecodeInt16SliceAt(v *[]int16, p int64) error           { return gerr(self.dec.dec_int16s(v, p)) }
func (self *tReaderAt) DecodeInt32SliceAt(v *[]int32, p int64) error           { return gerr(self.dec.dec_int32s(v, p)) }
func (self *tReaderAt) DecodeInt64SliceAt(v *[]int64, p int64) error           { return gerr(self.dec.dec_int64s(v, p)) }
func (self *tReaderAt) DecodeUintSliceAt(v *[]uint, p int64) error             { return gerr(self.dec.dec_uints(v, p)) }
func (self *tReaderAt) DecodeUint8SliceAt(v *[]uint8, p int64) error           { return gerr(self.dec.dec_uint8s(v, p)) }
func (self *tReaderAt) DecodeUint16SliceAt(v *[]uint16, p int64) error         { return gerr(self.dec.dec_uint16s(v, p)) }
func (self *tReaderAt) DecodeUint32SliceAt(v *[]uint32, p int64) error         { return gerr(self.dec.dec_uint32s(v, p)) }
func (self *tReaderAt) DecodeUint64SliceAt(v *[]uint64, p int64) error         { return gerr(self.dec.dec_uint64s(v, p)) }
func (self *tReaderAt) DecodeUintptrSliceAt(v *[]uintptr, p int64) error       { return gerr(self.dec.dec_uintptrs(v, p)) }
func (self *tReaderAt) DecodeFloat32SliceAt(v *[]float32, p int64) error       { return gerr(self.dec.dec_float32s(v, p)) }
func (self *tReaderAt) DecodeFloat64SliceAt(v *[]float64, p int64) error       { return gerr(self.dec.dec_float64s(v, p)) }
func (self *tReaderAt) DecodeComplex64SliceAt(v *[]complex64, p int64) error   { return gerr(self.dec.dec_cplx64s(v, p)) }
func (self *tReaderAt) DecodeComplex128SliceAt(v *[]complex128, p int64) error { return gerr(self.dec.dec_cplx128s(v, p)) }
func (self *tReaderAt) DecodeStringSliceAt(v *[]string, p int64) error         { return gerr(self.dec.dec_strings(v, p)) }
func (self *tReaderAt) DecodeSliceAt(v *[]interface{}, p int64) error          { return gerr(self.dec.dec_slice(v, p)) }
func (self *tReaderAt) DecodeSerializableAt(v Serializable, p int64) error     { return gerr(self.dec.dec_serial(v, p)) }
