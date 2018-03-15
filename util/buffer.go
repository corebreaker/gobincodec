package util

import (
	"io"
)

func TransferFrom(dest io.ReaderFrom, src io.Reader) error {
	return wrapError64(dest.ReadFrom(src))
}

func TransferTo(dest io.Writer, src io.WriterTo) error {
	return wrapError64(src.WriteTo(dest))
}
