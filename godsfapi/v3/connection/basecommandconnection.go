package connection

import (
	"encoding/json"
	"os"

	"github.com/Duet3D/DSF-APIs/godsfapi/v3/commands"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/httpendpoints"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/job"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/messages"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/state"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/usersessions"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/types"
)

// BaseCommandConnection for sending commands to the control server
type BaseCommandConnection struct {
	BaseConnection
}

// AddHttpEndpoint adds a new third-party HTTP endpoint in the format /machine/{ns}/{path}
func (bcc *BaseCommandConnection) AddHttpEndpoint(t httpendpoints.HttpEndpointType, ns, path string, isUploadRequest bool, backlog uint64) (*HttpEndpointUnixSocket, error) {
	r, err := bcc.PerformCommand(commands.NewAddHttpEndpoint(t, ns, path, isUploadRequest))
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
func (bcc *BaseCommandConnection) AddUserSession(access usersessions.AccessLevel, t usersessions.SessionType, origin string, originPort int) (int, error) {
	if originPort == -1 {
		originPort = os.Getpid()
	}
	r, err := bcc.PerformCommand(commands.NewAddUserSession(access, t, origin, originPort))
	if err != nil {
		return -1, err
	}
	return r.GetResult().(int), nil
}

// CheckPassword checks the given password (see M551)
func (bcc *BaseCommandConnection) CheckPassword(password string) (bool, error) {
	r, err := bcc.PerformCommand(commands.NewCheckPassword(password))
	if err != nil {
		return false, err
	}
	return r.IsSuccess(), nil
}

// RemoveHttpEndpoint removes an existing HTTP endpoint
func (bcc *BaseCommandConnection) RemoveHttpEndpoint(t httpendpoints.HttpEndpointType, ns, path string) (bool, error) {
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
func (bcc *BaseCommandConnection) GetFileInfo(fileName string) (*job.ParsedFileInfo, error) {
	r, err := bcc.PerformCommand(commands.NewGetFileInfo(fileName))
	if err != nil {
		return nil, err
	}
	pfi := r.GetResult().(job.ParsedFileInfo)
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
// Deprecated: Use GetObjectModel instead
func (bcc *BaseCommandConnection) GetMachineModel() (*machine.MachineModel, error) {
	return bcc.GetObjectModel()
}

// GetObjectModel retrieves the full object model of the machine.
// In subscription mode this is the first command that has to be called once a connection has
// been established
func (bcc *BaseCommandConnection) GetObjectModel() (*machine.MachineModel, error) {
	r, err := bcc.PerformCommand(commands.NewGetObjectModel())
	if err != nil {
		return nil, err
	}
	mm := r.GetResult().(machine.MachineModel)
	return &mm, nil
}

// GetSerializedMachineModel fetches the machine model as UTF-8 JSON
// Deprecated: Use GetSerializedObjectModel instead
func (bcc *BaseCommandConnection) GetSerializedMachineModel() (json.RawMessage, error) {
	return bcc.GetSerializedObjectModel()
}

// GetSerializedObjectModel fetches the object model as UTF-8 JSON
func (bcc *BaseCommandConnection) GetSerializedObjectModel() (json.RawMessage, error) {
	var raw json.RawMessage
	err := bcc.Send(commands.NewGetObjectModel())
	if err != nil {
		return raw, err
	}
	err = bcc.Receive(&raw)
	if err != nil {
		return raw, err
	}
	return raw, nil
}

// LockMachineModel locks the machine model for read/write Access
// It is MANDATORY to call UnlockObjectModel when write access has finished
// Deprecated: Use LockObjectModel instead
func (bcc *BaseCommandConnection) LockMachineModel() error {
	return bcc.LockObjectModel()
}

// LockObjectModel locks the machine model for read/write Access
// It is MANDATORY to call UnlockObjectModel when write access has finished
func (bcc *BaseCommandConnection) LockObjectModel() error {
	_, err := bcc.PerformCommand(commands.NewLockObjectModel())
	return err
}

// PatchObjectModel will apply a full patch to the object model. Use with care!
func (bcc *BaseCommandConnection) PatchObjectModel(key, value string) error {
	_, err := bcc.PerformCommand(commands.NewPatchObjectModel(key, value))
	return err
}

// SetMachineModel sets a given property to a certain value. Make sure to lock the object
// model before calling this.
// Deprecated: Use SetObjectModel instead
func (bcc *BaseCommandConnection) SetMachineModel(path, value string) (bool, error) {
	return bcc.SetObjectModel(path, value)
}

// SetObjectModel sets a given property to a certain value. Make sure to lock the object
// model before calling this.
func (bcc *BaseCommandConnection) SetObjectModel(path, value string) (bool, error) {
	r, err := bcc.PerformCommand(commands.NewSetObjectModel(path, value))
	if err != nil {
		return false, err
	}
	return r.IsSuccess(), nil
}

// SyncMachineModel waits for the full machine model to be updated from RepRapFirmware
// Deprecated: Use SyncObjectModel instead
func (bcc *BaseCommandConnection) SyncMachineModel() error {
	return bcc.SyncObjectModel()
}

// SyncObjectModel waits for the full object model to be updated from RepRapFirmware
func (bcc *BaseCommandConnection) SyncObjectModel() error {
	_, err := bcc.PerformCommand(commands.NewSyncObjectModel())
	return err
}

// UnlockMachineModel unlocks the machine model
// Deprecated: Use UnlockObjectModel instead
func (bcc *BaseCommandConnection) UnlockMachineModel() error {
	return bcc.UnlockObjectModel()
}

// UnlockObjectModel unlocks the object model
func (bcc *BaseCommandConnection) UnlockObjectModel() error {
	_, err := bcc.PerformCommand(commands.NewUnlockObjectModel())
	return err
}

// ResolvePath resolves a RepRapFirmware-style file path to a real file path
func (bcc *BaseCommandConnection) ResolvePath(path string) (string, error) {
	r, err := bcc.PerformCommand(commands.NewResolvePath(path))
	if err != nil {
		return "", err
	}
	return r.GetResult().(string), nil
}

// InstallPlugin to install or upgrade a plugin.
// pluginFile is the absolute file path to the plugin ZIP bundle
func (bcc *BaseCommandConnection) InstallPlugin(pluginFile string) error {
	_, err := bcc.PerformCommand(commands.NewInstallPlugin(pluginFile))
	return err
}

// SetPluginData sets custom plugin data in the object model
// plugin is the name of the plugin and is optional. Leave empty if not needed
func (bcc *BaseCommandConnection) SetPluginData(plugin, key, value string) error {
	_, err := bcc.PerformCommand(commands.NewSetPluginData(plugin, key, value))
	return err
}

// StartPlugin starts a plugin
func (bcc *BaseCommandConnection) StartPlugin(plugin string) error {
	_, err := bcc.PerformCommand(commands.NewStartPlugin(plugin))
	return err
}

// StopPlugin stops a plugin
func (bcc *BaseCommandConnection) StopPlugin(plugin string) error {
	_, err := bcc.PerformCommand(commands.NewStopPlugin(plugin))
	return err
}

// UninstallPlugin uninstalls a plugin
func (bcc *BaseCommandConnection) UninstallPlugin(plugin string) error {
	_, err := bcc.PerformCommand(commands.NewUninstallPlugin(plugin))
	return err
}

// Write an arbitrary generic message
func (bcc *BaseCommandConnection) WriteTextMessage(mType messages.MessageType, message string, outputMessage bool, logLevel *state.LogLevel) error {
	_, err := bcc.PerformCommand(commands.NewWriteMessage(mType, message, outputMessage, logLevel))
	return err
}

// Write an arbitrary generic message from an existing messages.Message instance
func (bcc *BaseCommandConnection) WriteMessage(message messages.Message, outputMessage bool, logLevel *state.LogLevel) error {
	_, err := bcc.PerformCommand(commands.NewWriteMessage(message.Type, message.Content, outputMessage, logLevel))
	return err
}

// SetUpdateStatus overrides the current machin status if a software update is in progress.
// The object model may not be locked when this is called.
func (bcc *BaseCommandConnection) SetUpdateStatus(updating bool) error {
	_, err := bcc.PerformCommand(commands.NewSetUpdateStatus(updating))
	return err
}
