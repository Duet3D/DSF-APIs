package machine

const (
	// DefaultName of the machine
	DefaultName = "My Duet"
	// DefaultHostName as fallback if Name is invalid
	DefaultHostName = "duet"
	// DefaultPassword of the machine
	DefaultPassword = "reprap"
)

// InterfaceType represents supported interface types
type InterfaceType string

const (
	// WiFi is a wireless network interface
	WiFi InterfaceType = "wifi"
	// LAN is a wired network interface
	LAN = "lan"
)

// NetworkProtocol represents supported network protocols
type NetworkProtocol string

const (
	// HTTP protocol
	HTTP NetworkProtocol = "http"
	// FTP protocol
	FTP = "ftp"
	// Telnet protocol
	Telnet = "telnet"
)

// Network holds information about the network subsytem
type Network struct {
	// Name of the machine
	Name string
	// Hostname of the machine
	Hostname string
	// Password required to access this machine
	Password string
	// Interfaces is a list of available network interfaces
	Interfaces []NetworkInterface
}

// NeworkInterface holds information about a network interface
type NetworkInterface struct {
	// Type of this network interface
	Type InterfaceType
	// FirmwareVersion of the network interface (empty for unknonw)
	// This is primarily intended for the ESP8266-based network interfaces as used on the Duet WiFi
	FirmwareVersion string
	// Speed of the network interface (in MBit, nil if unknown, 0 if not connected)
	Speed *uint64
	// Signal strength of the WiFi adapter (in dBm)
	Signal int64
	// MacAddress of the network adapter
	MacAddress string
	// ConfiguredIP is the IPv4 address of the network adapter
	ConfiguredIP string
	// ActualIP tis the actual IPv4 address of the network adapter
	ActualIP string
	// Subnet mask of the network adapter
	Subnet string
	// Gateway address for this network adapter
	Gateway string
	// NumReconnnects is the number of reconnect attempts
	NumReconnnects uint64
	// ActiveProtocols is a list of active network protocols
	ActiveProtocols []NetworkProtocol
}
