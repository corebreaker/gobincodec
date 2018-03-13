package bincodec

import "reflect"

type tCodecInt64 tTypeId

func (tCodecInt64) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecInt64) DecodeValue(dec Decoder) (reflect.Value, error) {
    return nil_value, nil
}
