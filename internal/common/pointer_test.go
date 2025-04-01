package common

import "testing"

func TestDerefOrEmpty_ValueString(t *testing.T) {
	var value = "42"

	res := DerefOrEmpty(&value)

	if res != value {
		t.Errorf("Mismatch between origin and dereferenced values. expected=%q, got=%q", value, res)
	}
}

func TestDerefOrEmpty_EmptyString(t *testing.T) {
	var value *string = nil

	res := DerefOrEmpty(value)

	if res != "" {
		t.Errorf("Expected empty string for nil-pointer but got=%q", res)
	}
}
