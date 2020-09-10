package types

// CodeChannel represents supported input code channels
type CodeChannel string

const (
	// HTTP is the code channel for HTTP requests
	HTTP CodeChannel = "HTTP"
	// Telnet is the code channel for Telnet requests
	Telnet = "Telnet"
	// File is the code channel for file jobs
	File = "File"
	// USB is the code channel for codes from USB
	USB = "USB"
	// Aux is the code channel of serial devices except USB (e.g. PanelDue)
	Aux = "Aux"
	// Trigger is the code channel running triggers or config.g
	Trigger = "Trigger"
	// Queue is the code channel for the code queue that executes a couple of
	// codes in-sync with moves
	Queue = "Queue"
	// LCD is the code channel for auxiliary LCD devices (e.g. PanelOne)
	LCD = "LCD"
	// SBC is the default code channel for requests of SBC
	SBC = "SBC"
	// Daemon is the code channel for running triggers or config.g
	Daemon = "Daemon"
	// Aux2 is the code channel for the second UART port
	Aux2 = "Aux2"
	// AutoPause is the code channel that executes macros on power fail,
	// heater faults and filament out
	AutoPause = "AutoPause"
	// Unknown code channel
	Unknown = "Unknown"

	// DefaultChannel is the default channel to use
	DefaultChannel CodeChannel = SBC
)

// AllChannels returns a slice containing all channels
func AllChannels() []CodeChannel {
	return []CodeChannel{HTTP, Telnet, File, USB, Aux, Trigger, Queue, LCD, SBC, Daemon, Aux2, AutoPause, Unknown}
}
