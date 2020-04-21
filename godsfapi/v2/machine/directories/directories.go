package directories

// Default values
const (
	DefaultFilamentsPath = "0:/filaments"
	DefaultFirmwarePath  = "0:/sys"
	DefaultGCodesPath    = "0:/gcodes"
	DefaultMacrosPath    = "0:/macros"
	DefaultMenuPath      = "0:/menu"
	DefaultScansPath     = "0:/scans"
	DefaultSystemPath    = "0:/sys"
	DefaultWebPath       = "0:/www"
)

// Directories holds information about the directory structure
type Directories struct {
	// Filaments is the path to filaments directory
	Filaments string `json:"filaments"`
	// Firmware is the path to firmware directory
	Firmware string `json:"firmware"`
	// GCodes is the path to the gcodes directory
	GCodes string `json:"gCodes"`
	// Macros is the path to the macros directory
	Macros string `json:"macros"`
	// Menu is the path to the menu directory (12864 displays)
	Menu string `json:"menu"`
	// Scans is the path to scans directory
	Scans string `json:"scans"`
	// System is the path to the sys directory
	System string `json:"system"`
	// Web is the path to the web directory
	Web string `json:"web"`
}

// NewDirectories returns an instance with all paths set to their defaults
func NewDirectories() *Directories {
	return &Directories{
		Filaments: DefaultFilamentsPath,
		Firmware:  DefaultFirmwarePath,
		GCodes:    DefaultGCodesPath,
		Macros:    DefaultMacrosPath,
		Menu:      DefaultMenuPath,
		Scans:     DefaultScansPath,
		System:    DefaultSystemPath,
		Web:       DefaultWebPath,
	}
}
