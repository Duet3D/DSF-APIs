package inputs

import "github.com/Duet3D/DSF-APIs/godsfapi/v3/types"

// Compatibility level for emulation
type Compatibility string

const (
	// Default means no emulation (same as RepRapFirmware)
	Default Compatibility = "Default"
	// RepRapFirmware emulation (i.e. no emulation)
	RepRapFirmware = "RepRapFirmware"
	// Marlin emulation
	Marlin = "Marlin"
	// Teacup emulation
	Teacup = "Teacup"
	// Sprinter emulation
	Sprinter = "Sprinter"
	// Repetier emulation
	Repetier = "Repetier"
	// NanoDLP emulation (special)
	NanoDLP = "NanoDLP"
)

// DistanceUnit used for positioning
type DistanceUnit string

const (
	// MM represents millimeters
	MM DistanceUnit = "MM"
	// Inch represents inches
	Inch = "Inch"
)

// InputChannelState is the state of a channel
type InputChannelState string

const (
	// AwaitingAcknowledgement waits for message acknowledgement
	AwaitingAcknowledgement InputChannelState = "awaitingAcknowledgement"
	// Idle for an idle channel
	Idle = "idle"
	// Executing if channel executes G/M/T-code
	Executing = "executing"
	// Waiting for more data
	Waiting = "waiting"
	// Reading a G/M/T-code
	Reading = "reading"
)

const (
	// DefaultFeedRate on a channel
	DefaultFeedRate = 50.0
)

// InputChannel holds information about G/M/T-code channels
type InputChannel struct {
	// AxesRelative represents usage of relative positioning
	AxesRelative bool `json:"axesRelative"`
	// Compatibility is the emulation used on this channel
	Compatibility Compatibility `json:"compatibility"`
	// DistanceUnit is the distance unit in use
	DistanceUnit DistanceUnit `json:"distanceUnit"`
	// DrivesRelative represents usage of relative extrusion
	DrivesRelative bool `json:"drivesRelative"`
	// FeedRate is the current feedrate in mm/s
	FeedRate float64 `json:"feedRate"`
	// InMacro if a macro is being processed
	InMacro bool `json:"inMacro"`
	// Name of this channel
	Name types.CodeChannel `json:"name"`
	// StackDepth is the depth of the stack
	StackDepth uint8 `json:"stackDepth"`
	// State of this channel
	State InputChannelState `json:"state"`
	// LineNumber is the number of the current line
	LineNumber int64 `json:"lineNumber"`
	// Volumetric represents usage of volumetric extrusion
	Volumetric bool `json:"volumetric"`
}

// NewInputChannel returns an InputChannel with default values set
func NewInputChannel(name types.CodeChannel) InputChannel {
	return InputChannel{
		Compatibility:  RepRapFirmware,
		DistanceUnit:   MM,
		DrivesRelative: true,
		FeedRate:       DefaultFeedRate,
		Name:           name,
		State:          Idle,
	}
}
