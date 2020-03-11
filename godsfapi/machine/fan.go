package machine

// Fan represents information about an attached fan
type Fan struct {
	// Value is the current speed on a scale betweem 0 to 1
	Value float64
	// Name of the fan
	Name string
	// Rpm is the current RPM of this fan TODO: can be nil?
	Rpm int64
	// Inverted represents the inversion state of the fan PWM signal
	Inverted bool
	// Frequency is the fan PWM frequency in Hz
	Frequency float64
	// Min speed of this fan on a scale between 0 and 1
	Min float64
	// Max speed of this fan on a scale between 0 and 1
	Max float64
	// Blip value indicating how long the fan is supposed to run at 100%
	// when turning it on to get it started (in s)
	Blip float64
	// Thermostatic control parameters
	Thermostatic Thermostatic
	// Pin number of the assigned fan
	Pin uint64
}

// Thermostatic parameters of a fan
type Thermostatic struct {
	// Control represents whether thermostatic control is enabled
	Control bool
	// Heaters is a list of heaters to minitor
	Heaters []int64
	// Temperature at which the fan will be turned on in degC
	Temperature float64
}
