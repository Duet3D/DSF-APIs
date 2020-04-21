package machine

import (
	"github.com/Duet3D/DSF-APIs/godsfapi/types"
)

// Channels holds all available code channels
type Channels struct {
	// HTTP is the G/M/T-code channel for HTTP requests
	HTTP Channel
	// Telnet is the G/M/T-code channel for Telnet requests
	Telnet Channel
	// File is the G/M/T-code channel for file jobs
	File Channel
	// USB is the G/M/T-code channel for USB
	USB Channel
	// AUX is the G/M/T-code channel for serial devices (e.g. UART, PanelDue)
	AUX Channel
	// Trigger is the G/M/T-code channel to deal with triggers and config.g
	Trigger Channel
	// CodeQueue is the G/M/T-code channel for the code queue
	CodeQueue Channel
	// LCD is the G/M/T-code channel for auxiliary LCD devices
	LCD Channel
	// SPI is the G/M/T-code channel for generic codes via SPI
	SPI Channel
	// Daemon is the code channel that executes the daemon process
	Daemon Channel
	// AutoPause is the G/M/T-code channel for auto pause events
	AutoPause Channel
}

// NewChannels creates a new Channels with default Compatibility set for certain channels
func NewChannels() Channels {
	return Channels{
		Telnet: Channel{Compatibility: Marlin},
		USB:    Channel{Compatibility: Marlin},
	}
}

// Get will return the Channel to the given types.CodeChannel.
// It will return SPI for unknown types.
func (ch *Channels) Get(cc types.CodeChannel) Channel {
	switch cc {
	case types.HTTP:
		return ch.HTTP
	case types.Telnet:
		return ch.Telnet
	case types.File:
		return ch.File
	case types.USB:
		return ch.USB
	case types.AUX:
		return ch.AUX
	case types.Trigger:
		return ch.Trigger
	case types.CodeQueue:
		return ch.CodeQueue
	case types.LCD:
		return ch.LCD
	case types.SPI:
		return ch.SPI
	case types.Daemon:
		return ch.Daemon
	case types.AutoPause:
		return ch.AutoPause
	default:
		return ch.SPI
	}
}
