# Atlas Implementation Guide

This document provides comprehensive guidelines for AI assistants developing Atlas features. **Follow these principles strictly** to maintain code quality and consistency with the lazygit codebase.

## Core Development Principles

### 1. Test-Driven Development (TDD)

**ALWAYS follow this sequence**:

1. **Write tests first** - Never write implementation code before tests
2. **Run tests** - Ensure they fail (red)
3. **Write minimal implementation** - Make tests pass (green)
4. **Refactor** - Improve code while keeping tests green
5. **Repeat** - One feature at a time

### 2. Study Existing Patterns First

**Before implementing ANY new feature**:

1. **Search lazygit codebase** for similar functionality
2. **Read existing implementations** in `pkg/gui/`, `pkg/commands/`, etc.
3. **Follow identical patterns** for:
   - File structure and naming
   - Function signatures and error handling
   - Test organization and scenarios
   - Import statements and dependencies

**Commands to use**:

- `find pkg/ -name "*.go" | grep <feature>`
- `grep -r "similar_functionality" pkg/`
- Read `pkg/gui/keybindings.go` before adding keybindings
- Read `pkg/gui/controllers/` for UI interaction patterns
- And more

### 3. Reuse Lazygit Functionality

**DO NOT reinvent the wheel**:

- **Keybinding system**: Use existing `types.Binding` and keybinding registration
- **UI panels**: Follow `gocui.View` and context patterns
- **Commands**: Reuse `oscommands` and `git_commands` where possible
- **Configuration**: Use existing config system patterns
- **Error handling**: Follow lazygit's error patterns exactly

### 4. File Organization Rules

**Package Structure**:

- Keep Atlas code in `pkg/atlas/` to avoid polluting lazygit core
- Follow Go package naming conventions
- Split functionality logically (mode.go, keybindings.go, ui.go, etc.)
- Mirror lazygit's file organization patterns

**Testing Structure**:

- One `*_test.go` file per feature file
- Use scenario-based testing like lazygit
- Group related tests in the same file
- Follow existing test naming conventions

### 5. Integration Guidelines

#### Keybinding Integration

**MUST follow lazygit's keybinding system**:

1. Study `pkg/gui/keybindings.go` structure
2. Use existing `types.Binding` struct
3. Register through proper keybinding handlers
4. Respect keybinding configuration system
5. Add to help/cheatsheet system

#### UI Integration

**MUST follow lazygit's UI patterns**:

1. Study `pkg/gui/layout.go` for panel management
2. Use existing view and context system
3. Follow theme and styling conventions
4. Handle terminal resizing properly
5. Maintain consistent user experience

#### Command Integration

**MUST reuse existing command infrastructure**:

1. Use `oscommands` for system operations
2. Use `git_commands` for git operations
3. Follow existing error handling patterns
4. Respect user configuration preferences
5. Maintain command logging and debugging

### 6. Code Quality Standards

#### Error Handling

```go
// GOOD - Follow lazygit patterns
func (aw *AtlasWrapper) SomeOperation() error {
    if err := someOperation(); err != nil {
        return err
    }
    return nil
}

// BAD - Don't ignore errors or use panic
func (aw *AtlasWrapper) SomeOperation() {
    someOperation() // No error handling
}
```

#### Function Signatures

**Follow lazygit conventions exactly**:

- Use consistent parameter naming
- Return `error` as last return value when needed
- Use receiver naming conventions (`self *Gui` or similar)
- Follow Go naming conventions (exported vs unexported)

#### Comments and Documentation

**Match lazygit's comment style**:

- Document exported functions and types
- Use Go doc comment conventions
- No excessive commenting for obvious code
- Focus on why, not what

### 7. Testing Requirements

#### Test Coverage

**MUST achieve comprehensive coverage**:

- Test all public methods
- Test error conditions
- Test edge cases and boundary conditions
- Test concurrent access if applicable

#### Test Quality

**Follow lazygit's testing patterns**:

- Use scenario-based testing structure
- Descriptive test names explaining the scenario
- Clear arrange-act-assert patterns
- Mock external dependencies appropriately

### 8. Development Workflow

#### One Feature at a Time

**NEVER mix multiple features**:

1. Complete one roadmap item fully before starting next
2. Ensure all tests pass before moving forward
3. Update changelog when feature is complete
4. Get feature working before adding enhancements

#### Code Review Checklist

**Before submitting any implementation**:

- [ ] All tests pass (`go test ./pkg/atlas/`)
- [ ] Code follows existing lazygit patterns
- [ ] No new external dependencies without justification
- [ ] Error handling follows conventions
- [ ] Documentation is complete and accurate
- [ ] Performance considerations addressed

#### Build and Integration

**Always verify**:

- [ ] Project builds without errors (`go build`)
- [ ] No breaking changes to existing lazygit functionality
- [ ] Integration with lazygit keybinding system works
- [ ] UI integration maintains lazygit look and feel

### 9. Reference Resources

#### Atlas Project Documents

**MUST READ these Atlas-specific documents before starting**:

- `ATLAS.md` - **Project vision and architecture overview** - Understand what Atlas is, its philosophy, and overall technical architecture
- `ATLAS_CHANGELOG.md` - **Historical implementation record** - Review completed features, key decisions, and architecture choices made
- `ATLAS_ROADMAP.md` - **Feature planning and priorities** - Understand what's planned, current progress, and acceptance criteria for features
- `ATLAS_IMPLEMENTATION.md` - **This document** - Development guidelines and patterns to follow

#### Essential Lazygit Files to Study

**Before implementing any feature**:

- `CONTRIBUTING.md` and linked files - Development guidelines
- `pkg/gui/gui.go` - Main GUI structure
- `pkg/gui/keybindings.go` - Keybinding patterns
- `pkg/gui/layout.go` - UI layout patterns
- `pkg/gui/controllers/` - Controller patterns
- `pkg/commands/` - Command execution patterns

#### Integration Points

**Key lazygit systems to understand**:

- Context system (`pkg/gui/context/`)
- View management (`pkg/gui/types/`)
- Command execution (`pkg/commands/`)
- Configuration (`pkg/config/`)
- Internationalization (`pkg/i18n/`)

### 10. Debugging and Troubleshooting

#### Development Debugging

**Use lazygit's debugging tools**:

- Run with `--debug` flag for logging
- Use `--logs` flag to view logs
- Follow existing logging patterns
- Use structured logging where available

#### Common Pitfalls to Avoid

- **Don't** create custom keybinding systems
- **Don't** bypass lazygit's UI management
- **Don't** ignore existing configuration systems
- **Don't** break lazygit's existing functionality
- **Don't** skip writing tests

## Summary

**The golden rule**: When in doubt, find how lazygit does it and follow that pattern exactly. Atlas should feel like a natural extension of lazygit, not a separate application grafted onto it.
