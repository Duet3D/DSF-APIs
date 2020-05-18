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

var acknowledge = NewBaseCommand("Acknowledge")
var cancel = NewBaseCommand("Cancel")
var getMachineModel = NewBaseCommand("GetMachineModel")
var ignore = NewBaseCommand("Ignore")
var syncMachineModel = NewBaseCommand("SyncMachineModel")
var lockMachineModel = NewBaseCommand("LockMachineModel")
var unlockMachineModel = NewBaseCommand("UnlockMachineModel")

// NewAcknowledge returns a Acknowledge command
func NewAcknowledge() *BaseCommand {
	return acknowledge
}

// NewCancel returns a Cancel command
func NewCancel() *BaseCommand {
	return cancel
}

// NewGetMachineModel returns a GetMachineModel command
func NewGetMachineModel() *BaseCommand {
	return getMachineModel
}

// NewIgnore returns an Ignore command
func NewIgnore() *BaseCommand {
	return ignore
}

// NewSyncMachineModel returns a SyncMachineModel command
func NewSyncMachineModel() *BaseCommand {
	return syncMachineModel
}

// NewLockMachineModel returns a LockMachineModel command
func NewLockMachineModel() *BaseCommand {
	return lockMachineModel
}

// NewUnlockMachineModel returns a UnlockMachineModel command
func NewUnlockMachineModel() *BaseCommand {
	return unlockMachineModel
}
