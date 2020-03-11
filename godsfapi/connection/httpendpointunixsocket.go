package connection

import (
	"net"

	"os"

	"github.com/Duet3D/DSF-APIs/godsfapi/types"
)

const (
	DefaultBacklog = 4
)

type HttpEndpointHandler interface {
	// Handle the client request
	Handle(h *HttpEndpointUnixSocket, c *HttpEndpointConnection)
}

type HttpEndpointUnixSocket struct {
	// EndpointType of this HTTP endpoint
	EndpointType types.HttpEndpointType
	// Namespace of this HTTO endpoint
	Namespace string
	// EndpointPath of this HTTP endpoint
	EndpointPath string
	// SocketPath to the UNIX socket file
	SocketPath string
	// socket listener
	socket net.Listener
	// Handler to handle individiual requests
	Handler HttpEndpointHandler
}

// NewHttpEndpointUnixSocket opens a new UNIX socket on the given file path
func NewHttpEndpointUnixSocket(t types.HttpEndpointType, ns, path, socketPath string, backlog uint64) (*HttpEndpointUnixSocket, error) {
	h := HttpEndpointUnixSocket{
		EndpointType: t,
		Namespace:    ns,
		EndpointPath: path,
		SocketPath:   socketPath,
	}
	os.Remove(h.SocketPath)

	var err error
	h.socket, err = net.Listen("unix", h.SocketPath)
	if err != nil {
		return nil, err
	}
	go h.accept()

	return &h, nil
}

// Close the socket connection and remove the corresponding socket file
func (h *HttpEndpointUnixSocket) Close() error {
	if h.socket == nil {
		return nil
	}
	err := h.socket.Close()
	h.socket = nil
	return err
}

// accept accepts incoming UNIX socket connections and forwards
// them to a handler
func (h *HttpEndpointUnixSocket) accept() {
	for {
		c, err := h.socket.Accept()
		if err != nil {
			// TODO: instead return?
			continue
		}
		hec := NewHttpEndpointConnection(c, h.EndpointType == types.WebSocket)
		if h.Handler != nil {
			go h.Handler.Handle(h, hec)
		} else {
			hec.SendResponse(500, "No event handler registered", types.StatusCode)
			hec.Close()
		}
	}
}
