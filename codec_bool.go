package bincodec

import "reflect"

type tCodecBool tTypeId

func (tCodecBool) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecBool) DecodeValue(dec Decoder) (reflect.Value, error) {
    return nil_value, nil
}
