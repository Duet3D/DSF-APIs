package types

// CodeType is the generic type of G/M/T-code or being a comment
type CodeType string

const (
	Comment CodeType = "C"
	GCode            = "G"
	MCode            = "M"
	TCode            = "T"
)
