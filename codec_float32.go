package bincodec

import "reflect"

type tCodecFloat32 tTypeId

func (tCodecFloat32) EncodeValue(enc Encoder, val interface{}) error {
    return nil
}

func (tCodecFloat32) DecodeValue(dec Decoder) (reflect.Value, error) {
    return nil_value, nil
}
