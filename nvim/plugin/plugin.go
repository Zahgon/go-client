package plugin

import (
	"reflect"

	"github.com/neovim/go-client/nvim"
)

// Plugin represents a remote plugin.
type Plugin struct {
	Nvim        *nvim.Nvim
	pluginSpecs []*pluginSpec

	// Event/pattern counters used to generate unique paths for autocmds.
	eventPathCounts map[string]int
}

// New returns an intialized plugin.
func New(v *nvim.Nvim) *Plugin { _ = "STUB: not implemented"; return nil }

// Disable support for "specs" method until path mechanism for supporting
// binary executables with Nvim is worked out.
// err := v.RegisterHandler("specs", func(path string) ([]*pluginSpec, error) {
//  return p.pluginSpecs, nil
// })

type pluginSpec struct {
	sm   string
	Type string            `msgpack:"type"`
	Name string            `msgpack:"name"`
	Sync bool              `msgpack:"sync"`
	Opts map[string]string `msgpack:"opts"`
}

func (spec *pluginSpec) path() string { _ = "STUB: not implemented"; return "" }

func isSync(f any) bool { _ = "STUB: not implemented"; return false }

func (p *Plugin) handle(fn any, spec *pluginSpec) { _ = "STUB: not implemented"; return }

// Handle registers fn as a MessagePack RPC handler for the specified method
// name. The function signature for fn is one of
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
func (p *Plugin) Handle(method string, fn any) { _ = "STUB: not implemented"; return }

// FunctionOptions specifies function options.
type FunctionOptions struct {
	// Name is the name of the function in Nvim. The name must be made of
	// alphanumeric characters and '_', and must start with a capital letter.
	Name string

	// Eval is an expression evaluated in Nvim. The result is passed to the
	// handler function.
	Eval string
}

// HandleFunction registers fn as a handler for a Nvim function. The function
// signature for fn is one of
//
//	func([v *nvim.Nvim,] args {arrayType} [, eval {evalType}]) ({resultType}, error)
//	func([v *nvim.Nvim,] args {arrayType} [, eval {evalType}]) error
//
// where {arrayType} is a type that can be unmarshaled from a MessagePack
// array, {evalType} is a type compatible with the Eval option expression and
// {resultType} is the type of function result.
//
// If options.Eval == "*", then HandleFunction constructs the expression to
// evaluate in Nvim from the type of fn's last argument. The last argument is
// assumed to be a pointer to a struct type with 'eval' field tags set to the
// expression to evaluate for each field. Nested structs are supported. The
// expression for the function
//
//	func example(eval *struct{
//		GOPATH string `eval:"$GOPATH"`
//		Cwd    string `eval:"getcwd()"`
//	})
//
// is
//
//	{'GOPATH': $GOPATH, Cwd: getcwd()}
func (p *Plugin) HandleFunction(options *FunctionOptions, fn any) {
	_ = "STUB: not implemented"
	return
}

// CommandOptions specifies command options.
type CommandOptions struct {
	// Name is the name of the command in Nvim. The name must be made of
	// alphanumeric characters and '_', and must start with a capital
	// letter.
	Name string

	// NArgs specifies the number command arguments.
	//
	//  0   No arguments are allowed
	//  1   Exactly one argument is required, it includes spaces
	//  *   Any number of arguments are allowed (0, 1, or many),
	//      separated by white space
	//  ?   0 or 1 arguments are allowed
	//  +   Arguments must be supplied, but any number are allowed
	NArgs string

	// Range specifies that the command accepts a range.
	//
	//  .   Range allowed, default is current line. The value
	//      "." is converted to "" for Nvim.
	//  %   Range allowed, default is whole file (1,$)
	//  N   A count (default N) which is specified in the line
	//      number position (like |:split|); allows for zero line
	//	    number.
	//
	//  :help :command-range
	Range string

	// Count specifies that the command accepts a count.
	//
	//  N   A count (default N) which is specified either in the line
	//	    number position, or as an initial argument (like |:Next|).
	//      Specifying -count (without a default) acts like -count=0
	//
	//  :help :command-count
	Count string

	// Addr sepcifies the domain for the range option
	//
	//  lines           Range of lines (this is the default)
	//  arguments       Range for arguments
	//  buffers         Range for buffers (also not loaded buffers)
	//  loaded_buffers  Range for loaded buffers
	//  windows         Range for windows
	//  tabs            Range for tab pages
	//
	//  :help command-addr
	Addr string

	// Eval is evaluated in Nvim and the result is passed as an argument.
	Eval string

	// Complete specifies command completion.
	//
	//  :help :command-complete
	Complete string

	// Bang specifies that the command can take a ! modifier (like :q or :w).
	Bang bool

	// Register specifies that the first argument to the command can be an
	// optional register name (like :del, :put, :yank).
	Register bool

	// Bar specifies that the command can be followed by a "|" and another
	// command.  A "|" inside the command argument is not allowed then. Also
	// checks for a " to start a comment.
	Bar bool
}

// HandleCommand registers fn as a handler for a Nvim command. The arguments
// to the function fn are:
//
//	v *nvim.Nvim        optional
//	args []string       when options.NArgs != ""
//	range [2]int        when options.Range == "." or Range == "%"
//	range int           when options.Range == N or Count != ""
//	bang bool           when options.Bang == true
//	register string     when options.Register == true
//	eval any            when options.Eval != ""
//
// The function fn must return an error.
//
// If options.Eval == "*", then HandleCommand constructs the expression to
// evaluate in Nvim from the type of fn's last argument. See the
// HandleFunction documentation for information on how the expression is
// generated.
func (p *Plugin) HandleCommand(options *CommandOptions, fn any) { _ = "STUB: not implemented"; return }

// AutocmdOptions specifies autocmd options.
type AutocmdOptions struct {
	// Event is the event name.
	Event string

	// Group specifies the autocmd group.
	Group string

	// Pattern specifies an autocmd pattern.
	//
	//  :help autocmd-patterns
	Pattern string

	// Nested allows nested autocmds.
	//
	//  :help autocmd-nested
	Nested bool

	// Once supplys the command is executed once, then removed ("one shot").
	//
	//  :help autocmd-once
	Once bool

	// Eval is evaluated in Nvim and the result is passed to the handler
	// function.
	Eval string
}

// HandleAutocmd registers fn as a handler an autocmnd event.
//
// If options.Eval == "*", then HandleAutocmd constructs the expression to
// evaluate in Nvim from the type of fn's last argument. See the HandleFunction
// documentation for information on how the expression is generated.
func (p *Plugin) HandleAutocmd(options *AutocmdOptions, fn any) { _ = "STUB: not implemented"; return }

// Compute unique path for event and pattern.

// RegisterForTests registers the plugin with Nvim. Use this method for testing
// plugins in an embedded instance of Nvim.
func (p *Plugin) RegisterForTests() error { _ = "STUB: not implemented"; return nil }

func eval(eval string, f any) string { _ = "STUB: not implemented"; return "" }

func structEval(t reflect.Type) string { _ = "STUB: not implemented"; return "" }

type byServiceMethod []*pluginSpec

func (a byServiceMethod) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a byServiceMethod) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (a byServiceMethod) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (p *Plugin) Manifest(host string) []byte { _ = "STUB: not implemented"; return nil }

// Sort for consistent order on output.
