package bincodec

import "reflect"

type tCodecUint tTypeId

func (tCodecUint) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecUint) DecodeValue(dec Decoder) (reflect.Value, error) {
    return reflect.ValueOf(nil), nil
}
