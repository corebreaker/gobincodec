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

		r: da,
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
func (da *tDecAt) decSerial(v *Serializable, at int64) (int, error)   { return 0, nil }
