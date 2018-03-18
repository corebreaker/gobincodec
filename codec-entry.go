package gobincodec

import (
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
)

type tCodecEntry struct {
	name     string
	datatype reflect.Type
	codec    base.IDesc
}

func (self *tCodecEntry) clone() *tCodecEntry {
	res := new(tCodecEntry)
	*res = *self

	return res
}
