package bincodec

import (
    "io"
    "reflect"
)

var (
    nil_value = reflect.ValueOf(nil)
)

type Serializable interface {
    Serialize(Encoder) error
    Deserialize(Decoder) error
}

type CodecBase interface {
    GetType(name string) reflect.Type
    GetNameFromType(datatype reflect.Type) string
    GetName(value interface{}) string
    HasName(name string) bool
    HasValue(value interface{}) bool
    RegisterType(datatype reflect.Type)
    Register(value interface{})
    RegisterName(name string, value interface{})
    RegisterTypeName(name string, datatype reflect.Type)
    UnRegister(value interface{}) string
    UnRegisterType(datatype reflect.Type) string
    UnRegisterName(name string) reflect.Type
    Reset()
}

type tCodecEntry struct {
    id       uint16
    name     string
    datatype reflect.Type
    codec    iCodec
}

func (self *tCodecEntry) clone() *tCodecEntry {
    res := new(tCodecEntry)
    *res = *self

    return res
}

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

func (self *tCodecSpec) read_at(r io.ReaderAt, offs int64) error {
    return nil
}

func (self *tCodecSpec) write(w io.Writer) error {
    return nil
}

func (self *tCodecSpec) write_at(w io.WriterAt, offs int64) error {
    return nil
}

func new_spec() *tCodecSpec {
    return &tCodecSpec{
        identifiers:       make(map[uint16]*tCodecEntry),
        forward_registry:  make(map[string]*tCodecEntry),
        backward_registry: make(map[reflect.Type]*tCodecEntry),
    }
}

type tCodecBase struct {
    spec *tCodecSpec
}

func (self *tCodecBase) Reset() {
    self.spec = new_spec()
}

func (self *tCodecBase) GetType(name string) reflect.Type {
    res, ok := self.spec.forward_registry[name]
    if !ok {
        return nil
    }

    return res.datatype
}

func (self *tCodecBase) GetNameFromType(datatype reflect.Type) string {
    res, ok := self.spec.backward_registry[datatype]
    if !ok {
        return ""
    }

    return res.name
}

func (self *tCodecBase) GetName(value interface{}) string {
    return self.GetNameFromType(reflect.TypeOf(value))
}

func (self *tCodecBase) HasName(name string) bool {
    _, res := self.spec.forward_registry[name]

    return res
}

func (self *tCodecBase) HasValue(value interface{}) bool {
    _, res := self.spec.backward_registry[reflect.TypeOf(value)]

    return res
}

func (self *tCodecBase) register(name string, datatype reflect.Type) {
}

func (self *tCodecBase) RegisterType(datatype reflect.Type) {
    self.RegisterTypeName(datatype.String(), datatype)
}

func (self *tCodecBase) Register(value interface{}) {
    self.RegisterType(reflect.TypeOf(value))
}

func (self *tCodecBase) RegisterName(name string, value interface{}) {
    self.RegisterTypeName(name, reflect.TypeOf(value))
}

func (self *tCodecBase) RegisterTypeName(name string, datatype reflect.Type) {
    id := self.spec.next
    self.spec.next++

    res := &tCodecEntry{
        id:       id,
        name:     name,
        datatype: datatype,
        codec:    self.compile_type(datatype),
    }

    self.UnRegisterName(name)
    self.UnRegisterType(datatype)

    self.spec.forward_registry[name] = res
    self.spec.backward_registry[datatype] = res
    self.spec.identifiers[id] = res
}

func (self *tCodecBase) UnRegister(value interface{}) string {
    return self.UnRegisterType(reflect.TypeOf(value))
}

func (self *tCodecBase) unregister(entry *tCodecEntry) {
    delete(self.spec.forward_registry, entry.name)
    delete(self.spec.backward_registry, entry.datatype)
    delete(self.spec.identifiers, entry.id)
}

func (self *tCodecBase) UnRegisterType(datatype reflect.Type) string {
    res, ok := self.spec.backward_registry[datatype]
    if !ok {
        return ""
    }

    self.unregister(res)

    return res.name
}

func (self *tCodecBase) UnRegisterName(name string) reflect.Type {
    res, ok := self.spec.forward_registry[name]
    if !ok {
        return nil
    }

    self.unregister(res)

    return res.datatype
}
