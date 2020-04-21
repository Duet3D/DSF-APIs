package machine

import "github.com/Duet3D/DSF-APIs/godsfapi/types"

// Job holds information about the current file job (if any)
type Job struct {
	// File holds ParsedFileInfo about the file being processed
	File types.ParsedFileInfo `json:"file"`
	// FilePosition is the current position in the file being processed in bytes
	FilePosition uint64 `json:"filePosition"`
	// LastFileName is the name of the last processed file
	LastFileName string `json:"lastFileName"`
	// LastFileAborted indicated if the last file was aborted (unexpected cancellation)
	LastFileAborted bool `json:"lastFileAborted"`
	// LastFileCancelled indicates if the last file was cancelled by the user
	LastFileCancelled bool `json:"lastFileCancelled"`
	// LastFileSimulated indicates if the last file was simulated
	LastFileSimulated bool `json:"lastFileSimulated"`
	// ExtrudedRaw is a list of virtual amounts of extruded filament according to the
	// G-code file in mm
	ExtrudedRaw []float64 `json:"extrudedRaw"`
	// Duration is the total duration of the current job in s
	Duration float64 `json:"duration"`
	// Layer number of the current layer or 0 if none has been started yet
	Layer float64 `json:"layer"`
	// LayerTime is time elapsed since the beginning of the current layer in s
	LayerTime float64 `json:"layerTime"`
	// Layers is a list of Layer information about past layers
	Layers []Layer `json:"layers"`
	// WarmUpDuration is the time needed to heat up the heaters in se
	WarmUpDuration float64 `json:"warmUpDuration"`
	// TimesLeft contains estimated remaining times
	TimesLeft TimesLeft `json:"timesLeft"`
}

// Layer holds information about a layer from a file being printed
type Layer struct {
	// Duration of the layer in s (nil if unknown)
	Duration *float64 `json:"duration"`
	// Height of the layer in mm (0 if unknown)
	Height float64 `json:"height"`
	// Filament represents the actual amount of filament extruded during
	// this layer in mm
	Filament []float64 `json:"filament"`
	// FractionPrinted represents the fraction of the whole file printed
	// during this layer on a scale between 0 and 1
	FractionPrinted float64 `json:"fractionPrinted"`
}

// TimesLeft holds information about estimated remaining times
type TimesLeft struct {
	// File progress based estimation in s (nil if unknown)
	File *float64 `json:"file"`
	// Filament consumption based estimation in s (nil if unknown)
	Filament *float64 `json:"filament"`
	// Layer progress based estimation in s (nil if unknown)
	Layer *float64 `json:"layer"`
}
