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

func (*DescArrayTime) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	count := v.Len()

	var out bytes.Buffer

	if _, err := util.EncodeSize(&out, count); err != nil {
		return 0, err
	}

	for i := 0; i < count; i++ {
		if _, err := timeCodec.Encode(spec, &out, v.Index(i)); err != nil {
			return 0, err
		}
	}

	return util.Write(w, out.Bytes())
}

func (*DescArrayTime) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	size, cnt, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	res := reflect.New(reflect.ArrayOf(size, reflect.TypeOf((*time.Time)(nil)).Elem())).Elem()
	if size == 0 {
		return &res, cnt, nil
	}

	for i := 0; i < size; i++ {
		val, n, err := timeCodec.Decode(spec, r)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		res.Index(i).Set(*val)
	}

	return &res, cnt, nil
}
