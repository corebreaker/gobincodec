package util

import (
	"github.com/corebreaker/goerrors"
)

func wrapError(n int, err error) error {
	return goerrors.DecorateError(err)
}

func wrapError64(n int64, err error) error {
	return goerrors.DecorateError(err)
}
