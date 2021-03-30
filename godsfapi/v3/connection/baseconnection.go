package connection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Duet3D/DSF-APIs/godsfapi/v3/commands"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/connection/initmessages"
)

const (
	// TaskCanceledException is the name of a remote exception to be checked for
	TaskCanceledException = "TaskCanceledException"
	// IncompatibleVersionException is the name of a remote exception to be checked for
	IncompatibleVersionException = "IncompatibleVersionException"
	// SocketDirectory is the default directory in which DSF-related UNIX sockets reside
	SocketDirectory = "/run/dsf"
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

var o sync.Once
var conns []Closer

// CloseOnSignals will call Close on a connection if SIGINT or SIGTERM is encountered
func CloseOnSignals(c Closer) {
	o.Do(func() {
		conns = make([]Closer, 0)
		sc := make(chan os.Signal)
		signal.Notify(sc, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		go func() {
			<-sc
			for _, c := range conns {
				c.Close()
			}
			os.Exit(1)
		}()
	})
	conns = append(conns, c)
}

// Closer is the interface implemented by all connections
type Closer interface {
	// Close the UNIX socket connection
	Close() error
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
		return fmt.Errorf("Incompatible API version (expected %d got %d)", initmessages.ProtocolVersion, sim.Version)
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
		return fmt.Errorf("Could not set connection type %s (%s: %s)", initMessage.GetMode(), br.GetErrorType(), br.GetErrorMessage())
	}
	if bc.Debug {
		log.Println("[DEBUG] <Connect> Connection established")
	}
	return nil
}

// Close the UNIX socket connection
func (bc *BaseConnection) Close() error {
	if bc == nil {
		return nil
	}
	if bc.socket != nil {
		if bc.Debug {
			log.Println("[DEBUG] <Close> Closing connection")
		}
		err := bc.socket.Close()
		if err != nil {
			log.Println("[ERROR] <Close> Error closing connection", err)
		}
		bc.socket = nil
		return err
	}
	return nil
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
	return br, fmt.Errorf("InternalServerError: %s, %s, %s", command.GetCommand(), br.GetErrorType(), br.GetErrorMessage())
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
