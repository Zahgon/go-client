package nvimtest

import (
	"testing"

	"github.com/neovim/go-client/nvim"
)

// NewChildProcess returns the new *Nvim, and registers cleanup to tb.Cleanup.
func NewChildProcess(tb testing.TB) *nvim.Nvim { _ = "STUB: not implemented"; return nil }

// Mimics a fresh install of Nvim. See :help --clean
// Use stdin/stdout as a msgpack-RPC channel, so applications can embed and control Nvim via the RPC API.
// Start without UI, and do not wait for nvim_ui_attach
// Clean packpath
