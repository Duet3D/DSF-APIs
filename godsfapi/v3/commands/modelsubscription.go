// Deprecated: This package was deprected, please visit https://github.com/Duet3D/dsf-go.
package commands

var acknowledge = NewBaseCommand("Acknowledge")

// NewAcknowledge returns a Acknowledge command
func NewAcknowledge() *BaseCommand {
	return acknowledge
}
