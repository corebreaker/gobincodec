package gobincodec

import (
	"reflect"
	//"github.com/corebreaker/gobincodec/desc"
	//"github.com/corebreaker/gobincodec/desc/base"
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

type tCodecBase struct {
	spec *tCodecSpec
}

func (self *tCodecBase) Reset() {
	self.spec = newSpec()
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
	res := &tCodecEntry{
		name:     name,
		datatype: datatype,
		codec:    self.compileType(datatype),
	}

	self.UnRegisterName(name)
	self.UnRegisterType(datatype)

	self.spec.forward_registry[name] = res
	self.spec.backward_registry[datatype] = res
	self.spec.identifiers[res.codec.GetId()] = res
}

func (self *tCodecBase) UnRegister(value interface{}) string {
	return self.UnRegisterType(reflect.TypeOf(value))
}

func (self *tCodecBase) unregister(entry *tCodecEntry) {
	delete(self.spec.forward_registry, entry.name)
	delete(self.spec.backward_registry, entry.datatype)
	delete(self.spec.identifiers, entry.codec.GetId())
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
