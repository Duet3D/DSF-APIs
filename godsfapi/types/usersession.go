package types

// AccessLevel defines what a user is allowed to do
type AccessLevel string

const (
	// Changes tot the system and/or operation are not permitted
	ReadOnly AccessLevel = "ReadOnly"
	// Changes tot he system and/or operation are permitted
	ReadWrite = "ReadWrite"
)

// SessionType is the type of user session
type SessionType string

const (
	// Local client
	Local SessionType = "Local"
	// HTTP remote client
	SessionTypeHTTP = "HTTP"
	// Telnet remote client
	SesstionTypeTelnet = "Telnet"
)
