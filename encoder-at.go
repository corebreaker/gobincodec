package bincodec

import (
    "io"
    "reflect"
    "time"
    "unsafe"
)

type tEncAt struct {
    w io.WriterAt
}

func (self *tEncAt) seek(p int64) Encoder {
    return &tSeekedEncoder{
        tSeekedCodec: tSeekedCodec{
            p: p,
        },

        w:  self,
    }
}

func (self *tEncAt) encode(v interface{}, at int64) (int, error)        { return 0, nil }
func (self *tEncAt) enc_value(v reflect.Value, at int64) (int, error)   { return 0, nil }
func (self *tEncAt) enc_bool(v bool, at int64) (int, error)             { return 0, nil }
func (self *tEncAt) enc_byte(v byte, at int64) (int, error)             { return 0, nil }
func (self *tEncAt) enc_int(v int, at int64) (int, error)               { return 0, nil }
func (self *tEncAt) enc_int8(v int8, at int64) (int, error)             { return 0, nil }
func (self *tEncAt) enc_int16(v int16, at int64) (int, error)           { return 0, nil }
func (self *tEncAt) enc_int32(v int32, at int64) (int, error)           { return 0, nil }
func (self *tEncAt) enc_int64(v int64, at int64) (int, error)           { return 0, nil }
func (self *tEncAt) enc_uint(v uint, at int64) (int, error)             { return 0, nil }
func (self *tEncAt) enc_uint8(v uint8, at int64) (int, error)           { return 0, nil }
func (self *tEncAt) enc_uint16(v uint16, at int64) (int, error)         { return 0, nil }
func (self *tEncAt) enc_uint32(v uint32, at int64) (int, error)         { return 0, nil }
func (self *tEncAt) enc_uint64(v uint64, at int64) (int, error)         { return 0, nil }
func (self *tEncAt) enc_uintptr(v uintptr, at int64) (int, error)       { return 0, nil }
func (self *tEncAt) enc_ptr(v unsafe.Pointer, at int64) (int, error)    { return 0, nil }
func (self *tEncAt) enc_float32(v float32, at int64) (int, error)       { return 0, nil }
func (self *tEncAt) enc_float64(v float64, at int64) (int, error)       { return 0, nil }
func (self *tEncAt) enc_cplx64(v complex64, at int64) (int, error)      { return 0, nil }
func (self *tEncAt) enc_cplx128(v complex128, at int64) (int, error)    { return 0, nil }
func (self *tEncAt) enc_string(v string, at int64) (int, error)         { return 0, nil }
func (self *tEncAt) enc_time(v time.Time, at int64) (int, error)        { return 0, nil }
func (self *tEncAt) enc_bools(v []bool, at int64) (int, error)          { return 0, nil }
func (self *tEncAt) enc_bytes(v []byte, at int64) (int, error)          { return 0, nil }
func (self *tEncAt) enc_ints(v []int, at int64) (int, error)            { return 0, nil }
func (self *tEncAt) enc_int8s(v []int8, at int64) (int, error)          { return 0, nil }
func (self *tEncAt) enc_int16s(v []int16, at int64) (int, error)        { return 0, nil }
func (self *tEncAt) enc_int32s(v []int32, at int64) (int, error)        { return 0, nil }
func (self *tEncAt) enc_int64s(v []int64, at int64) (int, error)        { return 0, nil }
func (self *tEncAt) enc_uints(v []uint, at int64) (int, error)          { return 0, nil }
func (self *tEncAt) enc_uint8s(v []uint8, at int64) (int, error)        { return 0, nil }
func (self *tEncAt) enc_uint16s(v []uint16, at int64) (int, error)      { return 0, nil }
func (self *tEncAt) enc_uint32s(v []uint32, at int64) (int, error)      { return 0, nil }
func (self *tEncAt) enc_uint64s(v []uint64, at int64) (int, error)      { return 0, nil }
func (self *tEncAt) enc_uintptrs(v []uintptr, at int64) (int, error)    { return 0, nil }
func (self *tEncAt) enc_float32s(v []float32, at int64) (int, error)    { return 0, nil }
func (self *tEncAt) enc_float64s(v []float64, at int64) (int, error)    { return 0, nil }
func (self *tEncAt) enc_cplx64s(v []complex64, at int64) (int, error)   { return 0, nil }
func (self *tEncAt) enc_cplx128s(v []complex128, at int64) (int, error) { return 0, nil }
func (self *tEncAt) enc_strings(v []string, at int64) (int, error)      { return 0, nil }
func (self *tEncAt) enc_slice(v []interface{}, at int64) (int, error)   { return 0, nil }
func (self *tEncAt) enc_serial(v Serializable, at int64) (int, error)   { return 0, nil }

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
    EncodeSliceAt([]interface{}, int64) error
    EncodeSerializableAt(Serializable, int64) error
}

type tWriterAt struct {
    tIoBase

    enc *tEncAt
}

func (self *tWriterAt) WriteHeaderAtBegin() error        { return self.spec.write_at(self.enc.w, 0) }
func (self *tWriterAt) WriteHeaderAt(offset int64) error { return self.spec.write_at(self.enc.w, offset) }

func (self *tWriterAt) CloneWriterAt(w io.WriterAt) WriterAt {
    return &tWriterAt{
        tIoBase: tIoBase{
            tCodecBase: tCodecBase{
                spec: self.spec.clone(),
            },
        },

        enc: &tEncAt{
            w: w,
        },
    }
}

func (self *tWriterAt) EncodeAt(v interface{}, p int64) error                 { return gerr(self.enc.encode(v, p)) }
func (self *tWriterAt) EncodeValueAt(v reflect.Value, p int64) error          { return gerr(self.enc.enc_value(v, p)) }
func (self *tWriterAt) EncodeBoolAt(v bool, p int64) error                    { return gerr(self.enc.enc_bool(v, p)) }
func (self *tWriterAt) EncodeByteAt(v byte, p int64) error                    { return gerr(self.enc.enc_byte(v, p)) }
func (self *tWriterAt) EncodeIntAt(v int, p int64) error                      { return gerr(self.enc.enc_int(v, p)) }
func (self *tWriterAt) EncodeInt8At(v int8, p int64) error                    { return gerr(self.enc.enc_int8(v, p)) }
func (self *tWriterAt) EncodeInt16At(v int16, p int64) error                  { return gerr(self.enc.enc_int16(v, p)) }
func (self *tWriterAt) EncodeInt32At(v int32, p int64) error                  { return gerr(self.enc.enc_int32(v, p)) }
func (self *tWriterAt) EncodeInt64At(v int64, p int64) error                  { return gerr(self.enc.enc_int64(v, p)) }
func (self *tWriterAt) EncodeUintAt(v uint, p int64) error                    { return gerr(self.enc.enc_uint(v, p)) }
func (self *tWriterAt) EncodeUint8At(v uint8, p int64) error                  { return gerr(self.enc.enc_uint8(v, p)) }
func (self *tWriterAt) EncodeUint16At(v uint16, p int64) error                { return gerr(self.enc.enc_uint16(v, p)) }
func (self *tWriterAt) EncodeUint32At(v uint32, p int64) error                { return gerr(self.enc.enc_uint32(v, p)) }
func (self *tWriterAt) EncodeUint64At(v uint64, p int64) error                { return gerr(self.enc.enc_uint64(v, p)) }
func (self *tWriterAt) EncodeUintptrAt(v uintptr, p int64) error              { return gerr(self.enc.enc_uintptr(v, p)) }
func (self *tWriterAt) EncodePtrAt(v unsafe.Pointer, p int64) error           { return gerr(self.enc.enc_ptr(v, p)) }
func (self *tWriterAt) EncodeFloat32At(v float32, p int64) error              { return gerr(self.enc.enc_float32(v, p)) }
func (self *tWriterAt) EncodeFloat64At(v float64, p int64) error              { return gerr(self.enc.enc_float64(v, p)) }
func (self *tWriterAt) EncodeComplex64At(v complex64, p int64) error          { return gerr(self.enc.enc_cplx64(v, p)) }
func (self *tWriterAt) EncodeComplex128At(v complex128, p int64) error        { return gerr(self.enc.enc_cplx128(v, p)) }
func (self *tWriterAt) EncodeStringAt(v string, p int64) error                { return gerr(self.enc.enc_string(v, p)) }
func (self *tWriterAt) EncodeTimeAt(v time.Time, p int64) error               { return gerr(self.enc.enc_time(v, p)) }
func (self *tWriterAt) EncodeBoolSliceAt(v []bool, p int64) error             { return gerr(self.enc.enc_bools(v, p)) }
func (self *tWriterAt) EncodeByteSliceAt(v []byte, p int64) error             { return gerr(self.enc.enc_bytes(v, p)) }
func (self *tWriterAt) EncodeIntSliceAt(v []int, p int64) error               { return gerr(self.enc.enc_ints(v, p)) }
func (self *tWriterAt) EncodeInt8SliceAt(v []int8, p int64) error             { return gerr(self.enc.enc_int8s(v, p)) }
func (self *tWriterAt) EncodeInt16SliceAt(v []int16, p int64) error           { return gerr(self.enc.enc_int16s(v, p)) }
func (self *tWriterAt) EncodeInt32SliceAt(v []int32, p int64) error           { return gerr(self.enc.enc_int32s(v, p)) }
func (self *tWriterAt) EncodeInt64SliceAt(v []int64, p int64) error           { return gerr(self.enc.enc_int64s(v, p)) }
func (self *tWriterAt) EncodeUintSliceAt(v []uint, p int64) error             { return gerr(self.enc.enc_uints(v, p)) }
func (self *tWriterAt) EncodeUint8SliceAt(v []uint8, p int64) error           { return gerr(self.enc.enc_uint8s(v, p)) }
func (self *tWriterAt) EncodeUint16SliceAt(v []uint16, p int64) error         { return gerr(self.enc.enc_uint16s(v, p)) }
func (self *tWriterAt) EncodeUint32SliceAt(v []uint32, p int64) error         { return gerr(self.enc.enc_uint32s(v, p)) }
func (self *tWriterAt) EncodeUint64SliceAt(v []uint64, p int64) error         { return gerr(self.enc.enc_uint64s(v, p)) }
func (self *tWriterAt) EncodeUintptrSliceAt(v []uintptr, p int64) error       { return gerr(self.enc.enc_uintptrs(v, p)) }
func (self *tWriterAt) EncodeFloat32SliceAt(v []float32, p int64) error       { return gerr(self.enc.enc_float32s(v, p)) }
func (self *tWriterAt) EncodeFloat64SliceAt(v []float64, p int64) error       { return gerr(self.enc.enc_float64s(v, p)) }
func (self *tWriterAt) EncodeComplex64SliceAt(v []complex64, p int64) error   { return gerr(self.enc.enc_cplx64s(v, p)) }
func (self *tWriterAt) EncodeComplex128SliceAt(v []complex128, p int64) error { return gerr(self.enc.enc_cplx128s(v, p)) }
func (self *tWriterAt) EncodeStringSliceAt(v []string, p int64) error         { return gerr(self.enc.enc_strings(v, p)) }
func (self *tWriterAt) EncodeSliceAt(v []interface{}, p int64) error          { return gerr(self.enc.enc_slice(v, p)) }
func (self *tWriterAt) EncodeSerializableAt(v Serializable, p int64) error    { return gerr(self.enc.enc_serial(v, p)) }
