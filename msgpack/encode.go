package msgpack

import (
	"reflect"
	"sync"
)

// Marshaler is the interface implemented by objects that can encode themselves
// to a MessagePack stream.
type Marshaler interface {
	MarshalMsgPack(e *Encoder) error
}

type encodeTypeError struct {
	Type reflect.Type
}

func (e *encodeTypeError) Error() string { _ = "STUB: not implemented"; return "" }

func encodeUnsupportedType(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

// Encode writes the MessagePack encoding of v to the stream.
//
// Encode traverses the value v recursively. If an encountered value implements
// the Marshaler interface Encode calls its MarshalMsgPack method to write the
// value to the stream.
//
// Otherwise, Encode uses the following type-dependent default encodings:
//
//	Go Type             MessagePack Type
//	bool                true or false
//	float32, float64    float64
//	string              string
//	[]byte              binary
//	slices, arrays      array
//	struct, map         map
//
// Struct values encode as maps or arrays. If any struct field tag specifies
// the "array" option, then the struct is encoded as an array. Otherwise, the
// struct is encoded as a map.  Each exported struct field becomes a member of
// the map unless
//   - the field's tag is "-", or
//   - the field is empty and its tag specifies the "omitempty" option.
//
// Anonymous struct fields are marshaled as if their inner exported fields
// were fields in the outer struct.
//
// The struct field tag "empty" specifies a default value when decoding and the
// empty value for the "omitempty" option.
//
// Pointer values encode as the value pointed to. A nil pointer encodes as the
// MessagePack nil value.
//
// Interface values encode as the value contained in the interface. A nil
// interface value encodes as the MessagePack nil value.
func (e *Encoder) Encode(v any) (err error) { _ = "STUB: not implemented"; return nil }

type encodeFunc func(e *Encoder, v reflect.Value)

type encodeBuilder struct {
	m map[reflect.Type]encodeFunc
}

var encodeFuncCache struct {
	sync.RWMutex
	m map[reflect.Type]encodeFunc
}

func encoderForType(t reflect.Type, b *encodeBuilder) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

// Add temporary entry to break recursion.

func (b *encodeBuilder) encoder(t reflect.Type) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func nilEncoder(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func boolEncoder(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func intEncoder(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func uintEncoder(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func floatEncoder(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func stringEncoder(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func byteSliceEncoder(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func interfaceEncoder(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

type ptrEncoder struct{ elem encodeFunc }

func (enc ptrEncoder) encode(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func (b *encodeBuilder) ptrEncoder(t reflect.Type) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

type mapEncoder struct{ key, elem encodeFunc }

func (enc *mapEncoder) encode(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func (b *encodeBuilder) mapEncoder(t reflect.Type) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

type sliceArrayEncoder struct{ elem encodeFunc }

func (enc sliceArrayEncoder) encodeArray(e *Encoder, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (b *encodeBuilder) arrayEncoder(t reflect.Type) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func (enc sliceArrayEncoder) encodeSlice(e *Encoder, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (b *encodeBuilder) sliceEncoder(t reflect.Type) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

var marshalerType = reflect.TypeOf((*Marshaler)(nil)).Elem()

func marshalPtrEncoder(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func marshalEncoder(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func (b *encodeBuilder) marshalEncoder(t reflect.Type) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

type marshalAddrEncoder struct{ f encodeFunc }

func (enc marshalAddrEncoder) encode(e *Encoder, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

type fieldEnc struct {
	name  string
	empty func(reflect.Value) bool
	f     encodeFunc
	index []int
}

type structEncoder []*fieldEnc

func (enc structEncoder) encode(e *Encoder, v reflect.Value) { _ = "STUB: not implemented"; return }

func (enc structEncoder) encodeArray(e *Encoder, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (b *encodeBuilder) structEncoder(t reflect.Type) encodeFunc {
	_ = "STUB: not implemented"
	return *new(encodeFunc)
}

func emptyFunc(f *field) func(reflect.Value) bool { _ = "STUB: not implemented"; return nil }

func lenEmpty(v reflect.Value) bool   { _ = "STUB: not implemented"; return false }
func boolEmpty(v reflect.Value) bool  { _ = "STUB: not implemented"; return false }
func intEmpty(v reflect.Value) bool   { _ = "STUB: not implemented"; return false }
func uintEmpty(v reflect.Value) bool  { _ = "STUB: not implemented"; return false }
func floatEmpty(v reflect.Value) bool { _ = "STUB: not implemented"; return false }
func nilEmpty(v reflect.Value) bool   { _ = "STUB: not implemented"; return false }
