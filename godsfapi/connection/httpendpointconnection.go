package connection

import (
	"encoding/json"
	"net"

	"github.com/Duet3D/DSF-APIs/godsfapi/commands"
	"github.com/Duet3D/DSF-APIs/godsfapi/types"
)

// HttpEndpointConnection is dealing with requests received from a custom HTTP endpoint
type HttpEndpointConnection struct {
	conn        net.Conn
	isWebSocket bool
	decoder     *json.Decoder
}

// NewHttpEndpointConnection creates a new instance of HttpEndpointConnection
func NewHttpEndpointConnection(c net.Conn, isWebSocket bool) *HttpEndpointConnection {
	return &HttpEndpointConnection{
		conn:        c,
		isWebSocket: isWebSocket,
		decoder:     json.NewDecoder(c),
	}
}

// Close closes the underlying connection
func (h *HttpEndpointConnection) Close() error {
	if h.conn == nil {
		return nil
	}
	err := h.conn.Close()
	h.conn = nil
	return err
}

// ReadRequest reads information about the last HTTP request. A call to this method may fail
func (h *HttpEndpointConnection) ReadRequest() (*commands.ReceivedHttpRequest, error) {
	rhr := commands.NewReceivedHttpRequest()
	err := h.Receive(rhr)
	if err != nil {
		return nil, err
	}
	return rhr, nil
}

// SendResponse sends a simple HTTP response to the client and closes this connection unless
// it is a WebSocket
func (h *HttpEndpointConnection) SendResponse(statusCode uint16, response string, t types.HttpResponseType) error {

	// Close this connection automatically if only one response can be sent
	if !h.isWebSocket {
		defer h.Close()
	}
	shr := commands.NewSendHttpResponse(statusCode, response, t)
	err := h.Send(shr)
	return err
}

// Receive a deserialized object
func (h *HttpEndpointConnection) Receive(responseContainer interface{}) error {
	if err := h.decoder.Decode(responseContainer); err != nil {
		return err
	}
	return nil
}

// ReceiveJson returns a server response as a JSON string
func (h *HttpEndpointConnection) ReceiveJson() (string, error) {
	var raw json.RawMessage
	err := h.Receive(&raw)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// Send arbitrary data
func (h *HttpEndpointConnection) Send(data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// log.Println(string(b))
	_, err = h.conn.Write(b)
	return err
}
