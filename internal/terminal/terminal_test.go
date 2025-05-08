package terminal

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveDimensions_HostNotTerminal_YieldDefault(t *testing.T) {
	// GIVEN
	oldIsTerminal := termIsTerminal
	defer func() { termIsTerminal = oldIsTerminal }()

	termIsTerminal = func(fd uintptr) bool { return false }

	// WHEN
	accessor := NewTerminalAccessor()
	width, height := accessor.ResolveDimensions()

	// THEN
	assert.Equal(t, -1, width)
	assert.Equal(t, -1, height)
}

func TestResolveDimensions_HostTerminalButYieldsErr_Panics(t *testing.T) {
	// GIVEN
	oldIsTerminal := termIsTerminal
	oldGetSize := termGetSize
	defer func() {
		termIsTerminal = oldIsTerminal
		termGetSize = oldGetSize
	}()

	termIsTerminal = func(fd uintptr) bool { return true }
	termGetSize = func(fd uintptr) (width int, height int, err error) {
		return -1, -1, errors.New("some internal error")
	}

	// WHEN & THEN
	accessor := NewTerminalAccessor()
	assert.Panics(t, func() { accessor.ResolveDimensions() })
}

func TestResolveDimensions_HostTerminal_PassDimensionsWithHeightCorrection(t *testing.T) {
	// GIVEN
	oldIsTerminal := termIsTerminal
	oldGetSize := termGetSize
	defer func() {
		termIsTerminal = oldIsTerminal
		termGetSize = oldGetSize
	}()

	expectedWidht := 50
	expectedHeight := 95
	termIsTerminal = func(fd uintptr) bool { return true }
	termGetSize = func(fd uintptr) (width int, height int, err error) {
		return expectedWidht, expectedHeight + stdHeightDeduction, nil
	}

	// WHEN
	accessor := NewTerminalAccessor()
	width, height := accessor.ResolveDimensions()

	// THEN
	assert.Equal(t, expectedWidht, width)
	assert.Equal(t, expectedHeight, height)

}
