package job

import (
	"time"
)

// Thumbnail holds image parsed out of GCode files
type Thumbnail struct {

	// EncodedImage is the base64 encoded image
	EncodedImage string `json:"encodedImage"`
	// Height of thumbail
	Height int64 `json:"height"`
	// Width of thumbail
	Width int64 `json:"width"`
}

// ParsedFileInfo holds information about a parsed G-code file
type ParsedFileInfo struct {
	// Filament is the filament consumption per extruder drive (in mm)
	Filament []float64 `json:"filament"`
	// FileName of the G-code file
	FileName string `json:"fileName"`
	// FirstLayerHeight is the height of the first layer or 0 if not found (in mm)
	FirstLayerHeight float64 `json:"firstLayerHeight"`
	// GeneratedBy is the name of the application that generated this file
	GeneratedBy string `json:"generatedBy"`
	// Height is the build height of the G-code job or 0 if not found (in mm)
	Height float64 `json:"height"`
	// LastModified is the last date and time the file was modified or nil if none is set
	LastModified *time.Time `json:"lastModified"` // TODO: This will probably need adjustment/custom type
	// LayerHeight is the height of each layer above the first or 0 if not found (in mm)
	LayerHeight float64 `json:"layerHeight"`
	// NumLayers is the number of total layers or 0 if unknown
	NumLayers int64 `json:"numLayers"`
	// PrintTime is the estimated job duration (in s)
	PrintTime *uint64 `json:"printTime"`
	// SimulatedTime is the estimated job duration from G-code simulation (in s)
	SimulatedTime *uint64 `json:"simulatedTime"`
	// Size of the file in bytes
	Size uint64 `json:"size"`
	// Thumbnails is a collection of thumbnails parsed from GCode
	Thumbnails []Thumbnail `json:"thumbnails"`
}
