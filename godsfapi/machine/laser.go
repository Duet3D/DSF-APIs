package machine

// Laser holds information about an attached laser diode
type Laser struct {
	// ActualPwm on a scale between 0 and 1
	ActualPwm float64 `json:"actualPwm"`
	// RequestedPwm on a scale between 0 and 1 from a G1 move
	RequestedPwm float64 `json:"requestedPwm"`
}
