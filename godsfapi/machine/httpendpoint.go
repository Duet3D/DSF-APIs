package machine

import "github.com/Duet3D/DSF-APIs/godsfapi/types"

// HttpEndpoint represents an extra HTTP endpoint
type HttpEndpoint struct {
	// EndpointType is the type of this endpoint
	EndpointType types.HttpEndpointType
	// Namespace of this endpoint
	Namespace string
	// Path to the endpoint
	Path string
	// UnixSocket is the path to the corresponding UNIX socket
	UnixSocket string
}
