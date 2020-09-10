package plugins

import "github.com/Duet3D/DSF-APIs/godsfapi/v3/types"

// Plugin represents a loaded plugin
type Plugin struct {
	// DwcFiles is a list of files for DWC
	DwcFiles []string `json:"dwcFiles"`
	// SbcFiles is a list of installed SBC files in the plugin directory
	SbcFiles []string `json:"sbcFiles"`
	// RrfFiles is a list of RRF files on the (virtual) SD excluding web files
	RrfFiles []string `json:"rrfFiles"`
	// Pid is the process ID of the plugin or -1 if not started
	// This may become 0 when the plugin has been stopped and the application
	// is being shut down
	Pid int64 `json:"pid"`
}

// PluginManifest holds information about the third-party plugin
type PluginManifest struct {
	// Name of the plugin
	Name string `json:"name"`
	// Author of the plugin
	Author string `json:"author"`
	// Version of the plugin
	Version string `json:"version"`
	// License of the plugin
	License string `json:"license"`
	// SourceRepository is a link to the source code repository
	SourceRepository string `json:"sourceRepository"`
	// DwcVersion is the major/minor compatible DWC version
	DwcVersion string `json:"dwcVersion"`
	// DwcDependencies is a list of DWC plugins this plugin depends on.
	// Circular dependencies are not supported.
	DwcDependencies []string `json:"dwcDependencies"`
	// DwcWebpackChunk is the name of the generated webpack chunk
	DwcWebpackChunk string `json:"dwcWebpackChunk"`
	// SbcRequired is set to true if a SBC is absolutely required for this plugin
	SbcRequired bool `json:"sbcRequired"`
	// SbcDsfVersion is the required DSF version for the plugin running on the SBC
	// (ignored if there is no SBC executable)
	SbcDsfVersion string `json:"sbcDsfVersion"`
	// SbcData is a list of objects holding key-value pairs of a plugin running on the SBC.
	// May be used to share data between plugins or between the SBC and web interface.
	SbcData map[string]string `json:"sbcData"`
	// SbcExecutable is the filename in the bin driectory used to start the plugin.
	// A plugin may provide different binaries in subdirectories per architecture.
	// Supported architectures are: arm, arm64, x86, x86_64
	SbcExecutable string `json:"sbcExecutable"`
	// SbcExecutableArguments are the command-line arguments for the executable
	SbcExecutableArguments string `json:"sbcExecutableArguments"`
	// SbcOutputRedirected defines if messages from stdout/stderr are output as generic messages
	SbcOutputRedirected bool `json:"sbcOutputRedirected"`
	// SbcPermissions is a list of permissions required by the plugin executable running on the SBC
	SbcPermissions []types.SbcPermissions `json:"sbcPermissions"`
	// SbcPackageDependencies is a list of packages this plugin depends on (host packages)
	SbcPackageDependencies []string `json:"sbcPackageDependencies"`
	// SbcPluginDependencies is a list of SBC plugins this plugin depends on.
	// Circular dependencies are not supported.
	SbcPluginDependencies []string `json:"sbcPluginDependencies"`
	// RrfVersion is the major/minore supported RRF version (optional)
	RrfVersion string `json:"rrfVersion"`
}
