package usersessions

// UserSession represents a user session
type UserSession struct {
	// Id is the identifier of this session
	Id int64 `json:"id"`
	// AccessLevel of this session
	AccessLevel AccessLevel `json:"accessLevel"`
	// SessionType of this session
	SessionType SessionType `json:"sessionType"`
	// Origin of this session. For remote sessions this equals the remote IP address
	Origin string `json:"origin"`
	// OriginId is the corresponding identifier. If it is a remote session it is the remote port
	// else it defaults to the PID of the current process
	OriginId int `json:"originId"`
}

// AccessLevel defines what a user is allowed to do
type AccessLevel string

const (
	// ReadOnly means changes to the system and/or operation are not permitted
	ReadOnly AccessLevel = "readOnly"
	// ReadWrite means changes to the system and/or operation are permitted
	ReadWrite = "readWrite"
)

// SessionType is the type of user session
type SessionType string

const (
	// Local client
	Local SessionType = "local"
	// HTTP remote client
	HTTP = "http"
	// Telnet remote client
	Telnet = "telnet"
)
