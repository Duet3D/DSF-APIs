package messages

import (
	"fmt"
	"time"
)

// MessageType is the generic type of a message
type MessageType int64

// Valid MessageType values
const (
	Success MessageType = iota
	Warning
	Error
)

// Message is a generic container for messages
type Message struct {
	// Time at which the message was generated
	Time time.Time `json:"time"`
	// Type of this message
	Type MessageType `json:"type"`
	// Content of this message
	Content string `json:"content"`
}

// String converts this message to a RepRapFirmware-style message
func (m Message) String() string {
	switch m.Type {
	case Error:
		return fmt.Sprintf("Error: %s", m.Content)
	case Warning:
		return fmt.Sprintf("Warning: %s", m.Content)
	default:
		return m.Content
	}
}
