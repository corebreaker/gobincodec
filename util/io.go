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

func ReadBool(r io.Reader) (bool, error) {
	var buf [1]byte

	if err := Read(r, buf[:]); err != nil {
		return false, err
	}

	return buf[0] != 0, nil
}

func WriteBool(w io.Writer, b bool) error {
	var buf [1]byte

	if b {
		buf[0] = 1
	}

	return Write(w, buf[:])
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

func TransferFrom(dest io.ReaderFrom, src io.Reader) error {
	return wrapError64(dest.ReadFrom(src))
}

func TransferTo(dest io.Writer, src io.WriterTo) error {
	return wrapError64(src.WriteTo(dest))
}
