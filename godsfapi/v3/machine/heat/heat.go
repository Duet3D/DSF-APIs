package heat

const (
	// AbsoluteZero temperature in degC
	AbsoluteZero = -273.15
)

// HeaterState represents available heater states
type HeaterState string

const (
	// Off for a turned off heater
	Off HeaterState = "off"
	// Standby for a heater in standby mode
	Standby = "standby"
	// Active for an active heater
	Active = "active"
	// Fault for a faulted heater
	Fault = "fault"
	// Tuning for a tuning heater
	Tuning = "tuning"
	// Offline for a heater that cannot be reached
	Offline = "offline"
)

// Default values for Heat
const (
	DefaultColdExtrudeTemperature = 160.0
	DefaultColdRetractTemperature = 90.0
)

// Heat holds information about the heat subsystem
type Heat struct {
	// Beds is a list of configured beds (indices)
	// Items of -1 indicate no element with at that index
	Beds []int64 `json:"beds"`
	// ChamberHeaters is a list of configured chamber heaters (indices)
	// Items of -1 indicate no element with at that index
	ChamberHeaters []int64 `json:"chamberHeaters"`
	// ColdExtrudeTemperature is the minimum temperature required for extrusion moves (in degC)
	ColdExtrudeTemperature float64 `json:"coldExtrudeTemperature"`
	// ColdRetractTemperature is the minimum temperature required for retraction moves (in degC)
	ColdRetractTemperature float64 `json:"coldRetractTemperature"`
	// Heaters is a list of configured heaters
	Heaters []Heater `json:"heaters"`
}

// Default values for Heater
const (
	DefaultMaxTemp = 285.0
	DefaultMinTemp = -10.0
)

// Heater holds information about a heater
type Heater struct {
	// Active temperature (in degC)
	Active float64 `json:"active"`
	// Current temperature (in degC)
	Current float64 `json:"current"`
	// Max temperature allowed for this heater (in degC)
	Max float64 `json:"max"`
	// Min temperature allowed for this heater (in degC)
	Min float64 `json:"min"`
	// Model hold information about the heater model
	Model HeaterModel `json:"model"`
	// Monitors of this heater
	Monitors []HeaterMonitor `json:"monitors"`
	// Name of the heater
	Name string `json:"name"`
	// Sensor number of this heater or -1 if not configured
	Sensor int64 `json:"sensor"`
	// Standby temperature for this heater (in degC)
	Standby float64 `json:"standby"`
	// State of the heater
	State *HeaterState `json:"state"`
}

// Default values for HeaterModel
const (
	DefaultDeadTime     = 5.5
	DefaultGain         = 340.0
	DefaultMaxPwm       = 1.0
	DefaultTimeConstant = 140.0
)

// HeaterModel holds information about the way a heater heats up
type HeaterModel struct {
	// DeadTime value
	DeadTime float64 `json:"deadTime"`
	// Enabled indicates if this heater is enabled
	Enabled bool `json:"enabled"`
	// Gain value
	Gain float64 `json:"gain"`
	// Inverted if the heater PWM signal is Inverted
	Inverted bool `json:"inverted"`
	// MaxPwm value for this heater (0 if unknown)
	MaxPwm float64 `json:"maxPwm"`
	// PID holds details about the PID controller
	PID HeaterModelPID `json:"pid"`
	// StandardVoltage or nil if unknown
	StandardVoltage *float64 `json:"standardVoltage"`
	// TimeConstant value
	TimeConstant float64 `json:"timeConstant"`
}

// HeaterModelPID holds details about the PID model of a heater
type HeaterModelPID struct {
	// Overridden indicates the usage of custom PID values
	Overridden bool `json:"overridden"`
	// P is the proportional value of the PID regulator
	P float64 `json:"p"`
	// I is the integral value of the PID regulator
	I float64 `json:"i"`
	// D is the derivative value pf the PID regulator
	D float64 `json:"d"`
	// Used indicates usage of PID control (instead of bang-bang)
	Used bool `json:"used"`
}

// HeaterMonitorAction is the action to take when a heater monitor is triggered
type HeaterMonitorAction int64

const (
	// GenerateFault generates a heater fault
	GenerateFault HeaterMonitorAction = iota
	// PermanentSwitchOff switches off the heater permanently
	PermanentSwitchOff
	// TemporarySwitchOff switch off the heater unilthe condition is no longer met
	TemporarySwitchOff
	// ShutDown the printer
	ShutDown
)

// HeaterMonitorCondition is the trigger condition for a heater monitor
type HeaterMonitorCondition string

const (
	// Disabled for a disabled heater monitor
	Disabled HeaterMonitorCondition = "disabled"
	// TooHigh if limit temperature has been exceeded
	TooHigh = "tooHigh"
	// TooLow if limit temperature has been undercut
	TooLow = "tooLow"
	// Undefined for unknown condition
	Undefined = "undefined"
)

// HeaterMonitor holds information about a heater monitor
type HeaterMonitor struct {
	// Action to perfrm when the trigger condition is met
	Action *HeaterMonitorAction `json:"action"`
	// Condition to meet to perform an action
	Condition HeaterMonitorCondition `json:"condition"`
	// Limit threshold for this heater monitor
	Limit *float64 `json:"limit"`
}
