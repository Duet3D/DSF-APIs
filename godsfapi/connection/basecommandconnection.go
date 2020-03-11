package connection

import (
	"encoding/json"
	"os"

	"github.com/Duet3D/DSF-APIs/godsfapi/commands"
	"github.com/Duet3D/DSF-APIs/godsfapi/machine"
	"github.com/Duet3D/DSF-APIs/godsfapi/types"
)

// BaseCommandConnection for sending commands to the control server
type BaseCommandConnection struct {
	BaseConnection
}

// AddHttpEndpoint adds a new third-party HTTP endpoint in the format /machine/{ns}/{path}
func (bcc *BaseCommandConnection) AddHttpEndpoint(t types.HttpEndpointType, ns, path string, backlog uint64) (*HttpEndpointUnixSocket, error) {
	r, err := bcc.PerformCommand(commands.NewAddHttpEndpoint(t, ns, path))
	if err != nil {
		return nil, err
	}
	socketPath := r.GetResult().(string)
	if backlog <= 0 {
		backlog = DefaultBacklog
	}
	heus, err := NewHttpEndpointUnixSocket(t, ns, path, socketPath, backlog)
	if err != nil {
		return nil, err
	}
	return heus, nil
}

// AddUserSession adds a new user session. Pass -1 as originPort to have it replaced by current PID.
func (bcc *BaseCommandConnection) AddUserSession(access types.AccessLevel, t types.SessionType, origin string, originPort int) (int, error) {
	if originPort == -1 {
		originPort = os.Getpid()
	}
	r, err := bcc.PerformCommand(commands.NewAddUserSession(access, t, origin, originPort))
	if err != nil {
		return -1, err
	}
	return r.GetResult().(int), nil
}

// RemoveHttpEndpoint removes an existing HTTP endpoint
func (bcc *BaseCommandConnection) RemoveHttpEndpoint(t types.HttpEndpointType, ns, path string) (bool, error) {
	r, err := bcc.PerformCommand(commands.NewRemoveHttpEndpoint(t, ns, path))
	if err != nil {
		return false, err
	}
	return r.IsSuccess(), nil
}

// RemoveUserSession removes an existing user session
func (bcc *BaseCommandConnection) RemoveUserSession(id int) (bool, error) {
	r, err := bcc.PerformCommand(commands.NewRemoveUserSession(id))
	if err != nil {
		return false, err
	}
	return r.IsSuccess(), nil
}

// Flush waits for all pending codes of the given channel to finish
func (bcc *BaseCommandConnection) Flush(channel types.CodeChannel) (bool, error) {
	r, err := bcc.PerformCommand(commands.NewFlush(channel))
	if err != nil {
		return false, err
	}
	return r.IsSuccess(), nil
}

// GetFileInfo gets the parsed G-code file information
func (bcc *BaseCommandConnection) GetFileInfo(fileName string) (*types.ParsedFileInfo, error) {
	r, err := bcc.PerformCommand(commands.NewGetFileInfo(fileName))
	if err != nil {
		return nil, err
	}
	pfi := r.GetResult().(types.ParsedFileInfo)
	return &pfi, nil
}

// PerformCode executes an arbitrary pre-parsed code
// Note that even with an error being nil the returned *commands.CodeResult
// can also be nil, e.g. when sending Asynchronous commands that will only be queued and have no result yet.
func (bcc *BaseCommandConnection) PerformCode(code *commands.Code) (*commands.CodeResult, error) {
	r, err := bcc.PerformCommand(code)
	if err != nil {
		return nil, err
	}
	if r.GetResult() != nil {
		cr := r.GetResult().(commands.CodeResult)
		return &cr, nil
	}
	return nil, nil
}

// PerformSimpleCode executes an arbitrary G/M/T-code in text form and returns the result as a string
func (bcc *BaseCommandConnection) PerformSimpleCode(code string, channel types.CodeChannel) (string, error) {
	r, err := bcc.PerformCommand(commands.NewSimpleCode(code, channel))
	if err != nil {
		return "", err
	}
	return r.GetResult().(string), nil
}

// GetMachineModel retrieves the full object model of the machine.
// In subscription mode this is the first command that has to be called once a connection has
// been established
func (bcc *BaseCommandConnection) GetMachineModel() (*machine.MachineModel, error) {
	r, err := bcc.PerformCommand(commands.NewGetMachineModel())
	if err != nil {
		return nil, err
	}
	mm := r.GetResult().(machine.MachineModel)
	return &mm, nil
}

// GetSerializedMachineModel fetches the machine model as UTF-8 JSON
func (bcc *BaseCommandConnection) GetSerializedMachineModel() (json.RawMessage, error) {
	var raw json.RawMessage
	err := bcc.Send(commands.NewGetMachineModel())
	if err != nil {
		return raw, err
	}
	err = bcc.Receive(&raw)
	if err != nil {
		return raw, err
	}
	return raw, nil
}

// ResolvePath resolves a RepRapFirmware-style file path to a real file path
func (bcc *BaseCommandConnection) ResolvePath(path string) (string, error) {
	r, err := bcc.PerformCommand(commands.NewResolvePath(path))
	if err != nil {
		return "", err
	}
	return r.GetResult().(string), nil
}

// LockMachineModel locks the machine model for read/write Access
// It is MANDATORY to call UnlockMachineModel when write access has finished
func (bcc *BaseCommandConnection) LockMachineModel() error {
	_, err := bcc.PerformCommand(commands.NewLockMachineModel())
	return err
}

// SetMachineModel sets a given property to a certain value. Make sure to lock the object
// model before calling this.
func (bcc *BaseCommandConnection) SetMachineModel(path, value string) (bool, error) {
	r, err := bcc.PerformCommand(commands.NewSetMachineModel(path, value))
	if err != nil {
		return false, err
	}
	return r.IsSuccess(), nil
}

// SyncMachineModel waits for the full machine model to be updated from RepRapFirmware
func (bcc *BaseCommandConnection) SyncMachineModel() error {
	_, err := bcc.PerformCommand(commands.NewSyncMachineModel())
	return err
}

// UnlockMachineModel unlocks the machine model
func (bcc *BaseCommandConnection) UnlockMachineModel() error {
	_, err := bcc.PerformCommand(commands.NewUnlockMachineModel())
	return err
}
