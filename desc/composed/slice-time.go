package composed

import (
	"bytes"
	"io"
	"reflect"
	"time"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

var timeCodec = new(DescSimpleTime)

type DescSliceTime struct{ base.DescBase }

func (*DescSliceTime) Encode(w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := timeCodec.Encode(&out, v.Index(i)); err != nil {
			return err
		}
	}

	var res bytes.Buffer

	if err := util.EncodeSize(&res, out.Len()); err != nil {
		return err
	}

	if err := util.TransferTo(&res, &out); err != nil {
		return err
	}

	return util.Write(w, res.Bytes())
}

func (*DescSliceTime) Decode(r io.Reader) (*reflect.Value, error) {
	size, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	code := make([]byte, size)
	if err := util.Read(r, code); err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(code)

	count, err := util.DecodeSize(buf)
	if err != nil {
		return nil, err
	}

	in := make([]time.Time, count)

	for i := 0; i < size; i++ {
		val, err := timeCodec.Decode(buf)
		if err != nil {
			return nil, err
		}

		in[i] = val.Interface().(time.Time)
	}

	res := reflect.ValueOf(in)

	return &res, nil
}
