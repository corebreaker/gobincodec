package util

import (
	"encoding/binary"
	"io"

	"github.com/corebreaker/goerrors"
)

func DecodeNum(r io.Reader, n interface{}) error {
	return goerrors.DecorateError(binary.Read(r, binary.BigEndian, n))
}

func EncodeNum(w io.Writer, n interface{}) error {
	return goerrors.DecorateError(binary.Write(w, binary.BigEndian, n))
}
