package util

import (
	"encoding/binary"
	"io"
	"reflect"

	"github.com/corebreaker/goerrors"
)

func DecodeNum(r io.Reader, n interface{}) (int, error) {
	if err := goerrors.DecorateError(binary.Read(r, binary.BigEndian, n)); err != nil {
		return 0, err
	}

	return binary.Size(n), nil
}

func EncodeNum(w io.Writer, n interface{}) (int, error) {
	if err := goerrors.DecorateError(binary.Write(w, binary.BigEndian, n)); err != nil {
		return 0, err
	}

	return binary.Size(n), nil
}

func GetInt(v reflect.Value) (res int64, ok bool) {
	defer func() {
		defer func() {
			defer func() {
				defer goerrors.DiscardPanic()

				err := recover()
				if err == nil {
					return
				}

				res = int64(real(v.Complex()))
				ok = true
			}()

			err := recover()
			if err == nil {
				return
			}

			res = int64(v.Float())
			ok = true
		}()

		err := recover()
		if err == nil {
			return
		}

		res = int64(v.Uint())
		ok = true
	}()

	return v.Int(), true
}

func GetUint(v reflect.Value) (res uint64, ok bool) {
	defer func() {
		defer func() {
			defer func() {
				defer goerrors.DiscardPanic()

				err := recover()
				if err == nil {
					return
				}

				res = uint64(real(v.Complex()))
				ok = true
			}()

			err := recover()
			if err == nil {
				return
			}

			res = uint64(v.Float())
			ok = true
		}()

		err := recover()
		if err == nil {
			return
		}

		res = uint64(v.Int())
		ok = true
	}()

	return v.Uint(), true
}

func GetFloat(v reflect.Value) (res float64, ok bool) {
	defer func() {
		defer func() {
			defer func() {
				defer goerrors.DiscardPanic()

				err := recover()
				if err == nil {
					return
				}

				res = real(v.Complex())
				ok = true
			}()

			err := recover()
			if err == nil {
				return
			}

			res = float64(v.Uint())
			ok = true
		}()

		err := recover()
		if err == nil {
			return
		}

		res = float64(v.Int())
		ok = true
	}()

	return v.Float(), true
}

func GetComplex(v reflect.Value) (res complex128, ok bool) {
	defer func() {
		defer func() {
			defer func() {
				defer goerrors.DiscardPanic()

				err := recover()
				if err == nil {
					return
				}

				res = complex(float64(v.Uint()), 0)
				ok = true
			}()

			err := recover()
			if err == nil {
				return
			}

			res = complex(float64(v.Int()), 0)
			ok = true
		}()

		err := recover()
		if err == nil {
			return
		}

		res = complex(v.Float(), 0)
		ok = true
	}()

	return v.Complex(), true
}
