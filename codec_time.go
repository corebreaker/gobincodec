package bincodec

import "reflect"

type tCodecTime tTypeId

func (tCodecTime) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecTime) DecodeValue(dec Decoder) (reflect.Value, error) {
    return reflect.ValueOf(nil), nil
}
