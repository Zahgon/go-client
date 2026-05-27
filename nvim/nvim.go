package nvim

import (
	"bytes"
	"context"
	"io"
	"net"
	"os/exec"
	"sync"
	"syscall"

	"github.com/neovim/go-client/msgpack"
	"github.com/neovim/go-client/msgpack/rpc"
)

//go:generate go run api_tool.go -generate api.go -deprecated api_deprecated.go

var embedProcAttr *syscall.SysProcAttr

// Nvim represents a remote instance of Nvim. It is safe to call Nvim methods
// concurrently.
type Nvim struct {
	ep *rpc.Endpoint

	// cmd is the child process, if any.
	cmd         *exec.Cmd
	serveCh     chan error
	channelID   int
	channelIDMu sync.Mutex

	// readMu prevents concurrent calls to read on the child process stdout pipe and
	// calls to cmd.Wait().
	readMu sync.Mutex
}

// Serve serves incoming mesages from the peer. Serve blocks until Nvim
// disconnects or there is an error.
//
// By default, the NewChildProcess and Dial functions start a goroutine to run Serve().
// Callers of the low-level New function are responsible for running Serve().
func (v *Nvim) Serve() error { _ = "STUB: not implemented"; return nil }

func (v *Nvim) startServe() { _ = "STUB: not implemented"; return }

// Close releases the resources used the client.
func (v *Nvim) Close() error { _ = "STUB: not implemented"; return nil }

// The child process should exit cleanly on call to v.ep.Close(). Kill
// the process if it does not exit as expected.

// ExitCode returns the exit code of the exited nvim process.
func (v *Nvim) ExitCode() int { _ = "STUB: not implemented"; return 0 }

// New creates an Nvim client. When connecting to Nvim over stdio, use stdin as
// r and stdout as w and c, When connecting to Nvim over a network connection,
// use the connection for r, w and c.
//
// The application must call Serve() to handle RPC requests and responses.
//
//	:help rpc-connecting
func New(r io.Reader, w io.Writer, c io.Closer, logf func(string, ...any)) (*Nvim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ChildProcessOption specifies an option for creating a child process.
type ChildProcessOption struct {
	f func(*childProcessOptions)
}

type childProcessOptions struct {
	ctx          context.Context
	logf         func(string, ...any)
	command      string
	dir          string
	args         []string
	env          []string
	serve        bool
	disableEmbed bool
}

// ChildProcessArgs specifies the command line arguments. The application must
// include the --embed flag or other flags that cause Nvim to use stdin/stdout
// as a MsgPack RPC channel.
func ChildProcessArgs(args ...string) ChildProcessOption {
	_ = "STUB: not implemented"
	return *new(ChildProcessOption)
}

// ChildProcessCommand specifies the command to run. NewChildProcess runs
// "nvim" by default.
func ChildProcessCommand(command string) ChildProcessOption {
	_ = "STUB: not implemented"
	return *new(ChildProcessOption)
}

// ChildProcessContext specifies the context to use when starting the command.
// The background context is used by defaullt.
func ChildProcessContext(ctx context.Context) ChildProcessOption {
	_ = "STUB: not implemented"
	return *new(ChildProcessOption)
}

// ChildProcessDir specifies the working directory for the process. The current
// working directory is used by default.
func ChildProcessDir(dir string) ChildProcessOption {
	_ = "STUB: not implemented"
	return *new(ChildProcessOption)
}

// ChildProcessEnv specifies the environment for the child process. The current
// process environment is used by default.
func ChildProcessEnv(env []string) ChildProcessOption {
	_ = "STUB: not implemented"
	return *new(ChildProcessOption)
}

// ChildProcessServe specifies whether Server should be run in a goroutine.
// The default is to run Serve().
func ChildProcessServe(serve bool) ChildProcessOption {
	_ = "STUB: not implemented"
	return *new(ChildProcessOption)
}

// ChildProcessLogf specifies function for logging output. The log.Printf
// function is used by default.
func ChildProcessLogf(logf func(string, ...any)) ChildProcessOption {
	_ = "STUB: not implemented"
	return *new(ChildProcessOption)
}

// ChildProcessDisableEmbed disables the --embed flag of nvim.
// See: https://neovim.io/doc/user/starting.html#--embed for details.
func ChildProcessDisableEmbed() ChildProcessOption {
	_ = "STUB: not implemented"
	return *new(ChildProcessOption)
}

// appendEmbedFlagIfNeeded appends the --embed flag, if it is not yet added.
// This behavior can be overriden by setting the ChildProcessDisableEmbed() process option.
func appendEmbedFlagIfNeeded(cpos *childProcessOptions) { _ = "STUB: not implemented"; return }

// NewChildProcess returns a client connected to stdin and stdout of a new
// child process.
func NewChildProcess(options ...ChildProcessOption) (*Nvim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DialOption specifies an option for dialing to Nvim.
type DialOption struct {
	f func(*dialOptions)
}

type dialOptions struct {
	ctx     context.Context
	logf    func(string, ...any)
	netDial func(ctx context.Context, network, address string) (net.Conn, error)
	serve   bool
}

// DialContext specifies the context to use when starting the command.
// The background context is used by default.
func DialContext(ctx context.Context) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialNetDial specifies a function used to dial a network connection. A
// default net.Dialer DialContext method is used by default.
func DialNetDial(f func(ctx context.Context, network, address string) (net.Conn, error)) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// DialServe specifies whether Server should be run in a goroutine.
// The default is to run Serve().
func DialServe(serve bool) DialOption { _ = "STUB: not implemented"; return *new(DialOption) }

// DialLogf specifies function for logging output. The log.Printf function is used by default.
func DialLogf(logf func(string, ...any)) DialOption {
	_ = "STUB: not implemented"
	return *new(DialOption)
}

// Dial dials an Nvim instance given an address in the format used by
// $NVIM_LISTEN_ADDRESS.
//
//	:help rpc-connecting
//	:help $NVIM_LISTEN_ADDRESS
func Dial(address string, options ...DialOption) (*Nvim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RegisterHandler registers fn as a MessagePack RPC handler for the named
// method. The function signature for fn is one of
//
//	func([v *nvim.Nvim,] {args}) ({resultType}, error)
//	func([v *nvim.Nvim,] {args}) error
//	func([v *nvim.Nvim,] {args})
//
// where {args} is zero or more arguments and {resultType} is the type of a
// return value. Call the handler from Nvim using the rpcnotify and rpcrequest
// functions:
//
//	:help rpcrequest()
//	:help rpcnotify()
func (v *Nvim) RegisterHandler(method string, fn any) error { _ = "STUB: not implemented"; return nil }

// ChannelID returns Nvim's channel id for this client.
func (v *Nvim) ChannelID() int { _ = "STUB: not implemented"; return 0 }

// TODO: log error and exit process?

func (v *Nvim) call(sm string, result any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// NewBatch creates a new batch.
func (v *Nvim) NewBatch() *Batch { _ = "STUB: not implemented"; return nil }

// Batch collects API function calls and executes them atomically.
//
// The function calls in the batch are executed without processing requests
// from other clients, redrawing or allowing user interaction in between.
// Functions that could fire autocommands or do event processing still might do
// so. For instance invoking the :sleep command might call timer callbacks.
//
// Call the Execute() method to execute the commands in the batch. Result
// parameters in the API function calls are set in the call to Execute.  If an
// API function call fails, all results proceeding the call are set and a
// *BatchError is returned.
//
// A Batch does not support concurrent calls by the application.
type Batch struct {
	err     error
	ep      *rpc.Endpoint
	enc     *msgpack.Encoder
	sms     []string
	results []any
	buf     bytes.Buffer
}

// Execute executes the API function calls in the batch.
func (b *Batch) Execute() error { _ = "STUB: not implemented"; return nil }

// emptyArgs represents a empty interface slice which use to empty args.
var emptyArgs = []any{}

func (b *Batch) call(sm string, result any, args ...any) { _ = "STUB: not implemented"; return }

// batchArg represents a batch call arguments.
type batchArg struct {
	n int
	p []byte
}

// compile time check whether the batchArg implements msgpack.Marshaler interface.
var _ msgpack.Marshaler = (*batchArg)(nil)

// MarshalMsgPack implements msgpack.Marshaler.
func (a *batchArg) MarshalMsgPack(enc *msgpack.Encoder) error {
	_ = "STUB: not implemented"
	return nil
}

// BatchError represents an error from a API function call in a Batch.
type BatchError struct {
	// Err is the error.
	Err error

	// Index is a zero-based index of the function call which resulted in the
	// error.
	Index int
}

// Error implements the error interface.
func (e *BatchError) Error() string { _ = "STUB: not implemented"; return "" }

func fixError(sm string, err error) error { _ = "STUB: not implemented"; return nil }

// ErrorList is a list of errors.
type ErrorList []error

// Error implements the error interface.
func (el ErrorList) Error() string { _ = "STUB: not implemented"; return "" }

// Request makes a any RPC request.
func (v *Nvim) Request(procedure string, result any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Request makes a any RPC request atomically as a part of batch request.
func (b *Batch) Request(procedure string, result any, args ...any) {
	_ = "STUB: not implemented"
	return
}

// Call calls a VimL function with the given arguments.
//
// Fails with VimL error, does not update "v:errmsg".
//
// fn is Function to call.
//
// args is Function arguments packed in an Array.
//
// result is the result of the function call.
func (v *Nvim) Call(fname string, result any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Call calls a VimL function with the given arguments.
//
// Fails with VimL error, does not update "v:errmsg".
//
// fn is Function to call.
//
// args is function arguments packed in an array.
//
// result is the result of the function call.
func (b *Batch) Call(fname string, result any, args ...any) { _ = "STUB: not implemented"; return }

// CallDict calls a VimL dictionary function with the given arguments.
//
// Fails with VimL error, does not update "v:errmsg".
//
// dict is dictionary, or string evaluating to a VimL "self" dict.
//
// fn is name of the function defined on the VimL dict.
//
// args is function arguments packed in an array.
//
// result is the result of the function call.
func (v *Nvim) CallDict(dict []any, fname string, result any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// CallDict calls a VimL dictionary function with the given arguments.
//
// Fails with VimL error, does not update "v:errmsg".
//
// dict is dictionary, or string evaluating to a VimL "self" dict.
//
// fn is name of the function defined on the VimL dict.
//
// args is Function arguments packed in an Array.
//
// result is the result of the function call.
func (b *Batch) CallDict(dict []any, fname string, result any, args ...any) {
	_ = "STUB: not implemented"
	return
}

// ExecLua execute Lua code.
//
// Parameters are available as `...` inside the chunk. The chunk can return a value.
//
// Only statements are executed. To evaluate an expression, prefix it
// with `return` is  "return my_function(...)".
//
// code is Lua code to execute.
//
// args is arguments to the code.
//
// The returned result value of Lua code if present or nil.
func (v *Nvim) ExecLua(code string, result any, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// ExecLua execute Lua code.
//
// Parameters are available as `...` inside the chunk. The chunk can return a value.
//
// Only statements are executed. To evaluate an expression, prefix it
// with `return` is  "return my_function(...)".
//
// code is Lua code to execute.
//
// args is arguments to the code.
//
// The returned result value of Lua code if present or nil.
func (b *Batch) ExecLua(code string, result any, args ...any) { _ = "STUB: not implemented"; return }

// Notify the user with a message.
//
// Relays the call to vim.notify. By default forwards your message in the
// echo area but can be overriden to trigger desktop notifications.
//
// msg is message to display to the user.
//
// logLevel is the LogLevel.
//
// opts is reserved for future use.
func (v *Nvim) Notify(msg string, logLevel LogLevel, opts map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// Notify the user with a message.
//
// Relays the call to vim.notify. By default forwards your message in the
// echo area but can be overriden to trigger desktop notifications.
//
// msg is message to display to the user.
//
// logLevel is the LogLevel.
//
// opts is reserved for future use.
func (b *Batch) Notify(msg string, logLevel LogLevel, opts map[string]any) {
	_ = "STUB: not implemented"
	return
}

// decodeExt decodes a MsgPack encoded number to go int value.
func decodeExt(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// encodeExt encodes n to MsgPack format.
func encodeExt(n int) []byte { _ = "STUB: not implemented"; return nil }

func unmarshalExt(dec *msgpack.Decoder, id int, v any) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
