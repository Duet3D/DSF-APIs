package machine

// Tool holds information about a configured tool
type Tool struct {
	// Number of the tool
	Number int64
	// Active temperature of the tool
	Active []float64
	// Standby temperature of the tool
	Standby []float64
	// Name of the tool
	Name string
	// FilamentExtruder is the extruder drive index for resolving the tool filament (-1 if undefined)
	FilamentExtruder int64
	// Filament is the name of the currently loaded filament
	Filament string
	// Fans is a list of associated fan indices
	Fans []int64
	// Heaters is a list of associated heater indices
	Heaters []int64
	// Extruders is a list of extruder drives of this tool
	Extruders []int64
	// Mix ratios of the associated extruder drives
	Mix []float64
	// Spindle index associated to this tool (-1 if none is defined)
	Spindle int64
	// Axes associated to this tool. At present only X and Y can be mapped per tool.
	// The order is the same as the visual axes, so by default the layout is
	// [
	//   [0],        // X
	//   [1]         // Y
	// ]
	// Make sure to set each item individually so the change events are called
	Axes [][]uint64
	// Offets for this tool (in mm).
	// The list is in the same order as Move.Axes
	Offsets []float64
	// OffsetsProbed bitmap of the axes which were probed
	OffsetsProbed int64
}
