package simple

import (
	"bytes"
	"io"
	"reflect"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

var strCodec = new(DescPrimitiveBool)

type DescSliceString struct{ base.DescBase }

func (*DescSliceString) Encode(w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := strCodec.Encode(&out, v.Index(i)); err != nil {
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

func (*DescSliceString) Decode(r io.Reader) (*reflect.Value, error) {
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

	in := make([]string, count)

	for i := 0; i < size; i++ {
		val, err := strCodec.Decode(buf)
		if err != nil {
			return nil, err
		}

		in[i] = val.String()
	}

	res := reflect.ValueOf(in)

	return &res, nil
}
