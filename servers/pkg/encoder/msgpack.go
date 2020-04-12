package encoder

import "github.com/ugorji/go/codec"

// MsgpackEncoder provides a encoding function of Msgpack.
type MsgpackEncoder struct {
	handle codec.Handle
}

// NewMsgpackEncoder returns a new MsgpackEncoder.
func NewMsgpackEncoder(h codec.Handle) *MsgpackEncoder {
	return &MsgpackEncoder{handle: h}
}

// Encode encodes a object `v` into bytes of Msgpack.
func (m *MsgpackEncoder) Encode(v interface{}) ([]byte, error) {
	var buf = []byte{}
	err := codec.NewEncoderBytes(&buf, m.handle).Encode(v)

	return buf, err
}
