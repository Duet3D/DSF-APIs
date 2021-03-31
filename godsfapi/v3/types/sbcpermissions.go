package types

type SbcPermissions uint64

const (
	// None for no permissions set (default value)
	None SbcPermissions = 1 << iota
	// CommandExecution to execute generic commands
	CommandExecution
	// CodeInterceptionRead to intercept codes but don't interact with them
	CodeInterceptionRead
	// CodeInterceptionReadWrite to intercept codes in a blocking way
	// with options to resolve or cancel them
	CodeInterceptionReadWrite
	// ManagePlugins to install, load, unload and uninstall plugins.
	// Grants FS access to all third-party plugins, too.
	ManagePlugins
	// ServicePlugins runtime information (for internal purposes only, do not use)
	ServicePlugins
	// ManageUserSession to manage user sessions
	ManageUserSession
	// ObjectModelRead to read from the object model
	ObjectModelRead
	// ObjectModelReadWrite to read from and write to the object model
	ObjectModelReadWrite
	// RegisterHttpEndpoints to create new HTTP endpoints
	RegisterHttpEndpoints
	// ReadFilaments to read files in 0:/filaments
	ReadFilaments
	// WriteFilaments to write files in 0:/filaments
	WriteFilaments
	// ReadFirmware to read files in 0:/firmware
	ReadFirmware
	// WriteFirmware to write files in 0:/firmware
	WriteFirmware
	// ReadGCodes to read files in 0:/gcodes
	ReadGCodes
	// WriteGCodes to write files in 0:/gcodes
	WriteGCodes
	// ReadMacros to read files in 0:/macros
	ReadMacros
	// WriteMacros to write files in 0:/macros
	WriteMacros
	// ReadMenu to read files in 0:/menu
	ReadMenu
	// WriteMenu to write files in 0:/menu
	WriteMenu
	// ReadSystem to read files in 0:/sys
	ReadSystem
	// WriteSystem to write files in 0:/sys
	WriteSystem
	// ReadWeb to read files in 0:/www
	ReadWeb
	// WriteWeb to write files in 0:/www
	WriteWeb
	// FileSystemAccess to access files including all subdirectories of the virtual SD directory as DSF user
	FileSystemAccess
	// LaunchProcess to launch new processes
	LaunchProcess
	// NetworkAccess to communicat over network (stand-alone)
	NetworkAccess
	// SuperUser to launch processes as root user (for full device control - potentially dangerous)
	SuperUser
)
