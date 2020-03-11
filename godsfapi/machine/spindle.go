package machine

// Spindle holds information about a CNC spindle
type Spindle struct {
	// Active RPM
	Active float64
	// Current RPM
	Current float64
}
