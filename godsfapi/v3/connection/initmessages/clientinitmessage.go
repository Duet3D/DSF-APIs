package initmessages

// ConnectionMode represents supported connection types for client connections
type ConnectionMode string

const (
	// ConnectionModeUnknown is an unknown connection type. If this is used the connection
	// is immediately terminated
	ConnectionModeUnknown ConnectionMode = "Unknown"
	// ConnectionModeCommand enters command mode. This allows clients to send general
	// purpose messages to the control server like G-codes or requests of the full
	// object model
	ConnectionModeCommand = "Command"
	// ConnectionModeIntercept enters interception mode. This allows clients to intercept
	// G/M/T-codes before or after they are initially processed or after they have been executed
	ConnectionModeIntercept = "Intercept"
	// ConnectionModeSubscribe enters subscription mode. In this mode object model updates are
	// transmitted to the client after each update
	ConnectionModeSubscribe = "Subscribe"
)

// ClientInitMessage is sent from the client to the server as response
// to a ServerInitMessage. It allows to select the connection mode.
type ClientInitMessage interface {
	// GetMode returns the connection mode
	GetMode() ConnectionMode
}

// BaseInitMessage holds the common members of all init messages
type BaseInitMessage struct {
	// Mode is the desired connection mode
	Mode ConnectionMode
	// Version number of the client-side API
	Version int64
}

// NewBaseInitMessage creates a new BaseInitMessage for the given ConnectionMode
func NewBaseInitMessage(mode ConnectionMode) BaseInitMessage {
	return BaseInitMessage{
		Mode:    mode,
		Version: ProtocolVersion,
	}
}

// GetMode returns the connection mode
func (bim *BaseInitMessage) GetMode() ConnectionMode {
	return bim.Mode
}

// commandInitMessage is a BaseInitMessage with a fixed mode and no further members
var commandInitMessage = NewBaseInitMessage(ConnectionModeCommand)

// NewCommandInitMessage returns a command init message
func NewCommandInitMessage() ClientInitMessage {
	return &commandInitMessage
}
