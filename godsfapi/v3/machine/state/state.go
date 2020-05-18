package state

const (
	// NoTool is the tool index if no tool is selected
	NoTool = -1
)

// State holds information about the machine state
type State struct {
	// AtxPower is the state of the ATX power pin (nil if not configured)
	AtxPower *bool `json:"atxPower"`
	// Beed holds information about a requested beep
	Beep BeepRequest `json:"beep"`
	// CurrentTool is the number of the currently selected tool or -1 if none is selected
	CurrentTool int64 `json:"currentTool"`
	// DisplayMessage is a persistent message to display (see M117)
	DisplayMessage string `json:"displayMessage"`
	// DsfVersion is the version of Duet SoftwareFramework
	DsfVersion string `json:"dsfVersion"`
	// GpOut is a list of general-purpose output ports
	GpOut []GpOutputPort `json:"gpOut"`
	// LaserPwm is laser PWM of the next commanded move on a scale of 0..1 or nil if not applicable
	LaserPwm *float64 `json:"laserPwm"`
	// LogFile being written to (empty if logging is disabled)
	LogFile string `json:"logFile"`
	// MessageBox holds details about a requested message box or nil if none is requested
	MessageBox *MessageBox `json:"messageBox"`
	// MachineMode the machine is currently in
	MachineMode MachineMode `json:"machineMode"`
	// NextTool is the number of the next to to be selected
	NextTool int64 `json:"nextTool"`
	// PowerFailScript is the script to execute when power fails
	PowerFailScript string `json:"powerFailScript"`
	// PreviousTool is the number of the previous tool
	PreviousTool int64 `json:"previousTool"`
	// RestorePoints is a list of restore points
	RestorePoints []RestorePoint `json:"restorePoints"`
	// Status the machine has currently
	Status MachineStatus `json:"status"`
	// UpTime is how long the mchine has been running (in s)
	UpTime uint64 `json:"upTime"`
}

// BeepRequest about a requested beep
type BeepRequest struct {
	// Duration of the requested beep (in ms)
	Duration uint64 `json:"duration"`
	// Frequency of the requested beep (in Hz)
	Frequency uint64 `json:"frequency"`
}

// GpOutputPort holds details about a general-purpose output port
type GpOutputPort struct {
	// Pwm value of this port in range 0..1
	Pwm float64 `json:"pwm"`
}

// MachineMode represents supported operation modes of the machine
type MachineMode string

const (
	// FFF is Fused Filament Fabrication (default)
	FFF MachineMode = "FFF"
	// CNC is computer numerical control
	CNC = "CNC"
	// Laser for laser operation mode (e.g. laser cutters)
	Laser = "Laser"
)

// MachineStatus represents possibile states of the firmware
type MachineStatus string

const (
	// Updating while firmware is being updated
	Updating MachineStatus = "updating"
	// Off if the machine is turned off (i.e. the input voltage is too low for operation)
	Off = "off"
	// Halted if the machine has encountered an emergency stop and is ready to reset
	Halted = "halted"
	// Pausing if the machine is baout to pause a file job
	Pausing = "pausing"
	// Paused if the machine has paused a file job
	Paused = "paused"
	// Resuming if the machine is about to resume a paused file job
	Resuming = "resuming"
	// Processing if the machine is processing a file job
	Processing = "processing"
	// Simulating while the machine is simulation a file job to determine its processing time
	Simulating = "simulating"
	// Busy if the machine is busy doing something (e.g. moving)
	Busy = "busy"
	// ChangingTool if the machine is chaging tools
	ChangingTool = "changingTool"
	// Idle if the machine is on but idle
	Idle = "idle"
)

// RestorePoint holds information about a restore point
type RestorePoint struct {
	// Coords are the axis coordinates of the restore point (in mm)
	Coords []float64 `json:"coords"`
	// ExtruderPos is the virtual extuder position at the start of this move
	ExtruderPos float64 `json:"extruderPos"`
	// FeedRate is the requested feed rate (in mm/s)
	FeedRate float64 `json:"feedRate"`
	// IoBits is the output port bits setting for this move or nil if not applicable
	IoBits *int64 `json:"ioBits"`
	// LaserPwm value in the range 0..1 or nil if not applicable
	LaserPwm *float64 `json:"laserPwm"`
	// SpindleSpeeds are the spindle RPMs that were set, negative if anticlockwise direction
	SpindleSpeeds []float64 `json:"spindleSpeeds"`
	// ToolNumber of the tool that was active
	ToolNumber int64 `json:"toolNumber"`
}
