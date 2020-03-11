package machine

import "github.com/Duet3D/DSF-APIs/godsfapi/commands"

// MachineModel represents the full machine model as maintained by DuetControlServer
type MachineModel struct {
	// Channels holds information about every available G/M/T-code channel
	Channels Channels
	// Directories holds information about the individual directories
	Directories Directories
	// Electronics holds information about the main and expansion boards
	Electronics Electronics
	// Fans is a list of configured fans
	Fans []Fan
	// Heat holds information about the heat subsystem
	Heat Heat
	// HttpEndpoints is a list of registered third-party HTTP endpoints
	HttpEndpoints []HttpEndpoint
	// Job holds information about the current file job (if any)
	Job Job
	// Lasers is a list of configured laser diodes
	Lasers []Laser
	// MessageBox holds information about message box requests
	MessageBox MessageBox
	// Messages is a list of generic messages that do not belong explicitly to codes
	// being executed. This includes status message, generic errors and outputs generated
	// by M118
	Messages []commands.Message
	// Move holds information about the move subsystem
	Move Move
	// Network holds information about connected network adapters
	Network Network
	// Scanner holds information about the 3D scanner subsystem
	Scanner Scanner
	// Sensors holds information about connected sensors including Z-probes and endstops
	Sensors Sensors
	// Spindles is a list of configured CNC spindles
	Spindles []Spindle
	// State holds information about the machine state
	State State
	// Storages is a list of configured storage devices
	Storages []Storage
	// Tools is a list of configure tools
	Tools []Tool
	// UserSessions is a list of user session
	UserSessions []UserSession
	// UserVariables is a list of user-defined variables
	UserVariables []UserVariable
}

// NewMachineModel creates a new MachineModel
func NewMachineModel() *MachineModel {
	return &MachineModel{}
}
