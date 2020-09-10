package initmessages

// SubscriptionMode represents supported subscription modes
type SubscriptionMode string

const (
	// SubscriptionModeFull receives full object model after each update
	// Generic messages may or may not be included in full object model. To keep
	// track of messages reliably it is strongly advised to creat a subscription
	// in Patch mode
	SubscriptionModeFull SubscriptionMode = "Full"
	// SubscriptionModePatch receives only updated JSON fragments of the object model
	SubscriptionModePatch = "Patch"
)

// SubscribeInitMessage enters subscription mode to receive either full object model or
// change patches after every update
type SubscribeInitMessage struct {
	BaseInitMessage
	// SubscriptionMode is the type of subscription
	SubscriptionMode SubscriptionMode
	// Filter is an optional filter path for mode Patch
	// Multiple filters can be used on one connection and they have to be delimited by one of these charaters: ['|', ',', ' ', '\r', '\n']
	// This setting is deprecated in favor of the new Filters list
	Filter string
	// Filters is an optional list of filter paths for mode Patch
	// The style of a filter is similar to XPath. For example, if you want to monitor only the current heater temperatures,
	// you can use the filter expression "heat/heaters[*]/current". Wildcards are supported either for full names or indices.
	// To get updates for an entire namespace, the ** wildcard can be used (for example heat/** for everything heat-related),
	// however it can be only used at the end of a filter expression.
	Filters []string
}

// NewSubscribeInitMessage returns a new SubscribeInitMessage for the given mode and filters
func NewSubscribeInitMessage(subMode SubscriptionMode, filters []string) ClientInitMessage {
	return &SubscribeInitMessage{
		BaseInitMessage:  NewBaseInitMessage(ConnectionModeSubscribe),
		SubscriptionMode: subMode,
		Filters:          filters,
	}
}
