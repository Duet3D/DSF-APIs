// Deprecated: This package was deprected, please visit https://github.com/Duet3D/dsf-go.
package commands

import (
	"fmt"
	"strings"

	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/messages"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/types"
)

// CodeResult is a list of code results
// Deprecated: This class is now deprecated. It will be replaced with messages.Message in foreseeable future
type CodeResult []messages.Message

func (cr CodeResult) String() string {
	var b strings.Builder
	for _, m := range cr {
		if m.Content != "" {
			b.WriteString(m.String())
			b.WriteByte('\n')
		}
	}
	return b.String()
}

// CodeFlags are bit masks to classify G/M/T-codes
type CodeFlags int64

const (
	// Asynchronous codes are considered finished as soon as they enter the code queue
	Asynchronous CodeFlags = 1 << iota
	// IsPreProcessed marks pre-processed codes
	IsPreProcessed
	// IsPostProcessed marks post-processed codes
	IsPostProcessed
	// IsFromMacro indicates code originating from macro
	IsFromMacro
	// IsNestedMacro indicates code originating from system macro
	IsNestedMacro
	// IsFromConfig indicates code originating from config.g or config.g.bak
	IsFromConfig
	// IsFromConfigOverride indicated code originating from config-override.g
	IsFromConfigOverride
	// EnforceAbsolutePosition marks code prefixed with G53
	EnforceAbsolutePosition
	// IsPrioritized will be sent to the firmware as soon as possible jumping all queued codes
	IsPrioritized
	// Unbuffered will execute this code circumventing any buffers
	// Do NOT process another code on the same channel before this code has been fully executed
	Unbuffered
	// IsFromFirmware indicates if this code was requested by the firmware
	IsFromFirmware
	// IsLastCode indicates if this is the last code on the line
	IsLastCode
	// CodeFlagsNone is a placeholder to indicate that no flags are set
	CodeFlagsNone = 0
)

// CodeType is the generic type of G/M/T-code or being a comment
type CodeType string

const (
	// Comment if this code is whole line comment
	Comment CodeType = "Q"
	// GCode if this code is a G-Code
	GCode = "G"
	// MCode if this code is a M-Code
	MCode = "M"
	// TCode if this code is a T-Code
	TCode = "T"
)

// KeywordType is the type of conditional G-code
type KeywordType byte

const (
	// None for no conditional code
	None KeywordType = iota
	// If condition
	If
	// ElseIf condition
	ElseIf
	// Else condition
	Else
	// While condition
	While
	// Break instruction (used in While)
	Break
	// Return instruction
	// Deprecated: was never supported by RRF
	Return
	// Abort instruction
	Abort
	// Var operation
	Var
	// Set operation
	Set
	// Echo operation
	Echo
	// Continue instruction (used in While)
	Continue
	// Global operation
	Global
)

func (k KeywordType) String() string {
	switch k {
	case Abort:
		return "abort"
	case Break:
		return "break"
	case Echo:
		return "echo"
	case Else:
		return "else"
	case ElseIf:
		return "elif"
	case If:
		return "if"
	case Return:
		return "return"
	case Set:
		return "set"
	case Var:
		return "var"
	case While:
		return "while"
	default:
		return ""
	}
}

// Code is a parsed representation of a generic G/M/T/code
type Code struct {
	BaseCommand
	// SourceConnection ID this code was received from. If this is 0, the code originates from an internal DCS task
	// Usually there is no need to populate this property. It is internally overwritten by the control server on receipt
	SourceConnection int64
	// Result of this code. This property is only set when the code has finished its execution
	// It remains nil if the code has been cancelled.
	Result CodeResult
	// Type of the code
	Type CodeType
	// Channel to send this code to
	Channel types.CodeChannel
	// LineNumber of this code
	LineNumber *int64
	// Indent are the number of whitespaces prefixing the command content
	Indent byte
	// Keyword type of conditional G-code
	Keyword KeywordType
	// KeywordArgument of the conditional G-code
	KeywordArgument string
	// MajorNumber of the code (e.g. 28 in G28)
	MajorNumber *int64
	// MinorNumber of the code (e.g. 3 in G54.3)
	MinorNumber *int8
	// Flags of this code
	Flags CodeFlags
	// Comment provided to this G/M/T-code
	Comment string
	// FilePosition of this code in bytes (optional)
	FilePosition *int64
	// Length of the original code in bytes (optional)
	Length *int64
	// Parameters are a list of parsed code parameters
	Parameters []CodeParameter
}

// NewCode instantiates a Code with default values
func NewCode() *Code {
	return &Code{
		BaseCommand: *NewBaseCommand("Code"),
		Type:        Comment,
		Channel:     types.DefaultChannel,
		Keyword:     None,
		Flags:       CodeFlagsNone,
	}
}

// Clone an existing Code into a new instance
func (c *Code) Clone() *Code {
	cc := *c
	cparams := make([]CodeParameter, 0)
	for _, p := range c.Parameters {
		cp := p.Clone()
		cparams = append(cparams, *cp)
	}
	cc.Parameters = cparams
	return &cc
}

// IsMajorNumber is a convenience function that checks if
// the MajorNumber of this Code instance is present and equal
// to the given value
func (c *Code) IsMajorNumber(n int64) bool {
	return c.MajorNumber != nil && *c.MajorNumber == n
}

// HasFlag checks if this code has the given flag set
func (c *Code) HasFlag(flag CodeFlags) bool {
	return (c.Flags & flag) != 0
}

// HasParameter returns whether or not a certain parameter is present without returning the
// CodeParameter instance
func (c *Code) HasParameter(letter string) bool {
	return c.Parameter(letter) != nil
}

// Parameter retrieves a parameter for the given letter. This will return nil in case there
// is no parameter with this letter. Lookup is case-insensitive.
func (c *Code) Parameter(letter string) *CodeParameter {
	l := strings.ToUpper(letter)
	for _, p := range c.Parameters {
		if l == strings.ToUpper(p.Letter) {
			return &p
		}
	}
	return nil
}

// ParameterOrDefault will return the Parameter for the given letter or return the given default value.
// Lookup is case-insensitive.
func (c *Code) ParameterOrDefault(letter string, value interface{}) *CodeParameter {
	p := c.Parameter(letter)
	if p != nil {
		return p
	}
	return NewSimpleCodeParameter(letter, value)
}

// ReplaceParameter will replace the first occurrence of a parameter with the given letter
func (c *Code) ReplaceParameter(letter string, np *CodeParameter) bool {
	for i, p := range c.Parameters {
		if p.Letter == letter {
			c.Parameters[i] = *np
			return true
		}
	}
	return false
}

// RemoveParameter removes all parameters with the given letter
func (c *Code) RemoveParameter(letter string) *CodeParameter {
	p := c.Parameter(letter)
	if p != nil {
		cp := make([]CodeParameter, 0)
		for _, p := range c.Parameters {
			if p.Letter != letter {
				cp = append(cp, p)
			}
		}
		c.Parameters = cp
	}
	return p
}

// GetUnprecedentedString reconstructs an unprecedented string from parameter list
func (c *Code) GetUnprecedentedString(quote bool) string {
	var b strings.Builder
	for _, p := range c.Parameters {
		if b.Len() > 0 {
			b.WriteString(" ")
		}
		b.WriteString(p.Letter)
		if quote && p.IsString {
			b.WriteString(`"`)
		}
		b.WriteString(p.AsString())
		if quote && p.IsString {
			b.WriteString(`"`)
		}
	}
	return b.String()
}

// String will convert the parsed code back to a text-based G/M/T-code
func (c *Code) String() string {
	if c.Keyword != None {
		if c.KeywordArgument != "" {
			return fmt.Sprintf("%s %s", c.Keyword.String(), c.KeywordArgument)
		}
		return c.Keyword.String()
	}
	if c.Type == Comment {
		return ";" + c.Comment
	}
	var b strings.Builder
	b.WriteString(c.ShortString())

	for _, p := range c.Parameters {
		b.WriteString(" ")
		b.WriteString(p.String())
	}

	if c.Comment != "" {
		if b.Len() > 0 {
			b.WriteString(" ")
		}
		b.WriteString(";")
		b.WriteString(c.Comment)
	}

	if len(c.Result) > 0 {
		b.WriteString(" => ")
		b.WriteString(strings.TrimRight(c.Result.String(), " "))
	}

	return b.String()
}

// ShortString converts only the command portion to text-based G/M/T-code (e.g. G28)
func (c *Code) ShortString() string {
	if c.Keyword != None {
		return c.Keyword.String()
	}
	if c.Type == Comment {
		return "(comment)"
	}

	var prefix string
	if c.HasFlag(EnforceAbsolutePosition) {
		prefix = "G53 "
	}
	if c.MajorNumber != nil {
		if c.MinorNumber != nil {
			return fmt.Sprintf("%s%s%d.%d", prefix, c.Type, *c.MajorNumber, *c.MinorNumber)
		}
		return fmt.Sprintf("%s%s%d", prefix, c.Type, *c.MajorNumber)
	}
	return fmt.Sprintf("%s%s", prefix, string(c.Type))
}
