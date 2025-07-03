# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

typingo is a CLI typing game implemented in Go. It uses the Bubble Tea framework to build a terminal UI and measures WPM (Words Per Minute) by typing randomly generated texts.

## Development Commands

### Build and Test
```bash
# Build
go build

# Run
./typingo

# Download dependencies
go mod download

# Tidy dependencies
go mod tidy

# Run tests (no test files currently exist)
go test ./...
```

### Lint
```bash
# Install golangci-lint (first time only)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run lint
golangci-lint run
```

### Release
```bash
# Install GoReleaser (first time only)
go install github.com/goreleaser/goreleaser@latest

# Test release locally (doesn't actually publish)
goreleaser release --snapshot --clean

# Actual release (automatically triggered when pushing a tag)
git tag v0.1.0
git push origin v0.1.0
```

## Architecture

### Directory Structure
- `cmd/` - CLI command definitions (using Cobra)
- `internal/game/` - Game logic (Bubble Tea model)
- `internal/texts/` - Text generation logic
- `main.go` - Entry point

### Bubble Tea MVU Pattern

This project follows the Bubble Tea Model-View-Update pattern:

1. **Model** (`internal/game/game.go`): Game struct that manages game state
2. **Update**: Updates state based on key inputs and timer events
3. **View**: Switches between 4 views (start, countdown, text, result) based on current state

### Key Message Types
- `tea.KeyMsg`: Key input handling
- `countdownMsg`: For countdown
- `tickMsg`: Timer updates (millisecond precision)

### Key Bindings
- Space: Start game
- r: Retry
- q: Quit
- Ctrl+C, Esc: Cancel

### State Transitions
1. Launch → Start screen
2. Press space → Countdown (3-2-1)
3. Countdown ends → Typing screen
4. All texts completed → Result screen
5. Press 'r' → Reset state → Back to start screen

## Development Notes

- CGO is disabled (for cross-platform builds)
- Version info is embedded in `internal/game.version` via LDFLAGS
- UI rendering uses lipgloss (main color: #00ADD8)
- Supports window resizing (handles WindowSizeMsg)