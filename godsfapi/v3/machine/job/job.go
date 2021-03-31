package job

// Job holds information about the current file job (if any)
type Job struct {
	// Build holds information about the current build or nil if not available
	Build *Build `json:"build"`
	// Duration is the total duration of the current job in s
	Duration *int `json:"duration"`
	// File holds ParsedFileInfo about the file being processed
	File ParsedFileInfo `json:"file"`
	// FilePosition is the current position in the file being processed in bytes
	FilePosition *uint64 `json:"filePosition"`
	// FirstLayerDuration is the duration of the first layer in s or nil if not available
	// Deprecated: No longer used, will always be nil
	FirstLayerDuration *int64 `json:"-"`
	// LastDuration is the total duration of the last job in s or nil if not available
	LastDuration *int64 `json:"lastDuration"`
	// LastFileName is the name of the last processed file
	LastFileName string `json:"lastFileName"`
	// LastFileAborted indicated if the last file was aborted (unexpected cancellation)
	LastFileAborted bool `json:"lastFileAborted"`
	// LastFileCancelled indicates if the last file was cancelled by the user
	LastFileCancelled bool `json:"lastFileCancelled"`
	// LastFileSimulated indicates if the last file was simulated
	LastFileSimulated bool `json:"lastFileSimulated"`
	// Layer number of the current layer or nil if none has been started yet
	Layer *int64 `json:"layer"`
	// LayerTime is time elapsed since the beginning of the current layer in s or nil if unknown
	LayerTime *float64 `json:"layerTime"`
	// Layers is a list of Layer information about past layers
	Layers []Layer `json:"layers"`
	// PauseDuration is total pause time since job stareted
	PauseDuration *int64 `json:"pauseDuration"`
	// TimesLeft contains estimated remaining times
	TimesLeft TimesLeft `json:"timesLeft"`
	// WarmUpDuration is the time needed to heat up the heaters in s or nil if unknown
	WarmUpDuration *int64 `json:"warmUpDuration"`
}

// Layer holds information about a layer from a file being printed
type Layer struct {
	// Duration of the layer (in s)
	Duration float64 `json:"duration"`
	// Filament represents the actual amount of filament extruded during
	// this layer in mm
	Filament []float64 `json:"filament"`
	// FractionPrinted represents the fraction of the whole file printed
	// during this layer on a scale between 0 and 1
	FractionPrinted float64 `json:"fractionPrinted"`
	// Height of the layer in mm (0 if unknown)
	Height float64 `json:"height"`
	// Temparatures are the last heater temparatures (in degC or nil if unknown)
	Temperatures []*float64 `json:"temperatures"`
}

// TimesLeft holds information about estimated remaining times
type TimesLeft struct {
	// File progress based estimation in s (nil if unknown)
	File *int64 `json:"file"`
	// Filament consumption based estimation in s (nil if unknown)
	Filament *int64 `json:"filament"`
	// Layer progress based estimation in s (nil if unknown)
	// Deprecated: No longer used, will always return nil
	Layer *int64 `json:"-"`
	// Slicer is time left base on slicer reports (see M73, in s or nil)
	Slicer *int64 `json:"slicer"`
}
