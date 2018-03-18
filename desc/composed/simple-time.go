package composed

import (
	"io"
	"reflect"
	"time"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSimpleTime struct{ base.DescBase }

func (*DescSimpleTime) Encode(_ base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	value := v.Interface().(time.Time)
	buf, err := util.MarshallTime(&value)
	if err != nil {
		return 0, err
	}

	res := append(buf, 0)
	copy(res[1:], res)
	res[0] = byte(len(buf) - 1)

	return util.Write(w, res)
}

func (*DescSimpleTime) Decode(_ base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	var size_buf [1]byte

	cnt1, err := util.Read(r, size_buf[:])
	if err != nil {
		return nil, 0, err
	}

	buf := make([]byte, size_buf[0])

	cnt2, err := util.Read(r, buf)
	if err != nil {
		return nil, 0, err
	}

	var t time.Time

	if err := util.UnmarshalTime(buf, &t); err != nil {
		return nil, 0, err
	}

	res := reflect.ValueOf(t)

	return &res, cnt1 + cnt2, nil
}
