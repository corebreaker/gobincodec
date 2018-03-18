package util

import (
	"github.com/corebreaker/goerrors"
)

func wrapError(n int, err error) (int, error) {
	return n, goerrors.DecorateError(err)
}

func wrapError64(n int64, err error) (int64, error) {
	return n, goerrors.DecorateError(err)
}
