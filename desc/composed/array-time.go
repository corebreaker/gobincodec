package composed

import (
	"bytes"
	"io"
	"reflect"
	"time"

	"github.com/corebreaker/gobincodec/util"
)

type DescArrayTime struct{ DescSliceTime }

func (*DescArrayTime) Decode(r io.Reader) (*reflect.Value, error) {
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

	res := reflect.New(reflect.ArrayOf(count, reflect.TypeOf((*time.Time)(nil)).Elem())).Elem()

	for i := 0; i < count; i++ {
		val, err := timeCodec.Decode(buf)
		if err != nil {
			return nil, err
		}

		res.Index(i).Set(*val)
	}

	return &res, nil
}
