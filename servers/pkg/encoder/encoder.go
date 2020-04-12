package encoder

// Encoder is the interface that implemented by the encoding function.
type Encoder interface {
	Encode(v interface{}) ([]byte, error)
}
