package types

import (
	"time"
)

// ParsedFileInfo holds information about a parsed G-code file
type ParsedFileInfo struct {
	// FileName of the G-code file
	FileName string
	// Size of the file in bytes
	Size uint64
	// LastModified is the last date and time the file was modified or nil if none is set
	LastModified *time.Time // TODO: This will probably need adjustment/custom type
	// Height is the build height of the G-code job or 0 if not found (in mm)
	Height float64
	// FirstLayerHeight is the height of the first layer or 0 if not found (in mm)
	FirstLayerHeight float64
	// LayerHeight is the height of each layer above the first or 0 if not found (in mm)
	LayerHeight float64
	// NumLayers is the number of total layers or nil if unknown
	NumLayers *uint64
	// Filament is the filament consumption per extruder drive (in mm)
	Filament []float64
	// GeneratedBy is the name of the application that generated this file
	GeneratedBy string
	// PrintTime is the estimated job duration (in s)
	PrintTime uint64
	// SimulatedTime is the estimated job duration from G-code simulation (in s)
	SimulatedTime uint64
}
