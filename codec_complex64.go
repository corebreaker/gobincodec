package bincodec

import "reflect"

type tCodecComplex64 tTypeId

func (tCodecComplex64) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecComplex64) DecodeValue(dec Decoder) (reflect.Value, error) {
    return nil_value, nil
}
