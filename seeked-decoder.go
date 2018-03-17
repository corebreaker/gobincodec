 package bincodec

import (
    "reflect"
    "time"
    "unsafe"
)

type tSeekedDecoder struct {
    tSeekedCodec

    r   *tDecAt
}

func (dec *tSeekedDecoder) Decode(values ...interface{}) error {
    for v := range values {
        if err := dec.next(dec.r.decode(v, dec.p)); err != nil {
            return err
        }
    }

    return nil
}

func (dec *tSeekedDecoder) DecodeValue(v reflect.Value) error           { return dec.next(dec.r.decValue(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeBool(v *bool) error                    { return dec.next(dec.r.decBool(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeByte(v *byte) error                    { return dec.next(dec.r.decByte(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeInt(v *int) error                      { return dec.next(dec.r.decInt(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeInt8(v *int8) error                    { return dec.next(dec.r.decInt8(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeInt16(v *int16) error                  { return dec.next(dec.r.decInt16(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeInt32(v *int32) error                  { return dec.next(dec.r.decInt32(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeInt64(v *int64) error                  { return dec.next(dec.r.decInt64(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUint(v *uint) error                    { return dec.next(dec.r.decUint(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUint8(v *uint8) error                  { return dec.next(dec.r.decUint8(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUint16(v *uint16) error                { return dec.next(dec.r.decUint16(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUint32(v *uint32) error                { return dec.next(dec.r.decUint32(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUint64(v *uint64) error                { return dec.next(dec.r.decUint64(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUintptr(v *uintptr) error              { return dec.next(dec.r.decUintptr(v, dec.p)) }
func (dec *tSeekedDecoder) DecodePtr(v *unsafe.Pointer) error           { return dec.next(dec.r.decPtr(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeFloat32(v *float32) error              { return dec.next(dec.r.decFloat32(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeFloat64(v *float64) error              { return dec.next(dec.r.decFloat64(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeComplex64(v *complex64) error          { return dec.next(dec.r.decCplx64(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeComplex128(v *complex128) error        { return dec.next(dec.r.decCplx128(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeString(v *string) error                { return dec.next(dec.r.decString(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeTime(v *time.Time) error               { return dec.next(dec.r.decTime(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeBoolSlice(v *[]bool) error             { return dec.next(dec.r.decBools(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeByteSlice(v *[]byte) error             { return dec.next(dec.r.decBytes(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeIntSlice(v *[]int) error               { return dec.next(dec.r.decInts(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeInt8Slice(v *[]int8) error             { return dec.next(dec.r.decInt8s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeInt16Slice(v *[]int16) error           { return dec.next(dec.r.decInt16s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeInt32Slice(v *[]int32) error           { return dec.next(dec.r.decInt32s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeInt64Slice(v *[]int64) error           { return dec.next(dec.r.decInt64s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUintSlice(v *[]uint) error             { return dec.next(dec.r.decUints(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUint8Slice(v *[]uint8) error           { return dec.next(dec.r.decUint8s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUint16Slice(v *[]uint16) error         { return dec.next(dec.r.decUint16s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUint32Slice(v *[]uint32) error         { return dec.next(dec.r.decUint32s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUint64Slice(v *[]uint64) error         { return dec.next(dec.r.decUint64s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeUintptrSlice(v *[]uintptr) error       { return dec.next(dec.r.decUintptrs(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeFloat32Slice(v *[]float32) error       { return dec.next(dec.r.decFloat32s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeFloat64Slice(v *[]float64) error       { return dec.next(dec.r.decFloat64s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeComplex64Slice(v *[]complex64) error   { return dec.next(dec.r.decCplx64s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeComplex128Slice(v *[]complex128) error { return dec.next(dec.r.decCplx128s(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeStringSlice(v *[]string) error         { return dec.next(dec.r.decStrings(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeSlice(v *[]interface{}) error          { return dec.next(dec.r.decSlice(v, dec.p)) }
func (dec *tSeekedDecoder) DecodeSerializable(v Serializable) error     { return dec.next(dec.r.decSerial(v, dec.p)) }
