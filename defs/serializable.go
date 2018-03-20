package defs

type Serializable interface {
	Serialize(Encoder) error
	Deserialize(Decoder) error
}
