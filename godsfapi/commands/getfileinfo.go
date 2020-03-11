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
