package plugins

import (
	"math"
	"strings"

	"github.com/Duet3D/DSF-APIs/godsfapi/v3/types"
)

// Plugin represents a loaded plugin
type Plugin struct {
	// DsfFiles is a list of files for the DSF plugin
	DsfFiles []string `json:"dsfFiles"`
	// DwcFiles is a list of files for  the DWC plugin
	DwcFiles []string `json:"dwcFiles"`
	// SdFiles is a list of files to be installed to the (virtual) SD excluding web files
	SdFiles []string `json:"sdFiles"`
	// Pid is the process ID of the plugin or -1 if not started
	// It is set to 0 while the plugin is being shut down
	Pid int64 `json:"pid"`
}

// PluginManifest holds information about the third-party plugin
type PluginManifest struct {
	// Id is the identifier of this plugin. May consist of letters and digits only (max length 32 chars)
	Id string `json:"id"`
	// Name of the plugin. May consist of letters, digits, dashed and underscores only (max length 64 chars)
	Name string `json:"name"`
	// Author of the plugin
	Author string `json:"author"`
	// Version of the plugin
	Version string `json:"version"`
	// License of the plugin
	License string `json:"license"`
	// Homepage is a link to the plugin homepage or source code repository
	Homepage string `json:"homepage"`
	// DwcVersion is the major/minor compatible DWC version
	DwcVersion string `json:"dwcVersion"`
	// DwcDependencies is a list of DWC plugins this plugin depends on.
	// Circular dependencies are not supported.
	DwcDependencies []string `json:"dwcDependencies"`
	// SbcRequired is set to true if a SBC is absolutely required for this plugin
	SbcRequired bool `json:"sbcRequired"`
	// SbcDsfVersion is the required DSF version for the plugin running on the SBC
	// (ignored if there is no SBC executable)
	SbcDsfVersion string `json:"sbcDsfVersion"`
	// SbcExecutable is the filename in the DSF directory used to start the plugin.
	// A plugin may provide different binaries in subdirectories per architecture.
	// Supported architectures are: arm, arm64, x86, x86_64
	SbcExecutable string `json:"sbcExecutable"`
	// SbcExtraExecutables is a list of other filenames in the DSF directory
	// that should be executable
	SbcExtraExecutables []string `json:"sbcExtraExecutables"`
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
	// Data is Custom plugin data to be populated in the object model
	// (DSF/DWC in SBC mode - or - DWC in standalone mode).
	// Before commands.SetPluginData can be used, corresponding properties must be registered via this property first!
	Data map[string]string `json:"data"`
}

// CheckVersion checks if the given version satisfies a required version
func CheckVersion(actual, required string) bool {
	if strings.TrimSpace(required) != "" {
		actualItems := strings.FieldsFunc(actual, func(r rune) bool {
			return r == '.' || r == '-' || r == '+'
		})
		requiredItems := strings.FieldsFunc(required, func(r rune) bool {
			return r == '.' || r == '-' || r == '+'
		})
		for i := 0; i < int(math.Min(float64(len(actualItems)), float64(len(requiredItems)))); i++ {
			if actualItems[i] != requiredItems[i] {
				return false
			}
		}
	}
	return true
}
