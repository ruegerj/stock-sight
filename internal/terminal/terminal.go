package terminal

import (
	"github.com/charmbracelet/x/term"
)

// due to some terminal configs displaying the current state and command input on multiple lines
const stdHeightDeduction int = 5

type TerminalAccessor struct{}

func NewTerminalAccesor() TerminalAccessor {
	return TerminalAccessor{}
}

func (hp TerminalAccessor) ResolveDimensions() (widht int, height int) {
	if !term.IsTerminal(0) {
		return -1, -1
	}

	widht, height, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	return widht, height - stdHeightDeduction
}
