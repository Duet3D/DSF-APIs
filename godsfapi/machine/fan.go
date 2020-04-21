package machine

// Fan represents information about an attached fan
type Fan struct {
	// Value is the current speed on a scale betweem 0 to 1
	Value float64 `json:"value"`
	// Name of the fan
	Name string `json:"name"`
	// Rpm is the current RPM of this fan TODO: can be nil?
	Rpm int64 `json:"rpm"`
	// Inverted represents the inversion state of the fan PWM signal
	Inverted bool `json:"inverted"`
	// Frequency is the fan PWM frequency in Hz
	Frequency float64 `json:"frequency"`
	// Min speed of this fan on a scale between 0 and 1
	Min float64 `json:"min"`
	// Max speed of this fan on a scale between 0 and 1
	Max float64 `json:"max"`
	// Blip value indicating how long the fan is supposed to run at 100%
	// when turning it on to get it started (in s)
	Blip float64 `json:"blip"`
	// Thermostatic control parameters
	Thermostatic Thermostatic `json:"thermostatic"`
	// Pin number of the assigned fan
	Pin uint64 `json:"pin"`
}

// Thermostatic parameters of a fan
type Thermostatic struct {
	// Control represents whether thermostatic control is enabled
	Control bool `json:"control"`
	// Heaters is a list of heaters to minitor
	Heaters []int64 `json:"heaters"`
	// Temperature at which the fan will be turned on in degC
	Temperature float64 `json:"temperature"`
}
