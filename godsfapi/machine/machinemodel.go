package machine

import "github.com/Duet3D/DSF-APIs/godsfapi/commands"

// MachineModel represents the full machine model as maintained by DuetControlServer
type MachineModel struct {
	// Channels holds information about every available G/M/T-code channel
	Channels Channels `json:"channels"`
	// Directories holds information about the individual directories
	Directories Directories `json:"directories"`
	// Electronics holds information about the main and expansion boards
	Electronics Electronics `json:"electronics"`
	// Fans is a list of configured fans
	Fans []Fan `json:"fans"`
	// Heat holds information about the heat subsystem
	Heat Heat `json:"heat"`
	// HttpEndpoints is a list of registered third-party HTTP endpoints
	HttpEndpoints []HttpEndpoint `json:"httpEndpoints"`
	// Job holds information about the current file job (if any)
	Job Job `json:"job"`
	// Lasers is a list of configured laser diodes
	Lasers []Laser `json:"lasers"`
	// MessageBox holds information about message box requests
	MessageBox MessageBox `json:"messageBox"`
	// Messages is a list of generic messages that do not belong explicitly to codes
	// being executed. This includes status message, generic errors and outputs generated
	// by M118
	Messages []commands.Message `json:"messages"`
	// Move holds information about the move subsystem
	Move Move `json:"move"`
	// Network holds information about connected network adapters
	Network Network `json:"network"`
	// Scanner holds information about the 3D scanner subsystem
	Scanner Scanner `json:"scanner"`
	// Sensors holds information about connected sensors including Z-probes and endstops
	Sensors Sensors `json:"sensors"`
	// Spindles is a list of configured CNC spindles
	Spindles []Spindle `json:"spindles"`
	// State holds information about the machine state
	State State `json:"state"`
	// Storages is a list of configured storage devices
	Storages []Storage `json:"storages"`
	// Tools is a list of configure tools
	Tools []Tool `json:"tools"`
	// UserSessions is a list of user session
	UserSessions []UserSession `json:"userSessions"`
	// UserVariables is a list of user-defined variables
	UserVariables []UserVariable `json:"userVariables"`
}

// NewMachineModel creates a new MachineModel
func NewMachineModel() *MachineModel {
	return &MachineModel{}
}
