package machine

import (
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/boards"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/directories"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/fans"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/heat"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/httpendpoints"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/inputs"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/job"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/limits"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/messages"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/move"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/network"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/scanner"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/sensors"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/spindles"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/state"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/tool"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/usersessions"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/uservariables"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/machine/volume"
)

// MachineModel represents the full machine model as maintained by DuetControlServer
type MachineModel struct {
	// Boards holds the list of connected boards
	Boards []boards.Board `json:"boards"`
	// Directories holds information about the individual directories
	Directories directories.Directories `json:"directories"`
	// Fans is a list of configured fans
	Fans []fans.Fan `json:"fans"`
	// Heat holds information about the heat subsystem
	Heat heat.Heat `json:"heat"`
	// HttpEndpoints is a list of registered third-party HTTP endpoints
	HttpEndpoints []httpendpoints.HttpEndpoint `json:"httpEndpoints"`
	// Inputs holds information about every available G/M/T-code channel
	Inputs inputs.Inputs `json:"inputs"`
	// Job holds information about the current file job (if any)
	Job job.Job `json:"job"`
	// Limits are machine configuration limits
	Limits limits.Limits `json:"limits"`
	// Messages is a list of generic messages that do not belong explicitly to codes
	// being executed. This includes status message, generic errors and outputs generated
	// by M118
	Messages []messages.Message `json:"messages"`
	// Move holds information about the move subsystem
	Move move.Move `json:"move"`
	// Network holds information about connected network adapters
	Network network.Network `json:"network"`
	// Scanner holds information about the 3D scanner subsystem
	Scanner scanner.Scanner `json:"scanner"`
	// Sensors holds information about connected sensors including Z-probes and endstops
	Sensors sensors.Sensors `json:"sensors"`
	// Spindles is a list of configured CNC spindles
	Spindles []spindles.Spindle `json:"spindles"`
	// State holds information about the machine state
	State state.State `json:"state"`
	// Tools is a list of configure tools
	Tools []tool.Tool `json:"tools"`
	// UserSessions is a list of user session
	UserSessions []usersessions.UserSession `json:"userSessions"`
	// UserVariables is a list of user-defined variables
	UserVariables []uservariables.UserVariable `json:"userVariables"`
	// Volumes holds a list of available mass storages
	Volumes []volume.Volume `json:"volumes"`
}

// NewMachineModel creates a new MachineModel
func NewMachineModel() *MachineModel {
	return &MachineModel{}
}
