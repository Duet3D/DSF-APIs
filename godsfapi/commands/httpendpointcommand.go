package commands

import "github.com/Duet3D/DSF-APIs/godsfapi/types"

// HttpEndpointCommand is used to either create or remove a custom HTTP endpoint
type HttpEndpointCommand struct {
	BaseCommand
	EndpointType types.HttpEndpointType
	Namespace    string
	Path         string
}

// NewAddHttpEndpoint registers a new HTTP endpoint via DuetWebServer. This will create a new HTTP endpoint under /machine/{Namespace}/{EndpointPath}.
// Returns a path to the UNIX socket which DuetWebServer will connect to whenever a matching HTTP request is received.
// A plugin using this command has to open a new UNIX socket with the given path that DuetWebServer can connect to
func NewAddHttpEndpoint(t types.HttpEndpointType, ns, path string) *HttpEndpointCommand {
	return &HttpEndpointCommand{
		BaseCommand:  *NewBaseCommand("AddHttpEndpoint"),
		EndpointType: t,
		Namespace:    ns,
		Path:         path,
	}
}

// NewRemoveHttpEndpoint removes an existing HTTP endpoint.
func NewRemoveHttpEndpoint(t types.HttpEndpointType, ns, path string) *HttpEndpointCommand {
	return &HttpEndpointCommand{
		BaseCommand:  *NewBaseCommand("RemoveHttpEndpoint"),
		EndpointType: t,
		Namespace:    ns,
		Path:         path,
	}
}
