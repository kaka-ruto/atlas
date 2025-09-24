# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Atlas is a Git TUI (Terminal User Interface) that extends Lazygit with AI capabilities for pair programming. It provides two modes:
- **Git Mode**: Standard Lazygit functionality for Git operations
- **AI Mode**: AI-powered pair programming

The project is built in Go and uses the same architecture as Lazygit, with additional AI mode functionality in the `pkg/atlas` package.

**IMPORTANT**: When working with this codebase, always read `ATLAS_IMPLEMENTATION.md` first to understand the current implementation details, design decisions, and specific Atlas features.

## Development Commands

### Building
```bash
make build          # Build the atlas binary
make install        # Install to $GOPATH/bin
go build            # Alternative build command
```

### Running
```bash
make run            # Build and run atlas
./atlas             # Run the built binary directly
go run main.go      # Run directly with Go
```

### Testing
```bash
make unit-test      # Run unit tests only
make test           # Run both unit and integration tests
go test ./... -short  # Alternative unit test command
```

### Development/Debugging
```bash
make run-debug      # Run in debug mode (use in one terminal)
make print-log      # Print logs (use in separate terminal)
go run main.go -debug    # Alternative debug command
go run main.go --logs    # Alternative log viewing
```

### Code Quality
```bash
make lint           # Run golangci-lint via script
make format         # Format code with gofumpt
make generate       # Generate auto-generated files
```

### Integration Testing
```bash
make integration-test-tui  # Run TUI integration tests
make integration-test-cli  # Run CLI integration tests
make integration-test-all  # Run all integration tests
```

## Architecture Overview

Atlas is built as a wrapper around Lazygit with the following key components:

### Core Architecture
- `main.go`: Entry point that starts the app with build info
- `pkg/app/`: Application initialization and startup logic
- `pkg/atlas/`: Atlas-specific functionality including mode switching
- `pkg/gui/`: GUI components and panels (inherited from Lazygit)
- `pkg/commands/`: Git command implementations
- `pkg/config/`: Configuration management

### Atlas-Specific Components
- `AtlasWrapper`: Main wrapper around Lazygit GUI that manages modes
- `Mode`: Enum defining GitMode and AIMode
- `AIModeState`: State management for AI-specific functionality

### Key Packages
- `pkg/gui/`: Contains the main GUI logic and panels
- `pkg/commands/`: Git command abstractions and implementations
- `pkg/app/`: Application entry point and initialization
- `pkg/atlas/`: Atlas mode switching and AI integration
- `pkg/integration/`: Integration test framework

## Mode System

Atlas operates in two distinct modes:
1. **Git Mode** (`GitMode`): Standard Lazygit functionality
2. **AI Mode** (`AIMode`): AI-powered pair programming

## Development Workflow

1. This project follows the same patterns as Lazygit
2. Use `make run-debug` and `make print-log` for development debugging
3. Code formatting uses `gofumpt` (stricter than `gofmt`)
4. Linting is enforced via golangci-lint with custom configuration
5. Tests must be added for new functionality
6. Follow Go idioms with some project-specific preferences

## Configuration Files

- `.golangci.yml`: Linter configuration with Atlas-specific rules
- `Makefile`: Build targets and development commands
- `go.mod`: Module dependencies (based on Lazygit dependencies plus Atlas additions)

## Testing Strategy

- Unit tests in `*_test.go` files
- Integration tests in `pkg/integration/`
- Debug with logs by setting `LOG_LEVEL` environment variable
- Use `LAZYGIT_LOG_PATH` for vendor directory debugging

## Important Notes

- Atlas inherits most functionality from Lazygit
- The main extension point is the mode system in `pkg/atlas/`
- Debug logging is available via the same mechanisms as Lazygit
- Build process may set version info via LDFLAGS