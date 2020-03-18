package types

// CodeChannel represents available code channels
type CodeChannel byte

const (
	// HTTP is the code channel for HTTP requests
	HTTP CodeChannel = iota
	// Telnet is the code channel for Telnet requests
	Telnet
	// File is the code channel for file jobs
	File
	// USB is the code channel for codes from USB
	USB
	// AUX is the code channel of serial devices except USB (e.g. PanelDue)
	AUX
	// Daemon is the code channel for running triggers or config.g
	Trigger
	// CodeQueue is the code channel for the code queue that executes a couple of
	// codes in-sync with moves
	CodeQueue
	// LCD is the code channel for auxiliary LCD devices (e.g. PanelOne)
	LCD
	// SPI is the default code channel for requests of SPI
	SPI
	// Daemon is the code channel that executes the daemon process
	Daemon
	// AutoPause is the code channel that executes macros on power fail,
	// heater faults and filament out
	AutoPause
	// DefaultChannel is the default channel to use
	DefaultChannel CodeChannel = SPI
)
