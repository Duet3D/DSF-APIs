package machine

// MessageBoxMode represents supported modes of displaying a message box
type MessageBoxMode uint64

const (
	// NoButtons displays a message box without any buttons
	NoButtons MessageBoxMode = iota
	// CloseOnly displays a message box with only a Close button
	CloseOnly
	// OkOnly displays a message box with only an Ok button which is supposed to send M292 when clicked
	OkOnly
	// OkCancel displays a message box with an Ok button that sends M292 P0 and
	// a Cancel button that sends M292 P1 when clicked
	OkCancel
)

// MessageBox holds information about the message box to show
type MessageBox struct {
	// Mode of the message box to display or nil if none is shown
	Mode *MessageBoxMode `json:"mode"`
	// Title of the message box
	Title string `json:"title"`
	// Message of the message box
	Message string `json:"message"`
	// AxisControls is a list of axis indices to show movement controls for
	AxisControls []uint8 `json:"axisControls"`
	// Seq is a counter that is incremented whenever a new message box is shown
	Seq int64 `json:"seq"`
}
