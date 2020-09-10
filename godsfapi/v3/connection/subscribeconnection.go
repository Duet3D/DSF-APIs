package connection

import (
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/commands"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/connection/initmessages"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine"
)

// SubscribeConnection is used to subscribe for object model updates
type SubscribeConnection struct {
	BaseConnection
	Mode initmessages.SubscriptionMode
	// Filter is the deprecated delimited filter expression
	Filter string
	// Filters is a list of filter expressions
	Filters []string
}

// Connect will send a SubscribeInitMessage to the control server
func (sc *SubscribeConnection) Connect(mode initmessages.SubscriptionMode, filters []string, socketPath string) error {
	sc.Mode = mode
	sc.Filters = filters
	sim := initmessages.NewSubscribeInitMessage(mode, filters)
	return sc.BaseConnection.Connect(sim, socketPath)
}

// GetMachineModel retrieves the full object model of the machine.
// In subscription mode this is the first command that has to be called once a connection has
// been established
func (sc *SubscribeConnection) GetMachineModel() (*machine.MachineModel, error) {
	m := machine.NewMachineModel()
	err := sc.Receive(m)
	if err != nil {
		return nil, err
	}
	err = sc.Send(commands.NewAcknowledge())
	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetMachineModelPatch receives a (partial) machine model update as JSON UTF-8 string.
// If the subscription mode is set to Patch, new update patches of the object model
// need to be applied manually. This method is intended to receive such fragments.
func (sc *SubscribeConnection) GetMachineModelPatch() (string, error) {
	j, err := sc.ReceiveJSONString()
	if err != nil {
		return "", err
	}
	err = sc.Send(commands.NewAcknowledge())
	if err != nil {
		return "", err
	}
	return j, nil
}
