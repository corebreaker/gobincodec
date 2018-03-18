package gobincodec

func gerr(i int, err error) error                               { return err }
func gerr(r interface{}, i int, err error) (interface{}, error) { return r, err }
