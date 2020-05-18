package limits

// Limits configured for the machine
type Limits struct {
	// Axes is the maximum number of axes or nil if unknown
	Axes *int64 `json:"axes"`
	// AxesPlusExtruders is the maximum number of axes + extruders or nil if unknown
	AxesPlusExtruders *int64 `json:"axesPlusExtruders"`
	// BedHeaters is the maximum number of bed heaters or nil if unknown
	BedHeaters *int64 `json:"bedHeaters"`
	// Boards is the maximum number of boards or nil if unknown
	Boards *int64 `json:"boards"`
	// ChamberHeaters is the maximum number of chamber heaters or nil if unknown
	ChamberHeaters *int64 `json:"chamberHeaters"`
	// Drivers is the maximum number of drivers or nil if unknown
	Drivers *int64 `json:"drivers"`
	// DriversPerAxis is the maximum number of drivers per axis of nil if unknown
	DriverPerAxis *int64 `json:"driverPerAxis"`
	// Extruders is the maxmimum number of extruders or nil if unknown
	Extruders *int64 `json:"extruders"`
	// ExtruderPerTool is the maximum number of extruders per tool or nil if unknown
	ExtruderPerTool *int64 `json:"extruderPerTool"`
	// Fans is the maximum number of fans of nil if unknown
	Fans *int64 `json:"fans"`
	// GpInPorts is the maxmimum number of general-purpose input ports or nil if unknown
	GpInPorts *int64 `json:"gpInPorts"`
	// GpOutPorts is the maxmimum number of general-purpose output ports or nil if unknown
	GpOutPorts *int64 `json:"gpOutPorts"`
	// Heaters is the maxmimum number of heaters or nil if unknown
	Heaters *int64 `json:"heaters"`
	// HeatersPerTool is the maximum number of heaters per tool or nil if unknown
	HeatersPerTool *int64 `json:"heatersPerTool"`
	// MonitorsPerHeater is the maximum number of monitors per heater or nil if unknown
	MonitorsPerHeater *int64 `json:"monitorsPerHeater"`
	// RestorePoints is the maximum number of restore points or nil if unknown
	RestorePoints *int64 `json:"restorePoints"`
	// Sensors is the maxmimum number of sensors or nil if unknown
	Sensors *int64 `json:"sensors"`
	// Spindles is the maxmimum number of spindles or nil if unknown
	Spindles *int64 `json:"spindles"`
	// Tools is the maximum number of tools or nil if unknown
	Tools *int64 `json:"tools"`
	// TrackedObjects is the maximum number of tracked objects or nil if unknown
	TrackedObjects *int64 `json:"trackedObjects"`
	// Triggers is the maximum number of triggers or nil if unknown
	Triggers *int64 `json:"triggers"`
	// Volumes is the maximum number of volumes or nil if unknown
	Volumes *int64 `json:"volumes"`
	// Workplaces is the maximum number of workplaces or nil if unknown
	Workplaces *int64 `json:"workplaces"`
	// ZProbeProgramBytes is the maximum number of Z-probe programming bytes or nil if unknown
	ZProbeProgramBytes *int64 `json:"zProbeProgramBytes"`
	// ZProbes is the maximum number of Z-probes or nil if unknown
	ZProbes *int64 `json:"zProbes"`
}
