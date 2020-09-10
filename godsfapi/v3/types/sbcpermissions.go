package types

type SbcPermissions string

const (
	// None for no permissions set (default value)
	None SbcPermissions = "none"
	// CommandExecution to execute generic commands
	CommandExecution = "commandExecution"
	// CodeInterceptionRead to intercept codes in a non-blocking way
	CodeInterceptionRead = "codeInterceptionRead"
	// CodeInterceptionReadWrite to intercept codes in a blocking way
	// with options to resolve or cancel them
	CodeInterceptionReadWrite = "codeInterceptionReadWrite"
	// ManagePlugins to install, load, unload and uninstall plugins
	ManagePlugins = "managePlugins"
	// ManageUserSession to manage user sessions
	ManageUserSession = "manageUserSessions"
	// ObjectModelRead to read from the object model
	ObjectModelRead = "objectModelRead"
	// ObjectModelReadWrite to read from and write to the object model
	ObjectModelReadWrite = "objectModelReadWrite"
	// RegisterHttpEndpoints to create new HTTP endpoints
	RegisterHttpEndpoints = "registerHttpEndpoints"
	// ReadFilaments to read files in 0:/filaments
	ReadFilaments = "readFilaments"
	// WriteFilaments to write files in 0:/filaments
	WriteFilaments = "writeFilaments"
	// ReadFirmware to read files in 0:/firmware
	ReadFirmware = "readFirmware"
	// WriteFirmware to write files in 0:/firmware
	WriteFirmware = "writeFirmware"
	// ReadGCodes to read files in 0:/gcodes
	ReadGCodes = "readGCodes"
	// WriteGCodes to write files in 0:/gcodes
	WriteGCodes = "writeGCodes"
	// ReadMacros to read files in 0:/macros
	ReadMacros = "readMacros"
	// WriteMacros to write files in 0:/macros
	WriteMacros = "writeMacros"
	// ReadMenu to read files in 0:/menu
	ReadMenu = "readMenu"
	// WriteMenu to write files in 0:/menu
	WriteMenu = "WriteMenu"
	// ReadSystem to read files in 0:/sys
	ReadSystem = "readSystem"
	// WriteSystem to write files in 0:/sys
	WriteSystem = "writeSystem"
	// ReadWeb to read files in 0:/www
	ReadWeb = "readWeb"
	// WriteWeb to write files in 0:/www
	WriteWeb = "writeWeb"
	// FileSystemAccess to access files outside the virtual SD directory (as DSF user)
	FileSystemAccess = "fileSystemAccess"
	// LaunchProcess to launch new processes
	LaunchProcess = "launchProcess"
	// NetworkAccess to communicat over network (stand-alone)
	NetworkAccess = "networkAccess"
	// SuperUser to launch processes as root user (for full device control - potentially dangerous)
	SuperUser = "superUser"
)
