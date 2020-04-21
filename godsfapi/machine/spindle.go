package machine

// Spindle holds information about a CNC spindle
type Spindle struct {
	// Active RPM
	Active float64 `json:"active"`
	// Current RPM
	Current float64 `json:"current"`
}
