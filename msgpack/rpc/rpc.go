// Package rpc implements MessagePack RPC.
package rpc

import (
	"bufio"
	"errors"
	"io"
	"reflect"
	"sync"

	"github.com/neovim/go-client/msgpack"
)

// kind represents a MessagePack RPC message kind.
type kind int

// list of kind.
const (
	requestMessage      kind = 0
	replyMessage        kind = 1
	notificationMessage kind = 2
)

// state represents a MessagePack RPC state.
type state int

// list of state.
const (
	stateInit state = iota
	stateClosed
)

var (
	// ErrClosed session closed error.
	ErrClosed = errors.New("msgpack/rpc: session closed")

	// ErrInternal msgpack-rpc internal error.
	ErrInternal = errors.New("msgpack/rpc: internal error")

	// ErrHandlerNotFunction handler type is not a function error.
	ErrHandlerNotFunction = errors.New("msgpack/rpc: handler not a function")

	// ErrInvalidHandlerReturn invalid handler function return type error.
	ErrInvalidHandlerReturn = errors.New("msgpack/rpc: handler return must be (), (error) or (valueType, error)")

	// ErrInvalidArgument invalid argument error.
	ErrInvalidArgument = errors.New("msgpack/rpc: invalid argument")
)

// Error represents a MessagePack RPC error.
type Error struct {
	Value any
}

// Error implements the error interface.
func (e Error) Error() string { _ = "STUB: not implemented"; return "" }

// Call represents a MessagePack RPC call.
type Call struct {
	Args   any
	Reply  any
	Err    error
	Done   chan *Call
	Method string
}

func (c *Call) done(e *Endpoint, err error) { _ = "STUB: not implemented"; return }

// ok

type handler struct {
	fn   reflect.Value
	args []reflect.Value
}

type notification struct {
	call   func([]reflect.Value) []reflect.Value
	next   *notification
	method string
	args   []reflect.Value
}

// Endpoint represents a MessagePack RPC peer.
type Endpoint struct {
	err  error
	logf func(fmt string, args ...any)

	done   chan struct{}
	closer io.Closer
	bw     *bufio.Writer
	enc    *msgpack.Encoder
	dec    *msgpack.Decoder

	handlers          map[string]*handler
	pending           map[uint64]*Call
	notificationsCond *sync.Cond

	arg           reflect.Value
	notifications []*notification
	state         state
	id            uint64

	mu              sync.Mutex
	handlersMu      sync.RWMutex
	encMu           sync.Mutex
	notificationsMu sync.Mutex
}

// Option is a configures a Endpoint.
type Option struct{ f func(*Endpoint) }

// WithExtensions configures Endpoint to define application-specific types.
func WithExtensions(extensions msgpack.ExtensionMap) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLogf sets the log function to Endpoint.
func WithLogf(f func(fmt string, args ...any)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// NewEndpoint returns a new endpoint with the specified options.
func NewEndpoint(r io.Reader, w io.Writer, c io.Closer, options ...Option) (*Endpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Endpoint) decodeUint(what string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *Endpoint) decodeString(what string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *Endpoint) skip(n int) error { _ = "STUB: not implemented"; return nil }

// Serve serves incoming requests. Serve blocks until the peer disconnects or
// there is an error.
func (e *Endpoint) Serve() error { _ = "STUB: not implemented"; return nil }

func (e *Endpoint) close(err error) error { _ = "STUB: not implemented"; return nil }

// Close releases the resources used by endpoint.
func (e *Endpoint) Close() error { _ = "STUB: not implemented"; return nil }

var errorType = reflect.ValueOf(new(error)).Elem().Type()

// Register registers handler fn for the specified method name.
//
// When servicing a call, the arguments to fn are the values in args followed
// by the values passed from the peer.
func (e *Endpoint) Register(method string, fn any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Call invokes the target method and waits for a response.
func (e *Endpoint) Call(method string, reply any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Go append method call to queue and returns the new Call.
func (e *Endpoint) Go(method string, done chan *Call, reply any, args ...any) *Call {
	_ = "STUB: not implemented"
	return nil
}

// Notify invokes the target method with non-blocking.
func (e *Endpoint) Notify(method string, args ...any) error { _ = "STUB: not implemented"; return nil }

func (e *Endpoint) createCall(h *handler) (func([]reflect.Value) []reflect.Value, []reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Decode plain arguments.

// Skip extra arguments

func (e *Endpoint) reply(id uint64, replyErr error, reply any) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Endpoint) handleRequest(messageLen int) error { _ = "STUB: not implemented"; return nil }

// messageType, id, method, args

func (e *Endpoint) handleReply(messageLen int) error { _ = "STUB: not implemented"; return nil }

// messageType, id, error, reply

func (e *Endpoint) handleNotification(messageLen int) error {
	_ = "STUB: not implemented"
	// messageType, method, args
	return nil
}

func (e *Endpoint) enqueNotification(n *notification) { _ = "STUB: not implemented"; return }

func (e *Endpoint) dequeueNotifications() []*notification { _ = "STUB: not implemented"; return nil }

// runNotifications runs notifications in a single goroutine to ensure that the
// notifications are processed in order by the application.
func (e *Endpoint) runNotifications() { _ = "STUB: not implemented"; return }

// Serve() enqueues nil on return
