package machine

import "github.com/Duet3D/DSF-APIs/godsfapi/types"

// UserSession represents a user session
type UserSession struct {
	// Id is the identifier of this session
	Id int64 `json:"id"`
	// AccessLevel of this session
	AccessLevel types.AccessLevel `json:"accessLevel"`
	// SessionType of this session
	SessionType types.SessionType `json:"sessionType"`
	// Origin of this session. For remote sessions this equals the remote IP address
	Origin string `json:"origin"`
	// OriginId is the corresponding identifier. If it is a remote session it is the remote port
	// else it defaults to the PID of the current process
	OriginId int `json:"originId"`
}
