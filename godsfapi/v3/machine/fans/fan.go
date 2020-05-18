package fans

// Fan represents information about an attached fan
type Fan struct {
	// ActualValue is the current speed on a scale betweem 0 to 1 or -1 if unknown
	ActualValue float64 `json:"actualValue"`
	// Blip value indicating how long the fan is supposed to run at 100%
	// when turning it on to get it started (in s)
	Blip float64 `json:"blip"`
	// Frequency is the fan PWM frequency in Hz
	Frequency float64 `json:"frequency"`
	// Max speed of this fan on a scale between 0 and 1
	Max float64 `json:"max"`
	// Min speed of this fan on a scale between 0 and 1
	Min float64 `json:"min"`
	// Name of the fan
	Name string `json:"name"`
	// RequestedValue for this fan on a scale between 0 to 1
	RequestedValue float64 `json:"requestedValue"`
	// Rpm is the current RPM of this fan or -1 if unknown/unset
	Rpm int64 `json:"rpm"`
	// Thermostatic control parameters
	Thermostatic Thermostatic `json:"thermostatic"`
}

// Thermostatic parameters of a fan
type Thermostatic struct {
	// Heaters is a list of heaters to monitor (indices)
	Heaters []int64 `json:"heaters"`
	// HighTemperature is the upper temperature range required to turn
	// on the fan (in degC)
	HighTemperature *float64 `json:"highTemperature"`
	// LowTemperature is the lower temperature range required to turn
	// on the fan (in degC)
	LowTemperature *float64 `json:"lowTemperature"`
}
