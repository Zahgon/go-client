package nvim

import (
	"io"
)

type bufferReader struct {
	err   error
	v     *Nvim
	lines [][]byte
	b     Buffer
}

// compile time check whether the bufferReader implements io.Reader interface.
var _ io.Reader = (*bufferReader)(nil)

// NewBufferReader returns a reader for the specified buffer. If b = 0, then
// the current buffer is used.
func NewBufferReader(v *Nvim, b Buffer) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

// Read implements io.Reader.
func (r *bufferReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
