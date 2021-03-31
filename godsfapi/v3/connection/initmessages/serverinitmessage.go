package initmessages

const (
	// ProcotolVersion is the version the server needs to have to be compatible with
	// this client
	ProtocolVersion = 11
)

// ServerInitMessage is sent by the server to the client in JSON format once a connection
// has been established
type ServerInitMessage struct {
	// Version of the server-side API. A client is supposed to check if received API level is
	// greater than or equal to ExpectedServerVersion (i.e. its own API level) once a connection
	// has been established in order to ensure that all of the required commands are actually
	// supported by the control server.
	Version int64
	// Id is the unique connection ID assigned by the control server to allow clients to track their commands
	Id int64
}

// IsCompatible checks if the returned server API version is compatible with this client
func (s *ServerInitMessage) IsCompatible() bool {
	return s.Version >= ProtocolVersion
}
