package limits

// Limits configured for the machine
type Limits struct {
	// Axes is the maximum number of axes or nil if unknown
	Axes *int `json:"axes"`
	// AxesPlusExtruders is the maximum number of axes + extruders or nil if unknown
	AxesPlusExtruders *int `json:"axesPlusExtruders"`
	// BedHeaters is the maximum number of bed heaters or nil if unknown
	BedHeaters *int `json:"bedHeaters"`
	// Boards is the maximum number of boards or nil if unknown
	Boards *int `json:"boards"`
	// ChamberHeaters is the maximum number of chamber heaters or nil if unknown
	ChamberHeaters *int `json:"chamberHeaters"`
	// Drivers is the maximum number of drivers or nil if unknown
	Drivers *int `json:"drivers"`
	// DriversPerAxis is the maximum number of drivers per axis of nil if unknown
	DriverPerAxis *int `json:"driverPerAxis"`
	// Extruders is the maxmimum number of extruders or nil if unknown
	Extruders *int `json:"extruders"`
	// ExtruderPerTool is the maximum number of extruders per tool or nil if unknown
	ExtruderPerTool *int `json:"extruderPerTool"`
	// Fans is the maximum number of fans of nil if unknown
	Fans *int `json:"fans"`
	// GpInPorts is the maxmimum number of general-purpose input ports or nil if unknown
	GpInPorts *int `json:"gpInPorts"`
	// GpOutPorts is the maxmimum number of general-purpose output ports or nil if unknown
	GpOutPorts *int `json:"gpOutPorts"`
	// Heaters is the maxmimum number of heaters or nil if unknown
	Heaters *int `json:"heaters"`
	// HeatersPerTool is the maximum number of heaters per tool or nil if unknown
	HeatersPerTool *int `json:"heatersPerTool"`
	// MonitorsPerHeater is the maximum number of monitors per heater or nil if unknown
	MonitorsPerHeater *int `json:"monitorsPerHeater"`
	// Sensors is the maxmimum number of sensors or nil if unknown
	Sensors *int `json:"sensors"`
	// Spindles is the maxmimum number of spindles or nil if unknown
	Spindles *int `json:"spindles"`
	// Triggers is the maximum number of triggers or nil if unknown
	Triggers *int `json:"triggers"`
	// Volumes is the maximum number of volumes or nil if unknown
	Volumes *int `json:"volumes"`
	// Workplaces is the maximum number of workplaces or nil if unknown
	Workplaces *int `json:"workplaces"`
	// ZProbeProgramBytes is the maximum number of Z-probe programming bytes or nil if unknown
	ZProbeProgramBytes *int `json:"zProbeProgramBytes"`
	// ZProbes is the maximum number of Z-probes or nil if unknown
	ZProbes *int `json:"zProbes"`
}
