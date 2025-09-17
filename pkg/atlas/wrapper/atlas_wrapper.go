// File: atlas_wrapper.go - Atlas State Wrapper
// Template: pkg/gui/types/search_state.go
// Purpose: Atlas state management wrapper with mode switching and AI state tracking
// Reference: Follows lazygit's pattern for state management and mode tracking
package wrapper

import (
	"github.com/jesseduffield/lazygit/pkg/gui"

	atlasTypes "github.com/kaka-ruto/atlas/pkg/atlas/types"
)

// AtlasWrapper wraps the lazygit GUI with AI capabilities
type AtlasWrapper struct {
	lazygitGui  *gui.Gui
	currentMode atlasTypes.Mode
	aiMode      *atlasTypes.AIModeState
}

// NewAtlasWrapper creates a new Atlas wrapper around lazygit
func NewAtlasWrapper(lazygitGui *gui.Gui) *AtlasWrapper {
	return &AtlasWrapper{
		lazygitGui:  lazygitGui,
		currentMode: atlasTypes.GitMode,
		aiMode: &atlasTypes.AIModeState{
			CurrentFile: "",
		},
	}
}

// IsAIMode returns true if currently in AI mode
func (aw *AtlasWrapper) IsAIMode() bool {
	return aw.currentMode == atlasTypes.AIMode
}

// IsGitMode returns true if currently in Git mode
func (aw *AtlasWrapper) IsGitMode() bool {
	return aw.currentMode == atlasTypes.GitMode
}

// SwitchToAIMode switches to AI mode
func (aw *AtlasWrapper) SwitchToAIMode() {
	aw.currentMode = atlasTypes.AIMode
}

// SwitchToGitMode switches to Git mode
func (aw *AtlasWrapper) SwitchToGitMode() {
	aw.currentMode = atlasTypes.GitMode
}

// GetCurrentMode returns the current mode
func (aw *AtlasWrapper) GetCurrentMode() atlasTypes.Mode {
	return aw.currentMode
}

// GetCurrentFile returns the currently selected file
func (aw *AtlasWrapper) GetCurrentFile() string {
	return aw.aiMode.CurrentFile
}

// SetCurrentFile sets the currently selected file
func (aw *AtlasWrapper) SetCurrentFile(filepath string) {
	aw.aiMode.CurrentFile = filepath
}

// HandleGitModeKeybinding handles the 'G' keybinding to switch to Git mode
func (aw *AtlasWrapper) HandleGitModeKeybinding() error {
	aw.SwitchToGitMode()
	return nil
}

// HandleAIModeKeybinding handles the 'A' keybinding to switch to AI mode
func (aw *AtlasWrapper) HandleAIModeKeybinding() error {
	aw.SwitchToAIMode()
	return nil
}

// GetModeDisplayName returns a human-readable name for the current mode
func (aw *AtlasWrapper) GetModeDisplayName() string {
	switch aw.currentMode {
	case atlasTypes.GitMode:
		return "Git"
	case atlasTypes.AIMode:
		return "AI"
	default:
		return "Unknown"
	}
}

// ShouldShowAIPanel returns true if the AI panel should be displayed
func (aw *AtlasWrapper) ShouldShowAIPanel() bool {
	return aw.currentMode == atlasTypes.AIMode
}

// ShouldShowGitPanels returns true if the Git panels should be displayed
func (aw *AtlasWrapper) ShouldShowGitPanels() bool {
	return aw.currentMode == atlasTypes.GitMode
}