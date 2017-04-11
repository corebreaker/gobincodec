package bincodec

import "reflect"

type tCodecFloat64 tTypeId

func (tCodecFloat64) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecFloat64) DecodeValue(dec Decoder) (reflect.Value, error) {
    return reflect.ValueOf(nil), nil
}
