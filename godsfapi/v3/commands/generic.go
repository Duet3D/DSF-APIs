package commands

import (
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/messages"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/state"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/types"
)

// CheckPassword checks if the given password is correct and matches the previously set value from M551.
// If no password was configured before or if it was set to "reprap" this will always return true.
type CheckPassword struct {
	BaseCommand
	Password string
}

// NewCheckPassword creates a new CheckPassword instance for the given password
func NewCheckPassword(password string) *CheckPassword {
	return &CheckPassword{
		BaseCommand: *NewBaseCommand("CheckPassword"),
		Password:    password,
	}
}

// EvaluateExpression can be used to evaluate an arbitrary expression on the given channel in RepRapFirmware
//
// Do not use this call to evaluation file-based or network-related fields because DSF and
// RRF models diverge in this regard
type EvaluateExpression struct {
	BaseCommand
	// Channel where the expression is evaluated
	Channel types.CodeChannel
	// Expression to evaluate
	Expression string
}

// NewEvaluateExpression creates a new EvaluateExpression instance for the given settings
func NewEvaluateExpression(channel types.CodeChannel, expression string) *EvaluateExpression {
	return &EvaluateExpression{
		BaseCommand: *NewBaseCommand("EvaluateExpression"),
		Channel:     channel,
		Expression:  expression,
	}
}

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

// SetUpdateStatus overrides the current status as reported by the object model
// when performing a software update
type SetUpdateStatus struct {
	BaseCommand
	// Updating sets whether an update is in progress
	Updating bool
}

// NewSetUpdateStatus creates a new SetUpdateStatus command
func NewSetUpdateStatus(updating bool) *SetUpdateStatus {
	return &SetUpdateStatus{
		BaseCommand: *NewBaseCommand("SetUpdateStatus"),
		Updating:    updating,
	}
}

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

// WriteMessage writes an arbitrary generic message.
// If neither OutputMessage nor LogMessage is true the message is
// written to the console output.
type WriteMessage struct {
	BaseCommand
	// Type of the message to write
	Type messages.MessageType
	// Content of the message to write
	Content string
	// OutputMessage on the console and via the object model
	OutputMessage bool
	// LogMessage writes the message to the log file (if applicable)
	// Deprecated: in favor of LogLevel
	LogMessage bool
	// LogLevel of this message
	LogLevel *state.LogLevel
}

// NewWriteMessage creates a new WriteMessage
func NewWriteMessage(mType messages.MessageType, content string, outputMessage bool, logLevel *state.LogLevel) *WriteMessage {
	return &WriteMessage{
		BaseCommand:   *NewBaseCommand("WriteMessage"),
		Type:          mType,
		Content:       content,
		OutputMessage: outputMessage,
		LogLevel:      logLevel,
	}
}
