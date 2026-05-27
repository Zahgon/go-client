package msgpack

import (
	"errors"
	"io"
)

var (
	// ErrLongStringOrBinary is the long string or binary error.
	ErrLongStringOrBinary = errors.New("msgpack: long string or binary")

	// ErrIllegalSize is the illegal array or map size error.
	ErrIllegalSize = errors.New("msgpack: illegal array or map size")
)

// Encoder writes values in MessagePack format.
type Encoder struct {
	buf         [32]byte
	w           io.Writer
	writeString func(string) (int, error)
	err         error // permanent error
}

// NewEncoder allocates and initializes a new Unpacker.
func NewEncoder(w io.Writer) *Encoder { _ = "STUB: not implemented"; return nil }

func (e *Encoder) writeStringUnopt(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type numCodes struct {
	c8  byte
	c16 byte
	c32 byte
	c64 byte
}

var (
	stringLenEncodings = &numCodes{
		c8:  string8Code,
		c16: string16Code,
		c32: string32Code,
		c64: 0,
	}
	binaryLenEncodings = &numCodes{
		c8:  binary8Code,
		c16: binary16Code,
		c32: binary32Code,
		c64: 0,
	}
	arrayLenEncodings = &numCodes{
		c8:  0,
		c16: array16Code,
		c32: array32Code,
		c64: 0,
	}
	mapLenEncodings = &numCodes{
		c8:  0,
		c16: map16Code,
		c32: map32Code,
		c64: 0,
	}
	extLenEncodings = &numCodes{
		c8:  ext8Code,
		c16: ext16Code,
		c32: ext32Code,
		c64: 0,
	}
	uintEncodings = &numCodes{
		c8:  uint8Code,
		c16: uint16Code,
		c32: uint32Code,
		c64: uint64Code,
	}
)

func (e *Encoder) encodeNum(fc *numCodes, v uint64) []byte { _ = "STUB: not implemented"; return nil }

// PackBool writes a Bool value to the MessagePack stream.
func (e *Encoder) PackBool(b bool) error { _ = "STUB: not implemented"; return nil }

// PackInt packs an Int value to the MessagePack stream.
func (e *Encoder) PackInt(v int64) error { _ = "STUB: not implemented"; return nil }

// Pack as unsigned for compatibility with other encoders.

// PackUint packs a Uint value to the message pack stream.
func (e *Encoder) PackUint(v uint64) error { _ = "STUB: not implemented"; return nil }

// Pack as signed for compatibility with other encoders.

// PackFloat writes a Float value to the MessagePack stream.
func (e *Encoder) PackFloat(f float64) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) packStringLen(n int64) error { _ = "STUB: not implemented"; return nil }

// PackString writes a String value to the MessagePack stream.
func (e *Encoder) PackString(v string) error { _ = "STUB: not implemented"; return nil }

// PackStringBytes writes a String value to the MessagePack stream.
func (e *Encoder) PackStringBytes(v []byte) error { _ = "STUB: not implemented"; return nil }

// PackBinary writes a Binary value to the MessagePack stream.
func (e *Encoder) PackBinary(v []byte) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) packArrayMapLen(fixMin int64, fc *numCodes, v int64) error {
	_ = "STUB: not implemented"
	return nil
}

// PackArrayLen write an Array length to the MessagePack stream. The
// application must write n objects to the stream following this call.
func (e *Encoder) PackArrayLen(n int64) error { _ = "STUB: not implemented"; return nil }

// PackMapLen write an Map length to the MessagePack stream. The application
// must write n key-value pairs to the stream following this call.
func (e *Encoder) PackMapLen(n int64) error { _ = "STUB: not implemented"; return nil }

// PackExtension writes an extension to the MessagePack stream.
func (e *Encoder) PackExtension(kind int, data []byte) error { _ = "STUB: not implemented"; return nil }

// PackNil writes a Nil value to the MessagePack stream.
func (e *Encoder) PackNil() error { _ = "STUB: not implemented"; return nil }

// PackRaw writes bytes directly to the MessagePack stream. It is the
// application's responsibility to ensure that the bytes are valid.
func (e *Encoder) PackRaw(p []byte) error { _ = "STUB: not implemented"; return nil }
