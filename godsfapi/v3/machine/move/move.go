package move

import "github.com/Duet3D/DSF-APIs/godsfapi/v3/types"

// Move holds information about the move subsystem
type Move struct {
	// Axes is a list of configured axes
	Axes []Axis `json:"axes"`
	// Calibration holds information about automatic calibration
	Calibration MoveCalibration `json:"calibration"`
	// Compensation holds information about the currently configured compensation options
	Compensation MoveCompensation `json:"compensation"`
	// CurrentMove holds information about the current move
	CurrentMove CurrentMove `json:"currentMove"`
	// Extruders is a list of configured extrudersr
	Extruders []Extruder `json:"extruders"`
	// Idle current reduction parameters
	Idle MotorsIdleControl `json:"idle"`
	// Kinematics holds information about the currently configured kinematics
	// Use one of the NewXXKinematics methods to convert it
	Kinematics Kinematics `json:"kinematics"`
	// PrintingAcceleration is maximum accelertion allowed while printing (in mm/s^2)
	PrintingAcceleration float64 `json:"printingAcceleration"`
	// Queue is a list of move queue items (DDA rings)
	Queue []MoveQueueItem `json:"queue"`
	// Shaping are the input shaping parameter
	Shaping MoveInputShaping `json:"shaping"`
	// SpeedFactor applied to every move (0.01..1 or greater)
	SpeedFactor float64 `json:"speedFactor"`
	// TravelAcceleration is maximum acceleration allowed while travelling (in mm/s^2)
	TravelAcceleration float64 `json:"travelAcceleration"`
	// VirtualPos is the virtual total extruder position
	VirtualPos float64 `json:"virtualPos"`
	// WorkspaceNumber is the index of the currently selected workplace (0..8)
	WorkplaceNumber int `json:"workplaceNumber"`
}

// MoveCalibration holds information about configured calibration options
type MoveCalibration struct {
	// Final calibration results (for Delta calibration)
	Final MoveDeviations `json:"final"`
	// Initial calibration resuls (for Delta calibration)
	Initial MoveDeviations `json:"initial"`
	// NumFactors is the number of factors used (for Delta calibration)
	NumFactors int64 `json:"numFactors"`
}

// MoveCompensation holds informatin about the configured compensation options
type MoveCompensation struct {
	// FadeHeight effective height before the bed compensation is turned off (in mm or nil if not configurea)
	FadeHeight *float64 `json:"fadeHeight"`
	// File is the full path to the currently used height map file or empty if none is used
	File string `json:"file"`
	// MeshDeviation are the deviations of the mesh grid of nil if not applicable
	MeshDeviation *MoveDeviations `json:"meshDeviation"`
	// ProbeGrid holds the settings of the current probe grid
	ProbeGrid ProbeGrid `json:"probeGrid"`
	// Skew holds information about the configured orthogonal axis parameters
	Skew Skew
	// Type is the type of compensation in use
	Type MoveCompensationType `json:"type"`
}

// MoveCompensationType are the supported compensation types
type MoveCompensationType string

const (
	// None for no compensations
	None MoveCompensationType = "none"
	// Mesh for mesh compensation
	Mesh = "mesh"
)

// MoveDeviations holds calibration or mesh grid results
type MoveDeviations struct {
	// Deviation RMS (in mm)
	Deviation float64 `json:"deviation"`
	// Mean deviation (in mm)
	Mean float64 `json:"mean"`
}

const (
	DefaultDamping             = 0.2
	DefaultFrequency           = 40
	DefaultMinimumAcceleration = 10
)

// MoveInputShaping describes parameters of input shaping
type MoveInputShaping struct {
	// Damping factor
	Damping float64 `json:"damping"`
	// Frequency in Hz
	Frequency float64 `json:"frequency"`
	// MinimumAcceleration in mm/s
	MinimumAcceleration float64 `json:"minimumAcceleration"`
	// Type of configured input shaping
	Type MoveInputShapingType `json:"type"`
}

// MoveInputShapingType are the possible input shaping methods
type MoveInputShapingType string

const (
	MoveInputShapingTypeNone MoveInputShapingType = "none"
	ZVD                                           = "ZVD"
	ZVDD                                          = "ZVDD"
	EI2                                           = "EI2"
	DAA                                           = "DAA"
)

// MoveQueueItem is information about a DDA ring
type MoveQueueItem struct {
	// GracePeriod is the minimum idle time in milliseconds before we should start a move
	GracePeriod uint64 `json:"gracePeriod"`
	// Length is the maximum number of moves that can be accomodated in the DDA ring
	Length int64 `json:"length"`
}

// Default values for Axis
const (
	DefaultJerk           = 15.0
	DefaultMaxTravel      = 200.0
	DefaultStepsPerMmAxis = 80.0
)

// Axis holds information about a configured axis
type Axis struct {
	// Acceleration of this axis (in mm/s^2)
	Acceleration float64 `json:"acceleration"`
	// Babystep amount (in mm)
	Babystep float64 `json:"babystep"`
	// Current of the motor (in mA)
	Current int64 `json:"current"`
	// Drivers is a list of assigned drivers
	Drivers []types.DriverId `json:"drivers"`
	// Homed indicates homing status
	Homed bool `json:"homed"`
	// Jerk of the motor (in mm/s)
	Jerk float64 `json:"jerk"`
	// Letter assigned to this axis (always upper-case)
	Letter string `json:"letter"`
	// MachinePosition is the current machine position (in mm or nil if unknown)
	MachinePosition *float64 `json:"machinePosition"`
	// Max travel of this axis (in mm)
	Max float64 `json:"max"`
	// MaxProbed is ture if the maximum was probed
	MaxProbed bool `json:"maxProbed"`
	// Microstepping of this axis
	Microstepping Microstepping `json:"microstepping"`
	// Min travel of this axis (in mm)
	Min float64 `json:"min"`
	// MinProbed is true if the minimum was probed
	MinProbed bool `json:"minProbed"`
	// Speed is the maximum speed (in mm/s)
	Speed float64 `json:"speed"`
	// StepsPerMm for this axis
	StepsPerMm float64 `json:"stepsPerMm"`
	// UserPosition (in mm or nil if unknown)
	UserPosition *float64 `json:"userPosition"`
	// Visible is true if the axis is not explicitely hidden
	Visible bool `json:"visible"`
	// WorkplaceOffsets for this axis (in mm)
	WorkplaceOffsets []float64 `json:"workplaceOffsets"`
}

// CurrentMove holds information about the current move
type CurrentMove struct {
	// Acceleration of the current move (in mm/s^2)
	Acceleration float64 `json:"acceleration"`
	// Deceleration of the current move (in mm/s^2)
	Deceleration float64 `json:"deceleration"`
	// LaserPwm of the current move as 0..1 or nil if not applicable
	LaserPwm *float64 `json:"laserPwm"`
	// RequestedSpeed of the current move (in mm/s)
	RequestedSpeed float64 `json:"requestedSpeed"`
	// TopSpeed actually reached for the current move (in mm/s)
	TopSpeed float64 `json:"topSpeed"`
}

// Default values for Extruder
const (
	DefaultMaxExtruderSpeed   = 100.0
	DefaultStepsPerMmExtruder = 420.0
)

// Extruder holds information about an extruder drive
type Extruder struct {
	// Acceleration of this extruder (in mm/s^2)
	Acceleration float64 `json:"acceleration"`
	// Current of the motor (in mA)
	Current int64 `json:"current"`
	// Driver is the assigned driver
	Driver types.DriverId `json:"driver"`
	// Filament is the name fo the currently loaded filament
	Filament string `json:"filament"`
	// Factor is the extrusion factor (1.0 equals 100%)
	Factor float64 `json:"factor"`
	// Jerk of the motor (in mm/s)
	Jerk float64 `json:"jerk"`
	// Microstepping of this extruder
	Microstepping Microstepping `json:"microstepping"`
	// NonLinear extrusion parameters (see M592)
	NonLinear ExtruderNonLinear `json:"nonLinear"`
	// Position of the extruder (in mm)
	Position float64 `json:"position"`
	// PressureAdvance (in s)
	PressureAdvance float64 `json:"pressureAdvance"`
	// RawPosition is the extruder position without factor applied (in mm)
	RawPosition float64 `json:"rawPosition"`
	// Speed is the maximum speed (in mm/s)
	Speed float64 `json:"speed"`
	// StepsPerMm for this extruder
	StepsPerMm float64 `json:"stepsPerMm"`
}

// ExtruderNonLinear contains non-linear extrusion parameters (see M592)
type ExtruderNonLinear struct {
	// A coefficient in the extrusion formula
	A float64 `json:"a"`
	// B coefficient in the extrusion formula
	B float64 `json:"b"`
	// UpperLimit of the nonlinear extrusion compensation
	UpperLimit float64 `json:"upperLimit"`
}

// Microstepping holds information about configured microstepping
type Microstepping struct {
	// Interpolated indicates if microstep interpolation is in use
	Interpolated bool `json:"interpolated"`
	// Value is the microstepping factor
	Value uint16 `json:"value"`
}

// MotorsIdleControl holds idle factor parameters for automatic MotorsIdleControl
// current reduction
type MotorsIdleControl struct {
	// Timeout after which the motor currents are reduced (in s)
	Timeout float64 `json:"timeout"`
	// Factor of the reduction on a scale between 0 and 1
	Factor float64 `json:"factor"`
}

// ProbeGrid holds information about the configured probe grid (see M557)
type ProbeGrid struct {
	// Axes are the axis letters of this heightmap
	Axes []string `json:"axes"`
	// Maxs are the end coordinates of the heightmap
	Maxs []float64 `json:"maxs"`
	// Mins are the start coordinates of the heightmap
	Mins []float64 `json:"mins"`
	// Radius is the probing radius on delta kinematics
	Radius float64 `json:"radius"`
	// Spacings between coordinates
	Spacings []float64 `json:"spacings"`
}

// Skew holds details about orthogonal axis compensation parameters
type Skew struct {
	// CompensateXY indicates if TanXY value is applied to the X or Y axis value
	CompensateXY bool `json:"compensateXY"`
	// TanXY is the tangent of the skew angle for XY axes
	TanXY float64 `json:"tanXY"`
	// TaxXZ is the tangent of the skew angle for XZ axes
	TanXZ float64 `json:"tanXZ"`
	// TaxYZ is the tangent of the skew angle for YZ axes
	TanYZ float64 `json:"tanYZ"`
}
