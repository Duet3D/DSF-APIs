package sensors

// Sensors holds information about sensors
type Sensors struct {
	// Analog is a list of analog sensors
	Analog []AnalogSensor `json:"analog"`
	// Endstops is a list of configured endstops
	Endstops []Endstop `json:"endstops"`
	// FilamentMonitors is a list of configured filament monitors
	FilamentMonitors FilamentMonitors `json:"filamentMonitors"`
	// Inputs is a list of general-purpose input ports
	GpIn []*GpInputPort `json:"gpIn"`
	// Probes is a list of configured probes
	Probes []Probe `json:"probes"`
}

// AnalogSensor represents an analog sensor
type AnalogSensor struct {
	// LastReading of the sensor (in degC) or nil if invalid
	LastReading *float64 `json:"lastReading"`
	// Name of this sensor or empty if not configured
	Name string `json:"name"`
	// Type of this sensor
	Type AnalogSensorType `json:"type"`
}

// AnalogSensorType represents supported analog sensor types
type AnalogSensorType string

// Valid AnalogSensorType values
const (
	Thermistor    AnalogSensorType = "thermistor"
	PT1000                         = "pt1000"
	MAX31865                       = "rtdmax31865"
	MAX31855                       = "thermocouplemax31855"
	MAX31856                       = "thermocouplemax31856"
	LinearAnalaog                  = "linearanalaog"
	DHT11                          = "dth11"
	DHT21                          = "dht21"
	DHT22                          = "dht22"
	DHTHumidity                    = "dhthumidity"
	CurrentLoop                    = "currentlooppyro"
	McuTemp                        = "mcutemp"
	Drivers                        = "drivers"
	DriversDuex                    = "driversduex"
	Unknown                        = "unknown"
)

// Endstop holds information about an endstop
type Endstop struct {
	// Triggered represents the curent trigger state
	Triggered bool `json:"triggered"`
	// Type of this endstop
	Type EndstopType `json:"type"`
	// Probe is the index of the use probe (only valid if Type == ZProbeAsEndstop)
	Probe *int64 `json:"probe"`
}

// EndstopType represents the type of a configured enstop
type EndstopType string

const (
	// InputPin for a generic input pin
	InputPin EndstopType = "inputPin"
	// ZProbeAsEndstop if the Z-probe acts as endstop
	ZProbeAsEndstop = "zProbeAsEndstop"
	// MotorStallAny stops all the drives when triggered
	MotorStallAny = "motorStallAny"
	// MotorStallIndividual stops individual drives when triggered
	MotorStallIndividual = "motorStallIndividual"
	// EndstopTypeUnknown is the unkown type
	EndstopTypeUnknown = "unknown"
)

// GpInputPort holds details about a general-purpose input port
type GpInputPort struct {
	// Value of this port in range 0..1
	Value float64 `json:"value"`
}

const (
	// DefaultMaxProbeCount of a probe
	DefaultMaxProbeCount = 1
	// DefaultProbingSpeed at which a probing move is performed (in mm/s)
	DefaultProbingSpeed = 2.0
	// DefaultTriggerThreshold of an analog probe
	DefaultTriggerThreshold = 500
	// DefaultTolerance for a deviation between to measurements (in mm)
	DefaultTolerance = 0.03
	// DefaultTravelSpeed between probing locations (in mm/s)
	DefaultTravelSpeed = 100.0
	// DefaultTriggerHeight at which the probe is triggered (in mm)
	DefaultTriggerHeight = 0.7
)

// Probe holds information about a configured probe
type Probe struct {
	// CalibrationTemperature in degC
	CalibrationTemperature float64 `json:"calibrationTemperature"`
	// DeployedByUser indicates if the user has deployed the probe
	DeployedByUser bool `json:"deployedByUser"`
	// DisablesHeaters is true if the heater(s) are disabled while probing
	DisablesHeaters bool `json:"disablesHeaters"`
	// DiveHeight is how far above the probe point a probing move starts (in mm)
	DiveHeight float64 `json:"diveHeight"`
	// MaxProbeCount is the maximum number of times to probe after a bad reading
	// was determined
	MaxProbeCount uint64 `json:"maxProbeCount"`
	// Offsets for X and Y (in mm)
	Offsets []float64 `json:"offsets"`
	// RecoveryTime (in s)
	RecoveryTime float64 `json:"recoveryTime"`
	// Speed at which probing is performed (in mm/s)
	Speed float64 `json:"speed"`
	// TemperatureCoefficient of the probe
	TemperatureCoefficient float64 `json:"temperatureCoefficient"`
	// Threshold at which the probe is considered to be triggered (0..1023)
	Threshold int64 `json:"threshold"`
	// Tolerance is the allowed deviation between two measurements (in mm)
	Tolerance float64 `json:"tolerance"`
	// TravelSpeed when probing multiple points (in mm/s)
	TravelSpeed float64 `json:"travelSpeed"`
	// TriggerHeight is th  Z height at which the probe is triggered (in mm)
	TriggerHeight float64 `json:"triggerHeight"`
	// Type of the configured probe
	Type ProbeType `json:"type"`
	// Value are the current analog values of the probe
	Value []int64 `json:"value"`
}

// ProbeType represents supported probe types
type ProbeType uint64

const (
	// None for no probe
	None ProbeType = iota
	// Analog is a simple unmodulated probe (like dc42's infrared probe)
	Analog
	// DumbModulated probe (like the original one shipped with RepRapPro Ormerod)
	DumbModulated
	// AlternateAnalog probe (like ultrasonic probe)
	AlternateAnalog
	// EndstopSwitch_Obsolete should not be used anymore (switch connected to endstop pin)
	EndstopSwitch_Obsolete
	// Digital is a switch that is triggered when the probe is activated (filtered)
	Digital
	// E1Switch_Obsolete should not be used anymore (switch connected to E1 endstop pin)
	E1Switch_Obsolete
	// ZSwitch_Obsolete should not be used anymore (switch connected to Z endstop pin)
	ZSwitch_Obsolete
	// UnfilteredDigital is a switch that is triggered when the probe is activated (unfiltered)
	UnfilteredDigital
	// BLTouch probe
	BLTouch
	// ZMotorStall provided by the stepper driver
	ZMotorStall
)
