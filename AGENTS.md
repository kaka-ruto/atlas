# Agent Guidelines for Atlas Repository

## Build/Lint/Test Commands

### Building

- `make build` - Build the atlas binary with debug symbols
- `make install` - Install to $GOPATH/bin
- `go build -gcflags='all=-N -l'` - Alternative build command

### Running

- `make run` - Build and run atlas
- `./atlas` - Run the built binary directly
- `make run-debug` - Run in debug mode (use with `make print-log` in separate terminal)

### Testing

- `make unit-test` - Run unit tests only (`go test ./... -short`)
- `make test` - Run both unit and integration tests
- `go test -run TestName ./path/to/package` - Run a single test
- `go test -v ./path/to/package` - Run tests with verbose output

### Code Quality

- `make lint` - Run golangci-lint via shim script
- `make format` - Format code with gofumpt (stricter than gofmt)

## Code Style Guidelines

### Formatting

- Use `gofumpt` for formatting (via `make format`)
- Tab indentation (configured in .editorconfig)
- Follow standard Go formatting conventions

### Linting

- Enabled linters: copyloopvar, errorlint, exhaustive, intrange, makezero, nakedret, nolintlint, prealloc, revive, thelper, tparallel, unconvert, unparam, wastedassign
- No naked returns allowed (max-func-lines: 0)
- Use `self` as receiver name (allowed despite revive warnings)
- Error strings should not be capitalized (ST1005 disabled)

### Imports

- Standard library imports first
- Third-party imports second
- Local/project imports last
- Use blank lines to separate import groups

### Naming Conventions

- Functions/variables: camelCase
- Exported identifiers: PascalCase
- Constants: ALL_CAPS with underscores
- Follow Go naming conventions (short, descriptive names)

### File Structure

- Strictly follow existing lazygit directory and file structure
- All atlat code should live in the atlas folder, outside of existing lazygit code
- Lazygit should still work as is
- Avoid merge conflicts with lazygit when we shall update lazygit

### Error Handling

- Return errors as last return value
- Use standard Go error patterns
- Check errors immediately after function calls
- Use error wrapping when appropriate

### Comments

- File headers: Include purpose and template references
- No unnecessary inline comments
- Export comments follow Go conventions

### Testing

- Use `testify/assert` for assertions
- Table-driven tests preferred
- Test files end with `_test.go`
- Use descriptive test names

### Types and Structs

- Use meaningful type names
- Embed interfaces when appropriate
- Follow Go idioms for struct composition

### Documentation

- Read ATLAS.md and other ATLAS markdown files
- Feel free to change any of those ATLAS md files as they may contain outdated documentation/roadmap
