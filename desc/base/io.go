package base

import (
	"io"

	"github.com/corebreaker/gobincodec/util"
)

func ReadDescId(r io.Reader) (DescId, int, error) {
	var descid DescId

	n, err := util.DecodeNum(r, &descid)
	if err != nil {
		return 0, 0, err
	}

	return descid, n, nil
}

func WriteDescId(w io.Writer, id DescId) (int, error) {
	return util.EncodeNum(w, id)
}

func ReadDescFromSpec(r io.Reader, spec ISpec) (IDesc, int, error) {
	id, n, err := ReadDescId(r)
	if err != nil {
		return nil, 0, err
	}

	return spec.DescFromId(id), n, nil
}

func WriteIdFromDesc(w io.Writer, desc IDesc) (int, error) {
	return util.EncodeNum(w, desc.GetId())
}
