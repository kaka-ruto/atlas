// File: atlas_gui.go - Atlas GUI Integration Layer
// Template: pkg/gui/gui.go
// Purpose: AtlasGUI wrapper that integrates Atlas functionality with lazygit GUI
// Reference: Follows lazygit's GUI initialization and controller attachment patterns
package gui

import (
	"time"

	appTypes "github.com/jesseduffield/lazygit/pkg/app/types"
	"github.com/jesseduffield/lazygit/pkg/gui"
	"github.com/jesseduffield/lazygit/pkg/gui/controllers"

	atlasControllers "github.com/kaka-ruto/atlas/pkg/atlas/controllers"
	"github.com/kaka-ruto/atlas/pkg/atlas/wrapper"
)

// AtlasGUI wraps the lazygit GUI with Atlas functionality
type AtlasGUI struct {
	*gui.Gui
	atlasWrapper    *wrapper.AtlasWrapper
	atlasController *atlasControllers.AtlasController
}

// NewAtlasGUI creates a new Atlas-enhanced GUI from an existing lazygit GUI
func NewAtlasGUI(lazygitGUI *gui.Gui) (*AtlasGUI, error) {
	if lazygitGUI == nil {
		return nil, &AtlasError{"cannot create AtlasGUI with nil GUI"}
	}

	// Create the Atlas wrapper around the lazygit GUI
	atlasWrapper := wrapper.NewAtlasWrapper(lazygitGUI)
	atlasController := atlasControllers.NewAtlasController(atlasWrapper)

	// Create the Atlas GUI wrapper
	atlasGUI := &AtlasGUI{
		Gui:             lazygitGUI,
		atlasWrapper:    atlasWrapper,
		atlasController: atlasController,
	}

	return atlasGUI, nil
}

// GetAtlasWrapper returns the Atlas wrapper for external access
func (ag *AtlasGUI) GetAtlasWrapper() *wrapper.AtlasWrapper {
	if ag == nil {
		return nil
	}
	return ag.atlasWrapper
}

// GetAtlasController returns the Atlas controller for keybinding registration
func (ag *AtlasGUI) GetAtlasController() *atlasControllers.AtlasController {
	if ag == nil {
		return nil
	}
	return ag.atlasController
}

// AttachAtlasController attaches the Atlas controller to the global context
// This should be called after the GUI is fully initialized but before running
func (ag *AtlasGUI) AttachAtlasController() {
	if ag == nil || ag.Gui == nil {
		return
	}

	// Safety check: make sure the GUI is fully initialized
	// State, Contexts, and Global context need to exist
	if ag.Gui.State == nil || ag.Gui.State.Contexts.Global == nil {
		// GUI not fully initialized yet, skip for now
		// TODO: Find a better way to integrate with lazygit's controller system
		return
	}

	// Get the global context from the lazygit GUI
	// The Global context is where we attach controllers that work across all views
	globalContext := ag.Gui.State.Contexts.Global

	// Attach our Atlas controller to the global context
	// This adds our 'G' and 'A' keybindings globally
	controllers.AttachControllers(globalContext, ag.atlasController)
}

// Run starts the Atlas GUI with enhanced functionality
func (ag *AtlasGUI) Run(startArgs appTypes.StartArgs) error {
	if ag == nil || ag.Gui == nil {
		return &AtlasError{"AtlasGUI not properly initialized"}
	}

	// Run the standard lazygit GUI with Atlas enhancements
	// Note: Controller attachment is now handled in runWithAtlasIntegration
	return ag.Gui.RunAndHandleError(startArgs)
}

// RunAndHandleError runs the Atlas GUI and handles any errors
// This method provides the same interface as gui.Gui for compatibility
func (ag *AtlasGUI) RunAndHandleError(startArgs appTypes.StartArgs) error {
	if ag == nil || ag.Gui == nil {
		return &AtlasError{"AtlasGUI not properly initialized"}
	}

	// Create a custom Run function that attaches our controller at the right time
	return ag.runWithAtlasIntegration(startArgs)
}

// runWithAtlasIntegration runs the GUI with proper Atlas controller integration
func (ag *AtlasGUI) runWithAtlasIntegration(startArgs appTypes.StartArgs) error {
	// We'll use a goroutine to attach the controller after a short delay
	// This ensures the GUI is fully initialized before we try to attach
	go func() {
		// Wait a bit for the GUI to initialize
		// This is a temporary solution until we find a better integration point
		time.Sleep(100 * time.Millisecond)
		ag.AttachAtlasController()
	}()

	// Run the standard lazygit GUI
	return ag.Gui.RunAndHandleError(startArgs)
}

// AtlasError represents Atlas-specific errors
type AtlasError struct {
	message string
}

func (e *AtlasError) Error() string {
	return "Atlas Error: " + e.message
}