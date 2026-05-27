package msgpack

import (
	"errors"
	"reflect"
	"sync"
)

// Unmarshaler is the interface implemented by objects that can decode
// themselves from a MessagePack stream.
type Unmarshaler interface {
	UnmarshalMsgPack(d *Decoder) error
}

// ErrInvalidDecodeArg is the invalid argument error.
var ErrInvalidDecodeArg = errors.New("msgpack: argument to Decode must be non-nil pointer, slice or map")

// DecodeConvertError describes a MessagePack value that was not appropriate
// for a value of a specific Go type.
type DecodeConvertError struct {
	// The MessagePack type of the value.
	SrcType Type
	// Option value.
	SrcValue any
	// Type of the Go value that could not be assigned to.
	DestType reflect.Type
}

// Error implements the error interface.
func (e *DecodeConvertError) Error() string { _ = "STUB: not implemented"; return "" }

func decodeUnsupportedType(ds *decodeState, v reflect.Value) { _ = "STUB: not implemented"; return }

// decodeState represents the state while decoding value.
type decodeState struct {
	*Decoder
	errSaved error
}

func (ds *decodeState) unpack() { _ = "STUB: not implemented"; return }

func (ds *decodeState) skip() { _ = "STUB: not implemented"; return }

func (ds *decodeState) saveErrorAndSkip(destValue reflect.Value, srcValue any) {
	_ = "STUB: not implemented"
	return
}

// Decode decodes the next value in the stream to v.
//
// Decode uses the inverse of the encodings that Encoder.Encode uses,
// allocating maps, slices, and pointers as necessary, with the following
// additional rules:
//
// To decode into a pointer, Decode first handles the case of a MessagePack
// nil. In that case, Decode sets the pointer to nil. Otherwise, Decode decodes
// the stream into the value pointed at by the pointer. If the pointer is nil,
// Decode allocates a new value for it to point to.
//
// To decode a MessagePack array into a slice, Decode sets the slice length to
// the length of the MessagePack array or reallocates the slice if there is
// insufficient capaicity. Slice elments are not cleared before decoding the
// element.
//
// To decode a MessagePack array into a Go array, Decode decodes the
// MessagePack array elements into corresponding Go array elements.  If the Go
// array is smaller than the MessagePack array, the additional MessagePack
// array elements are discarded. If the MessagePack array is smaller than the
// Go array, the additional Go array elements are set to zero values.
//
// If a MessagePack value is not appropriate for a given target type, or if a
// MessagePack number overflows the target type, Decode skips that field and
// completes the decoding as best it can.  If no more serious errors are
// encountered, Decode returns an DecodeConvertError describing the earliest
// such error.
func (d *Decoder) Decode(v any) (err error) { _ = "STUB: not implemented"; return nil }

var decodeFuncCache struct {
	sync.RWMutex
	m map[reflect.Type]decodeFunc
}

type decodeFunc func(*decodeState, reflect.Value)

type decodeBuilder struct {
	m map[reflect.Type]decodeFunc
}

func decoderForType(t reflect.Type, b *decodeBuilder) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

// Add temporary entry to break recursion

func (b *decodeBuilder) decoder(t reflect.Type) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func boolDecoder(ds *decodeState, v reflect.Value) { _ = "STUB: not implemented"; return }

func intDecoder(ds *decodeState, v reflect.Value) { _ = "STUB: not implemented"; return }

func uintDecoder(ds *decodeState, v reflect.Value) { _ = "STUB: not implemented"; return }

func floatDecoder(ds *decodeState, v reflect.Value) { _ = "STUB: not implemented"; return }

func stringDecoder(ds *decodeState, v reflect.Value) { _ = "STUB: not implemented"; return }

func byteSliceDecoder(ds *decodeState, v reflect.Value) { _ = "STUB: not implemented"; return }

// Nothing to do

// TODO: check if OK to set?

func interfaceDecoder(ds *decodeState, v reflect.Value) { _ = "STUB: not implemented"; return }

// We don't know how to make an object of this interface type.

type sliceArrayDecoder struct {
	elem decodeFunc
}

func (dec sliceArrayDecoder) decodeArray(ds *decodeState, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (b *decodeBuilder) arrayDecoder(t reflect.Type) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

func (dec sliceArrayDecoder) decodeSlice(ds *decodeState, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func (b *decodeBuilder) sliceDecoder(t reflect.Type) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

type mapDecoder struct {
	key  decodeFunc
	elem decodeFunc
}

func (dec *mapDecoder) decode(ds *decodeState, v reflect.Value) { _ = "STUB: not implemented"; return }

func (b *decodeBuilder) mapDecoder(t reflect.Type) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

type fieldDec struct {
	index []int
	f     decodeFunc
	empty reflect.Value
}

func (fd *fieldDec) setEmpty(v reflect.Value) { _ = "STUB: not implemented"; return }

type structArrayDecoder []*fieldDec

func (dec structArrayDecoder) decode(ds *decodeState, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

type structDecoder map[string]*fieldDec

func (dec structDecoder) decode(ds *decodeState, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

// Key

// Value

func (b *decodeBuilder) structDecoder(t reflect.Type) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

type ptrDecoder struct {
	elem decodeFunc
}

func (dec ptrDecoder) decode(ds *decodeState, v reflect.Value) { _ = "STUB: not implemented"; return }

func (b *decodeBuilder) ptrDecoder(t reflect.Type) decodeFunc {
	_ = "STUB: not implemented"
	return *new(decodeFunc)
}

var unmarshalerType = reflect.TypeOf((*Unmarshaler)(nil)).Elem()

func unmarshalDecoder(ds *decodeState, v reflect.Value) { _ = "STUB: not implemented"; return }

type unmarshalAddrDecoder struct{ f decodeFunc }

func (dec unmarshalAddrDecoder) decode(ds *decodeState, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

type extensionValue struct {
	kind int
	data []byte
}

func (ev extensionValue) MarshalMsgPack(e *Encoder) error { _ = "STUB: not implemented"; return nil }

func decodeNoReflect(ds *decodeState) (x any) { _ = "STUB: not implemented"; return *new(any) }
