package gobincodec

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

func (self *tSeekedCodec) next2(r interface{}, n int, err error) (interface{}, error) {
	if err != nil {
		return err
	}

	self.p += int64(n)

	return r, nil
}
