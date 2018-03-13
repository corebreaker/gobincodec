package bincodec

type tSeekedCodec struct {
    p int64
}

func (self *tSeekedCodec) next(n int, err error) error {
    if err != nil {
        return err
    }

    self.p += int64(n)

    return nil
}
