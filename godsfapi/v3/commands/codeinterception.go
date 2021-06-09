// Deprecated: This package was deprected, please visit https://github.com/Duet3D/dsf-go.
package commands

import "github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/messages"

var cancel = NewBaseCommand("Cancel")
var ignore = NewBaseCommand("Ignore")

// NewCancel returns a Cancel command
func NewCancel() *BaseCommand {
	return cancel
}

// NewIgnore returns an Ignore command
func NewIgnore() *BaseCommand {
	return ignore
}

// Resolve the code to intercept and return the given message details for its completion.
type Resolve struct {
	BaseCommand
	// Type of the resolving message
	Type messages.MessageType
	// Content of the resolving message
	Content string
}

// NewResolve creates a new Resolve for the given type and message
func NewResolve(mType messages.MessageType, content string) *Resolve {
	return &Resolve{
		BaseCommand: *NewBaseCommand("Resolve"),
		Type:        mType,
		Content:     content,
	}
}
