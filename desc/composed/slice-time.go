package composed

import (
	"io"
	"reflect"
	"time"

	"github.com/corebreaker/gobincodec/desc/base"
	"github.com/corebreaker/gobincodec/util"
)

type DescSliceTime struct{ DescArrayTime }

func (ds *DescSliceTime) TypeEquals(reflect.Type) bool {

}

func (ds *DescSliceTime) Convert(reflect.Value, reflect.Type) *reflect.Value {

}

func (ds *DescSliceTime) Encode(spec base.ISpec, w io.Writer, v reflect.Value) (int, error) {
	if util.IsNil(v) {
		return util.WriteBool(w, true)
	}

	cnt1, err := util.WriteBool(w, false)
	if err != nil {
		return 0, err
	}

	cnt2, err := ds.DescArrayTime.Encode(spec, w, v)
	if err != nil {
		return 0, err
	}

	return cnt1 + cnt2, nil
}

func (*DescSliceTime) Decode(spec base.ISpec, r io.Reader) (*reflect.Value, int, error) {
	isNil, cnt1, err := util.ReadBool(r)
	if err != nil {
		return nil, 0, err
	}

	if isNil {
		var val []time.Time

		res := reflect.ValueOf(&val).Elem()

		return &res, cnt1, nil
	}

	count, cnt2, err := util.DecodeSize(r)
	if err != nil {
		return nil, 0, err
	}

	cnt := cnt1 + cnt2
	buf := make([]time.Time, count)

	for i := 0; i < count; i++ {
		val, n, err := timeCodec.Decode(spec, r)
		if err != nil {
			return nil, 0, err
		}

		cnt += n
		buf[i] = val.Interface().(time.Time)
	}

	res := reflect.ValueOf(buf)

	return &res, cnt, nil
}
