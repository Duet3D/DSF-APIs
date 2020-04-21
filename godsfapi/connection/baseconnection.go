package connection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/Duet3D/DSF-APIs/godsfapi/commands"
	"github.com/Duet3D/DSF-APIs/godsfapi/connection/initmessages"
)

const (
	// TaskCanceledException is the name of a remote exception to be checked for
	TaskCanceledException = "TaskCanceledException"
	// IncompatibleVersionException is the name of a remote exception to be checked for
	IncompatibleVersionException = "IncompatibleVersionException"
	// SocketDirectory is the default directory in which DSF-related UNIX sockets reside
	SocketDirectory = "/var/run/dsf"
	// SocketFile is the default UNIX socket file for DuetControlServer
	SocketFile = "dcs.sock"
	// FullSocketPath is the default fully-qualified path to the UNIX socket for DuetControlServer
	FullSocketPath = SocketDirectory + "/" + SocketFile
)

// DecodeError is returned if a response from DCS could not be unmarshalled
type DecodeError struct {
	Target string
	Err    error
}

func (e *DecodeError) Unwrap() error { return e.Err }

func (e *DecodeError) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("Failed to unmarshal to type %s because of %v", e.Target, e.Err)
}

// BaseConnection provides common functionalities for more concrete implementations
type BaseConnection struct {
	socket  net.Conn
	decoder *json.Decoder
	id      int64
	Debug   bool
}

// Connect establishes a connecton to the given UNIX socket file
func (bc *BaseConnection) Connect(initMessage initmessages.ClientInitMessage, socketPath string) error {
	var err error
	bc.socket, err = net.Dial("unix", socketPath)
	if err != nil {
		return err
	}
	bc.decoder = json.NewDecoder(bc.socket)

	sim, err := bc.receiveServerInitMessage()
	if err != nil {
		return err
	}

	if !sim.IsCompatible() {
		return errors.New(fmt.Sprintf("Incompatible API version (expected %d got %d)", initmessages.ExpectedServerVersion, sim.Version))
	}

	bc.id = sim.Id

	err = bc.Send(initMessage)
	if err != nil {
		return err
	}

	br, err := bc.ReceiveResponse()
	if err != nil {
		return err
	}
	if !br.IsSuccess() {
		if br.GetErrorType() == IncompatibleVersionException {
			return errors.New(br.GetErrorMessage())
		}
		return errors.New(fmt.Sprintf("Could not set connection type %s (%s: %s)", initMessage.GetMode(), br.GetErrorType(), br.GetErrorMessage()))
	}

	return nil
}

// Close the UNIX socket connection
func (bc *BaseConnection) Close() {
	if bc.socket != nil {
		bc.socket.Close()
		bc.socket = nil
	}
}

// PerformCommand performs an arbitrary command
func (bc *BaseConnection) PerformCommand(command commands.Command) (commands.Response, error) {
	err := bc.Send(command)
	if err != nil {
		return nil, err
	}
	br, err := bc.ReceiveResponse()
	if err != nil {
		return nil, err
	}
	if br.IsSuccess() {
		return br, nil
	}

	// The following two returns intentionally return br instead of nil
	// so the user can work with the received data alongside a simple error object

	if br.GetErrorType() == TaskCanceledException {
		return br, errors.New(br.GetErrorMessage())
	}
	return br, errors.New(fmt.Sprintf("InternalServerError: %s, %s, %s", command.GetCommand(), br.GetErrorType(), br.GetErrorMessage()))
}

// ReceiveResponse receives a deserialized response from the server
func (bc *BaseConnection) ReceiveResponse() (commands.Response, error) {
	br := &commands.BaseResponse{}
	err := bc.Receive(br)
	if err != nil {
		return nil, err
	}
	return br, nil
}

// receiveServerInitMessage returns the ServerInitMessage
func (bc *BaseConnection) receiveServerInitMessage() (*initmessages.ServerInitMessage, error) {
	sim := &initmessages.ServerInitMessage{}
	err := bc.Receive(sim)
	if err != nil {
		return nil, err
	}
	return sim, nil
}

// Receive a deserialized object
func (bc *BaseConnection) Receive(responseContainer interface{}) error {
	if bc.Debug {
		var b json.RawMessage
		if err := bc.decoder.Decode(&b); err != nil {
			if err == io.EOF {
				return err
			}
			return &DecodeError{
				Err:    err,
				Target: fmt.Sprintf("%T", responseContainer),
			}
		}
		log.Println("[DEBUG] <Recv>", string(b))
		return json.Unmarshal(b, responseContainer)
	}
	if err := bc.decoder.Decode(responseContainer); err != nil {
		if err == io.EOF {
			return err
		}
		return &DecodeError{
			Err:    err,
			Target: fmt.Sprintf("%T", responseContainer),
		}
	}
	return nil
}

// ReceiveJson returns a server response as a JSON []byteg
func (bc *BaseConnection) ReceiveJson() ([]byte, error) {
	var raw json.RawMessage
	err := bc.Receive(&raw)
	if err != nil {
		return nil, err
	}
	return []byte(raw), nil
}

// ReceiveJSONString returns a server response as a JSON string
func (bc *BaseConnection) ReceiveJSONString() (string, error) {
	b, err := bc.ReceiveJson()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Send arbitrary data
func (bc *BaseConnection) Send(data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if bc.Debug {
		log.Println("[DEBUG] <Send>", string(b))
	}
	_, err = bc.socket.Write(b)
	return err
}
