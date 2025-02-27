package cmd

import "testing"

// TODO Delete me
func TestDummyOperation(t *testing.T) {
	expected := 2

	res := 1 + 1

	if res != expected {
		t.Errorf("failed to compute: expected=%d, got=%d", expected, res)
	}
}
