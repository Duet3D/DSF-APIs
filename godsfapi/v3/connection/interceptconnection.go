package connection

import (
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/commands"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/connection/initmessages"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/messages"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/types"
)

// InterceptConnection to intercept G/M/T-codes from the control server
//
// If this connection type is used to implement new G/M/T-codes, always call the commands.Flush
// command before further actions are started and make sure it returns true> before the code is further
// processed. This step is mandatory to guarantee that the new code is executed when all other codes have finished
// and not when a code is being fed for the internal G-code buffer. If the Flush command returns false, it
// is recommended to use CancelCode() to resolve the command. DCS follows the same pattern for
// internally processed codes, too.
// If a code from a macro file is intercepted, make sure to set the commands.CodeFlags.IsFromMacro
// flag if new codes are inserted, else they will be started when the macro file(s) have finished. This step
// is obsolete if a commands.SimpleCode is inserted.
type InterceptConnection struct {
	BaseCommandConnection
	Mode          initmessages.InterceptionMode
	Channels      []types.CodeChannel
	Filters       []string
	PriorityCodes bool
}

// Connect sends a InterceptInitMessage to the control server
// mode is the initmessages.InterceptionMode
// channels is an optional list of input channels to intercept codes from (empty list = all)
// filters to filter specific codes (see initmessages.InterceptInitMessage for details)
// priorityCodes to enable codes with CodeFlags.IsPrioritized
func (ic *InterceptConnection) Connect(mode initmessages.InterceptionMode, channels []types.CodeChannel, filters []string, priorityCodes bool, socketPath string) error {
	ic.Mode = mode
	if len(channels) == 0 {
		channels = types.AllChannels()
	}
	ic.Channels = channels
	ic.Filters = filters
	ic.PriorityCodes = priorityCodes
	iim := initmessages.NewInterceptInitMessage(mode, channels, filters, priorityCodes)
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

// Flush will wait for all previous codes of to finish
func (ic *InterceptConnection) Flush() (bool, error) {
	return ic.BaseCommandConnection.Flush(types.Unknown)
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
func (ic *InterceptConnection) ResolveCode(mType messages.MessageType, content string) error {
	return ic.Send(commands.NewResolve(mType, content))
}

// ResolveCodeMessage instructs the control server to resolved the last received code with
// the given message details
func (ic *InterceptConnection) ResolveCodeMessage(message messages.Message) error {
	return ic.Send(commands.NewResolve(message.Type, message.Content))
}
