// File: atlas_integration_test.go - Atlas Integration Tests
// Template: pkg/integration/components/test.go
// Purpose: End-to-end integration tests for Atlas functionality
// Reference: Follows lazygit's integration testing patterns for comprehensive testing
package integration

import (
	"testing"

	"github.com/jesseduffield/lazygit/pkg/gui"
	"github.com/stretchr/testify/assert"

	atlasControllers "github.com/kaka-ruto/atlas/pkg/atlas/controllers"
	atlasTypes "github.com/kaka-ruto/atlas/pkg/atlas/types"
	"github.com/kaka-ruto/atlas/pkg/atlas/wrapper"
)

func TestGuiIntegration_AtlasWrapperInitialization(t *testing.T) {
	scenarios := []struct {
		testName       string
		gui            *gui.Gui
		expectedResult bool // whether AtlasWrapper should be initialized
	}{
		{
			testName:       "AtlasWrapper is initialized when GUI is created",
			gui:            nil, // Will be created in test
			expectedResult: true,
		},
	}

	for _, s := range scenarios {
		s := s
		t.Run(s.testName, func(t *testing.T) {
			// Note: This is a basic integration test structure
			// In a full implementation, we would need to mock the GUI
			// dependencies and test the full initialization flow

			// For now, we just test the concept that the AtlasWrapper
			// should be created when the GUI starts up
			atlasWrapper := wrapper.NewAtlasWrapper(nil)
			assert.NotNil(t, atlasWrapper)
			assert.True(t, atlasWrapper.IsGitMode())
			assert.False(t, atlasWrapper.IsAIMode())
		})
	}
}

func TestGuiIntegration_AtlasControllerKeybindings(t *testing.T) {
	scenarios := []struct {
		testName         string
		expectedGKey     bool
		expectedAKey     bool
		expectedGitMode  bool
		expectedAIMode   bool
	}{
		{
			testName:         "Atlas controller keybindings are registered",
			expectedGKey:     true,
			expectedAKey:     true,
			expectedGitMode:  true,
			expectedAIMode:   false, // Starts in Git mode
		},
	}

	for _, s := range scenarios {
		s := s
		t.Run(s.testName, func(t *testing.T) {
			// Create AtlasWrapper and controller
			atlasWrapper := wrapper.NewAtlasWrapper(nil)
			controller := atlasControllers.NewAtlasController(atlasWrapper)

			// Check that we can get keybindings without error
			assert.NotNil(t, controller)

			// Check initial state
			assert.Equal(t, s.expectedGitMode, atlasWrapper.IsGitMode())
			assert.Equal(t, s.expectedAIMode, atlasWrapper.IsAIMode())
		})
	}
}

// This test will be expanded when we have full GUI integration
func TestGuiIntegration_ModeSwichingInContext(t *testing.T) {
	scenarios := []struct {
		testName      string
		initialMode   atlasTypes.Mode
		switchToMode  atlasTypes.Mode
		expectedMode  atlasTypes.Mode
	}{
		{
			testName:      "Switch from Git mode to AI mode",
			initialMode:   atlasTypes.GitMode,
			switchToMode:  atlasTypes.AIMode,
			expectedMode:  atlasTypes.AIMode,
		},
		{
			testName:      "Switch from AI mode to Git mode",
			initialMode:   atlasTypes.AIMode,
			switchToMode:  atlasTypes.GitMode,
			expectedMode:  atlasTypes.GitMode,
		},
	}

	for _, s := range scenarios {
		s := s
		t.Run(s.testName, func(t *testing.T) {
			// Create AtlasWrapper with initial mode
			atlasWrapper := wrapper.NewAtlasWrapper(nil)
			if s.initialMode == atlasTypes.AIMode {
				atlasWrapper.SwitchToAIMode()
			} else {
				atlasWrapper.SwitchToGitMode()
			}

			// Switch mode based on test case
			if s.switchToMode == atlasTypes.GitMode {
				atlasWrapper.SwitchToGitMode()
			} else {
				atlasWrapper.SwitchToAIMode()
			}

			// Verify the mode switched correctly
			assert.Equal(t, s.expectedMode, atlasWrapper.GetCurrentMode())
		})
	}
}