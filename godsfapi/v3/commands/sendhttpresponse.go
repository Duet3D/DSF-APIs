package commands

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
