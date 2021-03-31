package spindles

const (
	// DefaultMaxRpm is the maximum RPM of a spindle
	DefaultMaxRpm = 10000.0
	// DefaultTool mapping for a spindle
	DefaultTool = -1
)

// Spindle holds information about a CNC spindle
type Spindle struct {
	// Active RPM
	Active float64 `json:"active"`
	// Current RPM
	Current float64 `json:"current"`
	// Frequency in Hz
	Frequency int64 `json:"frequency"`
	// Min RPM when turned on
	Min float64 `json:"min"`
	// Max RPM
	Max float64 `json:"max"`
	// State is the current state
	State SpindleState `json:"state"`
}

// SpindleState are the possible states of a spindle
type SpindleState string

const (
	// Unconfigured if the spindle is not configured
	Unconfigured SpindleState = "unconfigured"
	// Stopped if the spindle is stopped (inactive)
	Stopped = "stopped"
	// Forward if the spindle spins clockwise
	Forward = "forward"
	// Reverse if the spindle spins counterclockwise
	Reverse = "reverse"
)
