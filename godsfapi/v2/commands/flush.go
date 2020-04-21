package commands

import "github.com/Duet3D/DSF-APIs/godsfapi/v2/types"

// Flush waits for all pending (macro) codes on the given channel to finish.
// This effectively guarantees that all buffered codes are processed by RRF
// before this command finishes.
// If the flush request is successful, true is returned
type Flush struct {
	BaseCommand
	// Channel is the CodeChannel to flush
	// This value is ignored if this request is processed while a code is
	// being intercepted.
	Channel types.CodeChannel
}

// NewFlush creates a flush command for the given CodeChannel
func NewFlush(channel types.CodeChannel) *Flush {
	return &Flush{
		BaseCommand: *NewBaseCommand("Flush"),
		Channel:     channel,
	}
}
