package gobincodec

import (
	"io"
	"reflect"
	"time"
	"unsafe"

	"github.com/corebreaker/gobincodec/defs"
)

type tEncAt struct {
	w io.WriterAt
}

func (ea *tEncAt) seek(p int64) Encoder {
	return &tSeekedEncoder{
		tSeekedCodec: tSeekedCodec{
			p: p,
		},

		w: ea,
	}
}

func (ea *tEncAt) encode(v interface{}, at int64) (int, error)          { return 0, nil }
func (ea *tEncAt) encValue(v reflect.Value, at int64) (int, error)      { return 0, nil }
func (ea *tEncAt) encBool(v bool, at int64) (int, error)                { return 0, nil }
func (ea *tEncAt) encByte(v byte, at int64) (int, error)                { return 0, nil }
func (ea *tEncAt) encInt(v int, at int64) (int, error)                  { return 0, nil }
func (ea *tEncAt) encInt8(v int8, at int64) (int, error)                { return 0, nil }
func (ea *tEncAt) encInt16(v int16, at int64) (int, error)              { return 0, nil }
func (ea *tEncAt) encInt32(v int32, at int64) (int, error)              { return 0, nil }
func (ea *tEncAt) encInt64(v int64, at int64) (int, error)              { return 0, nil }
func (ea *tEncAt) encUint(v uint, at int64) (int, error)                { return 0, nil }
func (ea *tEncAt) encUint8(v uint8, at int64) (int, error)              { return 0, nil }
func (ea *tEncAt) encUint16(v uint16, at int64) (int, error)            { return 0, nil }
func (ea *tEncAt) encUint32(v uint32, at int64) (int, error)            { return 0, nil }
func (ea *tEncAt) encUint64(v uint64, at int64) (int, error)            { return 0, nil }
func (ea *tEncAt) encUintptr(v uintptr, at int64) (int, error)          { return 0, nil }
func (ea *tEncAt) encPtr(v unsafe.Pointer, at int64) (int, error)       { return 0, nil }
func (ea *tEncAt) encFloat32(v float32, at int64) (int, error)          { return 0, nil }
func (ea *tEncAt) encFloat64(v float64, at int64) (int, error)          { return 0, nil }
func (ea *tEncAt) encCplx64(v complex64, at int64) (int, error)         { return 0, nil }
func (ea *tEncAt) encCplx128(v complex128, at int64) (int, error)       { return 0, nil }
func (ea *tEncAt) encString(v string, at int64) (int, error)            { return 0, nil }
func (ea *tEncAt) encTime(v time.Time, at int64) (int, error)           { return 0, nil }
func (ea *tEncAt) encBools(v []bool, at int64) (int, error)             { return 0, nil }
func (ea *tEncAt) encBytes(v []byte, at int64) (int, error)             { return 0, nil }
func (ea *tEncAt) encInts(v []int, at int64) (int, error)               { return 0, nil }
func (ea *tEncAt) encInt8s(v []int8, at int64) (int, error)             { return 0, nil }
func (ea *tEncAt) encInt16s(v []int16, at int64) (int, error)           { return 0, nil }
func (ea *tEncAt) encInt32s(v []int32, at int64) (int, error)           { return 0, nil }
func (ea *tEncAt) encInt64s(v []int64, at int64) (int, error)           { return 0, nil }
func (ea *tEncAt) encUints(v []uint, at int64) (int, error)             { return 0, nil }
func (ea *tEncAt) encUint8s(v []uint8, at int64) (int, error)           { return 0, nil }
func (ea *tEncAt) encUint16s(v []uint16, at int64) (int, error)         { return 0, nil }
func (ea *tEncAt) encUint32s(v []uint32, at int64) (int, error)         { return 0, nil }
func (ea *tEncAt) encUint64s(v []uint64, at int64) (int, error)         { return 0, nil }
func (ea *tEncAt) encUintptrs(v []uintptr, at int64) (int, error)       { return 0, nil }
func (ea *tEncAt) encFloat32s(v []float32, at int64) (int, error)       { return 0, nil }
func (ea *tEncAt) encFloat64s(v []float64, at int64) (int, error)       { return 0, nil }
func (ea *tEncAt) encCplx64s(v []complex64, at int64) (int, error)      { return 0, nil }
func (ea *tEncAt) encCplx128s(v []complex128, at int64) (int, error)    { return 0, nil }
func (ea *tEncAt) encStrings(v []string, at int64) (int, error)         { return 0, nil }
func (ea *tEncAt) encPtrs(v []unsafe.Pointer, at int64) (int, error)    { return 0, nil }
func (ea *tEncAt) encTimes(v []time.Time, at int64) (int, error)        { return 0, nil }
func (ea *tEncAt) encSlice(v []interface{}, at int64) (int, error)      { return 0, nil }
func (ea *tEncAt) encSerial(v defs.Serializable, at int64) (int, error) { return 0, nil }
