package httpendpoints

// HttpEndpoint represents an extra HTTP endpoint
type HttpEndpoint struct {
	// EndpointType is the type of this endpoint
	EndpointType HttpEndpointType `json:"endpointType"`
	// Namespace of this endpoint
	Namespace string `json:"namespace"`
	// Path to the endpoint
	Path string `json:"path"`
	// UnixSocket is the path to the corresponding UNIX socket
	UnixSocket string `json:"unixSocket"`
}

// HttpEndpointType represents supported HTTP request types
type HttpEndpointType string

const (
	// GET Request
	GET HttpEndpointType = "GET"
	// POST Request
	POST = "POST"
	// PUT Request
	PUT = "PUT"
	// PATCH Request
	PATCH = "PATCH"
	// TRACE Request
	TRACE = "TRACE"
	// DELETE Request
	DELETE = "DELETE"
	// OPTIONS Request
	OPTIONS = "OPTIONS"
	// WebSocket request. This has not been implemented yet but is reserved for future usage
	WebSocket = "WebSocket"
)
