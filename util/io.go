package util

import (
	"io"

	"github.com/corebreaker/goerrors"
)

func Read(r io.Reader, b []byte) (int, error) {
	return wrapError(r.Read(b))
}

func Write(w io.Writer, b []byte) (int, error) {
	return wrapError(w.Write(b))
}

func ReadBool(r io.Reader) (res bool, n int, err error) {
	var buf [1]byte

	if n, err = Read(r, buf[:]); err == nil {
		res = buf[0] != 0
	}

	return
}

func WriteBool(w io.Writer, b bool) (int, error) {
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

func TransferFrom(dest io.ReaderFrom, src io.Reader) (int64, error) {
	return wrapError64(dest.ReadFrom(src))
}

func TransferTo(dest io.Writer, src io.WriterTo) (int64, error) {
	return wrapError64(src.WriteTo(dest))
}
