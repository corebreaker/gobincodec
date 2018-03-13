package bincodec

import "reflect"

type tCodecInt8 tTypeId

func (tCodecInt8) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecInt8) DecodeValue(dec Decoder) (reflect.Value, error) {
    return nil_value, nil
}
