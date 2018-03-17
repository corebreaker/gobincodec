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

type DescArrayTime struct{ base.DescBase }

func (*DescArrayTime) Encode(spec base.ISpec, w io.Writer, v reflect.Value) error {
	count := v.Len()

	var out bytes.Buffer

	if err := util.EncodeSize(&out, count); err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		if err := timeCodec.Encode(spec, &out, v.Index(i)); err != nil {
			return err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayTime) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, error) {
	count, err := util.DecodeSize(r)
	if err != nil {
		return nil, err
	}

	res := reflect.New(reflect.ArrayOf(count, reflect.TypeOf((*time.Time)(nil)).Elem())).Elem()
	if count == 0 {
		return &res, nil
	}

	for i := 0; i < count; i++ {
		val, err := timeCodec.Decode(spec, r)
		if err != nil {
			return nil, err
		}

		res.Index(i).Set(*val)
	}

	return &res, nil
}
