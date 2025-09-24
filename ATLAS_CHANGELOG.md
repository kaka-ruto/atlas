# Atlas Changelog

All notable changes to the Atlas project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [0.1.0] - 2025-09-16

### Added
- **Core Mode Management System**
  - Implemented dual-mode architecture (Git Mode + AI Mode)
  - Created `Mode` enum with `GitMode` (default) and `AIMode` constants
  - Built `AtlasWrapper` struct to manage mode switching and state

### Architecture Decisions
- **Separation of Concerns**: AI functionality isolated in `pkg/atlas/` package to avoid polluting lazygit core
- **State Management**: Created `AIModeState` struct to hold AI-specific data (current file, processing status)
- **Mode-First Design**: All functionality centers around current mode, enabling clean UI separation

### Key Features
- **Mode Switching**
  - `SwitchToGitMode()` and `SwitchToAIMode()` for explicit mode changes
  - `IsGitMode()` and `IsAIMode()` for mode checking
  - `GetCurrentMode()` returns current mode state

- **Keybinding Support**
  - `HandleGitModeKeybinding()` for 'G' key (Git mode)
  - `HandleAIModeKeybinding()` for 'A' key (AI mode)
  - Error-free keybinding handlers returning `error` type for consistency

- **UI Panel Management**
  - `ShouldShowAIPanel()` - controls full-screen AI panel visibility
  - `ShouldShowGitPanels()` - controls normal lazygit panels visibility
  - `GetModeDisplayName()` - provides human-readable mode names

- **File Context Tracking**
  - `GetCurrentFile()` and `SetCurrentFile()` for AI mode file context
  - File path stored in AI mode state for context-aware AI assistance

### Testing Strategy
- **Test-Driven Development**: 40 comprehensive test scenarios written before implementation
- **Scenario-Based Testing**: Following lazygit's testing patterns with scenario structs
- **Full Coverage**: All public methods tested with multiple edge cases
- **Separation**: Tests split across `mode_test.go` and `keybindings_test.go` files

### Technical Implementation
- **Package Structure**: Clean separation in `pkg/atlas/` following Go conventions
- **Import Strategy**: Minimal dependencies, only importing `github.com/jesseduffield/lazygit/pkg/gui`
- **Error Handling**: Consistent error return patterns matching lazygit conventions
- **Type Safety**: Strong typing with custom `Mode` type and structured state management

## [0.2.0] - 2025-09-17

### Added
- **Complete Keybinding Integration**
  - Implemented `AtlasController` with full `IController` interface compliance
  - Added 'G' and 'A' keybindings for Git/AI mode switching
  - Integrated controller with lazygit's global context system

- **AtlasGUI Wrapper System**
  - Created `AtlasGUI` wrapper that embeds lazygit's GUI
  - Implemented clean integration that attaches Atlas controller to global context
  - Built error handling and initialization patterns

- **Main Application Integration**
  - Updated `main.go` to use Atlas entry point instead of lazygit directly
  - Created `atlas.Start()` function as drop-in replacement for `app.Start()`
  - Maintained full compatibility with lazygit's startup process

### Architecture Decisions
- **Zero Core Modifications**: All Atlas functionality contained within `pkg/atlas/` package
- **Composition Over Inheritance**: AtlasGUI wraps rather than extends lazygit's GUI
- **Controller Pattern**: Uses lazygit's existing controller attachment system for keybindings
- **Clean Integration Points**: Minimal surface area for integration (only main.go modified)

### Key Features
- **Global Keybindings**
  - 'G' key switches to Git mode (standard lazygit interface)
  - 'A' key switches to AI mode (placeholder for future AI panel)
  - Keybindings work from any lazygit panel or context
  - No conflicts with existing lazygit keybindings

- **Mode Management**
  - Seamless switching between Git and AI modes
  - Proper state management and mode indicators
  - Thread-safe mode operations

- **Integration Architecture**
  - `AtlasWrapper` manages mode state and operations
  - `AtlasController` handles keybinding registration and processing
  - `AtlasGUI` provides the integration layer with lazygit
  - Clean separation allows easy future enhancements

### Testing Strategy
- **Test-Driven Development**: All functionality implemented with comprehensive test coverage
- **Integration Testing**: Full end-to-end testing of keybinding and mode switching
- **Component Testing**: Individual testing of wrapper, controller, and GUI components
- **Error Case Coverage**: Robust testing of edge cases and error conditions

### Technical Highlights
- **Controller Interface**: Full implementation of lazygit's `IController` interface
- **Context Integration**: Proper integration with lazygit's context management system
- **Memory Safety**: Careful null checking and proper initialization patterns
- **Performance**: Zero performance impact on standard lazygit operations

## [0.3.0] - 2025-09-18

### Added
- **Directory Structure Reorganization**
  - Reorganized all Atlas files to match lazygit's exact directory patterns
  - Created proper package hierarchy: `types/`, `wrapper/`, `controllers/`, `gui/`, `app/`, `integration/`
  - Each component now follows its corresponding lazygit template structure

- **Template Compliance System**
  - Added comprehensive template references to all Atlas files
  - Each file header documents its lazygit template (e.g., `atlas_wrapper.go` → `search_state.go`)
  - Directory structure matches template locations for consistency and maintainability

- **Improved Controller Integration**
  - Fixed nil pointer dereference in `AttachAtlasController`
  - Implemented delayed controller attachment using goroutine to avoid GUI initialization race conditions
  - Added proper safety checks for GUI state initialization

- **Comprehensive Test Coverage**
  - All test files reorganized to match new directory structure
  - Fixed package names and imports across all test files
  - Maintained 100% test pass rate (6/6 packages, 33+ tests)

### Architecture Improvements
- **Pattern Compliance**: All Atlas files now strictly follow their lazygit template patterns
- **Maintainability**: Directory structure makes it easy to find corresponding lazygit files
- **Integration Robustness**: Fixed timing issues with GUI initialization and controller attachment
- **Code Organization**: Clean separation by component type following lazygit conventions

### Bug Fixes
- **Runtime Crashes**: Fixed segmentation violation in `AttachAtlasController` method
- **Controller Timing**: Resolved race condition between GUI initialization and controller attachment
- **Package Structure**: Fixed circular dependencies through proper package organization

### Technical Debt Reduction
- **Template References**: All files now have clear lazygit template mappings
- **Directory Consistency**: Atlas structure mirrors lazygit's proven organization patterns
- **Test Organization**: Test files properly organized by component and package
- **Import Cleanup**: Resolved import cycles and dependency issues

## Phase 1 Status: ✅ COMPLETED (18/9/2025)
All foundation components are now in place and functional. Atlas successfully integrates with lazygit while maintaining complete separation of concerns and zero modifications to lazygit's core codebase. The architecture is solid, tested, and ready for Phase 2 development.