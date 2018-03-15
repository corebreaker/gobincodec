package util

import (
	"io"

	"github.com/corebreaker/goerrors"
)

func Read(r io.Reader, b []byte) error {
	return wrapError(r.Read(b))
}

func Write(w io.Writer, b []byte) error {
	return wrapError(w.Write(b))
}

func ReadByte(r io.ByteReader) (byte, error) {
	res, err := r.ReadByte()
	if err != nil {
		return 0, goerrors.DecorateError(err)
	}

	return res, nil
}

func WriteByte(w io.ByteWriter, b byte) error {
	return goerrors.DecorateError(w.WriteByte(b))
}
