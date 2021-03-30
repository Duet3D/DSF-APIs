package sensors

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// FilamentMonitor holds information about a filament monitor
type FilamentMonitor map[string]interface{}

func toFilamentMonitor(src interface{}) (FilamentMonitor, error) {
	b, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}
	var fm FilamentMonitor
	err = json.Unmarshal(b, &fm)
	if err != nil {
		return nil, err
	}
	return fm, nil
}

// GetType returns the FilamentMonitorType of this instance
func (f FilamentMonitor) GetType() FilamentMonitorType {
	return FilamentMonitorType(f["type"].(string))
}

// FilamentMonitors is a slice of FilamentMonitor
type FilamentMonitors []FilamentMonitor

// ErrInvalidIndex is returned in case an invalid index is accessed
var ErrInvalidIndex = errors.New("Invalid index")

// GetAsBaseFilamentMonitor returns the instance at the given index as BaseFilamentMonitor
func (fm FilamentMonitors) GetAsBaseFilamentMonitor(i int) (*BaseFilamentMonitor, error) {
	if i < 0 || i > len(fm) {
		return nil, ErrInvalidIndex
	}
	f := fm[i]
	bfm := &BaseFilamentMonitor{}
	err := mapstructure.Decode(f, bfm)
	if err != nil {
		return nil, err
	}
	return bfm, nil
}

// GetAsSimpleFilamentMonitor returns the instance at the given index as SimpleFilamentMonitor
func (fm FilamentMonitors) GetAsSimpleFilamentMonitor(i int) (*SimpleFilamentMonitor, error) {
	if i < 0 || i > len(fm) {
		return nil, ErrInvalidIndex
	}
	f := fm[i]
	name := f.GetType()
	if name != Simple {
		return nil, fmt.Errorf("Not SimpleFilamentMonitor: %s", name)
	}
	sfm := &SimpleFilamentMonitor{}
	err := mapstructure.Decode(f, sfm)
	if err != nil {
		return nil, err
	}
	return sfm, nil
}

// GetAsLaserFilamentMonitor returns the instance at the given index as LaserFilamentMonitor
func (fm FilamentMonitors) GetAsLaserFilamentMonitor(i int) (*LaserFilamentMonitor, error) {
	if i < 0 || i > len(fm) {
		return nil, ErrInvalidIndex
	}
	f := fm[i]
	name := f.GetType()
	if name != Laser {
		return nil, fmt.Errorf("Not LaserFilamentMonitor: %s", name)
	}
	lfm := &LaserFilamentMonitor{}
	err := mapstructure.Decode(f, lfm)
	if err != nil {
		return nil, err
	}
	return lfm, nil
}

// GetAsPulsedFilamentMonitor returns the instance at the given index as PulsedFilamentMonitor
func (fm FilamentMonitors) GetAsPulsedFilamentMonitor(i int) (*PulsedFilamentMonitor, error) {
	if i < 0 || i > len(fm) {
		return nil, ErrInvalidIndex
	}
	f := fm[i]
	name := f.GetType()
	if name != Pulsed {
		return nil, fmt.Errorf("Not PulsedFilamentMonitor: %s", name)
	}
	pfm := &PulsedFilamentMonitor{}
	err := mapstructure.Decode(f, pfm)
	if err != nil {
		return nil, err
	}
	return pfm, nil
}

// GetAsRotatingMagnetFilamentMonitor returns the instance at the given index as RotatingMagnetFilamentMonitor
func (fm FilamentMonitors) GetAsRotatingMagnetFilamentMonitor(i int) (*RotatingMagnetFilamentMonitor, error) {
	if i < 0 || i > len(fm) {
		return nil, ErrInvalidIndex
	}
	f := fm[i]
	name := f.GetType()
	if name != RotatingMagnet {
		return nil, fmt.Errorf("Not RotatingMagnetFilamentMonitor: %s", name)
	}
	rmfm := &RotatingMagnetFilamentMonitor{}
	err := mapstructure.Decode(f, rmfm)
	if err != nil {
		return nil, err
	}
	return rmfm, nil
}

// FilamentMonitorStatus are the possible filament sensor statuses
type FilamentMonitorStatus string

const (
	// NoMonitor is present
	NoMonitor FilamentMonitorStatus = "noMonitor"
	// Ok for filament monitor working normally
	Ok = "ok"
	// NoDataReceived if no data received from remote filament monitor
	NoDataReceived = "noDataReceived"
	// NoFilament if no filament is detected
	NoFilament = "noFilament"
	// TooLittleMovement if the sensor reads less movement than expected
	TooLittleMovement = "tooLittleMovement"
	// TooMuchMovement if the sensor reads more movement than expected
	TooMuchMovement = "tooMuchMovement"
	// SensorError if the sensor encountered an error
	SensorError = "sensorError"
)

// BaseFilamentMonitor holds information about a filament monitor
type BaseFilamentMonitor struct {
	// Enabled indicates if this filament monitor is enabled
	Enabled bool `json:"enabled"`
	// Status is the last reported status of this filament monitor
	Status FilamentMonitorStatus `json:"status"`
	// Type of this filament monitor
	Type FilamentMonitorType `json:"type"`
}

// AsFilamentMonitor returns this instance as FilamentMonitor
func (bfm *BaseFilamentMonitor) AsFilamentMonitor() (FilamentMonitor, error) {
	return toFilamentMonitor(bfm)
}

// SimpleFilamentMonitor represents a simple filament monitor
type SimpleFilamentMonitor struct {
	BaseFilamentMonitor `mapstructure:",squash"`

	// FilamentPresent indicates if filament is present or nil if not available
	FilamentPresent *bool `json:"filamentPresent"`
}

// AsFilamentMonitor returns this instance as FilamentMonitor
func (sfm *SimpleFilamentMonitor) AsFilamentMonitor() (FilamentMonitor, error) {
	return toFilamentMonitor(sfm)
}

// FilamentMonitorProperties shared by all FilamentMonitorCalibrated and FilamentMonitorConfigured structs
type FilamentMonitorProperties struct {
	// PercentMax is the maximum allowed movement percentage (0..1 or greater)
	PercentMax float64 `json:"percentMax"`
	// PercentMin is the minimum allowed movement percentage (0..1 or greater)
	PercentMin float64 `json:"percentMin"`
}

// FilamentMonitorCalibrated shared by all concrete type structs
type FilamentMonitorCalibrated struct {
	FilamentMonitorProperties `mapstructure:",squash"`

	// TotalDistance extruded (in mm)
	TotalDistance float64 `json:"totalDistance"`
}

// FilamentMonitorConfigured shared by all concrete type structs
type FilamentMonitorConfigured struct {
	FilamentMonitorProperties `mapstructure:",squash"`

	// SampleDistance in mm
	SampleDistance float64 `json:"sampleDistance"`
}

// FilamentMonitorType represents supported filament monitors
type FilamentMonitorType string

const (
	// Simple filament monitor
	Simple FilamentMonitorType = "simple"
	// Laser for a laser-based monitor
	Laser = "laser"
	// Pulsed filament monitor
	Pulsed = "pulsed"
	// RotatingMagnet filament monitor
	RotatingMagnet = "rotatingMagnet"
	// Unkown for unknown sensor type
	Unkown = "unknown"
)

// LaserFilamentMonitor holds information about a laser filament monitor
type LaserFilamentMonitor struct {
	SimpleFilamentMonitor `mapstructure:",squash"`

	// Calibrated holds calibrated properties of this filament sensor
	Calibrated LaserFilamentMonitorCalibrated `json:"calibrated"`
	// Configured holds configured properties of this filament sensor
	Configured LaserFilamentMonitorConfigured `json:"configured"`
}

// AsFilamentMonitor returns this instance as FilamentMonitor
func (lfm *LaserFilamentMonitor) AsFilamentMonitor() (FilamentMonitor, error) {
	return toFilamentMonitor(lfm)
}

// LaserFilamentMonitorCalibrated reprent the calibrated properties
// of a laser filament monitor
type LaserFilamentMonitorCalibrated struct {
	FilamentMonitorCalibrated `mapstructure:",squash"`
	// Sensitivity from calibration
	Sensitivity float64 `json:"sensitivity"`
}

// LaserFilamentMonitorConfigured represents configured properties
// of a laser filament sensor
type LaserFilamentMonitorConfigured FilamentMonitorConfigured

// PulsedFilamentMonitor holds information about a pulsed filament monitor
type PulsedFilamentMonitor struct {
	BaseFilamentMonitor `mapstructure:",squash"`

	// Calibrated holds calibrated properties of this filament monitor
	Calibrated PulsedFilamentMonitorCalibrated `json:"calibrated"`
	// Configured holds configured properties of this filament monitor
	Configured PulsedFilamentMonitorConfigured `json:"configured"`
}

// AsFilamentMonitor returns this instance as FilamentMonitor
func (pfm *PulsedFilamentMonitor) AsFilamentMonitor() (FilamentMonitor, error) {
	return toFilamentMonitor(pfm)
}

// PulsedFilamentMonitorCalibrated represents calibrated properties of pulsed filament monitor
type PulsedFilamentMonitorCalibrated struct {
	FilamentMonitorCalibrated `mapstructure:",squash"`
	// MmPerPulse is extruded distance per pulse (in mm)
	MmPerPulse float64 `json:"mmPerPulse"`
}

// PulsedFilamentMonitorConfigured represents configured properties
// of a pulsed filament sensor
type PulsedFilamentMonitorConfigured struct {
	FilamentMonitorConfigured `mapstructure:",squash"`
	// MmPerPulse is extruded distance per pulse (in mm)
	MmPerPulse float64 `json:"mmPerPulse"`
}

// RotatingMagnetFilamentMonitor holds information about a rotating magnet filament monitor
type RotatingMagnetFilamentMonitor struct {
	SimpleFilamentMonitor `mapstructure:",squash"`

	// Calibrated holds calibrated properties of this filament monitor
	Calibrated RotatingMagnetFilamentMonitorCalibrated `json:"calibrated"`
	// Configured holds configured properties of this filament monitor
	Configured RotatingMagnetFilamentMonitorConfigured `json:"configured"`
}

// AsFilamentMonitor returns this instance as FilamentMonitor
func (rmfm *RotatingMagnetFilamentMonitor) AsFilamentMonitor() (FilamentMonitor, error) {
	return toFilamentMonitor(rmfm)
}

// RotatingMagnetFilamentMonitorCalibrated represents calibrated properties of pulsed filament monitor
type RotatingMagnetFilamentMonitorCalibrated struct {
	FilamentMonitorCalibrated `mapstructure:",squash"`
	// MmPerRev is extruded distance per revolution (in mm)
	MmPerRev float64 `json:"mmPerRev"`
}

// RotatingMagnetFilamentMonitorConfigured represents configured properties
// of a pulsed filament sensor
type RotatingMagnetFilamentMonitorConfigured struct {
	FilamentMonitorConfigured `mapstructure:",squash"`
	// MmPerRev is extruded distance per revolution (in mm)
	MmPerRev float64 `json:"mmPerRev"`
}
