package machine

// Storage holds information about a storage device
type Storage struct {
	// Mounted represents mount state
	Mounted bool
	// Speed of the storage device in bytes/s (0 for unknown)
	Speed uint64
	// Capacity is the total capacity of the storage device in bytes (0 for unknown)
	Capacity uint64
	// Free is the amount of space still available (nil if unknown)
	Free *uint64
	// OpenFiles is the number of currently open files or nil if unknown
	OpenFiles *uint64
	// Path is the logical path of the storage device
	Path string
}
