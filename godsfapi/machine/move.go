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
	Axes []Axis
	// BabystepZ is the current babystep amount in Z direction in mm
	BabystepZ float64
	// CurrentMove holds information about the current move
	CurrentMove CurrentMove
	// Compensation is the name of the currently used bed compensation
	// (one of "Mesh", "[n] Point", "None")
	Compensation string
	// HeightmapFile is the path to the current heightmap file if Compensation is "Mesh"
	HeightmapFile string
	// Drives is a list of configured drives
	Drives []Drive
	// Extruders is a list of configured extrudersr
	Extruders []Extruder
	// Kinematics holds information about the currently configured kinematics
	Kinematics Kinematics `json:"Geometry"`
	// Idle current reduction parameters
	Idle MotorsIdleControl
	// ProbeGrid holds information about the configured mesh compensation (see M557)
	ProbeGrid ProbeGrid
	// SpeedFactor applied to every regular move (1.0 equals 100%)
	SpeedFactor float64
	// CurrentWorkplace is the index of the selected workspace
	CurrentWorkplace int64
	// WorkplaceCoordinates are the axis offsets of each available workspace in mm
	WorkplaceCoordinates [][]float64
}

// Axis holds information about a configured axis
type Axis struct {
	// Letter assigned to this axis (always upper-case)
	Letter string
	// Drives is a list of drive indices assigned to this axis
	Drives []int64
	// Homed is true if the axis has been homed
	Homed bool
	// MachinePosition is the current machine position (in mm or nil if unknown)
	MachinePosition *float64
	// Min travel of this axis (in mm or nil if unknown)
	Min *float64
	// MinEndstop is the index of the endstop that is used for the low end
	// or nil if none is configured
	MinEndstop *int64
	// MinProbed is true if the minimum was probed
	MinProbed bool
	// Max travel of this axis (in mm or nil if unknown)
	Max *float64
	// MaxEndstop is the index of the endstop that is used for the high end
	// or nil if none is configured
	MaxEndstop *int64
	// MaxProbed is ture if the maximum was probed
	MaxProbed bool
	// Visible is true if the axis is not explicitely hidden
	Visible bool
}

// CurrentMove holds information about the current move
type CurrentMove struct {
	// RequestedSpeed of the current move (in mm/s)
	RequestedSpeed float64
	// TopSpeed actually reached for the current move (in mm/s)
	TopSpeed float64
}

// Drive holds information about a drive
type Drive struct {
	// Position is the current user position of this drive (in mm)
	Position float64
	// Microstepping configured for this drive
	Microstepping DriveMicrostepping
	// Current configured for this drive (in mA)
	Current uint64
	// Acceleration of this drive (in mm/s²)
	Acceleration float64
	// MinSpeed allowed for this drive (in mm/s)
	MinSpeed float64
	// MaxSpeed allowed for this drive (in mm/s)
	MaxSpeed float64
}

// DriveMicrostepping holds information about configured microstepping
type DriveMicrostepping struct {
	// Value of the microstepping (e.g. x16)
	Value uint64
	// Interpolated is true if interpolation is used
	Interpolated bool
}

// Extruder holds information about an extruder drive
type Extruder struct {
	// Drives is a list of drive indices of this extruder
	Drives []int64
	// Factor is the extrusion factor (1.0 equals 100%)
	Factor float64
	// NonLinear extrusion parameters (see M592)
	NonLinear ExtruderNonLinear
}

// ExtruderNonLinear contains non-linear extrusion parameters (see M592)
type ExtruderNonLinear struct {
	// A coefficient in the extrusion formula
	A float64
	// B coefficient in the extrusion formula
	B float64
	// UpperLimit of the nonlinear extrusion compensation
	UpperLimit float64
	// Temperature at which these values are valid in degC (future use only)
	Temperature float64
}

// Kinematics holds information about the configured kinematics
// Note to developers: this is called "Geometry" upstream
type Kinematics struct {
	// Type of currently configured kinematics
	Type KinematicsType
	// Anchors of a hangprinter A, B, C, Dz (10 values)
	Anchors []float64
	// PrintRadius for Hangprinter and Delta kinematics in mm
	PrintRadius float64
	// Diagonals for a delta
	Diagonals []float64
	// Radius of a delta in mm
	Radius float64
	// HomedHeight of a delta in mm
	HomedHeight float64
	// AngleCorrections ABC for delta kinematics
	AngleCorrections []float64
	// EndstopAdjustments of the XYZ axes in mm
	EndstopAdjustments []float64
	// Tilt values of the XY axes
	Tilt []float64
}

// MotorsIdleControl holds idle factor parameters for automatic MotorsIdleControl
// current reduction
type MotorsIdleControl struct {
	// Timeout after which the motor currents are reduced (in s)
	Timeout float64
	// Factor of the reduction on a scale between 0 and 1
	Factor float64
}

// ProbeGrid holds information about the configured probe grid (see M557)
type ProbeGrid struct {
	// XMin is the X start coordinate of the heightmap
	XMin float64
	// XMax is the X end coordinate of the heightmap
	XMax float64
	// XSpacing is the spacing between probe points in X direction
	XSpacing float64
	// YMin is the Y start coordinate of the heightmap
	YMin float64
	// YMax is the Y end coordinate of the heightmap
	YMax float64
	// YSpacing is the spacing between probe points in Y direction
	YSpacing float64
	// Radius is the probing radius on delta kinematics
	Radius float64
	// Spacing between the probe points for delta kinematics
	Spacing float64
}
