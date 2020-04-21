package commands

import "github.com/Duet3D/DSF-APIs/godsfapi/v2/types"

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
