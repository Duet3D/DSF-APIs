package machine

// Electronics holds information about the electronics used
type Electronics struct {
	// Version of Duet Software Framework package
	Version string
	// Type name of the main board
	Type string
	/// ShortName is the short code of the board
	ShortName string
	// Name is the full name of the board
	Name string
	// Revision of the board
	Revision string
	// Firmware of the attached main board
	Firmware Firmware
	// ProcessorId of the main board
	ProcessorId string
	// VIn represents input voltage details of the main board in V or nil if unknown
	VIn *MinMaxCurrent
	// McuTemp represents the MCU temperature details of the main boad in degC or nil if unknown
	McuTemp *MinMaxCurrent
	// ExpansionBoards is a list of attached expansion boards
	ExpansionBoards []ExpansionBoard
}

// Firmware holds information about the firmware version
type Firmware struct {
	// Name of the firmware
	Name string
	// Version of the firmare
	Version string
	// Date the firmware was built
	Date string
}

// MinMaxCurrent represents a data structure to hold current, min and max values
type MinMaxCurrent struct {
	// Current value
	Current float64
	// Minimum value encountered
	Min float64
	// Maximum value encountered
	Max float64
}

// ExpansionBoard represents information about an attached expansion board
type ExpansionBoard struct {
	// ShortName is the short code of the board
	ShortName string
	// Name is the full name of the attached expansion board
	Name string
	// Revision of the expansion board
	Revision string
	// Firmware of the expansion board
	Firmware Firmware
	// VIn represents input voltage details of the expansion board in V or nil if unknown
	VIn *MinMaxCurrent
	// McuTemp represents the MCU temperature details of the expansion board in degC or nil if unknown
	McuTemp *MinMaxCurrent
	// MaxHeaters is the maximum number of heater that can be attached to this board
	MaxHeaters int64
	// MaxMotors is the maximum number of motors that can be attched to this board
	MaxMotors int64
}
