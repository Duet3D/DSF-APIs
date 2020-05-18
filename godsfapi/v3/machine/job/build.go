package job

// BuildObject holds information about a detected build object
type BuildObject struct {
	// Cancelled indicates if this build object is cancelled
	Cancelled bool `json:"cancelled"`
	// Name of the build object (if any)
	Name string `json:"name"`
	// X coordinates of the build object (in mm or nil if not found)
	X []*float64 `json:"x"`
	// Y coordinates of the build object (in mm or nil if not found)
	Y []*float64 `json:"y"`
}

// Build holds information about the current build
type Build struct {
	// CurrentObject is the index of the current object being printed
	// or -1 if unknown
	CurrentObject int64 `json:"currentObject"`
	// M486Names if M486 names are being used
	M486Names bool `json:"m486Names"`
	// M486Numbers if M486 numbers are being used
	M486Numbers bool `json:"m486Numbers"`
	// Objects is a list of detected objects
	Objects []BuildObject `json:"objects"`
}
