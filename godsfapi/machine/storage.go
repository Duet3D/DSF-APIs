package machine

// Storage holds information about a storage device
type Storage struct {
	// Mounted represents mount state
	Mounted bool `json:"mounted"`
	// Speed of the storage device in bytes/s (0 for unknown)
	Speed uint64 `json:"speed"`
	// Capacity is the total capacity of the storage device in bytes (0 for unknown)
	Capacity uint64 `json:"capacity"`
	// Free is the amount of space still available (nil if unknown)
	Free *uint64 `json:"free"`
	// OpenFiles is the number of currently open files or nil if unknown
	OpenFiles *uint64 `json:"openFiles"`
	// Path is the logical path of the storage device
	Path string `json:"path"`
}
