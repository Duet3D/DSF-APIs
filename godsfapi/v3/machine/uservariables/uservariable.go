// Deprecated: This package was deprected, please visit https://github.com/Duet3D/dsf-go.
package uservariables

// UserVariable is a key-value pair for user-defined variables
type UserVariable struct {
	// Name (key) of the variable
	Name string `json:"name"`
	// Value of the variable
	Value string `json:"value"`
}
