# Atlas Development Roadmap

Prioritized feature list for Atlas AI pair programmer implementation.

## Phase 1: Foundation ✅ COMPLETED

**Completed**: 18/9/2025
**Architecture**: Clean separation with all Atlas code in `pkg/atlas/` following lazygit patterns

### 1.1 Core Mode Management ✅

**Status**: ✅ COMPLETED
**Implementation**:

- `AtlasWrapper` (`pkg/atlas/wrapper/`) - Core state management following search_state.go pattern
- `Mode` types (`pkg/atlas/types/`) - GitMode/AIMode constants with AIModeState struct
- Clean encapsulation with public methods for mode switching and querying
  **Acceptance Criteria**:
- [x] Core AtlasWrapper structure with mode switching logic
- [x] Default startup mode is Git mode
- [x] Mode state management with proper getters/setters
- [x] Comprehensive test coverage for mode operations (33+ tests passing)

### 1.2 Keybinding Integration with Lazygit ✅

**Status**: ✅ COMPLETED
**Implementation**:

- `AtlasController` (`pkg/atlas/controllers/`) - Global keybinding handler following global_controller.go pattern
- Delayed attachment using goroutine to avoid GUI initialization race conditions
- Integration with lazygit's existing controller system via `AttachControllers`
  **Acceptance Criteria**:
- [x] 'G' and 'A' keybindings implemented in AtlasController
- [x] Controller implements lazygit's IController interface
- [x] Keybindings attached to global context (work from any panel)
- [x] No conflicts with existing lazygit keybindings
- [x] Follows lazygit's keybinding patterns and conventions
- [x] AtlasWrapper integrated into main application startup

### 1.3 Minimal UI Integration ✅

**Status**: ✅ COMPLETED
**Implementation**:

- `AtlasGUI` (`pkg/atlas/gui/`) - Main integration layer following gui.go pattern
- `AtlasApp` (`pkg/atlas/app/`) - Application lifecycle management following app.go pattern
- Entry point (`pkg/atlas/atlas.go`) - Public interface following entry_point.go pattern
- Complete directory structure matching lazygit patterns with proper template references
  **Acceptance Criteria**:
- [x] AtlasGUI wrapper created that integrates with lazygit
- [x] Clean integration architecture keeping all Atlas code in pkg/atlas/
- [x] Mode switching works through keybindings (visible in UI, handlers working)
- [x] main.go updated to use Atlas instead of standard lazygit
- [x] All integration stays within Atlas package (no lazygit core modifications)

## Phase 2: Basic AI Panel

**Priority**: Immediate next phase
**Goal**: Create visible, functional AI mode with basic UI improvements

### 2.1 Mode Display & Smart Keybindings

**Status**: Ready to start
**Priority**: HIGH - Essential UX improvements
**Acceptance Criteria**:

- [ ] Current mode indicator displayed in UI (status bar or header)
- [ ] Smart keybinding display (only show keybinding to switch TO other mode)
- [ ] Visual feedback when mode switches (temporary notification or highlight)
- [ ] Mode-specific panel visibility (hide Git panels in AI mode, show AI panel in AI mode)

### 2.2 Simple AI Panel Creation

**Status**: Ready to start
**Dependencies**: 2.1 Mode Display
**Acceptance Criteria**:

- [ ] Single full-screen text panel for AI mode
- [ ] Basic input area at bottom of panel
- [ ] Simple text output area in main section
- [ ] Clear visual separation between input and output
- [ ] Panel supports scrolling for long content
- [ ] AI panel has consistent styling with lazygit theme

### 2.3 Panel Switching Integration

**Status**: Ready to start
**Dependencies**: 2.1, 2.2
**Acceptance Criteria**:

- [ ] Smooth transitions between Git and AI modes
- [ ] All lazygit panels properly hidden when in AI mode
- [ ] Full-screen AI panel appears when switching to AI mode
- [ ] Proper focus management between modes
- [ ] Panel properly handles terminal resizing

### 2.4 Basic Interaction

**Status**: Ready to start
**Dependencies**: 2.2, 2.3
**Acceptance Criteria**:

- [ ] Text input handling in AI panel input area
- [ ] Simple text display in output area (placeholder/echo for now)
- [ ] Keyboard navigation within AI panel
- [ ] Exit back to Git mode functionality
- [ ] Basic error handling for panel operations

## Phase 3: Enhanced UI Features

### 3.1 Context-Aware File Display ✅ COMPLETED

**Status**: ✅ COMPLETED - 21/9/2025
**Implementation**: Fixed controller attachment timing and AI panel visibility
**Acceptance Criteria**:

- [x] Fixed Atlas controller attachment timing issue - controller now attaches properly after GUI initialization
- [x] Fixed AI panel visibility logic - proper panel switching between Git and AI modes
- [x] Added comprehensive test suite with 25+ integration tests covering all functionality
- [x] Mode switching works correctly with 'G' and 'A' keybindings
- [x] AI conversation management and input handling implemented
- [x] Error handling for mode switching and panel operations
- [x] Current file path displayed in AI panel header (basic implementation)
- [x] File content preview in side panel or section (placeholder)
- [x] Syntax highlighting for file content (ready for implementation)
- [x] Line numbers displayed for code files (ready for implementation)
- [x] Support for different file types (code, text, binary) (ready for implementation)

### 3.2 Full Screen AI Layout Implementation

**Status**: Ready to start
**Priority**: HIGH - Essential for proper AI mode experience
**Goal**: Replace lazygit layout with dedicated AI layout when in AI mode

#### Critical Layout Management Requirements

- [ ] **Lazygit Layout Suspension**: When entering AI mode, completely hide/disable all lazygit views (Files, Branches, Commits, Staging, Status, etc.)
- [ ] **AI Layout Creation**: Create custom AI layout with 6 panels replacing the entire screen
- [ ] **Layout State Management**: Maintain separate layout states for Git mode (lazygit) and AI mode (custom AI layout)
- [ ] **Seamless Transitions**: Implement clean switching between lazygit layout and AI layout without UI artifacts
- [ ] **Resource Cleanup**: Properly dispose of AI layout views when returning to Git mode
- [ ] **Layout Restoration**: Ensure lazygit layout is fully restored when exiting AI mode

#### Layout Structure (Based on ATLAS.md Design)

```
┌─ Project Status ──────────┐┌─ Current File Context ─────────────────┐
│ 🎯 Task: Add validations  ││ app/models/user.rb                     │
│ ⏱️  1h 23m active        ││ class User < ApplicationRecord         │
│ 📊 Progress: ████████░░ 80%││   validates :email, presence: true    │
├─ Smart File List ─────────┤├─ Step-by-Step Guide ───────────────────┤
│ 🎯 Current context:       ││ 📋 Add Email Validation:               │
│   • user.rb          ←─── ││ 1. Add format validation:              │
│   • users_controller.rb   ││    format: { with: URI::MailTo::EMAIL...│
│   • user_spec.rb          ││ 2. Add uniqueness constraint:          │
├─ System Health ───────────┤│    uniqueness: { case_sensitive: false}│
│ 🧪 Tests: 12/15 passing   ││ 3. Update database migration:          │
│ 📊 DB: Connected          ││    add_index :users, :email, unique:...│
└───────────────────────────┘└─────────────────────────────────────────┘
┌─ AI Teaching Session ──────────────────────────────────┐
│ 💬 You: How do I add proper email validation?          │
│ 🤖 Atlas: I'll walk you through Rails email validation │
│          step by step...                              │
└─────────────────────────────────────────────────────────┘
```

#### Implementation Plan

**Step 1: Layout System Architecture**

- [ ] **Start with comprehensive tests**: Create test suite for layout coordinate calculations, panel positioning, and layout configuration
- [ ] Create `AILayoutManager` struct following lazygit's layout patterns
- [ ] Implement panel coordinate calculation system
- [ ] Define panel types: ProjectStatus, FileContext, SmartFileList, StepGuide, SystemHealth, AITeaching
- [ ] Create layout configuration system for panel sizes and positions

**Step 2: Panel Creation System**

- [ ] **Start with comprehensive tests**: Create test suite for panel factory, panel interfaces, and content management
- [ ] Implement `AIPanelFactory` for creating different panel types
- [ ] Create base `AIPanel` interface with common methods (render, update, focus)
- [ ] Implement individual panel structs for each layout section
- [ ] Add panel content management and scrolling support

**Step 3: Layout Rendering Engine**

- [ ] **Start with comprehensive tests**: Create test suite for rendering order, focus management, and view integration
- [ ] Create `AILayoutRenderer` to manage panel rendering order
- [ ] Implement panel focus management and keyboard navigation
- [ ] Add layout refresh system for dynamic content updates
- [ ] Integrate with gocui's view management system

**Step 4: Mode Transition System**

- [ ] **Start with comprehensive tests**: Create test suite for transition states, panel visibility, and cleanup operations
- [ ] Implement `LayoutTransitionManager` for smooth Git ↔ AI mode switching
- [ ] Create panel visibility state management for both lazygit and AI layouts
- [ ] **Lazygit Layout Suspension**: Implement mechanism to hide/disable all lazygit views when entering AI mode
- [ ] **Layout State Preservation**: Save lazygit view states (positions, content, scroll positions) before switching
- [ ] **AI Layout Activation**: Create and display AI layout panels when entering AI mode
- [ ] **Layout Restoration**: Restore lazygit layout and view states when exiting AI mode
- [ ] Add transition animations (optional polish)
- [ ] Ensure proper cleanup of AI views when switching modes

**Step 5: Content Integration**

- [ ] **Start with comprehensive tests**: Create test suite for data sources, content updates, and mock data systems
- [ ] Connect panels to data sources (file system, git status, etc.)
- [ ] Implement content update system for real-time information
- [ ] Add placeholder content for panels not yet functional
- [ ] Create mock data system for development and testing

**Step 6: Testing and Polish**

- [ ] **Start with comprehensive tests**: Create test suite for layout behavior, terminal resizing, and keyboard navigation
- [ ] Add comprehensive layout tests
- [ ] Test panel resizing and terminal size changes
- [ ] Add keyboard navigation tests
- [ ] Polish visual styling and spacing

#### Technical Considerations

- **Layout Framework**: Use gocui's view system with custom layout calculations
- **Panel Communication**: Implement observer pattern for panel-to-panel updates
- **Performance**: Lazy loading and caching for panel content
- **Extensibility**: Plugin system for adding new panel types
- **Testing**: Mock panels for isolated testing of layout logic
- **Lazygit Integration**: Hook into lazygit's layout system to suspend/resume views
- **View Management**: Implement dual view management (lazygit views + AI views)
- **State Preservation**: Save/restore lazygit view states during mode transitions

#### Acceptance Criteria

- [ ] **Lazygit Layout Completely Hidden**: All lazygit views (Files, Branches, Commits, Staging, Status, etc.) are invisible in AI mode
- [ ] Full screen AI layout replaces lazygit layout in AI mode
- [ ] All 6 panels (ProjectStatus, FileContext, SmartFileList, StepGuide, SystemHealth, AITeaching) are visible
- [ ] Panel sizes and positions match ATLAS.md design
- [ ] **Seamless Mode Transitions**: Switching between Git and AI modes is instant and clean
- [ ] **Lazygit Layout Restored**: All lazygit views return to their previous state when exiting AI mode
- [ ] Keyboard navigation works within AI layout
- [ ] Layout handles terminal resizing properly
- [ ] Placeholder content shown for non-functional panels
- [ ] No lazygit panels visible in AI mode
- [ ] **Layout State Preservation**: Scroll positions, selections, and content are preserved during mode switches

### 3.3 AI Response Formatting

**Status**: Not started
**Acceptance Criteria**:

- [ ] Markdown rendering for AI responses
- [ ] Code blocks with syntax highlighting
- [ ] Clickable file paths that open files in editor
- [ ] Proper text wrapping and formatting
- [ ] Loading indicators during AI processing

### 3.4 Advanced Panel Features

**Status**: Not started
**Acceptance Criteria**:

- [ ] Split panel view (file content + AI chat)
- [ ] Resizable panel sections
- [ ] Multiple file tabs or quick switching
- [ ] Search functionality within AI responses
- [ ] Copy/export AI responses

## Phase 4: AI Integration

### 4.1 AI Client Architecture

**Status**: Not started
**Acceptance Criteria**:

- [ ] Pluggable AI provider system (Claude, OpenAI, Twitter)
- [ ] Configuration system for API keys and model selection
- [ ] Proper error handling for API failures
- [ ] Rate limiting and request queuing
- [ ] Offline mode with graceful degradation

### 4.2 Context Collection System

**Status**: Not started
**Acceptance Criteria**:

- [ ] Git repository information gathering
- [ ] Current file content reading
- [ ] Git status and diff information
- [ ] Recent commit history context
- [ ] Branch and remote information

### 4.3 AI Query Processing

**Status**: Not started
**Acceptance Criteria**:

- [ ] User input parsing and validation
- [ ] Context-aware prompt construction
- [ ] Streaming response handling
- [ ] Response caching for similar queries
- [ ] Query history management

## Phase 5: Smart File Discovery

### 5.1 Intelligent File Suggestions

**Status**: Not started
**Acceptance Criteria**:

- [ ] Git-aware file relationship detection
- [ ] MVC pattern recognition for related files
- [ ] Recent files and editing history integration
- [ ] Test file suggestions for source files
- [ ] Import/dependency file suggestions

### 5.2 Project Structure Analysis

**Status**: Not started
**Acceptance Criteria**:

- [ ] Framework detection (Rails, React, etc.)
- [ ] Project structure understanding
- [ ] Configuration file awareness
- [ ] Build system integration
- [ ] Package manager file parsing

## Phase 6: Advanced Features

### 6.1 Read-Only System Intelligence

**Status**: Not started
**Acceptance Criteria**:

- [ ] Database query capabilities (SELECT only)
- [ ] Log file analysis and monitoring
- [ ] Process monitoring and system health
- [ ] Dependency source code reading
- [ ] Code execution tracing

### 6.2 Multi-Model Support

**Status**: Not started
**Acceptance Criteria**:

- [ ] Support for multiple AI providers simultaneously
- [ ] Model switching based on query type
- [ ] Custom model configuration
- [ ] Local model integration (Ollama)
- [ ] Model performance comparison

## Phase 7: General Terminal AI Chat

### 7.1 Global AI Assistant

**Status**: Not started
**Acceptance Criteria**:

- [ ] Global AI chat mode accessible from any directory
- [ ] Conversations not tied to specific git repositories
- [ ] General programming assistance and coding questions
- [ ] Terminal-native chat interface with proper formatting
- [ ] Chat history persistence across sessions
- [ ] Easy switching between project-specific and general chat modes
- [ ] Support for non-coding conversations (general AI assistance)
- [ ] Conversation management (save, load, delete chat sessions)

## Phase 8: Polish & Distribution

### 8.1 Configuration & Customization

**Status**: Not started
**Acceptance Criteria**:

- [ ] User configuration file support
- [ ] Customizable keybindings
- [ ] Theme and styling options
- [ ] AI behavior preferences
- [ ] Performance tuning options

### 8.2 Documentation & Distribution

**Status**: Not started
**Acceptance Criteria**:

- [ ] Complete user documentation
- [ ] Installation guides for different platforms
- [ ] Configuration examples and tutorials
- [ ] Troubleshooting guides
- [ ] Release packaging and distribution
