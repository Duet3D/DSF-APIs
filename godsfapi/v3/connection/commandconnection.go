package connection

import (
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/connection/initmessages"
)

// CommandConnection used to send commands to the control server
type CommandConnection struct {
	BaseCommandConnection
}

// Connect sends a CommandInitMessage to the server
func (cc *CommandConnection) Connect(socketPath string) error {
	return cc.BaseConnection.Connect(initmessages.NewCommandInitMessage(), socketPath)
}
