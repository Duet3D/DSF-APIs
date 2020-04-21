package machine

// Tool holds information about a configured tool
type Tool struct {
	// Number of the tool
	Number int64 `json:"number"`
	// Active temperature of the tool
	Active []float64 `json:"active"`
	// Standby temperature of the tool
	Standby []float64 `json:"standby"`
	// Name of the tool
	Name string `json:"name"`
	// FilamentExtruder is the extruder drive index for resolving the tool filament (-1 if undefined)
	FilamentExtruder int64 `json:"filamentExtruder"`
	// Filament is the name of the currently loaded filament
	Filament string `json:"filament"`
	// Fans is a list of associated fan indices
	Fans []int64 `json:"fans"`
	// Heaters is a list of associated heater indices
	Heaters []int64 `json:"heaters"`
	// Extruders is a list of extruder drives of this tool
	Extruders []int64 `json:"extruders"`
	// Mix ratios of the associated extruder drives
	Mix []float64 `json:"mix"`
	// Spindle index associated to this tool (-1 if none is defined)
	Spindle int64 `json:"spindle"`
	// Axes associated to this tool. At present only X and Y can be mapped per tool.
	// The order is the same as the visual axes, so by default the layout is
	// [
	//   [0],        // X
	//   [1]         // Y
	// ]
	// Make sure to set each item individually so the change events are called
	Axes [][]uint64 `json:"axes"`
	// Offets for this tool (in mm).
	// The list is in the same order as Move.Axes
	Offsets []float64 `json:"offsets"`
	// OffsetsProbed bitmap of the axes which were probed
	OffsetsProbed int64 `json:"offsetsProbed"`
}
