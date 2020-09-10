package commands

// GetFileInfo will initiate analysis of a G-code file and returns
// ParsedFileInfo when ready.
type GetFileInfo struct {
	BaseCommand
	// Filename of the file to analyse
	FileName string
}

// NewGetFileInfo creates a new GetFileInfo for the given file name
func NewGetFileInfo(fileName string) *GetFileInfo {
	return &GetFileInfo{
		BaseCommand: *NewBaseCommand("GetFileInfo"),
		FileName:    fileName,
	}
}

// ResolvePath will resolve a RepRapFirmware-style path to an actual file system path
type ResolvePath struct {
	BaseCommand
	// POath that is RepRapFirmware-compatible
	Path string
}

// NewResolvePath creates a new ResolvePath for the given path
func NewResolvePath(path string) *ResolvePath {
	return &ResolvePath{
		BaseCommand: *NewBaseCommand("ResolvePath"),
		Path:        path,
	}
}
