package bincodec

import "reflect"

type tCodecUint64 tTypeId

func (tCodecUint64) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecUint64) DecodeValue(dec Decoder) (reflect.Value, error) {
    return nil_value, nil
}
