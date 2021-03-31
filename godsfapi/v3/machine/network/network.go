package network

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
	// HTTPS protocol
	HTTPS = "https"
	// FTP protocol
	FTP = "ftp"
	// SFTP procotol
	SFTP = "sftp"
	// Telnet protocol
	Telnet = "telnet"
	// SSH protocol
	SSH = "ssh"
)

// Network holds information about the network subsytem
type Network struct {
	// CorsSite indicates if the web server will allow cross-origin requests via Access-Control-Allow-Origin header
	CorsSite bool `json:"corsSite"`
	// Hostname of the machine
	Hostname string `json:"hostname"`
	// Interfaces is a list of available network interfaces
	Interfaces []NetworkInterface `json:"interfaces"`
	// Name of the machine
	Name string `json:"name"`
}

// NetworkInterface holds information about a network interface
type NetworkInterface struct {
	// ActiveProtocols is a list of active network protocols
	ActiveProtocols []NetworkProtocol `json:"activeProtocols"`
	// ActualIP tis the actual IPv4 address of the network adapter
	ActualIP string `json:"actualIP"`
	// ConfiguredIP is the IPv4 address of the network adapter
	ConfiguredIP string `json:"configuredIP"`
	// DnsServer configured for IPv4 traffic
	DnsServer string `json:"dnsServer"`
	// FirmwareVersion of the network interface (empty for unknonw)
	// This is primarily intended for the ESP8266-based network interfaces as used on the Duet WiFi
	FirmwareVersion string `json:"firmwareVersion"`
	// Gateway address for this network adapter (IPv4)
	Gateway string `json:"gateway"`
	// Mac address of the network adapter
	Mac string `json:"mac"`
	// NumReconnnects is the number of reconnect attempts
	NumReconnnects *int64 `json:"numReconnnects"`
	// Signal strength of the WiFi adapter (in dBm)
	Signal *int64 `json:"signal"`
	// Speed of the network interface (in MBit, nil if unknown, 0 if not connected)
	Speed *uint64 `json:"speed"`
	// Subnet mask of the network adapter
	Subnet string `json:"subnet"`
	// Type of this network interface
	Type InterfaceType `json:"type"`
}
