package terminal

import (
	"github.com/charmbracelet/x/term"
)

// due to some terminal configs displaying the current state and command input on multiple lines
const stdHeightDeduction int = 5

// allow monkey-patching in tests
var termIsTerminal = term.IsTerminal
var termGetSize = term.GetSize

type TerminalAccessor interface {
	ResolveDimensions() (width, height int)
}

type ConcreteTerminaAccessor struct{}

func NewTerminalAccessor() TerminalAccessor {
	return ConcreteTerminaAccessor{}
}

func (ta ConcreteTerminaAccessor) ResolveDimensions() (width int, height int) {
	if !termIsTerminal(0) {
		return -1, -1
	}

	width, height, err := termGetSize(0)
	if err != nil {
		panic(err)
	}

	return width, height - stdHeightDeduction
}
