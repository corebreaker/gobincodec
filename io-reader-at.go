package gobincodec

import (
	"io"
	"reflect"
	"time"
	"unsafe"

	"github.com/corebreaker/gobincodec/defs"
)

type ReaderAt interface {
	CodecBase
	IoBase

	ReadHeaderAtBegin() error
	ReadHeaderAt(int64) error
	CloneReaderAt(io.ReaderAt) ReaderAt

	DecodeAt(interface{}, int64) error
	DecodeNextAt(int64) (interface{}, error)
	DecodeValueAt(int64) (*reflect.Value, error)
	DecodeBoolAt(int64) (bool, error)
	DecodeByteAt(int64) (byte, error)
	DecodeIntAt(int64) (int, error)
	DecodeInt8At(int64) (int8, error)
	DecodeInt16At(int64) (int16, error)
	DecodeInt32At(int64) (int32, error)
	DecodeInt64At(int64) (int64, error)
	DecodeUintAt(int64) (uint, error)
	DecodeUint8At(int64) (uint8, error)
	DecodeUint16At(int64) (uint16, error)
	DecodeUint32At(int64) (uint32, error)
	DecodeUint64At(int64) (uint64, error)
	DecodeUintptrAt(int64) (uintptr, error)
	DecodePtrAt(int64) (unsafe.Pointer, error)
	DecodeFloat32At(int64) (float32, error)
	DecodeFloat64At(int64) (float64, error)
	DecodeComplex64At(int64) (complex64, error)
	DecodeComplex128At(int64) (complex128, error)
	DecodeStringAt(int64) (string, error)
	DecodeTimeAt(int64) (*time.Time, error)
	DecodeBoolSliceAt(int64) ([]bool, error)
	DecodeByteSliceAt(int64) ([]byte, error)
	DecodeIntSliceAt(int64) ([]int, error)
	DecodeInt8SliceAt(int64) ([]int8, error)
	DecodeInt16SliceAt(int64) ([]int16, error)
	DecodeInt32SliceAt(int64) ([]int32, error)
	DecodeInt64SliceAt(int64) ([]int64, error)
	DecodeUintSliceAt(int64) ([]uint, error)
	DecodeUint8SliceAt(int64) ([]uint8, error)
	DecodeUint16SliceAt(int64) ([]uint16, error)
	DecodeUint32SliceAt(int64) ([]uint32, error)
	DecodeUint64SliceAt(int64) ([]uint64, error)
	DecodeUintptrSliceAt(int64) ([]uintptr, error)
	DecodeFloat32SliceAt(int64) ([]float32, error)
	DecodeFloat64SliceAt(int64) ([]float64, error)
	DecodeComplex64SliceAt(int64) ([]complex64, error)
	DecodeComplex128SliceAt(int64) ([]complex128, error)
	DecodeStringSliceAt(int64) ([]string, error)
	DecodePtrSliceAt(int64) ([]unsafe.Pointer, error)
	DecodeTimeSliceAt(int64) ([]time.Time, error)
	DecodeSliceAt(int64) ([]interface{}, error)
	DecodeSerializableAt(int64) (defs.Serializable, error)
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

func (ra *tReaderAt) DecodeAt(v interface{}, p int64) error     { return gerr(ra.dec.decode(v, p)) }
func (ra *tReaderAt) DecodeNextAt(p int64) (interface{}, error) { return gerrNext(ra.dec.decNext(p)) }
func (ra *tReaderAt) DecodeValueAt(p int64) (*reflect.Value, error) {
	return gerrVal(ra.dec.decValue(v, p))
}
func (ra *tReaderAt) DecodeBoolAt(p int64) (bool, error)          { return gerrBool(ra.dec.decBool(v, p)) }
func (ra *tReaderAt) DecodeByteAt(p int64) (byte, error)          { return gerrByte(ra.dec.decByte(v, p)) }
func (ra *tReaderAt) DecodeIntAt(p int64) (int, error)            { return gerrI(ra.dec.decInt(v, p)) }
func (ra *tReaderAt) DecodeInt8At(p int64) (int8, error)          { return gerrI8(ra.dec.decInt8(v, p)) }
func (ra *tReaderAt) DecodeInt16At(p int64) (int16, error)        { return gerrI16(ra.dec.decInt16(v, p)) }
func (ra *tReaderAt) DecodeInt32At(p int64) (int32, error)        { return gerrI32(ra.dec.decInt32(v, p)) }
func (ra *tReaderAt) DecodeInt64At(p int64) (int64, error)        { return gerrI64(ra.dec.decInt64(v, p)) }
func (ra *tReaderAt) DecodeUintAt(p int64) (uint, error)          { return gerrU(ra.dec.decUint(v, p)) }
func (ra *tReaderAt) DecodeUint8At(p int64) (uint8, error)        { return gerrU8(ra.dec.decUint8(v, p)) }
func (ra *tReaderAt) DecodeUint16At(p int64) (uint16, error)      { return gerrU16(ra.dec.decUint16(v, p)) }
func (ra *tReaderAt) DecodeUint32At(p int64) (uint32, error)      { return gerrU32(ra.dec.decUint32(v, p)) }
func (ra *tReaderAt) DecodeUint64At(p int64) (uint64, error)      { return gerrU64(ra.dec.decUint64(v, p)) }
func (ra *tReaderAt) DecodeUintptrAt(p int64) (uintptr, error)    { return gerrUp(ra.dec.decUintptr(v, p)) }
func (ra *tReaderAt) DecodePtrAt(p int64) (unsafe.Pointer, error) { return gerrUs(ra.dec.decPtr(v, p)) }
func (ra *tReaderAt) DecodeFloat32At(p int64) (float32, error)    { return gerrF1(ra.dec.decFloat32(v, p)) }
func (ra *tReaderAt) DecodeFloat64At(p int64) (float64, error)    { return gerrF2(ra.dec.decFloat64(v, p)) }
func (ra *tReaderAt) DecodeComplex64At(p int64) (complex64, error) {
	return gerrC1(ra.dec.decCplx64(v, p))
}
func (ra *tReaderAt) DecodeComplex128At(p int64) (complex128, error) {
	return gerrC2(ra.dec.decCplx128(v, p))
}
func (ra *tReaderAt) DecodeStringAt(p int64) (string, error)   { return gerrStr(ra.dec.decString(v, p)) }
func (ra *tReaderAt) DecodeTimeAt(p int64) (*time.Time, error) { return gerrTime(ra.dec.decTime(v, p)) }
func (ra *tReaderAt) DecodeBoolSliceAt(p int64) ([]bool, error) {
	return gerrBools(ra.dec.decBools(v, p))
}
func (ra *tReaderAt) DecodeByteSliceAt(p int64) ([]byte, error) {
	return gerrBytes(ra.dec.decBytes(v, p))
}
func (ra *tReaderAt) DecodeIntSliceAt(p int64) ([]int, error)   { return gerrIs(ra.dec.decInts(v, p)) }
func (ra *tReaderAt) DecodeInt8SliceAt(p int64) ([]int8, error) { return gerrI8s(ra.dec.decInt8s(v, p)) }
func (ra *tReaderAt) DecodeInt16SliceAt(p int64) ([]int16, error) {
	return gerrI16s(ra.dec.decInt16s(v, p))
}
func (ra *tReaderAt) DecodeInt32SliceAt(p int64) ([]int32, error) {
	return gerrI32s(ra.dec.decInt32s(v, p))
}
func (ra *tReaderAt) DecodeInt64SliceAt(p int64) ([]int64, error) {
	return gerrI64s(ra.dec.decInt64s(v, p))
}
func (ra *tReaderAt) DecodeUintSliceAt(p int64) ([]uint, error) { return gerrUs(ra.dec.decUints(v, p)) }
func (ra *tReaderAt) DecodeUint8SliceAt(p int64) ([]uint8, error) {
	return gerrU8s(ra.dec.decUint8s(v, p))
}
func (ra *tReaderAt) DecodeUint16SliceAt(p int64) ([]uint16, error) {
	return gerrU16s(ra.dec.decUint16s(v, p))
}
func (ra *tReaderAt) DecodeUint32SliceAt(p int64) ([]uint32, error) {
	return gerrU32s(ra.dec.decUint32s(v, p))
}
func (ra *tReaderAt) DecodeUint64SliceAt(p int64) ([]uint64, error) {
	return gerrU64s(ra.dec.decUint64s(v, p))
}
func (ra *tReaderAt) DecodeUintptrSliceAt(p int64) ([]uintptr, error) {
	return gerrUps(ra.dec.decUintptrs(v, p))
}
func (ra *tReaderAt) DecodePtrSliceAt(p int64) ([]unsafe.Pointer, error) {
	return gerrUss(ra.dec.decPtrs(v, p))
}
func (ra *tReaderAt) DecodeFloat32SliceAt(p int64) ([]float32, error) {
	return gerrF1s(ra.dec.decFloat32s(v, p))
}
func (ra *tReaderAt) DecodeFloat64SliceAt(p int64) ([]float64, error) {
	return gerrF2s(ra.dec.decFloat64s(v, p))
}
func (ra *tReaderAt) DecodeComplex64SliceAt(p int64) ([]complex64, error) {
	return gerrC1s(ra.dec.decCplx64s(v, p))
}
func (ra *tReaderAt) DecodeComplex128SliceAt(p int64) ([]complex128, error) {
	return gerrC2s(ra.dec.decCplx128s(v, p))
}
func (ra *tReaderAt) DecodeStringSliceAt(p int64) ([]string, error) {
	return gerrStrs(ra.dec.decStrings(v, p))
}
func (ra *tReaderAt) DecodeTimeSliceAt(p int64) ([]time.Time, error) {
	return gerrTimes(ra.dec.decTimes(v, p))
}
func (ra *tReaderAt) DecodeSliceAt(p int64) ([]interface{}, error) {
	return gerrSlice(ra.dec.decSlice(v, p))
}
func (ra *tReaderAt) DecodeSerializableAt(p int64) (defs.Serializable, error) {
	return gerrSer(ra.dec.decSerial(v, p))
}
