package commands

import "github.com/Duet3D/DSF-APIs/godsfapi/types"

// SimpleCode performs a simple G/M/T-code.
// On the server the code passed is converted to a full Code instance and on completion
// its CodeResult is transformed back into a basic string. This is useful for minimal extensions
// that do not require granular control of the code details.
// Important Note: Except for certain cases, it is NOT recommended for usage in
// connection.InterceptionConnection because it renders the internal code buffer useless.
type SimpleCode struct {
	BaseCommand
	// Code to parse and execute
	Code string
	// Channel to execute this code on
	Channel types.CodeChannel
}

// NewSimpleCode creates a new SimpleCode for the given code and channel.
func NewSimpleCode(code string, channel types.CodeChannel) *SimpleCode {
	return &SimpleCode{
		BaseCommand: *NewBaseCommand("SimpleCode"),
		Code:        code,
		Channel:     channel,
	}
}
