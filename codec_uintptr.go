package bincodec

import "reflect"

type tCodecUintptr tTypeId

func (tCodecUintptr) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecUintptr) DecodeValue(dec Decoder) (reflect.Value, error) {
    return reflect.ValueOf(nil), nil
}
