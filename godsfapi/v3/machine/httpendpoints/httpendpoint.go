package httpendpoints

const (
	// RepRapFirmwareNamespace is the namespace used for rr_ requests
	RepRapFirmwareNamespace = "rr_"
)

// HttpEndpoint represents an extra HTTP endpoint
type HttpEndpoint struct {
	// EndpointType is the type of this endpoint
	EndpointType HttpEndpointType `json:"endpointType"`
	// Namespace of this endpoint
	// May be RepRapFirmwareNamespace to register root-level rr_ requests (to emulate RRF poll requests)
	Namespace string `json:"namespace"`
	// Path to the endpoint
	Path string `json:"path"`
	// IsUploadRequest flages if this is a upload request
	// If set to true the whole body payload is written to a temporary file and the file path is
	// passed in the body field
	IsUploadRequest bool `json:"isUploadRequest"`
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
