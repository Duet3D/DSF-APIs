package commands

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
