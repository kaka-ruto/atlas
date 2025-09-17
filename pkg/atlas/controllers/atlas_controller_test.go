// File: atlas_controller_test.go - Atlas Controller Tests
// Template: pkg/gui/controllers/local_commits_controller_test.go
// Purpose: Tests for Atlas keybinding controller functionality
// Reference: Follows lazygit's controller testing patterns with scenario-based tests
package controllers

import (
	"testing"

	"github.com/jesseduffield/lazygit/pkg/config"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/stretchr/testify/assert"

	atlasTypes "github.com/kaka-ruto/atlas/pkg/atlas/types"
	"github.com/kaka-ruto/atlas/pkg/atlas/wrapper"
)

func TestAtlasController_GetKeybindings(t *testing.T) {
	scenarios := []struct {
		name                    string
		testName                string
		expectedBindingsCount   int
		expectedGitModeKey      interface{}
		expectedAIModeKey       interface{}
		expectedGitDescription  string
		expectedAIDescription   string
		expectedGitHandler      bool // whether handler exists
		expectedAIHandler       bool // whether handler exists
	}{
		{
			testName:                "returns correct keybindings for Atlas mode switching",
			expectedBindingsCount:   2,
			expectedGitModeKey:      'G',
			expectedAIModeKey:       'A',
			expectedGitDescription:  "Switch to Git mode",
			expectedAIDescription:   "Switch to AI mode",
			expectedGitHandler:      true,
			expectedAIHandler:       true,
		},
	}

	for _, s := range scenarios {
		s := s
		t.Run(s.testName, func(t *testing.T) {
			// Arrange
			atlasWrapper := wrapper.NewAtlasWrapper(nil)
			atlasWrapper.SwitchToGitMode()

			controller := NewAtlasController(atlasWrapper)

			opts := types.KeybindingsOpts{
				GetKey: func(key string) types.Key {
					return key // Simple passthrough for test
				},
				Config: config.KeybindingConfig{}, // Empty config for test
				Guards: types.KeybindingGuards{
					NoPopupPanel: func(f func() error) func() error { return f },
				},
			}

			// Act
			bindings := controller.GetKeybindings(opts)

			// Assert
			assert.Len(t, bindings, s.expectedBindingsCount)

			// Check Git mode keybinding
			gitBinding := bindings[0]
			assert.Equal(t, s.expectedGitModeKey, gitBinding.Key)
			assert.Equal(t, s.expectedGitDescription, gitBinding.Description)
			assert.NotNil(t, gitBinding.Handler)

			// Check AI mode keybinding
			aiBinding := bindings[1]
			assert.Equal(t, s.expectedAIModeKey, aiBinding.Key)
			assert.Equal(t, s.expectedAIDescription, aiBinding.Description)
			assert.NotNil(t, aiBinding.Handler)

			// Test that handlers work without error
			if s.expectedGitHandler {
				err := gitBinding.Handler()
				assert.NoError(t, err)
				assert.True(t, atlasWrapper.IsGitMode())
			}

			if s.expectedAIHandler {
				err := aiBinding.Handler()
				assert.NoError(t, err)
				assert.True(t, atlasWrapper.IsAIMode())
			}
		})
	}
}

func TestAtlasController_HandleGitModeKeybinding(t *testing.T) {
	scenarios := []struct {
		testName      string
		initialMode   atlasTypes.Mode
		expectedMode  atlasTypes.Mode
		shouldSucceed bool
	}{
		{
			testName:      "switches to Git mode from AI mode",
			initialMode:   atlasTypes.AIMode,
			expectedMode:  atlasTypes.GitMode,
			shouldSucceed: true,
		},
		{
			testName:      "remains in Git mode when already in Git mode",
			initialMode:   atlasTypes.GitMode,
			expectedMode:  atlasTypes.GitMode,
			shouldSucceed: true,
		},
	}

	for _, s := range scenarios {
		s := s
		t.Run(s.testName, func(t *testing.T) {
			// Arrange
			atlasWrapper := wrapper.NewAtlasWrapper(nil)
			if s.initialMode == atlasTypes.AIMode {
				atlasWrapper.SwitchToAIMode()
			} else {
				atlasWrapper.SwitchToGitMode()
			}

			controller := NewAtlasController(atlasWrapper)
			assert.NotNil(t, controller)

			// Act
			err := atlasWrapper.HandleGitModeKeybinding()

			// Assert
			if s.shouldSucceed {
				assert.NoError(t, err)
				assert.Equal(t, s.expectedMode, atlasWrapper.GetCurrentMode())
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestAtlasController_HandleAIModeKeybinding(t *testing.T) {
	scenarios := []struct {
		testName      string
		initialMode   atlasTypes.Mode
		expectedMode  atlasTypes.Mode
		shouldSucceed bool
	}{
		{
			testName:      "switches to AI mode from Git mode",
			initialMode:   atlasTypes.GitMode,
			expectedMode:  atlasTypes.AIMode,
			shouldSucceed: true,
		},
		{
			testName:      "remains in AI mode when already in AI mode",
			initialMode:   atlasTypes.AIMode,
			expectedMode:  atlasTypes.AIMode,
			shouldSucceed: true,
		},
	}

	for _, s := range scenarios {
		s := s
		t.Run(s.testName, func(t *testing.T) {
			// Arrange
			atlasWrapper := wrapper.NewAtlasWrapper(nil)
			if s.initialMode == atlasTypes.AIMode {
				atlasWrapper.SwitchToAIMode()
			} else {
				atlasWrapper.SwitchToGitMode()
			}

			controller := NewAtlasController(atlasWrapper)
			assert.NotNil(t, controller)

			// Act
			err := atlasWrapper.HandleAIModeKeybinding()

			// Assert
			if s.shouldSucceed {
				assert.NoError(t, err)
				assert.Equal(t, s.expectedMode, atlasWrapper.GetCurrentMode())
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestAtlasController_Creation(t *testing.T) {
	scenarios := []struct {
		testName      string
		atlasWrapper  *wrapper.AtlasWrapper
		expectNil     bool
	}{
		{
			testName: "creates controller with valid wrapper.AtlasWrapper",
			atlasWrapper: wrapper.NewAtlasWrapper(nil),
			expectNil: false,
		},
		{
			testName:     "creates controller with nil wrapper.AtlasWrapper",
			atlasWrapper: nil,
			expectNil:    false, // Controller should still be created, just with nil wrapper
		},
	}

	for _, s := range scenarios {
		s := s
		t.Run(s.testName, func(t *testing.T) {
			// Act
			controller := NewAtlasController(s.atlasWrapper)

			// Assert
			if s.expectNil {
				assert.Nil(t, controller)
			} else {
				assert.NotNil(t, controller)
				assert.Equal(t, s.atlasWrapper, controller.atlasWrapper)
			}
		})
	}
}