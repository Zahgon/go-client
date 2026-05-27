package msgpack

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
)

// Type represents the type of value in the MessagePack stream.
type Type int

// list of MessagePack types.
const (
	Invalid Type = iota
	Nil
	Bool
	Int
	Uint
	Float
	ArrayLen
	MapLen
	String
	Binary
	Extension
)

var typeNames = [...]string{
	Invalid:   "Invalid",
	Nil:       "Nil",
	Bool:      "Bool",
	Int:       "Int",
	Uint:      "Uint",
	Float:     "Float",
	ArrayLen:  "ArrayLen",
	MapLen:    "MapLen",
	String:    "String",
	Binary:    "Binary",
	Extension: "Extension",
}

// String returns a string representation of the Type.
func (t Type) String() string { _ = "STUB: not implemented"; return "" }

// ErrDataSizeTooLarge is the data size too large error.
var ErrDataSizeTooLarge = errors.New("msgpack: data size too large")

// Decoder reads MessagePack objects from an io.Reader.
type Decoder struct {
	extensions ExtensionMap
	err        error
	r          *bufio.Reader
	n          uint64
	p          []byte
	t          Type
	peek       bool
}

const bufioReaderSize = 4096

// NewDecoder allocates and initializes a new decoder.
func NewDecoder(r io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

// ExtensionMap specifies functions for converting MessagePack extensions to Go
// values.
//
// The key is the MessagePack extension type.
// The value is a function that converts the extension data to a Go value.
type ExtensionMap map[int]func([]byte) (any, error)

// SetExtensions specifies functions for converting MessagePack extensions to Go
// values.
func (d *Decoder) SetExtensions(extensions ExtensionMap) { _ = "STUB: not implemented"; return }

// Type returns the type of the current value in the stream.
func (d *Decoder) Type() Type {
	_ = "STUB: not implemented"

	// Extension returns the type of the current Extension value.
	return *new(Type)
}

func (d *Decoder) Extension() int {
	_ = "STUB: not implemented"

	// Bytes returns the current String, Binary or Extension value as a slice of
	// bytes.
	return 0
}

func (d *Decoder) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// BytesNoCopy returns the current String, Binary or Extension value as a slice
// of bytes. The underlying array may point to data that will be overwritten by
// a subsequent call to Unpack.
func (d *Decoder) BytesNoCopy() []byte {
	_ = "STUB: not implemented"

	// String returns the current String, Binary or Extension value as a string.
	return nil
}

func (d *Decoder) String() string {
	_ = "STUB: not implemented"

	// Int returns the current Int value.
	return ""
}

func (d *Decoder) Int() int64 {
	_ = "STUB: not implemented"

	// Uint returns the current Uint value.
	return 0
}

func (d *Decoder) Uint() uint64 {
	_ = "STUB: not implemented"

	// Len returns the current ArrayLen or MapLen value.
	return 0
}

func (d *Decoder) Len() int {
	_ = "STUB: not implemented"

	// Bool returns the current Bool value.
	return 0
}

func (d *Decoder) Bool() bool { _ = "STUB: not implemented"; return false }

// Float returns the current Float value.
func (d *Decoder) Float() float64 { _ = "STUB: not implemented"; return 0 }

// Unpack reads the next value from the MessagePack stream. Call Type to get the
// type of the current value. Call Bool, Uint, Int, Float, Bytes or Extension
// to get the value.
func (d *Decoder) Unpack() error { _ = "STUB: not implemented"; return nil }

// Don't call d.fatal here because we don't want io.EOF converted to
// io.ErrUnexpectedEOF

// Skip skips over any nested values in the stream.
func (d *Decoder) Skip() error { _ = "STUB: not implemented"; return nil }

func (d *Decoder) skipCount() int { _ = "STUB: not implemented"; return 0 }

var formats = [256]*struct {
	t    Type
	n    func(d *Decoder, code byte) (uint64, error)
	more bool
}{
	fixIntCodeMin: {
		t: Int,
		n: func(d *Decoder, code byte) (uint64, error) { return uint64(code), nil },
	},
	fixMapCodeMin: {
		t: MapLen,
		n: func(d *Decoder, code byte) (uint64, error) { return uint64(code) - uint64(fixMapCodeMin), nil },
	},
	fixArrayCodeMin: {
		t: ArrayLen,
		n: func(d *Decoder, code byte) (uint64, error) { return uint64(code) - uint64(fixArrayCodeMin), nil },
	},
	fixStringCodeMin: {
		t:    String,
		n:    func(d *Decoder, code byte) (uint64, error) { return uint64(code) - uint64(fixStringCodeMin), nil },
		more: true,
	},
	negFixIntCodeMin: {
		t: Int,
		n: func(d *Decoder, code byte) (uint64, error) { return uint64(int64(int8(code))), nil },
	},
	nilCode: {
		t: Nil,
		n: func(d *Decoder, code byte) (uint64, error) { return 0, nil },
	},
	falseCode: {
		t: Bool,
		n: func(d *Decoder, code byte) (uint64, error) { return 0, nil },
	},
	trueCode: {
		t: Bool,
		n: func(d *Decoder, code byte) (uint64, error) { return 1, nil },
	},
	float32Code: {
		t: Float,
		n: func(d *Decoder, code byte) (uint64, error) {
			n, err := d.read4(code)
			return math.Float64bits(float64(math.Float32frombits(uint32(n)))), err
		},
	},
	float64Code: {
		t: Float,
		n: (*Decoder).read8,
	},
	uint8Code: {
		t: Uint,
		n: (*Decoder).read1,
	},
	uint16Code: {
		t: Uint,
		n: (*Decoder).read2,
	},
	uint32Code: {
		t: Uint,
		n: (*Decoder).read4,
	},
	uint64Code: {
		t: Uint,
		n: (*Decoder).read8,
	},
	int8Code: {
		t: Int,
		n: func(d *Decoder, code byte) (uint64, error) {
			n, err := d.read1(code)
			return uint64(int64(int8(n))), err
		},
	},
	int16Code: {
		t: Int,
		n: func(d *Decoder, code byte) (uint64, error) {
			n, err := d.read2(code)
			return uint64(int64(int16(n))), err
		},
	},
	int32Code: {
		t: Int,
		n: func(d *Decoder, code byte) (uint64, error) {
			n, err := d.read4(code)
			return uint64(int64(int32(n))), err
		},
	},
	int64Code: {
		t: Int,
		n: (*Decoder).read8,
	},
	string8Code: {
		t:    String,
		n:    (*Decoder).read1,
		more: true,
	},
	string16Code: {
		t:    String,
		n:    (*Decoder).read2,
		more: true,
	},
	string32Code: {
		t:    String,
		n:    (*Decoder).read4,
		more: true,
	},
	binary8Code: {
		t:    Binary,
		n:    (*Decoder).read1,
		more: true,
	},
	binary16Code: {
		t:    Binary,
		n:    (*Decoder).read2,
		more: true,
	},
	binary32Code: {
		t:    Binary,
		n:    (*Decoder).read4,
		more: true,
	},
	array16Code: {
		t: ArrayLen,
		n: (*Decoder).read2,
	},
	array32Code: {
		t: ArrayLen,
		n: (*Decoder).read4,
	},
	map16Code: {
		t: MapLen,
		n: (*Decoder).read2,
	},
	map32Code: {
		t: MapLen,
		n: (*Decoder).read4,
	},
	fixExt1Code: {
		t:    Extension,
		n:    func(d *Decoder, code byte) (uint64, error) { return 1, nil },
		more: true,
	},
	fixExt2Code: {
		t:    Extension,
		n:    func(d *Decoder, code byte) (uint64, error) { return 2, nil },
		more: true,
	},
	fixExt4Code: {
		t:    Extension,
		n:    func(d *Decoder, code byte) (uint64, error) { return 4, nil },
		more: true,
	},
	fixExt8Code: {
		t:    Extension,
		n:    func(d *Decoder, code byte) (uint64, error) { return 8, nil },
		more: true,
	},
	fixExt16Code: {
		t:    Extension,
		n:    func(d *Decoder, code byte) (uint64, error) { return 16, nil },
		more: true,
	},
	ext8Code: {
		t:    Extension,
		n:    (*Decoder).read1,
		more: true,
	},
	ext16Code: {
		t:    Extension,
		n:    (*Decoder).read2,
		more: true,
	},
	ext32Code: {
		t:    Extension,
		n:    (*Decoder).read4,
		more: true,
	},
	unusedCode: {
		t: Invalid,
		n: func(d *Decoder, code byte) (uint64, error) {
			return 0, fmt.Errorf("msgpack: unknown format code %x", code)
		},
	},
}

func init() {
	for i := fixIntCodeMin + 1; i <= fixIntCodeMax; i++ {
		formats[i] = formats[fixIntCodeMin]
	}

	for i := fixMapCodeMin + 1; i <= fixMapCodeMax; i++ {
		formats[i] = formats[fixMapCodeMin]
	}

	for i := fixArrayCodeMin + 1; i <= fixArrayCodeMax; i++ {
		formats[i] = formats[fixArrayCodeMin]
	}

	for i := fixStringCodeMin + 1; i <= fixStringCodeMax; i++ {
		formats[i] = formats[fixStringCodeMin]
	}

	for i := negFixIntCodeMin + 1; i <= negFixIntCodeMax; i++ {
		formats[i] = formats[negFixIntCodeMin]
	}
}

func (d *Decoder) fatal(err error) error { _ = "STUB: not implemented"; return nil }

func (d *Decoder) read1(format byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *Decoder) read2(format byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *Decoder) read4(format byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *Decoder) read8(format byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
