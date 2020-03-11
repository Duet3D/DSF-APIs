package commands

// SetMachineModel sets an atomic property in the machine model. Mameksure to
// acquire the read/wrtie lock first.
type SetMachineModel struct {
	BaseCommand
	// PropertyPtath to the property in the machine model
	PropertyPath string
	// Value is the string representation of the value to set
	Value string
}

// NewSetMachineModel creates a new SetMachineModel for the given key-value pair
func NewSetMachineModel(path, val string) *SetMachineModel {
	return &SetMachineModel{
		BaseCommand:  *NewBaseCommand("SetMachineModel"),
		PropertyPath: path,
		Value:        val,
	}
}
