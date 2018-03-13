package bincodec

import "reflect"

type tCodecInt16 tTypeId

func (tCodecInt16) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecInt16) DecodeValue(dec Decoder) (reflect.Value, error) {
    return nil_value, nil
}
