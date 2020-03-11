package types

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
	HttpResponseTypeFile = "file"
)
