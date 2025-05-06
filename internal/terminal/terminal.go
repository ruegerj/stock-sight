package terminal

import (
	"github.com/charmbracelet/x/term"
)

// due to some terminal configs displaying the current state and command input on multiple lines
const stdHeightDeduction int = 5

// allow monkey-patching in tests
var termIsTerminal = term.IsTerminal
var termGetSize = term.GetSize

type TerminalAccessor struct{}

func NewTerminalAccessor() TerminalAccessor {
	return TerminalAccessor{}
}

func (hp TerminalAccessor) ResolveDimensions() (widht int, height int) {
	if !termIsTerminal(0) {
		return -1, -1
	}

	widht, height, err := termGetSize(0)
	if err != nil {
		panic(err)
	}

	return widht, height - stdHeightDeduction
}
