package commands

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
	// Body content as plain-text
	Body string
}

// NewReceivedHttpRequest creates a new default ReceivedHttpRequest
func NewReceivedHttpRequest() *ReceivedHttpRequest {
	return &ReceivedHttpRequest{}
}
