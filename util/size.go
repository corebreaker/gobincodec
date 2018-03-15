package util

import (
	"io"
)

var (
	sizes = [32]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 4, 4, 5, 6, 7, 8}
	masks = [8]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x1F, 0x1F, 0x0F, 0x07}
)

func DecodeSize(r io.Reader) (int, error) {
	var code_buf [1]byte

	if err := Read(r, code_buf[:]); err != nil {
		return 0, err
	}

	code := code_buf[0]
	size := sizes[code>>3]

	b := make([]byte, size)
	if size > 0 {
		if err := Read(r, b); err != nil {
			return 0, err
		}
	}

	type i = uint64

	d := func(pos uint) uint64 {
		return uint64(b[pos]) << (pos * 8)
	}

	res := uint64(code & masks[code>>5])
	switch size {
	case 0:

	case 1:
		res = (res << 8) | uint64(b[0])

	case 2:
		res = (res << 16) | d(1) | uint64(b[0])

	case 3:
		res = (res << 24) | d(2) | d(1) | uint64(b[0])

	case 4:
		res = (res << 32) | d(3) | d(2) | d(1) | uint64(b[0])

	case 5:
		res = (res << 40) | d(4) | d(3) | d(2) | d(1) | uint64(b[0])

	case 6:
		res = (res << 48) | d(5) | d(4) | d(3) | d(2) | d(1) | uint64(b[0])

	case 7:
		res = (res << 56) | d(6) | d(5) | d(4) | d(3) | d(2) | d(1) | uint64(b[0])

	case 8:
		res = d(7) | d(6) | d(5) | d(4) | d(3) | d(2) | d(1) | uint64(b[0])
	}

	for _, digit := range b {
		res = (res << 8) | uint64(digit)
	}

	return int(res), nil
}

func EncodeSize(w io.Writer, size int) error {
	var buf []byte

	sz := uint64(size)
	switch {
	case sz < 0x80:
		buf = []byte{
			byte(sz),
		}

	case sz < 0x2000:
		buf = []byte{
			byte(sz>>8) | 0x80,
			byte(sz),
		}

	case sz < 0x200000:
		buf = []byte{
			byte(sz>>16) | 0xA0,
			byte(sz >> 8),
			byte(sz),
		}

	case sz < 0x10000000:
		buf = []byte{
			byte(sz>>24) | 0xC0,
			byte(sz >> 16),
			byte(sz >> 8),
			byte(sz),
		}

	case sz < 0x1000000000:
		buf = []byte{
			byte(sz>>32) | 0xD0,
			byte(sz >> 24),
			byte(sz >> 16),
			byte(sz >> 8),
			byte(sz),
		}

	case sz < 0x80000000000:
		buf = []byte{
			byte(sz>>40) | 0xE0,
			byte(sz >> 32),
			byte(sz >> 24),
			byte(sz >> 16),
			byte(sz >> 8),
			byte(sz),
		}

	case sz < 0x8000000000000:
		buf = []byte{
			byte(sz>>48) | 0xE8,
			byte(sz >> 40),
			byte(sz >> 32),
			byte(sz >> 24),
			byte(sz >> 16),
			byte(sz >> 8),
			byte(sz),
		}

	case sz < 0x800000000000000:
		buf = []byte{
			byte(sz>>56) | 0xF0,
			byte(sz >> 48),
			byte(sz >> 40),
			byte(sz >> 32),
			byte(sz >> 24),
			byte(sz >> 16),
			byte(sz >> 8),
			byte(sz),
		}

	default:
		buf = []byte{
			0xFF,
			byte(sz >> 56),
			byte(sz >> 48),
			byte(sz >> 40),
			byte(sz >> 32),
			byte(sz >> 24),
			byte(sz >> 16),
			byte(sz >> 8),
			byte(sz),
		}
	}

	return Write(w, buf)
}
