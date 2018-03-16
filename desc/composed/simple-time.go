package composed

import (
	"io"
	"reflect"
	"time"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSimpleTime struct{ base.DescBase }

func (*DescSimpleTime) Encode(_ base.ISpec, w io.Writer, v reflect.Value) error {
	value := v.Interface().(time.Time)
	buf, err := util.MarshallTime(&value)
	if err != nil {
		return err
	}

	res := append(buf, 0)
	copy(res[1:], res)
	res[0] = byte(len(buf) - 1)

	return util.Write(w, res)
}

func (*DescSimpleTime) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, error) {
	var size_buf [1]byte

	if err := util.Read(r, size_buf[:]); err != nil {
		return nil, err
	}

	buf := make([]byte, size_buf[0])
	if err := util.Read(r, buf); err != nil {
		return nil, err
	}

	var t time.Time

	if err := util.UnmarshalTime(buf, &t); err != nil {
		return nil, err
	}

	res := reflect.ValueOf(t)

	return &res, nil
}
