package util

import (
	"encoding/binary"
	"io"

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
