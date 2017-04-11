package bincodec

import "reflect"

type tCodecUnsafePtr tTypeId

func (tCodecUnsafePtr) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecUnsafePtr) DecodeValue(dec Decoder) (reflect.Value, error) {
    return reflect.ValueOf(nil), nil
}
