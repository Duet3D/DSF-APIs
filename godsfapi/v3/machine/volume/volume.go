package volume

// Volume holds information about a storage device
type Volume struct {
	// Capacity is the total capacity of the storage device in bytes (0 for unknown)
	Capacity uint64 `json:"capacity"`
	// FreeSpace is the amount of space still available (nil if unknown)
	FreeSpace *uint64 `json:"freeSpace"`
	// Mounted represents mount state
	Mounted bool `json:"mounted"`
	// Name of this volume
	Name string `json:"name"`
	// OpenFiles is the number of currently open files or nil if unknown
	OpenFiles *uint64 `json:"openFiles"`
	// Path is the logical path of the storage device
	Path string `json:"path"`
	// Speed of the storage device in bytes/s (0 for unknown)
	Speed uint64 `json:"speed"`
}
