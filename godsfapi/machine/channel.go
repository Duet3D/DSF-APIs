package machine

// Compatibility level for emulation
type Compatibility uint64

const (
	// No emulation (same as RepRapFirmware)
	Me Compatibility = iota
	// RepRapFirmware emulation (i.e. no emulation)
	RepRapFirmware
	// Marlin emulation
	Marlin
	// Teacup emulation
	Teacup
	// Sprinter emulation
	Sprinter
	// Repetier emulation
	Repetier
	// NanoDLP emulation (special)
	NanoDLP
)

// Channel holds information about G/M/T-code channels
type Channel struct {
	// Compatibility is the emulation used on this channel
	Compatibility Compatibility `json:"compatibility"`
	// Feedrate is the current feedrate in mm/s
	Feedrate float64 `json:"feedrate"`
	// RelativeExtrusion represents usage of relative extrusion
	RelativeExtrusion bool `json:"relativeExtrusion"`
	// VolumetricExtrusion represents usage of volumetric extrusion
	VolumetricExtrusion bool `json:"volumetricExtrusion"`
	// RelativePositioning represents usage of relative positioning
	RelativePositioning bool `json:"relativePositioning"`
	// UsingInches represents the usage of inches instead of mm
	UsingInches bool `json:"usingInches"`
	// StackDepth is the depth of the stack
	StackDepth uint8 `json:"stackDepth"`
	// LineNumber is the number of the current line
	LineNumber int64 `json:"lineNumber"`
}
