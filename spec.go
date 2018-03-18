package gobincodec

import (
	"io"
	"reflect"
)

type tCodecSpec struct {
	next              uint16
	identifiers       map[uint16]*tCodecEntry
	forward_registry  map[string]*tCodecEntry
	backward_registry map[reflect.Type]*tCodecEntry
}

func (self *tCodecSpec) clone() *tCodecSpec {
	identifiers := make(map[uint16]*tCodecEntry)
	forward_registry := make(map[string]*tCodecEntry)
	backward_registry := make(map[reflect.Type]*tCodecEntry)

	for _, entry := range self.identifiers {
		res := entry.clone()

		identifiers[res.id] = res
		forward_registry[res.name] = res
		backward_registry[res.datatype] = res
	}

	return &tCodecSpec{
		next:              self.next,
		identifiers:       identifiers,
		forward_registry:  forward_registry,
		backward_registry: backward_registry,
	}
}

func (self *tCodecSpec) read(r io.Reader) error {
	return nil
}

func (self *tCodecSpec) readAt(r io.ReaderAt, offs int64) error {
	return nil
}

func (self *tCodecSpec) write(w io.Writer) error {
	return nil
}

func (self *tCodecSpec) writeAt(w io.WriterAt, offs int64) error {
	return nil
}

func newSpec() *tCodecSpec {
	return &tCodecSpec{
		identifiers:       make(map[uint16]*tCodecEntry),
		forward_registry:  make(map[string]*tCodecEntry),
		backward_registry: make(map[reflect.Type]*tCodecEntry),
	}
}
