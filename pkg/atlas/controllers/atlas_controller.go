// File: atlas_controller.go - Atlas Keybinding Controller
// Template: pkg/gui/controllers/global_controller.go
// Purpose: Handles global keybindings for Atlas mode switching ('G' and 'A' keys)
// Reference: Follows lazygit's controller pattern for handling global keybindings
package controllers

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/gui/types"

	"github.com/kaka-ruto/atlas/pkg/atlas/wrapper"
)

// AtlasController handles keybindings for Atlas mode switching
type AtlasController struct {
	atlasWrapper *wrapper.AtlasWrapper
}

// NewAtlasController creates a new Atlas controller
func NewAtlasController(atlasWrapper *wrapper.AtlasWrapper) *AtlasController {
	return &AtlasController{
		atlasWrapper: atlasWrapper,
	}
}

// GetKeybindings returns the keybindings for Atlas mode switching
func (ac *AtlasController) GetKeybindings(opts types.KeybindingsOpts) []*types.Binding {
	return []*types.Binding{
		{
			ViewName:        "",
			Key:             'G',
			Handler:         opts.Guards.NoPopupPanel(ac.handleGitModeKeybinding),
			Description:     "Switch to Git mode",
			DisplayOnScreen: true,
			Tag:             "atlas",
		},
		{
			ViewName:        "",
			Key:             'A',
			Handler:         opts.Guards.NoPopupPanel(ac.handleAIModeKeybinding),
			Description:     "Switch to AI mode",
			DisplayOnScreen: true,
			Tag:             "atlas",
		},
	}
}

// handleGitModeKeybinding handles the 'G' keybinding to switch to Git mode
func (ac *AtlasController) handleGitModeKeybinding() error {
	if ac.atlasWrapper != nil {
		return ac.atlasWrapper.HandleGitModeKeybinding()
	}
	return nil
}

// handleAIModeKeybinding handles the 'A' keybinding to switch to AI mode
func (ac *AtlasController) handleAIModeKeybinding() error {
	if ac.atlasWrapper != nil {
		return ac.atlasWrapper.HandleAIModeKeybinding()
	}
	return nil
}

// GetMouseKeybindings returns mouse keybindings (none for Atlas)
func (ac *AtlasController) GetMouseKeybindings(opts types.KeybindingsOpts) []*gocui.ViewMouseBinding {
	return nil
}

// GetOnClick returns onClick handler (none for Atlas)
func (ac *AtlasController) GetOnClick() func() error {
	return nil
}

// GetOnClickFocusedMainView returns onClick handler for main view (none for Atlas)
func (ac *AtlasController) GetOnClickFocusedMainView() func(mainViewName string, clickedLineIdx int) error {
	return nil
}

// Context returns the context (none for Atlas global controller)
func (ac *AtlasController) Context() types.Context {
	return nil
}

// GetOnRenderToMain returns onRenderToMain handler (none for Atlas)
func (ac *AtlasController) GetOnRenderToMain() func() {
	return nil
}

// GetOnFocus returns onFocus handler (none for Atlas)
func (ac *AtlasController) GetOnFocus() func(types.OnFocusOpts) {
	return nil
}

// GetOnFocusLost returns onFocusLost handler (none for Atlas)
func (ac *AtlasController) GetOnFocusLost() func(types.OnFocusLostOpts) {
	return nil
}