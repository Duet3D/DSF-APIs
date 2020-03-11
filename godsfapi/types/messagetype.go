package types

// MessageType is the generic type of a message
type MessageType int64

const (
	Success MessageType = iota
	Warning
	Error
)
