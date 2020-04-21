package scanner

// Scanner holds information about the 3D scanner subsytem
type Scanner struct {
	// Progress of the current action on scale between 0 and 1
	Progress float64 `json:"progress"`
	// Status of the 3D scanner
	Status ScannerStatus `json:"status"`
}

// ScannerStatus represents possible states of an attached 3D scanner
type ScannerStatus string

const (
	// Disconnected if the scanner is not present
	Disconnected ScannerStatus = "D"
	// Idle for a scanner that is registered and idle
	Idle = "I"
	// Scanning while the scanner is scanning
	Scanning = "S"
	// PostProcessing while the scanner is post-processing a file
	PostProcessing = "P"
	// Calibrating while the scanner is calibrating
	Calibrating = "C"
	// Uploading while the scanner is uploading
	Uploading = "U"
)
