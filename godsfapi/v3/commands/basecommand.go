package commands

// Command interface
type Command interface {
	// GetCommand returns the type of command
	GetCommand() string
}

// BaseCommand is the common base member of nearly all actual commands
type BaseCommand struct {
	Command string
}

// GetCommand returns the type of command
func (bc *BaseCommand) GetCommand() string {
	return bc.Command
}

// NewBaseCommand instantiates a new BaseCommand with the given name
func NewBaseCommand(command string) *BaseCommand {
	return &BaseCommand{Command: command}
}
