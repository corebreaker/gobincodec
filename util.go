package gobincodec

import (
	"time"
	"unsafe"

	"github.com/corebreaker/gobincodec/defs"
)

func gerr(i int, err error) error                                              { return err }
func gerrNext(r interface{}, i int, err error) (interface{}, error)            { return r, err }
func gerrVal(r *reflect.Value, i int, err error) (*reflect.Value, error)       { return r, err }
func gerrBool(r bool, i int, err error) (bool, error)                          { return r, err }
func gerrByte(r byte, i int, err error) (byte, error)                          { return r, err }
func gerrI(r int, i int, err error) (int, error)                               { return r, err }
func gerrI8(r int8, i int, err error) (int8, error)                            { return r, err }
func gerrI16(r int16, i int, err error) (int16, error)                         { return r, err }
func gerrI32(r int32, i int, err error) (int32, error)                         { return r, err }
func gerrI64(r int64, i int, err error) (int64, error)                         { return r, err }
func gerrU(r uint, i int, err error) (uint, error)                             { return r, err }
func gerrU8(r uint8, i int, err error) (uint8, error)                          { return r, err }
func gerrU16(r uint16, i int, err error) (uint16, error)                       { return r, err }
func gerrU32(r uint32, i int, err error) (uint32, error)                       { return r, err }
func gerrU64(r uint64, i int, err error) (uint64, error)                       { return r, err }
func gerrUp(r uintptr, i int, err error) (uintptr, error)                      { return r, err }
func gerrUs(r unsafe.Pointer, i int, err error) (unsafe.Pointer, error)        { return r, err }
func gerrF1(r float32, i int, err error) (float32, error)                      { return r, err }
func gerrF2(r float64, i int, err error) (float64, error)                      { return r, err }
func gerrC1(r complex64, i int, err error) (complex64, error)                  { return r, err }
func gerrC2(r complex128, i int, err error) (complex128, error)                { return r, err }
func gerrStr(r string, i int, err error) (string, error)                       { return r, err }
func gerrTime(r *time.Time, i int, err error) (*time.Time, error)              { return r, err }
func gerrBools(r []bool, i int, err error) ([]bool, error)                     { return r, err }
func gerrBytes(r []byte, i int, err error) ([]byte, error)                     { return r, err }
func gerrIs(r []int, i int, err error) ([]int, error)                          { return r, err }
func gerrI8s(r []int8, i int, err error) ([]int8, error)                       { return r, err }
func gerrI16s(r []int16, i int, err error) ([]int16, error)                    { return r, err }
func gerrI32s(r []int32, i int, err error) ([]int32, error)                    { return r, err }
func gerrI64s(r []int64, i int, err error) ([]int64, error)                    { return r, err }
func gerrUs(r []uint, i int, err error) ([]uint, error)                        { return r, err }
func gerrU8s(r []uint8, i int, err error) ([]uint8, error)                     { return r, err }
func gerrU16s(r []uint16, i int, err error) ([]uint16, error)                  { return r, err }
func gerrU32s(r []uint32, i int, err error) ([]uint32, error)                  { return r, err }
func gerrU64s(r []uint64, i int, err error) ([]uint64, error)                  { return r, err }
func gerrUps(r []uintptr, i int, err error) ([]uintptr, error)                 { return r, err }
func gerrUss(r []unsafe.Pointer, i int, err error) ([]unsafe.Pointer, error)   { return r, err }
func gerrF1s(r []float32, i int, err error) ([]float32, error)                 { return r, err }
func gerrF2s(r []float64, i int, err error) ([]float64, error)                 { return r, err }
func gerrC1s(r []complex64, i int, err error) ([]complex64, error)             { return r, err }
func gerrC2s(r []complex128, i int, err error) ([]complex128, error)           { return r, err }
func gerrStrs(r []string, i int, err error) ([]string, error)                  { return r, err }
func gerrTimes(r []time.Time, i int, err error) ([]time.Time, error)           { return r, err }
func gerrSlice(r interface{}, i int, err error) (interface{}, error)           { return r, err }
func gerrSer(r defs.Serializable, i int, err error) (defs.Serializable, error) { return r, err }
