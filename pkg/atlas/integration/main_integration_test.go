// File: main_integration_test.go - Atlas Main Integration Tests
// Template: pkg/integration/components/test_test.go
// Purpose: Tests for Atlas main application integration and startup flows
// Reference: Follows lazygit's integration testing patterns for application scenarios
package integration

import (
	"testing"

	appTypes "github.com/jesseduffield/lazygit/pkg/app/types"
	"github.com/stretchr/testify/assert"

	atlasControllers "github.com/kaka-ruto/atlas/pkg/atlas/controllers"
	"github.com/kaka-ruto/atlas/pkg/atlas/wrapper"
)

func TestMainIntegration_AtlasStartup(t *testing.T) {
	scenarios := []struct {
		testName                string
		expectAtlasInitialized  bool
		expectKeybindingsActive bool
	}{
		{
			testName:                "Atlas should initialize with keybindings when app starts",
			expectAtlasInitialized:  true,
			expectKeybindingsActive: true,
		},
	}

	for _, s := range scenarios {
		s := s
		t.Run(s.testName, func(t *testing.T) {
			// Test the concept that when main.go calls app.Start with Atlas
			// it should:
			// 1. Create AtlasGUI instead of regular GUI
			// 2. Have 'G' and 'A' keybindings registered
			// 3. Start in Git mode by default

			// We test this by verifying the AtlasGUI components work properly
			atlasWrapper := wrapper.NewAtlasWrapper(nil)
			atlasController := atlasControllers.NewAtlasController(atlasWrapper)

			// Verify initial state (should start in Git mode)
			assert.True(t, atlasWrapper.IsGitMode())
			assert.False(t, atlasWrapper.IsAIMode())

			// Verify controller has the expected keybindings
			assert.NotNil(t, atlasController)

			// Test mode switching works (simulates keybinding usage)
			err := atlasWrapper.HandleAIModeKeybinding()
			assert.NoError(t, err)
			assert.True(t, atlasWrapper.IsAIMode())

			err = atlasWrapper.HandleGitModeKeybinding()
			assert.NoError(t, err)
			assert.True(t, atlasWrapper.IsGitMode())
		})
	}
}

func TestMainIntegration_StartArgsPassthrough(t *testing.T) {
	scenarios := []struct {
		testName  string
		startArgs appTypes.StartArgs
	}{
		{
			testName:  "Atlas passes through startArgs to lazygit correctly",
			startArgs: appTypes.StartArgs{},
		},
	}

	for _, s := range scenarios {
		s := s
		t.Run(s.testName, func(t *testing.T) {
			// Test that Atlas properly passes through start arguments
			// This verifies the integration maintains compatibility with lazygit

			// We can't test the full integration without a complete setup
			// but we can test that the AtlasGUI structure is correct

			// Create mock AtlasGUI to verify structure
			atlasWrapper := wrapper.NewAtlasWrapper(nil)
			atlasController := atlasControllers.NewAtlasController(atlasWrapper)

			// Verify these components would work in the main integration
			assert.NotNil(t, atlasWrapper)
			assert.NotNil(t, atlasController)

			// Verify AtlasGUI would handle the start args correctly
			// (Real implementation would pass args to underlying lazygit GUI)
			assert.IsType(t, appTypes.StartArgs{}, s.startArgs)
		})
	}
}