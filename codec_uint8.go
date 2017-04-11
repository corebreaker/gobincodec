package bincodec

import "reflect"

type tCodecUint8 tTypeId

func (tCodecUint8) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecUint8) DecodeValue(dec Decoder) (reflect.Value, error) {
    return reflect.ValueOf(nil), nil
}
