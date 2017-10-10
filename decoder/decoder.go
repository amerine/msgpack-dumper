package decoder

import (
	"io"
)

// FBitDecoder handles decoding of fluent-bit msgpack messages.
type FBitDecoder struct {
}

// NewDecoder takes the provided io.Reader with a messagepack-encoded fluent-bit message
// and returns a pre-configured FBitDecoder.
func NewDecoder(r io.Reader) *FBitDecoder {
	dec := new(FBitDecoder)

	return dec
}

// NewDecoderBytes takes the provides []byte input and returns a preconfigured FBitDecoder.
func NewDecoderBytes(in []byte) *FBitDecoder {
	dec := new(FBitDecoder)
	return dec
}

// GetRecord returns a single messages from the payload.
func GetRecord(dec *FBitDecoder) (ret int, ts interface{}, rec map[interface{}]interface{}) {
	return -1, 0, nil
}
