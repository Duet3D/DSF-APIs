package move

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
	// DAA holds information about the configured Dynamic Acceleration Adjustment
	DAA DAA `json:"daa"`
	// Extruders is a list of configured extrudersr
	Extruders []Extruder `json:"extruders"`
	// Idle current reduction parameters
	Idle MotorsIdleControl `json:"idle"`
	// Kinematics holds information about the currently configured kinematics
	// Use one of the NewXXKinematics methods to convert it
	Kinematics Kinematics `json:"kinematics"`
	// PrintingAcceleration is maximum accelertion allowed while printing (in mm/s^2)
	PrintingAcceleration float64 `json:"printingAcceleration"`
	// SpeedFactor applied to every move (0.01..1 or greater)
	SpeedFactor float64 `json:"speedFactor"`
	// TravelAcceleration is maximum acceleration allowed while travelling (in mm/s^2)
	TravelAcceleration float64 `json:"travelAcceleration"`
	// VirtualPos is the virtual total extruder position
	VirtualPos float64 `json:"virtualPos"`
	// WorkspaceNumber is the index of the currently selected workspace
	WorkspaceNumber int `json:"workspaceNumber"`
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

// Default values for DAA
const (
	DefaultMinimumAcceleration = 10.0
)

// DAA holds information about Dynamic Acceleration Adjustment
type DAA struct {
	// Enabled indicates if DAA is enabled
	Enabled bool `json:"enabled"`
	// MinimumAcceleration allowed (in mm/s^2)
	MinimumAcceleration float64 `json:"minimumAcceleration"`
	// Period of the ringing that is supposed to be cancelled (in s)
	Period float64 `json:"period"`
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
	Driver string `json:"driver"`
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
	// XMin is the X start coordinate of the heightmap
	XMin float64 `json:"xMin"`
	// XMax is the X end coordinate of the heightmap
	XMax float64 `json:"xMax"`
	// XSpacing is the spacing between probe points in X direction
	XSpacing float64 `json:"xSpacing"`
	// YMin is the Y start coordinate of the heightmap
	YMin float64 `json:"yMin"`
	// YMax is the Y end coordinate of the heightmap
	YMax float64 `json:"yMax"`
	// YSpacing is the spacing between probe points in Y direction
	YSpacing float64 `json:"ySpacing"`
	// Radius is the probing radius on delta kinematics
	Radius float64 `json:"radius"`
}

// Skew holds details about orthogonal axis compensation parameters
type Skew struct {
	// TanXY is the tangent of the skew angle for XY axes
	TanXY float64
	// TaxXZ is the tangent of the skew angle for XZ axes
	TanXZ float64
	// TaxYZ is the tangent of the skew angle for YZ axes
	TanYZ float64
}
