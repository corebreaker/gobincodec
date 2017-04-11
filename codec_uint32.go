package bincodec

import "reflect"

type tCodecUint32 tTypeId

func (tCodecUint32) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecUint32) DecodeValue(dec Decoder) (reflect.Value, error) {
    return reflect.ValueOf(nil), nil
}
