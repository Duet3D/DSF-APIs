package initmessages

// InterceptionMode represents supported interception modes
type InterceptionMode string

const (
	// InterceptionModePre intercepts codes before they are internally processed by the control server
	InterceptionModePre InterceptionMode = "Pre"
	// InterceptionModePost intercepts codes after the initial processing of the control server
	// but before they are forwarded to the RepRapFirmware controller
	InterceptionModePost = "Post"
	// InterceptionModeExecuted receives notifications for executed codes. In this state the final
	// result can still be changed
	InterceptionModeExecuted = "Executed"
)

// InterceptInitMessage enters interception mode. Whenever a code is received the connection must respons with
// one of
// - commands.Ignore to pass through the code without modifications (i.e. it is ignored by the client)
// - commands.Resolve to resolve the current code and return a message (i.e. the client has handled this code)
// In addition the interceptor may issue custom commands once a code has been received.
// Do not attemt to perform commands before an intercepted code is received else the order of commands
// exectution cannot be guaranteed.
type InterceptInitMessage struct {
	BaseInitMessage
	// InterceptionMode selects when to intercept codes.
	InterceptionMode InterceptionMode
}

// NewInterceptInitMessage creates a new InterceptInitMessage for the given InterceptionMode
func NewInterceptInitMessage(iMode InterceptionMode) ClientInitMessage {
	return &InterceptInitMessage{
		BaseInitMessage:  NewBaseInitMessage(ConnectionModeIntercept),
		InterceptionMode: iMode,
	}
}
