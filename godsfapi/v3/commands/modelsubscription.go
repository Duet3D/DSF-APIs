package commands

var acknowledge = NewBaseCommand("Acknowledge")

// NewAcknowledge returns a Acknowledge command
func NewAcknowledge() *BaseCommand {
	return acknowledge
}
