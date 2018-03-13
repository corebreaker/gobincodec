package bincodec

import "reflect"

type tCodecInt32 tTypeId

func (tCodecInt32) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecInt32) DecodeValue(dec Decoder) (reflect.Value, error) {
    return nil_value, nil
}
