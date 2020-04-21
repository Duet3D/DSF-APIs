package machine

// MachineMode represents supported operation modes of the machine
type MachineMode string

const (
	// FFF is Fused Filament Fabrication (default)
	FFF MachineMode = "FFF"
	// CNC is computer numerical control
	CNC = "CNC"
	// MachineModeLaser for laser operation mode (e.g. laser cutters)
	MachineModeLaser = "Laser"
)

// MachineStatus represents possibile states of the firmware
type MachineStatus string

const (
	// Updating while firmware is being updated
	Updating MachineStatus = "Updating"
	// MachineStatusOff if the machine is turned off (i.e. the input voltage is too low for operation)
	MachineStatusOff = "Off"
	// Halted if the machine has encountered an emergency stop and is ready to reset
	Halted = "Halted"
	// Pausing if the machine is baout to pause a file job
	Pausing = "Pausing"
	// Paused if the machine has paused a file job
	Paused = "Paused"
	// Resuming if the machine is about to resume a paused file job
	Resuming = "Resuming"
	// Processing if the machine is processing a file job
	Processing = "Processing"
	// Simulating while the machine is simulation a file job to determine its processing time
	Simulating = "Simulating"
	// Busy if the machine is busy doing something (e.g. moving)
	Busy = "Busy"
	// ChangingTool if the machine is chaging tools
	ChangingTool = "ChangingTool"
	// MachineStatusIdle if the machine is on but idle
	MachineStatusIdle = "Idle"
)

const (
	// NoTool is the tool index if no tool is selected
	NoTool = -1
)

// State holds information about the machine state
type State struct {
	// AtxPower is the state of the ATX power pin (nil if not configured)
	AtxPower *bool `json:"atxPower"`
	// Beed holds information about a requested beep
	Beep BeepDetails `json:"beep"`
	// CurrentTool is the number of the currently selected tool or -1 if none is selected
	CurrentTool int64 `json:"currentTool"`
	// DisplayMessage is a persistent message to display (see M117)
	DisplayMessage string `json:"displayMessage"`
	// LogFile being written to (empty if logging is disabled)
	LogFile string `json:"logFile"`
	// Mode the machine is currently in
	Mode MachineMode `json:"mode"`
	// Status the machine has currently
	Status MachineStatus `json:"status"`
}

// BeepDetails about a requested beep
type BeepDetails struct {
	// Frequency of the requested beep (in Hz)
	Frequency uint64 `json:"frequency"`
	// Duration of the requested beep (in ms)
	Duration float64 `json:"duration"`
}
