package main

import "testing"

func TestAddSuccess(t *testing.T) {
	result := Add(20, 2)
	expected := uint64(23)

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
