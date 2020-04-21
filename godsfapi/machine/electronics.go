package machine

// Electronics holds information about the electronics used
type Electronics struct {
	// Version of Duet Software Framework package
	Version string `json:"version"`
	// Type name of the main board
	Type string `json:"type"`
	/// ShortName is the short code of the board
	ShortName string `json:"shortName"`
	// Name is the full name of the board
	Name string `json:"name"`
	// Revision of the board
	Revision string `json:"revision"`
	// Firmware of the attached main board
	Firmware Firmware `json:"firmware"`
	// ProcessorId of the main board
	ProcessorId string `json:"processorId"`
	// VIn represents input voltage details of the main board in V or nil if unknown
	VIn *MinMaxCurrent `json:"vIn"`
	// McuTemp represents the MCU temperature details of the main boad in degC or nil if unknown
	McuTemp *MinMaxCurrent `json:"mcuTemp"`
	// ExpansionBoards is a list of attached expansion boards
	ExpansionBoards []ExpansionBoard `json:"expansionBoards"`
}

// Firmware holds information about the firmware version
type Firmware struct {
	// Name of the firmware
	Name string `json:"name"`
	// Version of the firmare
	Version string `json:"version"`
	// Date the firmware was built
	Date string `json:"date"`
}

// MinMaxCurrent represents a data structure to hold current, min and max values
type MinMaxCurrent struct {
	// Current value
	Current float64 `json:"current"`
	// Minimum value encountered
	Min float64 `json:"min"`
	// Maximum value encountered
	Max float64 `json:"max"`
}

// ExpansionBoard represents information about an attached expansion board
type ExpansionBoard struct {
	// ShortName is the short code of the board
	ShortName string `json:"shortName"`
	// Name is the full name of the attached expansion board
	Name string `json:"name"`
	// Revision of the expansion board
	Revision string `json:"revision"`
	// Firmware of the expansion board
	Firmware Firmware `json:"firmware"`
	// VIn represents input voltage details of the expansion board in V or nil if unknown
	VIn *MinMaxCurrent `json:"vIn"`
	// McuTemp represents the MCU temperature details of the expansion board in degC or nil if unknown
	McuTemp *MinMaxCurrent `json:"mcuTemp"`
	// MaxHeaters is the maximum number of heater that can be attached to this board
	MaxHeaters int64 `json:"maxHeaters"`
	// MaxMotors is the maximum number of motors that can be attched to this board
	MaxMotors int64 `json:"maxMotors"`
}
