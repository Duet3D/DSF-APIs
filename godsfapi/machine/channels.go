package machine

import "github.com/Duet3D/DSF-APIs/godsfapi/types"

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
	// Aux is the G/M/T-code channel for serial devices (e.g. UART, PanelDue)
	Aux Channel
	// Daemon is the G/M/T-code channel to deal with triggers and config.g
	Daemon Channel
	// Queue is the G/M/T-code channel for the code queue
	Queue Channel
	// LCD is the G/M/T-code channel for auxiliary LCD devices
	LCD Channel
	// SBC is the G/M/T-code channel for generic codes via SBC
	SBC Channel
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
	case types.Aux:
		return ch.Aux
	case types.Daemon:
		return ch.Daemon
	case types.Queue:
		return ch.Queue
	case types.LCD:
		return ch.LCD
	case types.SBC:
		return ch.SBC
	case types.AutoPause:
		return ch.AutoPause
	default:
		return ch.SBC
	}
}
