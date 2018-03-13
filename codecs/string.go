package bincodec

import "reflect"

type tCodecString tTypeId

func (tCodecString) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecString) DecodeValue(dec Decoder) (reflect.Value, error) {
    return nil_value, nil
}
