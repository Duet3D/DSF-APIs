package commands

var getObjectModel = NewBaseCommand("GetObjectModel")
var syncObjectModel = NewBaseCommand("SyncObjectModel")
var lockObjectModel = NewBaseCommand("LockObjectModel")
var unlockObjectModel = NewBaseCommand("UnlockObjectModel")

// NewGetObjectModel returns a GetObjectModel command
func NewGetObjectModel() *BaseCommand {
	return getObjectModel
}

// NewSyncObjectModel returns a SyncObjectModel command
func NewSyncObjectModel() *BaseCommand {
	return syncObjectModel
}

// NewLockObjectModel returns a LockObjectModel command
func NewLockObjectModel() *BaseCommand {
	return lockObjectModel
}

// NewUnlockObjectModel returns a UnlockObjectModel command
func NewUnlockObjectModel() *BaseCommand {
	return unlockObjectModel
}

// SetObjectModel sets an atomic property in the machine model. Mameksure to
// acquire the read/wrtie lock first.
type SetObjectModel struct {
	BaseCommand
	// PropertyPtath to the property in the machine model
	PropertyPath string
	// Value is the string representation of the value to set
	Value string
}

// NewSetObjectModel creates a new SetObjectModel for the given key-value pair
func NewSetObjectModel(path, val string) *SetObjectModel {
	return &SetObjectModel{
		BaseCommand:  *NewBaseCommand("SetObjectModel"),
		PropertyPath: path,
		Value:        val,
	}
}

// PatchObjectModel applies as full patch to the object model. May be used only
// in non-SPI mode
type PatchObjectModel struct {
	BaseCommand
	// Key to update
	Key string
	// Patch to apply in JSON format
	Patch string
}

// NewPatchObjectModel creates a new SetObjectModel for the given key-patch pair
func NewPatchObjectModel(key, patch string) *PatchObjectModel {
	return &PatchObjectModel{
		BaseCommand: *NewBaseCommand("PatchObjectModel"),
		Key:         key,
		Patch:       patch,
	}
}
