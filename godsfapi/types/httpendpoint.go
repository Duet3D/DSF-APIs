package types

// Enumeration of supported HTTP request types
type HttpEndpointType string

const (
	// HTTP GET Request
	GET HttpEndpointType = "get"
	// HTTP POST Request
	POST = "post"
	// HTTP PUT Request
	PUT = "put"
	// HTTP PATCH Request
	PATCH = "patch"
	// HTTP TRACE Request
	TRACE = "trace"
	// HTTP DELETE Request
	DELETE = "delete"
	// HTTP OPTIONS Request
	OPTIONS = "options"
	// WebSocket request. This has not been implemented yet but is reserved for future usage
	WebSocket = "websocket"
)
