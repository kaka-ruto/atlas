// File: atlas_wrapper_test.go - Atlas Wrapper Tests
// Template: pkg/gui/context/list_renderer_test.go
// Purpose: Tests for Atlas wrapper and mode switching functionality
// Reference: Follows lazygit's testing patterns for wrapper components
package wrapper

import (
	"testing"

	"github.com/jesseduffield/lazygit/pkg/gui"
	"github.com/stretchr/testify/assert"

	atlasTypes "github.com/kaka-ruto/atlas/pkg/atlas/types"
)

func TestHandleGitModeKeybinding(t *testing.T) {
	scenarios := []struct {
		testName    string
		initialMode atlasTypes.Mode
		test        func(t *testing.T, wrapper *AtlasWrapper)
	}{
		{
			testName:    "switches to Git mode from AI mode",
			initialMode: atlasTypes.AIMode,
			test: func(t *testing.T, wrapper *AtlasWrapper) {
				assert.Equal(t, atlasTypes.AIMode, wrapper.currentMode)

				err := wrapper.HandleGitModeKeybinding()

				assert.NoError(t, err)
				assert.Equal(t, atlasTypes.GitMode, wrapper.currentMode)
				assert.True(t, wrapper.IsGitMode())
				assert.False(t, wrapper.IsAIMode())
			},
		},
		{
			testName:    "remains in Git mode when already in Git mode",
			initialMode: atlasTypes.GitMode,
			test: func(t *testing.T, wrapper *AtlasWrapper) {
				assert.Equal(t, atlasTypes.GitMode, wrapper.currentMode)

				err := wrapper.HandleGitModeKeybinding()

				assert.NoError(t, err)
				assert.Equal(t, atlasTypes.GitMode, wrapper.currentMode)
				assert.True(t, wrapper.IsGitMode())
				assert.False(t, wrapper.IsAIMode())
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.testName, func(t *testing.T) {
			mockGui := &gui.Gui{}
			wrapper := NewAtlasWrapper(mockGui)
			wrapper.currentMode = s.initialMode

			s.test(t, wrapper)
		})
	}
}

func TestHandleAIModeKeybinding(t *testing.T) {
	scenarios := []struct {
		testName    string
		initialMode atlasTypes.Mode
		test        func(t *testing.T, wrapper *AtlasWrapper)
	}{
		{
			testName:    "switches to AI mode from Git mode",
			initialMode: atlasTypes.GitMode,
			test: func(t *testing.T, wrapper *AtlasWrapper) {
				assert.Equal(t, atlasTypes.GitMode, wrapper.currentMode)

				err := wrapper.HandleAIModeKeybinding()

				assert.NoError(t, err)
				assert.Equal(t, atlasTypes.AIMode, wrapper.currentMode)
				assert.False(t, wrapper.IsGitMode())
				assert.True(t, wrapper.IsAIMode())
			},
		},
		{
			testName:    "remains in AI mode when already in AI mode",
			initialMode: atlasTypes.AIMode,
			test: func(t *testing.T, wrapper *AtlasWrapper) {
				assert.Equal(t, atlasTypes.AIMode, wrapper.currentMode)

				err := wrapper.HandleAIModeKeybinding()

				assert.NoError(t, err)
				assert.Equal(t, atlasTypes.AIMode, wrapper.currentMode)
				assert.False(t, wrapper.IsGitMode())
				assert.True(t, wrapper.IsAIMode())
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.testName, func(t *testing.T) {
			mockGui := &gui.Gui{}
			wrapper := NewAtlasWrapper(mockGui)
			wrapper.currentMode = s.initialMode

			s.test(t, wrapper)
		})
	}
}

func TestGetModeDisplayName(t *testing.T) {
	scenarios := []struct {
		testName        string
		mode            atlasTypes.Mode
		expectedDisplay string
	}{
		{
			testName:        "Git mode returns correct display name",
			mode:            atlasTypes.GitMode,
			expectedDisplay: "Git",
		},
		{
			testName:        "AI mode returns correct display name",
			mode:            atlasTypes.AIMode,
			expectedDisplay: "AI",
		},
	}

	for _, s := range scenarios {
		t.Run(s.testName, func(t *testing.T) {
			mockGui := &gui.Gui{}
			wrapper := NewAtlasWrapper(mockGui)
			wrapper.currentMode = s.mode

			display := wrapper.GetModeDisplayName()
			assert.Equal(t, s.expectedDisplay, display)
		})
	}
}

func TestShouldShowAIPanel(t *testing.T) {
	scenarios := []struct {
		testName        string
		mode            atlasTypes.Mode
		expectedResult  bool
	}{
		{
			testName:       "returns false for Git mode",
			mode:           atlasTypes.GitMode,
			expectedResult: false,
		},
		{
			testName:       "returns true for AI mode",
			mode:           atlasTypes.AIMode,
			expectedResult: true,
		},
	}

	for _, s := range scenarios {
		t.Run(s.testName, func(t *testing.T) {
			mockGui := &gui.Gui{}
			wrapper := NewAtlasWrapper(mockGui)
			wrapper.currentMode = s.mode

			result := wrapper.ShouldShowAIPanel()
			assert.Equal(t, s.expectedResult, result)
		})
	}
}

func TestShouldShowGitPanels(t *testing.T) {
	scenarios := []struct {
		testName        string
		mode            atlasTypes.Mode
		expectedResult  bool
	}{
		{
			testName:       "returns true for Git mode",
			mode:           atlasTypes.GitMode,
			expectedResult: true,
		},
		{
			testName:       "returns false for AI mode",
			mode:           atlasTypes.AIMode,
			expectedResult: false,
		},
	}

	for _, s := range scenarios {
		t.Run(s.testName, func(t *testing.T) {
			mockGui := &gui.Gui{}
			wrapper := NewAtlasWrapper(mockGui)
			wrapper.currentMode = s.mode

			result := wrapper.ShouldShowGitPanels()
			assert.Equal(t, s.expectedResult, result)
		})
	}
}