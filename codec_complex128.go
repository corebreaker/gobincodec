package bincodec

import "reflect"

type tCodecComplex128 tTypeId

func (tCodecComplex128) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecComplex128) DecodeValue(dec Decoder) (reflect.Value, error) {
    return reflect.ValueOf(nil), nil
}
