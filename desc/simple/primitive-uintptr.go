package simple

import (
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescPrimitiveUintptr struct{ DescPrimitiveUint64 }

func (*DescPrimitiveUintptr) TypeEquals(reflect.Type) bool {

}

func (*DescPrimitiveUintptr) Convert(reflect.Value, reflect.Type) *reflect.Value {
	defer util.DiscardPanic()

	if v.Kind() == reflect.String {
		i, err := strconv.Atoi(v.String())
		if err != nil {
			return nil
		}

		res := reflect.ValueOf(uintptr(i))

		return &res
	}

	i, ok := util.GetUint(b)
	if !ok {
		return nil
	}

	res := reflect.ValueOf(uintptr(i))

	return &res

}

func (*DescPrimitiveUintptr) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var num uint64

	cnt, err := util.DecodeNum(r, &num)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(uintptr(num))

	return &res, cnt, nil
}
