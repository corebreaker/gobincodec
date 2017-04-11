package bincodec

import "reflect"

type tCodecInt tTypeId

func (tCodecInt) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecInt) DecodeValue(dec Decoder) (reflect.Value, error) {
    return reflect.ValueOf(nil), nil
}
