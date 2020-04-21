package tool

// Default values for Tool
const (
	DefaultFilamentExtruder = -1
)

// Tool holds information about a configured tool
type Tool struct {
	// Active temperature of the tool (in degC)
	Active []float64 `json:"active"`
	// Axes associated to this tool. At present only X and Y can be mapped per tool.
	// The order is the same as the visual axes, so by default the layout is
	// [
	//   [0],        // X
	//   [1]         // Y
	// ]
	Axes [][]uint64 `json:"axes"`
	// Extruders is a list of extruder drives of this tool
	Extruders []int64 `json:"extruders"`
	// Fans is a list of associated fan indices
	Fans []int64 `json:"fans"`
	// FilamentExtruder is the extruder drive index for resolving the tool filament (-1 if undefined)
	FilamentExtruder int64 `json:"filamentExtruder"`
	// Heaters is a list of associated heater indices
	Heaters []int64 `json:"heaters"`
	// Mix ratios of the associated extruder drives
	Mix []float64 `json:"mix"`
	// Name of the tool
	Name string `json:"name"`
	// Number of the tool
	Number int64 `json:"number"`
	// Offets for this tool (in mm).
	// The list is in the same order as Move.Axes
	Offsets []float64 `json:"offsets"`
	// OffsetsProbed bitmap of the axes which were probed
	OffsetsProbed int64 `json:"offsetsProbed"`
	// Retraction are the firmware retraction settings or nil if not configured
	Retraction *ToolRetraction `json:"retraction"`
	// Standby temperature of the tool
	Standby []float64 `json:"standby"`
	// State is the current state if this tool
	State ToolState `json:"state"`
}

// ToolRetraction holds tool retraction parameters
type ToolRetraction struct {
	// ExtraRestart is the amount of additional filament to extrude when undoing a retraction (in mm)
	ExtraRestart float64 `json:"extraRestart"`
	// Length of retraction (in mm)
	Length float64 `json:"length"`
	// Speed of retraction (in mm/s)
	Speed float64 `json:"speed"`
	// UnretractSpeed (in mm/s)
	UnretractSpeed float64 `json:"unretractSpeed"`
	// ZHop is the amount of Z lift after doing a retraction (in mm)
	ZHop float64 `json:"zHop"`
}

// ToolState are the states of tool
type ToolState string

const (
	// Off for a turned off tooll
	Off ToolState = "off"
	// Active for an active tool
	Active = "active"
	// Standby for a tool in standby
	Standby = "standby"
)
