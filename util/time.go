package util

import (
	"time"

	"github.com/corebreaker/goerrors"
)

func MarshallTime(t *time.Time) ([]byte, error) {
	if t == nil {
		return nil, nil
	}

	res, err := t.MarshalBinary()
	if err != nil {
		return nil, goerrors.DecorateError(err)
	}

	return res, nil
}

func UnmarshalTime(b []byte, res *time.Time) error {
	if res == nil {
		return nil
	}

	return goerrors.DecorateError(res.UnmarshalBinary(b))
}
