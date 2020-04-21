package machine

// KinematicsType represents the supported kinmatics types
type KinematicsType string

const (
	// Cartesian kinematics
	Cartesian KinematicsType = "cartesian"
	// CoreXY kinematics
	CoreXY = "corexy"
	// CoreXYU is a CoreXY kinematics with extra U axis
	CoreXYU = "corexyu"
	// CoreXYUV is a CoreXY kinematics with extra UV axes
	CoreXYUV = "corexyuv"
	// CoreXZ kinmatics
	CoreXZ = "corexz"
	// Hangprinter kinematics
	Hangprinter = "hangprinter"
	// Delta kinematics
	Delta = "delta"
	// Polar kinematics
	Polar = "polar"
	// RotaryDelta kinematics
	RotaryDelta = "rotary delta"
	// SCARA kinematics
	SCARA = "scara"
	// Unknown kinematics
	Unknown = "unknown"
)

// Move holds information about the move subsystem
type Move struct {
	// Axes is a list of configured axes
	Axes []Axis `json:"axes"`
	// BabystepZ is the current babystep amount in Z direction in mm
	BabystepZ float64 `json:"babystepZ"`
	// CurrentMove holds information about the current move
	CurrentMove CurrentMove `json:"currentMove"`
	// Compensation is the name of the currently used bed compensation
	// (one of "Mesh", "[n] Point", "None")
	Compensation string `json:"compensation"`
	// HeightmapFile is the path to the current heightmap file if Compensation is "Mesh"
	HeightmapFile string `json:"heightmapFile"`
	// Drives is a list of configured drives
	Drives []Drive `json:"drives"`
	// Extruders is a list of configured extrudersr
	Extruders []Extruder `json:"extruders"`
	// Kinematics holds information about the currently configured kinematics
	Kinematics Kinematics `json:"Geometry"`
	// Idle current reduction parameters
	Idle MotorsIdleControl `json:"idle"`
	// ProbeGrid holds information about the configured mesh compensation (see M557)
	ProbeGrid ProbeGrid `json:"probeGrid"`
	// SpeedFactor applied to every regular move (1.0 equals 100%)
	SpeedFactor float64 `json:"speedFactor"`
	// CurrentWorkplace is the index of the selected workspace
	CurrentWorkplace int64 `json:"currentWorkplace"`
	// WorkplaceCoordinates are the axis offsets of each available workspace in mm
	WorkplaceCoordinates [][]float64 `json:"workplaceCoordinates"`
}

// Axis holds information about a configured axis
type Axis struct {
	// Letter assigned to this axis (always upper-case)
	Letter string `json:"letter"`
	// Drives is a list of drive indices assigned to this axis
	Drives []int64 `json:"drives"`
	// Homed is true if the axis has been homed
	Homed bool `json:"homed"`
	// MachinePosition is the current machine position (in mm or nil if unknown)
	MachinePosition *float64 `json:"machinePosition"`
	// Min travel of this axis (in mm or nil if unknown)
	Min *float64 `json:"min"`
	// MinEndstop is the index of the endstop that is used for the low end
	// or nil if none is configured
	MinEndstop *int64 `json:"minEndstop"`
	// MinProbed is true if the minimum was probed
	MinProbed bool `json:"minProbed"`
	// Max travel of this axis (in mm or nil if unknown)
	Max *float64 `json:"max"`
	// MaxEndstop is the index of the endstop that is used for the high end
	// or nil if none is configured
	MaxEndstop *int64 `json:"maxEndstop"`
	// MaxProbed is ture if the maximum was probed
	MaxProbed bool `json:"maxProbed"`
	// Visible is true if the axis is not explicitely hidden
	Visible bool `json:"visible"`
}

// CurrentMove holds information about the current move
type CurrentMove struct {
	// RequestedSpeed of the current move (in mm/s)
	RequestedSpeed float64 `json:"requestedSpeed"`
	// TopSpeed actually reached for the current move (in mm/s)
	TopSpeed float64 `json:"topSpeed"`
}

// Drive holds information about a drive
type Drive struct {
	// Position is the current user position of this drive (in mm)
	Position float64 `json:"position"`
	// Microstepping configured for this drive
	Microstepping DriveMicrostepping `json:"microstepping"`
	// Current configured for this drive (in mA)
	Current uint64 `json:"current"`
	// Acceleration of this drive (in mm/s²)
	Acceleration float64 `json:"acceleration"`
	// MinSpeed allowed for this drive (in mm/s)
	MinSpeed float64 `json:"minSpeed"`
	// MaxSpeed allowed for this drive (in mm/s)
	MaxSpeed float64 `json:"maxSpeed"`
}

// DriveMicrostepping holds information about configured microstepping
type DriveMicrostepping struct {
	// Value of the microstepping (e.g. x16)
	Value uint64 `json:"value"`
	// Interpolated is true if interpolation is used
	Interpolated bool `json:"interpolated"`
}

// Extruder holds information about an extruder drive
type Extruder struct {
	// Drives is a list of drive indices of this extruder
	Drives []int64 `json:"drives"`
	// Factor is the extrusion factor (1.0 equals 100%)
	Factor float64 `json:"factor"`
	// NonLinear extrusion parameters (see M592)
	NonLinear ExtruderNonLinear `json:"nonLinear"`
}

// ExtruderNonLinear contains non-linear extrusion parameters (see M592)
type ExtruderNonLinear struct {
	// A coefficient in the extrusion formula
	A float64 `json:"a"`
	// B coefficient in the extrusion formula
	B float64 `json:"b"`
	// UpperLimit of the nonlinear extrusion compensation
	UpperLimit float64 `json:"upperLimit"`
	// Temperature at which these values are valid in degC (future use only)
	Temperature float64 `json:"temperature"`
}

// Kinematics holds information about the configured kinematics
// Note to developers: this is called "Geometry" upstream
type Kinematics struct {
	// Type of currently configured kinematics
	Type KinematicsType `json:"type"`
	// Anchors of a hangprinter A, B, C, Dz (10 values)
	Anchors []float64 `json:"anchors"`
	// PrintRadius for Hangprinter and Delta kinematics in mm
	PrintRadius float64 `json:"printRadius"`
	// Diagonals for a delta
	Diagonals []float64 `json:"diagonals"`
	// Radius of a delta in mm
	Radius float64 `json:"radius"`
	// HomedHeight of a delta in mm
	HomedHeight float64 `json:"homedHeight"`
	// AngleCorrections ABC for delta kinematics
	AngleCorrections []float64 `json:"angleCorrections"`
	// EndstopAdjustments of the XYZ axes in mm
	EndstopAdjustments []float64 `json:"endstopAdjustments"`
	// Tilt values of the XY axes
	Tilt []float64 `json:"tilt"`
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
	// Spacing between the probe points for delta kinematics
	Spacing float64 `json:"spacing"`
}
