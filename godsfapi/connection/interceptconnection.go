package connection

import (
	"github.com/Duet3D/DSF-APIs/godsfapi/commands"
	"github.com/Duet3D/DSF-APIs/godsfapi/connection/initmessages"
	"github.com/Duet3D/DSF-APIs/godsfapi/types"
)

// InterceptConnection to intercept G/M/T-codes from the control server
type InterceptConnection struct {
	BaseCommandConnection
	Mode initmessages.InterceptionMode
}

// Connect sends a InterceptInitMessage to the control server
func (ic *InterceptConnection) Connect(mode initmessages.InterceptionMode, socketPath string) error {
	ic.Mode = mode
	iim := initmessages.NewInterceptInitMessage(mode)
	return ic.BaseConnection.Connect(iim, socketPath)
}

// ReceiveCode waits for a code to be intercepted
// Any other error than io.EOF requires the client to respond by either
// CancelCode(), IgnoreCode() or ResolveCode() because DCS will otherwise
// block while waiting for the Interceptor's response.
func (ic *InterceptConnection) ReceiveCode() (*commands.Code, error) {
	c := commands.NewCode()
	err := ic.Receive(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// CancelCode instructs the control server to cancel the last received code
func (ic *InterceptConnection) CancelCode() error {
	return ic.Send(commands.NewCancel())
}

// IgnoreCode tells the control server that this connection is not interested in the last received Code
// so it can continue with handling it.
func (ic *InterceptConnection) IgnoreCode() error {
	return ic.Send(commands.NewIgnore())
}

// ResolveCode instructs the control server to resolve the last received code with
// the given message details
func (ic *InterceptConnection) ResolveCode(mType types.MessageType, content string) error {
	return ic.Send(commands.NewResolve(mType, content))
}
