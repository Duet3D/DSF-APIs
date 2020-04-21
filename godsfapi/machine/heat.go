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
	Beds []BedOrChamber `json:"beds"`
	// Chambers is a list of configured chambers
	Chambers []BedOrChamber `json:"chambers"`
	// ColdExtrusionTemp is the minimum temperature required for extrusion moves (in degC)
	ColdExtrusionTemp float64 `json:"coldExtrusionTemp"`
	// ColdRetractTemp is the minimum temperature required for retraction moves (in degC)
	ColdRetractTemp float64 `json:"coldRetractTemp"`
	// Extra is a list of configured extra heaters
	Extra []ExtraHeater `json:"extra"`
	// Heaters is a list of configured heaters
	Heaters []Heater `json:"heaters"`
}

// BedOrChamber holds information about a bed or a chamber heater
type BedOrChamber struct {
	// Active temperatures (in degC)
	Active []float64 `json:"active"`
	// Standby temperatures (in degC)
	Standby []float64 `json:"standby"`
	// Name of the bed or chamber
	Name string `json:"name"`
	// Heaters is a list of heater indices controlled by this bed or chamber
	Heaters []int64 `json:"heaters"`
}

// ExtraHeater holds information about an extra heater (virtual)
type ExtraHeater struct {
	// Current temperature (in degC)
	Current float64 `json:"current"`
	// Name of the extra heater
	Name string `json:"name"`
	// State of the exra heater or nil if unknown/unset
	State *HeaterState `json:"state"`
	// Sensor number of the extra heater or nil if unknown/unset
	Sensor *int64 `json:"sensor"`
}

// Heater holds information about a heater
type Heater struct {
	// Current temperature (in degC)
	Current float64 `json:"current"`
	// Name of the heater
	Name string `json:"name"`
	// State of the heater
	State HeaterState `json:"state"`
	// Model hold information about the heater model
	Model HeaterModel `json:"model"`
	// Max allowed temperature of this heater (in degC)
	Max float64 `json:"max"`
	// Sensor number of this heater or nil if unknown
	Sensor *int64 `json:"sensor"`
}

// HeaterModel holds information about the way a heater heats up
type HeaterModel struct {
	// Gain value or nil if unknown
	Gain *float64 `json:"gain"`
	// TimeConstant or nil if unknown
	TimeConstant *float64 `json:"timeConstant"`
	// DeadTime of this heater or nil if unknown
	DeadTime *float64 `json:"deadTime"`
	// MaxPwm value for this heater (0 if unknown)
	MaxPwm float64 `json:"maxPwm"`
	// StandardVoltage of this heater (0 if unknown)
	StandardVoltage float64 `json:"standardVoltage"`
	// UsePID indicates usage of PID (instead of bang-bang)
	UsePID bool `json:"usePID"`
	// CustomPID indicates the usage of custom PID values
	CustomPID bool `json:"customPID"`
	// P is the proportional value of the PID regulator
	P float64 `json:"p"`
	// I is the integral value of the PID regulator
	I float64 `json:"i"`
	// D is the derivative value pf the PID regulator
	D float64 `json:"d"`
}
