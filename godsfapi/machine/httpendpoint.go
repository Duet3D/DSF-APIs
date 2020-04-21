package machine

import "github.com/Duet3D/DSF-APIs/godsfapi/types"

// HttpEndpoint represents an extra HTTP endpoint
type HttpEndpoint struct {
	// EndpointType is the type of this endpoint
	EndpointType types.HttpEndpointType `json:"endpointType"`
	// Namespace of this endpoint
	Namespace string `json:"namespace"`
	// Path to the endpoint
	Path string `json:"path"`
	// UnixSocket is the path to the corresponding UNIX socket
	UnixSocket string `json:"unixSocket"`
}
