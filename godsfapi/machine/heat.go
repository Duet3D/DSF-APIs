package machine

const (
	// AbsoluteZero temperature in degC
	AbsoluteZero = -273.15
)

// HeaterState represents available heater states
type HeaterState uint64

const (
	// Off for a turned off heater
	Off HeaterState = iota
	// Standby for a heater in standby mode
	Standby
	// Active for an active heater
	Active
	// Tuning for a tuning heater
	Tuning
	// Offline for a heater that cannot be reached
	Offline
)

// Heat holds information about the heat subsystem
type Heat struct {
	// Beds is a list of configured beds
	Beds []BedOrChamber
	// Chambers is a list of configured chambers
	Chambers []BedOrChamber
	// ColdExtrusionTemp is the minimum temperature required for extrusion moves (in degC)
	ColdExtrusionTemp float64
	// ColdRetractTemp is the minimum temperature required for retraction moves (in degC)
	ColdRetractTemp float64
	// Extra is a list of configured extra heaters
	Extra []ExtraHeater
	// Heaters is a list of configured heaters
	Heaters []Heater
}

// BedOrChamber holds information about a bed or a chamber heater
type BedOrChamber struct {
	// Active temperatures (in degC)
	Active []float64
	// Standby temperatures (in degC)
	Standby []float64
	// Name of the bed or chamber
	Name string
	// Heaters is a list of heater indices controlled by this bed or chamber
	Heaters []int64
}

// ExtraHeater holds information about an extra heater (virtual)
type ExtraHeater struct {
	// Current temperature (in degC)
	Current float64
	// Name of the extra heater
	Name string
	// State of the exra heater or nil if unknown/unset
	State *HeaterState
	// Sensor number of the extra heater or nil if unknown/unset
	Sensor *int64
}

// Heater holds information about a heater
type Heater struct {
	// Current temperature (in degC)
	Current float64
	// Name of the heater
	Name string
	// State of the heater
	State HeaterState
	// Model hold information about the heater model
	Model HeaterModel
	// Max allowed temperature of this heater (in degC)
	Max float64
	// Sensor number of this heater or nil if unknown
	Sensor *int64
}

// HeaterModel holds information about the way a heater heats up
type HeaterModel struct {
	// Gain value or nil if unknown
	Gain *float64
	// TimeConstant or nil if unknown
	TimeConstant *float64
	// DeadTime of this heater or nil if unknown
	DeadTime *float64
	// MaxPwm value for this heater (0 if unknown)
	MaxPwm float64
	// StandardVoltage of this heater (0 if unknown)
	StandardVoltage float64
	// UsePID indicates usage of PID (instead of bang-bang)
	UsePID bool
	// CustomPID indicates the usage of custom PID values
	CustomPID bool
	// P is the proportional value of the PID regulator
	P float64
	// I is the integral value of the PID regulator
	I float64
	// D is the derivative value pf the PID regulator
	D float64
}
