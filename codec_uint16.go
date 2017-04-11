package bincodec

import "reflect"

type tCodecUint16 tTypeId

func (tCodecUint16) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecUint16) DecodeValue(dec Decoder) (reflect.Value, error) {
    return reflect.ValueOf(nil), nil
}
