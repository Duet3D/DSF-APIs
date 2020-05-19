package boards

// Board holds information about the electronics used
type Board struct {
	// BootloaderFileName is filename of firmware binary
	BootloaderFileName string `json:"bootloaderFileName"`
	// CanAddress of this board or nil if not applicable
	CanAddress *int64 `json:"canAddress"`
	// FirmwareDate is the date of the build
	FirmwareDate string `json:"firmwareDate"`
	// FirmwareFileName is filename of the binary
	FirmwareFileName string `json:"firmwareFileName"`
	// FirmwareName is the name of the build
	FirmwareName string `json:"firmwareName"`
	// FirmwareVersion of the buld
	FirmwareVersion string `json:"firmwareVersion"`
	// IapFileNameSBC is the filename of the IAP binary that is used
	// for updates from the SBC or empty if unsupported
	IapFileNameSBC string `json:"iapFileNameSbc"`
	// IapFileNameSD is the filname of the IAP binary that is used
	// for updates from the SD card or empty if unsupported
	IapFileNameSD string `json:"iapFileNameSd"`
	// MaxHeaters is the maximum number of heaters this board can control
	MaxHeaters int64 `json:"maxHeaters"`
	// MaxMotors is the maximum number of motors this board can drive
	MaxMotors int64 `json:"maxMotors"`
	// McuTemp represents the MCU temperature details of the main boad in degC or nil if unknown
	McuTemp *MinMaxCurrent `json:"mcuTemp"`
	// Name is the full name of the board
	Name string `json:"name"`
	/// ShortName is the short code of the board
	ShortName string `json:"shortName"`
	// State of this board
	State BoardState `json:"state"`
	// Supports12864 indicates if this board supports external 12864 displays
	Supports12864 bool `json:"supports12864"`
	// UniqueId of the board
	UniqueId string `json:"uniqueId"`
	// V12 represents 12V rail details of the main board in V or nil if unknown
	V12 *MinMaxCurrent `json:"v12"`
	// VIn represents input voltage details of the main board in V or nil if unknown
	VIn *MinMaxCurrent `json:"vIn"`
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

// BoardState is a representation of possible expansion board states
type BoardState string

const (
	// Unknown state
	Unknown BoardState = "unknown"
	// Flashing new firmware
	Flashing = "flashing"
	// FlashFailed if flashing new firmware failed
	FlashFailed = "flashingFailed"
	// Resetting if the board is being reset
	Resetting = "resetting"
	// Running if the board is up and running
	Running = "running"
)
