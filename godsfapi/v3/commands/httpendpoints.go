package commands

import (
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/httpendpoints"
)

// HttpEndpointCommand is used to either create or remove a custom HTTP endpoint
type HttpEndpointCommand struct {
	BaseCommand
	// EndpointType is type of HTTP request
	EndpointType httpendpoints.HttpEndpointType
	// Namespace of the plugin wanting to create a new third-party endpoint
	Namespace string
	// Path to the endpoint to register
	Path string
	// Whether this is an upload request
	IsUploadRequest bool
}

// NewAddHttpEndpoint registers a new HTTP endpoint via DuetWebServer. This will create a new HTTP endpoint under /machine/{Namespace}/{EndpointPath}.
// Returns a path to the UNIX socket which DuetWebServer will connect to whenever a matching HTTP request is received.
// A plugin using this command has to open a new UNIX socket with the given path that DuetWebServer can connect to
func NewAddHttpEndpoint(t httpendpoints.HttpEndpointType, ns, path string, isUploadRequest bool) *HttpEndpointCommand {
	return &HttpEndpointCommand{
		BaseCommand:     *NewBaseCommand("AddHttpEndpoint"),
		EndpointType:    t,
		Namespace:       ns,
		Path:            path,
		IsUploadRequest: isUploadRequest,
	}
}

// NewRemoveHttpEndpoint removes an existing HTTP endpoint.
func NewRemoveHttpEndpoint(t httpendpoints.HttpEndpointType, ns, path string) *HttpEndpointCommand {
	return &HttpEndpointCommand{
		BaseCommand:  *NewBaseCommand("RemoveHttpEndpoint"),
		EndpointType: t,
		Namespace:    ns,
		Path:         path,
	}
}

// ReceivedHttpRequest is the notification sent by the webserver when a
// new HTTP request is received
type ReceivedHttpRequest struct {
	// SessionId of the corresponding user session. This is -1 if it is an anonymous request
	SessionId int64
	// Queries is a map of HTTP query parameters
	Queries map[string]string
	// Headers is a map of HTTP headers
	Headers map[string]string
	// ContentType is the type of the request body
	ContentType string
	// Body content as plain-text or the filename where the body payload was saved
	// if HttpEndpointCommand.IsUploadRequest is true
	Body string
}

// NewReceivedHttpRequest creates a new default ReceivedHttpRequest
func NewReceivedHttpRequest() *ReceivedHttpRequest {
	return &ReceivedHttpRequest{}
}

// SendHttpResponse responds to a received HTTP request
type SendHttpResponse struct {
	// StatusCode (HTTP or WebSocket) to return. If this is greater or equal to 1000 the WbeSocket is closed
	StatusCode uint16
	// Response is the content to return. If this is null or empty and a WebSocket is conencted the connection is closed
	Response string
	// ResponseType of the content to return. Ignored if a WebSocket is connected.
	ResponseType HttpResponseType
}

// NewSendHttpResponse creates a new SendHttpResponse for the given status code, response body and type.
func NewSendHttpResponse(statusCode uint16, response string, t HttpResponseType) *SendHttpResponse {
	return &SendHttpResponse{
		StatusCode:   statusCode,
		Response:     response,
		ResponseType: t,
	}
}

// HttpResponseType enumerates supported HTTP responses
type HttpResponseType string

const (
	// StatusCode without payload
	StatusCode HttpResponseType = "statuscode"
	// PlainText UTF-8 response
	PlainText = "plaintext"
	// JSON formatted response
	JSON = "json"
	// File content. Response must hold the absolute path to the file to return
	File = "file"
)
