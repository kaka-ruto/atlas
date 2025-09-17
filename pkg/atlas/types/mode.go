// File: mode.go - Atlas Mode Types
// Template: pkg/gui/types/common.go
// Purpose: Defines Atlas mode constants and state management structures
// Reference: Follows lazygit's pattern for defining GUI types and state enums
package types

// Mode represents the current mode of Atlas
type Mode int

const (
	GitMode Mode = iota
	AIMode
)

// AIModeState holds state specific to AI mode
type AIModeState struct {
	CurrentFile  string
	IsProcessing bool
}